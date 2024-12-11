package urlhandler

import (
	"errors"
	"github.com/labstack/echo/v4"
	urlsvc "github.com/ruhollahh/mini-url/service/url"
	"net/http"
)

func (h *Handler) Visit(c echo.Context) error {
	postfix := c.Param("postfix")

	originalURL, err := h.urlSvc.GetOriginalURL(postfix)
	if err != nil {
		if errors.Is(err, urlsvc.ErrNotFound) {
			return c.String(http.StatusNotFound, "Not Found")
		}

		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	return c.Redirect(http.StatusMovedPermanently, originalURL)
}
