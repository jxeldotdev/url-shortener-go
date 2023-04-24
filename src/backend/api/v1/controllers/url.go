package controllers

import (
	"encoding/json"
	"fmt"
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

	correctInput := models.Url{
		LongUrl:  input.LongUrl,
		ShortUrl: helper.GenUniqueShortUrl(),
		UserId:   user.Id,
	}

	ci, _ := json.Marshal(correctInput)
	fmt.Printf("Input: %s", ci)
	savedUrl, err := correctInput.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("Saved URL: %%v: %s", string(ci))
	context.JSON(http.StatusCreated, gin.H{"data": savedUrl})
}

func GetAllUrls(context *gin.Context) {
	var url models.Url
	urls, err := url.GetAll()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// return an empty array instead of null if there's nothing
	if urls == nil {
		urls = []models.Url{}
	}

	context.JSON(http.StatusOK, gin.H{"data": urls})
}

func RedirectToLongUrl(context *gin.Context) {
	shortUrlId := context.Param("shortUrl")
	fmt.Printf("Short URL: %s", shortUrlId)
	url, err := models.FindUrlByShortUrl(shortUrlId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("Redirecting TO: %s", url.LongUrl)
	context.Redirect(http.StatusFound, url.LongUrl)
}
