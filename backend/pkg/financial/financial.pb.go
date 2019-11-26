// Code generated by protoc-gen-go. DO NOT EDIT.
// source: financial/financial.proto

package financial

import (
	context "context"
	fmt "fmt"
	math "math"

	"github.com/Beaxhem/manage/backend/pkg/products"
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

type Request struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8d8ee334e6c4413, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

type PaymentParams struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Amount               float32  `protobuf:"fixed32,2,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PaymentParams) Reset()         { *m = PaymentParams{} }
func (m *PaymentParams) String() string { return proto.CompactTextString(m) }
func (*PaymentParams) ProtoMessage()    {}
func (*PaymentParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8d8ee334e6c4413, []int{1}
}

func (m *PaymentParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaymentParams.Unmarshal(m, b)
}
func (m *PaymentParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaymentParams.Marshal(b, m, deterministic)
}
func (m *PaymentParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentParams.Merge(m, src)
}
func (m *PaymentParams) XXX_Size() int {
	return xxx_messageInfo_PaymentParams.Size(m)
}
func (m *PaymentParams) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentParams.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentParams proto.InternalMessageInfo

func (m *PaymentParams) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PaymentParams) GetAmount() float32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type EmptyResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyResponse) Reset()         { *m = EmptyResponse{} }
func (m *EmptyResponse) String() string { return proto.CompactTextString(m) }
func (*EmptyResponse) ProtoMessage()    {}
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8d8ee334e6c4413, []int{2}
}

func (m *EmptyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyResponse.Unmarshal(m, b)
}
func (m *EmptyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyResponse.Marshal(b, m, deterministic)
}
func (m *EmptyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyResponse.Merge(m, src)
}
func (m *EmptyResponse) XXX_Size() int {
	return xxx_messageInfo_EmptyResponse.Size(m)
}
func (m *EmptyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyResponse proto.InternalMessageInfo

type Expenses struct {
	Expense              []*Expense `protobuf:"bytes,1,rep,name=expense,proto3" json:"expense,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Expenses) Reset()         { *m = Expenses{} }
func (m *Expenses) String() string { return proto.CompactTextString(m) }
func (*Expenses) ProtoMessage()    {}
func (*Expenses) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8d8ee334e6c4413, []int{3}
}

func (m *Expenses) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Expenses.Unmarshal(m, b)
}
func (m *Expenses) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Expenses.Marshal(b, m, deterministic)
}
func (m *Expenses) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Expenses.Merge(m, src)
}
func (m *Expenses) XXX_Size() int {
	return xxx_messageInfo_Expenses.Size(m)
}
func (m *Expenses) XXX_DiscardUnknown() {
	xxx_messageInfo_Expenses.DiscardUnknown(m)
}

var xxx_messageInfo_Expenses proto.InternalMessageInfo

func (m *Expenses) GetExpense() []*Expense {
	if m != nil {
		return m.Expense
	}
	return nil
}

type Expense struct {
	Timestamp            int64    `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Amount               float32  `protobuf:"fixed32,3,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Expense) Reset()         { *m = Expense{} }
func (m *Expense) String() string { return proto.CompactTextString(m) }
func (*Expense) ProtoMessage()    {}
func (*Expense) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8d8ee334e6c4413, []int{4}
}

func (m *Expense) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Expense.Unmarshal(m, b)
}
func (m *Expense) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Expense.Marshal(b, m, deterministic)
}
func (m *Expense) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Expense.Merge(m, src)
}
func (m *Expense) XXX_Size() int {
	return xxx_messageInfo_Expense.Size(m)
}
func (m *Expense) XXX_DiscardUnknown() {
	xxx_messageInfo_Expense.DiscardUnknown(m)
}

var xxx_messageInfo_Expense proto.InternalMessageInfo

func (m *Expense) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Expense) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Expense) GetAmount() float32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type Params struct {
	Name                 string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Amount               float32         `protobuf:"fixed32,2,opt,name=amount,proto3" json:"amount,omitempty"`
	To                   string          `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	Query                *products.Query `protobuf:"bytes,4,opt,name=query,proto3" json:"query,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8d8ee334e6c4413, []int{5}
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

func (m *Params) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Params) GetAmount() float32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Params) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *Params) GetQuery() *products.Query {
	if m != nil {
		return m.Query
	}
	return nil
}

