package simple

import (
	"github.com/dollarkillerx/zim/internal/manager/models"
	"github.com/rs/xid"
)

func (s *SimpleStorage) SuperAdminCreate() (*models.SuperAdmin, error) {
	sup := models.SuperAdmin{
		SupID: xid.New().String(),
		Token: xid.New().String(),
	}

	err := s.orm.Model(&models.SuperAdmin{}).Create(&sup).Error

	return &sup, err
}

func (s *SimpleStorage) SuperAdminGetBySuperAdminID(superAdminID string) (*models.SuperAdmin, error) {
	var sa models.SuperAdmin

	err := s.orm.Model(&models.SuperAdmin{}).
		Where("sup_id = ?", superAdminID).
		First(&sa).Error

	return &sa, err
}

func (s *SimpleStorage) SuperAdminDel(superAdminID string) error {
	err := s.orm.Model(&models.SuperAdmin{}).
		Where("sup_id = ?", superAdminID).
		Delete(&models.SuperAdmin{}).Error

	return err
}

func (s *SimpleStorage) SuperAdminReset(superAdminID string) (*models.SuperAdmin, error) {
	newToken := xid.New().String()
	err := s.orm.Model(&models.SuperAdmin{}).
		Where("sup_id = ?", superAdminID).
		Update("token", newToken).Error
	if err != nil {
		return nil, err
	}

	return &models.SuperAdmin{
		SupID: superAdminID,
		Token: newToken,
	}, nil
}
