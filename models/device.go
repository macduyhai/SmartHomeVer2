package models

import "time"

//Device : Cấu trúc device trogn data base
type Device struct {
	ID          int64      `gorm:"column:id;PRIMARY_KEY"`
	User_ID     int64      `gorm:"column:user_id"`
	Device_name string     `gorm:"column:device_name"`
	Mac         string     `gorm:"column:mac"`
	Video_name  string     `gorm:"column:video_name"`
	Video_size  int64      `gorm:"column:video_size"`
	Video_time  int64      `gorm:"column:video_time"`
	Status      int        `gorm:"column:status"`
	Location    string     `gorm:"column:location"`
	Map_long    string     `gorm:"column:map_long"`
	Map_lat     string     `gorm:"column:map_lat"`
	Expired     *time.Time `gorm:"column:expired"`
	CreateAt    *time.Time `gorm:"column:created_at"`
	Updated_at  *time.Time `gorm:"column:updated_at"`
}
