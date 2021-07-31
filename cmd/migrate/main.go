package main

import (
	"flag"
	"fmt"

	"github.com/no-de-lab/nodelab-server/db"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var (
	user     = flag.String("user", "nodelab", "username used for migration")
	pass     = flag.String("pass", "test", "password used for migration")
	host     = flag.String("host", "127.0.0.1", "host url for migration")
	database = flag.String("database", "nodelab", "database for migration")
	steps    = flag.Int("steps", 0, "number of steps to migrate")
)

func main() {
	flag.Parse()
	dsn := fmt.Sprintf("%s:%s@(%s:3306)/%s?parseTime=true&multiStatements=true", *user, *pass, *host, *database)

	db.Migrate(*steps, dsn)
}
