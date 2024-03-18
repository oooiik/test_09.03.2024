package database

import (
	"context"
	"database/sql"
	"github.com/oooiik/test_09.03.2024/internal/logger"

	_ "github.com/lib/pq"
)

type Interface interface {
	Conn() *sql.Conn
	Close()
}

type database struct {
	conn *sql.Conn
}

func New(driver string, dataSourceName string) Interface {
	logger.Debug("database@New", driver)

	open, err := sql.Open(driver, dataSourceName)
	if err != nil {
		logger.Fatal(err)
	}
	defer open.Close()

	ctx := context.Background()

	conn, err := open.Conn(ctx)
	if err != nil {
		logger.Fatal(err)
	}

	db := database{
		conn: conn,
	}

	return &db
}

func (d *database) Conn() *sql.Conn {
	return d.conn
}

func (d *database) Close() {
	d.conn.Close()
}
