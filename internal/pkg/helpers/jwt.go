package helpers

import (
	"time"
	"user-service/internal/pkg/helpers/errors"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(userID int) (tokenString string, refreshToken string, expiredAt time.Time, err error) {
	// Define the secret key
	secret := "your-secret-key"

	expiredAt = time.Now().Add(time.Hour)
	// Create the claims
	claims := jwt.MapClaims{
		"userID": userID,
		"exp":    expiredAt,
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// refresh token

	// Sign the token with the secret key
	tokenString, err = token.SignedString([]byte(secret))
	if err != nil {
		return "", "", time.Time{}, errors.BadRequest("error generating token")
	}

	// Return the token and other information
	return tokenString, "", expiredAt, nil
}
