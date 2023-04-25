package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jxeldotdev/url-backend/api/v1/controllers"
	"github.com/jxeldotdev/url-backend/config"
	"github.com/jxeldotdev/url-backend/internal/db"
	"github.com/jxeldotdev/url-backend/middleware"
	"github.com/jxeldotdev/url-backend/models"
	"github.com/stretchr/testify/assert"
)

var AppLogger *log.Logger
var cfg config.Config
var testUser string = "joelf"

func main() {
	config.PopulateConfigFromEnv(&cfg)
	r := serveApplication()
	r.Run()
}

func serveApplication() *gin.Engine {
	router := gin.Default()
	gin.SetMode(gin.TestMode)
	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controllers.Register)
	publicRoutes.POST("/login", controllers.Login)

	protectedRoutes := router.Group("/api/v1")
	protectedRoutes.Use(middleware.JWTAuthMiddleware())
	protectedRoutes.POST("/url", controllers.AddUrl)
	protectedRoutes.GET("/url", controllers.GetAllUrls)
	protectedRoutes.GET("/url/:id", controllers.GetSingleUrl)
	protectedRoutes.DELETE("/url/:id", controllers.DeleteUrl)
	protectedRoutes.GET("/user", controllers.GetAllUsers)
	return router
}

func TestInitDb(t *testing.T) {
	config.PopulateConfigFromEnv(&cfg)

	// Spin up database
	db.Connect(&cfg)

	db.Database.AutoMigrate(&models.User{})
	db.Database.AutoMigrate(&models.Url{})
}

func TestRegister(t *testing.T) {
	testRouter := serveApplication()
	var registerInput models.AuthenticationInput
	registerInput.Username = "joelfa"
	registerInput.Password = "testing123"

	fmt.Println("TEST")

	data, _ := json.Marshal(registerInput)

	req, err := http.NewRequest("POST", "/auth/register", bytes.NewBufferString(string(data)))

	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		fmt.Println(err)
	}

	resp := httptest.NewRecorder()

	testRouter.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestRegisterInvalidUsername(t *testing.T) {
	testRouter := serveApplication()
	var registerInput models.AuthenticationInput
	registerInput.Username = "joelfa"
	registerInput.Password = "wadojawkidasd!!4"

	fmt.Println("TEST")

	data, _ := json.Marshal(registerInput)

	req, err := http.NewRequest("POST", "/auth/register", bytes.NewBufferString(string(data)))

	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		fmt.Println(err)
	}

	resp := httptest.NewRecorder()

	testRouter.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

func TestLogin(t *testing.T) {
	testRouter := serveApplication()
	var authInput models.AuthenticationInput

	authInput.Username = "joelfa"
	authInput.Password = "testing123"

	data, _ := json.Marshal(authInput)

	req, err := http.NewRequest("POST", "/auth/login", bytes.NewBufferString(string(data)))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		fmt.Println(err)
	}
	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var response struct {
		Jwt string `json:"jwt"`
	}

	json.Unmarshal(body, &response)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestInvalidLogin(t *testing.T) {
	testRouter := serveApplication()
	var authInput models.AuthenticationInput

	authInput.Username = "joelfa"
	authInput.Password = "testing123aaaadad"

	data, _ := json.Marshal(authInput)

	req, err := http.NewRequest("POST", "/auth/login", bytes.NewBufferString(string(data)))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		fmt.Println(err)
	}
	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var response struct {
		Jwt string `json:"jwt"`
	}

	json.Unmarshal(body, &response)
	assert.Equal(t, http.StatusNotAcceptable, resp.Code)
}
