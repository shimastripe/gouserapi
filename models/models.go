package models

import "time"

type User struct {
	ID          uint         `gorm:"primary_key;AUTO_INCREMENT" json:"id,omitempty"`
	Profile     *Profile     `json:"profile,omitempty"`
	ProfileID   uint         `gorm:"unique" json:"profile_id,omitempty"`
	AccountName *AccountName `gorm:"unique" json:"account_name,omitempty"`
	Emails      []Email      `gorm:"unique" json:"emails,omitempty"`
	CreatedAt   *time.Time   `json:"created_at,omitempty"`
	UpdatedAt   *time.Time   `json:"updated_at,omitempty"`
}

// Belongs-to User
type Profile struct {
	ID       uint    `gorm:"primary_key;AUTO_INCREMENT" json:"id,omitempty"`
	Name     string  `json:"name,omitempty"`
	User     *User   `json:"user,omitempty"`
	Nation   *Nation `json:"nation,omitempty"`
	NationID uint    `json:"nation_id,omitempty"`
}

// Belongs-to Profile
type Nation struct {
	ID      uint     `gorm:"primary_key;AUTO_INCREMENT" json:"id,omitempty"`
	Country string   `json:"country,omitempty"`
	Profile *Profile `json:"profile,omitempty"`
}

// Has-one
type AccountName struct {
	ID          uint   `gorm:"primary_key;AUTO_INCREMENT" json:"id,omitempty"`
	UserID      uint   `gorm:"unique" json:"user_id,omitempty"`
	AccountName string `gorm:"unique" json:"account_name,omitempty"`
	User        *User  `json:"user,omitempty"`
}

// Has-many
type Email struct {
	ID     uint   `gorm:"primary_key;AUTO_INCREMENT" json:"id,omitempty"`
	UserID uint   `json:"user_id,omitempty"`
	Email  string `gorm:"unique" json:"email,omitempty"`
	User   *User  `json:"user,omitempty"`
}
