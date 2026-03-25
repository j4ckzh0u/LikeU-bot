package middleware

import (
	"net/http"
	"strings"

	"ai-coding-assistant/internal/service"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	Service *service.Service
}

type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func JWT() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return next(c)
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString == authHeader {
				return next(c)
			}

			token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte("secret"), nil
			})

			if err != nil || !token.Valid {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"code":    2002,
					"message": "invalid token",
				})
			}

			if claims, ok := token.Claims.(*Claims); ok {
				c.Set("user_id", claims.UserID)
			}

			return next(c)
		}
	}
}
