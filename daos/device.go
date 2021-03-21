package daos

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"

	"github.com/macduyhai/SmartHomeVer2/dtos"
	"github.com/macduyhai/SmartHomeVer2/models"
)

type DeviceDao interface {
	Add(device models.Device) (*models.Device, error)
	List(userID int64, username string) ([]models.Device, error)
	Delete(userID int64, mac string) (models.Device, error)
	Edit(userID int64, mac, deviceName, location string) (models.Device, error)
	Upload(request dtos.UploadRequest) (models.Device, error)
	Getstatus(userID int64, mac string) (models.Device, error)
}

type deviceDaoImpl struct {
	db *gorm.DB
}

func NewDeviceDao(db *gorm.DB) DeviceDao {
	return &deviceDaoImpl{db: db}
}

func (dao *deviceDaoImpl) Add(device models.Device) (*models.Device, error) {

	if err := dao.db.Create(&device).Error; err != nil {
		fmt.Println("insert database error")
		return nil, err
	}
	return &device, nil
}

func (dao *deviceDaoImpl) List(userID int64, username string) ([]models.Device, error) {
	devices := make([]models.Device, 0)
	if err := dao.db.Where("user_id = ?", userID).Find(&devices).Error; err != nil {
		return nil, err
	}

	return devices, nil
}

func (dao *deviceDaoImpl) Edit(userID int64, mac, deviceName, location string) (models.Device, error) {
	device := models.Device{}
	if err := dao.db.Where("user_id = ? AND mac =?", userID, mac).Find(&device).Error; err != nil {
		return device, err
	}
	if device.Device_name != deviceName {
		device.Device_name = deviceName
	}
	if device.Location != location {
		device.Location = location
	}

	dao.db.Save(&device)

	return device, nil
}
func (dao *deviceDaoImpl) Delete(userID int64, mac string) (models.Device, error) {
	device := models.Device{}
	if err := dao.db.Where("user_id = ? AND mac =?", userID, mac).Delete(&device).Error; err != nil {
		return device, err
	}
	return device, nil
}

func (dao *deviceDaoImpl) Upload(request dtos.UploadRequest) (models.Device, error) {
	device := models.Device{}
	if err := dao.db.Where("user_id = ? AND mac =?", request.User_ID, request.Mac).Find(&device).Error; err != nil {
		log.Println(err)
		return device, err
	}
	for _, file := range request.Files {

		device.Video_name = file.Video_name
		device.Video_size = file.Video_size
		device.Video_time = file.Video_time
	}
	// if device.State != state {
	// 	device.State = state

	// }
	dao.db.Save(&device)
	return device, nil
}
func (dao *deviceDaoImpl) Getstatus(userID int64, mac string) (models.Device, error) {
	device := models.Device{}
	if err := dao.db.Where("user_id = ? AND mac =?", userID, mac).Find(&device).Error; err != nil {
		return device, err
	}
	return device, nil

}
