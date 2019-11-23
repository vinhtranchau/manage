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
	FirstName            string      `protobuf:"bytes,1,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	MidName              string      `protobuf:"bytes,2,opt,name=MidName,proto3" json:"MidName,omitempty"`
	LastName             string      `protobuf:"bytes,3,opt,name=LastName,proto3" json:"LastName,omitempty"`
	PhoneNumber          string      `protobuf:"bytes,4,opt,name=PhoneNumber,proto3" json:"PhoneNumber,omitempty"`
	Email                string      `protobuf:"bytes,5,opt,name=Email,proto3" json:"Email,omitempty"`
	Password             string      `protobuf:"bytes,6,opt,name=Password,proto3" json:"Password,omitempty"`
	Role                 string      `protobuf:"bytes,7,opt,name=Role,proto3" json:"Role,omitempty"`
	ID                   string      `protobuf:"bytes,8,opt,name=ID,proto3" json:"ID,omitempty"`
	State                string      `protobuf:"bytes,9,opt,name=State,proto3" json:"State,omitempty"`
	City                 string      `protobuf:"bytes,10,opt,name=City,proto3" json:"City,omitempty"`
	Address              string      `protobuf:"bytes,11,opt,name=Address,proto3" json:"Address,omitempty"`
	CNPJ                 string      `protobuf:"bytes,12,opt,name=CNPJ,proto3" json:"CNPJ,omitempty"`
	CompanyName          string      `protobuf:"bytes,13,opt,name=CompanyName,proto3" json:"CompanyName,omitempty"`
	Category             string      `protobuf:"bytes,14,opt,name=Category,proto3" json:"Category,omitempty"`
	Zip                  string      `protobuf:"bytes,15,opt,name=Zip,proto3" json:"Zip,omitempty"`
	Username             string      `protobuf:"bytes,16,opt,name=Username,proto3" json:"Username,omitempty"`
	Companyid            string      `protobuf:"bytes,17,opt,name=Companyid,proto3" json:"Companyid,omitempty"`
	Permission           *Permission `protobuf:"bytes,18,opt,name=permission,proto3" json:"permission,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
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

func (m *User) GetPermission() *Permission {
	if m != nil {
		return m.Permission
	}
	return nil
}

type PermissionParams struct {
	Permission           *Permission `protobuf:"bytes,1,opt,name=permission,proto3" json:"permission,omitempty"`
	Id                   string      `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *PermissionParams) Reset()         { *m = PermissionParams{} }
func (m *PermissionParams) String() string { return proto.CompactTextString(m) }
func (*PermissionParams) ProtoMessage()    {}
func (*PermissionParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_a774a375f779fe02, []int{1}
}

func (m *PermissionParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PermissionParams.Unmarshal(m, b)
}
func (m *PermissionParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PermissionParams.Marshal(b, m, deterministic)
}
func (m *PermissionParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PermissionParams.Merge(m, src)
}
func (m *PermissionParams) XXX_Size() int {
	return xxx_messageInfo_PermissionParams.Size(m)
}
func (m *PermissionParams) XXX_DiscardUnknown() {
	xxx_messageInfo_PermissionParams.DiscardUnknown(m)
}

var xxx_messageInfo_PermissionParams proto.InternalMessageInfo

func (m *PermissionParams) GetPermission() *Permission {
	if m != nil {
		return m.Permission
	}
	return nil
}

func (m *PermissionParams) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Permission struct {
	Financial            bool     `protobuf:"varint,1,opt,name=financial,proto3" json:"financial,omitempty"`
	Quotes               bool     `protobuf:"varint,2,opt,name=quotes,proto3" json:"quotes,omitempty"`
	Stock                bool     `protobuf:"varint,3,opt,name=stock,proto3" json:"stock,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Permission) Reset()         { *m = Permission{} }
func (m *Permission) String() string { return proto.CompactTextString(m) }
func (*Permission) ProtoMessage()    {}
func (*Permission) Descriptor() ([]byte, []int) {
	return fileDescriptor_a774a375f779fe02, []int{2}
}

func (m *Permission) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Permission.Unmarshal(m, b)
}
func (m *Permission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Permission.Marshal(b, m, deterministic)
}
func (m *Permission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Permission.Merge(m, src)
}
func (m *Permission) XXX_Size() int {
	return xxx_messageInfo_Permission.Size(m)
}
func (m *Permission) XXX_DiscardUnknown() {
	xxx_messageInfo_Permission.DiscardUnknown(m)
}

var xxx_messageInfo_Permission proto.InternalMessageInfo

func (m *Permission) GetFinancial() bool {
	if m != nil {
		return m.Financial
	}
	return false
}

func (m *Permission) GetQuotes() bool {
	if m != nil {
		return m.Quotes
	}
	return false
}

