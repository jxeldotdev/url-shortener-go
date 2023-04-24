package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jxeldotdev/url-backend/helper"
	"github.com/jxeldotdev/url-backend/models"
	"gorm.io/gorm"
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

func UpdateUrl(context *gin.Context) {
	// PUT /api/v1/url/id
	var url models.Url
	urlId := context.Param("id")
	urlIdAsInt, strconvErr := strconv.ParseUint(urlId, 10, 64)
	if strconvErr != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": strconvErr.Error()})
	}
	context.BindJSON(&url)
	// need to update the url object with the req body
	url.Id = urlIdAsInt
	err := url.UpdateUrl()

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.Status(http.StatusNotFound)
			return

		}
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	context.JSON(http.StatusOK, &url)
}

func DeleteUrl(context *gin.Context) {
	// DELETE /api/v1/url/id
	var url models.Url
	urlId := context.Param("id")
	urlIdAsInt, strconvErr := strconv.ParseUint(urlId, 10, 64)
	if strconvErr != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": strconvErr.Error()})
	}
	err := url.DeleteUrl(urlIdAsInt)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.Status(http.StatusNotFound)
			return
		}
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.Status(http.StatusNoContent)
	return
}

func GetSingleUrl(context *gin.Context) {
	// GET /api/v1/url/id
	var url models.Url
	urlId := context.Param("id")
	urlIdAsInt, strconvErr := strconv.ParseUint(urlId, 10, 64)
	if strconvErr != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": strconvErr.Error()})
	}
	url, err := url.FindById(urlIdAsInt)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.Status(http.StatusNotFound)
			return
		}
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": url})
}

func GetAllUrls(context *gin.Context) {
	// GET /api/v1/url
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
	// GET /r/shortUrlId
	shortUrlId := context.Param("shortUrl")
	fmt.Printf("Short URL: %s", shortUrlId)
	url, err := models.FindUrlByShortUrl(shortUrlId)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.Status(http.StatusNotFound)
			return
		}
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("Redirecting TO: %s", url.LongUrl)
	context.Redirect(http.StatusFound, url.LongUrl)
}
