// Code generated by protoc-gen-go. DO NOT EDIT.
// source: authorization/authorization.proto

package authorization

import (
	context "context"
	fmt "fmt"
	math "math"

	"github.com/Beaxhem/manage/backend/internal/products"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type User struct {
	FirstName            string   `protobuf:"bytes,1,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	MidName              string   `protobuf:"bytes,2,opt,name=MidName,proto3" json:"MidName,omitempty"`
	LastName             string   `protobuf:"bytes,3,opt,name=LastName,proto3" json:"LastName,omitempty"`
	PhoneNumber          string   `protobuf:"bytes,4,opt,name=PhoneNumber,proto3" json:"PhoneNumber,omitempty"`
	Email                string   `protobuf:"bytes,5,opt,name=Email,proto3" json:"Email,omitempty"`
	Password             string   `protobuf:"bytes,6,opt,name=Password,proto3" json:"Password,omitempty"`
	Role                 string   `protobuf:"bytes,7,opt,name=Role,proto3" json:"Role,omitempty"`
	ID                   string   `protobuf:"bytes,8,opt,name=ID,proto3" json:"ID,omitempty"`
	State                string   `protobuf:"bytes,9,opt,name=State,proto3" json:"State,omitempty"`
	City                 string   `protobuf:"bytes,10,opt,name=City,proto3" json:"City,omitempty"`
	Address              string   `protobuf:"bytes,11,opt,name=Address,proto3" json:"Address,omitempty"`
	CNPJ                 string   `protobuf:"bytes,12,opt,name=CNPJ,proto3" json:"CNPJ,omitempty"`
	CompanyName          string   `protobuf:"bytes,13,opt,name=CompanyName,proto3" json:"CompanyName,omitempty"`
	Category             string   `protobuf:"bytes,14,opt,name=Category,proto3" json:"Category,omitempty"`
	Zip                  string   `protobuf:"bytes,15,opt,name=Zip,proto3" json:"Zip,omitempty"`
	Username             string   `protobuf:"bytes,16,opt,name=Username,proto3" json:"Username,omitempty"`
	Companyid            string   `protobuf:"bytes,17,opt,name=Companyid,proto3" json:"Companyid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_a774a375f779fe02, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *User) GetMidName() string {
	if m != nil {
		return m.MidName
	}
	return ""
}

func (m *User) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *User) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *User) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *User) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *User) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *User) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *User) GetCNPJ() string {
	if m != nil {
		return m.CNPJ
	}
	return ""
}

func (m *User) GetCompanyName() string {
	if m != nil {
		return m.CompanyName
	}
	return ""
}

