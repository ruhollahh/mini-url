package urlrepo

import (
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/ruhollahh/mini-url/entity"
	urlsvc "github.com/ruhollahh/mini-url/service/url"
)

func (r *Repository) Create(url entity.URL) (int64, error) {
	query := `
		INSERT INTO urls (original_url, short_postfix) 
		VALUES (?, ?)
	`
	result, err := r.db.Exec(query, url.OriginalURL.String(), url.ShortPostfix)
	if err != nil {
		// Check for duplicate entry error
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return 0, urlsvc.ErrDuplicate
		}
		return 0, fmt.Errorf("failed to insert URL: %w", err)
	}

	// Get the auto-incremented ID of the new record
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to retrieve last insert ID: %w", err)
	}

	return id, nil
}
