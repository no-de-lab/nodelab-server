package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	log "github.com/sirupsen/logrus"
)

func Migrate(steps int, dsn string) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.WithError(err).Fatal("failed to open mysql connection")
		return
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.WithError(err).Fatal("failed to create mysql driver")
		return
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations/",
		"mysql",
		driver,
	)
	if err != nil {
		log.WithError(err).Fatal("failed to create migrate instance")
		return
	}

	if err = m.Steps(steps); err != nil {
		log.WithError(err).Fatal(fmt.Sprintf("failed to migrate %d steps", steps))
	}
}
