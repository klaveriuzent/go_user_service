package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
	"userservice/controller"
	"userservice/database"
	"userservice/middleware"
	"userservice/schema"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const port = ":8000"

// Load all func
func main() {
	loadEnv()
	loadDatabase()
	serveApplication()
	testGetDateTime()
}

// Test the Go time package by printing the current date and time in various formats
func testGetDateTime() {
	currentTime := time.Now()

	// Print the individual elements of the current time
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		currentTime.Year(), currentTime.Month(), currentTime.Day(),
		currentTime.Hour(), currentTime.Minute(), currentTime.Second())

	// Print the current time in custom formats
	fmt.Println(currentTime.Format("2006-01-02 15:04:05"))
	fmt.Println(currentTime.Format("2006-January-02"))
	fmt.Println(currentTime.Format("2006-01-02 3:4:5 pm"))

}

// Load environment variables from .env.local file
func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// Connect to the database and perform auto-migration
func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(
		&schema.Role{},
		&schema.User{},
		&schema.UserRole{},
		&schema.Profile{},
		&schema.RoleApplication{},
		&schema.Account{},
		&schema.Address{},
		&schema.ActivityLog{},
	)
}

// Start the application server using Gin framework
func serveApplication() {
	gin.DisableConsoleColor() // Disable console color for logging

	// Create a log file for logging
	f, _ := os.Create("./logger.log")
	// Use the following code if you need to write the logs to file and console at the same time.
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := gin.Default()                    // Create a new Gin router
	router.Use(gin.Recovery())                 // Use Gin recovery middleware to recover from any panics
	router.Use(middleware.JSONLogMiddleware()) // Use custom middleware to log requests and responses in JSON format

	// Define public routes for authentication
	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controller.Register)
	publicRoutes.POST("/login", controller.Login)

	// Define protected routes for authorized users
	protectedRoutes := router.Group("/v1")
	protectedRoutes.Use(middleware.JWTAuthMiddleware())
	protectedRoutes.GET("/user/:ID", controller.UserGetProfiles)
	protectedRoutes.POST("/user/:ID/assign_roles", controller.UserAssignRole)
	protectedRoutes.POST("/user/:ID/assign_role_app", controller.UserAssignRoleApplication)
	protectedRoutes.PATCH("/user/user_account/:ID/edit", controller.UserAccountUpdate)

	// Start the server and log any errors
	if err := router.Run(port); err != nil {
		log.Printf("Failed to start server: %v\n", err)
	}
}
