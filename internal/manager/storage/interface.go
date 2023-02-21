package storage

import (
	"github.com/dollarkillerx/zim/internal/manager/models"
	"gorm.io/gorm"
)

type Interface interface {
	DB() *gorm.DB

	// super admin
	superAdmin

	// project
	project

	// user
	user

	// group
	group
}

type superAdmin interface {
	SuperAdminCreate() (*models.SuperAdmin, error)
	SuperAdminDel(superAdminID string) error
	SuperAdminReset(superAdminID string) error
}

type project interface {
	ProjectCreate(superAdminID string, projectName string) (*models.Project, error)
	ProjectDelete(projectID string) error
	ProjectReset(p)
}

type user interface {
}

type group interface {
}
