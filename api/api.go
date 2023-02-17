package api

import (
	"github.com/JamesTiberiusKirk/money-waste/api/route"
	stripehandlers "github.com/JamesTiberiusKirk/money-waste/events/stripe_handlers"
	"github.com/JamesTiberiusKirk/money-waste/server"
	"github.com/JamesTiberiusKirk/money-waste/session"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

// API api struct.
type API struct {
	rootAPIPath    string
	publicRoutes   []*route.Route
	stripeWebhooks []*route.Route
	echoGroup      *echo.Group
	sessionManager *session.Manager
	routes         server.RoutesMap
	stripeSig      string
}

// NewAPI new api instance.
func NewAPI(group *echo.Group, rootAPIPath string, db *gorm.DB,
	sesessionManager *session.Manager, stripeEventHandler *stripehandlers.ConfigMap,
	stripeSig string) *API {
	return &API{
		rootAPIPath:  rootAPIPath,
		publicRoutes: []*route.Route{},
		stripeWebhooks: []*route.Route{
			route.NewStripeWebhookRoute(stripeEventHandler),
		},
		echoGroup:      group,
		sessionManager: sesessionManager,
		routes:         server.RoutesMap{},
		stripeSig:      stripeSig,
	}
}

// Serve api.
func (a *API) Serve() {
	a.echoGroup.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	a.mapRoutes(&a.publicRoutes)
	a.mapRoutes(&a.stripeWebhooks, stripeSignatureMiddleware(a.stripeSig))
}

// GetRoutes returns available routes from this server.
func (a *API) GetRoutes() server.RoutesMap {
	return a.routes
}

func (a *API) SetRoutes(t string, r server.RoutesMap) {
	// NOOP
}

func (a *API) mapRoutes(routes *[]*route.Route, middlewares ...echo.MiddlewareFunc) {
	for _, r := range *routes {
		routes := r.Init("", a.echoGroup, middlewares...)

		for k, v := range routes {
			a.routes[k] = a.rootAPIPath + v
		}
	}
}
