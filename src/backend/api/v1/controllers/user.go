package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jxeldotdev/url-backend/helper"
	"github.com/jxeldotdev/url-backend/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Register(context *gin.Context) {
	var input models.AuthenticationInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Username: input.Username,
		Password: input.Password,
	}

	// check if user already exists
	userExists, err := models.FindUserByUsername(input.Username)

	if userExists.Username != "" {

		context.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		return
	}

	savedUser, err := user.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"user": savedUser})
}

func Login(context *gin.Context) {
	var input models.AuthenticationInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.BindJSON(&input)

	// does user exist
	user, err := models.FindUserByUsername(input.Username)

	if err != nil {
		if err != gorm.ErrRecordNotFound {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
			return
		} else {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	fmt.Printf("%+v", &input)
	loginValid, err := models.UserLoginCheck(input.Username, input.Password)

	if loginValid != true && err == bcrypt.ErrMismatchedHashAndPassword {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	jwt, err := helper.GenerateJWT(user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"jwt": jwt})
}

func GetAllUsers(context *gin.Context) {
	var user models.User

	users, err := user.GetAllUsers()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": users})
}
