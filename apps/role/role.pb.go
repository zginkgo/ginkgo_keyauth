// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.1
// source: apps/role/pb/role.proto

package role

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type QueryRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 分页参数
	// @gotags: json:"page"
	Page *PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page"`
	// 根据role名称批量获取role对象
	// @gotags: json:"role_names"
	RoleNames []string `protobuf:"bytes,2,rep,name=role_names,json=roleNames,proto3" json:"role_names"`
}

func (x *QueryRoleRequest) Reset() {
	*x = QueryRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_role_pb_role_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRoleRequest) ProtoMessage() {}

func (x *QueryRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_role_pb_role_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRoleRequest.ProtoReflect.Descriptor instead.
func (*QueryRoleRequest) Descriptor() ([]byte, []int) {
	return file_apps_role_pb_role_proto_rawDescGZIP(), []int{0}
}

func (x *QueryRoleRequest) GetPage() *PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *QueryRoleRequest) GetRoleNames() []string {
	if x != nil {
		return x.RoleNames
	}
	return nil
}

type PageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageSize   uint64 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageNumber uint64 `protobuf:"varint,2,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	Offset     int64  `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *PageRequest) Reset() {
	*x = PageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_role_pb_role_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageRequest) ProtoMessage() {}

func (x *PageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_role_pb_role_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageRequest.ProtoReflect.Descriptor instead.
func (*PageRequest) Descriptor() ([]byte, []int) {
	return file_apps_role_pb_role_proto_rawDescGZIP(), []int{1}
}

func (x *PageRequest) GetPageSize() uint64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *PageRequest) GetPageNumber() uint64 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *PageRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type RoleSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 角色id
	// @gotags: json:"total" bson:"total"
	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total" bson:"total"`
	// 角色id
	// @gotags: json:"items" bson:"items"
	Items []*Role `protobuf:"bytes,2,rep,name=items,proto3" json:"items" bson:"items"`
}

func (x *RoleSet) Reset() {
	*x = RoleSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_role_pb_role_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleSet) ProtoMessage() {}

func (x *RoleSet) ProtoReflect() protoreflect.Message {
	mi := &file_apps_role_pb_role_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleSet.ProtoReflect.Descriptor instead.
func (*RoleSet) Descriptor() ([]byte, []int) {
	return file_apps_role_pb_role_proto_rawDescGZIP(), []int{2}
}

func (x *RoleSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *RoleSet) GetItems() []*Role {
	if x != nil {
		return x.Items
	}
	return nil
}

type Role struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 角色id
	// @gotags: json:"id" bson:"_id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"_id"`
	// 角色创建时间
	// @gotags: json:"name" bson:"create_at"
	CreateAt int64 `protobuf:"varint,2,opt,name=create_at,json=createAt,proto3" json:"name" bson:"create_at"`
	// 角色名称
	// @gotags: json:"spec" bson:"spec"
	Spec *CreateRoleRequest `protobuf:"bytes,3,opt,name=spec,proto3" json:"spec" bson:"spec"`
}

func (x *Role) Reset() {
	*x = Role{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_role_pb_role_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Role) ProtoMessage() {}

func (x *Role) ProtoReflect() protoreflect.Message {
	mi := &file_apps_role_pb_role_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Role.ProtoReflect.Descriptor instead.
func (*Role) Descriptor() ([]byte, []int) {
	return file_apps_role_pb_role_proto_rawDescGZIP(), []int{3}
}

func (x *Role) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Role) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *Role) GetSpec() *CreateRoleRequest {
	if x != nil {
		return x.Spec
	}
	return nil
}

type CreateRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 角色名称
	// @gotags: json:"name" bson:"name"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" bson:"name"`
	// 角色标识
	// @gotags: json:"description" bson:"description"
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description" bson:"description"`
	// 角色的权限
	// @gotags: json:"permissions" bson:"permissions"
	Permissions []*Permission `protobuf:"bytes,3,rep,name=permissions,proto3" json:"permissions" bson:"permissions"`
	// 角色的一些额外属性
	// @gotags: json:"meta" bson:"meta"
	Meta map[string]string `protobuf:"bytes,5,rep,name=meta,proto3" json:"meta" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" bson:"meta"`
}

func (x *CreateRoleRequest) Reset() {
	*x = CreateRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_role_pb_role_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoleRequest) ProtoMessage() {}

