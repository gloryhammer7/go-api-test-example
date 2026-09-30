package tests

import (
	"example.com/company/qa-api/api"
	"example.com/company/qa-api/config"

	"github.com/stretchr/testify/suite"
)

type APISuite struct {
	suite.Suite

	API *api.API
}

func (s *APISuite) SetupSuite() {
	cfg := config.Load()

	s.API = api.NewAPI(cfg.BaseURL)
}