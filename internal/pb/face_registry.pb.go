// Code generated by protoc-gen-go. DO NOT EDIT.
// source: face_registry.proto

package DomFaceRegistry

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type RecognizeFaceResponse struct {
	PersonInfo           *Person  `protobuf:"bytes,1,opt,name=person_info,json=personInfo,proto3" json:"person_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecognizeFaceResponse) Reset()         { *m = RecognizeFaceResponse{} }
func (m *RecognizeFaceResponse) String() string { return proto.CompactTextString(m) }
func (*RecognizeFaceResponse) ProtoMessage()    {}
func (*RecognizeFaceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_31944c70042b1972, []int{0}
}

func (m *RecognizeFaceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecognizeFaceResponse.Unmarshal(m, b)
}
func (m *RecognizeFaceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecognizeFaceResponse.Marshal(b, m, deterministic)
}
func (m *RecognizeFaceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecognizeFaceResponse.Merge(m, src)
}
func (m *RecognizeFaceResponse) XXX_Size() int {
	return xxx_messageInfo_RecognizeFaceResponse.Size(m)
}
func (m *RecognizeFaceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RecognizeFaceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RecognizeFaceResponse proto.InternalMessageInfo

func (m *RecognizeFaceResponse) GetPersonInfo() *Person {
	if m != nil {
		return m.PersonInfo
	}
	return nil
}

type AddFaceRequest struct {
	PersonId             string     `protobuf:"bytes,1,opt,name=person_id,json=personId,proto3" json:"person_id,omitempty"`
	FaceImage            *FaceImage `protobuf:"bytes,2,opt,name=face_image,json=faceImage,proto3" json:"face_image,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *AddFaceRequest) Reset()         { *m = AddFaceRequest{} }
func (m *AddFaceRequest) String() string { return proto.CompactTextString(m) }
func (*AddFaceRequest) ProtoMessage()    {}
func (*AddFaceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_31944c70042b1972, []int{1}
}

func (m *AddFaceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddFaceRequest.Unmarshal(m, b)
}
func (m *AddFaceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddFaceRequest.Marshal(b, m, deterministic)
}
func (m *AddFaceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddFaceRequest.Merge(m, src)
}
func (m *AddFaceRequest) XXX_Size() int {
	return xxx_messageInfo_AddFaceRequest.Size(m)
}
func (m *AddFaceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddFaceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddFaceRequest proto.InternalMessageInfo

func (m *AddFaceRequest) GetPersonId() string {
	if m != nil {
		return m.PersonId
	}
	return ""
}

func (m *AddFaceRequest) GetFaceImage() *FaceImage {
	if m != nil {
		return m.FaceImage
	}
	return nil
}

type AddFaceResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddFaceResponse) Reset()         { *m = AddFaceResponse{} }
func (m *AddFaceResponse) String() string { return proto.CompactTextString(m) }
func (*AddFaceResponse) ProtoMessage()    {}
func (*AddFaceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_31944c70042b1972, []int{2}
}

func (m *AddFaceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddFaceResponse.Unmarshal(m, b)
}
func (m *AddFaceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddFaceResponse.Marshal(b, m, deterministic)
}
func (m *AddFaceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddFaceResponse.Merge(m, src)
}
func (m *AddFaceResponse) XXX_Size() int {
	return xxx_messageInfo_AddFaceResponse.Size(m)
}
func (m *AddFaceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddFaceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddFaceResponse proto.InternalMessageInfo

func (m *AddFaceResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type RemoveFaceRequest struct {
	FaceId               string   `protobuf:"bytes,1,opt,name=face_id,json=faceId,proto3" json:"face_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveFaceRequest) Reset()         { *m = RemoveFaceRequest{} }
func (m *RemoveFaceRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveFaceRequest) ProtoMessage()    {}
func (*RemoveFaceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_31944c70042b1972, []int{3}
}

func (m *RemoveFaceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveFaceRequest.Unmarshal(m, b)
}
func (m *RemoveFaceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveFaceRequest.Marshal(b, m, deterministic)
}
func (m *RemoveFaceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveFaceRequest.Merge(m, src)
}
func (m *RemoveFaceRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveFaceRequest.Size(m)
}
func (m *RemoveFaceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveFaceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveFaceRequest proto.InternalMessageInfo

func (m *RemoveFaceRequest) GetFaceId() string {
	if m != nil {
		return m.FaceId
	}
	return ""
}

type RemoveFaceResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveFaceResponse) Reset()         { *m = RemoveFaceResponse{} }
func (m *RemoveFaceResponse) String() string { return proto.CompactTextString(m) }
func (*RemoveFaceResponse) ProtoMessage()    {}
func (*RemoveFaceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_31944c70042b1972, []int{4}
}

func (m *RemoveFaceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveFaceResponse.Unmarshal(m, b)
}
func (m *RemoveFaceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveFaceResponse.Marshal(b, m, deterministic)
}
func (m *RemoveFaceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveFaceResponse.Merge(m, src)
}
func (m *RemoveFaceResponse) XXX_Size() int {
	return xxx_messageInfo_RemoveFaceResponse.Size(m)
}
func (m *RemoveFaceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveFaceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveFaceResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*RecognizeFaceResponse)(nil), "DomFaceRegistry.RecognizeFaceResponse")
	proto.RegisterType((*AddFaceRequest)(nil), "DomFaceRegistry.AddFaceRequest")
	proto.RegisterType((*AddFaceResponse)(nil), "DomFaceRegistry.AddFaceResponse")
	proto.RegisterType((*RemoveFaceRequest)(nil), "DomFaceRegistry.RemoveFaceRequest")
	proto.RegisterType((*RemoveFaceResponse)(nil), "DomFaceRegistry.RemoveFaceResponse")
}

func init() { proto.RegisterFile("face_registry.proto", fileDescriptor_31944c70042b1972) }

var fileDescriptor_31944c70042b1972 = []byte{
	// 232 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x3d, 0x4f, 0xc3, 0x30,
	0x10, 0x40, 0xd5, 0x0c, 0x85, 0x5c, 0xa4, 0x56, 0x18, 0x50, 0xab, 0xb2, 0x80, 0x27, 0x06, 0x94,
	0x01, 0x16, 0x18, 0x91, 0x10, 0x52, 0x37, 0xf0, 0x1f, 0xa8, 0x42, 0x7c, 0x0e, 0x1e, 0xe2, 0x33,
	0xb6, 0x41, 0x0a, 0xbf, 0x1e, 0xc5, 0x4e, 0xa2, 0x88, 0x6e, 0xfe, 0x78, 0x7e, 0xef, 0x64, 0x38,
	0x57, 0x55, 0x8d, 0x07, 0x87, 0x8d, 0xf6, 0xc1, 0x75, 0xa5, 0x75, 0x14, 0x88, 0xad, 0x5f, 0xa8,
	0x7d, 0xad, 0x6a, 0x14, 0xc3, 0xf1, 0xae, 0x08, 0x9d, 0x45, 0x9f, 0x6e, 0xf9, 0x3b, 0x5c, 0x0a,
	0xac, 0xa9, 0x31, 0xfa, 0x17, 0x13, 0xe5, 0x2d, 0x19, 0x8f, 0xec, 0x11, 0x0a, 0x8b, 0xce, 0x93,
	0x39, 0x68, 0xa3, 0x68, 0xbb, 0xb8, 0x5e, 0xdc, 0x16, 0xf7, 0x9b, 0xf2, 0x9f, 0xac, 0x7c, 0x8b,
	0x8c, 0x80, 0xc4, 0xee, 0x8d, 0x22, 0xfe, 0x09, 0xab, 0x67, 0x29, 0x13, 0xf5, 0xf5, 0x8d, 0x3e,
	0xb0, 0x2b, 0xc8, 0x47, 0x97, 0x8c, 0xa6, 0x5c, 0x9c, 0x0e, 0x0f, 0x24, 0x7b, 0x02, 0x88, 0x63,
	0xeb, 0xb6, 0x6a, 0x70, 0x9b, 0xc5, 0xce, 0xee, 0xa8, 0xd3, 0x6f, 0xf6, 0x3d, 0x21, 0x72, 0x35,
	0x2e, 0xf9, 0x0d, 0xac, 0xa7, 0xd2, 0x30, 0xf6, 0x0a, 0xb2, 0xa9, 0x91, 0x69, 0xc9, 0xef, 0xe0,
	0x4c, 0x60, 0x4b, 0x3f, 0x38, 0x9f, 0x67, 0x03, 0x27, 0x29, 0x39, 0x92, 0xcb, 0xe8, 0x94, 0xfc,
	0x02, 0xd8, 0x9c, 0x4e, 0xce, 0x8f, 0x65, 0xfc, 0xaa, 0x87, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x6b, 0x90, 0x2c, 0xb8, 0x5f, 0x01, 0x00, 0x00,
}
