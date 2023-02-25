package models

import "github.com/dollarkillerx/zim/pkg/base_models"

// Project 项目表
type Project struct {
	base_models.BaseModel
	ProjectID string `gorm:"type:char(36);uniqueIndex;comment:project id"`
	SupID     string `gorm:"type:char(36);index;comment:sup admin id"`
	Name      string `gorm:"type:varchar(300);comment:project name"`
	Token     string `gorm:"type:varchar(300);index;comment:project token"`
}
