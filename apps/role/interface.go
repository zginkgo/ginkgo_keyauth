package role

import (
	"context"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/rs/xid"
)

const (
	AppName = "role"
)

const (
	// DefaultPageSize 默认分页大小
	DefaultPageSize = 20
	// DefaultPageNumber 默认页号
	DefaultPageNumber = 1
)

var (
	validate = validator.New()
)

type Service interface {
	CreateRole(context.Context, *CreateRoleRequest) (*Role, error)
	RPCServer
}

func NewQueryRoleRequestWithName(names []string) *QueryRoleRequest {
	return &QueryRoleRequest{
		Page:      NewDefaultPageRequest(),
		RoleNames: names,
	}
}

// NewPageRequest 实例化
func NewPageRequest(ps uint, pn uint) *PageRequest {
	return &PageRequest{
		PageSize:   uint64(ps),
		PageNumber: uint64(pn),
	}
}

func NewDefaultPageRequest() *PageRequest {
	return NewPageRequest(DefaultPageSize, DefaultPageNumber)
}

func (s *RoleSet) HasPermission(req *PermissionRequest) (bool, *Role) {
	for i := range s.Items {
		if s.Items[i].HasPermission(req) {
			return true, s.Items[i]
		}
	}
	return false, nil
}

// HasPermission 单个角色如何判断有没有权限
func (r *Role) HasPermission(req *PermissionRequest) bool {
	for i := range r.Spec.Permissions {
		fmt.Printf("=======> 数据库查询的单个角色权限: %s \n", r.Spec.Permissions)
		fmt.Printf("=======> 用户请求接口的权限: req: %s \n", req)
		if r.Spec.Permissions[i].HasPermission(req) {
			return true
		}
	}
	return false

}

func (p *Permission) HasPermission(req *PermissionRequest) bool {
	// 放行所有功能
	if p.AllowAll {
		return true
	}

	// 确认是不是同一个服务
	if p.Service != req.Service {
		return false
	}

	// 查询查询所有的功能,确认是否允许
	fmt.Printf("=======> 功能 Featrues: : %s \n", p.Featrues)
	for i := range p.Featrues {
		f := p.Featrues[i]
		if f.Resource == req.Resource && f.Action == req.Action {
			return true
		}
	}

	return false
}

func NewRoleSet() *RoleSet {
	return &RoleSet{
		Items: []*Role{},
	}
}

func (s *RoleSet) Add(item *Role) {
	s.Items = append(s.Items, item)
}

func NewPermissionRequest(service, resource, action string) *PermissionRequest {
	return &PermissionRequest{
		Service:  service,
		Resource: resource,
		Action:   action,
	}
}

func (req *CreateRoleRequest) Validate() error {
	return validate.Struct(req)
}

func NewRole(req *CreateRoleRequest) *Role {
	return &Role{
		Id:       xid.New().String(),
		CreateAt: time.Now().UnixMilli(),
		Spec:     req,
	}
}

func NewCreateRoleRequest() *CreateRoleRequest {
	return &CreateRoleRequest{
		Permissions: []*Permission{},
		Meta:        map[string]string{},
	}
}

func NewDefaultRole() *Role {
	return &Role{
		Spec: NewCreateRoleRequest(),
	}
}
