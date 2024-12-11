package urlsvc

import (
	"errors"
	"fmt"
	"github.com/ruhollahh/mini-url/entity"
	"github.com/ruhollahh/mini-url/pkg/token"
	"net/url"
)

// CreateShortenedURL generates a new shortened URL and retries if a duplicate is detected
func (s *Service) CreateShortenedURL(originalURL *url.URL) (string, error) {
	for i := 0; i < s.cfg.MaxRetries; i++ {
		shortPostfix := token.Generate(s.cfg.PostfixLen)
		urlRecord := entity.URL{
			OriginalURL:  originalURL,
			ShortPostfix: shortPostfix,
		}

		_, err := s.repo.Create(urlRecord)
		if err == nil {
			shortenedURL := s.cfg.ShortDomainURL.JoinPath(shortPostfix)

			return shortenedURL.String(), nil
		}

		if !errors.Is(err, ErrDuplicate) {
			return "", fmt.Errorf("could not create url: %w", err)
		}
	}

	return "", ErrMaxRetries
}
