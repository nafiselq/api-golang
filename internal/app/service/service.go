package service

import (
	"github.com/e-ziswaf/eziswaf-api/config"
	"github.com/e-ziswaf/eziswaf-api/internal/app/commons"
	"github.com/e-ziswaf/eziswaf-api/internal/app/repository"
)

// Option anything any service object needed
type Option struct {
	config.AppConfig
	commons.Options
	*repository.Repository
}

// Services all service object injected here
type Services struct {
	HealthCheck IHealthCheckService
	Hello       IHelloService
	Donor       IDonorService
}
