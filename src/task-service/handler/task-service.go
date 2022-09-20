package handler

import (
	"context"
	"fmt"
	"time"

	log "go-micro.dev/v4/logger"

	pb "task-service/proto"

	p "task-service/persistence"

	Model "task-service/models"
)

var ps = p.Persistence{}

type TaskService struct{}

func (e *TaskService) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	log.Infof("Received TaskService.Call request: %v", req)
	rsp.Msg = "Hello " + req.Name
	return nil
}

func (e *TaskService) CreateTodo(ctx context.Context, req *pb.TodoItem, rsp *pb.Status) error {
	fmt.Println("Inside The Handler")
	fmt.Println(req.Text, req.Color)
	dbInstance, _ := ps.InitializeDB()
	time.Sleep(time.Second * 5)
	dbInstance.Model(&p.TaskItem{}).Create(&Model.TaskItem{
		Text:  "Read Some Book !",
		Color: "Yellow",
	})
	rsp.MsgType = "Ok got it !"
	return nil
}
