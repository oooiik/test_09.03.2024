package database

import (
	"database/sql"
	"github.com/oooiik/test_09.03.2024/internal/logger"

	_ "github.com/lib/pq"
)

type Interface interface {
	DB() *sql.DB
}

type database struct {
	db *sql.DB
}

func New(driver string, dataSourceName string) Interface {
	logger.Debug("database@New", driver)

	db, err := sql.Open(driver, dataSourceName)
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		logger.Fatal(err)
	}
	logger.Debug("database@New", driver, "Ping successful")

	n := database{
		db: db,
	}

	return &n
}

func (d *database) DB() *sql.DB {
	return d.db
}
