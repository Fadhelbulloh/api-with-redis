package route

import (
	"encoding/json"

	"github.com/Fadhelbulloh/api-with-redis/model/param"
	"github.com/Fadhelbulloh/api-with-redis/model/response"
	redisRepo "github.com/Fadhelbulloh/api-with-redis/repository/redist"
	"github.com/Fadhelbulloh/api-with-redis/service"
	"github.com/Fadhelbulloh/api-with-redis/util"
	"github.com/gin-gonic/gin"
	"github.com/kataras/golog"
)

func BasicService(router *gin.Engine) {
	router.GET("/basic/postgres", basicServicePostgres)
	router.GET("/basic/static", basicService)

}

func basicServicePostgres(c *gin.Context) {
	var (
		param    param.UserM
		response response.Response
	)

	// binding param
	if util.ErrHandler(500, c, c.Bind(&param)) {
		return
	}

	// transform param to byte
	byteParam, err := json.Marshal(&param)
	if util.ErrHandler(500, c, err) {
		return
	}

	// create key for radist
	hashQuery := util.Hashing(byteParam, "basic:search")

	// open reedis connection
	RedisCache := redisRepo.NewClient()

	// getting data from redist
	client, redisResult := RedisCache.Get(hashQuery)

	// defer Closing client after func executed
	defer client.Close()

	// checking redis
	if redisResult != "" {
		// redis is not empty, return response from redis
		if util.ErrHandler(500, c, json.Unmarshal([]byte(redisResult), &response)) {
			return
		}
		golog.Info("response from redis")

	} else {
		// redis is empty, accsessing db, return response from db
		response = service.BasicGetFromPostgres(param)
		if util.ErrorHandleResponse(c, response) {
			return
		}
		golog.Info("response from db")

		// insert response from db to redis
		if RedisCache.Set(10, hashQuery, &response) {
			return
		}
	}

	c.JSON(200, response)
}

func basicService(c *gin.Context) {
	var (
		param    param.UserM
		response response.Response
	)

	// binding param
	if util.ErrHandler(500, c, c.Bind(&param)) {
		return
	}

	// transform param to byte
	byteParam, err := json.Marshal(&param)
	if util.ErrHandler(500, c, err) {
		return
	}

	// create key for radist
	hashQuery := util.Hashing(byteParam, "basic:static")

	// open reedis connection
	RedisCache := redisRepo.NewClient()

	// getting data from redist
	client, redisResult := RedisCache.Get(hashQuery)

	// defer Closing client after func executed
	defer client.Close()

	// checking redis
	if redisResult != "" {
		// redis is not empty, return response from redis
		if util.ErrHandler(500, c, json.Unmarshal([]byte(redisResult), &response)) {
			return
		}
		golog.Info("response from redis")

	} else {
		// redis is empty, accsessing db, return response from db
		response = service.BasicGetStatic(param)
		if util.ErrorHandleResponse(c, response) {
			return
		}
		golog.Info("static response")

		// insert response  to redis
		if RedisCache.Set(10, hashQuery, &response) {
			return
		}
	}

	c.JSON(200, response)

}
