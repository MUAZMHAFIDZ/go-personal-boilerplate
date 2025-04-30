package utils

import "os"

// Kunci rahasia untuk JWT
var JwtKey = []byte(os.Getenv("JWT_SECRET"))
