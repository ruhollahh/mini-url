package urlhandler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/url"
)

func (h *Handler) Create(c echo.Context) error {
	originalURL := c.FormValue("url")
	if originalURL == "" {
		return c.String(http.StatusBadRequest, "URL is required")
	}
	parsedURL, err := url.Parse(originalURL)
	if err != nil {
		return c.String(http.StatusUnprocessableEntity, "URL must be valid")
	}

	shortenedURL, err := h.urlSvc.CreateShortenedURL(parsedURL)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	resp := fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>URL Shortened</title>
</head>
<body>
	<h1>URL Successfully Shortened!</h1>
	<p class="shortened-url">
		<a href="%s" target="_blank">%s</a>
	</p>
	<button onclick="window.location.href='/'">Shorten Another URL</button>
</body>
</html>`, shortenedURL, shortenedURL)

	return c.HTML(http.StatusOK, resp)
}
