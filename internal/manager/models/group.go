package models

import "github.com/dollarkillerx/zim/pkg/base_models"

// Group 群组表
type Group struct {
	base_models.BaseModel
	GroupID   string `gorm:"type:char(36);uniqueIndex,comment:group id"`
	ProjectID string `gorm:"type:char(36);index,comment:project id"`
}
