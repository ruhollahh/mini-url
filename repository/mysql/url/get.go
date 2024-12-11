package urlrepo

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/ruhollahh/mini-url/entity"
	urlsvc "github.com/ruhollahh/mini-url/service/url"
	"net/url"
)

func (r *Repository) GetByPostfix(shortPostfix string) (entity.URL, error) {
	query := `SELECT id, original_url, short_postfix FROM urls WHERE short_postfix = ?`
	var originalURL string
	var urlRecord entity.URL
	err := r.db.QueryRow(query, shortPostfix).Scan(&urlRecord.ID, &originalURL, &urlRecord.ShortPostfix)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.URL{}, urlsvc.ErrNotFound // No matching URL found
		}

		return entity.URL{}, fmt.Errorf("failed to retrieve url: %w", err)
	}
	urlRecord.OriginalURL, err = url.Parse(originalURL)
	if err != nil {
		return entity.URL{}, fmt.Errorf("parse: %w", err)
	}

	return urlRecord, nil
}
