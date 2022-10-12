package rpc

import (
	"fmt"
	"github.com/zginkgo/ginkgo_keyauth/apps/token"
	"github.com/infraboard/mcenter/client/rpc"
	"github.com/infraboard/mcenter/client/rpc/auth"
	"github.com/infraboard/mcenter/client/rpc/resolver"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	client *ClientSet
)

func SetGlobal(cli *ClientSet) {
	client = cli
}

// C Global
func C() *ClientSet {
	return client
}

// NewClient todo
// 传递注册中心地址
func NewClient(conf *rpc.Config) (*ClientSet, error) {
	zap.DevelopmentSetup()
	log := zap.L()

	//conn, err := grpc.Dial(
	//	conf.Address(),
	//	grpc.WithTransportCredentials(insecure.NewCredentials()),
	//	grpc.WithPerRPCCredentials(&Authentication{}),
	//)

	// resolver 进行解析的时候 需要mcenter客户端实例已经初始化
	conn, err := grpc.Dial(
		fmt.Sprintf("%s://%s", resolver.Scheme, "keyauth"), // Dial to "mcenter://keyauth"
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithPerRPCCredentials(auth.NewAuthentication(conf.ClientID, conf.ClientSecret)),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, err
	}

	return &ClientSet{
		conn: conn,
		log:  log,
	}, nil
}

// ClientSet 客户端
type ClientSet struct {
	conn *grpc.ClientConn
	log  logger.Logger
}

// Token 服务的SDK
func (c *ClientSet) Token() token.ServiceClient {
	return token.NewServiceClient(c.conn)
}
