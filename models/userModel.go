package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName  string    `gorm:"not null"`
	Username  string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null;"`
	Status    Status    `gorm:"type:ENUM('active', 'inactive');default:'active'"`
	Role      Role      `gorm:"type:ENUM('admin', 'operator');default:'operator'"`
	CreatedAt time.Time // Define the type as time.Time for the "created_at" column
	UpdatedAt time.Time // Define the type as time.Time for the "updated_at" column
}

type Status string
type Role string

const (
	Active   Status = "active"
	Inactive Status = "inactive"

	Admin    Role = "admin"
	Operator Role = "operator"
)

// EncryptPassword encrypts the user's password before creating or updating the user.
func (u *User) EncryptPassword() error {
	if len(u.Password) == 0 {
		return nil
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	return nil
}

// BeforeCreate is a Gorm hook called before creating a new user record.
func (u *User) BeforeCreate(tx *gorm.DB) error {
	return u.EncryptPassword()
}

// BeforeUpdate is a Gorm hook called before updating a user record.
func (u *User) BeforeUpdate(tx *gorm.DB) error {
	return u.EncryptPassword()
}
