package models

import "time"

//Device : Cấu trúc device trogn data base
type Device struct {
	ID              int64      `gorm:"column:id;PRIMARY_KEY"`
	User_ID         int64      `gorm:"column:user_id"`
	Controller_id   int64     `gorm:"column:controller_id"`
	Device_id   	string     `gorm:"column:device_id"`
	Name            string     `gorm:"column:name"`
	Type            string     `gorm:"column:type"`
	State       	bool       `gorm:"column:state"`
	CreateAt        *time.Time `gorm:"column:created_at"`
}

