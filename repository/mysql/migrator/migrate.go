package migrator

import (
	"fmt"
	"github.com/ruhollahh/mini-url/repository/mysql"
	"os"
)

func Migrate(config mysql.Config) error {
	db, err := mysql.New(config)
	if err != nil {
		return fmt.Errorf("could not open db: %w", err)
	}
	schema, err := os.ReadFile("./repository/mysql/migrator/sql/schema.sql")
	if err != nil {
		return fmt.Errorf("could not read schema: %w", err)
	}

	_, err = db.Exec(string(schema))
	if err != nil {
		return fmt.Errorf("could not execute schema: %w", err)
	}

	return nil
}
