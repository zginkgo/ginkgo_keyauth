package rpc

import (
	"context"
	"fmt"
	mcenter "github.com/infraboard/mcenter/client/rpc"
	"github.com/stretchr/testify/assert"
	"github.com/zginkgo/ginkgo_keyauth/apps/token"
	"testing"
)

// keyauth 客户端
// 需要配置注册中心地址
// 获取注册中心的客户端, 使用注册中心的客户端 查询 keyauth的地址
func TestBookQuery(t *testing.T) {
	should := assert.New(t)
	conf := mcenter.NewDefaultConfig()
	conf.Address = "127.0.0.1:18010"
	conf.ClientID = "t53ZcSeBXvQfItjDhTb1HTgx"
	conf.ClientSecret = "DkU6JwaJeOFhPusctwCe7yzAl6b3SCFM"

	// 传递Mcenter配置, 客户端通过Mcenter进行搜索, New一个用户中心的客户端
	keyauthClient, err := NewClient(conf)

	// 使用SDK 调用 keyauth 进行凭证的校验
	// c.Token().ValidateToken()

	if should.NoError(err) {
		resp, err := keyauthClient.Token().ValidateToken(
			context.Background(),
			token.NewValidateTokenRequest("wgFdYXhCTTHjoS30myNJzApf"),
		)
		should.NoError(err)

		fmt.Println(resp)
	}
}

func init() {
	// 提前加载好 mcenter客户端, resolver需要使用
	err := mcenter.LoadClientFromEnv()
	if err != nil {
		panic(err)
	}
}
