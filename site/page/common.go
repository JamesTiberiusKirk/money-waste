package page

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	internalServerError = "Internal Server Error"
	invalidData         = "Unable To Validate Data"
)

type Param struct {
	Key   string
	Value string
}

func redirect(c echo.Context, uri string, params ...Param) error {
	if len(params) > 0 {
		withQuery := fmt.Sprintf("%s?", uri)

		for _, p := range params {
			withQuery = fmt.Sprintf("%s%s=%s&", withQuery, p.Key, p.Value)
		}

		return c.Redirect(http.StatusSeeOther, withQuery)
	}

	return c.Redirect(http.StatusSeeOther, uri)
}
