package daos

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/macduyhai/SmartHomeVer2/models"
)

type DeviceDao interface {
	Add(device models.Device) (*models.Device, error)
	List(userID int64, username string) ([]models.Device, error)
	// Delete(user models.User) (*models.User, error)
	Edit(log models.Log) (*models.Log, error)
	// TurnOn(userID int64, begin *time.Time, end *time.Time) ([]models.Log, error)
	// TurnOff(userName string) (*models.User, error)
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
func (dao *deviceDaoImpl) Edit(userID int64, username, chip_id, name, typedv string) (*models.Device, error) {
	device := models.Device{}
	if err := dao.db.Where("user_id = ? AND chip_id =?", userID, chip_id).Find(&device).Error; err != nil {
		return nil, err
	}
	if device.Name != name {
		device.Name = name
	}
	if device.Type != typedv {
		device.Type = typedv
	}
	return &device, nil
}
