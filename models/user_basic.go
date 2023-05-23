package models

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdateAt  time.Time
	DeleteAt  gorm.DeletedAt `gorm:"index"`
}

type UserBasic struct {
	Model
	Name          string
	PassWord      string
	Avatar        string
	Gender        string `gorm:"column:gender;default:male;type:varchar(6) comment 'male便是男，female表示女'"`
	Phone         string `valid:"matches(^1[3-9]{1}\\d{9}$)"`
	Email         string `valid:"email"`
	Identity      string
	ClientIP      string `valid:"ipv4"`
	ClientPort    string
	Salt          string
	LogInTime     *time.Time `gorm:"column:log_in_time"`
	HeartBeatTime *time.Time `gorm:"column:heart_beat_time"`
	LogOutTime    *time.Time `gorm:"column:log_out_time"`
	IsLogOut      bool
	DeviceInfo    string
}

func (table *UserBasic) UserTableName() string {
	return "user_basic"
}
