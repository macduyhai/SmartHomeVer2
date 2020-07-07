package daos

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"SmartHomeVer2/models"
	"SmartHomeVer2/dtos"
)

type ControllerDao interface {
	Add(controller models.Controller,dv []dtos.Device) (*models.Controller, error)
	List(userID int64, username string) ([]models.Controller, error)
	Delete(userID int64, chip_id string) (models.Controller, error)
	// Edit(userID int64, username, chip_id, name, typedv string) (models.Controller, error)
	Control(userID int64, chip_id string, device_id string , state bool) (models.Controller,models.Device, error)
	Getstatus( chip_id string , station_mac string) (models.Controller, error)
}

type controllerDaoImpl struct {
	db *gorm.DB
}

func NewcontrollerDao(db *gorm.DB) ControllerDao {
	return &controllerDaoImpl{db: db}
}

func (dao *controllerDaoImpl) Add(controller models.Controller,dv []dtos.Device) (*models.Controller, error) {
	
	if err := dao.db.Create(&controller).Error; err != nil {
		fmt.Println("insert database controller error")
		return nil, err
	}
	for i:=int64(0);i<controller.Number_device;i++{
		device := models.Device{
			User_ID: controller.User_ID,
			Controller_id: controller.ID,
			Device_id: dv[i].Device_id,
			Name: dv[i].Name,
			Type: dv[i].Type,
			State: dv[i].State,
		}

		if err := dao.db.Create(&device).Error; err != nil {
			fmt.Println("insert database device error")
			return nil, err
		}
	}
		
	return &controller, nil
}
func (dao *controllerDaoImpl) List(userID int64, username string) ([]models.Controller, error) {
	controllers := make([]models.Controller, 0)
	if err := dao.db.Where("user_id =?", userID).Find(&controllers).Error; err != nil {
		return nil, err
	}
	// for i, c := range controllers {
	for i:=0;i<len(controllers);i++ {
		devices := make([]models.Device, 0)
		if err := dao.db.Where("controller_id =?", controllers[i].ID).Find(&devices).Error; err != nil {
			return nil, err
		}
		controllers[i].Devices =devices
		// fmt.Println("----------------------------------------0------------------------------------------------")
		// // fmt.Println(c)
		// fmt.Println(controllers[i].Devices)
		// fmt.Println("----------------------------------------1------------------------------------------------")
	}

	return  controllers, nil
}

func (dao *controllerDaoImpl) Delete(userID int64, chip_id string) (models.Controller, error) {
	controllers := models.Controller{}

	if err := dao.db.Where("user_id = ? AND chip_id =?", userID, chip_id).Find(&controllers).Error; err != nil {
		return controllers, err
	}

	devices := make([]models.Device, 0)
	if err := dao.db.Where("controller_id =?", controllers.ID).Delete(&devices).Error; err != nil {
		return controllers, err
	}

	if err := dao.db.Where("user_id = ? AND chip_id =?", userID, chip_id).Delete(&controllers).Error; err != nil {
		return controllers, err
	}
	return controllers, nil
}

// func (dao *controllerDaoImpl) Edit(userID int64, username, chip_id, name, typedv string) (models.controller, error) {
// 	controller := models.controller{}
// 	if err := dao.db.Where("user_id = ? AND chip_id =?", userID, chip_id).Find(&controller).Error; err != nil {
// 		return controller, err
// 	}
// 	if controller.Name != name {
// 		controller.Name = name
// 	}
// 	if controller.Type != typedv {
// 		controller.Type = typedv
// 	}

// 	dao.db.Save(&controller)

// 	return controller, nil
// }

func (dao *controllerDaoImpl) Control(userID int64, chip_id string , device_id string, state bool) (models.Controller,models.Device, error){
	controller := models.Controller{}
	device :=models.Device{}
	if err := dao.db.Where("user_id = ? AND chip_id =?", userID, chip_id).Find(&controller).Error; err != nil {
		return controller,device, nil
	}
	
	if err := dao.db.Where("controller_id = ? AND device_id =?", controller.ID, device_id).Find(&device).Error; err != nil {
		return controller,device, nil
	}
	if device.State != state {
		device.State = state
		
	}
	dao.db.Save(&controller)
	dao.db.Save(&device)
	return controller,device, nil
}
func (dao *controllerDaoImpl)Getstatus( chip_id string , station_mac string) (models.Controller, error){
	controller := models.Controller{}
	if err := dao.db.Where("station_mac = ? AND chip_id =?", station_mac, chip_id).Find(&controller).Error; err != nil {
		return controller, err
	}

	devices := make([]models.Device, 0)
	if err := dao.db.Where("controller_id =?", controller.ID).Find(&devices).Error; err != nil {
		return controller, err
	}
	controller.Devices =devices
	
	return controller, nil
	
}