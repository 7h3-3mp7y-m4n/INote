package main

import (
	"7h3-3mp7y-m4n/INote/internal/config"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if os.Getenv("GO_ENV") != "production" {
		_ = godotenv.Load()
	}
	cfg := config.LoadConfig()
	log.Printf("Starting server on port %s", cfg.AppPort)
	fmt.Println("Loaded config:")
	fmt.Printf("AppPort: %s\n", cfg.AppPort)
	fmt.Printf("PostgresHost: %s\n", cfg.PostgresHost)
	fmt.Printf("PostgresUser: %s\n", cfg.PostgresUser)
	fmt.Printf("PostgresPass: %s\n", cfg.PostgresPass)
	fmt.Printf("PostgresDB: %s\n", cfg.PostgresDB)
	fmt.Printf("RedisAddr: %s\n", cfg.RedisAddr)
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "UP"})
	})

	err := r.Run(":" + cfg.AppPort)
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
