package mysql

import (
	"database/sql"
	"fmt"
)

type Config struct {
	User            string
	Password        string
	Host            string
	Port            uint
	Database        string
	MultiStatements bool
}

func New(cfg Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(%s:%d)/%s?parseTime=true&multiStatements=%t", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database, cfg.MultiStatements))
	if err != nil {
		return nil, fmt.Errorf("failed to open db: %w", err)
	}

	err = db.Ping()
	if err != nil {
		cErr := db.Close()
		if cErr != nil {
			return nil, fmt.Errorf("failed to close db: %w", cErr)
		}

		return nil, fmt.Errorf("failed to ping db: %w", err)
	}

	return db, nil
}
