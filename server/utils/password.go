package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	HandleErrorWithPanic(err)
	return string(hashedPassword)
}

func CompareHashPassword(passwordHash string, candidatePassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(candidatePassword)) == nil
}

func GenerateJWTToken(ttl time.Duration, payload interface{}, JWTSectret string) string {
	claims := jwt.MapClaims{}
	claims["iat"] = time.Now().Unix()
	claims["sub"] = payload

	if ttl != time.Duration(0) {
		claims["exp"] = time.Now().Add(ttl)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwt, err := token.SignedString([]byte(JWTSectret))
	HandleErrorWithPanic(err)

	return jwt
}

func ValidateJWToken(tokenString string, JWTSignedKey string) (interface{}, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected method of signed jwt token")
		}
		return []byte(JWTSignedKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return nil, fmt.Errorf("Cannot get required claims from token")
	}

	return claims["sub"], nil
}
