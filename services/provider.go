package services

import (
	"SmartHomeVer2/config"
	"SmartHomeVer2/daos"
	"SmartHomeVer2/middlewares"
	"github.com/jinzhu/gorm"
)

type Provider interface {
	GetUserService() UserService
	GetAnalysisService() AnalysisService
	GetAverageService() AverageService
}

type providerImpl struct {
	config *config.Config
	db *gorm.DB
}

func (provider *providerImpl) GetUserService() UserService {
	userDao := daos.NewUserDao(provider.db)
	jwtClient := middlewares.NewJWT(provider.config.SecretKet)
	return NewUserService(provider.config, userDao, jwtClient)
}

func (provider *providerImpl) GetAverageService() AverageService {
	userDao := daos.NewUserDao(provider.db)
	return NewAverageService(userDao)
}

func (provider *providerImpl) GetAnalysisService() AnalysisService {
	return NewAnalysisService()
}

func NewProvider(conf *config.Config, db *gorm.DB) Provider {
	return &providerImpl{config: conf,db:db}
}
