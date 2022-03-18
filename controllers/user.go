package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/sushanthbc/hide_seek/models"
)

func Index(c *gin.Context) {
	var users []models.User
	res := models.DB.Find(&users)

	if res.Error != nil {
		logrus.Errorf("Could not fetch the records: %v", res.Error)
		c.JSON(http.StatusNotFound, "No Records Found")
	}

	c.JSON(http.StatusOK, users)
}
