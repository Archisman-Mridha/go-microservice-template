// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: chat_service.proto

package generated

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Message  string                 `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	SenderID int32                  `protobuf:"varint,4,opt,name=senderID,proto3" json:"senderID,omitempty"`
	SentAt   *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=sentAt,proto3" json:"sentAt,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_chat_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_chat_service_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Message) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Message) GetSenderID() int32 {
	if x != nil {
		return x.SenderID
	}
	return 0
}

func (x *Message) GetSentAt() *timestamppb.Timestamp {
	if x != nil {
		return x.SentAt
	}
	return nil
}

type Chat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	WithUserID  int32    `protobuf:"varint,2,opt,name=withUserID,proto3" json:"withUserID,omitempty"`
	LastMessage *Message `protobuf:"bytes,3,opt,name=lastMessage,proto3" json:"lastMessage,omitempty"`
}

func (x *Chat) Reset() {
	*x = Chat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chat) ProtoMessage() {}

func (x *Chat) ProtoReflect() protoreflect.Message {
	mi := &file_chat_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chat.ProtoReflect.Descriptor instead.
func (*Chat) Descriptor() ([]byte, []int) {
	return file_chat_service_proto_rawDescGZIP(), []int{1}
}

func (x *Chat) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Chat) GetWithUserID() int32 {
	if x != nil {
		return x.WithUserID
	}
	return 0
}

func (x *Chat) GetLastMessage() *Message {
	if x != nil {
		return x.LastMessage
	}
	return nil
}

type PaginationOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int32 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  int32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *PaginationOptions) Reset() {
	*x = PaginationOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaginationOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginationOptions) ProtoMessage() {}

func (x *PaginationOptions) ProtoReflect() protoreflect.Message {
	mi := &file_chat_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginationOptions.ProtoReflect.Descriptor instead.
func (*PaginationOptions) Descriptor() ([]byte, []int) {
	return file_chat_service_proto_rawDescGZIP(), []int{2}
}

func (x *PaginationOptions) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *PaginationOptions) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetChatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// userID is extracted from the JWT (at the gateway).
	UserID            int32              `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	PaginationOptions *PaginationOptions `protobuf:"bytes,2,opt,name=paginationOptions,proto3" json:"paginationOptions,omitempty"`
}

func (x *GetChatsRequest) Reset() {
	*x = GetChatsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChatsRequest) ProtoMessage() {}

func (x *GetChatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChatsRequest.ProtoReflect.Descriptor instead.
func (*GetChatsRequest) Descriptor() ([]byte, []int) {
	return file_chat_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetChatsRequest) GetUserID() int32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *GetChatsRequest) GetPaginationOptions() *PaginationOptions {
	if x != nil {
		return x.PaginationOptions
	}
	return nil
}

type GetChatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chats []*Chat `protobuf:"bytes,1,rep,name=chats,proto3" json:"chats,omitempty"`
}

func (x *GetChatsResponse) Reset() {
	*x = GetChatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChatsResponse) ProtoMessage() {}

func (x *GetChatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chat_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChatsResponse.ProtoReflect.Descriptor instead.
func (*GetChatsResponse) Descriptor() ([]byte, []int) {
	return file_chat_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetChatsResponse) GetChats() []*Chat {
	if x != nil {
		return x.Chats
	}
	return nil
}

type CreateChatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserIDs []int32 `protobuf:"varint,1,rep,packed,name=userIDs,proto3" json:"userIDs,omitempty"`
}

func (x *CreateChatRequest) Reset() {
	*x = CreateChatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateChatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateChatRequest) ProtoMessage() {}

func (x *CreateChatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateChatRequest.ProtoReflect.Descriptor instead.
func (*CreateChatRequest) Descriptor() ([]byte, []int) {
	return file_chat_service_proto_rawDescGZIP(), []int{5}
}

func (x *CreateChatRequest) GetUserIDs() []int32 {
	if x != nil {
		return x.UserIDs
	}
	return nil
}

type CreateChatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatID int32 `protobuf:"varint,1,opt,name=chatID,proto3" json:"chatID,omitempty"`
}

func (x *CreateChatResponse) Reset() {
	*x = CreateChatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateChatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateChatResponse) ProtoMessage() {}

