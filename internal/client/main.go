package main

import (
	"context"
	"log"
	"time"

	"github.com/hoangminhphuc/goph-shop/api"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := api.NewHelloServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resp, err := client.SayHello(ctx, &api.HelloRequest{Name: "Phuc"})
	if err != nil {
		log.Fatalf("Error calling SayHello: %v", err)
	}

	log.Printf("Response from server: %s", resp.Message)
}
