package models

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var CTX = context.Background()

func ConnectDatabase() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return rdb
}

// GET /feature_flags -> fetchAllFeatureFlags
// GET /feature_flags/:unique_key
//POST /feature_flags/:unique_key
//{
//   "staging": "true",
//   "production": "false",
//}
// "feat_toggle_button" : {
//     "staging": "true",
//     "production": "false",
//     "development": "true"
// }
