package main

import (
	"context"
	"log"
	"net"

	"github.com/hoangminhphuc/goph-shop/api"

	"google.golang.org/grpc"
)

type server struct {
	api.UnimplementedHelloServiceServer
}

func (s *server) SayHello(ctx context.Context, req *api.HelloRequest) (*api.HelloResponse, error) {
	log.Printf("Received: %s", req.Name)
	return &api.HelloResponse{Message: "Hello, " + req.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	api.RegisterHelloServiceServer(s, &server{})

	log.Println("gRPC server listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
