package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Task struct {
	gorm.Model
	Name  string `gorm:"not null"`
	State string `gorm:"not null;default:'waiting'"`
}

func CreateTask(taskName string, db *gorm.DB) uint {
	task := Task{Name: taskName}
	db.Create(&task)

	if db.NewRecord(task) {
		return task.ID
	} else {
		return task.ID
	}
}
