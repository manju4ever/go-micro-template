package persistence

import (
	GlobalModels "task-service/models"

	"gorm.io/gorm"
)

type TaskItem struct {
	gorm.Model
	GlobalModels.TaskItem `gorm:"embedded"`
}
