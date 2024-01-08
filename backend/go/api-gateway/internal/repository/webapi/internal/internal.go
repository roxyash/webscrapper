package internal

import (
	"webscrapper/api-gateway/internal/config"
	"webscrapper/pkg/logging"
)

type Internal struct {
	Auth
	Scrapper
}

var (
	conf   = config.GetConfig()
	logger = logging.GetLogger()
)

func New() *Internal {
	return &Internal{
		Auth:     NewAuth(),
		Scrapper: NewScrapper(),
	}
}
