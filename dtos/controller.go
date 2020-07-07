package dtos

import (
	"time"

	"SmartHomeVer2/models"
)

//------------------------------- Request -------------------------------------
type Device struct {
	// Controller_id   string     `json:"controller_id"`  // Chip_ID = Controller_id
	Device_id  		string     `json:"device_id"`
	Name            string     `json:"name"`
	Type            string     `json:"type"`
	State       	bool      `json:"state"`
}
type Controller struct {
	Chip_ID         string `json:"chip_id"`
	Flash_Chip_ID   string `json:"flash_chip_id"`
	Soft_AP_IP      string `json:"soft_ap_ip"`
	Station_MAC     string `json:"station_mac"`
	Serial          string `json:"serial"`
	Name            string `json:"name"`
	Type            string `json:"type"`
	Active			bool   `json:"active"`
	Number_device	int64 `json:"number_device"`
	Devices 		[]Device
}
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
	Active			bool   `json:"active"`
	Number_device	int64 `json:"number_device"`
	Devices 		[]Device
}
type AddResponse struct {
	User_ID     int64      `json:"user_id"`
	Station_MAC string     `json:"station_mac"`
	Chip_ID     string     `json:"chip_id"`
	Name        string     `json:"name"`
	Type        string     `json:"type"`
	Active    	bool       `json:"active"`
	Number_device	int64 `json:"number_device"`
	CreateAt    *time.Time `json:"createat"`
}

type ListRequest struct {
	User_ID  int64  `json:"user_id"`
	Username string `json:"username"`
}
type ListResponse struct {
	User_ID  int64  `json:"user_id"`
	Username string `json:"username"`
	Controllers  []models.Controller
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
	User_ID  int64  `json:"user_id"`
	Chip_ID  string `json:"chip_id"`
	Device_id string `json:"device_id"`
	State bool	`json:"state"`
}
type ControlResponse struct {
	Chip_ID  string `json:"chip_id"`
	Device_id string `json:"device_id"`
	State bool	`json:"state"`
}
type GetstatusRequest struct {
	Station_MAC  string  `json:"station_mac"`
	Chip_ID  string `json:"chip_id"`

}
type GetstatusResponse struct {
	Station_MAC  string  `json:"station_mac"`
	Chip_ID  string `json:"chip_id"`
	Devices [] models.Device
}

//---------------------------------------------------
type DeviceResponse struct { // DeleteResponse , TurnOnResponse, TurnOffResponse
	Status string `json:"status"`
}



