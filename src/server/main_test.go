package main

import (
	"context"
	"reflect"
	"testing"
	"time"

	pb "gRPCHelloWorld/helloworld"
)

func Test_server_SayHello(t *testing.T) {
	tests := []struct {
		name    string
		s       *server
		in      *pb.HelloRequest
		want    *pb.HelloReply
		wantErr bool
	}{
		{
			//success
			name:    "world",
			s:       &server{},
			in:      &pb.HelloRequest{Name: "world"},
			want:    &pb.HelloReply{Message: "Hello world"},
			wantErr: false,
		},
		/*	{
			//test failed due to wrong expectation
			name:    "naga",
			s:       &server{},
			in:      &pb.HelloRequest{Name: "naga"},
			want:    &pb.HelloReply{Message: "Hello world"},
			wantErr: false,
		}, */
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &server{}
			ctx, _ := context.WithTimeout(context.Background(), time.Second)
			got, err := s.SayHello(ctx, tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("server.SayHello() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("server.SayHello() = %v, want %v", got, tt.want)
			}
		})
	}
}
