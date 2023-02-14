package site

import (
	"log"
	"net/http"

	"github.com/JamesTiberiusKirk/money-waste/session"
	"github.com/labstack/echo/v4"

	echoSession "github.com/labstack/echo-contrib/session"
)

func sessionAuthMiddleware(sessionManager *session.Manager) echo.MiddlewareFunc {
	return echoSession.MiddlewareWithConfig(echoSession.Config{
		Skipper: func(c echo.Context) bool {
			auth := sessionManager.IsAuthenticated(c)
			if !auth {
				log.Println("-----------------------------------redirecting")
				_ = c.Redirect(http.StatusSeeOther, "/login")
			}

			return auth
		},
		Store: sessionManager.Jar,
	})
}
