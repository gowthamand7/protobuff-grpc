// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: ytservice/video.proto

package ytservice

import (
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

type Video struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string              `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string              `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string              `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	User        *User               `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`
	Comments    map[string]*Comment `protobuf:"bytes,5,rep,name=comments,proto3" json:"comments,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Video) Reset() {
	*x = Video{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ytservice_video_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Video) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Video) ProtoMessage() {}

func (x *Video) ProtoReflect() protoreflect.Message {
	mi := &file_ytservice_video_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Video.ProtoReflect.Descriptor instead.
func (*Video) Descriptor() ([]byte, []int) {
	return file_ytservice_video_proto_rawDescGZIP(), []int{0}
}

func (x *Video) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Video) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Video) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Video) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Video) GetComments() map[string]*Comment {
	if x != nil {
		return x.Comments
	}
	return nil
}

type Videos struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Videos map[string]*Video `protobuf:"bytes,1,rep,name=videos,proto3" json:"videos,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Videos) Reset() {
	*x = Videos{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ytservice_video_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Videos) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Videos) ProtoMessage() {}

func (x *Videos) ProtoReflect() protoreflect.Message {
	mi := &file_ytservice_video_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Videos.ProtoReflect.Descriptor instead.
func (*Videos) Descriptor() ([]byte, []int) {
	return file_ytservice_video_proto_rawDescGZIP(), []int{1}
}

func (x *Videos) GetVideos() map[string]*Video {
	if x != nil {
		return x.Videos
	}
	return nil
}

type VideoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *VideoRequest) Reset() {
	*x = VideoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ytservice_video_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoRequest) ProtoMessage() {}

func (x *VideoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ytservice_video_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoRequest.ProtoReflect.Descriptor instead.
func (*VideoRequest) Descriptor() ([]byte, []int) {
	return file_ytservice_video_proto_rawDescGZIP(), []int{2}
}

func (x *VideoRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CreateVideoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	UserId      string `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *CreateVideoRequest) Reset() {
	*x = CreateVideoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ytservice_video_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateVideoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateVideoRequest) ProtoMessage() {}

