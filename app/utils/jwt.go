package utils

import (
	"gotodo/app/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func CreateLoginSignature(u *models.User) (string, error) {
	mySigningKey := []byte(os.Getenv("JWT_SECRET_KEY"))

	// Create the Claims
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(3 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ID:        u.ID,
		Issuer:    u.Email,
		Subject:   u.Name,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(mySigningKey)
}

func ParseAndVerify(tokenString string) (any, error) {
	user := models.User{}
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
		user.ID = claims.ID
		user.Email = claims.Issuer
		user.Name = claims.Subject

		return user, nil
	}
	return user, err
}
