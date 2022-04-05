package user

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func generateToken(dayExpired int, data jwt.MapClaims) (string, error) {
	var hmacSampleSecret = []byte("Tran Trong Kim")
	createdAt := time.Now()
	expToken := createdAt.AddDate(0, 0, dayExpired)
	data["created_at"] = createdAt
	data["exp"] = expToken
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	return token.SignedString(hmacSampleSecret)

}
