package models

import (
	_ "github.com/go-sql-driver/mysql"
)

type Task struct {
	Name string	`gorm:"not null"`
	Date string `gorm:"not null"`
}
