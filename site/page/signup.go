package page

import (
	"fmt"

	"github.com/JamesTiberiusKirk/money-waste/models"
	"github.com/JamesTiberiusKirk/money-waste/session"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const (
	signupPageURI = "/signup"
)

type SignupPage struct {
	db             *gorm.DB
	sessionManager *session.Manager
	signupSecret   string
}

type SignupPageData struct {
	Message    string
	Validation struct {
		models.User
		RepeatPassword string
	}
	Misc string
}

func NewSignupPage(db *gorm.DB, sessionManager *session.Manager, signupSecret string) *Page {
	deps := &SignupPage{
		db:             db,
		sessionManager: sessionManager,
		signupSecret:   signupSecret,
	}

	return &Page{
		MenuID:      "signupPage",
		Title:       "Signup",
		Frame:       true,
		Path:        "/signup",
		Template:    "signup.gohtml",
		Deps:        deps,
		GetPageData: deps.GetPageData,
		PostHandler: deps.PostHandler,
	}
}

func (p *SignupPage) GetPageData(c echo.Context) (any, error) {
	data := SignupPageData{
		Message: c.QueryParam("message"),
	}

	data.Validation.Email = c.QueryParam("email")
	data.Validation.Username = c.QueryParam("username")
	data.Validation.Password = c.QueryParam("password")
	data.Validation.RepeatPassword = c.QueryParam("repeat_password")

	data.Misc = fmt.Sprintf("%+v", data.Validation)

	return data, nil
}

func (p *SignupPage) PostHandler(c echo.Context) error {
	user := struct {
		models.User
		RepeatPassword string
		SignupSecret   string
	}{
		User: models.User{
			Email:    c.FormValue("email"),
			Username: c.FormValue("username"),
			Password: c.FormValue("password"),
		},
		RepeatPassword: c.FormValue("repeat_password"),
		SignupSecret:   c.FormValue("admin_secret"),
	}

	notPassed, err := user.Validate()
	if err != nil {
		logrus.
			WithError(err).
			Error("Error validating signup form")
		return redirect(c, signupPageURI, Param{Key: "error", Value: internalServerError})
	}

	if user.Password != user.RepeatPassword {
		notPassed = append(notPassed, "repeat_password")
	}

	if user.SignupSecret != p.signupSecret {
		notPassed = append(notPassed, "admin_secret")
	}

	if len(notPassed) > 0 {
		params := []Param{
			{
				Key:   "error",
				Value: invalidData,
			},
		}
		for _, fields := range notPassed {
			params = append(params, Param{Key: fields, Value: "not valid"})
		}
		return redirect(c, signupPageURI, params...)
	}

	err = user.SetPassword(user.Password)
	if err != nil {
		return redirect(c, signupPageURI, Param{Key: "error", Value: internalServerError})
	}

	result := p.db.WithContext(c.Request().Context()).Create(&user.User)
	if result.Error != nil {
		msg := "failed to insert user into db"
		logrus.
			WithError(result.Error).
			Error(msg)

		return redirect(c, signupPageURI, Param{Key: "error", Value: internalServerError})
	}

	p.sessionManager.InitSession(user.User, c)
	return redirect(c, adminPageURI)
}
