package protocol

import (
	"context"
	"github.com/zginkgo/ginkgo_keyauth/apps"
	"github.com/zginkgo/ginkgo_keyauth/conf"
	"github.com/infraboard/mcenter/apps/instance"
	"github.com/infraboard/mcenter/client/rpc"
	"github.com/infraboard/mcenter/client/rpc/lifecycle"
	"github.com/infraboard/mcube/grpc/middleware/recovery"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"
	"net"
	"time"
)

// NewGRPCService todo
func NewGRPCService() *GRPCService {
	log := zap.L().Named("GRPC Service")

	rc := recovery.NewInterceptor(recovery.NewZapRecoveryHandler())
	grpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor(
		rc.UnaryServerInterceptor(),
	))

	// 控制Grpc启动其他服务, 比如注册中心,或者心跳
	ctx, cancel := context.WithCancel(context.Background())
	return &GRPCService{
		svr:    grpcServer,
		l:      log,
		c:      conf.C(),
		ctx:    ctx,
		cancel: cancel,
	}
}

type GRPCService struct {
	svr *grpc.Server
	l   logger.Logger
	c   *conf.Config

	ctx    context.Context
	cancel context.CancelFunc
	// 控制实例上线和下线
	lf lifecycle.Lifecycler
}

// Start 启动GRPC服务
func (s *GRPCService) Start() {
	// 装载所有GRPC服务
	apps.LoadGrpcApp(s.svr)

	// 启动HTTP服务
	lis, err := net.Listen("tcp", s.c.App.GRPC.Addr())
	if err != nil {
		s.l.Errorf("listen grpc tcp conn error, %s", err)
		return
	}

	// 1. 理论上, 我们需要等待GRPC服务启动成功后,才注册
	// 但是 GRPC server 没有提供成功时的回调

	// 假设GRPC Server 1秒后能正常启动
	time.AfterFunc(1*time.Second, s.registry)

	s.l.Infof("GRPC 服务监听地址: %s", s.c.App.GRPC.Addr())
	if err := s.svr.Serve(lis); err != nil {
		if err == grpc.ErrServerStopped {
			s.l.Info("service is stopped")
		}

		s.l.Error("start grpc service error, %s", err.Error())
		return
	}
}

func (s *GRPCService) registry() {
	// 1. 获取 mcenter sdk 实例
	// sdk 提供注册方法
	req := instance.NewRegistryRequest()
	req.Address = s.c.App.GRPC.Addr()
	lf, err := rpc.C().Registry(s.ctx, req)
	if err != nil {
		s.l.Errorf("registry to mcenter error, %s", err)
		return
	}
	s.l.Infof("registry to mcenter success")

	// 注销时需要使用
	s.lf = lf
}

// Stop 启动GRPC服务
func (s *GRPCService) Stop() error {
	// 提前 剔除注册中心的地址
	if s.lf != nil {
		if err := s.lf.UnRegistry(s.ctx); err != nil {
			s.l.Errorf("unregistry error, %s", err)
		} else {
			s.l.Info("unregistry success")
		}
	}
	s.svr.GracefulStop()
	return nil
}
