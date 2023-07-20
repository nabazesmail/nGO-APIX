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

func GetAllUsers() ([]*models.User, error) {
	var users []*models.User
	result := initializers.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func GetUserByID(userID string) (*models.User, error) {
	var user models.User
	result := initializers.DB.First(&user, userID)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func UpdateUser(user *models.User) error {
	result := initializers.DB.Save(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func DeleteUser(user *models.User) error {
	result := initializers.DB.Delete(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
