package models

import "time"

type Controller struct {
	ID              int64      `gorm:"column:id;PRIMARY_KEY"`
	User_ID         int64      `gorm:"column:user_id"`
	Chip_ID         string     `gorm:"column:chip_id"`
	Flash_Chip_ID   string     `gorm:"column:flash_chip_id"`
	IDE_Flash_Size  string     `gorm:"column:ide_flash_size"`
	Real_Flash_Size string     `gorm:"column:real_flash_size"`
	Soft_AP_IP      string     `gorm:"column:soft_ap_ip"`
	Soft_AP_MAC     string     `gorm:"column:soft_ap_mac"`
	Station_MAC     string     `gorm:"column:station_mac"`
	Serial          string     `gorm:"column:serial"`
	Name            string     `gorm:"column:name"`
	Type            string     `gorm:"column:type"`
	Active       	bool       `gorm:"column:active"`
	Number_device   int64      `gorm:"column:number_device"`
	CreateAt        *time.Time `gorm:"column:created_at"`
}