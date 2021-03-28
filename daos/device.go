package daos

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/jinzhu/gorm"

	"github.com/macduyhai/SmartHomeVer2/dtos"
	"github.com/macduyhai/SmartHomeVer2/models"
)

type DeviceDao interface {
	Add(device models.Device) (*models.Device, error)
	List(userID int64, username string) ([]models.Device, error)
	Delete(userID int64, mac string) (models.Device, error)
	Edit(userID int64, mac, deviceName, location string) (models.Device, error)
	Upload(request dtos.UploadRequest) (*dtos.UploadResponse, error)
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

func (dao *deviceDaoImpl) Upload(request dtos.UploadRequest) (*dtos.UploadResponse, error) {
	response := dtos.UploadResponse{
		User_ID: request.User_ID,
	}
	user := models.User{}

	// Cap nhat thong tin user
	if err := dao.db.Where("id = ? ", request.User_ID).Find(&user).Error; err != nil {
		log.Println()
		log.Println(err)
		return &response, err
	}

	// Check thong so video upload
	var total_size int64 = 0
	var count int64 = 0
	for _, file := range request.Files {
		count = count + 1
		if strings.ContainsAny(file.Video_name, " ") == true {
			err := errors.New("Video name don't have space character")
			return &response, err
		}
		total_size = total_size + file.Video_size
		if total_size > user.Max_size {
			err := errors.New("Total size > Max size")
			return &response, err
		}

	}
	for _, file := range request.Files {
		// Add thong tin cho user
		medias := models.Media{}
		medias.User_ID = request.User_ID
		medias.Video_name = file.Video_name
		medias.Video_size = file.Video_size
		medias.Video_time = file.Video_time
		response.Video_name = append(response.Video_name, file.Video_name)

		if err := dao.db.Create(&medias).Error; err != nil {
			fmt.Println("insert database media error")
			return &response, err
		}
	}
	user.Number_video = count
	user.Total_size = total_size
	dao.db.Save(&user)

	// for _, file := range request.Files {
	// 	device.Video_name = file.Video_name
	// 	if strings.ContainsAny(device.Video_name, " ") != true {
	// 		err := errors.New("Video name don't have space character")
	// 		return device, err
	// 	}
	// 	device.Video_size = file.Video_size
	// 	device.Video_time = file.Video_time
	// }
	// if err := dao.db.Where("user_id = ? AND mac =?", request.User_ID, request.Mac).Find(&device).Error; err != nil {
	// 	log.Println(err)
	// 	return device, err
	// }

	// if device.State != state {
	// 	device.State = state

	// }

	return &response, nil
}
func (dao *deviceDaoImpl) Getstatus(userID int64, mac string) (models.Device, error) {
	device := models.Device{}
	if err := dao.db.Where("user_id = ? AND mac =?", userID, mac).Find(&device).Error; err != nil {
		return device, err
	}
	return device, nil

}
