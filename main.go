package main

import (
	"crypto/subtle"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sushanthbc/hide_seek/controllers"
	"github.com/sushanthbc/hide_seek/models"
	"net/http"
	"os"
)

func main() {
	r := gin.Default()
	err := godotenv.Load(".env")
	if err != nil {
		panic(err.Error())
	}
	models.ConnectDatabase()

	apiKeyAuth := r.Group("/")
	apiKeyAuth.Use(ApiAuthRequired())
	{
		apiKeyAuth.GET("/users", func(c *gin.Context) {
			controllers.Index(c)
		})
	}

	err = r.Run()

	if err != nil {
		panic(err.Error())
	}
}

func ApiAuthRequired() gin.HandlerFunc {
	return func(context *gin.Context) {
		actualApiKey := []byte(context.Request.Header.Get("x-api-key"))
		expectedApikey := []byte(os.Getenv("API_KEY"))
		if subtle.ConstantTimeCompare(actualApiKey, expectedApikey) == 0 {
			context.JSON(http.StatusUnauthorized, "Failed to authenticate")
			context.Abort()
			return
		}
	}
}