func (m *User) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *User) GetZip() string {
	if m != nil {
		return m.Zip
	}
	return ""
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetCompanyid() string {
	if m != nil {
		return m.Companyid
	}
	return ""
}

type LoginParams struct {
	Email                string   `protobuf:"bytes,1,opt,name=Email,proto3" json:"Email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginParams) Reset()         { *m = LoginParams{} }
func (m *LoginParams) String() string { return proto.CompactTextString(m) }
func (*LoginParams) ProtoMessage()    {}
func (*LoginParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_a774a375f779fe02, []int{1}
}

func (m *LoginParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginParams.Unmarshal(m, b)
}
func (m *LoginParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginParams.Marshal(b, m, deterministic)
}
func (m *LoginParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginParams.Merge(m, src)
}
func (m *LoginParams) XXX_Size() int {
	return xxx_messageInfo_LoginParams.Size(m)
}
func (m *LoginParams) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginParams.DiscardUnknown(m)
}

var xxx_messageInfo_LoginParams proto.InternalMessageInfo

func (m *LoginParams) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *LoginParams) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type Response struct {
	Token                string   `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	Users                []*User  `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_a774a375f779fe02, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Response) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

type Params struct {
	Id                   string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Query                *products.Query `protobuf:"bytes,2,opt,name=query,proto3" json:"query,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_a774a375f779fe02, []int{3}
}

func (m *Params) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Params.Unmarshal(m, b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Params.Marshal(b, m, deterministic)
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return xxx_messageInfo_Params.Size(m)
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Params) GetQuery() *products.Query {
	if m != nil {
		return m.Query
	}
	return nil
}

type Stats struct {
	Quotes               uint32   `protobuf:"varint,1,opt,name=quotes,proto3" json:"quotes,omitempty"`
	Completed            uint32   `protobuf:"varint,2,opt,name=completed,proto3" json:"completed,omitempty"`
	Services             uint32   `protobuf:"varint,3,opt,name=services,proto3" json:"services,omitempty"`
	New                  uint32   `protobuf:"varint,4,opt,name=new,proto3" json:"new,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Stats) Reset()         { *m = Stats{} }
func (m *Stats) String() string { return proto.CompactTextString(m) }
func (*Stats) ProtoMessage()    {}
func (*Stats) Descriptor() ([]byte, []int) {
	return fileDescriptor_a774a375f779fe02, []int{4}
}

func (m *Stats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Stats.Unmarshal(m, b)
}
func (m *Stats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Stats.Marshal(b, m, deterministic)
}
func (m *Stats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Stats.Merge(m, src)
}
func (m *Stats) XXX_Size() int {
	return xxx_messageInfo_Stats.Size(m)
}
func (m *Stats) XXX_DiscardUnknown() {
	xxx_messageInfo_Stats.DiscardUnknown(m)
}

var xxx_messageInfo_Stats proto.InternalMessageInfo

func (m *Stats) GetQuotes() uint32 {
	if m != nil {
		return m.Quotes
	}
	return 0
}

func (m *Stats) GetCompleted() uint32 {
	if m != nil {
		return m.Completed
	}
	return 0
}

func (m *Stats) GetServices() uint32 {
	if m != nil {
		return m.Services
	}
	return 0
}

func (m *Stats) GetNew() uint32 {
	if m != nil {
		return m.New
	}
	return 0
}

func init() {
	proto.RegisterType((*User)(nil), "authorization.User")
	proto.RegisterType((*LoginParams)(nil), "authorization.LoginParams")
	proto.RegisterType((*Response)(nil), "authorization.Response")
	proto.RegisterType((*Params)(nil), "authorization.Params")
	proto.RegisterType((*Stats)(nil), "authorization.Stats")
}

func init() { proto.RegisterFile("authorization/authorization.proto", fileDescriptor_a774a375f779fe02) }

var fileDescriptor_a774a375f779fe02 = []byte{
	// 542 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x54, 0x5b, 0x6f, 0xd3, 0x4c,
	0x10, 0x55, 0x9d, 0xe6, 0x36, 0xf9, 0x9c, 0xf6, 0x5b, 0x02, 0x5d, 0x45, 0x3c, 0x84, 0x48, 0x48,
	0xe5, 0x25, 0x95, 0x82, 0x10, 0x08, 0x21, 0x55, 0x55, 0x0a, 0xa8, 0xa5, 0x44, 0xc1, 0x85, 0x17,
	0xde, 0xb6, 0xf1, 0x28, 0x5d, 0x35, 0xf6, 0xba, 0xbb, 0x6b, 0xaa, 0xf0, 0xbb, 0x78, 0xe2, 0xd7,
	0xa1, 0x9d, 0xb5, 0x73, 0x53, 0xf3, 0x36, 0xe7, 0xcc, 0x39, 0xe3, 0xb9, 0x6c, 0x02, 0x2f, 0x44,
	0x6e, 0x6f, 0x95, 0x96, 0xbf, 0x85, 0x95, 0x2a, 0x3d, 0xd9, 0x40, 0x83, 0x4c, 0x2b, 0xab, 0x58,
	0xb8, 0x41, 0x76, 0x8f, 0x32, 0xad, 0xe2, 0x7c, 0x6a, 0xcd, 0x49, 0x19, 0x78, 0x5d, 0xff, 0x4f,
	0x05, 0xf6, 0x7f, 0x18, 0xd4, 0xec, 0x39, 0x34, 0x3f, 0x49, 0x6d, 0xec, 0x58, 0x24, 0xc8, 0xf7,
	0x7a, 0x7b, 0xc7, 0xcd, 0x68, 0x45, 0x30, 0x0e, 0xf5, 0xaf, 0x32, 0xa6, 0x5c, 0x40, 0xb9, 0x12,
	0xb2, 0x2e, 0x34, 0xae, 0x44, 0x61, 0xab, 0x50, 0x6a, 0x89, 0x59, 0x0f, 0x5a, 0x93, 0x5b, 0x95,
	0xe2, 0x38, 0x4f, 0x6e, 0x50, 0xf3, 0x7d, 0x4a, 0xaf, 0x53, 0xac, 0x03, 0xd5, 0x8f, 0x89, 0x90,
	0x73, 0x5e, 0xa5, 0x9c, 0x07, 0xae, 0xe6, 0x44, 0x18, 0xf3, 0xa0, 0x74, 0xcc, 0x6b, 0xbe, 0x66,
	0x89, 0x19, 0x83, 0xfd, 0x48, 0xcd, 0x91, 0xd7, 0x89, 0xa7, 0x98, 0xb5, 0x21, 0xb8, 0x38, 0xe7,
	0x0d, 0x62, 0x82, 0x8b, 0x73, 0x57, 0xf5, 0xda, 0x0a, 0x8b, 0xbc, 0xe9, 0xab, 0x12, 0x70, 0xce,
	0x91, 0xb4, 0x0b, 0x0e, 0xde, 0xe9, 0x62, 0x37, 0xd7, 0x59, 0x1c, 0x6b, 0x34, 0x86, 0xb7, 0xfc,
	0x5c, 0x05, 0x24, 0xf5, 0x78, 0x72, 0xc9, 0xff, 0x2b, 0xd4, 0xe3, 0xc9, 0xa5, 0x9b, 0x67, 0xa4,
	0x92, 0x4c, 0xa4, 0x0b, 0x1a, 0x37, 0xf4, 0xf3, 0xac, 0x51, 0xae, 0xf3, 0x91, 0xb0, 0x38, 0x53,
	0x7a, 0xc1, 0xdb, 0xbe, 0xf3, 0x12, 0xb3, 0x43, 0xa8, 0xfc, 0x94, 0x19, 0x3f, 0x20, 0xda, 0x85,
	0x4e, 0xed, 0x76, 0x9f, 0xba, 0x62, 0x87, 0x5e, 0x5d, 0x62, 0x77, 0x8f, 0xa2, 0xb0, 0x8c, 0xf9,
	0xff, 0xfe, 0x1e, 0x4b, 0xa2, 0x7f, 0x0a, 0xad, 0x2b, 0x35, 0x93, 0xe9, 0x44, 0x68, 0x91, 0x98,
	0xd5, 0x1a, 0xf7, 0x76, 0xad, 0x31, 0xd8, 0x5c, 0x63, 0xff, 0x0b, 0x34, 0x22, 0x34, 0x99, 0x4a,
	0x0d, 0x3a, 0xf7, 0x77, 0x75, 0x87, 0x69, 0xe9, 0x26, 0xc0, 0x5e, 0x41, 0x35, 0x37, 0xa8, 0x0d,
	0x0f, 0x7a, 0x95, 0xe3, 0xd6, 0xf0, 0xc9, 0x60, 0xf3, 0x99, 0xb9, 0x46, 0x23, 0xaf, 0xe8, 0x9f,
	0x42, 0xad, 0x68, 0xa4, 0x0d, 0x81, 0x8c, 0x8b, 0x3a, 0x81, 0x8c, 0xd9, 0x4b, 0xa8, 0xde, 0xe7,
	0xa8, 0x17, 0xf4, 0xfd, 0xd6, 0xf0, 0x60, 0xb0, 0x7c, 0x7e, 0xdf, 0x1c, 0x1d, 0xf9, 0x6c, 0xff,
	0xce, 0x1f, 0xcc, 0xb0, 0x67, 0x50, 0xbb, 0xcf, 0x95, 0x45, 0x43, 0x35, 0xc2, 0xa8, 0x40, 0x6e,
	0x1b, 0x53, 0x95, 0x64, 0x73, 0xb4, 0xe8, 0x67, 0x09, 0xa3, 0x15, 0xe1, 0x06, 0x35, 0xa8, 0x7f,
	0xc9, 0x29, 0x1a, 0x7a, 0x83, 0x61, 0xb4, 0xc4, 0x6e, 0xeb, 0x29, 0x3e, 0xd0, 0xdb, 0x0b, 0x23,
	0x17, 0x0e, 0xff, 0x06, 0xd0, 0x39, 0x5b, 0x9f, 0xe5, 0xda, 0x6b, 0xd9, 0x3b, 0xb7, 0x93, 0x99,
	0x34, 0x16, 0x35, 0x7b, 0x6c, 0xdc, 0xee, 0xd1, 0x16, 0xb9, 0xdc, 0xe0, 0x07, 0xa8, 0xd2, 0x39,
	0x58, 0x77, 0x4b, 0xb1, 0x76, 0xa4, 0xdd, 0xee, 0x37, 0x50, 0xff, 0x8c, 0x96, 0x7e, 0x85, 0x4f,
	0xb7, 0x34, 0x85, 0xf5, 0xb1, 0x6e, 0xd8, 0x7b, 0x68, 0x14, 0x36, 0xb3, 0xcb, 0xb7, 0xf3, 0x93,
	0x6f, 0xc9, 0xeb, 0x77, 0xbe, 0xc3, 0xdb, 0xd9, 0xa2, 0x49, 0x7c, 0x53, 0xa3, 0xbf, 0x8d, 0xd7,
	0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xe8, 0x90, 0xf3, 0x31, 0x83, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthorizationServiceClient is the client API for AuthorizationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthorizationServiceClient interface {
	Register(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response, error)
	Login(ctx context.Context, in *LoginParams, opts ...grpc.CallOption) (*Response, error)
	GetUser(ctx context.Context, in *Params, opts ...grpc.CallOption) (*User, error)
	GetUsers(ctx context.Context, in *Params, opts ...grpc.CallOption) (*Response, error)
	GetStats(ctx context.Context, in *Params, opts ...grpc.CallOption) (*Stats, error)
}

type authorizationServiceClient struct {
	cc *grpc.ClientConn
}

func NewAuthorizationServiceClient(cc *grpc.ClientConn) AuthorizationServiceClient {
	return &authorizationServiceClient{cc}
}

func (c *authorizationServiceClient) Register(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/authorization.AuthorizationService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationServiceClient) Login(ctx context.Context, in *LoginParams, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/authorization.AuthorizationService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationServiceClient) GetUser(ctx context.Context, in *Params, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/authorization.AuthorizationService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationServiceClient) GetUsers(ctx context.Context, in *Params, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/authorization.AuthorizationService/GetUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationServiceClient) GetStats(ctx context.Context, in *Params, opts ...grpc.CallOption) (*Stats, error) {
	out := new(Stats)
	err := c.cc.Invoke(ctx, "/authorization.AuthorizationService/GetStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorizationServiceServer is the server API for AuthorizationService service.
type AuthorizationServiceServer interface {
	Register(context.Context, *User) (*Response, error)
	Login(context.Context, *LoginParams) (*Response, error)
	GetUser(context.Context, *Params) (*User, error)
	GetUsers(context.Context, *Params) (*Response, error)
	GetStats(context.Context, *Params) (*Stats, error)
}

// UnimplementedAuthorizationServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAuthorizationServiceServer struct {
}

func (*UnimplementedAuthorizationServiceServer) Register(ctx context.Context, req *User) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (*UnimplementedAuthorizationServiceServer) Login(ctx context.Context, req *LoginParams) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedAuthorizationServiceServer) GetUser(ctx context.Context, req *Params) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (*UnimplementedAuthorizationServiceServer) GetUsers(ctx context.Context, req *Params) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
func (*UnimplementedAuthorizationServiceServer) GetStats(ctx context.Context, req *Params) (*Stats, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStats not implemented")
}

func RegisterAuthorizationServiceServer(s *grpc.Server, srv AuthorizationServiceServer) {
	s.RegisterService(&_AuthorizationService_serviceDesc, srv)
}

func _AuthorizationService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authorization.AuthorizationService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServiceServer).Register(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorizationService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authorization.AuthorizationService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServiceServer).Login(ctx, req.(*LoginParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorizationService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Params)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authorization.AuthorizationService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServiceServer).GetUser(ctx, req.(*Params))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorizationService_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Params)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServiceServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authorization.AuthorizationService/GetUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServiceServer).GetUsers(ctx, req.(*Params))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorizationService_GetStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Params)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServiceServer).GetStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authorization.AuthorizationService/GetStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServiceServer).GetStats(ctx, req.(*Params))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthorizationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "authorization.AuthorizationService",
	HandlerType: (*AuthorizationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _AuthorizationService_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _AuthorizationService_Login_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _AuthorizationService_GetUser_Handler,
		},
		{
			MethodName: "GetUsers",
			Handler:    _AuthorizationService_GetUsers_Handler,
		},
		{
			MethodName: "GetStats",
			Handler:    _AuthorizationService_GetStats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authorization/authorization.proto",
}