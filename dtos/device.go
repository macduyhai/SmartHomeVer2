package dtos

import (
	"time"

	"github.com/macduyhai/SmartHomeVer2/models"
)

//------------------------------- Request -------------------------------------
type AddRequest struct {
	User_ID     int64  `json:"user_id"`
	Mac         string `json:"mac"`
	Device_name string `json:"device_name"`
	Location    string `json:"location"`
	Key         string `json:"key"`
}
type AddResponse struct {
	Mac         string     `json:"mac"`
	Device_name string     `json:"chip_id"`
	Video_name  string     `json:"video_name"`
	Video_size  int64      `json:"video_size"`
	Video_time  int64      `json:"video_time"`
	Status      int        `json:"status"`
	Location    string     `json:"location"`
	Expired     *time.Time `json:"expired"`
	CreateAt    *time.Time `json:"createat"`
}

type ListRequest struct {
	User_ID  int64  `json:"user_id"`
	Username string `json:"username"`
	Key      string `json:"key"`
}
type ListResponse struct {
	User_ID  int64  `json:"user_id"`
	Username string `json:"username"`
	Devices  []models.Device
}
type EditRequest struct {
	User_ID     int64  `json:"user_id"`
	Mac         string `json:"mac"`
	Device_Name string `json:"device_name"`
	Location    string `json:"location"`
	Key         string `json:"key"`
}
type EditResponse struct {
	User_ID  int64  `json:"user_id"`
	Username string `json:"username"`
	Device   models.Device
}
type DeleteRequest struct {
	User_ID int64  `json:"user_id"`
	Mac     string `json:"mac"`
	Key     string `json:"key"`
}
type UploadRequest struct {
	User_ID int64        `json:"user_id"`
	Files   []FileUpload `json:"files"`
	Key     string       `json:"key"`
}
type PushRequest struct {
	User_ID    int64  `json:"user_id"`
	Mac        string `json:"mac"`
	Video_name string `json:"video_name"`
}
type FileUpload struct {
	Video_name string `json:"video_name"`
	Video_size int64  `json:"video_size"`
	Video_time int64  `json:"video_time"`
}

type UploadResponse struct {
	User_ID    int64    `json:"user_id"`
	Total_size int64    `json:"total_size"`
	Max_size   int64    `json:"max_size"`
	Video_name []string `json:"video_name"`
}

// type UploadResponse struct {
// 	Device models.Device
// }
type GetstatusRequest struct {
	User_ID int64  `json:"user_id"`
	Mac     string `json:"mac"`
	Key     string `json:"key"`
}
type GetstatusResponse struct {
	Mac    string `json:"mac"`
	Status int    `json:"status"`
}

//---------------------------------------------------
type DeviceResponse struct { // DeleteResponse , TurnOnResponse, TurnOffResponse
	Status string `json:"status"`
}

type Device struct {
	Mac        string     `json:"mac"`
	Video_name string     `json:"video_name"`
	Video_size int64      `json:"video_size"`
	Video_time int64      `json:"video_time"`
	Status     int        `json:"status"`
	Location   string     `json:"location"`
	Expired    *time.Time `json:"expired"`
	CreateAt   *time.Time `json:"createat"`
}

// ID         int64      `gorm:"column:id;PRIMARY_KEY"`
// User_ID    int64      `gorm:"column:user_id"`
// Mac        string     `gorm:"column:mac"`
// Video_name string     `gorm:"column:video_name"`
// Video_size int64      `gorm:"column:video_size"`
// Video_time int64      `gorm:"column:video_time"`
// Status     int        `gorm:"column:status"`
// Expired    *time.Time `gorm:"column:expired"`
// Location   string     `gorm:"column:location"`
// State      bool       `gorm:"column:state"`
// CreateAt   *time.Time `gorm:"column:created_at"`
// Updated_at *time.Time `gorm:"column:updated_at"`
