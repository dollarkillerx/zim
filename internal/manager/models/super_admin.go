package models

import "github.com/dollarkillerx/zim/pkg/models"

// SuperAdmin 最高角色表
type SuperAdmin struct {
	models.BaseModel
	SupID string `gorm:"type:char(36);index,comment:sup admin id"`
	Token string `gorm:"type:varchar(300);index,comment:sup admin token"`
}
