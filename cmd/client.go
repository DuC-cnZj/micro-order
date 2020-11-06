package main

import (
	"context"
	order "github.com/DuC-cnZj/micro-order/protos"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"log"
)

func main() {
	dial, err := grpc.Dial("localhost:6666", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc dial error %v", err)
	}

	client := order.NewOrderClient(dial)
	list, err := client.List(context.Background(), &empty.Empty{})
	if err != nil {
		log.Fatalf("err call order svc list: %v", err)
	}
	log.Println(list)
}
