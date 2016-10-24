// Code generated by protoc-gen-go.
// source: proto/mail/mail.proto
// DO NOT EDIT!

/*
Package mail is a generated protocol buffer package.

It is generated from these files:
	proto/mail/mail.proto

It has these top-level messages:
	Message
	Recipient
	SendTemplateRequest
	SendTemplateResponse
*/
package mail

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type Message struct {
	From     string `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	FromName string `protobuf:"bytes,2,opt,name=fromName" json:"fromName,omitempty"`
	To       string `protobuf:"bytes,3,opt,name=to" json:"to,omitempty"`
	ToName   string `protobuf:"bytes,4,opt,name=toName" json:"toName,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Recipient struct {
	Name  string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Email string `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
}

func (m *Recipient) Reset()                    { *m = Recipient{} }
func (m *Recipient) String() string            { return proto.CompactTextString(m) }
func (*Recipient) ProtoMessage()               {}
func (*Recipient) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type SendTemplateRequest struct {
	Message   *Message          `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	Template  string            `protobuf:"bytes,2,opt,name=template" json:"template,omitempty"`
	Variables map[string]string `protobuf:"bytes,3,rep,name=variables" json:"variables,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *SendTemplateRequest) Reset()                    { *m = SendTemplateRequest{} }
func (m *SendTemplateRequest) String() string            { return proto.CompactTextString(m) }
func (*SendTemplateRequest) ProtoMessage()               {}
func (*SendTemplateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SendTemplateRequest) GetMessage() *Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *SendTemplateRequest) GetVariables() map[string]string {
	if m != nil {
		return m.Variables
	}
	return nil
}

type SendTemplateResponse struct {
	MessageID string `protobuf:"bytes,1,opt,name=messageID" json:"messageID,omitempty"`
}

func (m *SendTemplateResponse) Reset()                    { *m = SendTemplateResponse{} }
func (m *SendTemplateResponse) String() string            { return proto.CompactTextString(m) }
func (*SendTemplateResponse) ProtoMessage()               {}
func (*SendTemplateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*Message)(nil), "Message")
	proto.RegisterType((*Recipient)(nil), "Recipient")
	proto.RegisterType((*SendTemplateRequest)(nil), "SendTemplateRequest")
	proto.RegisterType((*SendTemplateResponse)(nil), "SendTemplateResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Mailing service

type MailingClient interface {
	SendTemplate(ctx context.Context, in *SendTemplateRequest, opts ...client.CallOption) (*SendTemplateResponse, error)
}

type mailingClient struct {
	c           client.Client
	serviceName string
}

func NewMailingClient(serviceName string, c client.Client) MailingClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "mailing"
	}
	return &mailingClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *mailingClient) SendTemplate(ctx context.Context, in *SendTemplateRequest, opts ...client.CallOption) (*SendTemplateResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Mailing.SendTemplate", in)
	out := new(SendTemplateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Mailing service

type MailingHandler interface {
	SendTemplate(context.Context, *SendTemplateRequest, *SendTemplateResponse) error
}

func RegisterMailingHandler(s server.Server, hdlr MailingHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Mailing{hdlr}, opts...))
}

type Mailing struct {
	MailingHandler
}

func (h *Mailing) SendTemplate(ctx context.Context, in *SendTemplateRequest, out *SendTemplateResponse) error {
	return h.MailingHandler.SendTemplate(ctx, in, out)
}

func init() { proto.RegisterFile("proto/mail/mail.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x91, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0x6d, 0x3b, 0xb7, 0xf5, 0x4d, 0x86, 0x3c, 0x37, 0x29, 0xc3, 0xc3, 0x88, 0x97, 0x9d,
	0x2a, 0x4c, 0x05, 0x11, 0x3d, 0x08, 0x0a, 0x7a, 0x98, 0x87, 0x2a, 0xde, 0x33, 0x7d, 0x8e, 0x60,
	0x9b, 0xd4, 0x26, 0x1b, 0xec, 0x1f, 0xf5, 0xef, 0x91, 0x64, 0xa9, 0x73, 0xd0, 0x4b, 0xfb, 0xbe,
	0x2f, 0x2f, 0xef, 0xf7, 0x3d, 0x02, 0xc3, 0xb2, 0x52, 0x46, 0x9d, 0x15, 0x5c, 0xe4, 0xee, 0x93,
	0x3a, 0xcd, 0x38, 0x74, 0x66, 0xa4, 0x35, 0x5f, 0x10, 0x22, 0xb4, 0x3e, 0x2b, 0x55, 0x24, 0xc1,
	0x38, 0x98, 0xc4, 0x99, 0xab, 0x71, 0x04, 0x5d, 0xfb, 0x7f, 0xe6, 0x05, 0x25, 0xa1, 0xf3, 0xff,
	0x34, 0xf6, 0x21, 0x34, 0x2a, 0x89, 0x9c, 0x1b, 0x1a, 0x85, 0xc7, 0xd0, 0x36, 0xca, 0x75, 0xb6,
	0x9c, 0xe7, 0x15, 0xbb, 0x84, 0x38, 0xa3, 0x77, 0x51, 0x0a, 0x92, 0xc6, 0x42, 0xa4, 0x6d, 0xf1,
	0x10, 0x5b, 0xe3, 0x00, 0xf6, 0xc9, 0x46, 0xf2, 0x84, 0x8d, 0x60, 0x3f, 0x01, 0x1c, 0xbd, 0x90,
	0xfc, 0x78, 0xa5, 0xa2, 0xcc, 0xb9, 0xa1, 0x8c, 0xbe, 0x97, 0xa4, 0x0d, 0x32, 0xe8, 0x14, 0x9b,
	0xc4, 0x6e, 0x48, 0x6f, 0xda, 0x4d, 0xfd, 0x06, 0x59, 0x7d, 0x60, 0x63, 0x1b, 0x7f, 0xad, 0x8e,
	0x5d, 0x6b, 0xbc, 0x83, 0x78, 0xc5, 0x2b, 0xc1, 0xe7, 0x39, 0xe9, 0x24, 0x1a, 0x47, 0x93, 0xde,
	0xf4, 0x34, 0x6d, 0x00, 0xa5, 0x6f, 0x75, 0xd7, 0x83, 0x34, 0xd5, 0x3a, 0xdb, 0xde, 0x1a, 0xdd,
	0x40, 0x7f, 0xf7, 0x10, 0x0f, 0x21, 0xfa, 0xa2, 0xb5, 0xdf, 0xca, 0x96, 0x76, 0xa9, 0x15, 0xcf,
	0x97, 0x35, 0x7f, 0x23, 0xae, 0xc3, 0xab, 0x80, 0x5d, 0xc0, 0x60, 0x17, 0xa7, 0x4b, 0x25, 0x35,
	0xe1, 0x09, 0xc4, 0x3e, 0xff, 0xd3, 0xbd, 0x9f, 0xb4, 0x35, 0xa6, 0x8f, 0xd0, 0x99, 0x71, 0x91,
	0x0b, 0xb9, 0xc0, 0x5b, 0x38, 0xf8, 0x3f, 0x00, 0x07, 0x4d, 0xf1, 0x47, 0xc3, 0xb4, 0x89, 0xc2,
	0xf6, 0xe6, 0x6d, 0xf7, 0xf2, 0xe7, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x92, 0xfb, 0x0d, 0x6c,
	0x12, 0x02, 0x00, 0x00,
}