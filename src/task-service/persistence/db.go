package persistence

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Persistence struct {
	dbInstance *gorm.DB
}

func (p *Persistence) InitializeDB() (dbInstance *gorm.DB, err error) {
	dsn := "host=localhost user=admin password=admin dbname=todo port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("[todo][database] Database connection error:", err)
		return nil, err
	}
	dbInstance = db
	// Migrate All Models fro Here
	p.dbInstance.AutoMigrate(&TaskItem{})
	// Test Database Connection
	fmt.Println("[database] Connection to DB Sucessful !")
	return db, nil
}
