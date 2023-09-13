package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context){
	// Get the authentication header value
	authHeader := c.GetHeader("Authorization")

	// Check if the header is empty or doesn't start with "Bearer"
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error" : "Unauthorized",
		})
		c.Abort()
		return
	}

	//Extract the token from the header
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	//Parse the JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token)(interface{}, error){
		//Provided the key for verifying the token's signature (replace with your actual key)
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error" : err.Error(),
		})
		c.Abort()
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		c.Set("jwtClaims", claims)
	}
	//if the token is valid, proceed with the next middleware/handler
	c.Next()
}