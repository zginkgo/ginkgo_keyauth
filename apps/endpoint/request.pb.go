// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.1
// source: apps/endpoint/pb/request.proto

package endpoint

import (
	http "github.com/zginkgo/ginkgo_mcube/pb/http"
	page "github.com/zginkgo/ginkgo_mcube/pb/page"
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

// RegistryRequest 服务注册请求
type RegistryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"client_id" validate:"required"
	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id" validate:"required"`
	// @gotags: json:"client_secret" validate:"required"
	ClientSecret string `protobuf:"bytes,2,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret" validate:"required"`
	// @gotags: json:"version" validate:"required,lte=32"
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version" validate:"required,lte=32"`
	// @gotags: json:"entries"
	Entries []*http.Entry `protobuf:"bytes,4,rep,name=entries,proto3" json:"entries"`
}

func (x *RegistryRequest) Reset() {
	*x = RegistryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_endpoint_pb_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegistryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistryRequest) ProtoMessage() {}

func (x *RegistryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_endpoint_pb_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistryRequest.ProtoReflect.Descriptor instead.
func (*RegistryRequest) Descriptor() ([]byte, []int) {
	return file_apps_endpoint_pb_request_proto_rawDescGZIP(), []int{0}
}

func (x *RegistryRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *RegistryRequest) GetClientSecret() string {
	if x != nil {
		return x.ClientSecret
	}
	return ""
}

func (x *RegistryRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *RegistryRequest) GetEntries() []*http.Entry {
	if x != nil {
		return x.Entries
	}
	return nil
}

// RegistryReponse todo
type RegistryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"message"
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message"`
}

func (x *RegistryResponse) Reset() {
	*x = RegistryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_endpoint_pb_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegistryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistryResponse) ProtoMessage() {}

