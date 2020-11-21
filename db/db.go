package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	c "github.com/no-de-lab/nodelab-server/config"
)

func NewDatabase(config *c.Configuration) *sqlx.DB {
	connectionStr := fmt.Sprintf("%s:%s@(%s:3306)/%s", config.Database.Username, config.Database.Password, config.Database.Host, config.Database.Database)
	db := sqlx.MustConnect("mysql", connectionStr)

	return db
}
