// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dsrhub.proto

package idl_dsrhub

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CreateDSRRequest struct {
	Regulation           string   `protobuf:"bytes,1,opt,name=regulation,proto3" json:"regulation,omitempty"`
	StatusCallbackUrl    string   `protobuf:"bytes,2,opt,name=status_callback_url,json=statusCallbackUrl,proto3" json:"status_callback_url,omitempty"`
	SubjectRequestId     string   `protobuf:"bytes,3,opt,name=subject_request_id,json=subjectRequestId,proto3" json:"subject_request_id,omitempty"`
	SubjectRequestType   string   `protobuf:"bytes,4,opt,name=subject_request_type,json=subjectRequestType,proto3" json:"subject_request_type,omitempty"`
	IdentityType         string   `protobuf:"bytes,5,opt,name=identity_type,json=identityType,proto3" json:"identity_type,omitempty"`
	IdentityFormat       string   `protobuf:"bytes,6,opt,name=identity_format,json=identityFormat,proto3" json:"identity_format,omitempty"`
	IdentityValue        string   `protobuf:"bytes,7,opt,name=identity_value,json=identityValue,proto3" json:"identity_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateDSRRequest) Reset()         { *m = CreateDSRRequest{} }
func (m *CreateDSRRequest) String() string { return proto.CompactTextString(m) }
func (*CreateDSRRequest) ProtoMessage()    {}
func (*CreateDSRRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_872837866b3f167b, []int{0}
}

func (m *CreateDSRRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateDSRRequest.Unmarshal(m, b)
}
func (m *CreateDSRRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateDSRRequest.Marshal(b, m, deterministic)
}
func (m *CreateDSRRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateDSRRequest.Merge(m, src)
}
func (m *CreateDSRRequest) XXX_Size() int {
	return xxx_messageInfo_CreateDSRRequest.Size(m)
}
func (m *CreateDSRRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateDSRRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateDSRRequest proto.InternalMessageInfo

func (m *CreateDSRRequest) GetRegulation() string {
	if m != nil {
		return m.Regulation
	}
	return ""
}

func (m *CreateDSRRequest) GetStatusCallbackUrl() string {
	if m != nil {
		return m.StatusCallbackUrl
	}
	return ""
}

func (m *CreateDSRRequest) GetSubjectRequestId() string {
	if m != nil {
		return m.SubjectRequestId
	}
	return ""
}

func (m *CreateDSRRequest) GetSubjectRequestType() string {
	if m != nil {
		return m.SubjectRequestType
	}
	return ""
}

func (m *CreateDSRRequest) GetIdentityType() string {
	if m != nil {
		return m.IdentityType
	}
	return ""
}

func (m *CreateDSRRequest) GetIdentityFormat() string {
	if m != nil {
		return m.IdentityFormat
	}
	return ""
}

func (m *CreateDSRRequest) GetIdentityValue() string {
	if m != nil {
		return m.IdentityValue
	}
	return ""
}

type CreateDSRResponse struct {
	Regulation           string   `protobuf:"bytes,1,opt,name=regulation,proto3" json:"regulation,omitempty"`
	RequestStatus        string   `protobuf:"bytes,2,opt,name=request_status,json=requestStatus,proto3" json:"request_status,omitempty"`
	SubjectRequestId     string   `protobuf:"bytes,3,opt,name=subject_request_id,json=subjectRequestId,proto3" json:"subject_request_id,omitempty"`
	SubjectRequestType   string   `protobuf:"bytes,4,opt,name=subject_request_type,json=subjectRequestType,proto3" json:"subject_request_type,omitempty"`
	IdentityType         string   `protobuf:"bytes,5,opt,name=identity_type,json=identityType,proto3" json:"identity_type,omitempty"`
	IdentityFormat       string   `protobuf:"bytes,6,opt,name=identity_format,json=identityFormat,proto3" json:"identity_format,omitempty"`
	IdentityValue        string   `protobuf:"bytes,7,opt,name=identity_value,json=identityValue,proto3" json:"identity_value,omitempty"`
	ControllerId         string   `protobuf:"bytes,8,opt,name=controller_id,json=controllerId,proto3" json:"controller_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateDSRResponse) Reset()         { *m = CreateDSRResponse{} }
func (m *CreateDSRResponse) String() string { return proto.CompactTextString(m) }
func (*CreateDSRResponse) ProtoMessage()    {}
func (*CreateDSRResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_872837866b3f167b, []int{1}
}

func (m *CreateDSRResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateDSRResponse.Unmarshal(m, b)
}
func (m *CreateDSRResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateDSRResponse.Marshal(b, m, deterministic)
}
func (m *CreateDSRResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateDSRResponse.Merge(m, src)
}
func (m *CreateDSRResponse) XXX_Size() int {
	return xxx_messageInfo_CreateDSRResponse.Size(m)
}
func (m *CreateDSRResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateDSRResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateDSRResponse proto.InternalMessageInfo

func (m *CreateDSRResponse) GetRegulation() string {
	if m != nil {
		return m.Regulation
	}
	return ""
}

func (m *CreateDSRResponse) GetRequestStatus() string {
	if m != nil {
		return m.RequestStatus
	}
	return ""
}

func (m *CreateDSRResponse) GetSubjectRequestId() string {
	if m != nil {
		return m.SubjectRequestId
	}
	return ""
}

func (m *CreateDSRResponse) GetSubjectRequestType() string {
	if m != nil {
		return m.SubjectRequestType
	}
	return ""
}

func (m *CreateDSRResponse) GetIdentityType() string {
	if m != nil {
		return m.IdentityType
	}
	return ""
}

func (m *CreateDSRResponse) GetIdentityFormat() string {
	if m != nil {
		return m.IdentityFormat
	}
	return ""
}

func (m *CreateDSRResponse) GetIdentityValue() string {
	if m != nil {
		return m.IdentityValue
	}
	return ""
}

func (m *CreateDSRResponse) GetControllerId() string {
	if m != nil {
		return m.ControllerId
	}
	return ""
}

type ExchangeIdentityRequest struct {
	IdentityType         string   `protobuf:"bytes,1,opt,name=identity_type,json=identityType,proto3" json:"identity_type,omitempty"`
	IdentityFormat       string   `protobuf:"bytes,2,opt,name=identity_format,json=identityFormat,proto3" json:"identity_format,omitempty"`
	IdentityValue        string   `protobuf:"bytes,3,opt,name=identity_value,json=identityValue,proto3" json:"identity_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExchangeIdentityRequest) Reset()         { *m = ExchangeIdentityRequest{} }
func (m *ExchangeIdentityRequest) String() string { return proto.CompactTextString(m) }
func (*ExchangeIdentityRequest) ProtoMessage()    {}
func (*ExchangeIdentityRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_872837866b3f167b, []int{2}
}

func (m *ExchangeIdentityRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExchangeIdentityRequest.Unmarshal(m, b)
}
func (m *ExchangeIdentityRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExchangeIdentityRequest.Marshal(b, m, deterministic)
}
func (m *ExchangeIdentityRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExchangeIdentityRequest.Merge(m, src)
}
func (m *ExchangeIdentityRequest) XXX_Size() int {
	return xxx_messageInfo_ExchangeIdentityRequest.Size(m)
}
func (m *ExchangeIdentityRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExchangeIdentityRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExchangeIdentityRequest proto.InternalMessageInfo

func (m *ExchangeIdentityRequest) GetIdentityType() string {
	if m != nil {
		return m.IdentityType
	}
	return ""
}

func (m *ExchangeIdentityRequest) GetIdentityFormat() string {
	if m != nil {
		return m.IdentityFormat
	}
	return ""
}

func (m *ExchangeIdentityRequest) GetIdentityValue() string {
	if m != nil {
		return m.IdentityValue
	}
	return ""
}

type ExchangeIdentityResponse struct {
	IdentityType         string   `protobuf:"bytes,1,opt,name=identity_type,json=identityType,proto3" json:"identity_type,omitempty"`
	IdentityFormat       string   `protobuf:"bytes,2,opt,name=identity_format,json=identityFormat,proto3" json:"identity_format,omitempty"`
	IdentityValue        string   `protobuf:"bytes,3,opt,name=identity_value,json=identityValue,proto3" json:"identity_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExchangeIdentityResponse) Reset()         { *m = ExchangeIdentityResponse{} }
func (m *ExchangeIdentityResponse) String() string { return proto.CompactTextString(m) }
func (*ExchangeIdentityResponse) ProtoMessage()    {}
func (*ExchangeIdentityResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_872837866b3f167b, []int{3}
}

func (m *ExchangeIdentityResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExchangeIdentityResponse.Unmarshal(m, b)
}
func (m *ExchangeIdentityResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExchangeIdentityResponse.Marshal(b, m, deterministic)
}
func (m *ExchangeIdentityResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExchangeIdentityResponse.Merge(m, src)
}
func (m *ExchangeIdentityResponse) XXX_Size() int {
	return xxx_messageInfo_ExchangeIdentityResponse.Size(m)
}
func (m *ExchangeIdentityResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExchangeIdentityResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExchangeIdentityResponse proto.InternalMessageInfo

func (m *ExchangeIdentityResponse) GetIdentityType() string {
	if m != nil {
		return m.IdentityType
	}
	return ""
}

func (m *ExchangeIdentityResponse) GetIdentityFormat() string {
	if m != nil {
		return m.IdentityFormat
	}
	return ""
}

func (m *ExchangeIdentityResponse) GetIdentityValue() string {
	if m != nil {
		return m.IdentityValue
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateDSRRequest)(nil), "idl.dsrhub.CreateDSRRequest")
	proto.RegisterType((*CreateDSRResponse)(nil), "idl.dsrhub.CreateDSRResponse")
	proto.RegisterType((*ExchangeIdentityRequest)(nil), "idl.dsrhub.ExchangeIdentityRequest")
	proto.RegisterType((*ExchangeIdentityResponse)(nil), "idl.dsrhub.ExchangeIdentityResponse")
}

func init() { proto.RegisterFile("dsrhub.proto", fileDescriptor_872837866b3f167b) }

var fileDescriptor_872837866b3f167b = []byte{
	// 577 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x94, 0x4d, 0x6f, 0xd3, 0x30,
	0x18, 0xc7, 0xe5, 0x0c, 0x0a, 0x7b, 0xb4, 0x8e, 0xd6, 0x43, 0x10, 0xaa, 0x31, 0xaa, 0x74, 0x15,
	0xa8, 0xda, 0x1a, 0x54, 0x2e, 0xa8, 0xb7, 0xb2, 0x81, 0x28, 0x42, 0x02, 0xb5, 0xc0, 0x81, 0x4b,
	0xe4, 0x24, 0x26, 0xcd, 0xf0, 0xe2, 0x60, 0x3b, 0x83, 0x5e, 0x41, 0x42, 0x88, 0x0b, 0x52, 0xf9,
	0x0e, 0x7c, 0xa0, 0xc2, 0x47, 0xe0, 0x83, 0xa0, 0x3a, 0x49, 0x57, 0xfa, 0x02, 0xbb, 0x21, 0x71,
	0x4a, 0xf5, 0x7f, 0x7e, 0xf6, 0xf3, 0x7f, 0x5e, 0x6a, 0xd8, 0xf0, 0xa5, 0x18, 0x24, 0x6e, 0x33,
	0x16, 0x5c, 0x71, 0x0c, 0xa1, 0xcf, 0x9a, 0xa9, 0x52, 0xd9, 0x0e, 0x38, 0x0f, 0x18, 0xb5, 0x49,
	0x1c, 0xda, 0x24, 0x8a, 0xb8, 0x22, 0x2a, 0xe4, 0x91, 0x4c, 0xc9, 0xca, 0x9e, 0xfe, 0x78, 0xfb,
	0x01, 0x8d, 0xf6, 0xe5, 0x5b, 0x12, 0x04, 0x54, 0xd8, 0x3c, 0xd6, 0xc4, 0x22, 0x6d, 0x7d, 0x5a,
	0x83, 0xd2, 0x81, 0xa0, 0x44, 0xd1, 0xc3, 0x7e, 0xaf, 0x47, 0xdf, 0x24, 0x54, 0x2a, 0xbc, 0x03,
	0x20, 0x68, 0x90, 0x30, 0x4d, 0x9a, 0xa8, 0x8a, 0x6e, 0xad, 0xf7, 0x66, 0x14, 0xdc, 0x84, 0x2d,
	0xa9, 0x88, 0x4a, 0xa4, 0xe3, 0x11, 0xc6, 0x5c, 0xe2, 0xbd, 0x76, 0x12, 0xc1, 0x4c, 0x43, 0x83,
	0xe5, 0x34, 0x74, 0x90, 0x45, 0x9e, 0x0b, 0x86, 0xf7, 0x00, 0xcb, 0xc4, 0x3d, 0xa2, 0x9e, 0x72,
	0x44, 0x9a, 0xc2, 0x09, 0x7d, 0x73, 0x4d, 0xe3, 0xa5, 0x2c, 0x92, 0xe5, 0xee, 0xfa, 0xf8, 0x36,
	0x5c, 0x9e, 0xa7, 0xd5, 0x30, 0xa6, 0xe6, 0x39, 0xcd, 0xe3, 0xdf, 0xf9, 0x67, 0xc3, 0x98, 0xe2,
	0x1a, 0x14, 0x43, 0x9f, 0x46, 0x2a, 0x54, 0xc3, 0x14, 0x3d, 0xaf, 0xd1, 0x8d, 0x5c, 0xd4, 0xd0,
	0x4d, 0xb8, 0x34, 0x85, 0x5e, 0x71, 0x71, 0x4c, 0x94, 0x59, 0xd0, 0xd8, 0x66, 0x2e, 0x3f, 0xd0,
	0x2a, 0xae, 0xc3, 0x54, 0x71, 0x4e, 0x08, 0x4b, 0xa8, 0x79, 0x41, 0x73, 0xd3, 0x1c, 0x2f, 0x26,
	0x62, 0xfb, 0xc9, 0xa8, 0xf3, 0x18, 0x1e, 0x8d, 0xd1, 0x4c, 0x63, 0xc6, 0x68, 0x49, 0x9d, 0x63,
	0xb4, 0xb4, 0x9c, 0x31, 0x9a, 0xcb, 0x62, 0xfd, 0x30, 0xa0, 0x3c, 0x33, 0x0a, 0x19, 0xf3, 0x48,
	0xd2, 0xbf, 0xce, 0xa2, 0x0e, 0x9b, 0xf9, 0xb5, 0x69, 0xe3, 0xb3, 0x31, 0x14, 0x33, 0xb5, 0xaf,
	0xc5, 0xff, 0x72, 0x04, 0x93, 0xa4, 0x1e, 0x8f, 0x94, 0xe0, 0x8c, 0x51, 0x31, 0xa9, 0xe7, 0x62,
	0x9a, 0xf4, 0x54, 0xec, 0xfa, 0xd6, 0x37, 0x04, 0x57, 0xef, 0xbf, 0xf3, 0x06, 0x24, 0x0a, 0x68,
	0x37, 0x3b, 0x9e, 0x2f, 0xfa, 0x82, 0x6b, 0x74, 0x36, 0xd7, 0xc6, 0x19, 0x5d, 0xaf, 0x2d, 0x5b,
	0x9c, 0x2b, 0xa3, 0xce, 0x16, 0x94, 0x17, 0xe7, 0xff, 0x05, 0x81, 0xb9, 0x68, 0x34, 0x5b, 0x83,
	0x7f, 0xe1, 0xb4, 0xf5, 0xd1, 0x80, 0xe2, 0x61, 0xbf, 0xf7, 0x30, 0x71, 0xfb, 0x54, 0x9c, 0x84,
	0x1e, 0xc5, 0x1c, 0xd6, 0xa7, 0x2b, 0x8a, 0xb7, 0x9b, 0xa7, 0x8f, 0x52, 0x73, 0xfe, 0x11, 0xa9,
	0x5c, 0x5f, 0x11, 0x4d, 0x0b, 0xb2, 0xea, 0xef, 0xbf, 0xff, 0xfc, 0x6a, 0xdc, 0xb0, 0x2a, 0x76,
	0x8a, 0xd8, 0x9e, 0x46, 0x1c, 0x5f, 0x8a, 0x7c, 0xe7, 0xda, 0xa8, 0x81, 0x3f, 0x20, 0x28, 0xcd,
	0x37, 0x05, 0xd7, 0x66, 0xaf, 0x5e, 0x31, 0xdb, 0xca, 0xee, 0x9f, 0xa1, 0xcc, 0xc6, 0xae, 0xb6,
	0xb1, 0x63, 0x5d, 0xcb, 0x6d, 0xd0, 0x8c, 0x74, 0xf2, 0x56, 0xb4, 0x51, 0xe3, 0xde, 0x67, 0xf4,
	0xb2, 0x16, 0x84, 0x6a, 0x72, 0x93, 0xc7, 0x8f, 0x73, 0x34, 0xfb, 0x84, 0x3e, 0x73, 0xd2, 0x9f,
	0xa3, 0x8e, 0x8b, 0xef, 0x02, 0xa4, 0x2d, 0xab, 0x76, 0x9e, 0x76, 0x2d, 0x0b, 0x57, 0x07, 0x4a,
	0xc5, 0xb2, 0x6d, 0xdb, 0xab, 0xce, 0xb7, 0x0a, 0x8c, 0x28, 0x2a, 0x55, 0xc3, 0x40, 0x46, 0xab,
	0x44, 0xe2, 0x98, 0x85, 0x9e, 0xfe, 0x6b, 0xdb, 0x47, 0x92, 0x47, 0xed, 0x05, 0xc5, 0x2d, 0xe8,
	0x97, 0xfb, 0xce, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xda, 0xa1, 0x88, 0x52, 0x21, 0x06, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DSRHubServiceClient is the client API for DSRHubService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DSRHubServiceClient interface {
	CreateDSR(ctx context.Context, in *CreateDSRRequest, opts ...grpc.CallOption) (*CreateDSRResponse, error)
	ExchangeIdentity(ctx context.Context, in *ExchangeIdentityRequest, opts ...grpc.CallOption) (*ExchangeIdentityResponse, error)
}

type dSRHubServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDSRHubServiceClient(cc grpc.ClientConnInterface) DSRHubServiceClient {
	return &dSRHubServiceClient{cc}
}

func (c *dSRHubServiceClient) CreateDSR(ctx context.Context, in *CreateDSRRequest, opts ...grpc.CallOption) (*CreateDSRResponse, error) {
	out := new(CreateDSRResponse)
	err := c.cc.Invoke(ctx, "/idl.dsrhub.DSRHubService/CreateDSR", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dSRHubServiceClient) ExchangeIdentity(ctx context.Context, in *ExchangeIdentityRequest, opts ...grpc.CallOption) (*ExchangeIdentityResponse, error) {
	out := new(ExchangeIdentityResponse)
	err := c.cc.Invoke(ctx, "/idl.dsrhub.DSRHubService/ExchangeIdentity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DSRHubServiceServer is the server API for DSRHubService service.
type DSRHubServiceServer interface {
	CreateDSR(context.Context, *CreateDSRRequest) (*CreateDSRResponse, error)
	ExchangeIdentity(context.Context, *ExchangeIdentityRequest) (*ExchangeIdentityResponse, error)
}

// UnimplementedDSRHubServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDSRHubServiceServer struct {
}

func (*UnimplementedDSRHubServiceServer) CreateDSR(ctx context.Context, req *CreateDSRRequest) (*CreateDSRResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDSR not implemented")
}
func (*UnimplementedDSRHubServiceServer) ExchangeIdentity(ctx context.Context, req *ExchangeIdentityRequest) (*ExchangeIdentityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExchangeIdentity not implemented")
}

func RegisterDSRHubServiceServer(s *grpc.Server, srv DSRHubServiceServer) {
	s.RegisterService(&_DSRHubService_serviceDesc, srv)
}

func _DSRHubService_CreateDSR_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDSRRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DSRHubServiceServer).CreateDSR(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/idl.dsrhub.DSRHubService/CreateDSR",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DSRHubServiceServer).CreateDSR(ctx, req.(*CreateDSRRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DSRHubService_ExchangeIdentity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExchangeIdentityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DSRHubServiceServer).ExchangeIdentity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/idl.dsrhub.DSRHubService/ExchangeIdentity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DSRHubServiceServer).ExchangeIdentity(ctx, req.(*ExchangeIdentityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DSRHubService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "idl.dsrhub.DSRHubService",
	HandlerType: (*DSRHubServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDSR",
			Handler:    _DSRHubService_CreateDSR_Handler,
		},
		{
			MethodName: "ExchangeIdentity",
			Handler:    _DSRHubService_ExchangeIdentity_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dsrhub.proto",
}
