package postgres

import (
	"fmt"
	"os"

	"github.com/Fadhelbulloh/api-with-redis/model/param"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresClient struct {
	db  *gorm.DB
	err error
}

func PostgresConn() *postgresClient {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Shanghai",
		os.Getenv("HOST"), os.Getenv("USERPG"), os.Getenv("PASSWORD"), os.Getenv("DB"), os.Getenv("PGPORT"), os.Getenv("SSLMODE"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return &postgresClient{db: db, err: err}
}

func (pg *postgresClient) Get(user_m param.UserM) (*gorm.DB, error) {
	if pg.err != nil {
		return nil, pg.err
	}

	search := pg.db.Find(&user_m)
	if search.Error != nil {
		return nil, search.Error
	}
	return search, nil
}
