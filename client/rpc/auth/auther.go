package auth

import (
	"github.com/zginkgo/ginkgo_keyauth/apps/token"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewKeyauthAuther(auth token.ServiceClient) *KeyauthAuther {
	return &KeyauthAuther{
		auth: auth,
		log:  zap.L().Named("http.auther"),
	}
}

// KeyauthAuther 用keyauth提供的中间件
type KeyauthAuther struct {
	log  logger.Logger
	auth token.ServiceClient
}
