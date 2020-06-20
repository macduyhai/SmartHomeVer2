package dtos

import (
	"time"

	"github.com/macduyhai/SmartHomeVer2/models"
)

//------------------------------- Request -------------------------------------
type AddRequest struct {
	User_ID         int64  `json:"user_id"`
	Chip_ID         string `json:"chip_id"`
	Flash_Chip_ID   string `json:"flash_chip_id"`
	IDE_Flash_Size  string `json:"ide_flash_size"`
	Real_Flash_Size string `json:"real_flash_size"`
	Soft_AP_IP      string `json:"soft_ap_ip"`
	Soft_AP_MAC     string `json:"soft_ap_mac"`
	Station_MAC     string `json:"station_mac"`
	Serial          string `json:"serial"`
	Name            string `json:"name"`
	Type            string `json:"type"`
}
type AddResponse struct {
	User_ID     int64      `json:"user_id"`
	Station_MAC string     `json:"station_mac"`
	Chip_ID     string     `json:"chip_id"`
	Name        string     `json:"name"`
	Type        string     `json:"type"`
	State       bool       `json:"state"`
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
	User_ID  int64  `json:"user_id"`
	Username string `json:"username"`
	Chip_ID  string `json:"chip_id"`
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
	Chip_ID         string     `json:"chip_id"`
	Flash_Chip_ID   string     `json:"flash_chip_id"`
	IDE_Flash_Size  string     `json:"ide_flash_size"`
	Real_Flash_Size string     `json:"real_flash_size"`
	Soft_AP_IP      string     `json:"soft_ap_ip"`
	Soft_AP_MAC     string     `json:"soft_ap_mac"`
	Station_MAC     string     `json:"station_mac"`
	Serial          string     `json:"serial"`
	Name            string     `json:"name"`
	Type            string     `json:"type"`
	CreateAt        *time.Time `json:"createat"`
}
