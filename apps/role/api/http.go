package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/response"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/zginkgo/ginkgo_keyauth/apps"
	"github.com/zginkgo/ginkgo_keyauth/apps/role"
)

var (
	h = &handler{}
)

type handler struct {
	service role.Service
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(role.AppName)
	h.service = apps.GetGrpcApp(role.AppName).(role.Service)
	return nil
}

func (h *handler) Name() string {
	return role.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{h.Name()}

	ws.Route(ws.POST("/").To(h.CreateRole).
		Doc("create a role").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(role.CreateRoleRequest{}).
		Writes(response.NewData(role.Role{})))
}

func init() {
	apps.RegistryRESTfulApp(h)
}
