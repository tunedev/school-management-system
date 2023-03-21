package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tunedev/school-management-system/internal/config"
	"github.com/tunedev/school-management-system/pkg/database"
)

func main() {
	// Load configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("failed to load configuration na the actual error be this: %v", err)
	}
	fmt.Println("NA THE ENV BE THIS OOH =====>>>>",cfg.DSN())

	// Connect to database
	_, err = database.NewDatabase(cfg.DSN())
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Initialize Gin router
	r := gin.Default()

	// Initialize HTTP handlers

	// Define routes
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to the School Management System API"})
	})

	// Start server
	err = r.Run(":8080")
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
