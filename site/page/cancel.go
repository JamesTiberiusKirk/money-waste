package page

import (
	"github.com/labstack/echo/v4"
)

const (
	cancelPageURI = "/canceled"
)

type CancelPage struct {
}

type CancelPageData struct {
	Message string
}

func NewCancelPage() *Page {
	deps := &CancelPage{}

	return &Page{
		MenuID:      "cancelPage",
		Path:        cancelPageURI,
		Frame:       true,
		Template:    "message.gohtml",
		Deps:        deps,
		GetPageData: deps.GetPageData,
	}
}

func (l *CancelPage) GetPageData(c echo.Context) (any, error) {
	return CancelPageData{
		Message: "Canceled",
	}, nil
}
