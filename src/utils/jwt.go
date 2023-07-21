package utils

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/nabazesmail/gopher/src/models"
)

// JWTSecretKey is your JWT secret key.
var JWTSecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

// GenerateJWTToken generates a new JWT token for the provided user.
func GenerateJWTToken(user *models.User, secretKey []byte) (string, error) {
	// Create a new token with the user's ID as the subject (sub) claim.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		// You can add more user information to the token as needed.
		"username": user.Username,
		"fullName": user.FullName,
		"role":     user.Role,
		"status":   user.Status,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expiration time (24 hours from now).
	})

	// Sign the token with the provided secret key.
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// VerifyJWTToken verifies the JWT token and returns the claims if the token is valid.
func VerifyJWTToken(tokenString string, secretKey []byte) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method of the token.
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	// Check if the token is valid and contains valid claims.
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}

// UserResponse represents the user information to be returned in the API response
type UserResponse struct {
	ID       uint   `json:"id"`
	FullName string `json:"fullName"`
	Username string `json:"username"`
	Status   string `json:"status"`
	Role     string `json:"role"`
}
