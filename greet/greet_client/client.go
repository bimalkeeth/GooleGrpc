package main

import (
	"GooleGrpc/greet/greetpb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {
	fmt.Println("Hello Client")
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Coud not connect", err)
	}
	defer conn.Close()
	c := greetpb.NewGreetServiceClient(conn)

	res, err := c.Greet(context.Background(), &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			First_Name: "Bimal",
			Last_Name:  "Kaluarachchi",
		},
	})
	if err != nil {
		log.Fatal("Error occurred")
	}
	fmt.Println(res.GetResult())

	resStream, err := c.GreatManyTimes(context.Background(), &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			First_Name: "Bimal",
			Last_Name:  "Kaluarachchi",
		},
	})
	if err != nil {
		log.Fatal("Error occurred")
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("Error occurred")
		}
		fmt.Println(msg.GetResult())
	}

}
