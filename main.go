package main

import (
	"fmt"
	"log"
	"time"
	"userservice/controller"
	"userservice/database"
	"userservice/helper"
	"userservice/middleware"
	"userservice/schema"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	loadDatabase()
	serveApplication()
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
}

func serveApplication() {
	router := gin.Default()

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

	genArticleId, _ := helper.GenerateArticleId(3)
	fmt.Println(genArticleId)

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controller.Register)
	publicRoutes.POST("/login", controller.Login)

	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(middleware.JWTAuthMiddleware())
	protectedRoutes.GET("/user/:ID", controller.UserGetProfiles)
	protectedRoutes.POST("/user/:ID/assign_roles", controller.UserAssignRole)
	protectedRoutes.GET("/article/:ID", controller.ArticleFindById)
	protectedRoutes.POST("/article", controller.ArticleAddNew)
	protectedRoutes.GET("/article", controller.AllArticles)
	protectedRoutes.PATCH("/article/:ID/edit", controller.ArticleUpdate)

	router.Run(":8000")
	fmt.Println("Server running on port 8000")
}
