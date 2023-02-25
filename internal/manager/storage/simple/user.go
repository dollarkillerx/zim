package simple

import (
	"github.com/dollarkillerx/zim/internal/manager/models"
	"github.com/rs/xid"
)

func (s *SimpleStorage) UserCreate(projectID string) (*models.User, error) {
	user := models.User{
		UserID:    xid.New().String(),
		ProjectID: projectID,
	}

	err := s.orm.Model(&models.User{}).Create(&user).Error

	return &user, err
}

func (s *SimpleStorage) UserDel(projectID string, userID string) error {
	err := s.orm.Model(&models.User{}).
		Where("project_id = ?", projectID).
		Where("user_id = ?", userID).
		Delete(&models.User{}).Error

	return err
}
