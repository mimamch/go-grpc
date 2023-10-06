package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/mimamch/gorpc/proto/hello"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

type HelloService struct {
	hello.UnimplementedHelloServiceServer
}

func (h *HelloService) SayHello(ctx context.Context, req *hello.HelloRequest) (*hello.HelloResponse, error) {
	fmt.Println("Received request " + req.Name)
	if len(req.Name) < 3 {
		return nil, status.Error(400, "name is too short")
	}
	return &hello.HelloResponse{
		Message: fmt.Sprintf("Hello %s", req.Name),
	}, nil
}

func main() {

	port := flag.Int("port", 50051, "the server port")
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()

	hello.RegisterHelloServiceServer(server, &HelloService{})
	log.Printf("Starting server on port %d", *port)
	server.Serve(lis)
}
