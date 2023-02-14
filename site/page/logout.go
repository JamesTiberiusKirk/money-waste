package page

import (
	"github.com/JamesTiberiusKirk/money-waste/session"
	"github.com/labstack/echo/v4"
)

const (
	logoutPageURI = "/logout"
)

type LogoutPage struct {
	session *session.Manager
}

func NewLogoutPage(session *session.Manager) *Page {
	deps := &LogoutPage{
		session: session,
	}

	return &Page{
		MenuID:      "logoutPage",
		Path:        logoutPageURI,
		Frame:       false,
		Template:    "logout_page.gohtml",
		Deps:        deps,
		GetPageData: deps.GetPageData,
	}
}

func (l *LogoutPage) GetPageData(c echo.Context) (any, error) {
	l.session.TerminateSession(c)
	redirect(c, loginPageURI)
	return "", nil
}
