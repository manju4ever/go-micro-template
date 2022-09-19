package main

/*
	This is where Gin / Fiber API routes will be created to call other microservice !
*/

import (
	"context"
	"fmt"

	Schema "todoservice/proto"

	natsbroker "github.com/go-micro/plugins/v4/broker/nats"
	natsregistry "github.co`m/go-micro/plugins/v4/registry/nats"
	natsport "github.com/go-micro/plugins/v4/transport/nats"
	micro "go-micro.dev/v4"
)

func main() {
	service := micro.NewService(
		micro.Registry(natsregistry.NewRegistry()),
		micro.Broker(natsbroker.NewBroker()),
		micro.Transport(natsport.NewTransport()),
	)
	service.Init()
	client := service.Client()
	req := client.NewRequest("todo-service", "TodoService.GetAllTodos", &Schema.Void{})
	res := &Schema.AllTodoItems{}
	ctx := context.Background()
	if err := client.Call(ctx, req, res); err != nil {
		fmt.Println("[Error] ", err, res)
		return
	}

	for _, item := range res.Items {
		fmt.Println(item)
	}

}
