package db

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	c "github.com/no-de-lab/nodelab-server/config"
)

const (
	maxIdleConns = 10
	maxOpenConns = 10
	maxConnLife  = 5 * time.Minute
)

func NewDatabase(config *c.Configuration) *sqlx.DB {
	connectionStr := fmt.Sprintf("%s:%s@(%s:3306)/%s?%s", config.Database.Username, config.Database.Password, config.Database.Host, config.Database.Database, "parseTime=true")
	db := sqlx.MustConnect("mysql", connectionStr)

	db.SetMaxIdleConns(maxIdleConns)
	db.SetMaxOpenConns(maxOpenConns)
	db.SetConnMaxLifetime(maxConnLife)

	return db
}
