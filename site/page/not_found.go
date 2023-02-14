package page

import (
	"github.com/labstack/echo/v4"
)

const (
	notFoundPageURI = "/*"
)

type NotFoundPage struct {
}

type NotFoundPageData struct {
}

func NewNotFoundPage() *Page {
	deps := &NotFoundPage{}

	return &Page{
		MenuID:      "notFoundPage",
		Title:       "NotFound",
		Frame:       true,
		Path:        notFoundPageURI,
		Template:    "not_found.gohtml",
		Deps:        deps,
		GetPageData: deps.GetPageData,
	}
}

func (p *NotFoundPage) GetPageData(c echo.Context) (any, error) {
	return NotFoundPageData{}, nil
}
