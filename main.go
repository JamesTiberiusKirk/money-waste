package main

import (
	"github.com/JamesTiberiusKirk/money-waste/api"
	stripeconnector "github.com/JamesTiberiusKirk/money-waste/connectors/stripe_connector"
	stripehandlers "github.com/JamesTiberiusKirk/money-waste/events/stripe_handlers"
	"github.com/JamesTiberiusKirk/money-waste/session"
	"github.com/JamesTiberiusKirk/money-waste/site"
)

func main() {
	initLogger()

	config := buildConfig()

	db := initDB(config)
	e := initServer()

	sessionManager := session.New()
	stripeConnector := stripeconnector.NewStripeConnector(
		config.Stripe.Public,
		config.Stripe.Secret,
		"http://localhost:3000/success",
		"http://localhost:3000/canceled",
	)
	stripeEventHandler := stripehandlers.NewConfigMap(db)

	apiServer := api.NewAPI(e.Group(config.HTTP.RootAPIPath), config.HTTP.RootAPIPath, db,
		sessionManager, stripeEventHandler)
	siteServer := site.NewSite(e, config.HTTP.RootSitePath, db, sessionManager, config.Debug,
		config.SignupSecret, stripeConnector)

	apiServer.Serve()
	apiRoutes := apiServer.GetRoutes()

	siteServer.SetRoutes("api", apiRoutes)
	siteServer.Serve()

	e.Logger.Fatal(e.Start(config.HTTP.Port))
}
