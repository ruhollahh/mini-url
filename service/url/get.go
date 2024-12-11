package urlsvc

import (
	"fmt"
)

func (s *Service) GetOriginalURL(shortPostfix string) (string, error) {
	urlRecord, err := s.repo.GetByPostfix(shortPostfix)
	if err != nil {
		return "", fmt.Errorf("getByPostfix: %w", err)
	}

	return urlRecord.OriginalURL.String(), nil
}
