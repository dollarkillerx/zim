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
	SuperAdminReset(superAdminID string) (*models.SuperAdmin, error)
}

type project interface {
	ProjectCreate(superAdminID string, projectName string) (*models.Project, error)
	ProjectDelete(projectID string) error
	ProjectReset(supID string, projectID string, projectName string) (*models.Project, error)
	ProjectList(supID string) ([]models.Project, error)
}

type user interface {
	UserCreate(projectID string) (*models.User, error)
	UserDel(projectID string, userID string) error
	UserRelevance(projectID string, userID1 string, userID2 string) error
	UserUnRelevance(projectID string, userID1 string, userID2 string) error
	UserFriendsList(projectID string, userID string) (total int, userFriends []models.UserFriend, err error)
	UserOnline(projectID string, users []string) ([]models.UserOnline, error)
}

type group interface {
	GroupCreate(projectID string) (*models.Group, error)
	GroupDel(projectID string, groupID string) error
	GroupUserRelevance(projectID string, groupID string, userID string) error
	GroupUserUnRelevance(projectID string, groupID string, userID string) error
	GroupDissolve(projectID string, groupID string) error
	GroupUserList(projectID string, groupID string) (total int, userIDs []string, err error)
}
