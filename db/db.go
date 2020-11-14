package db

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	c "github.com/no-de-lab/nodelab-server/config"
	log "github.com/sirupsen/logrus"
)

var (
	DatabaseConnectionError = errors.New("database connection error")
)

func NewDatabase(config *c.Configuration) *sqlx.DB {
	connectionStr := fmt.Sprintf("%s:%s@(%s:3306)/%s", config.Database.Username, config.Database.Password, config.Database.Host, config.Database.Database)
	db, err := sqlx.Connect("mysql", connectionStr)

	if err != nil {
		log.Fatal("database connection error")
		panic(DatabaseConnectionError)
	}

	return db
}
