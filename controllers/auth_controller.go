package controllers

import (
	"go-personal-boilerplate/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Signup handler untuk mendaftar user baru
func Signup(c *gin.Context) {
	var userRequest struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Parsing request body
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Daftarkan user baru
	user, err := services.RegisterUser(userRequest.Name, userRequest.Email, userRequest.Password, c.Writer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	// Kirim response dengan status Created
	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user":    user,
	})
}

// Login handler untuk autentikasi user
func Login(c *gin.Context) {
	var loginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Parsing request body
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Autentikasi user
	user, err := services.AuthenticateUser(loginRequest.Email, loginRequest.Password, c.Writer)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Kirim response sukses dengan status OK
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user":    user,
	})
}
