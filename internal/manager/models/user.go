package models

import "github.com/dollarkillerx/zim/pkg/base_models"

// User 用户表
type User struct {
	base_models.BaseModel
	UserID    string `gorm:"type:char(36);uniqueIndex;comment:user id"`
	ProjectID string `gorm:"type:char(36);index;comment:project id"`
}

type UserFriend struct {
	UserID string
}

type UserOnline struct {
	UserID         string
	LastOnlineTime int64
}
