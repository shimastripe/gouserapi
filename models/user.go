package models

import "time"

type User struct {
	ID          uint      `gorm:"primary_key" json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	AccountName string    `json:"account_name,omitempty"`
	Email       string    `json:"email,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}
