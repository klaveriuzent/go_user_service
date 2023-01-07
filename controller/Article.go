package controller

import (
	"net/http"
	"userservice/helper"
	"userservice/model"

	"github.com/gin-gonic/gin"
)

func ArticleAddNew(context *gin.Context) {
	var input model.Article
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := helper.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.UserId = user.Id
	unique, _ := helper.GenerateArticleId(3)
	input.Id = unique
	savedEntry, err := input.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": savedEntry})
}

func AllArticles(context *gin.Context) {
	user, err := helper.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": user.Articles})
}

func ArticleFindById(context *gin.Context) { // Get model if exist
	id := context.Param("ID")
	updatedEntry, err := model.ArticleFindById(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": updatedEntry})
}

func ArticleUpdate(context *gin.Context) {
	// Get model if exist
	id := context.Param("ID")
	data_entries, err := model.ArticleFindById(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	user, err := helper.CurrentUser(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate input
	var input model.UpdateArticle
	input.UserID = user.Id
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ddt := model.Entry{Content: input.Content}
	// database.Database.Model(&entryContent).Updates(ddt)

	updatedEntry, err := data_entries.ChangeData(id, input)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": updatedEntry})
}
