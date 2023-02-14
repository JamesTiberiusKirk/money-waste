package main

import (
	"log"

	env "github.com/Netflix/go-env"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Debug bool `env:"DEBUG"`
	DB    struct {
		Name string `env:"DB_NAME"`
		Host string `env:"DB_HOST"`
		Port int    `env:"DB_PORT"`
		User string `env:"DB_USER"`
		Pass string `env:"DB_PASS"`
	}
	HTTP struct {
		Port         string `env:"HTTP_PORT"`
		RootSitePath string `env:"HTTP_ROOT_PATH"`
		RootAPIPath  string `env:"HTTP_ROOT_API_PATH"`
	}
	Stripe struct {
		Public string `env:"STRIPE_PUBLIC_KEY"`
		Secret string `env:"STRIPE_SECRET_KEY"`
	}
	SessionSecret string `env:"SESSION_SECRET"`
	SignupSecret  string `env:"SIGNUP_SECRET"`
	Extras        env.EnvSet
}

func buildConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var c Config
	es, err := env.UnmarshalFromEnviron(&c)
	if err != nil {
		logrus.
			WithError(err).
			Fatal("unable to get env variables")
	}

	c.Extras = es

	return c
}
