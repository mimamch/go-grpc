package main

import (
	"context"
	"fmt"
	"log"

	"github.com/mimamch/gorpc/proto/hello"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func main() {
	dial, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed Connecting %v", err)
	}

	c := hello.NewHelloServiceClient(dial)
	sendHello(c, "mimamch")
}

func sendHello(conn hello.HelloServiceClient, data string) {
	response, err := conn.SayHello(context.Background(), &hello.HelloRequest{
		Name: data,
	})
	if err != nil {
		if errStatus, _ := status.FromError(err); errStatus != nil && errStatus.Code() == 400 {
			fmt.Printf("Invalid Argument %v", errStatus.Message())
			return
		}
		fmt.Printf("Failed Calling SayHello %v", err)
		return
	}
	fmt.Printf("Response from Server: %s", response.Message)
}
