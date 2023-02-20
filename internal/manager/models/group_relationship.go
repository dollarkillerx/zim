package models

import "github.com/dollarkillerx/zim/pkg/models"

// GroupRelationship 群组角色关系表
type GroupRelationship struct {
	models.BaseModel
	ProjectID string `gorm:"type:char(36);index,comment:project id"`
	GroupID   string `gorm:"type:char(36);index,comment:group id"`
	UserID    string `gorm:"type:char(36);index,comment:user id"`
}
