package models

import "time"

type Media struct {
	ID         int64      `gorm:"column:id;PRIMARY_KEY"`
	User_ID    int64      `gorm:"column:user_id"`
	Video_name string     `gorm:"column:video_name"`
	Video_size int64      `gorm:"column:video_size"`
	Video_time int64      `gorm:"column:video_time"`
	CreateAt   *time.Time `gorm:"column:created_at"`
}
