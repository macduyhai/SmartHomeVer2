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
	List(request dtos.ListRequest) (*dtos.ListResponse, error)
	Delete(request dtos.DeleteRequest) (*dtos.DeviceResponse, error)
	Edit(request dtos.EditRequest) (*dtos.EditResponse, error)
	Control(request dtos.ControlRequest) (*dtos.ControlResponse, error)
	Getstatus(request dtos.GetstatusRequest) (*dtos.GetstatusResponse, error)
}

type deviceServiceImpl struct {
	config    *config.Config
	deviceDao daos.DeviceDao
}

func NewDeviceService(conf *config.Config, deviceDao daos.DeviceDao, jwt middlewares.JWT) DeviceService {
	return &deviceServiceImpl{config: conf,
		deviceDao: deviceDao,
	}
}
func (service *deviceServiceImpl) Getstatus(request dtos.GetstatusRequest) (*dtos.GetstatusResponse, error) {
	device, err := service.deviceDao.Getstatus(request.Chip_ID, request.Station_MAC)
	if (device == models.Device{}) {
		return nil, err
	}
	response := dtos.GetstatusResponse{
		Chip_ID:     device.Chip_ID,
		Station_MAC: device.Station_MAC,
		State:       device.State,
	}
	return &response, nil
}

// Control
func (service *deviceServiceImpl) Control(request dtos.ControlRequest) (*dtos.ControlResponse, error) {
	device, err := service.deviceDao.Control(request.User_ID, request.Chip_ID, request.State)
	if err != nil {
		return nil, err
	}

	if device.State == true {
		fmt.Println("Bật đèn")
		s := "{\"chip_id\":" + device.Chip_ID + "," + "\"station_mac\":" + device.Station_MAC + "," + "\"value\":\"1\"}"
		fmt.Println(s)
		PublishData(device.Chip_ID, s)
	} else {
		fmt.Println("Tắt đèn")
		s := "{\"chip_id\":" + device.Chip_ID + "," + "\"station_mac\":" + device.Station_MAC + "," + "\"value\":\"0\"}"
		fmt.Println(s)
		PublishData(device.Chip_ID, s)
	}

	response := dtos.ControlResponse{
		Chip_ID: device.Chip_ID,
		State:   device.State,
	}
	return &response, nil
}

func (service *deviceServiceImpl) Delete(request dtos.DeleteRequest) (*dtos.DeviceResponse, error) {
	_, err := service.deviceDao.Delete(request.User_ID, request.Chip_ID)
	if err != nil {
		return nil, err
	}
	response := dtos.DeviceResponse{
		Status: "deleted",
	}
	return &response, nil
}
func (service *deviceServiceImpl) Edit(request dtos.EditRequest) (*dtos.EditResponse, error) {
	device, err := service.deviceDao.Edit(request.User_ID, request.Username, request.Chip_ID, request.Name, request.Type)
	if (device == models.Device{}) {
		return nil, err
	}
	response := dtos.EditResponse{
		User_ID:  request.User_ID,
		Username: request.Username,
		Device:   device,
	}
	return &response, nil
}

func (service *deviceServiceImpl) List(request dtos.ListRequest) (*dtos.ListResponse, error) {
	devices, err := service.deviceDao.List(request.User_ID, request.Username)
	if err != nil {
		return nil, err
	}
	response := dtos.ListResponse{
		User_ID:  request.User_ID,
		Username: request.Username,
		Devices:  devices,
	}
	return &response, nil
}
func (service *deviceServiceImpl) Add(request dtos.AddRequest) (*dtos.AddResponse, error) {

	dv := models.Device{
		User_ID:         request.User_ID,
		Chip_ID:         request.Chip_ID,
		Flash_Chip_ID:   request.Flash_Chip_ID,
		IDE_Flash_Size:  request.IDE_Flash_Size,
		Real_Flash_Size: request.Real_Flash_Size,
		Soft_AP_IP:      request.Soft_AP_IP,
		Soft_AP_MAC:     request.Soft_AP_MAC,
		Station_MAC:     request.Station_MAC,
		Serial:          request.Serial,
		Name:            request.Name,
		Type:            request.Type,
		State:           false,
	}

	device, err := service.deviceDao.Add(dv)
	if err != nil {
		return nil, err
	}
	response := dtos.AddResponse{
		User_ID:     request.User_ID,
		Station_MAC: device.Station_MAC,
		Chip_ID:     device.Chip_ID,
		Name:        device.Name,
		Type:        device.Type,
		State:       device.State,
		CreateAt:    device.CreateAt,
	}

	return &response, nil
}
