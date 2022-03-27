package main

import (
	pb "RPC/hello/proto/hello"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	Address = "localhost:50052"
)

type HelloService struct {
	pb.HelloServer
}

var helloService = HelloService{}

func (h HelloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResopnse, error) {
	hr := new(pb.HelloResopnse)
	hr.Message = fmt.Sprintf("Hello %s", in)

	return hr, nil
}
func (h HelloService) mustEmbedUnimplementedHelloServer() {}
func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterHelloServer(s, &HelloService{})

	log.Println("Listen on " + Address)
	s.Serve(listen)
}
