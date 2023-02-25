package simple

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"

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
		exKey := fmt.Sprintf("user-online-%s-%s", projectID, v)

		result, err := s.redisClient.Get(context.TODO(), exKey).Int64()
		if err != nil {
			if err != redis.Nil {
				return nil, err
			}
			result = 0
		}

		onlineUsers = append(onlineUsers, models.UserOnline{
			UserID:         v,
			LastOnlineTime: result,
		})
	}

	return onlineUsers, nil
}

func (s *SimpleStorage) UserOnlinePing(projectID string, users []string) error {
	for _, v := range users {
		exKey := fmt.Sprintf("user-online-%s-%s", projectID, v)
		err := s.redisClient.Set(context.TODO(), exKey, time.Now().Unix(), 0).Err()
		if err != nil {
			return err
		}
	}

	return nil
}
