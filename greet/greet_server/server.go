package main

import (
	"GooleGrpc/greet/greetpb"
	"fmt"
	"github.com/micro/util/go/lib/net"
	"google.golang.org/grpc"
	"log"
)

type server struct{}

func main() {
	fmt.Println("Hello world")

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to serve %v", err)
	}
}
