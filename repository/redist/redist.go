package redisRepo

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/kataras/golog"
)

type redisClient struct {
	client *redis.Client
}

// Open new Redis Client
func NewClient() *redisClient {

	client := redis.NewClient(&redis.Options{
		Addr:     "",
		Password: "",
		DB:       0,
	})
	return &redisClient{client: client}
}

// GetData get data from redis if gin is in release mode
func (rds *redisClient) Get(criteriaKey string) (*redis.Client, string) {
	var (
		redisResult string
		err         error
	)

	if !gin.IsDebugging() {
		redisResult, err = rds.client.Get(criteriaKey).Result()
		if err != nil {
			golog.Warn("REDIS EMPTY")
		}
		fmt.Printf("\ncriteriaKey: %v\n", criteriaKey)

	}

	return rds.client, redisResult
}

// Insert data to redis with expiration time in minutes if gin is in release mode, return true if there's error
func (rds *redisClient) Set(duration time.Duration, hashQuery string, response interface{}) bool {
	if !gin.IsDebugging() {
		rs, err := json.Marshal(response)
		if err != nil {
			golog.Error(err)
			return true
		}

		minute := duration * time.Minute
		err = rds.client.Set(hashQuery, string(rs), minute).Err()
		if err != nil {
			golog.Error(err)
			return true
		}

	}
	return false
}

// Insert data to redis with no expiration time if gin is in release mode, return true if there's error
func (rds *redisClient) SetUnlimited(hashQuery string, response interface{}) bool {
	if !gin.IsDebugging() {
		rs, err := json.Marshal(response)
		if err != nil {
			golog.Error(err)
			return true
		}

		err = rds.client.Set(hashQuery, string(rs), 0).Err()
		if err != nil {
			golog.Error(err)
			return true
		}

	}
	return false
}
