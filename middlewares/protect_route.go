package middlewares

import (
	"fmt"
	"go-personal-boilerplate/models"
	"go-personal-boilerplate/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jinzhu/gorm"
)

// ProtectRoute is a middleware that checks if the user has a valid JWT token
func ProtectRoute(c *gin.Context) {
	// Get JWT token from cookie
	cookie, err := c.Cookie("jwt")
	if err != nil {
		// If cookie is not found, return 401 Unauthorized response
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: No token found"})
		c.Abort() // Prevent further handling
		return
	}

	claims := &utils.Claims{}
	tokenString := cookie

	// Parse and verify the token
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is HS256
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		// Return the key for verification
		return []byte(utils.JwtKey), nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Invalid token"})
		c.Abort() // Prevent further handling
		return
	}

	// Find user by ID from the token claims (using UUID)
	db := utils.InitDB()
	var user models.User
	err = db.Where("id = ?", claims.UserID).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		c.Abort() // Prevent further handling
		return
	}

	// Store the user in the request context
	c.Set("user", user)

	// Call the next handler
	c.Next()
}
