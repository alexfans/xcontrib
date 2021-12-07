package echohelper

import (
	"github.com/labstack/echo/v4"
	"strconv"
)

const (
	maxLimit     = 500
	defaultLimit = 20
	defaultPage  = 1
)

func GetPage(c echo.Context) (offset int, limit int) {
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = defaultLimit
	}
	if limit < 0 {
		limit = defaultLimit
	}
	if limit > 500 {
		limit = defaultLimit
	}

	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		page = defaultPage
	}
	if page <= 0 {
		page = defaultPage
	}
	offset = limit * (page - 1)
	return
}
