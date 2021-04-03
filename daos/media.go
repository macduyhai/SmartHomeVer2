package daos

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/macduyhai/SmartHomeVer2/dtos"
	"github.com/macduyhai/SmartHomeVer2/models"
)

type MediaDao interface {
	Add(request dtos.AddMediaRequest) (*dtos.AddMediaResponse, error)
	List(userID int64) ([]models.Media, error)
	Delete(ID int64, userID int64) (models.Media, error)
}

type mediaDaoImpl struct {
	db *gorm.DB
}

func NewMediaDao(db *gorm.DB) MediaDao {
	return &mediaDaoImpl{db: db}
}

func (dao *mediaDaoImpl) Add(request dtos.AddMediaRequest) (*dtos.AddMediaResponse, error) {

	response := dtos.AddMediaResponse{
		User_ID: request.User_ID,
	}
	user := models.User{}

	// Cap nhat thong tin user
	if err := dao.db.Where("id = ? ", request.User_ID).Find(&user).Error; err != nil {
		log.Println("Load user error")
		log.Println(err)
		return &response, err
	}

	// Check thong so video upload
	var total_size int64 = user.Total_size
	var count int64 = 0
	for _, file := range request.Files {
		count = count + 1
		if strings.ContainsAny(file.Video_name, " ") == true {
			err := errors.New("Video name no space character")
			return &response, err
		}
		log.Print("File size: ")
		log.Println(file.Video_size)
		total_size = total_size + file.Video_size
		if total_size > user.Max_size {
			err := errors.New("Total size > Max size")
			return &response, err
		}

	}
	for _, file := range request.Files {
		// Add thong tin cho Media table
		var media models.Media
		media.User_ID = request.User_ID
		media.Video_name = file.Video_name
		media.Video_size = file.Video_size
		media.Video_time = file.Video_time
		response.Video_name = append(response.Video_name, file.Video_name)
		if err := dao.db.Where("user_id = ? AND video_name = ? ", request.User_ID, file.Video_name).Find(&media).Error; err != nil {
			if err := dao.db.Create(&media).Error; err != nil {
				fmt.Println("insert database media error")
				return &response, err
			}
		} else {
			err := errors.New("File exits")
			return &response, err
		}

	}
	user.Number_video = count
	user.Total_size = total_size
	dao.db.Save(&user)
	response.Total_size = total_size
	response.Max_size = user.Max_size

	return &response, nil
}

func (dao *mediaDaoImpl) List(userID int64) ([]models.Media, error) {
	medias := make([]models.Media, 0)
	if err := dao.db.Where("user_id = ?", userID).Find(&medias).Error; err != nil {
		return nil, err
	}

	return medias, nil
}

func (dao *mediaDaoImpl) Delete(ID int64, userID int64) (models.Media, error) {
	media := models.Media{}
	user := models.User{}
	// Cap nhat thong tin user
	if err := dao.db.Where("id = ? ", userID).Find(&user).Error; err != nil {
		log.Println(err)
		return media, err
	}
	if err := dao.db.Where("id = ? ", ID).Find(&media).Error; err != nil {
		log.Println(err)
		return media, err
	}
	user.Total_size = user.Total_size - media.Video_size
	dao.db.Save(&user)
	// Xoa file tren storage
	file_name := media.Video_name
	log.Println(userID)
	//
	// var path = "./storage/" + string(userID) + "/" + file_name
	var path = "./storage/" + fmt.Sprint(userID) + "/" + file_name

	log.Println(path)
	err := os.Remove(path)
	if err != nil {
		log.Println(err)
		return media, err
	}
	// Xoa an ghi trong DB
	if err := dao.db.Where("id = ?", ID).Delete(&media).Error; err != nil {
		return media, err
	}
	return media, nil
}
