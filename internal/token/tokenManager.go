package token

import (
	"errors"
	"time"

	"github.com/edelwei88/bytebuild-go/internal/postgres"
	"github.com/edelwei88/bytebuild-go/internal/postgres/models"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm/clause"
)

func GenerateJWT(user models.User) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  user.ID,
		"role": user.Role.Name,
		"iss":  "bytebuild-go",
		"exp":  time.Now().Add(time.Hour).Unix(),
		"iat":  time.Now().Unix(),
	})

	t, err := claims.SignedString(JWTSecretKey)
	if err != nil {
		return "", errors.New("failed to create JWT token")
	}

	return t, nil
}

func GetUserByJWT(tokenString string) (models.User, error) {
	var user models.User
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return JWTSecretKey, nil
	})
	if err != nil {
		return user, errors.New("failed to parse JWT")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return user, errors.New("token expired")
		}

		postgres.Postgres.Preload(clause.Associations).First(&user, claims["sub"])
		if user.ID == 0 {
			return user, errors.New("user not found")
		}

		return user, nil
	} else {
		return user, errors.New("failed to parse JWT")
	}
}
