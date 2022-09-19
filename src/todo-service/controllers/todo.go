package controllers

import (
	"encoding/json"
	"fmt"
	Schema "todo-service/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type TodoItemDB struct {
	gorm.Model
	Item Schema.TodoItem `gorm:"embedded"`
}

var dbInstance *gorm.DB = nil

func InitializeDB() (instance *gorm.DB) {
	dsn := "host=localhost user=admin password=admin dbname=todo port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	dbInstance = db
	if err != nil {
		fmt.Println("[todo][database] Database connection error:", err)
	}
	dbInstance.Debug().AutoMigrate(&TodoItemDB{})
	fmt.Println("[database] Connection to DB Sucessful !")
	return dbInstance
}

func CreateNewTodo(item Schema.TodoItem) (status bool, err error) {
	dbInstance.Create(&TodoItemDB{Item: item})
	return true, nil
}

func UpdateItemStatus(itemId int, newStatus string) (status bool, err error) {
	dbInstance.Model(TodoItemDB{}).Where("id = ?", itemId).Update("status", newStatus)
	return true, nil
}

func GetAllTodos() (items []TodoItemDB, err error) {
	var allItems []TodoItemDB
	dbInstance.Model(TodoItemDB{}).Find(&allItems)
	for _, exact := range allItems {
		obj, _ := json.Marshal(exact)
		fmt.Println(string(obj))
	}
	return allItems, nil
}
