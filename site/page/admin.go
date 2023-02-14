package page

import (
	"github.com/JamesTiberiusKirk/money-waste/models"
	echo "github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

const (
	adminPageURI = "/admin"
)

type AdminPage struct {
	db *gorm.DB
}

type AdminPageData struct {
	MessageAmount int
	Messages      []models.Transaction
	Error         string
}

func NewAdminPage(db *gorm.DB) *Page {
	deps := &AdminPage{
		db: db,
	}

	return &Page{
		MenuID:      "adminPage",
		Title:       "Admin Page",
		Frame:       true,
		Template:    "admin.gohtml",
		Path:        adminPageURI,
		Deps:        deps,
		GetPageData: deps.GetPageData,
	}
}

func (a *AdminPage) GetPageData(c echo.Context) (any, error) {
	messages := []models.Transaction{}

	result := a.db.Find(&messages)
	if result.Error != nil {
		return nil, result.Error
	}

	return AdminPageData{
		Messages:      messages,
		MessageAmount: len(messages),
	}, nil
}
