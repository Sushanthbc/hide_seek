package main

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/sushanthbc/hide_seek/models"
)

func main() {
	r := gin.Default()

	db := models.ConnectDatabase()

	users, err := models.GetAllUser(db)

	if err != nil {
		panic("failed to fetch")
	}

	j, _ := json.Marshal(users)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": j,
		})
	})

	r.Run()
}
