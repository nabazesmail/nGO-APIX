// repository/repository.go
package repository

import (
	"github.com/nabazesmail/gopher/src/initializers"
	"github.com/nabazesmail/gopher/src/models"
)

// inserting user to db
func CreateUser(user *models.User) error {
	result := initializers.DB.Create(user)
	return result.Error
}

// fetching all users from db
func GetAllUsers() ([]*models.User, error) {
	var users []*models.User
	result := initializers.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

// fetching user form db by Id
func GetUserByID(userID string) (*models.User, error) {
	var user models.User
	result := initializers.DB.First(&user, userID)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

// updating user in db
func UpdateUser(user *models.User) error {
	result := initializers.DB.Save(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// deleting user from db
func DeleteUser(user *models.User) error {
	result := initializers.DB.Delete(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// fetching user by username
func GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	result := initializers.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
