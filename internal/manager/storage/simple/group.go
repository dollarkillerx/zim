package simple

import (
	"github.com/dollarkillerx/zim/internal/manager/models"
	"github.com/rs/xid"
)

func (s *SimpleStorage) GroupCreate(projectID string) (*models.Group, error) {
	group := models.Group{
		GroupID:   xid.New().String(),
		ProjectID: projectID,
	}

	err := s.orm.Model(&models.Group{}).Create(&group).Error

	return &group, err
}

func (s *SimpleStorage) GroupDel(projectID string, groupID string) error {
	err := s.orm.Model(&models.Group{}).
		Where("project_id = ?", projectID).
		Where("group_id = ?", groupID).Error

	return err
}
