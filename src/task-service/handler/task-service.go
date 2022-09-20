package handler

import (
	"context"
	"fmt"

	log "go-micro.dev/v4/logger"

	pb "task-service/proto"

	p "task-service/persistence"

	Model "task-service/models"
)

var dbInstance, _ = (&p.Persistence{}).Init()

type TaskService struct{}

func (e *TaskService) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	log.Infof("Received TaskService.Call request: %v", req)
	rsp.Msg = "Hello " + req.Name
	return nil
}

func (e *TaskService) CreateTodo(ctx context.Context, req *pb.TodoItem, rsp *pb.Status) error {
	fmt.Println("[CreateTodo] Inside Handler")
	dbInstance.Create(&p.TaskItem{
		TaskItem: Model.TaskItem{
			Text:  req.Text,
			Color: req.Color,
		},
	})
	rsp.MsgType = "[CreateTodo] Task Created Successfully"
	return nil
}

func (e *TaskService) GetAllTodos(ctx context.Context, req *pb.Void, rsp *pb.TodoItems) error {
	fmt.Println("[GetAllTodos] Inside Handler")
	var results []Model.TaskItem = nil
	dbInstance.Find(&results)
	itemsToSend := make([]*pb.TodoItem, 1, 1)
	for _, each := range results {
		// fmt.Println(each.Text, " - ", each.Color)
		itemsToSend = append(itemsToSend, &pb.TodoItem{
			Text:  each.Text,
			Color: each.Color,
		})
	}
	rsp.Items = itemsToSend
	return nil
}
