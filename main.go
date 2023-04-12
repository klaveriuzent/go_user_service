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

func main() {
	loadEnv()
	loadDatabase()
	serveApplication()
	testGetDateTime()
}

func testGetDateTime() {
	currentTime := time.Now()

	// individual elements of time can
	// also be called to print accordingly
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		currentTime.Year(), currentTime.Month(), currentTime.Day(),
		currentTime.Hour(), currentTime.Minute(), currentTime.Second())

	// formatting time using
	// custom formats
	fmt.Println(currentTime.Format("2006-01-02 15:04:05"))
	fmt.Println(currentTime.Format("2006-January-02"))
	fmt.Println(currentTime.Format("2006-01-02 3:4:5 pm"))

}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

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
		&schema.Corporation{},
		&schema.ProfileCorporations{},
		&schema.ActivityLog{},
	)
}

func serveApplication() {
	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create("logger.log")
	// Use the following code if you need to write the logs to file and console at the same time.
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := gin.Default()
	router.Use(gin.Recovery())
	router.Use(middleware.JSONLogMiddleware())

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controller.Register)
	publicRoutes.POST("/login", controller.Login)

	protectedRoutes := router.Group("/v1")
	protectedRoutes.Use(middleware.JWTAuthMiddleware())
	protectedRoutes.GET("/user/:ID", controller.UserGetProfiles)
	protectedRoutes.POST("/user/:ID/assign_roles", controller.UserAssignRole)
	protectedRoutes.POST("/user/:ID/assign_corporate", controller.UserAssignCorporation)
	protectedRoutes.POST("/user/:ID/assign_role_app", controller.UserAssignRoleApplication)
	protectedRoutes.PATCH("/user/user_account/:ID/edit", controller.UserAccountUpdate)
	//Corporation
	protectedRoutes.POST("/corporation", controller.CorporationAddNew)
	protectedRoutes.PATCH("/corporation/:ID/edit", controller.CorporationUpdate)

	protectedLogRoutes := router.Group("/log")
	protectedLogRoutes.Use(middleware.JWTAuthMiddleware())

	if err := router.Run(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
