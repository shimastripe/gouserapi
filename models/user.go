package models

import "time"

type User struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	Name        string    `json:"name"`
	AccountName string    `json:"account_name"`
	Email       string    `json:"email"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
