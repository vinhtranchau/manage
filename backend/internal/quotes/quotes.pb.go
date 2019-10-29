// Code generated by protoc-gen-go. DO NOT EDIT.
// source: quotes/quotes.proto

package quotes

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

type Status int32

const (
	Status_NEW               Status = 0
	Status_SUPPLIER_APPLIED  Status = 1
	Status_CLIENT_APPLIED    Status = 2
	Status_SUPPLIER_REJECTED Status = 3
	Status_CLIENT_REJECTED   Status = 4
)

var Status_name = map[int32]string{
	0: "NEW",
	1: "SUPPLIER_APPLIED",
	2: "CLIENT_APPLIED",
	3: "SUPPLIER_REJECTED",
	4: "CLIENT_REJECTED",
}

var Status_value = map[string]int32{
	"NEW":               0,
	"SUPPLIER_APPLIED":  1,
	"CLIENT_APPLIED":    2,
	"SUPPLIER_REJECTED": 3,
	"CLIENT_REJECTED":   4,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_15679ebe12a64146, []int{0}
}

type ShippingStatus int32

const (
	ShippingStatus_PREPARATION ShippingStatus = 0
	ShippingStatus_DELIVERED   ShippingStatus = 1
	ShippingStatus_COMPLETED   ShippingStatus = 2
	ShippingStatus_CANCELLED   ShippingStatus = 3
)

var ShippingStatus_name = map[int32]string{
	0: "PREPARATION",
	1: "DELIVERED",
	2: "COMPLETED",
	3: "CANCELLED",
}

var ShippingStatus_value = map[string]int32{
	"PREPARATION": 0,
	"DELIVERED":   1,
	"COMPLETED":   2,
	"CANCELLED":   3,
}

func (x ShippingStatus) String() string {
	return proto.EnumName(ShippingStatus_name, int32(x))
}

func (ShippingStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_15679ebe12a64146, []int{1}
}

type QuoteResponse struct {
	Quotes               []*Quote `protobuf:"bytes,1,rep,name=quotes,proto3" json:"quotes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QuoteResponse) Reset()         { *m = QuoteResponse{} }
func (m *QuoteResponse) String() string { return proto.CompactTextString(m) }
func (*QuoteResponse) ProtoMessage()    {}
func (*QuoteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_15679ebe12a64146, []int{0}
}

func (m *QuoteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuoteResponse.Unmarshal(m, b)
}
func (m *QuoteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuoteResponse.Marshal(b, m, deterministic)
}
func (m *QuoteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuoteResponse.Merge(m, src)
}
func (m *QuoteResponse) XXX_Size() int {
	return xxx_messageInfo_QuoteResponse.Size(m)
}
func (m *QuoteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QuoteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QuoteResponse proto.InternalMessageInfo

func (m *QuoteResponse) GetQuotes() []*Quote {
	if m != nil {
		return m.Quotes
	}
	return nil
}

type QuoteParams struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Jwt                  string   `protobuf:"bytes,2,opt,name=jwt,proto3" json:"jwt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QuoteParams) Reset()         { *m = QuoteParams{} }
func (m *QuoteParams) String() string { return proto.CompactTextString(m) }
func (*QuoteParams) ProtoMessage()    {}
func (*QuoteParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_15679ebe12a64146, []int{1}
}

func (m *QuoteParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuoteParams.Unmarshal(m, b)
}
func (m *QuoteParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuoteParams.Marshal(b, m, deterministic)
}
func (m *QuoteParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuoteParams.Merge(m, src)
}
func (m *QuoteParams) XXX_Size() int {
	return xxx_messageInfo_QuoteParams.Size(m)
}
func (m *QuoteParams) XXX_DiscardUnknown() {
	xxx_messageInfo_QuoteParams.DiscardUnknown(m)
}

var xxx_messageInfo_QuoteParams proto.InternalMessageInfo

func (m *QuoteParams) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *QuoteParams) GetJwt() string {
	if m != nil {
		return m.Jwt
	}
	return ""
}

type Response struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_15679ebe12a64146, []int{2}
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

