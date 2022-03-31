package controllers

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/sushanthbc/hide_seek/models"
	"log"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context, rdb *redis.Client) {
	randomValue := rand.Int31()
	log.Printf("random value: %v", randomValue)
	statusCmd := rdb.Set(models.CTX, fmt.Sprintf("key-%d", randomValue), rand.Int31(), 0)

	log.Printf("Status Command: %v", statusCmd)
	fetchedValue := rdb.Get(models.CTX, fmt.Sprintf("key-%d", randomValue))

	c.JSON(http.StatusOK, gin.H{
		"data": fetchedValue.Val(),
	})
}
