package site

import (
	"html/template"
	"net/http"

	stripeconnector "github.com/JamesTiberiusKirk/money-waste/connectors/stripe_connector"
	"github.com/JamesTiberiusKirk/money-waste/server"
	"github.com/JamesTiberiusKirk/money-waste/session"
	"github.com/JamesTiberiusKirk/money-waste/site/page"
	"github.com/JamesTiberiusKirk/money-waste/site/renderer"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

const (
	siteName = "site"
)

// Site site struct with config and dependencies.
type Site struct {
	dev            bool
	rootSitePath   string
	publicPages    []*page.Page
	authedPages    []*page.Page
	notFoundPage   *page.Page
	staticFolders  map[string]string
	sessionManager *session.Manager
	echo           *echo.Echo
	frameTmpls     map[string]string
	tmplFuncs      template.FuncMap
	routes         map[string]server.RoutesMap
}

// NewSite init Site.
func NewSite(e *echo.Echo, rootSitePath string, db *gorm.DB,
	sessionManager *session.Manager, dev bool, signupSecret string, sc *stripeconnector.StripeConnector) *Site {
	return &Site{
		dev:          dev,
		rootSitePath: rootSitePath,
		publicPages: []*page.Page{
			page.NewLoginPage(db, sessionManager),
			page.NewSignupPage(db, sessionManager, signupSecret),
			page.NewHomePage(db, sc),
			page.NewWhyPage(),
			page.NewAboutPage(),
			page.NewSucessPage(),
			page.NewCancelPage(),
		},
		authedPages: []*page.Page{
			page.NewAdminPage(db),
			page.NewLogoutPage(sessionManager),
		},
		staticFolders: map[string]string{
			"/static": "site/static/",
			"/assets": "site/assets/",
		},
		notFoundPage:   page.NewNotFoundPage(),
		sessionManager: sessionManager,
		echo:           e,
		frameTmpls: map[string]string{
			"frame":    "frame.gohtml",
			"no_frame": "no_frame.gohtml",
		},
		tmplFuncs: template.FuncMap{
			"stringify": stringyfyJSON,
		},
		routes: map[string]server.RoutesMap{
			"site": {},
		},
	}
}

// Serve to start the server.
func (s *Site) Serve() {
	s.buildRenderer()

	s.mapPages(&s.publicPages)
	s.mapPages(&s.authedPages, sessionAuthMiddleware(s.sessionManager))

	// Mapping 404 page
	s.echo.GET(s.rootSitePath+s.notFoundPage.Path,
		s.notFoundPage.GetPageHandler(http.StatusNotFound, *s.sessionManager, s.routes))

	s.mapStatic()
}

// GetRoutes to get routes which have been made in the server.
func (s *Site) GetRoutes() server.RoutesMap {
	return s.routes["site"]
}

// SetRoutes which would be used in the templating engine.
func (s *Site) SetRoutes(t string, r server.RoutesMap) {
	s.routes[t] = r
}

func (s *Site) buildRenderer() {
	s.echo.Renderer = renderer.New(renderer.Config{
		Root:         "site/page/templates",
		Master:       s.frameTmpls["frame"],
		NoFrame:      s.frameTmpls["no_frame"],
		Funcs:        s.tmplFuncs,
		DisableCache: true,
	})
}

func (s *Site) mapStatic() {
	for k, v := range s.staticFolders {
		s.echo.Static(k, v)
	}
}

func (s *Site) mapPages(pages *[]*page.Page, middlewares ...echo.MiddlewareFunc) {
	for _, p := range *pages {
		route := s.rootSitePath + p.Path
		s.routes[siteName][p.MenuID] = route
	}

	for _, p := range *pages {
		route := s.rootSitePath + p.Path
		s.echo.GET(route, p.GetPageHandler(http.StatusOK, *s.sessionManager, s.routes), middlewares...)

		if p.PostHandler != nil {
			s.echo.POST(route, p.PostHandler, middlewares...)
		}

		if p.DeleteHandler != nil {
			s.echo.DELETE(route, p.DeleteHandler, middlewares...)
		}

		if p.PutHandler != nil {
			s.echo.PUT(route, p.PutHandler, middlewares...)
		}
	}
}
