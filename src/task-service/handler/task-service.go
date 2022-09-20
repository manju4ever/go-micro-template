package handler

import (
	"context"
	"fmt"

	log "go-micro.dev/v4/logger"

	pb "task-service/proto"

	p "task-service/persistence"

	Model "task-service/models"
)

var ps = p.Persistence{}
var dbInstance, _ = ps.InitializeDB()

type TaskService struct{}

func (e *TaskService) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	log.Infof("Received TaskService.Call request: %v", req)
	rsp.Msg = "Hello " + req.Name
	return nil
}

func (e *TaskService) CreateTodo(ctx context.Context, req *pb.TodoItem, rsp *pb.Status) error {
	fmt.Println("[CreateTodo] Inside Create Handler !")
	dbInstance.Create(&p.TaskItem{
		TaskItem: Model.TaskItem{
			Text:  req.Text,
			Color: req.Color,
		},
	})
	rsp.MsgType = "Task Created Successfully !"
	return nil
}
