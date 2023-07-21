package models

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName       string    `gorm:"not null"`
	Username       string    `gorm:"unique;not null"`
	Password       string    `gorm:"not null;"`
	Status         Status    `gorm:"type:ENUM('active', 'inactive');default:'active'"`
	Role           Role      `gorm:"type:ENUM('admin', 'operator');default:'operator'"`
	ProfilePicture string    // this field for profile picture name
	CreatedAt      time.Time //  the type as time.Time for the "created_at" column
	UpdatedAt      time.Time //  the type as time.Time for the "updated_at" column
}

type Status string
type Role string

const (
	Active   Status = "active"
	Inactive Status = "inactive"

	Admin    Role = "admin"
	Operator Role = "operator"
)

// SerializeUser serializes the user data to a JSON string.
func (u *User) Serialize() (string, error) {
	userJSON, err := json.Marshal(u)
	if err != nil {
		return "", err
	}
	return string(userJSON), nil
}

// DeserializeUser deserializes the JSON string to a User object.
func DeserializeUser(data string) (*User, error) {
	var user User
	if err := json.Unmarshal([]byte(data), &user); err != nil {
		return nil, err
	}
	return &user, nil
}
