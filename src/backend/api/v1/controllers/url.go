package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jxeldotdev/url-backend/helper"
	"github.com/jxeldotdev/url-backend/models"
)

func AddUrl(context *gin.Context) {
	var input models.Url
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := helper.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.UserId = user.ID

	savedUrl, err := input.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": savedUrl})

}

func GetAllUrls(context *gin.Context) {
	user, err := helper.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": user.Urls})
}
