// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/vessel/vessel.proto

/*/
5	Package vessel is a generated protocol buffer package.
6	It is generated from these files:
7		proto/vessel/vessel.proto
8	It has these top-level messages:
9		Vessel
10		Specification
11		Response
12	*/

package vessel

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Vessel struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Capacity             int32    `protobuf:"varint,2,opt,name=capacity,proto3" json:"capacity,omitempty"`
	MaxWeight            int32    `protobuf:"varint,3,opt,name=max_weight,json=maxWeight,proto3" json:"max_weight,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Available            bool     `protobuf:"varint,5,opt,name=available,proto3" json:"available,omitempty"`
	OwnerId              string   `protobuf:"bytes,6,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Vessel) Reset()         { *m = Vessel{} }
func (m *Vessel) String() string { return proto.CompactTextString(m) }
func (*Vessel) ProtoMessage()    {}
func (*Vessel) Descriptor() ([]byte, []int) {
	return fileDescriptor_04ef66862bb50716, []int{0}
}

func (m *Vessel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vessel.Unmarshal(m, b)
}
func (m *Vessel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vessel.Marshal(b, m, deterministic)
}
func (m *Vessel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vessel.Merge(m, src)
}
func (m *Vessel) XXX_Size() int {
	return xxx_messageInfo_Vessel.Size(m)
}
func (m *Vessel) XXX_DiscardUnknown() {
	xxx_messageInfo_Vessel.DiscardUnknown(m)
}

var xxx_messageInfo_Vessel proto.InternalMessageInfo

func (m *Vessel) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Vessel) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *Vessel) GetMaxWeight() int32 {
	if m != nil {
		return m.MaxWeight
	}
	return 0
}

func (m *Vessel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Vessel) GetAvailable() bool {
	if m != nil {
		return m.Available
	}
	return false
}

func (m *Vessel) GetOwnerId() string {
	if m != nil {
		return m.OwnerId
	}
	return ""
}

type Specification struct {
	Capacity             int32    `protobuf:"varint,1,opt,name=capacity,proto3" json:"capacity,omitempty"`
	MaxWeight            int32    `protobuf:"varint,2,opt,name=max_weight,json=maxWeight,proto3" json:"max_weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Specification) Reset()         { *m = Specification{} }
func (m *Specification) String() string { return proto.CompactTextString(m) }
func (*Specification) ProtoMessage()    {}
func (*Specification) Descriptor() ([]byte, []int) {
	return fileDescriptor_04ef66862bb50716, []int{1}
}

func (m *Specification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Specification.Unmarshal(m, b)
}
func (m *Specification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Specification.Marshal(b, m, deterministic)
}
func (m *Specification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Specification.Merge(m, src)
}
func (m *Specification) XXX_Size() int {
	return xxx_messageInfo_Specification.Size(m)
}
func (m *Specification) XXX_DiscardUnknown() {
	xxx_messageInfo_Specification.DiscardUnknown(m)
}

var xxx_messageInfo_Specification proto.InternalMessageInfo

func (m *Specification) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *Specification) GetMaxWeight() int32 {
	if m != nil {
		return m.MaxWeight
	}
	return 0
}

type Response struct {
	Vessel               *Vessel   `protobuf:"bytes,1,opt,name=vessel,proto3" json:"vessel,omitempty"`
	Vessels              []*Vessel `protobuf:"bytes,2,rep,name=vessels,proto3" json:"vessels,omitempty"`
	Created              bool      `protobuf:"varint,3,opt,name=created,proto3" json:"created,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_04ef66862bb50716, []int{2}
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

func (m *Response) GetVessel() *Vessel {
	if m != nil {
		return m.Vessel
	}
	return nil
}

func (m *Response) GetVessels() []*Vessel {
	if m != nil {
		return m.Vessels
	}
	return nil
}

func (m *Response) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func init() {
	proto.RegisterType((*Vessel)(nil), "vessel.Vessel")
	proto.RegisterType((*Specification)(nil), "vessel.Specification")
	proto.RegisterType((*Response)(nil), "vessel.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for VesselService service

type VesselServiceClient interface {
	FindAvailable(ctx context.Context, in *Specification, opts ...client.CallOption) (*Response, error)
	Create(ctx context.Context, in *Vessel, opts ...client.CallOption) (*Response, error)
}

type vesselServiceClient struct {
	c           client.Client
	serviceName string
}

func NewVesselServiceClient(serviceName string, c client.Client) VesselServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "vessel"
	}
	return &vesselServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *vesselServiceClient) FindAvailable(ctx context.Context, in *Specification, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "VesselService.FindAvailable", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vesselServiceClient) Create(ctx context.Context, in *Vessel, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "VesselService.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for VesselService service

type VesselServiceHandler interface {
	FindAvailable(context.Context, *Specification, *Response) error
	Create(context.Context, *Vessel, *Response) error
}

func RegisterVesselServiceHandler(s server.Server, hdlr VesselServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&VesselService{hdlr}, opts...))
}

