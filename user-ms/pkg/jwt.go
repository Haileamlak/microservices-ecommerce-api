package pkg

import (
	// "strings"
	// "log"
	"time"
	"user-ms/internal/domain"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(user *domain.User) (string, *domain.AppError) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", domain.InternalErr("Failed to generate token")
	}

	return tokenString, nil
}

func VerifyToken(token string) (string, *domain.AppError) {
	// tokenParts := strings.Split(token, " ")
	// if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
	// 	return "", domain.UnauthorizedErr("Invalid token")
	// }

	tokenString := token

	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", domain.UnauthorizedErr("Invalid token")
		}

		return []byte("secret"), nil
	})
	if err != nil {
		return "", domain.UnauthorizedErr("Invalid token")
	}

	return parsedToken.Claims.(jwt.MapClaims)["id"].(string), nil
}