func (m *Permission) GetStock() bool {
	if m != nil {
		return m.Stock
	}
	return false
}

type PasswordChange struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Old                  string   `protobuf:"bytes,2,opt,name=old,proto3" json:"old,omitempty"`
	New                  string   `protobuf:"bytes,3,opt,name=new,proto3" json:"new,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PasswordChange) Reset()         { *m = PasswordChange{} }
func (m *PasswordChange) String() string { return proto.CompactTextString(m) }
func (*PasswordChange) ProtoMessage()    {}
func (*PasswordChange) Descriptor() ([]byte, []int) {
	return fileDescriptor_a774a375f779fe02, []int{3}
}

func (m *PasswordChange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PasswordChange.Unmarshal(m, b)
}
func (m *PasswordChange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PasswordChange.Marshal(b, m, deterministic)
}
func (m *PasswordChange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PasswordChange.Merge(m, src)
}
func (m *PasswordChange) XXX_Size() int {
	return xxx_messageInfo_PasswordChange.Size(m)
}
func (m *PasswordChange) XXX_DiscardUnknown() {
	xxx_messageInfo_PasswordChange.DiscardUnknown(m)
}

var xxx_messageInfo_PasswordChange proto.InternalMessageInfo

func (m *PasswordChange) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PasswordChange) GetOld() string {
	if m != nil {
		return m.Old
	}
	return ""
}

func (m *PasswordChange) GetNew() string {
	if m != nil {
		return m.New
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
	return fileDescriptor_a774a375f779fe02, []int{4}
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
	return fileDescriptor_a774a375f779fe02, []int{5}
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
	return fileDescriptor_a774a375f779fe02, []int{6}
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
	return fileDescriptor_a774a375f779fe02, []int{7}
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
	proto.RegisterType((*PermissionParams)(nil), "authorization.PermissionParams")
	proto.RegisterType((*Permission)(nil), "authorization.Permission")
	proto.RegisterType((*PasswordChange)(nil), "authorization.PasswordChange")
	proto.RegisterType((*LoginParams)(nil), "authorization.LoginParams")
	proto.RegisterType((*Response)(nil), "authorization.Response")
	proto.RegisterType((*Params)(nil), "authorization.Params")
	proto.RegisterType((*Stats)(nil), "authorization.Stats")
}

func init() { proto.RegisterFile("authorization/authorization.proto", fileDescriptor_a774a375f779fe02) }

var fileDescriptor_a774a375f779fe02 = []byte{
	// 679 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xeb, 0x6a, 0xdb, 0x4a,
	0x10, 0xc6, 0xd7, 0xc8, 0xe3, 0x63, 0x27, 0xd9, 0x93, 0x73, 0xb2, 0xc7, 0xe4, 0xd0, 0x54, 0x50,
	0x48, 0xff, 0x24, 0xe0, 0x52, 0x7a, 0xa1, 0x10, 0x82, 0xd3, 0x4b, 0xd2, 0xc4, 0xb8, 0x4a, 0x0b,
	0xa5, 0xd0, 0x1f, 0x1b, 0x6b, 0xeb, 0x2c, 0xb1, 0xb4, 0xca, 0xee, 0xba, 0xc1, 0x7d, 0xcd, 0x3e,
	0x48, 0x5f, 0xa1, 0xcc, 0xae, 0x24, 0xdb, 0x6a, 0x44, 0xe9, 0xbf, 0xf9, 0xbe, 0xb9, 0xec, 0x7c,
	0x33, 0x23, 0x1b, 0xee, 0xb3, 0x99, 0xb9, 0x92, 0x4a, 0x7c, 0x63, 0x46, 0xc8, 0xf8, 0x60, 0x05,
	0xed, 0x27, 0x4a, 0x1a, 0x49, 0x3a, 0x2b, 0x64, 0x6f, 0x3b, 0x51, 0x32, 0x9c, 0x8d, 0x8d, 0x3e,
	0xc8, 0x0c, 0x17, 0xe7, 0xff, 0xa8, 0x41, 0xfd, 0x83, 0xe6, 0x8a, 0xec, 0x40, 0xeb, 0x95, 0x50,
	0xda, 0x0c, 0x59, 0xc4, 0x69, 0x65, 0xb7, 0xb2, 0xd7, 0x0a, 0x16, 0x04, 0xa1, 0xb0, 0x76, 0x2e,
	0x42, 0xeb, 0xab, 0x5a, 0x5f, 0x06, 0x49, 0x0f, 0xbc, 0x33, 0x96, 0xa6, 0xd5, 0xac, 0x2b, 0xc7,
	0x64, 0x17, 0xda, 0xa3, 0x2b, 0x19, 0xf3, 0xe1, 0x2c, 0xba, 0xe4, 0x8a, 0xd6, 0xad, 0x7b, 0x99,
	0x22, 0x5b, 0xd0, 0x78, 0x19, 0x31, 0x31, 0xa5, 0x0d, 0xeb, 0x73, 0x00, 0x6b, 0x8e, 0x98, 0xd6,
	0xb7, 0x52, 0x85, 0xb4, 0xe9, 0x6a, 0x66, 0x98, 0x10, 0xa8, 0x07, 0x72, 0xca, 0xe9, 0x9a, 0xe5,
	0xad, 0x4d, 0xba, 0x50, 0x3d, 0x39, 0xa6, 0x9e, 0x65, 0xaa, 0x27, 0xc7, 0x58, 0xf5, 0xc2, 0x30,
	0xc3, 0x69, 0xcb, 0x55, 0xb5, 0x00, 0x33, 0x07, 0xc2, 0xcc, 0x29, 0xb8, 0x4c, 0xb4, 0x51, 0xd7,
	0x51, 0x18, 0x2a, 0xae, 0x35, 0x6d, 0x3b, 0x5d, 0x29, 0xb4, 0xd1, 0xc3, 0xd1, 0x29, 0xfd, 0x2b,
	0x8d, 0x1e, 0x8e, 0x4e, 0x51, 0xcf, 0x40, 0x46, 0x09, 0x8b, 0xe7, 0x56, 0x6e, 0xc7, 0xe9, 0x59,
	0xa2, 0xb0, 0xf3, 0x01, 0x33, 0x7c, 0x22, 0xd5, 0x9c, 0x76, 0x5d, 0xe7, 0x19, 0x26, 0x1b, 0x50,
	0xfb, 0x24, 0x12, 0xba, 0x6e, 0x69, 0x34, 0x31, 0x1a, 0x67, 0x1f, 0x63, 0xb1, 0x0d, 0x17, 0x9d,
	0x61, 0xdc, 0x47, 0x5a, 0x58, 0x84, 0x74, 0xd3, 0xed, 0x23, 0x27, 0xc8, 0x33, 0x80, 0x84, 0xab,
	0x48, 0x68, 0x2d, 0x64, 0x4c, 0xc9, 0x6e, 0x65, 0xaf, 0xdd, 0xff, 0x6f, 0x7f, 0xf5, 0x10, 0x46,
	0x79, 0x40, 0xb0, 0x14, 0xec, 0x7f, 0x86, 0x8d, 0x85, 0x67, 0xc4, 0x14, 0x8b, 0x74, 0xa1, 0x5c,
	0xe5, 0x0f, 0xca, 0xe1, 0xec, 0x45, 0x98, 0x1e, 0x45, 0x55, 0x84, 0xfe, 0x47, 0x80, 0x45, 0x24,
	0xaa, 0xf8, 0x22, 0x62, 0x16, 0x8f, 0x05, 0x9b, 0xda, 0xba, 0x5e, 0xb0, 0x20, 0xc8, 0xbf, 0xd0,
	0xbc, 0x99, 0x49, 0xc3, 0xb5, 0xcd, 0xf7, 0x82, 0x14, 0xe1, 0xfe, 0xb4, 0x91, 0xe3, 0x6b, 0x7b,
	0x50, 0x5e, 0xe0, 0x80, 0x7f, 0x0c, 0xdd, 0xec, 0x0a, 0x06, 0x57, 0x2c, 0x9e, 0xf0, 0xf4, 0xed,
	0x4a, 0xf6, 0x36, 0x4e, 0x58, 0x4e, 0xb3, 0x66, 0xd0, 0x44, 0x26, 0xe6, 0xb7, 0xe9, 0x61, 0xa2,
	0xe9, 0x1f, 0x42, 0xfb, 0x4c, 0x4e, 0x44, 0xa6, 0x3c, 0x3f, 0xc0, 0x4a, 0xd9, 0x01, 0x56, 0x57,
	0x0f, 0xd0, 0x7f, 0x0b, 0x5e, 0xc0, 0x75, 0x22, 0x63, 0xcd, 0x31, 0xfb, 0xbd, 0xbc, 0xe6, 0x71,
	0x96, 0x6d, 0x01, 0x79, 0x08, 0x8d, 0x99, 0xe6, 0x0a, 0x55, 0xd5, 0xf6, 0xda, 0xfd, 0xbf, 0x0b,
	0x83, 0xc4, 0x15, 0x07, 0x2e, 0xc2, 0x3f, 0x84, 0x66, 0xda, 0x48, 0x51, 0xcb, 0x03, 0x68, 0xdc,
	0xcc, 0xb8, 0x9a, 0xdb, 0xf7, 0xdb, 0xfd, 0xf5, 0xfd, 0xfc, 0xc3, 0x7d, 0x87, 0x74, 0xe0, 0xbc,
	0xfe, 0xb5, 0x3b, 0x75, 0xbd, 0x34, 0x4b, 0xac, 0xd1, 0xc9, 0x67, 0xb9, 0x03, 0xad, 0xb1, 0x8c,
	0x92, 0x29, 0x37, 0xdc, 0x69, 0xe9, 0x04, 0x0b, 0x02, 0x85, 0x6a, 0xae, 0xbe, 0x8a, 0x31, 0xd7,
	0x76, 0x48, 0x9d, 0x20, 0xc7, 0xd9, 0xec, 0xea, 0x96, 0x46, 0xb3, 0xff, 0xbd, 0x06, 0x5b, 0x47,
	0xcb, 0x5a, 0x2e, 0x5c, 0x2c, 0x79, 0x8a, 0x33, 0x99, 0x08, 0x6d, 0xb8, 0x22, 0x77, 0xc9, 0xed,
	0x6d, 0x17, 0xc8, 0x7c, 0x82, 0x2f, 0xa0, 0x61, 0xd7, 0x41, 0x7a, 0x85, 0x88, 0xa5, 0x25, 0x95,
	0x67, 0xbf, 0x81, 0xae, 0x3b, 0x85, 0xfc, 0xe7, 0xe1, 0xff, 0xe2, 0xd5, 0xae, 0x5c, 0x4c, 0x79,
	0xa5, 0xc7, 0xb0, 0xf6, 0x9a, 0x1b, 0xfb, 0x4b, 0xf8, 0xcf, 0x2f, 0x25, 0x6c, 0x13, 0x77, 0xe9,
	0x22, 0xcf, 0xc1, 0x4b, 0xd3, 0x74, 0x59, 0x5e, 0xe9, 0x93, 0x4f, 0x6c, 0xae, 0xdb, 0x5e, 0x49,
	0xee, 0x56, 0x81, 0x76, 0xc1, 0xe7, 0xb0, 0x99, 0xaa, 0xce, 0x3f, 0x34, 0x4d, 0xee, 0x95, 0x7e,
	0xae, 0xbf, 0xe9, 0xe3, 0xb2, 0x69, 0xff, 0x09, 0x1e, 0xfd, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x1b,
	0x1d, 0x86, 0xe7, 0x56, 0x06, 0x00, 0x00,
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
	ChangePassword(ctx context.Context, in *PasswordChange, opts ...grpc.CallOption) (*Response, error)
	GetUser(ctx context.Context, in *Params, opts ...grpc.CallOption) (*User, error)
	GetUsers(ctx context.Context, in *Params, opts ...grpc.CallOption) (*Response, error)
	GetStats(ctx context.Context, in *Params, opts ...grpc.CallOption) (*Stats, error)
	ChangePermissions(ctx context.Context, in *PermissionParams, opts ...grpc.CallOption) (*Response, error)
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

func (c *authorizationServiceClient) ChangePassword(ctx context.Context, in *PasswordChange, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/authorization.AuthorizationService/ChangePassword", in, out, opts...)
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

func (c *authorizationServiceClient) ChangePermissions(ctx context.Context, in *PermissionParams, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/authorization.AuthorizationService/ChangePermissions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorizationServiceServer is the server API for AuthorizationService service.
type AuthorizationServiceServer interface {
	Register(context.Context, *User) (*Response, error)
	Login(context.Context, *LoginParams) (*Response, error)
	ChangePassword(context.Context, *PasswordChange) (*Response, error)
	GetUser(context.Context, *Params) (*User, error)
	GetUsers(context.Context, *Params) (*Response, error)
	GetStats(context.Context, *Params) (*Stats, error)
	ChangePermissions(context.Context, *PermissionParams) (*Response, error)
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
func (*UnimplementedAuthorizationServiceServer) ChangePassword(ctx context.Context, req *PasswordChange) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePassword not implemented")
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
func (*UnimplementedAuthorizationServiceServer) ChangePermissions(ctx context.Context, req *PermissionParams) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePermissions not implemented")
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

func _AuthorizationService_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PasswordChange)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServiceServer).ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authorization.AuthorizationService/ChangePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServiceServer).ChangePassword(ctx, req.(*PasswordChange))
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

func _AuthorizationService_ChangePermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PermissionParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServiceServer).ChangePermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authorization.AuthorizationService/ChangePermissions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServiceServer).ChangePermissions(ctx, req.(*PermissionParams))
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
			MethodName: "ChangePassword",
			Handler:    _AuthorizationService_ChangePassword_Handler,
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
		{
			MethodName: "ChangePermissions",
			Handler:    _AuthorizationService_ChangePermissions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authorization/authorization.proto",
}
