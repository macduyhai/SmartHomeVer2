package models

import "time"

//Device : Cấu trúc device trogn data base
type Device struct {
	ID        int64      `gorm:"column:id;PRIMARY_KEY"`
	Mac       string     `gorm:"column:mac"`
	Serial    string     `gorm:"column:serial"`
	idDevice  string     `gorm:"column:id_device"`
	Name      string     `gorm:"column:name"`
	Type      string     `gorm:"column:type"`
	LastState bool       `gorm:"column:laststate"`
	NewState  bool       `gorm:"column:newstate"`
	CreateAt  *time.Time `gorm:"column:created_at"`
}
