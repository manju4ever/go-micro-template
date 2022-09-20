package main

import (
	"context"
	"fmt"
	"task-service/handler"
	pb "task-service/proto"
	"time"

	ps "task-service/persistence"

	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
)

var (
	service = "task-service"
	version = "latest"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name(service),
		micro.Version(version),
	)

	srv.Init()

	// Register handler
	pb.RegisterTaskServiceHandler(srv.Server(), new(handler.TaskService))
	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}

	p := ps.Persistence{}
	p.InitializeDB()

	time.Sleep(time.Second * 6)

	client := srv.Client()

	req := client.NewRequest("task-service", "TaskService.CreateTodo", &pb.TodoItem{Text: "Read some stuff", Color: "Yellow"})
	res := &pb.Status{}
	if err := client.Call(context.Background(), req, res); err != nil {
		fmt.Println("[error] Something went wrong:", err)
	}

	fmt.Println("I'm at the end !", res)

}