func (x *CreateRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_role_pb_role_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoleRequest.ProtoReflect.Descriptor instead.
func (*CreateRoleRequest) Descriptor() ([]byte, []int) {
	return file_apps_role_pb_role_proto_rawDescGZIP(), []int{4}
}

func (x *CreateRoleRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateRoleRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateRoleRequest) GetPermissions() []*Permission {
	if x != nil {
		return x.Permissions
	}
	return nil
}

func (x *CreateRoleRequest) GetMeta() map[string]string {
	if x != nil {
		return x.Meta
	}
	return nil
}

// 那个服务 的哪些权限
type Permission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 服务名称
	// @gotags: json:"service" bson:"service"
	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service" bson:"service"`
	// 是否允许访问所有功能
	// @gotags: json:"allow_all" bson:"allow_all
	AllowAll bool `protobuf:"varint,2,opt,name=allow_all,json=allowAll,proto3" json:"allow_all"`
	// 服务功能
	// @gotags: json:"featrues" bson:"featrues"
	Featrues []*Featrue `protobuf:"bytes,3,rep,name=featrues,proto3" json:"featrues" bson:"featrues"`
}

func (x *Permission) Reset() {
	*x = Permission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_role_pb_role_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Permission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Permission) ProtoMessage() {}

func (x *Permission) ProtoReflect() protoreflect.Message {
	mi := &file_apps_role_pb_role_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Permission.ProtoReflect.Descriptor instead.
func (*Permission) Descriptor() ([]byte, []int) {
	return file_apps_role_pb_role_proto_rawDescGZIP(), []int{5}
}

func (x *Permission) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *Permission) GetAllowAll() bool {
	if x != nil {
		return x.AllowAll
	}
	return false
}

func (x *Permission) GetFeatrues() []*Featrue {
	if x != nil {
		return x.Featrues
	}
	return nil
}

type Featrue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 资源名称
	// @gotags: json:"resource" bson:"resource"
	Resource string `protobuf:"bytes,3,opt,name=resource,proto3" json:"resource" bson:"resource"`
	// 资源操作
	// @gotags: json:"action" bson:"action"
	Action string `protobuf:"bytes,4,opt,name=action,proto3" json:"action" bson:"action"`
}

func (x *Featrue) Reset() {
	*x = Featrue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_role_pb_role_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Featrue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Featrue) ProtoMessage() {}

func (x *Featrue) ProtoReflect() protoreflect.Message {
	mi := &file_apps_role_pb_role_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Featrue.ProtoReflect.Descriptor instead.
func (*Featrue) Descriptor() ([]byte, []int) {
	return file_apps_role_pb_role_proto_rawDescGZIP(), []int{6}
}

func (x *Featrue) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

func (x *Featrue) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

type PermissionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 服务名称
	// @gotags: json:"service"
	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service"`
	// 资源名称
	// @gotags: json:"resource"
	Resource string `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource"`
	// 资源操作
	// @gotags: json:"action"
	Action string `protobuf:"bytes,3,opt,name=action,proto3" json:"action"`
}

func (x *PermissionRequest) Reset() {
	*x = PermissionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_role_pb_role_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PermissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionRequest) ProtoMessage() {}

func (x *PermissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_role_pb_role_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionRequest.ProtoReflect.Descriptor instead.
func (*PermissionRequest) Descriptor() ([]byte, []int) {
	return file_apps_role_pb_role_proto_rawDescGZIP(), []int{7}
}

func (x *PermissionRequest) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *PermissionRequest) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

func (x *PermissionRequest) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

var File_apps_role_pb_role_proto protoreflect.FileDescriptor

var file_apps_role_pb_role_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x72,
	0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x67, 0x69, 0x6e, 0x6b, 0x67,
	0x6f, 0x5f, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x67,
	0x0a, 0x10, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x34, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x67, 0x69, 0x6e, 0x6b, 0x67, 0x6f, 0x5f, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x6f, 0x6c, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x72, 0x6f,
	0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x63, 0x0a, 0x0b, 0x50, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x50, 0x0a, 0x07,
	0x52, 0x6f, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x2f, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67,
	0x69, 0x6e, 0x6b, 0x67, 0x6f, 0x5f, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x72, 0x6f,
	0x6c, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x6f,
	0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x74, 0x12, 0x3a, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x26, 0x2e, 0x67, 0x69, 0x6e, 0x6b, 0x67, 0x6f, 0x5f, 0x6b, 0x65, 0x79, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x22,
	0x8b, 0x02, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x41, 0x0a, 0x0b, 0x70,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x67, 0x69, 0x6e, 0x6b, 0x67, 0x6f, 0x5f, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x44,
	0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x67,
	0x69, 0x6e, 0x6b, 0x67, 0x6f, 0x5f, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x72, 0x6f,
	0x6c, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04,
	0x6d, 0x65, 0x74, 0x61, 0x1a, 0x37, 0x0a, 0x09, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x7d, 0x0a,
	0x0a, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x61,
	0x6c, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x41,
	0x6c, 0x6c, 0x12, 0x38, 0x0a, 0x08, 0x66, 0x65, 0x61, 0x74, 0x72, 0x75, 0x65, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x69, 0x6e, 0x6b, 0x67, 0x6f, 0x5f, 0x6b, 0x65,
	0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x72,
	0x75, 0x65, 0x52, 0x08, 0x66, 0x65, 0x61, 0x74, 0x72, 0x75, 0x65, 0x73, 0x22, 0x3d, 0x0a, 0x07,
	0x46, 0x65, 0x61, 0x74, 0x72, 0x75, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x61, 0x0a, 0x11, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x57,
	0x0a, 0x03, 0x52, 0x50, 0x43, 0x12, 0x50, 0x0a, 0x09, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x6f,
	0x6c, 0x65, 0x12, 0x25, 0x2e, 0x67, 0x69, 0x6e, 0x6b, 0x67, 0x6f, 0x5f, 0x6b, 0x65, 0x79, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x6f,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x67, 0x69, 0x6e, 0x6b,
	0x67, 0x6f, 0x5f, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e,
	0x52, 0x6f, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x67, 0x69, 0x6e, 0x6b, 0x67, 0x6f, 0x2f, 0x67, 0x69,
	0x6e, 0x6b, 0x67, 0x6f, 0x5f, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x70, 0x70,
	0x73, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_role_pb_role_proto_rawDescOnce sync.Once
	file_apps_role_pb_role_proto_rawDescData = file_apps_role_pb_role_proto_rawDesc
)

