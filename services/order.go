package services

import (
	context "context"
	"github.com/DuC-cnZj/micro-order/protos"
	"github.com/golang/protobuf/ptypes/empty"
	"log"
)

type OrderSvc struct {
	order.UnimplementedOrderServer
}

//https://github.com/grpc/grpc-go/issues/3794
func (o *OrderSvc) List(ctx context.Context, empty *empty.Empty) (*order.Response, error) {
	log.Println("order svc: list().")

	return &order.Response{
		Code: 200,
		Data: "success from order svc List()",
	}, nil
}

func (o *OrderSvc) Create(ctx context.Context, request *order.Request) (*order.Response, error) {
	log.Println("order svc: create().")

	return &order.Response{
		Code: 200,
		Data: "success from order svc Create(), data from request: " + request.UserName + request.Num,
	}, nil
}

