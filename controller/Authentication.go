package controller

import (
	"fmt"
	"net/http"
	"strings"
	"time"
	"userservice/helper"
	"userservice/model"
	"userservice/schema"

	"github.com/gin-gonic/gin"
)

// @Summary Register a new user
// @Tags Authentication
// @Description Register a new user with the given credentials
// @Accept json
// @Produce json
// @Param input body schema.Register true "Registration details"
// @Success 201 {string} string "Registration is Completed"
// @Failure 400 {string} string "Bad Request"
// @Router /auth/register [post]
func Register(context *gin.Context) {
	var input schema.Register

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := helper.GenerateUserId(3)
	accountID, _ := helper.GenerateAccountId(3)

	user := model.User{
		Id:       userID,
		Email:    input.Email,
		Username: input.Username,
		Password: input.Password,
	}

	savedUser, err := user.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	account := model.Account{
		Id:     accountID,
		UserId: savedUser.Id,
	}

	savedAccount, err := account.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(savedUser)
	fmt.Println(savedAccount)

	context.JSON(http.StatusCreated, gin.H{"msg": "Registration is Completed"})
}

// @Summary Login with existing user credentials
// @Tags Authentication
// @Description with the given credentials
// @Accept json
// @Produce json
// @Param input body schema.Login true "Login details"
// @Success 200 {string} string "Login Successful"
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "User Not Found"
// @Failure 401 {string} string "Wrong Password"
// @Router /auth/login [post]
func Login(context *gin.Context) {
	var input schema.Login

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user model.User
	if strings.Contains(input.Username, "@") {
		// Jika input mengandung '@', artinya itu email
		foundUser, err := model.FindUserByEmail(input.Username)
		if err != nil {
			context.JSON(http.StatusNotFound, gin.H{"msg": "User Not Found"})
			return
		}
		user = foundUser
	} else {
		foundUser, err := model.FindUserByUsername(input.Username)
		if err != nil {
			context.JSON(http.StatusNotFound, gin.H{"msg": "User Not Found"})
			return
		}
		user = foundUser
	}

	err := user.ValidatePassword(input.Password)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"msg": "Wrong Password"})
		return
	}

	account, err := model.FindAccountByUserId(user.Id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"msg": "Account Not Found"})
		return
	}

	currentTime := time.Now()
	account.LastLoginAt = &currentTime

	err = model.UpdateAccount(&account)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update last login time"})
		return
	}

	jwt, err := helper.GenerateJWT(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"token": jwt})
}
