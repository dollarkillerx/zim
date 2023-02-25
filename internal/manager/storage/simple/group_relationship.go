package simple

import "github.com/dollarkillerx/zim/internal/manager/models"

func (s *SimpleStorage) GroupUserRelevance(projectID string, groupID string, userID string) error {
	var ex int64
	err := s.orm.Model(&models.GroupRelationship{}).
		Where("project_id = ?", projectID).
		Where("group_id = ?", groupID).
		Where("user_id = ?", userID).
		Count(&ex).Error
	if err != nil {
		return err
	}

	if ex != 0 {
		return nil
	}

	err = s.orm.Model(&models.GroupRelationship{}).Create(&models.GroupRelationship{
		ProjectID: projectID,
		UserID:    userID,
		GroupID:   groupID,
	}).Error

	return err
}

func (s *SimpleStorage) GroupUserUnRelevance(projectID string, groupID string, userID string) error {
	err := s.orm.Model(&models.GroupRelationship{}).
		Where("project_id = ?", projectID).
		Where("group_id = ?", groupID).
		Where("user_id = ?", userID).Delete(&models.GroupRelationship{}).Error

	return err
}

func (s *SimpleStorage) GroupDissolve(projectID string, groupID string) error {
	err := s.orm.Model(&models.GroupRelationship{}).
		Where("project_id = ?", projectID).
		Where("group_id = ?", groupID).
		Delete(&models.GroupRelationship{}).Error

	return err
}

func (s *SimpleStorage) GroupUserList(projectID string, groupID string) (total int, userIDs []string, err error) {
	var grs []models.GroupRelationship

	err = s.orm.Model(&models.GroupRelationship{}).
		Where("project_id = ?", projectID).
		Where("group_id = ?", groupID).
		Find(&grs).Error

	for _, v := range grs {
		userIDs = append(userIDs, v.UserID)
	}

	return len(grs), userIDs, err
}