func (x *RegistryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apps_endpoint_pb_request_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistryResponse.ProtoReflect.Descriptor instead.
func (*RegistryResponse) Descriptor() ([]byte, []int) {
	return file_apps_endpoint_pb_request_proto_rawDescGZIP(), []int{1}
}

func (x *RegistryResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// DescribeEndpointRequest todo
type DescribeEndpointRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
}

func (x *DescribeEndpointRequest) Reset() {
	*x = DescribeEndpointRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_endpoint_pb_request_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeEndpointRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeEndpointRequest) ProtoMessage() {}

func (x *DescribeEndpointRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_endpoint_pb_request_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeEndpointRequest.ProtoReflect.Descriptor instead.
func (*DescribeEndpointRequest) Descriptor() ([]byte, []int) {
	return file_apps_endpoint_pb_request_proto_rawDescGZIP(), []int{2}
}

func (x *DescribeEndpointRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// QueryEndpointRequest 查询应用列表
type QueryEndpointRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"page"
	Page *page.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page"`
	// @gotags: json:"service_ids"
	ServiceIds []string `protobuf:"bytes,2,rep,name=service_ids,json=serviceIds,proto3" json:"service_ids"`
	// @gotags: json:"path"
	Path string `protobuf:"bytes,3,opt,name=path,proto3" json:"path"`
	// @gotags: json:"method"
	Method string `protobuf:"bytes,4,opt,name=method,proto3" json:"method"`
	// @gotags: json:"function_name"
	FunctionName string `protobuf:"bytes,5,opt,name=function_name,json=functionName,proto3" json:"function_name"`
	// @gotags: json:"resources"
	Resources []string `protobuf:"bytes,6,rep,name=resources,proto3" json:"resources"`
	// @gotags: json:"labels"
	Labels map[string]string `protobuf:"bytes,7,rep,name=labels,proto3" json:"labels" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// @gotags: json:"permission_enable"
	PermissionEnable BoolQuery `protobuf:"varint,8,opt,name=permission_enable,json=permissionEnable,proto3,enum=zginkgo.ginkgo_keyauth.endpoint.BoolQuery" json:"permission_enable"`
}

func (x *QueryEndpointRequest) Reset() {
	*x = QueryEndpointRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_endpoint_pb_request_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryEndpointRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryEndpointRequest) ProtoMessage() {}

func (x *QueryEndpointRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_endpoint_pb_request_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryEndpointRequest.ProtoReflect.Descriptor instead.
func (*QueryEndpointRequest) Descriptor() ([]byte, []int) {
	return file_apps_endpoint_pb_request_proto_rawDescGZIP(), []int{3}
}

func (x *QueryEndpointRequest) GetPage() *page.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *QueryEndpointRequest) GetServiceIds() []string {
	if x != nil {
		return x.ServiceIds
	}
	return nil
}

func (x *QueryEndpointRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *QueryEndpointRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *QueryEndpointRequest) GetFunctionName() string {
	if x != nil {
		return x.FunctionName
	}
	return ""
}

func (x *QueryEndpointRequest) GetResources() []string {
	if x != nil {
		return x.Resources
	}
	return nil
}

func (x *QueryEndpointRequest) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *QueryEndpointRequest) GetPermissionEnable() BoolQuery {
	if x != nil {
		return x.PermissionEnable
	}
	return BoolQuery_ALL
}

// DeleteEndpointRequest todo
type DeleteEndpointRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"service_id"
	ServiceId string `protobuf:"bytes,1,opt,name=service_id,json=serviceId,proto3" json:"service_id"`
}

func (x *DeleteEndpointRequest) Reset() {
	*x = DeleteEndpointRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_endpoint_pb_request_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEndpointRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEndpointRequest) ProtoMessage() {}

func (x *DeleteEndpointRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_endpoint_pb_request_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEndpointRequest.ProtoReflect.Descriptor instead.
func (*DeleteEndpointRequest) Descriptor() ([]byte, []int) {
	return file_apps_endpoint_pb_request_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteEndpointRequest) GetServiceId() string {
	if x != nil {
		return x.ServiceId
	}
	return ""
}

// QueryResourceRequest todo
type QueryResourceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"page"
	Page *page.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page"`
	// @gotags: json:"service_ids"
	ServiceIds []string `protobuf:"bytes,2,rep,name=service_ids,json=serviceIds,proto3" json:"service_ids"`
	// @gotags: json:"resources"
	Resources []string `protobuf:"bytes,3,rep,name=resources,proto3" json:"resources"`
	// @gotags: json:"permission_enable"
	PermissionEnable BoolQuery `protobuf:"varint,4,opt,name=permission_enable,json=permissionEnable,proto3,enum=zginkgo.ginkgo_keyauth.endpoint.BoolQuery" json:"permission_enable"`
}

func (x *QueryResourceRequest) Reset() {
	*x = QueryResourceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_endpoint_pb_request_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryResourceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryResourceRequest) ProtoMessage() {}

func (x *QueryResourceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_endpoint_pb_request_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryResourceRequest.ProtoReflect.Descriptor instead.
func (*QueryResourceRequest) Descriptor() ([]byte, []int) {
	return file_apps_endpoint_pb_request_proto_rawDescGZIP(), []int{5}
}

func (x *QueryResourceRequest) GetPage() *page.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *QueryResourceRequest) GetServiceIds() []string {
	if x != nil {
		return x.ServiceIds
	}
	return nil
}

func (x *QueryResourceRequest) GetResources() []string {
	if x != nil {
		return x.Resources
	}
	return nil
}

func (x *QueryResourceRequest) GetPermissionEnable() BoolQuery {
	if x != nil {
		return x.PermissionEnable
	}
	return BoolQuery_ALL
}

var File_apps_endpoint_pb_request_proto protoreflect.FileDescriptor

var file_apps_endpoint_pb_request_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f,
	0x70, 0x62, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1f, 0x7a, 0x67, 0x69, 0x6e, 0x6b, 0x67, 0x6f, 0x2e, 0x67, 0x69, 0x6e, 0x6b, 0x67, 0x6f,
	0x5f, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x1a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x67,
	0x69, 0x6e, 0x6b, 0x67, 0x6f, 0x2f, 0x67, 0x69, 0x6e, 0x6b, 0x67, 0x6f, 0x5f, 0x6d, 0x63, 0x75,
	0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x7a, 0x67, 0x69, 0x6e, 0x6b, 0x67, 0x6f, 0x2f, 0x67, 0x69, 0x6e, 0x6b, 0x67, 0x6f,
	0x5f, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x65,
	0x6e, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x61, 0x70, 0x70, 0x73,
	0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x70, 0x62, 0x2f, 0x65, 0x6e, 0x75,
	0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa9, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3a, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x7a, 0x67, 0x69, 0x6e, 0x6b,
	0x67, 0x6f, 0x2e, 0x67, 0x69, 0x6e, 0x6b, 0x67, 0x6f, 0x5f, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e,
	0x68, 0x74, 0x74, 0x70, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72,
	0x69, 0x65, 0x73, 0x22, 0x2c, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x29, 0x0a, 0x17, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x45, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xd1, 0x03, 0x0a,
	0x14, 0x51, 0x75, 0x65, 0x72, 0x79, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x7a, 0x67, 0x69, 0x6e, 0x6b, 0x67, 0x6f, 0x2e, 0x67, 0x69,
	0x6e, 0x6b, 0x67, 0x6f, 0x5f, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e,
	0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49,
	0x64, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x23,
	0x0a, 0x0d, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x12, 0x59, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x41, 0x2e, 0x7a, 0x67, 0x69, 0x6e, 0x6b, 0x67, 0x6f, 0x2e, 0x67, 0x69, 0x6e, 0x6b,
	0x67, 0x6f, 0x5f, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x57, 0x0a, 0x11,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x7a, 0x67, 0x69, 0x6e, 0x6b, 0x67,
	0x6f, 0x2e, 0x67, 0x69, 0x6e, 0x6b, 0x67, 0x6f, 0x5f, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x52, 0x10, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x45,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x36, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x22, 0xea, 0x01, 0x0a, 0x14, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x3a, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x26, 0x2e, 0x7a, 0x67, 0x69, 0x6e, 0x6b, 0x67, 0x6f, 0x2e, 0x67, 0x69, 0x6e, 0x6b, 0x67, 0x6f,
	0x5f, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x73, 0x12, 0x1c,
	0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x57, 0x0a, 0x11,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x7a, 0x67, 0x69, 0x6e, 0x6b, 0x67,
	0x6f, 0x2e, 0x67, 0x69, 0x6e, 0x6b, 0x67, 0x6f, 0x5f, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x52, 0x10, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x45,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x67, 0x69, 0x6e, 0x6b, 0x67, 0x6f, 0x2f, 0x67, 0x69, 0x6e, 0x6b,
	0x67, 0x6f, 0x5f, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_endpoint_pb_request_proto_rawDescOnce sync.Once
	file_apps_endpoint_pb_request_proto_rawDescData = file_apps_endpoint_pb_request_proto_rawDesc
)

func file_apps_endpoint_pb_request_proto_rawDescGZIP() []byte {
	file_apps_endpoint_pb_request_proto_rawDescOnce.Do(func() {
		file_apps_endpoint_pb_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_endpoint_pb_request_proto_rawDescData)
	})
	return file_apps_endpoint_pb_request_proto_rawDescData
}

