package impl

import (
	"github.com/zginkgo/ginkgo_keyauth/apps"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/zginkgo/ginkgo_keyauth/apps/policy"
	"github.com/zginkgo/ginkgo_keyauth/apps/role"
	"github.com/zginkgo/ginkgo_keyauth/conf"
	"google.golang.org/grpc"
)

var (
	// Service 服务实例
	svr = &service{}
)

type service struct {
	col  *mongo.Collection
	log  logger.Logger
	role role.Service

	policy.UnimplementedRPCServer
}

func (s *service) Config() error {
	// 依赖MongoDB的DB对象
	db, err := conf.C().Mongo.GetDB()
	if err != nil {
		return err
	}

	// 获取一个Collection对象, 通过Collection对象 来进行CRUD
	s.col = db.Collection(s.Name())
	s.log = zap.L().Named(s.Name())
	s.role = apps.GetInternalApp(role.AppName).(role.Service)
	return nil
}

func (s *service) Name() string {
	return policy.AppName
}

func (s *service) Registry(server *grpc.Server) {
	policy.RegisterRPCServer(server, svr)
}

func init() {
	apps.RegistryInternalApp(svr)
	apps.RegistryGrpcApp(svr)
}
