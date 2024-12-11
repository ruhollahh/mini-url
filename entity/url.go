package entity

import "net/url"

type URL struct {
	ID           int64
	OriginalURL  *url.URL
	ShortPostfix string
}
