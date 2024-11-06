package middleware

import (
	"errors"
	"fmt"
	"monolithic-app/common"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("check auth ne")
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			panic(common.ErrInvalidRequest(errors.New("authorization header is blank")))
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			panic(common.ErrInvalidRequest(errors.New("authorization header is invalid")))
			return
		}

		tokenString := parts[1]

		secretKey := os.Getenv("JWT_SIGNER_KEY")
		if secretKey == "" {
			panic(common.ErrInternal(errors.New("JWT secret key is not set")))
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secretKey), nil
		})

		if err != nil {
			panic(common.ErrInvalidRequest(err))
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("user", claims) // Lưu claims vào context
		} else {
			panic(common.ErrInvalidRequest(errors.New("invalid token")))
			return
		}
		c.Next()
	}
}
