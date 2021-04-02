package dtos

import (
	"time"

	"github.com/macduyhai/SmartHomeVer2/models"
)

type AddMediaRequest struct {
	User_ID int64        `json:"user_id"`
	Files   []FileUpload `json:"files"`
	Key     string       `json:"key"`
}

type AddMediaResponse struct {
	User_ID    int64    `json:"user_id"`
	Total_size int64    `json:"total_size"`
	Max_size   int64    `json:"max_size"`
	Video_name []string `json:"video_name"`
}
type ListMediaRequest struct {
	User_ID int64  `json:"user_id"`
	Key     string `json:"key"`
}
type ListMediaResponse struct {
	User_ID int64 `json:"user_id"`
	Medias  []models.Media
}
type DeleteMediaRequest struct {
	ID      int64  `json:"id"`
	User_ID int64  `json:"user_id"`
	Key     string `json:"key"`
}
type Media struct {
	Video_name string     `json:"video_name"`
	Video_size int64      `json:"video_size"`
	Video_time int64      `json:"video_time"`
	CreateAt   *time.Time `json:"createat"`
}
type MediaResponse struct { // DeleteResponse , TurnOnResponse, TurnOffResponse
	Status string `json:"status"`
}
