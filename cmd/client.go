package main

import (
	"context"
	"github.com/DuC-cnZj/goods_proto"
	order "github.com/DuC-cnZj/micro-order/protos"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"log"
)

func main() {
	dial, err := grpc.Dial("localhost:6666", grpc.WithInsecure())
	goodsDial, err := grpc.Dial("localhost:9999", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc dial error %v", err)
	}

	client := order.NewOrderClient(dial)
	goodsClient := goods_proto.NewBBQClient(goodsDial)
	list, err := client.List(context.Background(), &empty.Empty{})
	if err != nil {
		log.Fatalf("err call order svc list: %v", err)
	}
	eat, _ := goodsClient.Eat(context.Background(), &goods_proto.Request{
		Name: "from order call",
	})
	log.Println(list, eat)
}
