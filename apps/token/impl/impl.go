package impl

import (
	"context"
	"github.com/zginkgo/ginkgo_keyauth/apps"
	"github.com/zginkgo/ginkgo_keyauth/apps/token"
	"github.com/zginkgo/ginkgo_keyauth/apps/user"
	"github.com/zginkgo/ginkgo_keyauth/conf"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"google.golang.org/grpc"
)

var (
	svr = &impl{}
)

type impl struct {
	col *mongo.Collection
	log logger.Logger
	token.UnimplementedServiceServer
	user user.ServiceServer
}

func (s *impl) Config() error {
	// 依赖MongoDB的DB对象
	db, err := conf.C().Mongo.GetDB()
	if err != nil {
		return err
	}

	// 获取一个Collection对象, 通过Collection对象来进行CRUD
	s.col = db.Collection(s.Name())
	s.log = zap.L().Named(s.Name())
	s.user = apps.GetGrpcApp(user.AppName).(user.ServiceServer)

	// 创建索引
	indexs := []mongo.IndexModel{
		{
			Keys: bsonx.Doc{
				{Key: "refresh_token", Value: bsonx.Int32(-1)},
			},
			Options: options.Index().SetUnique(true),
		},
	}

	_, err = s.col.Indexes().CreateMany(context.Background(), indexs)
	if err != nil {
		return err
	}
	return nil
}

func (s *impl) Name() string {
	return token.AppName
}

func (s *impl) Registry(server *grpc.Server) {
	token.RegisterServiceServer(server, svr)
}

func init() {
	apps.RegistryGrpcApp(svr)
}
