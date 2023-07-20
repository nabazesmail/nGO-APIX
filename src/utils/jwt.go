package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/nabazesmail/gopher/src/models"
)

// GenerateJWTToken generates a new JWT token for the provided user
func GenerateJWTToken(user *models.User, secretKey []byte) (string, error) {
	// Create a new token with the user's ID as the subject (sub) claim
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		// You can add more user information to the token as needed
		"username": user.Username,
		"fullName": user.FullName,
		"role":     user.Role,
		"status":   user.Status,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expiration time (24 hours from now)
	})

	// Sign the token with the provided secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
