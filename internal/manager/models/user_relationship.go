package models

import "github.com/dollarkillerx/zim/pkg/models"

// UserRelationship 用户关系表
type UserRelationship struct {
	models.BaseModel
	ProjectID string `gorm:"type:char(36);index,comment:project id"`
	UserID    string `gorm:"type:char(36);index,comment:user id"`
	FriendID  string `gorm:"type:char(36);index,comment:friend id"`
}
