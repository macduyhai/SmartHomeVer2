package services

import (
	"fmt"
	"log"
	"net/url"

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
	Upload(request dtos.UploadRequest) (*dtos.UploadResponse, error)
	Getstatus(request dtos.GetstatusRequest) (*dtos.GetstatusResponse, error)
	Push(request dtos.PushRequest) (*dtos.DeviceResponse, error)
}

type deviceServiceImpl struct {
	config     *config.Config
	deviceDao  daos.DeviceDao
	keyService EncryptService
}

func NewDeviceService(conf *config.Config, deviceDao daos.DeviceDao, jwt middlewares.JWT) DeviceService {
	return &deviceServiceImpl{config: conf,
		deviceDao: deviceDao,
	}
}

//--------------------------Upload video to device used MQTT -----------------------------
func (service *deviceServiceImpl) Push(request dtos.PushRequest) (*dtos.DeviceResponse, error) {
	err := service.keyService.CheckKey(request.User_ID, request.Key)
	if err != nil {
		return nil, err
	}
	deviceID := request.Device_ID
	mediaID := request.Media_ID

	device, err := service.deviceDao.Push(deviceID, mediaID)
	if err != nil {
		return nil, err
	}
	Url, err := url.Parse("http://192.168.2.9:9090/api/v1/media/download/")
	if err != nil {
		log.Fatal(err)
	}

	Url.Path += fmt.Sprint(request.User_ID) + "/" + request.Video_name

	s := "{\"mac\":" + `"` + device.Mac + `",` + "\"url\":" + `"` + Url.String() + `"}`
	log.Println(s)
	PublishData(device.Mac, s) // mqtt push device
	log.Println("Response")
	response := dtos.DeviceResponse{
		Status: "suscess",
	}
	log.Println("Return")
	return &response, nil
}

//--------------------------------------------------------------------------------------
func (service *deviceServiceImpl) Add(request dtos.AddRequest) (*dtos.AddResponse, error) {
	err := service.keyService.CheckKey(request.User_ID, request.Key)
	if err != nil {
		return nil, err
	}
	dv := models.Device{
		User_ID:     request.User_ID,
		Mac:         request.Mac,
		Device_name: request.Device_name,
		Location:    request.Location,
		Map_long:    request.Map_long,
		Map_lat:     request.Map_lat,
		Status:      1,
	}

	device, err := service.deviceDao.Add(dv)
	if err != nil {
		return nil, err
	}
	response := dtos.AddResponse{
		Mac:         device.Mac,
		Device_name: device.Device_name,
		Video_name:  device.Video_name,
		Video_size:  device.Video_size,
		Video_time:  device.Video_time,
		Status:      device.Status,
		Location:    device.Location,
		Map_long:    device.Map_long,
		Map_lat:     device.Map_lat,
		Expired:     device.Expired,
		CreateAt:    device.CreateAt,
	}

	return &response, nil
}

func (service *deviceServiceImpl) List(request dtos.ListRequest) (*dtos.ListResponse, error) {
	err := service.keyService.CheckKey(request.User_ID, request.Key)
	if err != nil {
		log.Println(err)
		return nil, err
	} else {
		// log.Println("Key okie")
	}
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
func (service *deviceServiceImpl) Delete(request dtos.DeleteRequest) (*dtos.DeviceResponse, error) {
	err1 := service.keyService.CheckKey(request.User_ID, request.Key)
	if err1 != nil {
		return nil, err1
	}
	_, err := service.deviceDao.Delete(request.User_ID, request.Mac)
	if err != nil {
		return nil, err
	}
	response := dtos.DeviceResponse{
		Status: "deleted",
	}
	return &response, nil
}

func (service *deviceServiceImpl) Edit(request dtos.EditRequest) (*dtos.EditResponse, error) {
	err := service.keyService.CheckKey(request.User_ID, request.Key)
	if err != nil {
		return nil, err
	}
	device, err := service.deviceDao.Edit(request.User_ID, request.Mac, request.Device_Name, request.Location)
	if (device == models.Device{}) {
		return nil, err
	}
	response := dtos.EditResponse{
		User_ID: request.User_ID,
		Device:  device,
	}
	return &response, nil
}
func (service *deviceServiceImpl) Getstatus(request dtos.GetstatusRequest) (*dtos.GetstatusResponse, error) {
	err := service.keyService.CheckKey(request.User_ID, request.Key)
	if err != nil {
		return nil, err
	}
	device, err := service.deviceDao.Getstatus(request.User_ID, request.Mac)
	if (device == models.Device{}) {
		return nil, err
	}
	response := dtos.GetstatusResponse{
		Mac:    device.Mac,
		Status: device.Status,
	}
	return &response, nil
}

// Upload
func (service *deviceServiceImpl) Upload(request dtos.UploadRequest) (*dtos.UploadResponse, error) {
	err := service.keyService.CheckKey(request.User_ID, request.Key)
	if err != nil {
		return nil, err
	}
	//---------------------
	response, err := service.deviceDao.Upload(request)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// if device.State == true {
	// 	fmt.Println("Bật đèn")
	// 	s := "{\"chip_id\":" + device.Chip_ID + "," + "\"station_mac\":" + device.Station_MAC + "," + "\"value\":\"1\"}"
	// 	fmt.Println(s)
	// 	PublishData(device.Chip_ID, s)
	// } else {
	// 	fmt.Println("Tắt đèn")
	// 	s := "{\"chip_id\":" + device.Chip_ID + "," + "\"station_mac\":" + device.Station_MAC + "," + "\"value\":\"0\"}"
	// 	fmt.Println(s)
	// 	PublishData(device.Chip_ID, s)
	// }

	return response, nil
}
