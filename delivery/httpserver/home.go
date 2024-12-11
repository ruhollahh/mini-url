package httpserver

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func home(c echo.Context) error {
	return c.HTML(http.StatusOK, `<!DOCTYPE html>
<html>
<head>
    <title>URL Shortener</title>
</head>
<body>
    <h1>Shorten Your URL</h1>
    <form action="/urls/create" method="POST">
        <label for="url">Enter URL:</label>
        <input type="url" id="url" name="url" required placeholder="https://google.com" />
        <button type="submit">Shorten</button>
    </form>
</body>
</html>`)
}
