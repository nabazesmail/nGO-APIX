package models

import (
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
	ProfilePicture string    // New field for profile picture URL
	CreatedAt      time.Time // Define the type as time.Time for the "created_at" column
	UpdatedAt      time.Time // Define the type as time.Time for the "updated_at" column
}

type Status string
type Role string

const (
	Active   Status = "active"
	Inactive Status = "inactive"

	Admin    Role = "admin"
	Operator Role = "operator"
)
