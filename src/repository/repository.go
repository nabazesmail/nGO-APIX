// repository/repository.go
package repository

import (
	"github.com/nabazesmail/gopher/src/initializers"
	"github.com/nabazesmail/gopher/src/models"
)

func CreateUser(user *models.User) error {
	result := initializers.DB.Create(user)
	return result.Error
}
