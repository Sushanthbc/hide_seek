package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sushanthbc/hide_seek/controllers"
	"github.com/sushanthbc/hide_seek/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/users", func(c *gin.Context) {
		controllers.AllUsers(c)
	})

	r.POST("/users", func(c *gin.Context) {
		controllers.CreateNewUser(c)
	})

	r.POST("/authenticate", func(c *gin.Context) {
		controllers.AuthenticateUser(c)
	})

	r.Run()
}
