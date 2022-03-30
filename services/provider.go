package services

import (
	"github.com/jinzhu/gorm"
	"github.com/macduyhai/SmartHomeVer2/config"
	"github.com/macduyhai/SmartHomeVer2/daos"
	"github.com/macduyhai/SmartHomeVer2/middlewares"
)

type Provider interface {
	GetUserService() UserService
	GetDeviceService() DeviceService
	GetMediaService() MediaService
}

type providerImpl struct {
	config *config.Config
	db     *gorm.DB
}

func (provider *providerImpl) GetMediaService() MediaService {
	mediaDao := daos.NewMediaDao(provider.db)
	jwtClient := middlewares.NewJWT(provider.config.SecretKey)
	return NewMediaService(provider.config, mediaDao, jwtClient)
}
func (provider *providerImpl) GetDeviceService() DeviceService {
	deviceDao := daos.NewDeviceDao(provider.db)
	jwtClient := middlewares.NewJWT(provider.config.SecretKey)
	return NewDeviceService(provider.config, deviceDao, jwtClient)
}

//---------------------------------------
func (provider *providerImpl) GetUserService() UserService {
	userDao := daos.NewUserDao(provider.db)
	jwtClient := middlewares.NewJWT(provider.config.SecretKey)
	return NewUserService(provider.config, userDao, jwtClient)
}

func NewProvider(conf *config.Config, db *gorm.DB) Provider {
	return &providerImpl{config: conf, db: db}
}
