package main

import (
	pb "RPC/Demo03/service"
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/alts"
)

func main(){
	tc := alts.NewClientCreds(alts.DefaultClientOptions())
	cc, err := grpc.Dial("localhost:8001",grpc.WithTransportCredentials(tc))
	if err!=nil{
		log.Fatalf("fail to dial: %v",err)
	}

	defer cc.Close()
	client := pb.NewRouteGuideClient(cc)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	f, err2 := client.GetFeature(ctx, &pb.Point{Latitude: 1, Longitude: 1})
	if err2!=nil{
		log.Fatalf("fail to GetFeature: %v",err)
	}
	fmt.Printf("f.String(): %v\n", f.String())
}