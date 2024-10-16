package auth

import (
	"errors"
	"time"

	"github.com/Tes-sudo/online-learning-platform/user-service/models"
	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("your_secret_key")

type Claims struct{
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(user *models.UserModel)(string,error){
	// Set token expiration time to 24 hours from now
	expiration := time.Now().Add(24*time.Hour)

	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiration),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtKey)
}

// ValidateToken checks if the provided token is valid
func ValidateToken(tokenString string) (*Claims, error) {
	// Initialize a new instance of `Claims`
	claims := &Claims{}

	// Parse the JWT string and store the result in `claims`.
	// This method will return an error if the token is invalid
	// or if the signature does not match
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	// If there's an error in parsing...
	if err != nil {
		return nil, err
	}

	// Check if the token is valid
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	// If we get here, then the token is valid. Return the claims
	return claims, nil
}