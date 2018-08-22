package main

import (
	"GooleGrpc/calculator/calculatorpb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("Hello Client")
	conn, err := grpc.Dial(":50052", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Coud not connect", err)
	}
	defer conn.Close()
	c := calculatorpb.NewCalculatorServiceClient(conn)

	req := &calculatorpb.SumRequest{
		FirstNumber:  10,
		SecondNumber: 20,
	}

	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatal("Error occurred")
	}
	fmt.Println(res.SumResult)
}
