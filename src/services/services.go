// services/services.go
package services

import (
	"errors"
	"regexp"

	"github.com/nabazesmail/gopher/src/models"
	"github.com/nabazesmail/gopher/src/repository"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(body *models.User) (*models.User, error) {
	// Validate the input
	if body.FullName == "" || body.Username == "" || body.Password == "" {
		return nil, errors.New("all fields must be provided")
	}

	// Validate username using regex (allow only characters)
	usernameRegex := regexp.MustCompile("^[a-zA-Z]+$")
	if !usernameRegex.MatchString(body.Username) {
		return nil, errors.New("username must contain only characters")
	}

	if len(body.Password) < 8 || len(body.Password) > 15 {
		return nil, errors.New("password must be between 8 and 15 characters")
	}

	// Validate status and role (if provided)
	if body.Status != models.Active && body.Status != models.Inactive {
		return nil, errors.New("invalid status value")
	}

	if body.Role != models.Admin && body.Role != models.Operator {
		return nil, errors.New("invalid role value")
	}

	// Hash the password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create a new User instance with the hashed password
	user := &models.User{
		FullName: body.FullName,
		Username: body.Username,
		Password: string(hashedPassword),
		Status:   body.Status,
		Role:     body.Role,
	}

	// Save the user in the database
	err = repository.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