func (x *CreateChatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chat_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateChatResponse.ProtoReflect.Descriptor instead.
func (*CreateChatResponse) Descriptor() ([]byte, []int) {
	return file_chat_service_proto_rawDescGZIP(), []int{6}
}

func (x *CreateChatResponse) GetChatID() int32 {
	if x != nil {
		return x.ChatID
	}
	return 0
}

type GetChatMessagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatId            int32              `protobuf:"varint,1,opt,name=chatId,proto3" json:"chatId,omitempty"`
	PaginationOptions *PaginationOptions `protobuf:"bytes,2,opt,name=paginationOptions,proto3" json:"paginationOptions,omitempty"`
}

func (x *GetChatMessagesRequest) Reset() {
	*x = GetChatMessagesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChatMessagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChatMessagesRequest) ProtoMessage() {}

func (x *GetChatMessagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChatMessagesRequest.ProtoReflect.Descriptor instead.
func (*GetChatMessagesRequest) Descriptor() ([]byte, []int) {
	return file_chat_service_proto_rawDescGZIP(), []int{7}
}

func (x *GetChatMessagesRequest) GetChatId() int32 {
	if x != nil {
		return x.ChatId
	}
	return 0
}

func (x *GetChatMessagesRequest) GetPaginationOptions() *PaginationOptions {
	if x != nil {
		return x.PaginationOptions
	}
	return nil
}

type GetChatMessagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message []*Message `protobuf:"bytes,1,rep,name=message,proto3" json:"message,omitempty"`
}

func (x *GetChatMessagesResponse) Reset() {
	*x = GetChatMessagesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChatMessagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChatMessagesResponse) ProtoMessage() {}

func (x *GetChatMessagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chat_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChatMessagesResponse.ProtoReflect.Descriptor instead.
func (*GetChatMessagesResponse) Descriptor() ([]byte, []int) {
	return file_chat_service_proto_rawDescGZIP(), []int{8}
}

func (x *GetChatMessagesResponse) GetMessage() []*Message {
	if x != nil {
		return x.Message
	}
	return nil
}

type NewMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// userID is extracted from the JWT (at the gateway).
	UserID  int32    `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	ChatID  int32    `protobuf:"varint,2,opt,name=chatID,proto3" json:"chatID,omitempty"`
	Message *Message `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *NewMessage) Reset() {
	*x = NewMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewMessage) ProtoMessage() {}

func (x *NewMessage) ProtoReflect() protoreflect.Message {
	mi := &file_chat_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewMessage.ProtoReflect.Descriptor instead.
func (*NewMessage) Descriptor() ([]byte, []int) {
	return file_chat_service_proto_rawDescGZIP(), []int{9}
}

func (x *NewMessage) GetUserID() int32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *NewMessage) GetChatID() int32 {
	if x != nil {
		return x.ChatID
	}
	return 0
}

func (x *NewMessage) GetMessage() *Message {
	if x != nil {
		return x.Message
	}
	return nil
}

var File_chat_service_proto protoreflect.FileDescriptor

var file_chat_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x83, 0x01, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x49, 0x44, 0x12, 0x32, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x06,
	0x73, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x22, 0x6f, 0x0a, 0x04, 0x43, 0x68, 0x61, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e,
	0x0a, 0x0a, 0x77, 0x69, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x77, 0x69, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x37,
	0x0a, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x41, 0x0a, 0x11, 0x50, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x78, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x43, 0x68, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x4d, 0x0a, 0x11, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x11, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x22, 0x3c, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x05, 0x63, 0x68, 0x61, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x05, 0x63, 0x68, 0x61,
	0x74, 0x73, 0x22, 0x2d, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x73, 0x22, 0x2c, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x44, 0x22,
	0x7f, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x68, 0x61,
	0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49,
	0x64, 0x12, 0x4d, 0x0a, 0x11, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63,
	0x68, 0x61, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x11, 0x70,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x22, 0x4a, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63,
	0x68, 0x61, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x6d, 0x0a, 0x0a,
	0x4e, 0x65, 0x77, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x44, 0x12, 0x2f, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x68,
	0x61, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x81, 0x03, 0x0a, 0x0b,
	0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x04, 0x50,
	0x69, 0x6e, 0x67, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x12, 0x49, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x74, 0x73, 0x12,
	0x1d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47,
	0x65, 0x74, 0x43, 0x68, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x43, 0x68, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f,
	0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x12, 0x1f, 0x2e, 0x63,
	0x68, 0x61, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e,
	0x63, 0x68, 0x61, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x5e, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x12, 0x24, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3e, 0x0a, 0x04, 0x43, 0x68, 0x61, 0x74, 0x12, 0x18, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4e, 0x65, 0x77, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x1a, 0x18, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x4e, 0x65, 0x77, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42,
	0x0c, 0x5a, 0x0a, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chat_service_proto_rawDescOnce sync.Once
	file_chat_service_proto_rawDescData = file_chat_service_proto_rawDesc
)