type VesselService struct {
	VesselServiceHandler
}

func (h *VesselService) FindAvailable(ctx context.Context, in *Specification, out *Response) error {
	return h.VesselServiceHandler.FindAvailable(ctx, in, out)
}

func (h *VesselService) Create(ctx context.Context, in *Vessel, out *Response) error {
	return h.VesselServiceHandler.Create(ctx, in, out)
}

func init() { proto.RegisterFile("proto/vessel/vessel.proto", fileDescriptor_04ef66862bb50716) }

var fileDescriptor_04ef66862bb50716 = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x75, 0xd3, 0x36, 0x4d, 0x47, 0x5a, 0x64, 0x40, 0xd8, 0x16, 0x85, 0x90, 0x83, 0xe4, 0x20,
	0x15, 0xea, 0xc5, 0xab, 0x08, 0x82, 0x1e, 0xb7, 0xa0, 0xc7, 0xb2, 0xdd, 0x1d, 0x75, 0xa1, 0x4d,
	0x42, 0x12, 0xd2, 0xfa, 0x6f, 0xfc, 0xa9, 0xc2, 0xe4, 0x43, 0x5a, 0xc5, 0xd3, 0xce, 0x7b, 0xf3,
	0xf6, 0xf1, 0xf6, 0x2d, 0x4c, 0xb3, 0x3c, 0x2d, 0xd3, 0x9b, 0x8a, 0x8a, 0x82, 0x36, 0xcd, 0x31,
	0x67, 0x0e, 0xfd, 0x1a, 0x45, 0x5f, 0x02, 0xfc, 0x17, 0x1e, 0x71, 0x02, 0x9e, 0xb3, 0x52, 0x84,
	0x22, 0x1e, 0x29, 0xcf, 0x59, 0x9c, 0x41, 0x60, 0x74, 0xa6, 0x8d, 0x2b, 0x3f, 0xa5, 0x17, 0x8a,
	0x78, 0xa0, 0x3a, 0x8c, 0x97, 0x00, 0x5b, 0xbd, 0x5f, 0xed, 0xc8, 0xbd, 0x7f, 0x94, 0xb2, 0xc7,
	0xdb, 0xd1, 0x56, 0xef, 0x5f, 0x99, 0x40, 0x84, 0x7e, 0xa2, 0xb7, 0x24, 0xfb, 0x6c, 0xc6, 0x33,
	0x5e, 0xc0, 0x48, 0x57, 0xda, 0x6d, 0xf4, 0x7a, 0x43, 0x72, 0x10, 0x8a, 0x38, 0x50, 0x3f, 0x04,
	0x4e, 0x21, 0x48, 0x77, 0x09, 0xe5, 0x2b, 0x67, 0xa5, 0xcf, 0xb7, 0x86, 0x8c, 0x9f, 0x6c, 0xf4,
	0x0c, 0xe3, 0x65, 0x46, 0xc6, 0xbd, 0x39, 0xa3, 0x4b, 0x97, 0x26, 0x07, 0xc1, 0xc4, 0xbf, 0xc1,
	0xbc, 0xa3, 0x60, 0x51, 0x05, 0x81, 0xa2, 0x22, 0x4b, 0x93, 0x82, 0xf0, 0x0a, 0x9a, 0x12, 0xd8,
	0xe4, 0x74, 0x31, 0x99, 0x37, 0x0d, 0xd5, 0x7d, 0xa8, 0x66, 0x8b, 0x31, 0x0c, 0xeb, 0xa9, 0x90,
	0x5e, 0xd8, 0xfb, 0x43, 0xd8, 0xae, 0x51, 0xc2, 0xd0, 0xe4, 0xa4, 0x4b, 0xb2, 0x5c, 0x49, 0xa0,
	0x5a, 0xb8, 0xd8, 0xc1, 0xb8, 0x16, 0x2f, 0x29, 0xaf, 0x9c, 0x21, 0xbc, 0x83, 0xf1, 0xa3, 0x4b,
	0xec, 0x7d, 0x57, 0xc0, 0x79, 0x6b, 0x7a, 0xf0, 0xd6, 0xd9, 0x59, 0x4b, 0x77, 0xb1, 0xaf, 0xc1,
	0x7f, 0x60, 0x57, 0x3c, 0xca, 0xf1, 0x5b, 0x1b, 0x9d, 0xac, 0x7d, 0xfe, 0xee, 0xdb, 0xef, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x23, 0x2f, 0x31, 0xf4, 0x0b, 0x02, 0x00, 0x00,
}
