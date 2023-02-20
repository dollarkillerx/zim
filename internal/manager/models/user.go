package models

import "github.com/dollarkillerx/zim/pkg/models"

// User 用户表
type User struct {
	models.BaseModel
	UserID    string `gorm:"type:char(36);uniqueIndex,comment:user id"`
	ProjectID string `gorm:"type:char(36);index,comment:project id"`
}
