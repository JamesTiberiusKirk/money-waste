// Package session contains an implementation of a session manager.
package session

import (
	"github.com/JamesTiberiusKirk/money-waste/models"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

const (
	sessionName     = "session"
	sessionLifeSpan = 216000 // 1 hour
)

// Manager maintains a record of open sessions.
type Manager struct {
	Jar *sessions.CookieStore
}

// New returns an instantiated session manager.
func New() *Manager {
	authKey := securecookie.GenerateRandomKey(64)
	encryptionKey := securecookie.GenerateRandomKey(32)

	return &Manager{
		Jar: sessions.NewCookieStore(authKey, encryptionKey),
	}
}

// InitSession will store a new session or refresh an existing one.
func (m *Manager) InitSession(user models.User, c echo.Context) {
	sess, _ := m.Jar.Get(c.Request(), sessionName)

	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   sessionLifeSpan,
		HttpOnly: true,
	}

	sess.Values["email"] = user.Email
	sess.Values["username"] = user.Username
	_ = sess.Save(c.Request(), c.Response())
}

// TerminateSession will cease tracking the session for the current user.
func (m *Manager) TerminateSession(c echo.Context) {
	sess, _ := m.Jar.Get(c.Request(), sessionName)
	// MaxAge < 0 means delete imediately
	sess.Options.MaxAge = -1
	_ = sess.Save(c.Request(), c.Response())
}

// IsAuthenticated checks that a provided request is born from an active session.
// As long as there is an active session, true is returned, else false.
func (m *Manager) IsAuthenticated(c echo.Context) bool {
	sess, _ := m.Jar.Get(c.Request(), sessionName)
	return sess.Values["email"] != nil
}

// GetUser checks that a provided request is born from an active session.
// As long as there is an active session, User is returned, else empty User.
func (m *Manager) GetUser(c echo.Context) (models.User, error) {
	sess, err := m.Jar.Get(c.Request(), sessionName)
	if err != nil {
		return models.User{}, err
	}

	if sess.Values == nil {
		return models.User{}, nil
	}

	email, ok := sess.Values["email"]
	if !ok {
		return models.User{}, nil
	}

	username, ok := sess.Values["username"]
	if !ok {
		return models.User{}, nil
	}

	return models.User{
		Email:    email.(string),
		Username: username.(string),
	}, nil
}
