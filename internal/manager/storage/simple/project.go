package simple

import (
	"github.com/dollarkillerx/zim/internal/manager/models"
	"github.com/rs/xid"
)

func (s *SimpleStorage) ProjectCreate(superAdminID string, projectName string) (*models.Project, error) {
	project := models.Project{
		ProjectID: xid.New().String(),
		SupID:     superAdminID,
		Name:      projectName,
		Token:     xid.New().String(),
	}

	err := s.orm.Model(&models.Project{}).Create(&project).Error

	return &project, err
}

func (s *SimpleStorage) ProjectDelete(projectID string) error {
	// 更多的东西
	err := s.orm.Model(&models.Project{}).
		Where("project_id = ?", projectID).Error

	return err
}

func (s *SimpleStorage) ProjectReset(supID string, projectID string, projectName string) (*models.Project, error) {
	err := s.orm.Model(&models.Project{}).
		Where("sup_id = ?", supID).
		Where("project_id = ?", projectID).Update("name", projectName).Error

	return &models.Project{
		ProjectID: projectID,
		SupID:     supID,
		Name:      projectName,
	}, err
}

func (s *SimpleStorage) ProjectList(supID string) (projects []models.Project, err error) {
	err = s.DB().Model(&models.Project{}).
		Where("sup_id = ?", supID).Find(&projects).Error
	return
}
