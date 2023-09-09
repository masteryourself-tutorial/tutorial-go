package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name          string
	Password      string
	Phone         string
	Email         string
	Identity      string
	ClientIp      string
	LoginTime     time.Time
	HeartbeatTime time.Time
	LogoutTime    time.Time
	IsLogout      bool
	DeviceInfo    string
}

func (u User) TableName() string {
	return "chat_user"
}
