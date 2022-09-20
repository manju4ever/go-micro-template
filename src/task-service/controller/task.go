package controller

import (
	"fmt"
	Models "task-service/models"
)

type TaskController struct{}

func (t *TaskController) CreateNewTask() {
	someItem := &Models.TaskItem{Text: "Hello", Color: "Red"}
	fmt.Println(someItem)
}

func (t *TaskController) GetAllTasks() {

}
