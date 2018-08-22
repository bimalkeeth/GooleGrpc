package main

import (
	"GooleGrpc/greet/greetpb"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"

	"context"
	"strconv"
	"time"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {

	fmt.Println("Greet function invoked")

	firstName := req.GetGreeting().GetFirst_Name()
	lastName := req.GetGreeting().GetLast_Name()
	result := "Hello" + firstName + " " + lastName

	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res, nil
}

func (*server) GreatManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreatManyTimesServer) error {

	fmt.Println("Greet Many Times  function invoked")
	firstName := req.GetGreeting().GetFirst_Name()
	lastName := req.GetGreeting().GetLast_Name()

	for i := 0; i < 100; i++ {
		result := "Hello" + firstName + " " + lastName + "" + strconv.Itoa(i)
		res := &greetpb.GreetManyTimesResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(10 * time.Millisecond)
	}
	return nil
}

const (
	port = ":50051"
)

func main() {
	fmt.Println("Hello world")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to serve %v", err)
	}
}
