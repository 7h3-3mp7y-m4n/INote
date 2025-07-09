package main

import (
	"7h3-3mp7y-m4n/INote/internal/config"
	"7h3-3mp7y-m4n/INote/internal/handlers"
	"fmt"
	"log"
	"os"

	"7h3-3mp7y-m4n/INote/internal/db"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// func main() {

// 	if os.Getenv("GO_ENV") != "production" {
// 		_ = godotenv.Load()
// 	}
// 	cfg := config.LoadConfig()
// 	log.Printf("Starting server on port %s", cfg.AppPort)
// 	fmt.Println("Loaded config:")
// 	fmt.Printf("AppPort: %s\n", cfg.AppPort)
// 	fmt.Printf("PostgresHost: %s\n", cfg.PostgresHost)
// 	fmt.Printf("PostgresUser: %s\n", cfg.PostgresUser)
// 	fmt.Printf("PostgresPass: %s\n", cfg.PostgresPass)
// 	fmt.Printf("PostgresDB: %s\n", cfg.PostgresDB)
// 	fmt.Printf("RedisAddr: %s\n", cfg.RedisAddr)
// 	r := gin.Default()
// 	database, err := db.NewDB()
// 	if err != nil {
// 		log.Fatalf("failed to connect db: %v", err)
// 	}
// 	defer database.Pool.Close()

// 	fmt.Println("DB connected")
// 	r.GET("/health", func(c *gin.Context) {
// 		c.JSON(200, gin.H{"status": "UP"})
// 	})
// 	r.POST("/notes", handlers.CreateNoteHandler(db))
// 	err = r.Run(":" + cfg.AppPort)
// 	if err != nil {
// 		log.Fatal("Failed to start server:", err)
// 	}
// }

func main() {
	// Load .env in non-production environments
	if os.Getenv("GO_ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			log.Println("Warning: could not load .env file")
		}
	}

	// Load config
	cfg := config.LoadConfig()
	fmt.Println("Loaded config:")
	fmt.Printf("AppPort: %s\n", cfg.AppPort)
	fmt.Printf("PostgresHost: %s\n", cfg.PostgresHost)
	fmt.Printf("PostgresUser: %s\n", cfg.PostgresUser)
	fmt.Printf("PostgresPass: %s\n", cfg.PostgresPass)
	fmt.Printf("PostgresDB: %s\n", cfg.PostgresDB)
	fmt.Printf("RedisAddr: %s\n", cfg.RedisAddr)

	// Init DB
	database, err := db.NewDB()
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}
	defer database.Pool.Close()
	fmt.Println("✅ DB connected")

	// Init Gin server
	r := gin.Default()

	// Health check route
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "UP"})
	})
	// Note routes
	r.POST("/notes", handlers.CreateNoteHandler(database))
	// Later: r.GET("/notes/:slug", handlers.GetNoteHandler(database))

	// Start server
	log.Printf("🚀 Starting server on port %s...", cfg.AppPort)
	if err := r.Run(":" + cfg.AppPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
