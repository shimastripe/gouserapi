package models

import "time"

type User struct {
	ID          uint         `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Profile     *Profile     `json:"profile,omitempty"`
	ProfileID   uint         `gorm:"unique" json:"profile_id"`
	AccountName *AccountName `gorm:"unique" json:"account_name,omitempty"`
	Emails      []Email      `gorm:"unique" json:"emails"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}

// Belongs-to User
type Profile struct {
	ID       uint    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Name     string  `json:"name"`
	User     *User   `json:"user,omitempty"`
	Nation   *Nation `json:"nation,omitempty"`
	NationID uint    `json:"nation_id"`
}

// Belongs-to Profile
type Nation struct {
	ID      uint     `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Country string   `json:"country"`
	Profile *Profile `json:"profile,omitempty"`
}

// Has-one
type AccountName struct {
	ID          uint   `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	UserID      uint   `gorm:"unique" json:"user_id"`
	AccountName string `gorm:"unique" json:"account_name"`
	User        *User  `json:"user,omitempty"`
}

// Has-many
type Email struct {
	ID     uint   `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	UserID uint   `json:"user_id"`
	Email  string `gorm:"unique" json:"email"`
	User   *User  `json:"user,omitempty"`
}
