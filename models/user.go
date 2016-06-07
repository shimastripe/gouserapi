package models

import "time"

type User struct {
	ID          uint `gorm:"primary_key"`
	Name        string
	AccountName string
	Email       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
