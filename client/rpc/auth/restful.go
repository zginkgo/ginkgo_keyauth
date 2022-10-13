package auth

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/response"
	"github.com/zginkgo/ginkgo_keyauth/apps/policy"
	"github.com/zginkgo/ginkgo_keyauth/apps/token"
	"github.com/zginkgo/ginkgo_keyauth/common/utils"
)

// RestfulAuthHandlerFunc Go-Restful auth Middleware
// 通过 r.Fillter(<middleware>) 添加中间件
func (a *KeyauthAuther) RestfulAuthHandlerFunc(
	req *restful.Request,
	resp *restful.Response,
	chain *restful.FilterChain,
) {
	// 1. 能不能获取路由装饰信息
	meta := req.SelectedRoute().Metadata()
	a.log.Debug(meta)

	// 获取meta信息, get , 判断是否开启认证
	var isAuthEnable bool
	if authV, ok := meta[label.Auth]; ok {
		switch v := authV.(type) {
		case bool:
			isAuthEnable = v
		case string:
			isAuthEnable = v == "true"
		}
	}

	var username string
	if isAuthEnable {
		// 处理Request对象
		// *restful.Request

		// 1. 认证中间件, 需要获取Token
		tkStr := utils.GetToken(req.Request)

		// 2. 到用户中心验证token合法性, 依赖用户中心的Grpc Clinet
		validateReq := token.NewValidateTokenRequest(tkStr)
		tk, err := a.auth.ValidateToken(req.Request.Context(), validateReq)
		if err != nil {
			response.Failed(resp.ResponseWriter, exception.NewUnauthorized(err.Error()))
			return
		}

		// 处理完成后 需要直接中断返回
		// *restful.Response
		// resp.Header()
		// resp.Write()

		// 如果处理ok需要把一些中间结果 然后后面的请求也能方式, 需要把结果保存中 上下文中
		// 原生 c.Request.Context()
		// 如何:  context Get
		// 认证完成后, 我们需要用户名称或者其他信息 传递下去

		// Set Context
		req.SetAttribute("token", tk)
	}

	// 获取meta信息, get, 判断是否开启鉴权
	var isPermEnable bool
	if authV, ok := meta[label.Permission]; ok {
		switch v := authV.(type) {
		case bool:
			isPermEnable = v
		case string:
			isPermEnable = v == "true"
		}
	}

	// 认证后,才能鉴权
	if isAuthEnable && isPermEnable {
		permReq := policy.NewValidatePermissionRequest()
		token2 := req.Attribute("token")
		tokenset, _ := token2.(*token.Token)
		permReq.Username = tokenset.Data.UserName

		permReq.Service = a.serviceName
		if meta != nil {
			if v, ok := meta[label.Resource]; ok {
				permReq.Resource, _ = v.(string)
			}
			if v, ok := meta[label.Action]; ok {
				permReq.Action, _ = v.(string)
			}
		}

		_, err := a.perm.ValidatePermission(req.Request.Context(), permReq)
		if err != nil {
			response.Failed(resp.ResponseWriter, exception.NewPermissionDeny(err.Error()))
			return
		}
	}

	// chain 用于将请求flow下去
	chain.ProcessFilter(req, resp)
}
