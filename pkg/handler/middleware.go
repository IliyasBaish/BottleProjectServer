package handler

import (
	"errors"
	"net/http"
	"strings"

	"example.com/server/pkg/jwt_auth"
	"github.com/gin-gonic/gin"
)

func Middleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	if headerParts[0] != "Bearer" {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	username, err := jwt_auth.ParseToken(headerParts[1], []byte("secret"))
	if err != nil {
		status := http.StatusBadRequest
		if err == errors.New("Ivalid Access Token") {
			status = http.StatusUnauthorized
		}

		c.AbortWithStatus(status)
		return
	}
	c.Request.Header.Add("userobg", username)
}
