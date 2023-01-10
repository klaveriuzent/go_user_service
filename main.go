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

func main() {
	loadEnv()
	loadDatabase()
	serveApplication()
	testGetDateTime()
}

func testGetDateTime() {
	current_time := time.Now()

	// individual elements of time can
	// also be called to print accordingly
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		current_time.Year(), current_time.Month(), current_time.Day(),
		current_time.Hour(), current_time.Minute(), current_time.Second())

	// formatting time using
	// custom formats
	fmt.Println(current_time.Format("2006-01-02 15:04:05"))
	fmt.Println(current_time.Format("2006-January-02"))
	fmt.Println(current_time.Format("2006-01-02 3:4:5 pm"))

}
func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&schema.Role{})
	database.Database.AutoMigrate(&schema.User{})
	database.Database.AutoMigrate(&schema.UserRole{})
	database.Database.AutoMigrate(&schema.Profile{})
	database.Database.AutoMigrate(&schema.Application{})
	database.Database.AutoMigrate(&schema.RoleApplication{})
	database.Database.AutoMigrate(&schema.Account{})
	database.Database.AutoMigrate(&schema.Article{})
	database.Database.AutoMigrate(&schema.Address{})
	database.Database.AutoMigrate(&schema.Corporation{})
	database.Database.AutoMigrate(&schema.AddressCorporation{})
	database.Database.AutoMigrate(&schema.ProfileCorporations{})
	database.Database.AutoMigrate(&schema.ActivityLog{})

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
	//Application
	protectedRoutes.POST("/application", controller.ApplicationAddNew)
	protectedRoutes.PATCH("/application/:ID/edit", controller.ApplicationUpdate)
	//Corporation
	protectedRoutes.POST("/corporation", controller.CorporationAddNew)
	protectedRoutes.PATCH("/corporation/:ID/edit", controller.CorporationUpdate)
	// Article
	protectedRoutes.POST("/article", controller.ArticleAddNew)
	protectedRoutes.GET("/article/:ID", controller.ArticleFindById)
	protectedRoutes.GET("/article", controller.AllArticles)
	protectedRoutes.PATCH("/article/:ID/edit", controller.ArticleUpdate)

	protectedLogRoutes := router.Group("/log")
	protectedLogRoutes.Use(middleware.JWTAuthMiddleware())

	router.Run(":8000")
	fmt.Println("Server running on port 8000")
}
