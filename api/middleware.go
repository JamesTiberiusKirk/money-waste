package api

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	"github.com/stripe/stripe-go/v74/webhook"
)

// func sessionAuthMiddleware(sessionManager *session.Manager) echo.MiddlewareFunc {
// 	return echoSession.MiddlewareWithConfig(echoSession.Config{
// 		Skipper: func(c echo.Context) bool {
// 			skip := sessionManager.IsAuthenticated(c)
// 			if !skip {
// 				_ = c.NoContent(http.StatusUnauthorized)
// 			}
//
// 			return skip
// 		},
// 		Store: sessionManager.Jar,
// 	})
// }

func stripeSignatureMiddleware(signature string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			payload, err := io.ReadAll(c.Request().Body)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error reading request body: %v\n", err)
				return c.NoContent(http.StatusServiceUnavailable)
			}

			_, err = webhook.ConstructEvent(
				payload,
				c.Request().Header.Get("Stripe-Signature"),
				signature,
			)

			if err != nil {
				logrus.
					WithError(err).
					Error("Error verifying webhook signature")
				return c.NoContent(http.StatusBadRequest)
			}

			fmt.Println("PASSING MiddlewareFunc")
			return next(c)
		}
	}
}
