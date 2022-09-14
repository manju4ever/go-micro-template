package controllers

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type TodoItem struct {
	gorm.Model
	Text   string
	Status string
}

func InitDB() {

	dsn := "host=localhost user=admin password=admin dbname=todo port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("[todo][database] Database connection error:", err)
	}

	fmt.Println("[database] Connection to DB Sucessful !")

	db.AutoMigrate(&TodoItem{})

	// db.Where("text is NULL AND status is NULL").Delete(&TodoItem{})

	// Add Some Items
	db.Create(&TodoItem{Text: "Should read some stuff on go", Status: "new"})
	db.Create(&TodoItem{Text: "Clean my garden", Status: "in-progress"})

}
