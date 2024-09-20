package middleware

import (
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/willykurniawan01/linknau-test/app/helpers"
)

type JwtMiddleware struct{}

func (jm *JwtMiddleware) VerifyJwt(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")

	if tokenString == "" {
		helpers.ResponseApi(c, "INVALID_JWT_TOKEN", map[string]interface{}{})
		c.Abort()
		return
	}

	if !strings.HasPrefix(tokenString, "Bearer ") {
		helpers.ResponseApi(c, "INVALID_JWT_TOKEN", map[string]interface{}{
			"error": "Authorization header format must be Bearer <token>",
		})
		c.Abort()
		return
	}

	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	jwtSecretKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	})

	if err != nil || !token.Valid {
		helpers.ResponseApi(c, "INVALID_JWT_TOKEN", map[string]interface{}{
			"error": err.Error(),
		})
		c.Abort()
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if exp, ok := claims["expired_at"].(float64); ok {
			if time.Unix(int64(exp), 0).Before(time.Now()) {
				helpers.ResponseApi(c, "EXPIRED_JWT_TOKEN", map[string]interface{}{
					"error": "JWT token has expired",
				})
				c.Abort()
				return
			}
		} else {
			helpers.ResponseApi(c, "INVALID_JWT_TOKEN", map[string]interface{}{
				"error": "JWT token missing expiration",
			})
			c.Abort()
			return
		}
	}
	c.Next()
}
