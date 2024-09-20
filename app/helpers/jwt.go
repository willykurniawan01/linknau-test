package helpers

import (
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func GenerateJWT(user gin.H) (string, error) {
	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	jwtExpiration, _ := strconv.Atoi(os.Getenv("JWT_EXPIRATION"))
	expirationTime := time.Now().Add(time.Duration(jwtExpiration) * time.Minute)

	claims := &jwt.MapClaims{
		"user":       user,
		"expired_at": expirationTime.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecretKey))
}
