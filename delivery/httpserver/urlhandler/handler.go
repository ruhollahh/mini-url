package urlhandler

import (
	urlsvc "github.com/ruhollahh/mini-url/service/url"
)

type Handler struct {
	urlSvc *urlsvc.Service
}

func New(urlSvc *urlsvc.Service) *Handler {
	return &Handler{
		urlSvc: urlSvc,
	}
}
