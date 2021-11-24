package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "sm2/api"
	"sm2/internal/config"
	"sm2/internal/server"
)

func main() {
	configs := config.GetConfig()
	listen, err := net.Listen("tcp", configs.GRpc.Addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCmbSm2Server(s, &server.Sm2GrpcServer{})
	fmt.Println("gRpc server listen on ", configs.GRpc.Addr)
	err = s.Serve(listen)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
