// services/services.go
package services

import (
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"mime/multipart"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/nabazesmail/gopher/src/models"
	"github.com/nabazesmail/gopher/src/repository"
	"github.com/nabazesmail/gopher/src/utils"
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
		log.Printf("Error hashing password: %s", err)
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
		log.Printf("Error saving user in the database: %s", err)
		return nil, err
	}

	return user, nil
}

func GetAllUsers() ([]*models.User, error) {
	users, err := repository.GetAllUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func GetUserByID(userID string) (*models.User, error) {
	if userID == "" {
		return nil, errors.New("user ID must be provided")
	}

	user, err := repository.GetUserByID(userID)
	if err != nil {
		log.Printf("Error fetching user by ID: %s", err)
		return nil, err
	}

	return user, nil
}

func UpdateUserByID(userID string, body *models.User) (*models.User, error) {
	if userID == "" {
		return nil, errors.New("user ID must be provided")
	}

	user, err := repository.GetUserByID(userID)
	if err != nil {
		log.Printf("Error fetching user by ID: %s", err)
		return nil, err
	}

	if user == nil {
		return nil, nil // User not found
	}

	// Update user fields if they are provided in the request body
	if body.FullName != "" {
		user.FullName = body.FullName
	}

	if body.Username != "" {
		user.Username = body.Username
	}

	if body.Password != "" {
		// Hash the password using bcrypt
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Printf("Error hashing password: %s", err)
			return nil, err
		}
		user.Password = string(hashedPassword)
	}

	if body.Status != "" {
		user.Status = body.Status
	}

	if body.Role != "" {
		user.Role = body.Role
	}

	// Save the updated user in the database
	err = repository.UpdateUser(user)
	if err != nil {
		log.Printf("Error updating user: %s", err)
		return nil, err
	}

	return user, nil
}

func DeleteUserByID(userID string) error {
	if userID == "" {
		return errors.New("user ID must be provided")
	}

	user, err := repository.GetUserByID(userID)
	if err != nil {
		log.Printf("Error fetching user by ID: %s", err)
		return err
	}

	if user == nil {
		return nil // User not found
	}

	// Delete the user from the database
	err = repository.DeleteUser(user)
	if err != nil {
		log.Printf("Error deleting user: %s", err)
		return err
	}

	return nil
}

func AuthenticateUser(body *models.User) (string, error) {
	// Find the user by username in the database
	user, err := repository.GetUserByUsername(body.Username)
	if err != nil {
		log.Printf("Error fetching user by username: %s", err)
		return "", err
	}

	if user == nil {
		return "", errors.New("user not found")
	}

	// Compare the provided password with the hashed password in the database
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		log.Printf("Password verification failed for user %s: %s", user.Username, err)
		return "", errors.New("incorrect password")
	}

	// Generate a JWT token using the utils package
	tokenString, err := utils.GenerateJWTToken(user, []byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		log.Printf("Error generating JWT token: %s", err)
		return "", errors.New("failed to generate JWT token")
	}

	return tokenString, nil
}

func UpdateUserProfilePicture(userID string, fileHeader *multipart.FileHeader) (*models.User, error) {
	// Find the user by ID in the database
	user, err := repository.GetUserByID(userID)
	if err != nil {
		log.Printf("Error fetching user by ID: %s", err)
		return nil, err
	}

	if user == nil {
		return nil, nil // User not found
	}

	// Check if the uploaded file is an image
	if !isImageFile(fileHeader) {
		return nil, errors.New("invalid file format, only images are allowed")
	}

	// Create a unique filename for the uploaded image
	filename := generateUniqueFilename(fileHeader)

	// Create the file path for storing the uploaded image
	filePath := filepath.Join("public/uploads", filename)

	// Open the uploaded file
	file, err := fileHeader.Open()
	if err != nil {
		log.Printf("Error opening uploaded file: %s", err)
		return nil, err
	}
	defer file.Close()

	// Create the destination file
	dst, err := os.Create(filePath)
	if err != nil {
		log.Printf("Error creating destination file: %s", err)
		return nil, err
	}
	defer dst.Close()

	// Copy the file data to the destination file
	_, err = io.Copy(dst, file)
	if err != nil {
		log.Printf("Error copying file data: %s", err)
		return nil, err
	}

	// Update the user's profile picture URL in the database with the original filename
	user.ProfilePicture = fileHeader.Filename
	if err := repository.UpdateUser(user); err != nil {
		log.Printf("Error updating user's profile picture: %s", err)
		return nil, err
	}

	return user, nil
}

// Helper function to check if the uploaded file is an image
func isImageFile(fileHeader *multipart.FileHeader) bool {
	// Extract the file extension from the uploaded file's header
	ext := filepath.Ext(fileHeader.Filename)
	ext = strings.ToLower(ext)
	return ext == ".jpg" || ext == ".jpeg" || ext == ".png" || ext == ".gif"
}

// Helper function to generate a unique filename for the uploaded image
func generateUniqueFilename(fileHeader *multipart.FileHeader) string {
	ext := filepath.Ext(fileHeader.Filename)
	// Create a unique filename using the original filename, timestamp, and a random number
	return fmt.Sprintf("%s_%d%d%s", strings.TrimSuffix(fileHeader.Filename, ext), time.Now().UnixNano(), rand.Intn(10000), ext)
}