func (m *Response) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Stats struct {
	Active               uint32   `protobuf:"varint,1,opt,name=active,proto3" json:"active,omitempty"`
	Pending              uint32   `protobuf:"varint,2,opt,name=pending,proto3" json:"pending,omitempty"`
	Denied               uint32   `protobuf:"varint,3,opt,name=denied,proto3" json:"denied,omitempty"`
	New                  uint32   `protobuf:"varint,4,opt,name=new,proto3" json:"new,omitempty"`
	Preparation          uint32   `protobuf:"varint,5,opt,name=preparation,proto3" json:"preparation,omitempty"`
	Production           uint32   `protobuf:"varint,6,opt,name=production,proto3" json:"production,omitempty"`
	Quality              uint32   `protobuf:"varint,7,opt,name=quality,proto3" json:"quality,omitempty"`
	Ready                uint32   `protobuf:"varint,8,opt,name=ready,proto3" json:"ready,omitempty"`
	Pickedup             uint32   `protobuf:"varint,9,opt,name=pickedup,proto3" json:"pickedup,omitempty"`
	Delivered            uint32   `protobuf:"varint,10,opt,name=delivered,proto3" json:"delivered,omitempty"`
	Installed            uint32   `protobuf:"varint,11,opt,name=installed,proto3" json:"installed,omitempty"`
	Completed            uint32   `protobuf:"varint,12,opt,name=completed,proto3" json:"completed,omitempty"`
	Cancelled            uint32   `protobuf:"varint,13,opt,name=cancelled,proto3" json:"cancelled,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Stats) Reset()         { *m = Stats{} }
func (m *Stats) String() string { return proto.CompactTextString(m) }
func (*Stats) ProtoMessage()    {}
func (*Stats) Descriptor() ([]byte, []int) {
	return fileDescriptor_15679ebe12a64146, []int{3}
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

func (m *Stats) GetActive() uint32 {
	if m != nil {
		return m.Active
	}
	return 0
}

func (m *Stats) GetPending() uint32 {
	if m != nil {
		return m.Pending
	}
	return 0
}

func (m *Stats) GetDenied() uint32 {
	if m != nil {
		return m.Denied
	}
	return 0
}

func (m *Stats) GetNew() uint32 {
	if m != nil {
		return m.New
	}
	return 0
}

func (m *Stats) GetPreparation() uint32 {
	if m != nil {
		return m.Preparation
	}
	return 0
}

func (m *Stats) GetProduction() uint32 {
	if m != nil {
		return m.Production
	}
	return 0
}

func (m *Stats) GetQuality() uint32 {
	if m != nil {
		return m.Quality
	}
	return 0
}

func (m *Stats) GetReady() uint32 {
	if m != nil {
		return m.Ready
	}
	return 0
}

func (m *Stats) GetPickedup() uint32 {
	if m != nil {
		return m.Pickedup
	}
	return 0
}

func (m *Stats) GetDelivered() uint32 {
	if m != nil {
		return m.Delivered
	}
	return 0
}

func (m *Stats) GetInstalled() uint32 {
	if m != nil {
		return m.Installed
	}
	return 0
}

func (m *Stats) GetCompleted() uint32 {
	if m != nil {
		return m.Completed
	}
	return 0
}

func (m *Stats) GetCancelled() uint32 {
	if m != nil {
		return m.Cancelled
	}
	return 0
}

type QuoteProduct struct {
	Product              *products.Product `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
	Qty                  uint32            `protobuf:"varint,2,opt,name=qty,proto3" json:"qty,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *QuoteProduct) Reset()         { *m = QuoteProduct{} }
func (m *QuoteProduct) String() string { return proto.CompactTextString(m) }
func (*QuoteProduct) ProtoMessage()    {}
func (*QuoteProduct) Descriptor() ([]byte, []int) {
	return fileDescriptor_15679ebe12a64146, []int{4}
}

func (m *QuoteProduct) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuoteProduct.Unmarshal(m, b)
}
func (m *QuoteProduct) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuoteProduct.Marshal(b, m, deterministic)
}
func (m *QuoteProduct) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuoteProduct.Merge(m, src)
}
func (m *QuoteProduct) XXX_Size() int {
	return xxx_messageInfo_QuoteProduct.Size(m)
}
func (m *QuoteProduct) XXX_DiscardUnknown() {
	xxx_messageInfo_QuoteProduct.DiscardUnknown(m)
}

var xxx_messageInfo_QuoteProduct proto.InternalMessageInfo

func (m *QuoteProduct) GetProduct() *products.Product {
	if m != nil {
		return m.Product
	}
	return nil
}

func (m *QuoteProduct) GetQty() uint32 {
	if m != nil {
		return m.Qty
	}
	return 0
}

type Suppliers struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Valid                bool     `protobuf:"varint,2,opt,name=valid,proto3" json:"valid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Suppliers) Reset()         { *m = Suppliers{} }
func (m *Suppliers) String() string { return proto.CompactTextString(m) }
func (*Suppliers) ProtoMessage()    {}
func (*Suppliers) Descriptor() ([]byte, []int) {
	return fileDescriptor_15679ebe12a64146, []int{5}
}

func (m *Suppliers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Suppliers.Unmarshal(m, b)
}
func (m *Suppliers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Suppliers.Marshal(b, m, deterministic)
}
func (m *Suppliers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Suppliers.Merge(m, src)
}
func (m *Suppliers) XXX_Size() int {
	return xxx_messageInfo_Suppliers.Size(m)
}
func (m *Suppliers) XXX_DiscardUnknown() {
	xxx_messageInfo_Suppliers.DiscardUnknown(m)
}

var xxx_messageInfo_Suppliers proto.InternalMessageInfo

func (m *Suppliers) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Suppliers) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

