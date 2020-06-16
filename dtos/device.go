package dtos

import "time"

//------------------------------- Request -------------------------------------
type AddRequest struct {
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

//------------------------------- Response -------------------------------------
type DeviceResponse struct { // DeleteResponse , TurnOnResponse, TurnOffResponse
	Status string `json:"status"`
}
type AddResponse struct {
	Station_MAC string     `json:"station_mac"`
	Chip_ID     int        `json:"chip_id"`
	Name        string     `json:"name"`
	Type        string     `json:"type"`
	NewState    bool       `json:"newstate"`
	CreateAt    *time.Time `json:"createat"`
}

type EditResponse struct {
	Mac       string `json:"mac"`
	idDevice  int    `json:"id_device"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	LastState bool   `json:"laststate"`
}

type DeleteRequest struct {
	Mac      string `json:"mac"`
	Name     string `json:"name"`
	idDevice int    `json:"id_device"`
	Serial   string `json:"serial"`
}
type EditRequest struct {
	Mac       string `json:"mac"`
	Name      int    `json:"name"`
	Type      string `json:"type"`
	Serial    string `json:"serial"`
	LastState bool   `json:"laststate"`
}
type TurnOnRequest struct {
	Mac      string `json:"mac"`
	idDevice int    `json:"id_device"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	NewState bool   `json:"newstate"`
}
type TurnOffRequest struct {
	Mac      string `json:"mac"`
	idDevice int    `json:"id_device"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	NewState bool   `json:"newstate"`
}

type Device struct {
	Chip_ID         string `json:"chip_id"`
	Flash_Chip_ID   string `json:"flash_chip_id"`
	IDE_Flash_Size  string `json:"ide_flash_size"`
	Real_Flash_Size string `json:"real_flash_size"`
	Soft_AP_IP      string `json:"soft_ap_ip"`
	Soft_AP_MAC     string `json:"soft_ap_mac"`
	Station_MAC     string `json:"Station_mac"`
	Value           bool   `json:"value"`
	Serial          string `json:"serial"`
}