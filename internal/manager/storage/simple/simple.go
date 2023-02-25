package simple

import (
	"github.com/dollarkillerx/common/pkg/client"
	"github.com/dollarkillerx/common/pkg/conf"
	"github.com/dollarkillerx/common/pkg/lib"
	"github.com/dollarkillerx/common/pkg/logger"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type SimpleStorage struct {
	orm         *gorm.DB
	redisClient *redis.Client
}

func (s *SimpleStorage) DB() *gorm.DB {
	return s.orm
}

func NewSimpleStorage(pgsqlConf conf.PostgresConfiguration, redisConf conf.RedisConfiguration) (*SimpleStorage, error) {
	orm, err := client.PostgresClient(pgsqlConf, &gorm.Config{
		Logger: lib.GormLoggerNew(logger.GetLogger()),
	})
	if err != nil {
		logger.Errorf(err.Error())
		return nil, errors.WithStack(err)
	}

	redisClient, err := client.RedisClient(redisConf)
	if err != nil {
		logger.Errorf(err.Error())
		return nil, errors.WithStack(err)
	}

	simpleStorage := &SimpleStorage{
		orm:         orm,
		redisClient: redisClient,
	}

	return simpleStorage, nil
}
