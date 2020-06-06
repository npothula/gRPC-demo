package main

import (
	"context"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	pb "gRPCHelloWorld/helloworld"
)

const (
	port = ":50051"
)

// server is used to implement HelloWorldServer.
type server struct {
	//pb.UnimplementedHelloWorldServer
}

// SayHello implements helloworld.HelloWorldServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

// Run ...
func (s *server) Run() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		os.Exit(1)
	}
	log.Printf("Listenging at port %s", port)

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterHelloWorldServer(grpcServer, &server{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	srvr := server{}
	srvr.Run()
}