type Payoff struct {
	Supplierid           string          `protobuf:"bytes,1,opt,name=supplierid,proto3" json:"supplierid,omitempty"`
	Timestamp            int64           `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Toreceive            float32         `protobuf:"fixed32,3,opt,name=toreceive,proto3" json:"toreceive,omitempty"`
	Paid                 float32         `protobuf:"fixed32,4,opt,name=paid,proto3" json:"paid,omitempty"`
	Quoteid              string          `protobuf:"bytes,5,opt,name=quoteid,proto3" json:"quoteid,omitempty"`
	Sectors              []*SectorProfit `protobuf:"bytes,6,rep,name=sectors,proto3" json:"sectors,omitempty"`
	Profitless           float32         `protobuf:"fixed32,7,opt,name=profitless,proto3" json:"profitless,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Payoff) Reset()         { *m = Payoff{} }
func (m *Payoff) String() string { return proto.CompactTextString(m) }
func (*Payoff) ProtoMessage()    {}
func (*Payoff) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8d8ee334e6c4413, []int{6}
}

func (m *Payoff) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Payoff.Unmarshal(m, b)
}
func (m *Payoff) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Payoff.Marshal(b, m, deterministic)
}
func (m *Payoff) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Payoff.Merge(m, src)
}
func (m *Payoff) XXX_Size() int {
	return xxx_messageInfo_Payoff.Size(m)
}
func (m *Payoff) XXX_DiscardUnknown() {
	xxx_messageInfo_Payoff.DiscardUnknown(m)
}

var xxx_messageInfo_Payoff proto.InternalMessageInfo

func (m *Payoff) GetSupplierid() string {
	if m != nil {
		return m.Supplierid
	}
	return ""
}

func (m *Payoff) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Payoff) GetToreceive() float32 {
	if m != nil {
		return m.Toreceive
	}
	return 0
}

func (m *Payoff) GetPaid() float32 {
	if m != nil {
		return m.Paid
	}
	return 0
}

func (m *Payoff) GetQuoteid() string {
	if m != nil {
		return m.Quoteid
	}
	return ""
}

func (m *Payoff) GetSectors() []*SectorProfit {
	if m != nil {
		return m.Sectors
	}
	return nil
}

func (m *Payoff) GetProfitless() float32 {
	if m != nil {
		return m.Profitless
	}
	return 0
}

type SectorProfit struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Amount               float32  `protobuf:"fixed32,2,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SectorProfit) Reset()         { *m = SectorProfit{} }
func (m *SectorProfit) String() string { return proto.CompactTextString(m) }
func (*SectorProfit) ProtoMessage()    {}
func (*SectorProfit) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8d8ee334e6c4413, []int{7}
}

func (m *SectorProfit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SectorProfit.Unmarshal(m, b)
}
func (m *SectorProfit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SectorProfit.Marshal(b, m, deterministic)
}
func (m *SectorProfit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SectorProfit.Merge(m, src)
}
func (m *SectorProfit) XXX_Size() int {
	return xxx_messageInfo_SectorProfit.Size(m)
}
func (m *SectorProfit) XXX_DiscardUnknown() {
	xxx_messageInfo_SectorProfit.DiscardUnknown(m)
}

var xxx_messageInfo_SectorProfit proto.InternalMessageInfo

func (m *SectorProfit) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SectorProfit) GetAmount() float32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type Response struct {
	Income               []*Payoff `protobuf:"bytes,1,rep,name=income,proto3" json:"income,omitempty"`
	Banks                []*Bank   `protobuf:"bytes,2,rep,name=banks,proto3" json:"banks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8d8ee334e6c4413, []int{8}
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

func (m *Response) GetIncome() []*Payoff {
	if m != nil {
		return m.Income
	}
	return nil
}

func (m *Response) GetBanks() []*Bank {
	if m != nil {
		return m.Banks
	}
	return nil
}

