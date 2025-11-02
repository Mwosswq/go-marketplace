package postgres

import (
	"database/sql"

	"go.uber.org/zap"
)

func New(logger *zap.Logger, connStr string) *sql.DB {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logger.Fatal("Error while connecting database: ", zap.Error(err))
	}

	if err := db.Ping(); err != nil {
		logger.Fatal("Connection failed: ", zap.Error(err))
	}
	return db
}
