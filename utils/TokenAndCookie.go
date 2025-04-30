package utils

import (
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET")) // Ganti dengan kunci rahasia yang lebih aman

// Struktur Payload JWT
type Claims struct {
	UserID string `json:"userId"`
	jwt.RegisteredClaims
}

// TokenAndCookie membuat token JWT dan menyimpannya dalam cookie
func TokenAndCookie(userID string, w http.ResponseWriter) {
	// Tentukan waktu kedaluwarsa token
	expirationTime := time.Now().Add(10 * 24 * time.Hour) // Token berlaku 10 hari

	// Membuat klaim
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Buat token menggunakan SigningMethodHS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Tandatangani token dengan kunci rahasia
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "Tidak bisa membuat token", http.StatusInternalServerError)
		return
	}

	// Set cookie "jwt" dengan token yang baru dibuat
	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    tokenString,
		Expires:  expirationTime,
		HttpOnly: true,                    // Melindungi dari XSS
		SameSite: http.SameSiteStrictMode, // Melindungi dari CSRF
	})

	// Response sukses (opsional)
	// w.Write([]byte("Token telah dibuat dan disimpan dalam cookie"))
}
