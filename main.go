package main

import (
	order "github.com/DuC-cnZj/micro-order/protos"
	"github.com/DuC-cnZj/micro-order/services"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":6666")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	order.RegisterOrderServer(server, &services.OrderSvc{})

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
