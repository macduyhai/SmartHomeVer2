package services

import (
	"fmt"

	"github.com/macduyhai/SmartHomeVer2/middlewares"
	"github.com/macduyhai/SmartHomeVer2/models"

	"github.com/macduyhai/SmartHomeVer2/config"
	"github.com/macduyhai/SmartHomeVer2/daos"
	"github.com/macduyhai/SmartHomeVer2/dtos"
	// "github.com/macduyhai/SmartHomeVer2/middlewares"
)

type DeviceService interface {
	Add(request dtos.AddRequest) (*dtos.AddResponse, error)
	// Delete(request dtos.DeleteRequest) (*dtos.DeviceResponse, error)
	// Edit(request dtos.EditRequest) (*dtos.EditResponse, error)
	// TurnOn(request dtos.TurnOnRequest) (*dtos.DeviceResponse, error)
	// TurnOff(request dtos.TurnOffRequest) (*dtos.DeviceResponse, error)
}

type deviceServiceImpl struct {
	config    *config.Config
	deviceDao daos.DeviceDao
}

func NewDeviceService(conf *config.Config, deviceDao daos.DeviceDao, jwt middlewares.JWT) DeviceService {
	return &deviceServiceImpl{config: conf,
		deviceDao: deviceDao,
		jwt:       jwt,
	}
}

func (service *deviceServiceImpl) Add(request dtos.AddRequest) (*dtos.AddResponse, error) {
	fmt.Println(request)
	dv := models.Device{
		Chip_ID:         request.Chip_ID,
		Flash_Chip_ID:   request.Flash_Chip_ID,
		IDE_Flash_Size:  request.IDE_Flash_Size,
		Real_Flash_Size: request.Real_Flash_Size,
		Soft_AP_IP:      request.Soft_AP_IP,
		Soft_AP_MAC:     request.Soft_AP_MAC,
		Station_MAC:     request.Station_MAC,
		Serial:          request.Serial,
		Name:            request.Name,
		LastState:       false,
		NewState:        false,
	}
	fmt.Println(dv)
	device, err := service.deviceDao.Add(dv)
	if err != nil {
		return nil, err
	}
	fmt.Println("Service 3")
	response := dtos.AddResponse{
		Station_MAC: device.Station_MAC,
		Chip_ID:     device.Chip_ID,
		Name:        device.Name,
		Type:        device.Type,
		NewState:    device.NewState,
		CreateAt:    device.CreateAt,
	}

	return &response, nil
}