func file_chat_service_proto_rawDescGZIP() []byte {
	file_chat_service_proto_rawDescOnce.Do(func() {
		file_chat_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_chat_service_proto_rawDescData)
	})
	return file_chat_service_proto_rawDescData
}

var file_chat_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_chat_service_proto_goTypes = []any{
	(*Message)(nil),                 // 0: chat_service.Message
	(*Chat)(nil),                    // 1: chat_service.Chat
	(*PaginationOptions)(nil),       // 2: chat_service.PaginationOptions
	(*GetChatsRequest)(nil),         // 3: chat_service.GetChatsRequest
	(*GetChatsResponse)(nil),        // 4: chat_service.GetChatsResponse
	(*CreateChatRequest)(nil),       // 5: chat_service.CreateChatRequest
	(*CreateChatResponse)(nil),      // 6: chat_service.CreateChatResponse
	(*GetChatMessagesRequest)(nil),  // 7: chat_service.GetChatMessagesRequest
	(*GetChatMessagesResponse)(nil), // 8: chat_service.GetChatMessagesResponse
	(*NewMessage)(nil),              // 9: chat_service.NewMessage
	(*timestamppb.Timestamp)(nil),   // 10: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),           // 11: google.protobuf.Empty
}
var file_chat_service_proto_depIdxs = []int32{
	10, // 0: chat_service.Message.sentAt:type_name -> google.protobuf.Timestamp
	0,  // 1: chat_service.Chat.lastMessage:type_name -> chat_service.Message
	2,  // 2: chat_service.GetChatsRequest.paginationOptions:type_name -> chat_service.PaginationOptions
	1,  // 3: chat_service.GetChatsResponse.chats:type_name -> chat_service.Chat
	2,  // 4: chat_service.GetChatMessagesRequest.paginationOptions:type_name -> chat_service.PaginationOptions
	0,  // 5: chat_service.GetChatMessagesResponse.message:type_name -> chat_service.Message
	0,  // 6: chat_service.NewMessage.message:type_name -> chat_service.Message
	11, // 7: chat_service.ChatService.Ping:input_type -> google.protobuf.Empty
	3,  // 8: chat_service.ChatService.GetChats:input_type -> chat_service.GetChatsRequest
	5,  // 9: chat_service.ChatService.CreateChat:input_type -> chat_service.CreateChatRequest
	7,  // 10: chat_service.ChatService.GetChatMessages:input_type -> chat_service.GetChatMessagesRequest
	9,  // 11: chat_service.ChatService.Chat:input_type -> chat_service.NewMessage
	11, // 12: chat_service.ChatService.Ping:output_type -> google.protobuf.Empty
	4,  // 13: chat_service.ChatService.GetChats:output_type -> chat_service.GetChatsResponse
	6,  // 14: chat_service.ChatService.CreateChat:output_type -> chat_service.CreateChatResponse
	8,  // 15: chat_service.ChatService.GetChatMessages:output_type -> chat_service.GetChatMessagesResponse
	9,  // 16: chat_service.ChatService.Chat:output_type -> chat_service.NewMessage
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_chat_service_proto_init() }
func file_chat_service_proto_init() {
	if File_chat_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chat_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Message); i {
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
		file_chat_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Chat); i {
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
		file_chat_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*PaginationOptions); i {
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
		file_chat_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetChatsRequest); i {
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
		file_chat_service_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetChatsResponse); i {
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
		file_chat_service_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*CreateChatRequest); i {
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
		file_chat_service_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*CreateChatResponse); i {
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
		file_chat_service_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*GetChatMessagesRequest); i {
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
		file_chat_service_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*GetChatMessagesResponse); i {
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
		file_chat_service_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*NewMessage); i {
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
			RawDescriptor: file_chat_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chat_service_proto_goTypes,
		DependencyIndexes: file_chat_service_proto_depIdxs,
		MessageInfos:      file_chat_service_proto_msgTypes,
	}.Build()
	File_chat_service_proto = out.File
	file_chat_service_proto_rawDesc = nil
	file_chat_service_proto_goTypes = nil
	file_chat_service_proto_depIdxs = nil
}