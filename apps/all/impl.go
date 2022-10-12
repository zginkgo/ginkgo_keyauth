package all

import (
	// 注册所有GRPC服务模块, 暴露给框架GRPC服务加载
	_ "github.com/zginkgo/ginkgo_keyauth/apps/book/impl"
	_ "github.com/zginkgo/ginkgo_keyauth/apps/token/impl"
	_ "github.com/zginkgo/ginkgo_keyauth/apps/user/impl"
)
