package controller

import (
	"net/http"
	"userservice/helper"
	"userservice/model"

	"github.com/gin-gonic/gin"
)

func CorporationAddNew(context *gin.Context) {
	var input model.Corporation
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	unique, _ := helper.GenerateCorporationId(3)
	input.Id = unique

	savedEntry, err := input.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"data": savedEntry})
}

func CorporationUpdate(context *gin.Context) {
	// Get model if exist
	id := context.Param("ID")
	data_entries, err := model.FindCorporateById(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input model.UpdateCorporation
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
