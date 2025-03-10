package db

import (
	"database/sql"
	"fmt"

	"github.com/dinizgab/golang-tests/internal/config"
	_ "github.com/lib/pq"
)

type Database interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
    Ping() error
}

type databaseImpl struct {
	conn *sql.DB
}

func New(config *config.DBConfig) (Database, error) {
	db, err := sql.Open("postgres",
		fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			config.Host,
			config.Port,
			config.User,
			config.Password,
			config.DBName,
		))

	if err != nil {
		return nil, fmt.Errorf("db.New: %w", err)
	}

	return &databaseImpl{db}, nil
}

func (d *databaseImpl) Exec(query string, args ...interface{}) (sql.Result, error) {
	return d.conn.Exec(query, args...)
}

func (d *databaseImpl) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return d.conn.Query(query, args...)
}

func (d *databaseImpl) QueryRow(query string, args ...interface{}) *sql.Row {
	return d.conn.QueryRow(query, args...)
}

func (d *databaseImpl) Ping() error {
    return d.conn.Ping()
}
