package dtos

import "time"

// ID        int64      `gorm:"column:id;PRIMARY_KEY"`
// Mac       string     `gorm:"column:mac"`
// Serial    string     `gorm:"column:serial"`
// idDevice  string     `gorm:"column:id_device"`
// Name      string     `gorm:"column:name"`
// Type      string     `gorm:"column:type"`
// LastState bool       `gorm:"column:laststate"`
// NewState  bool       `gorm:"column:newstate"`
// CreateAt  *time.Time `gorm:"column:created_at"`
type DeviceResponse struct { // DeleteResponse , TurnOnResponse, TurnOffResponse
	Status string `json:"status"`
}
type AddResponse struct {
	Mac      string `json:"mac"`
	idDevice int    `json:"id_device"`
	Name     string `json:"name"`
	Type     string `json:"type"`
}

type EditResponse struct {
	Mac       string `json:"mac"`
	idDevice  int    `json:"id_device"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	LastState bool   `json:"laststate"`
}

type AddRequest struct {
	Mac      string     `json:"mac"`
	Name     string     `json:"name"`
	Type     string     `json:"type"`
	Serial   string     `json:"serial"`
	CreateAt *time.Time `json:"createat"`
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