var file_apps_endpoint_pb_request_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_apps_endpoint_pb_request_proto_goTypes = []interface{}{
	(*RegistryRequest)(nil),         // 0: zginkgo.ginkgo_keyauth.endpoint.RegistryRequest
	(*RegistryResponse)(nil),        // 1: zginkgo.ginkgo_keyauth.endpoint.RegistryResponse
	(*DescribeEndpointRequest)(nil), // 2: zginkgo.ginkgo_keyauth.endpoint.DescribeEndpointRequest
	(*QueryEndpointRequest)(nil),    // 3: zginkgo.ginkgo_keyauth.endpoint.QueryEndpointRequest
	(*DeleteEndpointRequest)(nil),   // 4: zginkgo.ginkgo_keyauth.endpoint.DeleteEndpointRequest
	(*QueryResourceRequest)(nil),    // 5: zginkgo.ginkgo_keyauth.endpoint.QueryResourceRequest
	nil,                             // 6: zginkgo.ginkgo_keyauth.endpoint.QueryEndpointRequest.LabelsEntry
	(*http.Entry)(nil),              // 7: zginkgo.ginkgo_mcube.http.Entry
	(*page.PageRequest)(nil),        // 8: zginkgo.ginkgo_mcube.page.PageRequest
	(BoolQuery)(0),                  // 9: zginkgo.ginkgo_keyauth.endpoint.BoolQuery
}
var file_apps_endpoint_pb_request_proto_depIdxs = []int32{
	7, // 0: zginkgo.ginkgo_keyauth.endpoint.RegistryRequest.entries:type_name -> zginkgo.ginkgo_mcube.http.Entry
	8, // 1: zginkgo.ginkgo_keyauth.endpoint.QueryEndpointRequest.page:type_name -> zginkgo.ginkgo_mcube.page.PageRequest
	6, // 2: zginkgo.ginkgo_keyauth.endpoint.QueryEndpointRequest.labels:type_name -> zginkgo.ginkgo_keyauth.endpoint.QueryEndpointRequest.LabelsEntry
	9, // 3: zginkgo.ginkgo_keyauth.endpoint.QueryEndpointRequest.permission_enable:type_name -> zginkgo.ginkgo_keyauth.endpoint.BoolQuery
	8, // 4: zginkgo.ginkgo_keyauth.endpoint.QueryResourceRequest.page:type_name -> zginkgo.ginkgo_mcube.page.PageRequest
	9, // 5: zginkgo.ginkgo_keyauth.endpoint.QueryResourceRequest.permission_enable:type_name -> zginkgo.ginkgo_keyauth.endpoint.BoolQuery
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_apps_endpoint_pb_request_proto_init() }
func file_apps_endpoint_pb_request_proto_init() {
	if File_apps_endpoint_pb_request_proto != nil {
		return
	}
	file_apps_endpoint_pb_enum_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_apps_endpoint_pb_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegistryRequest); i {
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
		file_apps_endpoint_pb_request_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegistryResponse); i {
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
		file_apps_endpoint_pb_request_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeEndpointRequest); i {
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
		file_apps_endpoint_pb_request_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryEndpointRequest); i {
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
		file_apps_endpoint_pb_request_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEndpointRequest); i {
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
		file_apps_endpoint_pb_request_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryResourceRequest); i {
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
			RawDescriptor: file_apps_endpoint_pb_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apps_endpoint_pb_request_proto_goTypes,
		DependencyIndexes: file_apps_endpoint_pb_request_proto_depIdxs,
		MessageInfos:      file_apps_endpoint_pb_request_proto_msgTypes,
	}.Build()
	File_apps_endpoint_pb_request_proto = out.File
	file_apps_endpoint_pb_request_proto_rawDesc = nil
	file_apps_endpoint_pb_request_proto_goTypes = nil
	file_apps_endpoint_pb_request_proto_depIdxs = nil
}
