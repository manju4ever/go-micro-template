package main

/*
	This is where Gin / Fiber API routes will be created to call other microservice !
*/

import (
	"context"
	"fmt"
	TaskProto "taskservice/proto"

	natsbroker "github.com/go-micro/plugins/v4/broker/nats"
	natsregistry "github.com/go-micro/plugins/v4/registry/nats"
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
	taskBroker := TaskProto.NewTaskService("task-service", service.Client())
	res, err := taskBroker.CreateTodo(context.Background(), &TaskProto.TodoItem{Text: "Read books on spirituality", Color: "Yellow"})
	if err != nil {
		fmt.Println("Client Call to CreateTodo Failed", err)
	}
	fmt.Println(res.MsgType)

	allTasks, err := taskBroker.GetAllTodos(context.Background(), &TaskProto.Void{})
	if err != nil {
		fmt.Println("Client Call to GetAllTodos Failed", err)
	}
	for idx, eachItem := range allTasks.Items {
		fmt.Println("Item:", idx+1, " -- ", eachItem)
	}

}
