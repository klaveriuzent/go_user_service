package controller

import (
	"fmt"
	"net/http"
	"userservice/helper"
	"userservice/model"
	"userservice/schema"

	"github.com/gin-gonic/gin"
)

type RegisterSchema schema.Register
type LoginSchema schema.Login

func Register(context *gin.Context) {
	var input RegisterSchema

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input_roles := input.Role
	data_roles, err := model.FindRoleMapByName(input_roles)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID, _ := helper.GenerateUserId(3)
	profileId, _ := helper.GenerateProfileId(3)
	accountId, _ := helper.GenerateAccountId(3)

	roleMap := []schema.Role{}
	for _, element := range data_roles {
		roleMap = append(roleMap, schema.Role{Id: element.Id})
	}

	user := model.User{
		Id:       userID,
		Roles:    roleMap,
		Email:    input.Email,
		Username: input.Username,
		Password: input.Password,
	}

	savedUser, err := user.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	profile := model.Profile{
		Id:       profileId,
		Username: input.Username,
		Email:    input.Email,
		UserId:   savedUser.Id,
	}

	savedProfile, err := profile.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	account := model.Account{
		Id:            accountId,
		UserId:        savedUser.Id,
		ApplicationId: input.ApplicationId,
	}

	savedAccount, err := account.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(savedUser)
	fmt.Println(savedProfile)
	fmt.Println(savedAccount)
	context.JSON(http.StatusCreated, gin.H{"msg": "Registration is Completed"})
}

func Login(context *gin.Context) {
	var input LoginSchema

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := model.FindUserByUsername(input.Username)
	fmt.Println(&user, err)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"msg": "User Not Found"})
		return
	}

	err = user.ValidatePassword(input.Password)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"msg": "Wrong Password"})
		return
	}

	jwt, err := helper.GenerateJWT(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"token": jwt})
}
