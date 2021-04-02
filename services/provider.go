package services

import (
	"github.com/jinzhu/gorm"
	"github.com/macduyhai/SmartHomeVer2/config"
	"github.com/macduyhai/SmartHomeVer2/daos"
	"github.com/macduyhai/SmartHomeVer2/middlewares"
)

type Provider interface {
	GetUserService() UserService
	GetAnalysisService() AnalysisService
	GetAverageService() AverageService
	GetDeviceService() DeviceService
	GetMediaService() MediaService
}

type providerImpl struct {
	config *config.Config
	db     *gorm.DB
}

func (provider *providerImpl) GetMediaService() MediaService {
	mediaDao := daos.NewMediaDao(provider.db)
	jwtClient := middlewares.NewJWT(provider.config.SecretKet)
	return NewMediaService(provider.config, mediaDao, jwtClient)
}
func (provider *providerImpl) GetDeviceService() DeviceService {
	deviceDao := daos.NewDeviceDao(provider.db)
	jwtClient := middlewares.NewJWT(provider.config.SecretKet)
	return NewDeviceService(provider.config, deviceDao, jwtClient)
}

//---------------------------------------
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
	return &providerImpl{config: conf, db: db}
}
