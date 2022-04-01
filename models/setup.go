package models

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"os"
)

var CTX = context.Background()

func ConnectDatabase() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%v:%v", os.Getenv("HIDE_SEEK_DATABASE_HOST"),
			os.Getenv(os.Getenv("HIDE_SEEK_DATABASE_PORT"))),
		Password: "",
		DB:       0,
	})

	return rdb
}
