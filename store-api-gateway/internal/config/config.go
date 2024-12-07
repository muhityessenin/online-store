package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"os"
	"path/filepath"
)

type Config struct {
	UserURL    string
	OrderURL   string
	PaymentURL string
	ProductURL string
}

func LoadConfig() (cfg Config, err error) {

	root, err := os.Getwd()
	if err != nil {
		return
	}

	err = godotenv.Load(filepath.Join(root, ".env"))
	if err != nil {
		cfg.UserURL = os.Getenv("userURL")
		cfg.OrderURL = os.Getenv("orderURL")
		cfg.PaymentURL = os.Getenv("paymentURL")
		cfg.ProductURL = os.Getenv("productURL")

		return cfg, nil
	}
	if err = envconfig.Process("", &cfg); err != nil {
		return
	}

	return
}
