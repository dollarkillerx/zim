package simple

import "github.com/dollarkillerx/zim/internal/manager/models"

func (s *SimpleStorage) autoMigrate() {
	s.orm.AutoMigrate(
		&models.SuperAdmin{},
		&models.Project{},
		&models.User{},
		&models.UserRelationship{},
		&models.Group{},
		&models.GroupRelationship{},
	)
}
