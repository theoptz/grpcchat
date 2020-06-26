// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/chat/chat.proto

package chat

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type LoginRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_bf91dfeaf7880616, []int{0}
}
func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (dst *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(dst, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type LoginResponse struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_bf91dfeaf7880616, []int{1}
}
func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (dst *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(dst, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type LogoutRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogoutRequest) Reset()         { *m = LogoutRequest{} }
func (m *LogoutRequest) String() string { return proto.CompactTextString(m) }
func (*LogoutRequest) ProtoMessage()    {}
func (*LogoutRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_bf91dfeaf7880616, []int{2}
}
func (m *LogoutRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogoutRequest.Unmarshal(m, b)
}
func (m *LogoutRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogoutRequest.Marshal(b, m, deterministic)
}
func (dst *LogoutRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogoutRequest.Merge(dst, src)
}
func (m *LogoutRequest) XXX_Size() int {
	return xxx_messageInfo_LogoutRequest.Size(m)
}
func (m *LogoutRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LogoutRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LogoutRequest proto.InternalMessageInfo

func (m *LogoutRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type LogoutResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogoutResponse) Reset()         { *m = LogoutResponse{} }
func (m *LogoutResponse) String() string { return proto.CompactTextString(m) }
func (*LogoutResponse) ProtoMessage()    {}
func (*LogoutResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_bf91dfeaf7880616, []int{3}
}
func (m *LogoutResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogoutResponse.Unmarshal(m, b)
}
func (m *LogoutResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogoutResponse.Marshal(b, m, deterministic)
}
func (dst *LogoutResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogoutResponse.Merge(dst, src)
}
func (m *LogoutResponse) XXX_Size() int {
	return xxx_messageInfo_LogoutResponse.Size(m)
}
func (m *LogoutResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LogoutResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LogoutResponse proto.InternalMessageInfo

type StreamRequest struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamRequest) Reset()         { *m = StreamRequest{} }
func (m *StreamRequest) String() string { return proto.CompactTextString(m) }
func (*StreamRequest) ProtoMessage()    {}
func (*StreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_bf91dfeaf7880616, []int{4}
}
func (m *StreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamRequest.Unmarshal(m, b)
}
func (m *StreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamRequest.Marshal(b, m, deterministic)
}
func (dst *StreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamRequest.Merge(dst, src)
}
func (m *StreamRequest) XXX_Size() int {
	return xxx_messageInfo_StreamRequest.Size(m)
}
func (m *StreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamRequest proto.InternalMessageInfo

func (m *StreamRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type StreamResponse struct {
	Timestamp *timestamp.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Types that are valid to be assigned to Event:
	//	*StreamResponse_UserLogin
	//	*StreamResponse_UserLogout
	//	*StreamResponse_UserMessage
	//	*StreamResponse_ServerShutdown
	Event                isStreamResponse_Event `protobuf_oneof:"event"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *StreamResponse) Reset()         { *m = StreamResponse{} }
func (m *StreamResponse) String() string { return proto.CompactTextString(m) }
func (*StreamResponse) ProtoMessage()    {}
func (*StreamResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_bf91dfeaf7880616, []int{5}
}
func (m *StreamResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamResponse.Unmarshal(m, b)
}
func (m *StreamResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamResponse.Marshal(b, m, deterministic)
}
func (dst *StreamResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamResponse.Merge(dst, src)
}
func (m *StreamResponse) XXX_Size() int {
	return xxx_messageInfo_StreamResponse.Size(m)
}
func (m *StreamResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamResponse proto.InternalMessageInfo

type isStreamResponse_Event interface {
	isStreamResponse_Event()
}

type StreamResponse_UserLogin struct {
	UserLogin *StreamResponse_Login `protobuf:"bytes,2,opt,name=user_login,json=userLogin,proto3,oneof"`
}
type StreamResponse_UserLogout struct {
	UserLogout *StreamResponse_Logout `protobuf:"bytes,3,opt,name=user_logout,json=userLogout,proto3,oneof"`
}
type StreamResponse_UserMessage struct {
	UserMessage *StreamResponse_Message `protobuf:"bytes,4,opt,name=user_message,json=userMessage,proto3,oneof"`
}
type StreamResponse_ServerShutdown struct {
	ServerShutdown *StreamResponse_Shutdown `protobuf:"bytes,5,opt,name=server_shutdown,json=serverShutdown,proto3,oneof"`
}

func (*StreamResponse_UserLogin) isStreamResponse_Event()      {}
func (*StreamResponse_UserLogout) isStreamResponse_Event()     {}
func (*StreamResponse_UserMessage) isStreamResponse_Event()    {}
func (*StreamResponse_ServerShutdown) isStreamResponse_Event() {}

func (m *StreamResponse) GetEvent() isStreamResponse_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *StreamResponse) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *StreamResponse) GetUserLogin() *StreamResponse_Login {
	if x, ok := m.GetEvent().(*StreamResponse_UserLogin); ok {
		return x.UserLogin
	}
	return nil
}

func (m *StreamResponse) GetUserLogout() *StreamResponse_Logout {
	if x, ok := m.GetEvent().(*StreamResponse_UserLogout); ok {
		return x.UserLogout
	}
	return nil
}

func (m *StreamResponse) GetUserMessage() *StreamResponse_Message {
	if x, ok := m.GetEvent().(*StreamResponse_UserMessage); ok {
		return x.UserMessage
	}
	return nil
}

func (m *StreamResponse) GetServerShutdown() *StreamResponse_Shutdown {
	if x, ok := m.GetEvent().(*StreamResponse_ServerShutdown); ok {
		return x.ServerShutdown
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*StreamResponse) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _StreamResponse_OneofMarshaler, _StreamResponse_OneofUnmarshaler, _StreamResponse_OneofSizer, []interface{}{
		(*StreamResponse_UserLogin)(nil),
		(*StreamResponse_UserLogout)(nil),
		(*StreamResponse_UserMessage)(nil),
		(*StreamResponse_ServerShutdown)(nil),
	}
}

func _StreamResponse_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*StreamResponse)
	// event
	switch x := m.Event.(type) {
	case *StreamResponse_UserLogin:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.UserLogin); err != nil {
			return err
		}
	case *StreamResponse_UserLogout:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.UserLogout); err != nil {
			return err
		}
	case *StreamResponse_UserMessage:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.UserMessage); err != nil {
			return err
		}
	case *StreamResponse_ServerShutdown:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ServerShutdown); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("StreamResponse.Event has unexpected type %T", x)
	}
	return nil
}

func _StreamResponse_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*StreamResponse)
	switch tag {
	case 2: // event.user_login
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StreamResponse_Login)
		err := b.DecodeMessage(msg)
		m.Event = &StreamResponse_UserLogin{msg}
		return true, err
	case 3: // event.user_logout
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StreamResponse_Logout)
		err := b.DecodeMessage(msg)
		m.Event = &StreamResponse_UserLogout{msg}
		return true, err
	case 4: // event.user_message
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StreamResponse_Message)
		err := b.DecodeMessage(msg)
		m.Event = &StreamResponse_UserMessage{msg}
		return true, err
	case 5: // event.server_shutdown
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StreamResponse_Shutdown)
		err := b.DecodeMessage(msg)
		m.Event = &StreamResponse_ServerShutdown{msg}
		return true, err
	default:
		return false, nil
	}
}

func _StreamResponse_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*StreamResponse)
	// event
	switch x := m.Event.(type) {
	case *StreamResponse_UserLogin:
		s := proto.Size(x.UserLogin)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *StreamResponse_UserLogout:
		s := proto.Size(x.UserLogout)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *StreamResponse_UserMessage:
		s := proto.Size(x.UserMessage)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *StreamResponse_ServerShutdown:
		s := proto.Size(x.ServerShutdown)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type StreamResponse_Login struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamResponse_Login) Reset()         { *m = StreamResponse_Login{} }
func (m *StreamResponse_Login) String() string { return proto.CompactTextString(m) }
func (*StreamResponse_Login) ProtoMessage()    {}
func (*StreamResponse_Login) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_bf91dfeaf7880616, []int{5, 0}
}
func (m *StreamResponse_Login) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamResponse_Login.Unmarshal(m, b)
}
func (m *StreamResponse_Login) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamResponse_Login.Marshal(b, m, deterministic)
}
func (dst *StreamResponse_Login) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamResponse_Login.Merge(dst, src)
}
func (m *StreamResponse_Login) XXX_Size() int {
	return xxx_messageInfo_StreamResponse_Login.Size(m)
}
func (m *StreamResponse_Login) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamResponse_Login.DiscardUnknown(m)
}

var xxx_messageInfo_StreamResponse_Login proto.InternalMessageInfo

func (m *StreamResponse_Login) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type StreamResponse_Logout struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamResponse_Logout) Reset()         { *m = StreamResponse_Logout{} }
func (m *StreamResponse_Logout) String() string { return proto.CompactTextString(m) }
func (*StreamResponse_Logout) ProtoMessage()    {}
func (*StreamResponse_Logout) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_bf91dfeaf7880616, []int{5, 1}
}
func (m *StreamResponse_Logout) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamResponse_Logout.Unmarshal(m, b)
}
func (m *StreamResponse_Logout) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamResponse_Logout.Marshal(b, m, deterministic)
}
func (dst *StreamResponse_Logout) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamResponse_Logout.Merge(dst, src)
}
func (m *StreamResponse_Logout) XXX_Size() int {
	return xxx_messageInfo_StreamResponse_Logout.Size(m)
}
func (m *StreamResponse_Logout) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamResponse_Logout.DiscardUnknown(m)
}

var xxx_messageInfo_StreamResponse_Logout proto.InternalMessageInfo

func (m *StreamResponse_Logout) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type StreamResponse_Message struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamResponse_Message) Reset()         { *m = StreamResponse_Message{} }
func (m *StreamResponse_Message) String() string { return proto.CompactTextString(m) }
func (*StreamResponse_Message) ProtoMessage()    {}
func (*StreamResponse_Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_bf91dfeaf7880616, []int{5, 2}
}
func (m *StreamResponse_Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamResponse_Message.Unmarshal(m, b)
}
func (m *StreamResponse_Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamResponse_Message.Marshal(b, m, deterministic)
}
func (dst *StreamResponse_Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamResponse_Message.Merge(dst, src)
}
func (m *StreamResponse_Message) XXX_Size() int {
	return xxx_messageInfo_StreamResponse_Message.Size(m)
}
func (m *StreamResponse_Message) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamResponse_Message.DiscardUnknown(m)
}

var xxx_messageInfo_StreamResponse_Message proto.InternalMessageInfo

func (m *StreamResponse_Message) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *StreamResponse_Message) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type StreamResponse_Shutdown struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamResponse_Shutdown) Reset()         { *m = StreamResponse_Shutdown{} }
func (m *StreamResponse_Shutdown) String() string { return proto.CompactTextString(m) }
func (*StreamResponse_Shutdown) ProtoMessage()    {}
func (*StreamResponse_Shutdown) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_bf91dfeaf7880616, []int{5, 3}
}
func (m *StreamResponse_Shutdown) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamResponse_Shutdown.Unmarshal(m, b)
}
func (m *StreamResponse_Shutdown) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamResponse_Shutdown.Marshal(b, m, deterministic)
}
func (dst *StreamResponse_Shutdown) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamResponse_Shutdown.Merge(dst, src)
}
func (m *StreamResponse_Shutdown) XXX_Size() int {
	return xxx_messageInfo_StreamResponse_Shutdown.Size(m)
}
func (m *StreamResponse_Shutdown) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamResponse_Shutdown.DiscardUnknown(m)
}

var xxx_messageInfo_StreamResponse_Shutdown proto.InternalMessageInfo

func init() {
	proto.RegisterType((*LoginRequest)(nil), "chat.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "chat.LoginResponse")
	proto.RegisterType((*LogoutRequest)(nil), "chat.LogoutRequest")
	proto.RegisterType((*LogoutResponse)(nil), "chat.LogoutResponse")
	proto.RegisterType((*StreamRequest)(nil), "chat.StreamRequest")
	proto.RegisterType((*StreamResponse)(nil), "chat.StreamResponse")
	proto.RegisterType((*StreamResponse_Login)(nil), "chat.StreamResponse.Login")
	proto.RegisterType((*StreamResponse_Logout)(nil), "chat.StreamResponse.Logout")
	proto.RegisterType((*StreamResponse_Message)(nil), "chat.StreamResponse.Message")
	proto.RegisterType((*StreamResponse_Shutdown)(nil), "chat.StreamResponse.Shutdown")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ChatClient is the client API for Chat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error)
	Stream(ctx context.Context, opts ...grpc.CallOption) (Chat_StreamClient, error)
}

type chatClient struct {
	cc *grpc.ClientConn
}

func NewChatClient(cc *grpc.ClientConn) ChatClient {
	return &chatClient{cc}
}

func (c *chatClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/chat.Chat/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error) {
	out := new(LogoutResponse)
	err := c.cc.Invoke(ctx, "/chat.Chat/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) Stream(ctx context.Context, opts ...grpc.CallOption) (Chat_StreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Chat_serviceDesc.Streams[0], "/chat.Chat/Stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatStreamClient{stream}
	return x, nil
}

type Chat_StreamClient interface {
	Send(*StreamRequest) error
	Recv() (*StreamResponse, error)
	grpc.ClientStream
}

type chatStreamClient struct {
	grpc.ClientStream
}

func (x *chatStreamClient) Send(m *StreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatStreamClient) Recv() (*StreamResponse, error) {
	m := new(StreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatServer is the server API for Chat service.
type ChatServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Logout(context.Context, *LogoutRequest) (*LogoutResponse, error)
	Stream(Chat_StreamServer) error
}

func RegisterChatServer(s *grpc.Server, srv ChatServer) {
	s.RegisterService(&_Chat_serviceDesc, srv)
}

func _Chat_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.Chat/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.Chat/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).Logout(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServer).Stream(&chatStreamServer{stream})
}

type Chat_StreamServer interface {
	Send(*StreamResponse) error
	Recv() (*StreamRequest, error)
	grpc.ServerStream
}

type chatStreamServer struct {
	grpc.ServerStream
}

func (x *chatStreamServer) Send(m *StreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatStreamServer) Recv() (*StreamRequest, error) {
	m := new(StreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Chat_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chat.Chat",
	HandlerType: (*ChatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Chat_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _Chat_Logout_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _Chat_Stream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "pkg/chat/chat.proto",
}

func init() { proto.RegisterFile("pkg/chat/chat.proto", fileDescriptor_chat_bf91dfeaf7880616) }

var fileDescriptor_chat_bf91dfeaf7880616 = []byte{
	// 409 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x4d, 0x4f, 0xf2, 0x40,
	0x10, 0xc7, 0x5b, 0x68, 0xe1, 0x61, 0x78, 0x79, 0xcc, 0xc2, 0xa1, 0x29, 0x18, 0x4d, 0x13, 0x13,
	0xbc, 0x14, 0x83, 0x31, 0x6a, 0x4c, 0x4c, 0xd4, 0x4b, 0x0f, 0x7a, 0x29, 0xde, 0x49, 0xd1, 0xb5,
	0x10, 0x68, 0xb7, 0x76, 0x77, 0xf1, 0x4b, 0xf9, 0x95, 0xfc, 0x2e, 0x86, 0x7d, 0x29, 0x34, 0xa9,
	0x97, 0xa6, 0x33, 0xfb, 0xfb, 0xcf, 0xce, 0x7f, 0x67, 0xa0, 0x9f, 0xad, 0xe3, 0xc9, 0xdb, 0x32,
	0x62, 0xe2, 0xe3, 0x67, 0x39, 0x61, 0x04, 0x59, 0xbb, 0x7f, 0xf7, 0x24, 0x26, 0x24, 0xde, 0xe0,
	0x89, 0xc8, 0x2d, 0xf8, 0xc7, 0x84, 0xad, 0x12, 0x4c, 0x59, 0x94, 0x64, 0x12, 0xf3, 0x3c, 0xe8,
	0x3c, 0x93, 0x78, 0x95, 0x86, 0xf8, 0x93, 0x63, 0xca, 0x10, 0x02, 0x2b, 0x8d, 0x12, 0xec, 0x98,
	0xa7, 0xe6, 0xb8, 0x15, 0x8a, 0x7f, 0xef, 0x0c, 0xba, 0x8a, 0xa1, 0x19, 0x49, 0x29, 0x46, 0x03,
	0xb0, 0x19, 0x59, 0xe3, 0x54, 0x51, 0x32, 0x50, 0x18, 0xe1, 0x4c, 0xd7, 0xaa, 0xc6, 0x8e, 0xa0,
	0xa7, 0x31, 0x59, 0xce, 0x3b, 0x87, 0xee, 0x8c, 0xe5, 0x38, 0x4a, 0xb4, 0xd0, 0x81, 0x66, 0x82,
	0x29, 0x8d, 0x62, 0xdd, 0x87, 0x0e, 0xbd, 0x9f, 0x3a, 0xf4, 0x34, 0xab, 0x9a, 0xb9, 0x81, 0x56,
	0x61, 0x4a, 0xe0, 0xed, 0xa9, 0xeb, 0x4b, 0xdb, 0xbe, 0xb6, 0xed, 0xbf, 0x6a, 0x22, 0xdc, 0xc3,
	0xe8, 0x0e, 0x80, 0x53, 0x9c, 0xcf, 0x37, 0x3b, 0x73, 0x4e, 0x4d, 0x49, 0xc5, 0x1b, 0x96, 0xef,
	0xf0, 0x85, 0xfd, 0xc0, 0x08, 0x5b, 0x3b, 0x5e, 0x04, 0xe8, 0x1e, 0xda, 0x5a, 0x4c, 0x38, 0x73,
	0xea, 0x42, 0x3d, 0xfc, 0x4b, 0x4d, 0x38, 0x0b, 0x8c, 0x10, 0x94, 0x9c, 0x70, 0x86, 0x1e, 0xa0,
	0x23, 0xf4, 0xda, 0xa8, 0x25, 0x0a, 0x8c, 0x2a, 0x0b, 0xbc, 0x48, 0x26, 0x30, 0x42, 0x71, 0xa7,
	0x0a, 0x51, 0x00, 0xff, 0x29, 0xce, 0xb7, 0x38, 0x9f, 0xd3, 0x25, 0x67, 0xef, 0xe4, 0x2b, 0x75,
	0x6c, 0x51, 0xe5, 0xb8, 0xb2, 0xca, 0x4c, 0x41, 0x81, 0x11, 0xf6, 0xa4, 0x4e, 0x67, 0xdc, 0x21,
	0xd8, 0xd2, 0x55, 0xc5, 0xf8, 0xdd, 0x11, 0x34, 0x54, 0xcf, 0x55, 0xa7, 0xd7, 0xd0, 0xd4, 0xfd,
	0x54, 0x1c, 0x1f, 0x8e, 0xb2, 0x56, 0x1a, 0xa5, 0x0b, 0xf0, 0x4f, 0xdf, 0xff, 0xd8, 0x04, 0x1b,
	0x6f, 0x71, 0xca, 0xa6, 0xdf, 0x26, 0x58, 0x4f, 0xcb, 0x88, 0xa1, 0x69, 0xd1, 0x91, 0xf4, 0x72,
	0xb8, 0xa4, 0x6e, 0xbf, 0x94, 0x53, 0x5b, 0x64, 0xa0, 0xab, 0xa2, 0xd1, 0x3d, 0xb0, 0x5f, 0x47,
	0x77, 0x50, 0x4e, 0x16, 0xb2, 0x5b, 0x68, 0xc8, 0x97, 0xd2, 0xb2, 0xd2, 0x32, 0x6a, 0x59, 0xf9,
	0x31, 0x3d, 0x63, 0x6c, 0x5e, 0x98, 0x8b, 0x86, 0x58, 0xb0, 0xcb, 0xdf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xa1, 0x32, 0x83, 0xfb, 0x82, 0x03, 0x00, 0x00,
}
