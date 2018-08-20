package main

import (
	"GooleGrpc/greet/greetpb"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"

	"context"
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
