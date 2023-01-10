package controller

import (
	"net/http"
	"userservice/helper"
	"userservice/model"
	"userservice/schema"

	"github.com/gin-gonic/gin"
)

func UserAssignRole(context *gin.Context) {
	id := context.Param("ID")
	var input schema.AssignRole

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
	roleMap := []schema.Role{}
	for _, element := range data_roles {
		roleMap = append(roleMap, schema.Role{Id: element.Id})
	}
	users, err := model.FindUserById(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"msg": "User Not Found"})
		return
	}
	updateRoles, err := users.UserAssignRoles(id, roleMap)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"user": updateRoles})
}

func UserGetProfiles(context *gin.Context) { // Get model if exist
	id := context.Param("ID")
	user, err := model.FindUserById(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": user})
}

func UserAssignCorporation(context *gin.Context) { // Get model if exist
	id := context.Param("ID")
	var input schema.AssignCorporation

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input_corporation := input.Corporation
	data_roles, err := model.FindCorporationMapById(input_corporation)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	corpMap := []schema.Corporation{}
	for _, element := range data_roles {
		corpMap = append(corpMap, schema.Corporation{Id: element.Id})
	}
	profiles, err := model.FindProfileByUserId(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"msg": "User Not Found"})
		return
	}
	updateCorporations, err := profiles.ProfileAssignCorporation(profiles.Id, corpMap)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	users, err := model.FindUserById(updateCorporations.UserId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"user": users})
}

func UserAssignRoleApplication(context *gin.Context) { // Get model if exist
	id := context.Param("ID")
	var input schema.AssignRoleApplication

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input_role_application := input.RoleApplication
	data_roles, err := model.FindRoleApplicationMapByName(input_role_application)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	roleMap := []schema.RoleApplication{}
	for _, element := range data_roles {
		roleMap = append(roleMap, schema.RoleApplication{Id: element.Id})
	}
	accounts, err := model.FindAccountByUserId(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"msg": "User Not Found"})
		return
	}
	updateRoleApplication, err := accounts.AccountAssignRoleApplication(accounts.Id, roleMap)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	users, err := model.FindUserById(updateRoleApplication.UserId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"user": users})
}

func UserAccountUpdate(context *gin.Context) {
	// Get model if exist
	id := context.Param("ID")
	data_entries, err := model.FindAccountById(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input model.UpdateAccount
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedEntry, err := data_entries.ChangeData(id, input)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": updatedEntry})
}

func UserProfileAddNew(context *gin.Context) {
	var input model.Profile
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	unique, _ := helper.GenerateProfileId(3)
	input.Id = unique
	savedEntry, err := input.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": savedEntry})
}
