package helpers

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"net/http"
	"strings"
	"time"
)

var (
	secretKey = []byte(viper.GetString("secret"))
)

func CreateAccessToken(email string, isAdmin bool) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":   email,
		"isAdmin": isAdmin,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenStr string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secretKey, nil
	})
	return token, err
}

func ExtractEmailFromToken(tokenStr string) (string, error) {
	token, err := ParseToken(tokenStr)
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		email, ok := claims["email"].(string)
		if !ok {
			return "", errors.New("email not valid")
		}
		return email, nil
	}
	return "", errors.New("invalid token")
}

func ExtractIsAdminFromToken(tokenStr string) (bool, error) {
	token, err := ParseToken(tokenStr)
	if err != nil {
		return false, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		isAdmin, ok := claims["isAdmin"].(bool)
		if !ok {
			return false, errors.New("admin role required")
		}
		return isAdmin, nil
	}
	return false, errors.New("invalid token")
}

func GetCurrentUser(r *http.Request) (string, bool, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", false, errors.New("no auth header")
	}

	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
	email, err := ExtractEmailFromToken(tokenStr)
	if err != nil {
		return "", false, err
	}
	isAdmin, err := ExtractIsAdminFromToken(tokenStr)
	if err != nil {
		return "", false, err
	}
	return email, isAdmin, nil
}