func file_apps_role_pb_role_proto_rawDescGZIP() []byte {
	file_apps_role_pb_role_proto_rawDescOnce.Do(func() {
		file_apps_role_pb_role_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_role_pb_role_proto_rawDescData)
	})
	return file_apps_role_pb_role_proto_rawDescData
}

var file_apps_role_pb_role_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_apps_role_pb_role_proto_goTypes = []interface{}{
	(*QueryRoleRequest)(nil),  // 0: ginkgo_keyauth.role.QueryRoleRequest
	(*PageRequest)(nil),       // 1: ginkgo_keyauth.role.PageRequest
	(*RoleSet)(nil),           // 2: ginkgo_keyauth.role.RoleSet
	(*Role)(nil),              // 3: ginkgo_keyauth.role.Role
	(*CreateRoleRequest)(nil), // 4: ginkgo_keyauth.role.CreateRoleRequest
	(*Permission)(nil),        // 5: ginkgo_keyauth.role.Permission
	(*Featrue)(nil),           // 6: ginkgo_keyauth.role.Featrue
	(*PermissionRequest)(nil), // 7: ginkgo_keyauth.role.PermissionRequest
	nil,                       // 8: ginkgo_keyauth.role.CreateRoleRequest.MetaEntry
}
var file_apps_role_pb_role_proto_depIdxs = []int32{
	1, // 0: ginkgo_keyauth.role.QueryRoleRequest.page:type_name -> ginkgo_keyauth.role.PageRequest
	3, // 1: ginkgo_keyauth.role.RoleSet.items:type_name -> ginkgo_keyauth.role.Role
	4, // 2: ginkgo_keyauth.role.Role.spec:type_name -> ginkgo_keyauth.role.CreateRoleRequest
	5, // 3: ginkgo_keyauth.role.CreateRoleRequest.permissions:type_name -> ginkgo_keyauth.role.Permission
	8, // 4: ginkgo_keyauth.role.CreateRoleRequest.meta:type_name -> ginkgo_keyauth.role.CreateRoleRequest.MetaEntry
	6, // 5: ginkgo_keyauth.role.Permission.featrues:type_name -> ginkgo_keyauth.role.Featrue
	0, // 6: ginkgo_keyauth.role.RPC.QueryRole:input_type -> ginkgo_keyauth.role.QueryRoleRequest
	2, // 7: ginkgo_keyauth.role.RPC.QueryRole:output_type -> ginkgo_keyauth.role.RoleSet
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_apps_role_pb_role_proto_init() }
func file_apps_role_pb_role_proto_init() {
	if File_apps_role_pb_role_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_role_pb_role_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryRoleRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apps_role_pb_role_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apps_role_pb_role_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleSet); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apps_role_pb_role_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Role); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apps_role_pb_role_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRoleRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apps_role_pb_role_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Permission); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apps_role_pb_role_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Featrue); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apps_role_pb_role_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PermissionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_apps_role_pb_role_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apps_role_pb_role_proto_goTypes,
		DependencyIndexes: file_apps_role_pb_role_proto_depIdxs,
		MessageInfos:      file_apps_role_pb_role_proto_msgTypes,
	}.Build()
	File_apps_role_pb_role_proto = out.File
	file_apps_role_pb_role_proto_rawDesc = nil
	file_apps_role_pb_role_proto_goTypes = nil
	file_apps_role_pb_role_proto_depIdxs = nil
}
