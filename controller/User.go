package controller

import (
	"net/http"
	"userservice/helper"
	"userservice/model"
	"userservice/schema"

	"github.com/gin-gonic/gin"
)

// @Summary Assign role to a user
// @Tags User Management
// @ Assign role to a user with the given ID
// @Accept json
// @Produce json
// @Param ID path string true "User ID"
// @Param input body schema.AssignRole true "Role details"
// @Success 201 {string} string "Role assigned to user"
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "User Not Found"
// @Router /users/{ID}/assign-role [post]
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

// Summary Get profiles of a user
// @Tags User Management
// Description Get profiles of a user with the given ID
// @Accept json
// @Produce json
// @Param ID path string true "User ID"
// @Success 200 {string} string "Profiles of user retrieved"
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Record not found"
// @Router /users/{ID}/profiles [get]
func UserGetProfiles(context *gin.Context) { // Get model if exist
	id := context.Param("ID")
	user, err := model.FindUserById(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": user})
}

// @Summary Assign role application to a user's account
// @Tags User Management
// @Description Assign role application to a user's account with the given ID
// @Accept json
// @Produce json
// @Param ID path string true "User ID"
// @Param input body schema.AssignRoleApplication true "Role Application details"
// @Success 201 {string} string "Role Application assigned to user's account"
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "User Not Found"
// @Router /users/{ID}/assign-role-application [post]
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

// @Summary Update user's account information
// @Tags User Management
// @Description Update user's account information with the given ID
// @Accept json
// @Produce json
// @Param ID path string true "User ID"
// @Param input body schema.UpdateAccount true "Account details"
// @Success 200 {string} string "User's account information updated"
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Record not found"
// @Router /users/{ID}/account [put]
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

// @Summary Add new profile for a user
// @Tags User Management
// @Description Add new profile for a user
// @Accept json
// @Produce json
// @Param input body schema.Profile true "Profile details"
// @Success 201 {string} string "New profile added for user"
// @Failure 400 {string} string "Bad Request"
// @Router /users/profiles [post]
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
