package server

import (
	"github.com/dollarkillerx/zim/api/manager"
	"github.com/dollarkillerx/zim/internal/manager/conf"
	"github.com/dollarkillerx/zim/internal/manager/storage"
	"github.com/dollarkillerx/zim/internal/manager/storage/simple"
)

type ManagerServer struct {
	storage storage.Interface
	manager.UnimplementedManagerServer
}

func NewManagerServer() (*ManagerServer, error) {
	simpleStorage, err := simple.NewSimpleStorage(*conf.GetConfig().PostgresConfiguration, *conf.GetConfig().RedisConfiguration)
	if err != nil {
		return nil, err
	}

	return &ManagerServer{
		storage: simpleStorage,
	}, nil
}

func (m *ManagerServer) Run() error {
	grpcNewServer()
	return nil
}
