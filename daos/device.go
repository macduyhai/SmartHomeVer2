package daos

import (
	"github.com/jinzhu/gorm"

	"github.com/macduyhai/SmartHomeVer2/models"
)

type DeviceDao interface {
	AddAdd(device models.Device) (*models.Device, error)
	// Delete(user models.User) (*models.User, error)
	// Edit(log models.Log) (*models.Log, error)
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
		return nil, err
	}
	return &device, nil
}
