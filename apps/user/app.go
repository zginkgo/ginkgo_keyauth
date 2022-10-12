package user

import (
	"github.com/zginkgo/ginkgo_keyauth/common/utils"
	"github.com/go-playground/validator"
	"github.com/rs/xid"
	"net/http"
	"strconv"
	"time"
)

const (
	AppName = "user"
)

const (
	DefaultDomain = "default"
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

func (req *CreateUserRequest) Validate() error {
	return validate.Struct(req)
}

// NewUser 保存Hash过后的Password
func NewUser(req *CreateUserRequest) *User {
	return &User{
		Id:       xid.New().String(),
		CreateAt: time.Now().UnixMilli(),
		Data:     req,
	}
}

func (u *User) CheckPassword(password string) bool {
	return utils.CheckPasswordHash(password, u.Data.Password)
}

func NewUserSet() *UserSet {
	return &UserSet{
		Items: []*User{},
	}
}

func (s *UserSet) Add(item *User) {
	s.Items = append(s.Items, item)
}

func NewDefaultUser() *User {
	return NewUser(NewCreateUserRequest())
}

func NewCreateUserRequest() *CreateUserRequest {
	return &CreateUserRequest{
		Domain: DefaultDomain,
	}
}

func NewQueryUserRequestFromHTTP(r *http.Request) *QueryUserRequest {
	qs := r.URL.Query()

	return &QueryUserRequest{
		Page:     NewPageRequestFromHTTP(r),
		Keywords: qs.Get("keywords"),
	}
}

// NewPageRequestFromHTTP 从HTTP请求中加载分页请求
func NewPageRequestFromHTTP(req *http.Request) *PageRequest {
	qs := req.URL.Query()

	ps := qs.Get("page_size")
	pn := qs.Get("page_number")
	os := qs.Get("offset")

	psUint64, _ := strconv.ParseUint(ps, 10, 64)
	pnUint64, _ := strconv.ParseUint(pn, 10, 64)
	osInt64, _ := strconv.ParseInt(os, 10, 64)

	if psUint64 == 0 {
		psUint64 = DefaultPageSize
	}
	if pnUint64 == 0 {
		pnUint64 = DefaultPageNumber
	}

	return &PageRequest{
		PageSize:   psUint64,
		PageNumber: pnUint64,
		Offset:     osInt64,
	}
}

func NewPutUserRequest(id string) *UpdateUserRequest {
	return &UpdateUserRequest{
		Id:         id,
		UpdateMode: UpdateMode_PUT,
		UpdateAt:   time.Now().UnixMicro(),
		Data:       NewCreateUserRequest(),
	}
}

func NewPatchUserRequest(id string) *UpdateUserRequest {
	return &UpdateUserRequest{
		Id:         id,
		UpdateMode: UpdateMode_PATCH,
		UpdateAt:   time.Now().UnixMicro(),
		Data:       NewCreateUserRequest(),
	}
}

func NewDescribeUserRequestById(id string) *DescribeUserRequest {
	return &DescribeUserRequest{
		DescribeBy: DescribeBy_USER_ID,
		UserId:     id,
	}
}

func NewDescribeUserRequestByName(domain, name string) *DescribeUserRequest {
	return &DescribeUserRequest{
		DescribeBy: DescribeBy_USER_NAME,
		Domain:     domain,
		UserName:   name,
	}
}
