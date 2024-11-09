package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	helper "github.com/roh4nyh/swaggo/helpers"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("Authorization")
		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No Authorization header found"})
			c.Abort()
			return
		}

		claims, err := helper.ValidateToken(clientToken)
		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Set("firstname", claims.Firstname)
		c.Set("lastname", claims.Lastname)
		c.Set("role", claims.Role)
		c.Set("userId", claims.UserId)

		// log.Printf("claims: %+v\n", claims)

		c.Next()
	}
}

func AuthenticateAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetString("role")
		if role != "ADMIN" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "UnAuthenticated access to this resource"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func AuthenticateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetString("role")
		if role != "USER" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "UnAuthenticated access to this resource"})
			c.Abort()
			return
		}

		c.Next()
	}
}
