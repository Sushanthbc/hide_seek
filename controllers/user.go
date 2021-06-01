package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/sushanthbc/hide_seek/models"
	"golang.org/x/crypto/bcrypt"
)

type create_user struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Email     string `json:"email" binding:"required"`
}

type authenticate_user struct {
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
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

	var createParams create_user

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

	// make sure kick in unique validation
	if res.Error != nil {
		err_message := fmt.Sprintf("Could not create the user: %v", err.Error())
		logrus.Errorf(err_message)
		c.JSON(http.StatusNotFound, err_message)
		return
	}

	c.JSON(http.StatusCreated, "Created")
}

func AuthenticateUser(c *gin.Context) {
	var (
		authenticateUser authenticate_user
		user             models.User
	)

	if err := c.ShouldBindJSON(&authenticateUser); err != nil {
		err_message := fmt.Sprintf("Validation failed: %v", err.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"error": err_message,
		})
		return
	}

	res := models.DB.First(&user, "email = ?", authenticateUser.Email)

	if err := res.Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "Invalid Credentials, please try again",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword(user.PasswordHash, ([]byte)(authenticateUser.Password))

	if err != nil {
		logrus.Errorf("Invalid user name or password: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "Invalid Credentials, please try again",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Authenticated Successfully",
	})

}
