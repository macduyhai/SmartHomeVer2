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
	Key    string `json:"key"`
	Mac    string `json:"mac"`
	Id     int    `json:"id"`
	Value  int    `json:"value"`
	Serial string `json:"serial"`
}
