package page

import (
	"github.com/labstack/echo/v4"
)

const (
	whyPageURI = "/why"
)

type WhyPage struct {
}

type WhyPageData struct {
}

func NewWhyPage() *Page {
	deps := &WhyPage{}

	return &Page{
		MenuID:      "whyPage",
		Title:       "Why Page",
		Frame:       true,
		Path:        whyPageURI,
		Template:    "why.gohtml",
		Deps:        deps,
		GetPageData: deps.GetPageData,
	}
}

func (p *WhyPage) GetPageData(c echo.Context) (any, error) {
	return WhyPageData{}, nil
}
