package page

import (
	"strconv"

	stripeconnector "github.com/JamesTiberiusKirk/money-waste/connectors/stripe_connector"
	"github.com/JamesTiberiusKirk/money-waste/models"
	echo "github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const (
	homePageURI     = "/"
	metaDataMessage = "message"
)

type HomePage struct {
	db *gorm.DB
	sc *stripeconnector.StripeConnector
}

type HomePageData struct {
	WasteageAmount int64
}

func NewHomePage(db *gorm.DB, sc *stripeconnector.StripeConnector) *Page {
	deps := &HomePage{
		db: db,
		sc: sc,
	}

	return &Page{
		MenuID:      "homePage",
		Title:       "Home Page",
		Frame:       true,
		Path:        homePageURI,
		Template:    "homepage/homepage.gohtml",
		Deps:        deps,
		GetPageData: deps.GetPageData,
		PostHandler: deps.Post,
	}
}

func (p *HomePage) GetPageData(c echo.Context) (any, error) {
	count := int64(0)
	result := p.db.Model(&models.Transaction{}).Count(&count)
	if result.Error != nil {
		return nil, result.Error
	}

	return HomePageData{
		WasteageAmount: count,
	}, nil
}

func (p *HomePage) Post(c echo.Context) error {
	amount, err := strconv.ParseFloat(c.FormValue("amount"), 64)
	if err != nil {
		msg := "failed to convert to float64"
		logrus.
			WithError(err).
			Error(msg)

		return redirect(c, homePageURI, Param{Key: "error", Value: internalServerError})
	}

	message := models.Transaction{
		Message: c.FormValue("message_text"),
		Amount:  amount,
	}

	logrus.Printf("Message: %+v\n", message)

	ss, err := p.sc.StartCheckoutSession(stripeconnector.Product{
		Description: "Wasting money in the worse way you can.",
		Name:        "Money Waste",
		Price:       message.Amount,
		MetaData: map[string]string{
			metaDataMessage: message.Message,
		},
	})
	if err != nil {
		msg := "failed to start stripe session"
		logrus.
			WithError(err).
			Error(msg)

		return redirect(c, homePageURI, Param{Key: "error", Value: internalServerError})
	}

	// result := p.db.WithContext(c.Request().Context()).Create(&message)
	// if result.Error != nil {
	// 	msg := "failed to insert user into db"
	// 	logrus.
	// 		WithError(result.Error).
	// 		Error(msg)
	//
	// 	return redirect(c, homePageURI, Param{Key: "error", Value: internalServerError})
	// }

	return redirect(c, ss.URL)
	// return redirect(c, homePageURI)
}
