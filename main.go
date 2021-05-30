package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sushanthbc/hide_seek/models"
	"gorm.io/gorm"
)

func main() {
	r := gin.Default()

	db := models.ConnectDatabase()

	r.GET("/users", func(c *gin.Context) {
		allUsers(db, c)
	})

	r.Run()
}

func allUsers(db *gorm.DB, c *gin.Context) {
	res := models.GetAllUser(db)

	if err := res.Error; err != nil {
		c.JSON(http.StatusNotFound, fmt.Sprintf("Could not fetch all the users: %v", err))
	}

	c.JSON(http.StatusOK, res)
}
