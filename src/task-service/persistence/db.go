package persistence

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Persistence struct{}

var dbInstance *gorm.DB = nil

func (p *Persistence) Init() (dbInstance *gorm.DB, err error) {
	dsn := "host=localhost user=admin password=admin dbname=todo port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("[persistence][Init] Database connection error:", err)
		return nil, err
	}

	// Assign as singleton
	dbInstance = db

	// Migrate All Models from Here
	dbInstance.AutoMigrate(&TaskItem{})

	// Test Database Connection
	fmt.Println("[persistence][Init] Connection to DB Sucessful")
	return db, nil
}

func (p *Persistence) getInstance() (dbInstance *gorm.DB) {
	return dbInstance
}

func (p *Persistence) purge() {
	dbInstance = nil
}
