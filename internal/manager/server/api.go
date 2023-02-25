package server

import (
	"context"
	"log"
	"net"

	"github.com/dollarkillerx/common/pkg/client"
	"github.com/dollarkillerx/common/pkg/logger"
	"github.com/dollarkillerx/grpc_discover"
	"github.com/dollarkillerx/zim/api/manager"
	"github.com/dollarkillerx/zim/api/protocol"
	"github.com/dollarkillerx/zim/internal/manager/conf"
	"github.com/dollarkillerx/zim/pkg/enums"
	"google.golang.org/grpc"
)

func (m *ManagerServer) SuperAdminCreate(ctx context.Context, empty *protocol.Empty) (*manager.SuperAdmin, error) {
	sup, err := m.storage.SuperAdminCreate()

	if err != nil {
		logger.Error(err)
	}

	return &manager.SuperAdmin{
		SupId:    sup.SupID,
		SupToken: sup.Token,
	}, err
}

func (m *ManagerServer) SuperAdminDel(ctx context.Context, id *manager.SuperAdminId) (*protocol.Empty, error) {
	err := m.storage.SuperAdminDel(id.SupId)

	if err != nil {
		logger.Error(err)
	}

	return &protocol.Empty{}, err
}

func (m *ManagerServer) SuperAdminReset(ctx context.Context, id *manager.SuperAdminId) (*manager.SuperAdmin, error) {
	reset, err := m.storage.SuperAdminReset(id.SupId)

	if err != nil {
		logger.Error(err)
	}

	return &manager.SuperAdmin{
		SupId:    reset.SupID,
		SupToken: reset.Token,
	}, err
}

func (m *ManagerServer) ProjectCreate(ctx context.Context, request *manager.ProjectCreateRequest) (*manager.Project, error) {
	project, err := m.storage.ProjectCreate(request.SupId, request.ProjectName)

	if err != nil {
		logger.Error(err)
	}

	return &manager.Project{
		ProjectId:    project.ProjectID,
		SupId:        project.SupID,
		ProjectName:  project.Name,
		ProjectToken: project.Token,
	}, err
}

func (m *ManagerServer) ProjectDel(ctx context.Context, request *manager.ProjectDelRequest) (*protocol.Empty, error) {
	err := m.storage.ProjectDelete(request.ProjectId)

	if err != nil {
		logger.Error(err)
	}

	return &protocol.Empty{}, err
}

func (m *ManagerServer) ProjectReset(ctx context.Context, request *manager.ProjectResetRequest) (*manager.Project, error) {
	reset, err := m.storage.ProjectReset(request.SupId, request.ProjectId, request.ProjectName)

	if err != nil {
		logger.Error(err)
	}

	return &manager.Project{
		ProjectId:    reset.ProjectID,
		SupId:        reset.SupID,
		ProjectName:  reset.Name,
		ProjectToken: reset.Token,
	}, err
}

func (m *ManagerServer) ProjectList(ctx context.Context, request *manager.ProjectListRequest) (*manager.ProjectListResponse, error) {
	list, err := m.storage.ProjectList(request.SupId)

	if err != nil {
		logger.Error(err)
	}

	var response = manager.ProjectListResponse{
		Projects: []*manager.Project{},
	}

	for _, v := range list {
		response.Projects = append(response.Projects, &manager.Project{
			ProjectId:   v.ProjectID,
			SupId:       v.SupID,
			ProjectName: v.Name,
			//ProjectToken: v.Token,
		})
	}

	return &response, err
}

func (m *ManagerServer) UserCreate(ctx context.Context, request *manager.UserCreateRequest) (*manager.User, error) {
	us, err := m.storage.UserCreate(request.ProjectId)

	if err != nil {
		logger.Error(err)
	}

	return &manager.User{
		ProjectId: us.ProjectID,
		UserId:    us.UserID,
	}, err
}

func (m *ManagerServer) UserDel(ctx context.Context, user *manager.User) (*protocol.Empty, error) {
	err := m.storage.UserDel(user.ProjectId, user.UserId)

	if err != nil {
		logger.Error(err)
	}

	return &protocol.Empty{}, err
}

func (m *ManagerServer) UserRelevance(ctx context.Context, request *manager.UserRelevanceRequest) (*protocol.Empty, error) {
	err := m.storage.UserRelevance(request.ProjectId, request.UserId1, request.UserId2)

	if err != nil {
		logger.Error(err)
	}

	return &protocol.Empty{}, err
}

func (m *ManagerServer) UserUnRelevance(ctx context.Context, request *manager.UserRelevanceRequest) (*protocol.Empty, error) {
	err := m.storage.UserUnRelevance(request.ProjectId, request.UserId1, request.UserId2)

	if err != nil {
		logger.Error(err)
	}

	return &protocol.Empty{}, err
}

