package main

import (
	"fmt"

	"GooleGrpc/calculator/calculatorpb"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct{}

const (
	port = ":50052"
)

func main() {

	fmt.Println("Calculator Server")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Failed to listen %v", err)
		s := grpc.NewServer()

		calculatorpb.RegisterCalculatorServiceServer(s, &server{})
		if err := s.Serve(lis); err != nil {
			log.Fatal("Failed to serve %v", err)
		}
	}
}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {

	fmt.Printf("Received sum %v", req)
	firstNumber := req.GetFirstNumber()
	secondNumber := req.GetSecondNumber()
	sum := firstNumber + secondNumber
	res := &calculatorpb.SumResponse{
		SumResult: sum,
	}
	return res, nil
}
