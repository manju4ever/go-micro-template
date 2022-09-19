package main

import (
	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"

	natsb "github.com/go-micro/plugins/v4/broker/nats"
	natsr "github.com/go-micro/plugins/v4/registry/nats"
	natst "github.com/go-micro/plugins/v4/transport/nats"

	TodoController "todo-service/controllers"
	"todo-service/handler"
	Schema "todo-service/models"
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
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
	pb.RegisterTodoServiceHandler(srv.Server(), new(handler.TodoService))
	TodoController.InitializeDB()
	TodoController.CreateNewItem(Schema.TodoItem{Text: "Read some stuff", Status: "in-progress", Tags: "reading, writing"})
	TodoController.UpdateItemStatus(14, "done")
	TodoController.GetAllItems()
}
