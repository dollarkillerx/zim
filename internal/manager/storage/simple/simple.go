package simple

import (
	"github.com/dollarkillerx/common/pkg/client"
	"github.com/dollarkillerx/common/pkg/conf"
	"github.com/dollarkillerx/common/pkg/logger"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"

	"log"
	"os"
	"time"
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
		//Logger: lib.GormLoggerNew(logger.GetLogger()),
		Logger: gormLogger.New(
			//将标准输出作为Writer
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			gormLogger.Config{
				//设定慢查询时间阈值为1ms
				SlowThreshold: 3 * time.Second,
				//设置日志级别，只有Warn和Info级别会输出慢查询日志
				LogLevel: gormLogger.Info,
			},
		),
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

	simpleStorage.autoMigrate()

	return simpleStorage, nil
}
