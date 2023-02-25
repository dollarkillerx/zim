package models

import "github.com/dollarkillerx/zim/pkg/base_models"

// SuperAdmin 最高角色表
type SuperAdmin struct {
	base_models.BaseModel
	SupID string `gorm:"type:char(36);uniqueIndex;comment:sup admin id"`
	Token string `gorm:"type:varchar(300);index;comment:sup admin token"`
}
