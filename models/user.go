package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	gorm.Model
	Name        string
	AccountName string
	Email       string
}