type Quote struct {
	Id             string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	QuotationPrice float32         `protobuf:"fixed32,2,opt,name=quotationPrice,proto3" json:"quotationPrice,omitempty"`
	Phonenumber    string          `protobuf:"bytes,3,opt,name=phonenumber,proto3" json:"phonenumber,omitempty"`
	Timestamp      int64           `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Email          string          `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Supplierids    []*Suppliers    `protobuf:"bytes,6,rep,name=supplierids,proto3" json:"supplierids,omitempty"`
	Products       []*QuoteProduct `protobuf:"bytes,7,rep,name=products,proto3" json:"products,omitempty"`
	Jwt            string          `protobuf:"bytes,9,opt,name=jwt,proto3" json:"jwt,omitempty"`
	Name           string          `protobuf:"bytes,10,opt,name=name,proto3" json:"name,omitempty"`
	City           string          `protobuf:"bytes,11,opt,name=city,proto3" json:"city,omitempty"`
	Address        string          `protobuf:"bytes,12,opt,name=address,proto3" json:"address,omitempty"`
	Zip            string          `protobuf:"bytes,13,opt,name=zip,proto3" json:"zip,omitempty"`
	Size           string          `protobuf:"bytes,14,opt,name=size,proto3" json:"size,omitempty"`
	Userid         string          `protobuf:"bytes,15,opt,name=userid,proto3" json:"userid,omitempty"`
	Status         Status          `protobuf:"varint,16,opt,name=status,proto3,enum=quotes.Status" json:"status,omitempty"`
	SumPrice       float32         `protobuf:"fixed32,17,opt,name=sumPrice,proto3" json:"sumPrice,omitempty"`
	PaidPrice      float32         `protobuf:"fixed32,18,opt,name=paidPrice,proto3" json:"paidPrice,omitempty"`
	// uint32 quoteDate = 2;
	Shipping             ShippingStatus `protobuf:"varint,19,opt,name=shipping,proto3,enum=quotes.ShippingStatus" json:"shipping,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Quote) Reset()         { *m = Quote{} }
func (m *Quote) String() string { return proto.CompactTextString(m) }
func (*Quote) ProtoMessage()    {}
func (*Quote) Descriptor() ([]byte, []int) {
	return fileDescriptor_15679ebe12a64146, []int{6}
}

func (m *Quote) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Quote.Unmarshal(m, b)
}
func (m *Quote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Quote.Marshal(b, m, deterministic)
}
func (m *Quote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Quote.Merge(m, src)
}
func (m *Quote) XXX_Size() int {
	return xxx_messageInfo_Quote.Size(m)
}
func (m *Quote) XXX_DiscardUnknown() {
	xxx_messageInfo_Quote.DiscardUnknown(m)
}

var xxx_messageInfo_Quote proto.InternalMessageInfo

func (m *Quote) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Quote) GetQuotationPrice() float32 {
	if m != nil {
		return m.QuotationPrice
	}
	return 0
}

func (m *Quote) GetPhonenumber() string {
	if m != nil {
		return m.Phonenumber
	}
	return ""
}

func (m *Quote) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Quote) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Quote) GetSupplierids() []*Suppliers {
	if m != nil {
		return m.Supplierids
	}
	return nil
}

func (m *Quote) GetProducts() []*QuoteProduct {
	if m != nil {
		return m.Products
	}
	return nil
}

func (m *Quote) GetJwt() string {
	if m != nil {
		return m.Jwt
	}
	return ""
}

func (m *Quote) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Quote) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *Quote) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Quote) GetZip() string {
	if m != nil {
		return m.Zip
	}
	return ""
}

func (m *Quote) GetSize() string {
	if m != nil {
		return m.Size
	}
	return ""
}

func (m *Quote) GetUserid() string {
	if m != nil {
		return m.Userid
	}
	return ""
}

func (m *Quote) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_NEW
}

func (m *Quote) GetSumPrice() float32 {
	if m != nil {
		return m.SumPrice
	}
	return 0
}

func (m *Quote) GetPaidPrice() float32 {
	if m != nil {
		return m.PaidPrice
	}
	return 0
}

func (m *Quote) GetShipping() ShippingStatus {
	if m != nil {
		return m.Shipping
	}
	return ShippingStatus_PREPARATION
}

func init() {
	proto.RegisterEnum("quotes.Status", Status_name, Status_value)
	proto.RegisterEnum("quotes.ShippingStatus", ShippingStatus_name, ShippingStatus_value)
	proto.RegisterType((*QuoteResponse)(nil), "quotes.QuoteResponse")
	proto.RegisterType((*QuoteParams)(nil), "quotes.QuoteParams")
	proto.RegisterType((*Response)(nil), "quotes.Response")
	proto.RegisterType((*Stats)(nil), "quotes.Stats")
	proto.RegisterType((*QuoteProduct)(nil), "quotes.QuoteProduct")
	proto.RegisterType((*Suppliers)(nil), "quotes.Suppliers")
	proto.RegisterType((*Quote)(nil), "quotes.Quote")
}

func init() { proto.RegisterFile("quotes/quotes.proto", fileDescriptor_15679ebe12a64146) }

var fileDescriptor_15679ebe12a64146 = []byte{
	// 843 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xdd, 0x6e, 0x1a, 0x47,
	0x14, 0x0e, 0x60, 0x03, 0x7b, 0xd6, 0x8b, 0xf1, 0xd8, 0x49, 0x47, 0xa8, 0xaa, 0x2c, 0xa4, 0x46,
	0x56, 0x52, 0xd9, 0x2d, 0xae, 0x7a, 0xef, 0x9a, 0x55, 0x44, 0x45, 0x08, 0x19, 0xbb, 0xed, 0x65,
	0x35, 0x61, 0x46, 0xc9, 0xb4, 0xb0, 0xac, 0x77, 0x66, 0x6d, 0x91, 0xe7, 0xe9, 0x4b, 0xf4, 0x91,
	0xfa, 0x16, 0xd5, 0x39, 0x33, 0xbb, 0x86, 0x24, 0x17, 0xbd, 0x62, 0xbe, 0xef, 0x3b, 0x67, 0xcf,
	0xef, 0x0c, 0x70, 0x7c, 0x57, 0xae, 0x9d, 0xb6, 0x17, 0xfe, 0xe7, 0x3c, 0x2f, 0xd6, 0x6e, 0xcd,
	0xda, 0x1e, 0x0d, 0xbe, 0xca, 0x8b, 0xb5, 0x2a, 0x17, 0xce, 0x5e, 0x54, 0x07, 0x6f, 0x30, 0xfc,
	0x09, 0x92, 0xb7, 0x68, 0x22, 0xb4, 0xcd, 0xd7, 0x99, 0xd5, 0xec, 0x5b, 0x08, 0x3e, 0xbc, 0x71,
	0xda, 0x3a, 0x8b, 0x47, 0xc9, 0x79, 0xf8, 0xa0, 0x37, 0x0b, 0xe2, 0xf0, 0x02, 0x62, 0x22, 0xe6,
	0xb2, 0x90, 0x2b, 0xcb, 0x7a, 0xd0, 0x9c, 0x8c, 0x79, 0xe3, 0xb4, 0x71, 0x16, 0x89, 0xe6, 0x64,
	0xcc, 0xfa, 0xd0, 0xfa, 0xf3, 0xc1, 0xf1, 0x26, 0x11, 0x78, 0x1c, 0x0e, 0xa0, 0x5b, 0xc7, 0xe8,
	0x41, 0xd3, 0xa8, 0xca, 0xda, 0xa8, 0xe1, 0xbf, 0x4d, 0xd8, 0xbf, 0x71, 0xd2, 0x59, 0xf6, 0x0c,
	0xda, 0x72, 0xe1, 0xcc, 0xbd, 0x26, 0x35, 0x11, 0x01, 0x31, 0x0e, 0x9d, 0x5c, 0x67, 0xca, 0x64,
	0xef, 0xe9, 0x9b, 0x89, 0xa8, 0x20, 0x7a, 0x28, 0x9d, 0x19, 0xad, 0x78, 0xcb, 0x7b, 0x78, 0x84,
	0x19, 0x64, 0xfa, 0x81, 0xef, 0x11, 0x89, 0x47, 0x76, 0x0a, 0x71, 0x5e, 0xe8, 0x5c, 0x16, 0xd2,
	0x99, 0x75, 0xc6, 0xf7, 0x49, 0xd9, 0xa6, 0xd8, 0x37, 0x00, 0xa1, 0x3d, 0x68, 0xd0, 0x26, 0x83,
	0x2d, 0x06, 0xb3, 0xb8, 0x2b, 0xe5, 0xd2, 0xb8, 0x0d, 0xef, 0xf8, 0x2c, 0x02, 0x64, 0x27, 0xb0,
	0x5f, 0x68, 0xa9, 0x36, 0xbc, 0x4b, 0xbc, 0x07, 0x6c, 0x00, 0xdd, 0xdc, 0x2c, 0xfe, 0xd2, 0xaa,
	0xcc, 0x79, 0x44, 0x42, 0x8d, 0xd9, 0xd7, 0x10, 0x29, 0xbd, 0x34, 0xf7, 0xba, 0xd0, 0x8a, 0x03,
	0x89, 0x8f, 0x04, 0xaa, 0x26, 0xb3, 0x4e, 0x2e, 0x97, 0x5a, 0xf1, 0xd8, 0xab, 0x35, 0x81, 0xea,
	0x62, 0xbd, 0xca, 0x97, 0xda, 0x69, 0xc5, 0x0f, 0xbc, 0x5a, 0x13, 0xa4, 0xca, 0x6c, 0xa1, 0xc9,
	0x37, 0x09, 0x6a, 0x45, 0x0c, 0x5f, 0xc3, 0x81, 0x1f, 0x9c, 0x2f, 0x8b, 0xbd, 0x84, 0x4e, 0xa8,
	0x90, 0x5a, 0x1e, 0x8f, 0x8e, 0xce, 0xeb, 0x15, 0x09, 0x36, 0xa2, 0xb2, 0xc0, 0xa6, 0xde, 0xb9,
	0x4d, 0x18, 0x01, 0x1e, 0x87, 0x3f, 0x40, 0x74, 0x53, 0xe6, 0xf9, 0xd2, 0xe8, 0xc2, 0x7e, 0x3a,
	0x57, 0xec, 0xca, 0xbd, 0x5c, 0x1a, 0x45, 0x0e, 0x5d, 0xe1, 0xc1, 0xf0, 0xef, 0x3d, 0xd8, 0xa7,
	0x14, 0x3e, 0xb3, 0x7f, 0x0e, 0x3d, 0x5c, 0x2f, 0x1a, 0xc6, 0xbc, 0x30, 0x0b, 0x4d, 0x8e, 0x4d,
	0xf1, 0x09, 0x4b, 0x93, 0xfc, 0xb0, 0xce, 0x74, 0x56, 0xae, 0xde, 0xe9, 0x82, 0x06, 0x1f, 0x89,
	0x6d, 0x0a, 0x7b, 0xe0, 0xcc, 0x4a, 0x5b, 0x27, 0x57, 0x39, 0xed, 0x40, 0x4b, 0x3c, 0x12, 0x98,
	0x97, 0x5e, 0x49, 0xb3, 0xa4, 0x1d, 0x88, 0x84, 0x07, 0xec, 0x12, 0x62, 0x1b, 0x4a, 0x31, 0xca,
	0xf2, 0x36, 0xad, 0xff, 0x51, 0xb5, 0xfe, 0x75, 0x95, 0x62, 0xdb, 0x8a, 0x7d, 0x0f, 0xdd, 0xaa,
	0x5d, 0xbc, 0x43, 0x1e, 0x27, 0x3b, 0x17, 0xa6, 0x6a, 0x61, 0x6d, 0x55, 0x5d, 0x8d, 0xa8, 0xbe,
	0x1a, 0x8c, 0xc1, 0x5e, 0x26, 0x57, 0x9a, 0xb6, 0x20, 0x12, 0x74, 0x46, 0x6e, 0x81, 0x7b, 0x16,
	0x7b, 0x0e, 0xcf, 0xb8, 0x7e, 0x52, 0xa9, 0x42, 0x5b, 0x4b, 0x43, 0x8f, 0x44, 0x05, 0xf1, 0x9b,
	0x1f, 0x4d, 0x4e, 0xc3, 0x8e, 0x04, 0x1e, 0xd1, 0xdf, 0x9a, 0x8f, 0x9a, 0xf7, 0xbc, 0x3f, 0x9e,
	0xf1, 0xaa, 0x94, 0x16, 0xd3, 0xe6, 0x87, 0xc4, 0x06, 0xc4, 0x9e, 0x43, 0xdb, 0x3a, 0xe9, 0x4a,
	0xcb, 0xfb, 0xa7, 0x8d, 0xb3, 0xde, 0xa8, 0x57, 0xd7, 0x4c, 0xac, 0x08, 0x2a, 0xae, 0xb3, 0x2d,
	0x57, 0x7e, 0x30, 0x47, 0x34, 0x98, 0x1a, 0x63, 0xc3, 0x73, 0x69, 0x94, 0x17, 0x19, 0x89, 0x8f,
	0x04, 0x1b, 0x41, 0xd7, 0x7e, 0x30, 0x79, 0x8e, 0xf7, 0xf7, 0x98, 0x62, 0x3c, 0xab, 0x63, 0x04,
	0x3e, 0xc4, 0xaa, 0xed, 0x5e, 0xbc, 0x87, 0xb6, 0xe7, 0x58, 0x07, 0x5a, 0xb3, 0xf4, 0xf7, 0xfe,
	0x13, 0x76, 0x02, 0xfd, 0x9b, 0x5f, 0xe7, 0xf3, 0xe9, 0x24, 0x15, 0x7f, 0x5c, 0xd1, 0xef, 0xb8,
	0xdf, 0x60, 0x0c, 0x7a, 0xd7, 0xd3, 0x49, 0x3a, 0xbb, 0xad, 0xb9, 0x26, 0x7b, 0x0a, 0x47, 0xb5,
	0xa5, 0x48, 0x7f, 0x49, 0xaf, 0x6f, 0xd3, 0x71, 0xbf, 0xc5, 0x8e, 0xe1, 0x30, 0x98, 0xd6, 0xe4,
	0xde, 0x8b, 0x19, 0xf4, 0x76, 0x93, 0x60, 0x87, 0x10, 0xcf, 0x45, 0x3a, 0xbf, 0x12, 0x57, 0xb7,
	0x93, 0x37, 0xb3, 0xfe, 0x13, 0x96, 0x40, 0x34, 0x4e, 0xa7, 0x93, 0xdf, 0x52, 0x41, 0x11, 0x13,
	0x88, 0xae, 0xdf, 0xbc, 0x9e, 0x4f, 0xd3, 0x5b, 0x0a, 0x86, 0xf0, 0x6a, 0x76, 0x9d, 0x4e, 0xa7,
	0x18, 0x64, 0xf4, 0x4f, 0x33, 0x5c, 0xb1, 0x1b, 0x5d, 0xdc, 0x63, 0xf5, 0x2f, 0xa1, 0x3b, 0xd3,
	0x0f, 0x7e, 0xe5, 0x77, 0x9f, 0xd3, 0x41, 0xbf, 0x82, 0xf5, 0xdb, 0x78, 0x09, 0xd1, 0x2b, 0xed,
	0x48, 0xb5, 0xec, 0xf0, 0xf1, 0x2e, 0xbe, 0x2d, 0x75, 0xb1, 0x19, 0x3c, 0xdd, 0x7d, 0x8d, 0x2b,
	0xa7, 0x11, 0x1c, 0x54, 0x4e, 0x3f, 0x6f, 0x26, 0x63, 0x76, 0xbc, 0xbb, 0x83, 0xf4, 0x46, 0x0f,
	0x76, 0x43, 0xb3, 0x1f, 0x21, 0x1e, 0x6b, 0x7c, 0x31, 0x3c, 0xfc, 0xa2, 0xcb, 0xe7, 0xe9, 0x7d,
	0x07, 0x51, 0xaa, 0x8c, 0xfb, 0x9f, 0xc5, 0x5c, 0x40, 0xf2, 0x4a, 0x3b, 0xec, 0xaa, 0xb1, 0xce,
	0x2c, 0xbe, 0x50, 0x50, 0xb2, 0xbd, 0x6b, 0xf6, 0x5d, 0x9b, 0xfe, 0x95, 0x2e, 0xff, 0x0b, 0x00,
	0x00, 0xff, 0xff, 0xaa, 0xae, 0x02, 0xec, 0xcd, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QuoteServiceClient is the client API for QuoteService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QuoteServiceClient interface {
	NewQuote(ctx context.Context, in *Quote, opts ...grpc.CallOption) (*Response, error)
	GetQuotes(ctx context.Context, in *products.Query, opts ...grpc.CallOption) (*QuoteResponse, error)
	GetQuoteByID(ctx context.Context, in *QuoteParams, opts ...grpc.CallOption) (*Quote, error)
	DeleteQuote(ctx context.Context, in *QuoteParams, opts ...grpc.CallOption) (*Response, error)
	EditQuote(ctx context.Context, in *Quote, opts ...grpc.CallOption) (*Response, error)
	GetStatistics(ctx context.Context, in *products.Query, opts ...grpc.CallOption) (*Stats, error)
}

type quoteServiceClient struct {
	cc *grpc.ClientConn
}

func NewQuoteServiceClient(cc *grpc.ClientConn) QuoteServiceClient {
	return &quoteServiceClient{cc}
}

func (c *quoteServiceClient) NewQuote(ctx context.Context, in *Quote, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/quotes.QuoteService/NewQuote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quoteServiceClient) GetQuotes(ctx context.Context, in *products.Query, opts ...grpc.CallOption) (*QuoteResponse, error) {
	out := new(QuoteResponse)
	err := c.cc.Invoke(ctx, "/quotes.QuoteService/GetQuotes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quoteServiceClient) GetQuoteByID(ctx context.Context, in *QuoteParams, opts ...grpc.CallOption) (*Quote, error) {
	out := new(Quote)
	err := c.cc.Invoke(ctx, "/quotes.QuoteService/GetQuoteByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quoteServiceClient) DeleteQuote(ctx context.Context, in *QuoteParams, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/quotes.QuoteService/DeleteQuote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quoteServiceClient) EditQuote(ctx context.Context, in *Quote, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/quotes.QuoteService/EditQuote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quoteServiceClient) GetStatistics(ctx context.Context, in *products.Query, opts ...grpc.CallOption) (*Stats, error) {
	out := new(Stats)
	err := c.cc.Invoke(ctx, "/quotes.QuoteService/GetStatistics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QuoteServiceServer is the server API for QuoteService service.
type QuoteServiceServer interface {
	NewQuote(context.Context, *Quote) (*Response, error)
	GetQuotes(context.Context, *products.Query) (*QuoteResponse, error)
	GetQuoteByID(context.Context, *QuoteParams) (*Quote, error)
	DeleteQuote(context.Context, *QuoteParams) (*Response, error)
	EditQuote(context.Context, *Quote) (*Response, error)
	GetStatistics(context.Context, *products.Query) (*Stats, error)
}

// UnimplementedQuoteServiceServer can be embedded to have forward compatible implementations.
type UnimplementedQuoteServiceServer struct {
}

func (*UnimplementedQuoteServiceServer) NewQuote(ctx context.Context, req *Quote) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewQuote not implemented")
}
func (*UnimplementedQuoteServiceServer) GetQuotes(ctx context.Context, req *products.Query) (*QuoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQuotes not implemented")
}
func (*UnimplementedQuoteServiceServer) GetQuoteByID(ctx context.Context, req *QuoteParams) (*Quote, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQuoteByID not implemented")
}
func (*UnimplementedQuoteServiceServer) DeleteQuote(ctx context.Context, req *QuoteParams) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteQuote not implemented")
}
func (*UnimplementedQuoteServiceServer) EditQuote(ctx context.Context, req *Quote) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditQuote not implemented")
}
func (*UnimplementedQuoteServiceServer) GetStatistics(ctx context.Context, req *products.Query) (*Stats, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatistics not implemented")
}

func RegisterQuoteServiceServer(s *grpc.Server, srv QuoteServiceServer) {
	s.RegisterService(&_QuoteService_serviceDesc, srv)
}

func _QuoteService_NewQuote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Quote)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuoteServiceServer).NewQuote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quotes.QuoteService/NewQuote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuoteServiceServer).NewQuote(ctx, req.(*Quote))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuoteService_GetQuotes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(products.Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuoteServiceServer).GetQuotes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quotes.QuoteService/GetQuotes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuoteServiceServer).GetQuotes(ctx, req.(*products.Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuoteService_GetQuoteByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuoteParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuoteServiceServer).GetQuoteByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quotes.QuoteService/GetQuoteByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuoteServiceServer).GetQuoteByID(ctx, req.(*QuoteParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuoteService_DeleteQuote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuoteParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuoteServiceServer).DeleteQuote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quotes.QuoteService/DeleteQuote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuoteServiceServer).DeleteQuote(ctx, req.(*QuoteParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuoteService_EditQuote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Quote)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuoteServiceServer).EditQuote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quotes.QuoteService/EditQuote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuoteServiceServer).EditQuote(ctx, req.(*Quote))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuoteService_GetStatistics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(products.Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuoteServiceServer).GetStatistics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quotes.QuoteService/GetStatistics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuoteServiceServer).GetStatistics(ctx, req.(*products.Query))
	}
	return interceptor(ctx, in, info, handler)
}

var _QuoteService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "quotes.QuoteService",
	HandlerType: (*QuoteServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewQuote",
			Handler:    _QuoteService_NewQuote_Handler,
		},
		{
			MethodName: "GetQuotes",
			Handler:    _QuoteService_GetQuotes_Handler,
		},
		{
			MethodName: "GetQuoteByID",
			Handler:    _QuoteService_GetQuoteByID_Handler,
		},
		{
			MethodName: "DeleteQuote",
			Handler:    _QuoteService_DeleteQuote_Handler,
		},
		{
			MethodName: "EditQuote",
			Handler:    _QuoteService_EditQuote_Handler,
		},
		{
			MethodName: "GetStatistics",
			Handler:    _QuoteService_GetStatistics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "quotes/quotes.proto",
}
