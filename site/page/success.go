package page

import (
	"github.com/labstack/echo/v4"
)

const (
	sucessPageURI = "/success"
)

type SucessPage struct {
}

type SucessPageData struct {
	Message string
}

func NewSucessPage() *Page {
	deps := &SucessPage{}

	return &Page{
		MenuID:      "sucessPage",
		Path:        sucessPageURI,
		Frame:       true,
		Template:    "message.gohtml",
		Deps:        deps,
		GetPageData: deps.GetPageData,
	}
}

func (l *SucessPage) GetPageData(c echo.Context) (any, error) {
	return SucessPageData{
		Message: "Success!",
	}, nil
}
