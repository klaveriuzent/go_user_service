package controller

import (
	"net/http"
	"userservice/model"

	"github.com/gin-gonic/gin"
)

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
