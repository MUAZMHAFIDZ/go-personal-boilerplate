package services

import (
	"errors"
	"go-personal-boilerplate/models"
	"go-personal-boilerplate/utils"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword will hash the password using bcrypt
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// VerifyPassword will verify the hashed password
func VerifyPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

// RegisterUser performs the user registration process and sends JWT in a cookie
func RegisterUser(name, email, password string, w http.ResponseWriter) (*models.User, error) {
	// Hash the password
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return nil, err
	}

	// Create new user model
	newUser := models.User{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
	}

	// Initialize DB connection
	db := utils.InitDB()

	// Save the user to the database
	if err := db.Create(&newUser).Error; err != nil {
		return nil, err
	}

	// Generate JWT token and send it in a cookie
	utils.TokenAndCookie(newUser.ID, w) // Send token after registration

	return &newUser, nil
}

// AuthenticateUser verifies the login credentials of the user and sends JWT in a cookie
func AuthenticateUser(email, password string, w http.ResponseWriter) (*models.User, error) {
	// Initialize DB connection
	db := utils.InitDB()

	// Find user from the database by email
	var user models.User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Verify the password
	if !VerifyPassword(user.Password, password) {
		return nil, errors.New("invalid credentials")
	}

	// Generate JWT token and send it in a cookie
	utils.TokenAndCookie(user.ID, w) // Send token after authentication

	return &user, nil
}