func (m *ManagerServer) UserFriendsList(ctx context.Context, request *manager.UserFriendsListRequest) (*manager.UserFriendsListResponse, error) {
	total, friends, err := m.storage.UserFriendsList(request.ProjectId, request.UserId)

	if err != nil {
		logger.Error(err)
	}

	var ufs = manager.UserFriendsListResponse{
		Friends: []*manager.UserFriend{},
		Total:   total,
	}

	for _, v := range friends {
		ufs.Friends = append(ufs.Friends, &manager.UserFriend{
			UserId: v.UserID,
		})
	}

	return &manager.UserFriendsListResponse{
		Friends: nil,
		Total:   total,
	}, err
}

func (m *ManagerServer) UserOnline(ctx context.Context, request *manager.UserOnlineRequest) (*manager.UserOnlineResponse, error) {
	online, err := m.storage.UserOnline(request.ProjectId, request.Users)

	if err != nil {
		logger.Error(err)
	}

	var result = manager.UserOnlineResponse{
		UserOnline: []*manager.UserOnline{},
	}

	for _, v := range online {
		result.UserOnline = append(result.UserOnline, &manager.UserOnline{
			UserId:         v.UserID,
			LastOnlineTime: v.LastOnlineTime,
		})
	}

	return &result, err
}

func (m *ManagerServer) GroupCreate(ctx context.Context, request *manager.GroupCreateRequest) (*manager.Group, error) {
	gc, err := m.storage.GroupCreate(request.ProjectId)

	if err != nil {
		logger.Error(err)
	}

	return &manager.Group{
		ProjectId: gc.ProjectID,
		GroupId:   gc.GroupID,
	}, err
}

func (m *ManagerServer) GroupDel(ctx context.Context, group *manager.Group) (*protocol.Empty, error) {
	err := m.storage.GroupDel(group.ProjectId, group.GroupId)

	if err != nil {
		logger.Error(err)
	}

	return &protocol.Empty{}, err
}

func (m *ManagerServer) GroupUserRelevance(ctx context.Context, request *manager.GroupUserRelevanceRequest) (*protocol.Empty, error) {
	err := m.storage.GroupUserRelevance(request.ProjectId, request.GroupId, request.UserId)

	if err != nil {
		logger.Error(err)
	}

	return &protocol.Empty{}, err
}

func (m *ManagerServer) GroupUserUnRelevance(ctx context.Context, request *manager.GroupUserUnRelevanceRequest) (*protocol.Empty, error) {
	err := m.storage.GroupUserUnRelevance(request.ProjectId, request.GroupId, request.UserId)

	if err != nil {
		logger.Error(err)
	}

	return &protocol.Empty{}, err
}

func (m *ManagerServer) GroupDissolve(ctx context.Context, request *manager.GroupDissolveRequest) (*protocol.Empty, error) {
	err := m.storage.GroupDissolve(request.ProjectId, request.GroupId)

	if err != nil {
		logger.Error(err)
	}

	return &protocol.Empty{}, err
}

func (m *ManagerServer) GroupUserList(ctx context.Context, request *manager.GroupUserListRequest) (*manager.GroupUserListResponse, error) {
	total, ds, err := m.storage.GroupUserList(request.ProjectId, request.GroupId)

	if err != nil {
		logger.Error(err)
	}

	var result = manager.GroupUserListResponse{
		UserIds: []string{},
		Total:   total,
	}
	for _, v := range ds {
		result.UserIds = append(result.UserIds, v)
	}

	return &result, err
}

func grpcNewServer() *grpc.Server {
	//keepParams := grpc.KeepaliveParams(keepalive.ServerParameters{
	//	MaxConnectionIdle:     time.Duration(c.IdleTimeout),
	//	MaxConnectionAgeGrace: time.Duration(c.ForceCloseWait),
	//	Time:                  time.Duration(c.KeepAliveInterval),
	//	Timeout:               time.Duration(c.KeepAliveTimeout),
	//	MaxConnectionAge:      time.Duration(c.MaxLifeTime),
	//})

	lis, err := net.Listen("tcp", conf.GetConfig().ListenAddress)
	if err != nil {
		logger.Errorf("failed to listen: %v", err)
	}

	plugin, err := grpc_discover.NewETCDPlugin(client.ETCDOption(*conf.GetConfig().ETCDConfiguration))
	if err != nil {
		logger.Panic(err)
	}

	// 注册服务 registration service
	serverID, err := plugin.Register(enums.DiscoverManager.String(), lis.Addr().String())
	if err != nil {
		logger.Panic(err)
	}

	plugin.AutoUnRegister(serverID)

	grpcServer := grpc.NewServer()

	managerServer, err := NewManagerServer()
	if err != nil {
		logger.Panic(err)
	}

	manager.RegisterManagerServer(grpcServer, managerServer)

	go func() {
		logger.Infof("server listening at %v", lis.Addr())
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	return grpcServer
}
