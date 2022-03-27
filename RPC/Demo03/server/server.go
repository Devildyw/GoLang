package main

import (
	"log"
	"net"

	pb "RPC/Demo03/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/alts"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:8001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	altsTC := alts.NewServerCreds(alts.DefaultServerOptions())
	server := grpc.NewServer(grpc.Creds(altsTC))
	pb.RegisterRouteGuideServer(server, &pb.UnimplementedRouteGuideServer{})
	server.Serve(lis)
}
