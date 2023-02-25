package simple

import (
	"github.com/dollarkillerx/zim/internal/manager/models"

	"log"
)

func (s *SimpleStorage) autoMigrate() {
	err := s.orm.AutoMigrate(
		&models.SuperAdmin{},
		&models.Project{},
		&models.User{},
		&models.UserRelationship{},
		&models.Group{},
		&models.GroupRelationship{},
	)

	if err != nil {
		log.Println(err)
	}
}
