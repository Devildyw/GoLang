package main

import (
	pb "RPC/hello/proto/hello"
	"context"
	"log"

	"google.golang.org/grpc"
)


const(
	Address = "localhost:50052"
)

func main(){
	cc, err := grpc.Dial(Address, grpc.WithInsecure())
	if err!=nil{
		log.Fatalln(err)
	}
	defer cc.Close()

	hc := pb.NewHelloClient(cc)

	req:=&pb.HelloRequest{Name: "gRPC",};
	res, err2 := hc.SayHello(context.Background(),req)
	if err2!=nil{
		log.Fatalln(err2)
	}
	log.Println(res.Message)

}