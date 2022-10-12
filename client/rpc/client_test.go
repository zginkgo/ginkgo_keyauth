package rpc

import (
	"context"
	"fmt"
	mcenter "github.com/infraboard/mcenter/client/rpc"
	"github.com/stretchr/testify/assert"
	"github.com/zginkgo/ginkgo_keyauth/apps/policy"
	"testing"
)

// keyauth 客户端
// 需要配置注册中心地址
// 获取注册中心的客户端, 使用注册中心的客户端 查询 keyauth的地址
//func TestBookQuery(t *testing.T) {
//	should := assert.New(t)
//	conf := mcenter.NewDefaultConfig()
//	conf.Address = "127.0.0.1:18010"
//	conf.ClientID = "t53ZcSeBXvQfItjDhTb1HTgx"
//	conf.ClientSecret = "DkU6JwaJeOFhPusctwCe7yzAl6b3SCFM"
//
//	// 传递Mcenter配置, 客户端通过Mcenter进行搜索, New一个用户中心的客户端
//	keyauthClient, err := NewClient(conf)
//
//	// 使用SDK 调用 keyauth 进行凭证的校验
//	// c.Token().ValidateToken()
//
//	// 进行服务功能注册
//	//keyauthClient.Endpoint().RegistryEndpoint()
//
//	// 鉴权校验
//	//keyauthClient.Policy().ValidatePermission()
//
//	if should.NoError(err) {
//		resp, err := keyauthClient.Token().ValidateToken(
//			context.Background(),
//			token.NewValidateTokenRequest("wgFdYXhCTTHjoS30myNJzApf"),
//		)
//		should.NoError(err)
//
//		fmt.Println(resp)
//	}
//}

// 测试鉴权
func TestValidatePermission(t *testing.T) {
	should := assert.New(t)

	conf := mcenter.NewDefaultConfig()
	conf.Address = "127.0.0.1:18010"
	conf.ClientID = "t53ZcSeBXvQfItjDhTb1HTgx"
	conf.ClientSecret = "DkU6JwaJeOFhPusctwCe7yzAl6b3SCFM"

	// 传递Mcenter配置, 客户端通过Mcenter进行搜索, New一个用户中心的客户端

	keyauthClient, err := NewClient(conf)
	if should.NoError(err) {
		req := policy.NewValidatePermissionRequest()
		req.Username = "member3"
		req.Service = "cmdb"
		req.Resource = "secret"
		req.Namespace = "default"
		req.Action = "delete12345"
		//names := []string{"cmdb", "keyauth", "member"}
		//p, err := keyauthClient.Role().QueryRole(context.TODO(), role.NewQueryRoleRequestWithName(names))
		//p, err := keyauthClient.Policy().QueryPolicy(context.TODO(), policy.NewQueryPolicyRequest())
		p, err := keyauthClient.Policy().ValidatePermission(context.TODO(), req)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(p, err)
		t.Log(p)
	}
}

func init() {
	// 提前加载好 mcenter客户端, resolver需要使用
	err := mcenter.LoadClientFromEnv()
	if err != nil {
		panic(err)
	}
}
