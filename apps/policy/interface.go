package policy

import (
	"context"
	"github.com/go-playground/validator"
	"github.com/rs/xid"
	"time"
)

const (
	AppName = "policy"
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
	CreatePolicy(context.Context, *CreatePolicyRequest) (*Policy, error)
	RPCServer
}

func NewQueryPolicyRequest() *QueryPolicyRequest {
	return &QueryPolicyRequest{
		Page: NewDefaultPageRequest(),
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

func NewDefaultPolicy() *Policy {
	return &Policy{
		Spec: NewCreatePolicyRequest(),
	}
}

func NewCreatePolicyRequest() *CreatePolicyRequest {
	return &CreatePolicyRequest{}
}

func NewPolicySet() *PolicySet {
	return &PolicySet{
		Items: []*Policy{},
	}
}

func (s *PolicySet) Add(item *Policy) {
	s.Items = append(s.Items, item)
}

// Roles role 名称的列表
func (s *PolicySet) Roles() (roles []string) {
	for i := range s.Items {
		roles = append(roles, s.Items[i].Spec.Role)
	}
	return
}

// GetPolicyByRole 根据Role名称查询集合里面的策略
func (s *PolicySet) GetPolicyByRole(role string) *Policy {
	for i := range s.Items {
		if s.Items[i].Spec.Role == role {
			return s.Items[i]
		}
	}

	return nil
}

func (req *CreatePolicyRequest) Validate() error {
	return validate.Struct(req)
}

func NewPolicy(req *CreatePolicyRequest) (*Policy, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	return &Policy{
		Id:       xid.New().String(),
		CreateAt: time.Now().UnixMilli(),
		Spec:     req,
	}, nil
}

func NewValidatePermissionRequest() *ValidatePermissionRequest {
	return &ValidatePermissionRequest{}
}