type Bank struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Money                float32  `protobuf:"fixed32,2,opt,name=money,proto3" json:"money,omitempty"`
	Color                string   `protobuf:"bytes,3,opt,name=color,proto3" json:"color,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Bank) Reset()         { *m = Bank{} }
func (m *Bank) String() string { return proto.CompactTextString(m) }
func (*Bank) ProtoMessage()    {}
func (*Bank) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8d8ee334e6c4413, []int{9}
}

func (m *Bank) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bank.Unmarshal(m, b)
}
func (m *Bank) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bank.Marshal(b, m, deterministic)
}
func (m *Bank) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bank.Merge(m, src)
}
func (m *Bank) XXX_Size() int {
	return xxx_messageInfo_Bank.Size(m)
}
func (m *Bank) XXX_DiscardUnknown() {
	xxx_messageInfo_Bank.DiscardUnknown(m)
}

var xxx_messageInfo_Bank proto.InternalMessageInfo

func (m *Bank) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Bank) GetMoney() float32 {
	if m != nil {
		return m.Money
	}
	return 0
}

func (m *Bank) GetColor() string {
	if m != nil {
		return m.Color
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "financial.Request")
	proto.RegisterType((*PaymentParams)(nil), "financial.PaymentParams")
	proto.RegisterType((*EmptyResponse)(nil), "financial.EmptyResponse")
	proto.RegisterType((*Expenses)(nil), "financial.Expenses")
	proto.RegisterType((*Expense)(nil), "financial.Expense")
	proto.RegisterType((*Params)(nil), "financial.Params")
	proto.RegisterType((*Payoff)(nil), "financial.Payoff")
	proto.RegisterType((*SectorProfit)(nil), "financial.SectorProfit")
	proto.RegisterType((*Response)(nil), "financial.Response")
	proto.RegisterType((*Bank)(nil), "financial.Bank")
}

func init() { proto.RegisterFile("financial/financial.proto", fileDescriptor_a8d8ee334e6c4413) }

var fileDescriptor_a8d8ee334e6c4413 = []byte{
	// 546 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x4f, 0x8f, 0xd3, 0x3e,
	0x10, 0x55, 0xd2, 0x6d, 0xd3, 0x4c, 0x7f, 0xfd, 0x15, 0x0c, 0x62, 0x43, 0x85, 0x50, 0x15, 0x09,
	0xa9, 0x48, 0xa8, 0x2b, 0x5a, 0xd0, 0x02, 0x47, 0xc4, 0x6e, 0xc5, 0x05, 0x15, 0x97, 0x23, 0x17,
	0x6f, 0x3a, 0x95, 0xac, 0x36, 0x76, 0x1a, 0x3b, 0x0b, 0xf9, 0xb0, 0x1c, 0xf9, 0x1e, 0xc8, 0xce,
	0x9f, 0x26, 0xab, 0x22, 0xe0, 0x36, 0xf3, 0xfc, 0x9c, 0x79, 0x6f, 0xc6, 0x13, 0x78, 0xbc, 0xe5,
	0x82, 0x89, 0x88, 0xb3, 0xfd, 0x45, 0x1d, 0xcd, 0x92, 0x54, 0x6a, 0x49, 0xfc, 0x1a, 0x18, 0x9f,
	0x27, 0xa9, 0xdc, 0x64, 0x91, 0x56, 0x17, 0x55, 0x50, 0x70, 0x42, 0x1f, 0x3c, 0x8a, 0x87, 0x0c,
	0x95, 0x0e, 0x2f, 0x61, 0xb8, 0x62, 0x79, 0x8c, 0x42, 0xaf, 0x58, 0xca, 0x62, 0x45, 0xfe, 0x07,
	0x97, 0x6f, 0x02, 0x67, 0xe2, 0x4c, 0x7d, 0xea, 0xf2, 0x0d, 0x79, 0x04, 0x3d, 0x16, 0xcb, 0x4c,
	0xe8, 0xc0, 0x9d, 0x38, 0x53, 0x97, 0x96, 0x59, 0x38, 0x82, 0xe1, 0x55, 0x9c, 0xe8, 0x9c, 0xa2,
	0x4a, 0xa4, 0x50, 0x18, 0xbe, 0x81, 0xfe, 0xd5, 0xf7, 0x04, 0x85, 0x42, 0x45, 0x5e, 0x80, 0x87,
	0x45, 0x1c, 0x38, 0x93, 0xce, 0x74, 0x30, 0x27, 0xb3, 0xa3, 0xce, 0x92, 0x45, 0x2b, 0x4a, 0xb8,
	0x06, 0xaf, 0xc4, 0xc8, 0x13, 0xf0, 0x35, 0x8f, 0x51, 0x69, 0x16, 0x27, 0x56, 0x44, 0x87, 0x1e,
	0x01, 0x42, 0xe0, 0x4c, 0xb0, 0x18, 0xad, 0x12, 0x9f, 0xda, 0xb8, 0xa1, 0xaf, 0xd3, 0xd2, 0xb7,
	0x83, 0x5e, 0xe9, 0xa8, 0xba, 0xe5, 0x9c, 0xbc, 0xd5, 0x72, 0x65, 0xdc, 0x6b, 0x69, 0xbf, 0xe4,
	0x53, 0x57, 0x4b, 0xf2, 0x0c, 0xba, 0x87, 0x0c, 0xd3, 0x3c, 0x38, 0x9b, 0x38, 0xd3, 0xc1, 0x7c,
	0x34, 0xab, 0x3b, 0xf9, 0xd9, 0xc0, 0xb4, 0x38, 0x0d, 0x7f, 0x3a, 0xa6, 0x5a, 0x2e, 0xb7, 0x5b,
	0xf2, 0x14, 0x40, 0x65, 0x49, 0xb2, 0xe7, 0x98, 0xd6, 0x7d, 0x6c, 0x20, 0x6d, 0x87, 0xee, 0x5d,
	0x87, 0xe6, 0x54, 0xa6, 0x18, 0x21, 0xbf, 0xc5, 0xd2, 0xd0, 0x11, 0x30, 0x4e, 0x12, 0xc6, 0x37,
	0x56, 0x8c, 0x4b, 0x6d, 0x4c, 0x02, 0xf0, 0x0e, 0x99, 0xd4, 0xc8, 0x37, 0x41, 0xd7, 0x16, 0xab,
	0x52, 0xf2, 0x12, 0x3c, 0x85, 0x91, 0x96, 0xa9, 0x0a, 0x7a, 0x76, 0x08, 0xe7, 0x8d, 0x21, 0xac,
	0xed, 0xc9, 0x2a, 0x95, 0x5b, 0xae, 0x69, 0xc5, 0x33, 0xe2, 0x13, 0x0b, 0xed, 0x51, 0xa9, 0xc0,
	0xb3, 0x65, 0x1a, 0x48, 0xf8, 0x0e, 0xfe, 0x6b, 0x5e, 0xfc, 0x97, 0xd6, 0x86, 0x5f, 0xa1, 0x5f,
	0xbd, 0x15, 0xf2, 0x1c, 0x7a, 0x5c, 0x44, 0x32, 0xae, 0x9e, 0xc7, 0xfd, 0x86, 0xb2, 0xa2, 0x8f,
	0xb4, 0x24, 0x98, 0x09, 0xdc, 0x30, 0xb1, 0x53, 0x81, 0x6b, 0x99, 0xa3, 0x06, 0xf3, 0x3d, 0x13,
	0x3b, 0x5a, 0x9c, 0x86, 0xd7, 0x70, 0x66, 0xd2, 0x93, 0x8a, 0x1e, 0x42, 0x37, 0x96, 0x02, 0xf3,
	0x52, 0x50, 0x91, 0x18, 0x34, 0x92, 0x7b, 0x99, 0x96, 0xd3, 0x2e, 0x92, 0xf9, 0x0f, 0x17, 0xee,
	0x5d, 0x57, 0x15, 0xd6, 0x98, 0xde, 0xf2, 0x08, 0xc9, 0x02, 0xfc, 0x25, 0xea, 0x8f, 0x85, 0xa0,
	0xb6, 0x56, 0xf3, 0xc2, 0xc6, 0x0f, 0x1a, 0x50, 0xed, 0xf1, 0x15, 0x78, 0x9f, 0xf0, 0x9b, 0x15,
	0x75, 0x57, 0xf4, 0x38, 0x68, 0xae, 0x43, 0x73, 0x8b, 0xc8, 0x02, 0xfa, 0x4b, 0xd4, 0x86, 0xa4,
	0x08, 0x69, 0x7d, 0xd6, 0xee, 0xeb, 0xe9, 0x52, 0x97, 0x30, 0xfc, 0x22, 0x3f, 0xa0, 0xd2, 0x5c,
	0x30, 0xcd, 0xa5, 0xf8, 0x6b, 0x8d, 0xaf, 0x61, 0xb0, 0x44, 0x5d, 0xaf, 0xed, 0x1f, 0xae, 0xd5,
	0xbc, 0xb7, 0xd0, 0x59, 0xb1, 0x9c, 0x04, 0xed, 0xa9, 0x1d, 0x7f, 0x22, 0xbf, 0xf7, 0x77, 0xd3,
	0xb3, 0x7f, 0xa0, 0xc5, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7a, 0x2c, 0x03, 0xc2, 0xc2, 0x04,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FinancialServiceClient is the client API for FinancialService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FinancialServiceClient interface {
	GetIncome(ctx context.Context, in *Params, opts ...grpc.CallOption) (*Response, error)
	NewBank(ctx context.Context, in *Bank, opts ...grpc.CallOption) (*EmptyResponse, error)
	GetBanks(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	ToDestination(ctx context.Context, in *Params, opts ...grpc.CallOption) (*Response, error)
	GetExpenses(ctx context.Context, in *Params, opts ...grpc.CallOption) (*Expenses, error)
	Pay(ctx context.Context, in *PaymentParams, opts ...grpc.CallOption) (*EmptyResponse, error)
}

type financialServiceClient struct {
	cc *grpc.ClientConn
}

func NewFinancialServiceClient(cc *grpc.ClientConn) FinancialServiceClient {
	return &financialServiceClient{cc}
}

func (c *financialServiceClient) GetIncome(ctx context.Context, in *Params, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/financial.FinancialService/GetIncome", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financialServiceClient) NewBank(ctx context.Context, in *Bank, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/financial.FinancialService/NewBank", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financialServiceClient) GetBanks(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/financial.FinancialService/GetBanks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financialServiceClient) ToDestination(ctx context.Context, in *Params, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/financial.FinancialService/ToDestination", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financialServiceClient) GetExpenses(ctx context.Context, in *Params, opts ...grpc.CallOption) (*Expenses, error) {
	out := new(Expenses)
	err := c.cc.Invoke(ctx, "/financial.FinancialService/GetExpenses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financialServiceClient) Pay(ctx context.Context, in *PaymentParams, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/financial.FinancialService/Pay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FinancialServiceServer is the server API for FinancialService service.
type FinancialServiceServer interface {
	GetIncome(context.Context, *Params) (*Response, error)
	NewBank(context.Context, *Bank) (*EmptyResponse, error)
	GetBanks(context.Context, *Request) (*Response, error)
	ToDestination(context.Context, *Params) (*Response, error)
	GetExpenses(context.Context, *Params) (*Expenses, error)
	Pay(context.Context, *PaymentParams) (*EmptyResponse, error)
}

// UnimplementedFinancialServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFinancialServiceServer struct {
}

func (*UnimplementedFinancialServiceServer) GetIncome(ctx context.Context, req *Params) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIncome not implemented")
}
func (*UnimplementedFinancialServiceServer) NewBank(ctx context.Context, req *Bank) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewBank not implemented")
}
func (*UnimplementedFinancialServiceServer) GetBanks(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBanks not implemented")
}
func (*UnimplementedFinancialServiceServer) ToDestination(ctx context.Context, req *Params) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ToDestination not implemented")
}
func (*UnimplementedFinancialServiceServer) GetExpenses(ctx context.Context, req *Params) (*Expenses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExpenses not implemented")
}
func (*UnimplementedFinancialServiceServer) Pay(ctx context.Context, req *PaymentParams) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pay not implemented")
}

func RegisterFinancialServiceServer(s *grpc.Server, srv FinancialServiceServer) {
	s.RegisterService(&_FinancialService_serviceDesc, srv)
}

func _FinancialService_GetIncome_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Params)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinancialServiceServer).GetIncome(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/financial.FinancialService/GetIncome",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinancialServiceServer).GetIncome(ctx, req.(*Params))
	}
	return interceptor(ctx, in, info, handler)
}

func _FinancialService_NewBank_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Bank)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinancialServiceServer).NewBank(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/financial.FinancialService/NewBank",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinancialServiceServer).NewBank(ctx, req.(*Bank))
	}
	return interceptor(ctx, in, info, handler)
}

func _FinancialService_GetBanks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinancialServiceServer).GetBanks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/financial.FinancialService/GetBanks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinancialServiceServer).GetBanks(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _FinancialService_ToDestination_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Params)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinancialServiceServer).ToDestination(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/financial.FinancialService/ToDestination",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinancialServiceServer).ToDestination(ctx, req.(*Params))
	}
	return interceptor(ctx, in, info, handler)
}

func _FinancialService_GetExpenses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Params)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinancialServiceServer).GetExpenses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/financial.FinancialService/GetExpenses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinancialServiceServer).GetExpenses(ctx, req.(*Params))
	}
	return interceptor(ctx, in, info, handler)
}

func _FinancialService_Pay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaymentParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinancialServiceServer).Pay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/financial.FinancialService/Pay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinancialServiceServer).Pay(ctx, req.(*PaymentParams))
	}
	return interceptor(ctx, in, info, handler)
}

var _FinancialService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "financial.FinancialService",
	HandlerType: (*FinancialServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetIncome",
			Handler:    _FinancialService_GetIncome_Handler,
		},
		{
			MethodName: "NewBank",
			Handler:    _FinancialService_NewBank_Handler,
		},
		{
			MethodName: "GetBanks",
			Handler:    _FinancialService_GetBanks_Handler,
		},
		{
			MethodName: "ToDestination",
			Handler:    _FinancialService_ToDestination_Handler,
		},
		{
			MethodName: "GetExpenses",
			Handler:    _FinancialService_GetExpenses_Handler,
		},
		{
			MethodName: "Pay",
			Handler:    _FinancialService_Pay_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "financial/financial.proto",
}
