package service

import (
	"github.com/Fadhelbulloh/api-with-redis/model/param"
	"github.com/Fadhelbulloh/api-with-redis/model/response"
	"github.com/Fadhelbulloh/api-with-redis/repository/postgres"
)

// BasicGetFromPostgres get all service from postgres
func BasicGetFromPostgres(params param.UserM) response.Response {
	var response response.Response

	db := postgres.PostgresConn()

	var user_m param.UserM
	search, err := db.Get(user_m)
	if err != nil {
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
