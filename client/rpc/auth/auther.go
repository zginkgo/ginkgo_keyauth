package auth

import (
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/zginkgo/ginkgo_keyauth/apps/audit"
	"github.com/zginkgo/ginkgo_keyauth/apps/policy"
	"github.com/zginkgo/ginkgo_keyauth/apps/token"
	"github.com/zginkgo/ginkgo_keyauth/client/rpc"
)

func NewKeyauthAuther(client *rpc.ClientSet, serviceName string) *KeyauthAuther {
	return &KeyauthAuther{
		auth:        client.Token(),
		perm:        client.Policy(),
		audit:       client.Audit(),
		log:         zap.L().Named("http.auther"),
		serviceName: serviceName,
	}
}

// KeyauthAuther 用keyauth提供的中间件
type KeyauthAuther struct {
	log         logger.Logger
	auth        token.ServiceClient
	audit       audit.RPCClient
	perm        policy.RPCClient
	serviceName string
}
