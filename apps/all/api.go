package all

import (
	_ "github.com/zginkgo/ginkgo_keyauth/apps/token/api"
	_ "github.com/zginkgo/ginkgo_keyauth/apps/user/api"
	// 注册所有HTTP服务模块,暴露给HTTP服务器加载
	_ "github.com/zginkgo/ginkgo_keyauth/apps/book/api"
)
