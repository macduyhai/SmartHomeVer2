package daos

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/macduyhai/SmartHomeVer2/models"
)

type DeviceDao interface {
	Add(device models.Device) (*models.Device, error)
	List(userID int64, username string) ([]models.Device, error)
	Delete(userID int64, chip_id string) (models.Device, error)
	Edit(userID int64, username, chip_id, name, typedv string) (models.Device, error)
	Control(userID int64, chip_id string, state bool) (models.Device, error)
	Getstatus(chip_id string, station_mac string) (models.Device, error)
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
	if err := dao.db.Where("user_id = ?AND username =?", userID, username).Find(&devices).Error; err != nil {
		return nil, err
	}

	return devices, nil
}
func (dao *deviceDaoImpl) Edit(userID int64, username, chip_id, name, typedv string) (models.Device, error) {
	device := models.Device{}
	if err := dao.db.Where("user_id = ? AND chip_id =?", userID, chip_id).Find(&device).Error; err != nil {
		return device, err
	}
	if device.Name != name {
		device.Name = name
	}
	if device.Type != typedv {
		device.Type = typedv
	}

	dao.db.Save(&device)

	return device, nil
}
func (dao *deviceDaoImpl) Delete(userID int64, chip_id string) (models.Device, error) {
	device := models.Device{}
	if err := dao.db.Where("user_id = ? AND chip_id =?", userID, chip_id).Delete(&device).Error; err != nil {
		return device, err
	}
	return device, nil
}
func (dao *deviceDaoImpl) Control(userID int64, chip_id string, state bool) (models.Device, error) {
	device := models.Device{}
	if err := dao.db.Where("user_id = ? AND chip_id =?", userID, chip_id).Find(&device).Error; err != nil {
		return device, err
	}
	if device.State != state {
		device.State = state

	}
	dao.db.Save(&device)
	return device, nil
}
func (dao *deviceDaoImpl) Getstatus(chip_id string, station_mac string) (models.Device, error) {
	device := models.Device{}
	if err := dao.db.Where("station_mac = ? AND chip_id =?", station_mac, chip_id).Find(&device).Error; err != nil {
		return device, err
	}
	return device, nil

}
