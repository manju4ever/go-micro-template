package main

import (
	"context"
	"fmt"
	"time"

	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"

	natsb "github.com/go-micro/plugins/v4/broker/nats"
	natsr "github.com/go-micro/plugins/v4/registry/nats"
	natst "github.com/go-micro/plugins/v4/transport/nats"

	TodoController "todo-service/controllers"
	"todo-service/handler"
	pb "todo-service/proto"
)

var (
	service = "todo-service"
	version = "latest"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name(service),
		micro.Version(version),
		micro.Broker(natsb.NewBroker()),
		micro.Registry(natsr.NewRegistry()),
		micro.Transport(natst.NewTransport()),
	)
	srv.Init()
	TodoController.InitializeDB()

	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
	pb.RegisterTodoServiceHandler(srv.Server(), new(handler.TodoService))

	time.Sleep(time.Second * 4)

	client := srv.Client()
	req := client.NewRequest("todo-service", "TodoService.GetAllTodos", &pb.Void{})
	res := &pb.AllTodoItems{}
	ctx := context.Background()
	if err := client.Call(ctx, req, res); err != nil {
		fmt.Println("[Error] ", err, res)
		return
	}

	for _, item := range res.Items {
		fmt.Println(item)
	}
}
