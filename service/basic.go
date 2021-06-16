package service

import (
	"fmt"
	"os"

	"github.com/Fadhelbulloh/api-with-redis/model/param"
	"github.com/Fadhelbulloh/api-with-redis/model/response"
	"github.com/kataras/golog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// BasicGetFromPostgres get all service from postgres
func BasicGetFromPostgres(params param.UserM) response.Response {
	var response response.Response

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Shanghai",
		os.Getenv("HOST"), os.Getenv("USERPG"), os.Getenv("PASSWORD"), os.Getenv("DB"), os.Getenv("PGPORT"), os.Getenv("SSLMODE"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		golog.Error("failed to connect database")
		response.Failed("failed to connect database")
		return response
	}
	var user_m param.UserM
	search := db.Find(&user_m)
	if search.Error != nil {
		response.Failed("error db execution")
		return response
	}

	response.Success(user_m, 0, int(search.RowsAffected))
	return response
}

// BasicGetStatic dummy service
func BasicGetStatic(params param.UserM) response.Response {
	var response response.Response

	users := param.UserM{
		ID:          123,
		PhoneNumber: 877123456,
		FirstName:   "deni",
		Billing:     1000000,
	}

	response.Success(users, 0, 1)
	return response
}
