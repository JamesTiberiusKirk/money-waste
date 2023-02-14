package main

import (
	"fmt"

	"github.com/JamesTiberiusKirk/money-waste/models"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB(c Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=GMT",
		c.DB.Host, c.DB.User, c.DB.Pass, c.DB.Name, c.DB.Port)
	db, err := gorm.Open(postgres.Open(dsn), nil)
	if err != nil {
		logrus.WithError(err).Fatal("failed to init db")
	}

	err = db.AutoMigrate(models.GetAllModels()...)
	if err != nil {
		logrus.WithError(err).Fatal("unable to migrate db")
	}

	return db
}

func initLogger() {
	// For a json logger
	// logrus.SetFormatter(&logrus.JSONFormatter{
	// 	FieldMap: logrus.FieldMap{
	// 		logrus.FieldKeyLevel: "severity",
	// 		logrus.FieldKeyTime:  "log_time",
	// 	},
	// })
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:          true,
		DisableLevelTruncation: true,
		PadLevelText:           true,
	})
}

func initServer() *echo.Echo {
	e := echo.New()

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:       true,
		LogStatus:    true,
		LogUserAgent: true,
		LogLatency:   true,
		LogError:     true,
		LogRemoteIP:  true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			mwLog := logrus.WithFields(logrus.Fields{
				"URI":       values.URI,
				"status":    values.Status,
				"agent":     values.UserAgent,
				"latency":   values.Latency,
				"remote_ip": values.RemoteIP,
			})

			if values.Error != nil {
				mwLog.
					WithError(values.Error).
					Error("request error")
				return values.Error
			}

			mwLog.Info("request")

			return nil
		},
	}))

	e.Use(middleware.Gzip())
	e.Use(middleware.Recover())

	return e
}
