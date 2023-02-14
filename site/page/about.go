package page

import (
	"github.com/labstack/echo/v4"
)

const (
	aboutPageURI = "/about"
)

type AboutPage struct {
}

type AboutPageData struct {
}

func NewAboutPage() *Page {
	deps := &AboutPage{}

	return &Page{
		MenuID:      "aboutPage",
		Title:       "About Page",
		Frame:       true,
		Path:        aboutPageURI,
		Template:    "about.gohtml",
		Deps:        deps,
		GetPageData: deps.GetPageData,
	}
}

func (p *AboutPage) GetPageData(c echo.Context) (any, error) {
	return AboutPageData{}, nil
}
