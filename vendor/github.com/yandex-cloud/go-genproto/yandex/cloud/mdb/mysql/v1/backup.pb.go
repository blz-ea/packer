// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/mdb/mysql/v1/backup.proto

package mysql

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
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

// A MySQL backup. For more information, see
// the [documentation](/docs/managed-mysql/concepts/backup).
type Backup struct {
	// ID of the backup.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the folder that the backup belongs to.
	FolderId  string               `protobuf:"bytes,2,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// ID of the MySQL cluster that the backup was created for.
	SourceClusterId string `protobuf:"bytes,4,opt,name=source_cluster_id,json=sourceClusterId,proto3" json:"source_cluster_id,omitempty"`
	// Time when the backup operation was started.
	StartedAt            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Backup) Reset()         { *m = Backup{} }
func (m *Backup) String() string { return proto.CompactTextString(m) }
func (*Backup) ProtoMessage()    {}
func (*Backup) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9951bfdbe6ddc67, []int{0}
}

func (m *Backup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Backup.Unmarshal(m, b)
}
func (m *Backup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Backup.Marshal(b, m, deterministic)
}
func (m *Backup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Backup.Merge(m, src)
}
func (m *Backup) XXX_Size() int {
	return xxx_messageInfo_Backup.Size(m)
}
func (m *Backup) XXX_DiscardUnknown() {
	xxx_messageInfo_Backup.DiscardUnknown(m)
}

var xxx_messageInfo_Backup proto.InternalMessageInfo

func (m *Backup) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Backup) GetFolderId() string {
	if m != nil {
		return m.FolderId
	}
	return ""
}

func (m *Backup) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Backup) GetSourceClusterId() string {
	if m != nil {
		return m.SourceClusterId
	}
	return ""
}

func (m *Backup) GetStartedAt() *timestamp.Timestamp {
	if m != nil {
		return m.StartedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*Backup)(nil), "yandex.cloud.mdb.mysql.v1.Backup")
}

func init() {
	proto.RegisterFile("yandex/cloud/mdb/mysql/v1/backup.proto", fileDescriptor_a9951bfdbe6ddc67)
}

var fileDescriptor_a9951bfdbe6ddc67 = []byte{
	// 294 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xb1, 0x4e, 0xeb, 0x30,
	0x14, 0x86, 0x95, 0xde, 0xde, 0x8a, 0x9a, 0x01, 0x11, 0x31, 0x84, 0xa0, 0x8a, 0x8a, 0x01, 0x55,
	0x48, 0xb5, 0x15, 0x98, 0x10, 0x53, 0xd3, 0x89, 0x35, 0x62, 0x62, 0x89, 0x6c, 0x1f, 0x37, 0x58,
	0xd8, 0x71, 0x48, 0xec, 0x88, 0x3e, 0x21, 0x4f, 0xc1, 0xce, 0x63, 0x20, 0xec, 0x74, 0xc8, 0x80,
	0xd8, 0xec, 0x73, 0x3e, 0xff, 0x9f, 0xf4, 0x1b, 0x5d, 0xef, 0x69, 0x0d, 0xe2, 0x9d, 0x70, 0x65,
	0x1c, 0x10, 0x0d, 0x8c, 0xe8, 0x7d, 0xf7, 0xa6, 0x48, 0x9f, 0x11, 0x46, 0xf9, 0xab, 0x6b, 0x70,
	0xd3, 0x1a, 0x6b, 0xe2, 0xf3, 0xc0, 0x61, 0xcf, 0x61, 0x0d, 0x0c, 0x7b, 0x0e, 0xf7, 0x59, 0xba,
	0x18, 0x45, 0xf4, 0x54, 0x49, 0xa0, 0x56, 0x9a, 0x3a, 0xbc, 0x4c, 0x2f, 0x2b, 0x63, 0x2a, 0x25,
	0x88, 0xbf, 0x31, 0xb7, 0x23, 0x56, 0x6a, 0xd1, 0x59, 0xaa, 0x87, 0xe8, 0xab, 0xcf, 0x08, 0xcd,
	0x72, 0xef, 0x8a, 0xcf, 0xd0, 0x44, 0x42, 0x12, 0x2d, 0xa3, 0xd5, 0x3c, 0x9f, 0x7e, 0x7d, 0x64,
	0x51, 0x31, 0x91, 0x10, 0x5f, 0xa0, 0xf9, 0xce, 0x28, 0x10, 0x6d, 0x29, 0x21, 0x99, 0xfc, 0x2c,
	0x8b, 0xa3, 0x30, 0x78, 0x84, 0xf8, 0x1e, 0x21, 0xde, 0x0a, 0x6a, 0x05, 0x94, 0xd4, 0x26, 0xff,
	0x96, 0xd1, 0xea, 0xf8, 0x36, 0xc5, 0xc1, 0x89, 0x0f, 0x4e, 0xfc, 0x74, 0x70, 0x16, 0xf3, 0x81,
	0xde, 0xd8, 0xf8, 0x06, 0x9d, 0x76, 0xc6, 0xb5, 0x5c, 0x94, 0x5c, 0xb9, 0xce, 0x86, 0xfc, 0xa9,
	0xcf, 0x3f, 0x09, 0x8b, 0x6d, 0x98, 0x07, 0x4d, 0x67, 0x69, 0x3b, 0x68, 0xfe, 0xff, 0xad, 0x19,
	0xe8, 0x8d, 0xcd, 0x01, 0x2d, 0x46, 0xe5, 0xd1, 0x46, 0x8e, 0x0a, 0x7c, 0xde, 0x56, 0xd2, 0xbe,
	0x38, 0x86, 0xb9, 0xd1, 0x24, 0x90, 0xeb, 0xd0, 0x65, 0x65, 0xd6, 0x95, 0xa8, 0x7d, 0x3a, 0xf9,
	0xf5, 0x9f, 0x1e, 0xfc, 0x81, 0xcd, 0x3c, 0x76, 0xf7, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x8e, 0xfb,
	0xfd, 0x7a, 0xd1, 0x01, 0x00, 0x00,
}
