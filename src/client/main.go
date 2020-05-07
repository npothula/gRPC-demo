package main

import (
	"context"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"

	pb "gRPCHelloWorld/helloworld"
)

const (
	address    = "localhost:50051"
	defaultMsg = "world"
)

func sendMsg() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		os.Exit(1)
	}
	defer conn.Close()
	hwClient := pb.NewHelloWorldClient(conn)

	// Contact the server and print out its response.
	name := defaultMsg
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	response, err := hwClient.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	} else {
		log.Printf("Greeting: %s", response.GetMessage())
	}
}

func main() {
	sendMsg()
}