func (x *CreateVideoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ytservice_video_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateVideoRequest.ProtoReflect.Descriptor instead.
func (*CreateVideoRequest) Descriptor() ([]byte, []int) {
	return file_ytservice_video_proto_rawDescGZIP(), []int{3}
}

func (x *CreateVideoRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateVideoRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateVideoRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type UpdateVideoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *UpdateVideoRequest) Reset() {
	*x = UpdateVideoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ytservice_video_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateVideoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateVideoRequest) ProtoMessage() {}

func (x *UpdateVideoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ytservice_video_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateVideoRequest.ProtoReflect.Descriptor instead.
func (*UpdateVideoRequest) Descriptor() ([]byte, []int) {
	return file_ytservice_video_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateVideoRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateVideoRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateVideoRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type ListVideoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *ListVideoRequest) Reset() {
	*x = ListVideoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ytservice_video_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListVideoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListVideoRequest) ProtoMessage() {}

func (x *ListVideoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ytservice_video_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListVideoRequest.ProtoReflect.Descriptor instead.
func (*ListVideoRequest) Descriptor() ([]byte, []int) {
	return file_ytservice_video_proto_rawDescGZIP(), []int{5}
}

func (x *ListVideoRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

var File_ytservice_video_proto protoreflect.FileDescriptor

var file_ytservice_video_proto_rawDesc = []byte{
	0x0a, 0x15, 0x79, 0x74, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x79, 0x74, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x14, 0x79, 0x74, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe3, 0x01, 0x0a, 0x05, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x45, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1e, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x78, 0x0a, 0x06,
	0x56, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x12, 0x2b, 0x0a, 0x06, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x2e,
	0x56, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x73, 0x1a, 0x41, 0x0a, 0x0b, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x1e, 0x0a, 0x0c, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x65, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x5c, 0x0a,
	0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x2b, 0x0a, 0x10, 0x4c,
	0x69, 0x73, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x32, 0xe2, 0x01, 0x0a, 0x0c, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x08, 0x47, 0x65, 0x74,
	0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x0d, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x06, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x22, 0x00, 0x12, 0x2c,
	0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x13, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x06, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x0b,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x13, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x06, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x22, 0x00, 0x12, 0x26, 0x0a, 0x0b, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x0d, 0x2e, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x06, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x22, 0x00, 0x12, 0x29, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12,
	0x11, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x07, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x22, 0x00, 0x42, 0x30, 0x5a,
	0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x77, 0x74,
	0x68, 0x61, 0x6d, 0x61, 0x6e, 0x64, 0x37, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x79, 0x74, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ytservice_video_proto_rawDescOnce sync.Once
	file_ytservice_video_proto_rawDescData = file_ytservice_video_proto_rawDesc
)

func file_ytservice_video_proto_rawDescGZIP() []byte {
	file_ytservice_video_proto_rawDescOnce.Do(func() {
		file_ytservice_video_proto_rawDescData = protoimpl.X.CompressGZIP(file_ytservice_video_proto_rawDescData)
	})
	return file_ytservice_video_proto_rawDescData
}

var file_ytservice_video_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_ytservice_video_proto_goTypes = []interface{}{
	(*Video)(nil),              // 0: Video
	(*Videos)(nil),             // 1: Videos
	(*VideoRequest)(nil),       // 2: VideoRequest
	(*CreateVideoRequest)(nil), // 3: CreateVideoRequest
	(*UpdateVideoRequest)(nil), // 4: UpdateVideoRequest
	(*ListVideoRequest)(nil),   // 5: ListVideoRequest
	nil,                        // 6: Video.CommentsEntry
	nil,                        // 7: Videos.VideosEntry
	(*User)(nil),               // 8: User
	(*Comment)(nil),            // 9: Comment
}
var file_ytservice_video_proto_depIdxs = []int32{
	8,  // 0: Video.user:type_name -> User
	6,  // 1: Video.comments:type_name -> Video.CommentsEntry
	7,  // 2: Videos.videos:type_name -> Videos.VideosEntry
	9,  // 3: Video.CommentsEntry.value:type_name -> Comment
	0,  // 4: Videos.VideosEntry.value:type_name -> Video
	2,  // 5: VideoService.GetVideo:input_type -> VideoRequest
	3,  // 6: VideoService.CreateVideo:input_type -> CreateVideoRequest
	4,  // 7: VideoService.UpdateVideo:input_type -> UpdateVideoRequest
	2,  // 8: VideoService.DeleteVideo:input_type -> VideoRequest
	5,  // 9: VideoService.ListVideo:input_type -> ListVideoRequest
	0,  // 10: VideoService.GetVideo:output_type -> Video
	0,  // 11: VideoService.CreateVideo:output_type -> Video
	0,  // 12: VideoService.UpdateVideo:output_type -> Video
	0,  // 13: VideoService.DeleteVideo:output_type -> Video
	1,  // 14: VideoService.ListVideo:output_type -> Videos
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_ytservice_video_proto_init() }
func file_ytservice_video_proto_init() {
	if File_ytservice_video_proto != nil {
		return
	}
	file_ytservice_comment_proto_init()
	file_ytservice_user_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ytservice_video_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Video); i {
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
		file_ytservice_video_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Videos); i {
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
		file_ytservice_video_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoRequest); i {
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
		file_ytservice_video_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateVideoRequest); i {
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
		file_ytservice_video_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateVideoRequest); i {
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
		file_ytservice_video_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListVideoRequest); i {
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
			RawDescriptor: file_ytservice_video_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ytservice_video_proto_goTypes,
		DependencyIndexes: file_ytservice_video_proto_depIdxs,
		MessageInfos:      file_ytservice_video_proto_msgTypes,
	}.Build()
	File_ytservice_video_proto = out.File
	file_ytservice_video_proto_rawDesc = nil
	file_ytservice_video_proto_goTypes = nil
	file_ytservice_video_proto_depIdxs = nil
}
