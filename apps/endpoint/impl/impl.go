package impl

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"

	"github.com/zginkgo/ginkgo_keyauth/apps"
	"github.com/zginkgo/ginkgo_keyauth/apps/endpoint"
	"github.com/zginkgo/ginkgo_keyauth/conf"
)

var (
	// Service 服务实例
	svr = &service{}
)

type service struct {
	col *mongo.Collection
	log logger.Logger
	endpoint.UnimplementedRPCServer
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
	return nil
}

func (s *service) Name() string {
	return endpoint.AppName
}

func (s *service) Registry(server *grpc.Server) {
	endpoint.RegisterRPCServer(server, svr)
}

func init() {
	apps.RegistryInternalApp(svr)
	apps.RegistryGrpcApp(svr)
}
