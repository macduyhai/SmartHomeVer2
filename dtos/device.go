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
}
type ListResponse struct {
	User_ID  int64  `json:"user_id"`
	Username string `json:"username"`
	Devices  []models.Device
}
type EditRequest struct {
	User_ID  int64  `json:"user_id"`
	Username string `json:"username"`
	Chip_ID  string `json:"chip_id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
}
type EditResponse struct {
	User_ID  int64  `json:"user_id"`
	Username string `json:"username"`
	Device   models.Device
}
type DeleteRequest struct {
	User_ID int64  `json:"user_id"`
	Mac     string `json:"mac"`
}
type ControlRequest struct {
	User_ID int64  `json:"user_id"`
	Chip_ID string `json:"chip_id"`
	State   bool   `json:"state"`
}
type ControlResponse struct {
	Chip_ID string `json:"chip_id"`
	State   bool   `json:"state"`
}
type GetstatusRequest struct {
	Station_MAC string `json:"station_mac"`
	Chip_ID     string `json:"chip_id"`
}
type GetstatusResponse struct {
	Station_MAC string `json:"station_mac"`
	Chip_ID     string `json:"chip_id"`
	State       bool   `json:"state"`
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
