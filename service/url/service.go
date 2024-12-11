package urlsvc

import (
	"errors"
	"github.com/ruhollahh/mini-url/entity"
	"net/url"
)

type Config struct {
	ShortDomainURL *url.URL
	MaxRetries     int
	PostfixLen     int
}

var (
	ErrDuplicate  = errors.New("duplicate shortened postfix")
	ErrMaxRetries = errors.New("max retries reached for generating a unique shortened URL")
	ErrNotFound   = errors.New("record not found")
)

type Repo interface {
	Create(url entity.URL) (int64, error)
	GetByPostfix(string) (entity.URL, error)
}

type Service struct {
	cfg  Config
	repo Repo
}

func New(cfg Config, repo Repo) *Service {
	return &Service{
		cfg:  cfg,
		repo: repo,
	}
}
