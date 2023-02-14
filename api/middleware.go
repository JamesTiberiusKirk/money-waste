package api

import (
	"net/http"

	"github.com/JamesTiberiusKirk/money-waste/session"
	"github.com/labstack/echo/v4"

	echoSession "github.com/labstack/echo-contrib/session"
)

func sessionAuthMiddleware(sessionManager *session.Manager) echo.MiddlewareFunc {
	return echoSession.MiddlewareWithConfig(echoSession.Config{
		Skipper: func(c echo.Context) bool {
			skip := sessionManager.IsAuthenticated(c)
			if !skip {
				_ = c.NoContent(http.StatusUnauthorized)
			}

			return skip
		},
		Store: sessionManager.Jar,
	})
}
