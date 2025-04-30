package main

import (
	"go-personal-boilerplate/routes"
	"go-personal-boilerplate/utils"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	db := utils.InitDB()
	if db == nil {
		log.Fatal("Failed to connect to the database")
	}

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// r.Use(func(c *gin.Context) {
	// 	log.Println("Middleware CORS aktif untuk:", c.Request.Method, c.Request.URL.Path)
	// 	c.Next()
	// })

	routes.SetupRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}

// r.Use(cors.New(cors.Config{
// 	AllowAllOrigins: true,
// 	AllowMethods:    []string{"GET", "POST", "PUT", "DELETE"},
// 	AllowHeaders:    []string{"Content-Type", "Authorization"},
// 	ExposeHeaders:   []string{"Content-Length"},
// 	MaxAge:          12 * time.Hour,
// }))
