package simple

import (
	"context"

	"github.com/dollarkillerx/zim/internal/manager/models"
)

func (s *SimpleStorage) UserRelevance(projectID string, userID1 string, userID2 string) error {
	var ex int64
	err := s.orm.Model(&models.UserRelationship{}).
		Where("project_id = ?", projectID).
		Where("user_id = ?", userID1).Where("friend_id = ?", userID2).
		Count(&ex).Error
	if err != nil {
		return err
	}

	if ex != 0 {
		return nil
	}

	err = s.orm.Model(&models.UserRelationship{}).
		Where("project_id = ?", projectID).
		Where("user_id = ?", userID2).Where("friend_id = ?", userID1).
		Count(&ex).Error
	if err != nil {
		return err
	}

	if ex != 0 {
		return nil
	}

	err = s.orm.Model(&models.UserRelationship{}).Create(&models.UserRelationship{
		ProjectID: projectID,
		UserID:    userID1,
		FriendID:  userID2,
	}).Error

	return err
}

func (s *SimpleStorage) UserUnRelevance(projectID string, userID1 string, userID2 string) error {
	err := s.orm.Model(&models.UserRelationship{}).
		Where("project_id = ?", projectID).
		Where("user_id = ?", userID1).
		Where("friend_id = ?", userID2).
		Delete(&models.UserRelationship{}).Error
	if err != nil {
		return err
	}

	err = s.orm.Model(&models.UserRelationship{}).
		Where("project_id = ?", projectID).
		Where("user_id = ?", userID2).
		Where("friend_id = ?", userID1).
		Delete(&models.UserRelationship{}).Error
	if err != nil {
		return err
	}

	return err
}

func (s *SimpleStorage) UserFriendsList(projectID string, userID string) (total int, userFriends []models.UserFriend, err error) {
	err = s.orm.Model(&models.UserRelationship{}).
		Where("project_id = ?", projectID).
		Where(s.orm.Where("user_id = ?", userID).Or("friend_id = ?", userID)).
		Find(&userFriends).Error

	return len(userFriends), userFriends, err
}

func (s *SimpleStorage) UserOnline(projectID string, users []string) ([]models.UserOnline, error) {
	var onlineUsers []models.UserOnline

	for _, v := range users {
		result, err := s.redisClient.Exists(context.TODO(), v).Result()
		if err != nil {
			return nil, err
		}

		if result != 0 {
			onlineUsers = append(onlineUsers, models.UserOnline{
				UserID: v,
				Online: true,
			})
		} else {
			onlineUsers = append(onlineUsers, models.UserOnline{
				UserID: v,
				Online: false,
			})
		}
	}

	return onlineUsers, nil
}
