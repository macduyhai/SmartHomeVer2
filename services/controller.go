package services

import (
	"fmt"

	"SmartHomeVer2/middlewares"
	"SmartHomeVer2/models"

	"SmartHomeVer2/config"
	"SmartHomeVer2/daos"
	"SmartHomeVer2/dtos"
	// "SmartHomeVer2/middlewares"
)

type ControllerService interface {
	Add(request dtos.AddRequest) (*dtos.AddResponse, error)
	List(request dtos.ListRequest) (*dtos.ListResponse, error)
	// Delete(request dtos.DeleteRequest) (*dtos.controllerResponse, error)
	// Edit(request dtos.EditRequest) (*dtos.EditResponse, error)
	// Control(request dtos.ControlRequest) (*dtos.ControlResponse, error)
	// Getstatus(request dtos.GetstatusRequest) (*dtos.GetstatusResponse, error)
}

type controllerServiceImpl struct {
	config    *config.Config
	controllerDao daos.ControllerDao
}

func NewcontrollerService(conf *config.Config, controllerDao daos.ControllerDao, jwt middlewares.JWT) ControllerService {
	return &controllerServiceImpl{config: conf,
		controllerDao: controllerDao,
	}
}

func (service *controllerServiceImpl) Add(request dtos.AddRequest) (*dtos.AddResponse, error) {
	fmt.Println("Add device")
	contrl := models.Controller{
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
		Active:        	 true,
		Number_device:	request.Number_device,
	}

	dv:= request.Devices

	// controller, err := service.controllerDao.Add(contrl)
	controller, err := service.controllerDao.Add(contrl,dv)
	if err != nil {
		return nil, err
	}
	
	response := dtos.AddResponse{
		User_ID:     request.User_ID,
		Station_MAC: controller.Station_MAC,
		Chip_ID:     controller.Chip_ID,
		Name:        controller.Name,
		Type:        controller.Type,
		Active:    controller.Active,
		Number_device: controller.Number_device,
		CreateAt:    controller.CreateAt,
	}

	return &response, nil
}

func (service *controllerServiceImpl) List(request dtos.ListRequest) (*dtos.ListResponse, error) {
	controllers, err := service.controllerDao.List(request.User_ID, request.Username)
	if err != nil {
		return nil, err
	}
	
	response := dtos.ListResponse{
		User_ID:  request.User_ID,
		Username: request.Username,
		Controllers:  controllers,
	}
	return &response, nil
}

// func (service *controllerServiceImpl)Getstatus(request dtos.GetstatusRequest) (*dtos.GetstatusResponse, error){
// 	controller, err := service.controllerDao.Getstatus(request.Chip_ID, request.Station_MAC)
// 	if (controller == models.controller{}) {
// 		return nil, err
// 	}
// 	response := dtos.GetstatusResponse{
// 		Chip_ID: controller.Chip_ID,
// 		Station_MAC:  controller.Station_MAC,
// 		State:   controller.State,
// 	}
// 	return &response, nil
// }
// // Control
// func (service *controllerServiceImpl) Control(request dtos.ControlRequest) (*dtos.ControlResponse, error){
// 	controller, err := service.controllerDao.Control(request.User_ID, request.Chip_ID,request.State)
// 	if err != nil {
// 		return nil, err
// 	}
		
// 	if controller.State == true {
// 		fmt.Println("Bật đèn")
// 		s := "{\"chip_id\":" +controller.Chip_ID + "," + "\"station_mac\":" + controller.Station_MAC + "," + "\"value\":\"1\"}"
// 		fmt.Println(s)
// 		PublishData(controller.Chip_ID, s)
// 	}else {
// 		fmt.Println("Tắt đèn")
// 		s := "{\"chip_id\":" +controller.Chip_ID + "," + "\"station_mac\":" + controller.Station_MAC + "," + "\"value\":\"0\"}"
// 		fmt.Println(s)
// 		PublishData(controller.Chip_ID, s)
// 	}

// 	response := dtos.ControlResponse{
// 		Chip_ID: controller.Chip_ID,
// 		State: controller.State,
// 	}
// 	return &response, nil
// }

// func (service *controllerServiceImpl) Delete(request dtos.DeleteRequest) (*dtos.controllerResponse, error) {
// 	_, err := service.controllerDao.Delete(request.User_ID, request.Chip_ID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	response := dtos.controllerResponse{
// 		Status: "deleted",
// 	}
// 	return &response, nil
// }
// func (service *controllerServiceImpl) Edit(request dtos.EditRequest) (*dtos.EditResponse, error) {
// 	controller, err := service.controllerDao.Edit(request.User_ID, request.Username, request.Chip_ID, request.Name, request.Type)
// 	if (controller == models.controller{}) {
// 		return nil, err
// 	}
// 	response := dtos.EditResponse{
// 		User_ID:  request.User_ID,
// 		Username: request.Username,
// 		controller:   controller,
// 	}
// 	return &response, nil
// }




