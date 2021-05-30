package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/sushanthbc/hide_seek/models"
	"golang.org/x/crypto/bcrypt"
)

type CreateUser struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Email     string `json:"email" binding:"required"`
}

func AllUsers(c *gin.Context) {
	var users []models.User

	res := models.DB.Find(&users)

	if res.Error != nil {
		logrus.Errorf("Could not fetch the records: %v", res.Error)
		c.JSON(http.StatusNotFound, "No Records Found")
	}

	c.JSON(http.StatusOK, users)
}

func CreateNewUser(c *gin.Context) {

	var createParams CreateUser

	if err := c.ShouldBindJSON(&createParams); err != nil {
		err_message := fmt.Sprintf("Validation failed: %v", err.Error())
		logrus.Errorf(err_message)
		c.JSON(http.StatusNotFound, gin.H{
			"error": err_message,
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(createParams.Password), bcrypt.DefaultCost)

	if err != nil {
		err_message := fmt.Sprintf("Could not create the user: %v", err.Error())
		logrus.Errorf(err_message)
		c.JSON(http.StatusNotFound, gin.H{
			"error": err_message,
		})
		return
	}

	user := models.User{
		FirstName:    createParams.FirstName,
		LastName:     createParams.LastName,
		PasswordHash: hash,
		Email:        createParams.Email,
	}

	res := models.DB.Create(&user)

	// make sure kick in unique valdiation
	if res.Error != nil {
		err_message := fmt.Sprintf("Could not create the user: %v", err.Error())
		logrus.Errorf(err_message)
		c.JSON(http.StatusNotFound, err_message)
		return
	}

	c.JSON(http.StatusCreated, "Created")
}
