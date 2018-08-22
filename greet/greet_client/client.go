package main

import (
	"GooleGrpc/greet/greetpb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
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
	doBiDirectionalStream(c)
}

func doBiDirectionalStream(c greetpb.GreetServiceClient) {
	fmt.Println("Bidirectional client server streaming")
	stream, err := c.GreetEveryOne(context.Background())
	if err != nil {
		log.Fatal("Error occurred in streaming  %v", err)
	}

	requests := []*greetpb.GreetEveryOneRequest{
		{
			Greeting: &greetpb.Greeting{
				First_Name: "Stephane",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				First_Name: "July",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				First_Name: "Tom",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				First_Name: "Robert",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				First_Name: "Lucy",
			},
		},
	}

	waitc := make(chan struct{})
	go func() {
		for _, req := range requests {
			fmt.Println("Sending Message %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	go func() {

		for {
			res, err := stream.Recv()

			if err == io.EOF {
				close(waitc)
				break
			}
			if err != nil {
				log.Fatal("Error occurred in streaming rec %v", err)
				break
			}
			fmt.Printf("Received %v", res.GetResult())
		}
	}()

	<-waitc
}
