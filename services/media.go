package services

import (
	"log"

	"github.com/macduyhai/SmartHomeVer2/middlewares"

	"github.com/macduyhai/SmartHomeVer2/config"
	"github.com/macduyhai/SmartHomeVer2/daos"
	"github.com/macduyhai/SmartHomeVer2/dtos"
	// "github.com/macduyhai/SmartHomeVer2/middlewares"
)

type MediaService interface {
	AddMedia(request dtos.AddMediaRequest) (*dtos.AddMediaResponse, error)
	ListMedia(request dtos.ListMediaRequest) (*dtos.ListMediaResponse, error)
	DeleteMedia(request dtos.DeleteMediaRequest) (*dtos.MediaResponse, error)
}

type mediaServiceImpl struct {
	config   *config.Config
	mediaDao daos.MediaDao
}

func NewMediaService(conf *config.Config, mediaDao daos.MediaDao, jwt middlewares.JWT) MediaService {
	return &mediaServiceImpl{config: conf,
		mediaDao: mediaDao,
	}
}
func (service *mediaServiceImpl) AddMedia(request dtos.AddMediaRequest) (*dtos.AddMediaResponse, error) {
	err := CheckKey(request.User_ID, request.Key)
	if err != nil {
		return nil, err
	}
	log.Println(request)
	//---------------------
	response, err := service.mediaDao.Add(request)
	if err != nil {
		log.Println("Service media errr")
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

func (service *mediaServiceImpl) ListMedia(request dtos.ListMediaRequest) (*dtos.ListMediaResponse, error) {
	err := CheckKey(request.User_ID, request.Key)
	if err != nil {
		log.Println(err)
		return nil, err
	} else {
		// log.Println("Key okie")
	}
	medias, err := service.mediaDao.List(request.User_ID)
	if err != nil {
		return nil, err
	}
	response := dtos.ListMediaResponse{
		User_ID: request.User_ID,
		Medias:  medias,
	}
	return &response, nil
}

func (service *mediaServiceImpl) DeleteMedia(request dtos.DeleteMediaRequest) (*dtos.MediaResponse, error) {
	err1 := CheckKey(request.User_ID, request.Key)
	if err1 != nil {
		return nil, err1
	}
	_, err := service.mediaDao.Delete(request.ID, request.User_ID)
	if err != nil {
		return nil, err
	}
	response := dtos.MediaResponse{
		Status: "deleted",
	}
	return &response, nil
}
