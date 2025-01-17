// Code generated by protoc-gen-go. DO NOT EDIT.
// source: goshare/user.proto

package goshare

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

// enum 用户类型
type UserType int32

const (
	UserType_UT_NORMAL       UserType = 0
	UserType_UT_BRANCH_ADMIN UserType = 1
	UserType_UT_SYSTEM_ADMIN UserType = 2
)

var UserType_name = map[int32]string{
	0: "UT_NORMAL",
	1: "UT_BRANCH_ADMIN",
	2: "UT_SYSTEM_ADMIN",
}

var UserType_value = map[string]int32{
	"UT_NORMAL":       0,
	"UT_BRANCH_ADMIN": 1,
	"UT_SYSTEM_ADMIN": 2,
}

func (x UserType) String() string {
	return proto.EnumName(UserType_name, int32(x))
}

func (UserType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d95acd2a4cac0d87, []int{0}
}

// UserStatus 状态
type UserStatus int32

const (
	UserStatus_US_NORMAL  UserStatus = 0
	UserStatus_US_FROZEN  UserStatus = 1
	UserStatus_US_DELETED UserStatus = 2
)

var UserStatus_name = map[int32]string{
	0: "US_NORMAL",
	1: "US_FROZEN",
	2: "US_DELETED",
}

var UserStatus_value = map[string]int32{
	"US_NORMAL":  0,
	"US_FROZEN":  1,
	"US_DELETED": 2,
}

func (x UserStatus) String() string {
	return proto.EnumName(UserStatus_name, int32(x))
}

func (UserStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d95acd2a4cac0d87, []int{1}
}

// 用户
type User struct {
	Id                   string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	Name                 string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	Type                 UserType   `protobuf:"varint,3,opt,name=type,proto3,enum=goshare.UserType" json:"type"`
	Status               UserStatus `protobuf:"varint,4,opt,name=status,proto3,enum=goshare.UserStatus" json:"status"`
	Roles                []int64    `protobuf:"varint,5,rep,packed,name=roles,proto3" json:"roles"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_d95acd2a4cac0d87, []int{0}
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

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetType() UserType {
	if m != nil {
		return m.Type
	}
	return UserType_UT_NORMAL
}

func (m *User) GetStatus() UserStatus {
	if m != nil {
		return m.Status
	}
	return UserStatus_US_NORMAL
}

func (m *User) GetRoles() []int64 {
	if m != nil {
		return m.Roles
	}
	return nil
}

type UserList struct {
	List                 []*User  `protobuf:"bytes,1,rep,name=list,proto3" json:"list"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserList) Reset()         { *m = UserList{} }
func (m *UserList) String() string { return proto.CompactTextString(m) }
func (*UserList) ProtoMessage()    {}
func (*UserList) Descriptor() ([]byte, []int) {
	return fileDescriptor_d95acd2a4cac0d87, []int{1}
}

func (m *UserList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserList.Unmarshal(m, b)
}
func (m *UserList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserList.Marshal(b, m, deterministic)
}
func (m *UserList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserList.Merge(m, src)
}
func (m *UserList) XXX_Size() int {
	return xxx_messageInfo_UserList.Size(m)
}
func (m *UserList) XXX_DiscardUnknown() {
	xxx_messageInfo_UserList.DiscardUnknown(m)
}

var xxx_messageInfo_UserList proto.InternalMessageInfo

func (m *UserList) GetList() []*User {
	if m != nil {
		return m.List
	}
	return nil
}

// Permission 权限
type Permission struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	Menus                []string `protobuf:"bytes,3,rep,name=menus,proto3" json:"menus"`
	Apis                 []string `protobuf:"bytes,4,rep,name=apis,proto3" json:"apis"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Permission) Reset()         { *m = Permission{} }
func (m *Permission) String() string { return proto.CompactTextString(m) }
func (*Permission) ProtoMessage()    {}
func (*Permission) Descriptor() ([]byte, []int) {
	return fileDescriptor_d95acd2a4cac0d87, []int{2}
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

func (m *Permission) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Permission) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Permission) GetMenus() []string {
	if m != nil {
		return m.Menus
	}
	return nil
}

func (m *Permission) GetApis() []string {
	if m != nil {
		return m.Apis
	}
	return nil
}

// PermissionTreeNode 权限树节点
type PermissionTreeNode struct {
	Permission           *Permission           `protobuf:"bytes,1,opt,name=permission,proto3" json:"permission"`
	Selected             bool                  `protobuf:"varint,2,opt,name=selected,proto3" json:"selected"`
	Children             []*PermissionTreeNode `protobuf:"bytes,3,rep,name=children,proto3" json:"children"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *PermissionTreeNode) Reset()         { *m = PermissionTreeNode{} }
func (m *PermissionTreeNode) String() string { return proto.CompactTextString(m) }
func (*PermissionTreeNode) ProtoMessage()    {}
func (*PermissionTreeNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_d95acd2a4cac0d87, []int{3}
}

func (m *PermissionTreeNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PermissionTreeNode.Unmarshal(m, b)
}
func (m *PermissionTreeNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PermissionTreeNode.Marshal(b, m, deterministic)
}
func (m *PermissionTreeNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PermissionTreeNode.Merge(m, src)
}
func (m *PermissionTreeNode) XXX_Size() int {
	return xxx_messageInfo_PermissionTreeNode.Size(m)
}
func (m *PermissionTreeNode) XXX_DiscardUnknown() {
	xxx_messageInfo_PermissionTreeNode.DiscardUnknown(m)
}

var xxx_messageInfo_PermissionTreeNode proto.InternalMessageInfo

func (m *PermissionTreeNode) GetPermission() *Permission {
	if m != nil {
		return m.Permission
	}
	return nil
}

func (m *PermissionTreeNode) GetSelected() bool {
	if m != nil {
		return m.Selected
	}
	return false
}

func (m *PermissionTreeNode) GetChildren() []*PermissionTreeNode {
	if m != nil {
		return m.Children
	}
	return nil
}

// PermissionList 权限列表
type PermissionList struct {
	List                 []*Permission `protobuf:"bytes,1,rep,name=list,proto3" json:"list"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *PermissionList) Reset()         { *m = PermissionList{} }
func (m *PermissionList) String() string { return proto.CompactTextString(m) }
func (*PermissionList) ProtoMessage()    {}
func (*PermissionList) Descriptor() ([]byte, []int) {
	return fileDescriptor_d95acd2a4cac0d87, []int{4}
}

func (m *PermissionList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PermissionList.Unmarshal(m, b)
}
func (m *PermissionList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PermissionList.Marshal(b, m, deterministic)
}
func (m *PermissionList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PermissionList.Merge(m, src)
}
func (m *PermissionList) XXX_Size() int {
	return xxx_messageInfo_PermissionList.Size(m)
}
func (m *PermissionList) XXX_DiscardUnknown() {
	xxx_messageInfo_PermissionList.DiscardUnknown(m)
}

var xxx_messageInfo_PermissionList proto.InternalMessageInfo

func (m *PermissionList) GetList() []*Permission {
	if m != nil {
		return m.List
	}
	return nil
}

// 角色
type UserRole struct {
	Id                   int64         `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Name                 string        `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	Discription          string        `protobuf:"bytes,3,opt,name=discription,proto3" json:"discription"`
	Permissions          []*Permission `protobuf:"bytes,4,rep,name=permissions,proto3" json:"permissions"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *UserRole) Reset()         { *m = UserRole{} }
func (m *UserRole) String() string { return proto.CompactTextString(m) }
func (*UserRole) ProtoMessage()    {}
func (*UserRole) Descriptor() ([]byte, []int) {
	return fileDescriptor_d95acd2a4cac0d87, []int{5}
}

func (m *UserRole) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRole.Unmarshal(m, b)
}
func (m *UserRole) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRole.Marshal(b, m, deterministic)
}
func (m *UserRole) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRole.Merge(m, src)
}
func (m *UserRole) XXX_Size() int {
	return xxx_messageInfo_UserRole.Size(m)
}
func (m *UserRole) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRole.DiscardUnknown(m)
}

var xxx_messageInfo_UserRole proto.InternalMessageInfo

func (m *UserRole) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserRole) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserRole) GetDiscription() string {
	if m != nil {
		return m.Discription
	}
	return ""
}

func (m *UserRole) GetPermissions() []*Permission {
	if m != nil {
		return m.Permissions
	}
	return nil
}

// 用户会话
type Session struct {
	Id                   string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	Token                string     `protobuf:"bytes,2,opt,name=token,proto3" json:"token"`
	Type                 UserType   `protobuf:"varint,3,opt,name=type,proto3,enum=goshare.UserType" json:"type"`
	Status               UserStatus `protobuf:"varint,4,opt,name=status,proto3,enum=goshare.UserStatus" json:"status"`
	Roles                []int64    `protobuf:"varint,5,rep,packed,name=roles,proto3" json:"roles"`
	Deadline             int64      `protobuf:"varint,6,opt,name=deadline,proto3" json:"deadline"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Session) Reset()         { *m = Session{} }
func (m *Session) String() string { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()    {}
func (*Session) Descriptor() ([]byte, []int) {
	return fileDescriptor_d95acd2a4cac0d87, []int{6}
}

func (m *Session) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Session.Unmarshal(m, b)
}
func (m *Session) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Session.Marshal(b, m, deterministic)
}
func (m *Session) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Session.Merge(m, src)
}
func (m *Session) XXX_Size() int {
	return xxx_messageInfo_Session.Size(m)
}
func (m *Session) XXX_DiscardUnknown() {
	xxx_messageInfo_Session.DiscardUnknown(m)
}

var xxx_messageInfo_Session proto.InternalMessageInfo

func (m *Session) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Session) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Session) GetType() UserType {
	if m != nil {
		return m.Type
	}
	return UserType_UT_NORMAL
}

func (m *Session) GetStatus() UserStatus {
	if m != nil {
		return m.Status
	}
	return UserStatus_US_NORMAL
}

func (m *Session) GetRoles() []int64 {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *Session) GetDeadline() int64 {
	if m != nil {
		return m.Deadline
	}
	return 0
}

func init() {
	proto.RegisterEnum("goshare.UserType", UserType_name, UserType_value)
	proto.RegisterEnum("goshare.UserStatus", UserStatus_name, UserStatus_value)
	proto.RegisterType((*User)(nil), "goshare.User")
	proto.RegisterType((*UserList)(nil), "goshare.UserList")
	proto.RegisterType((*Permission)(nil), "goshare.Permission")
	proto.RegisterType((*PermissionTreeNode)(nil), "goshare.PermissionTreeNode")
	proto.RegisterType((*PermissionList)(nil), "goshare.PermissionList")
	proto.RegisterType((*UserRole)(nil), "goshare.UserRole")
	proto.RegisterType((*Session)(nil), "goshare.Session")
}

func init() { proto.RegisterFile("goshare/user.proto", fileDescriptor_d95acd2a4cac0d87) }

var fileDescriptor_d95acd2a4cac0d87 = []byte{
	// 506 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0x4f, 0x8f, 0xd2, 0x40,
	0x1c, 0xb5, 0x14, 0x58, 0xf8, 0x91, 0x45, 0x9c, 0xe5, 0xd0, 0xe8, 0xa5, 0x36, 0x31, 0x92, 0x55,
	0x4b, 0xc2, 0xc6, 0x18, 0xbd, 0xb1, 0x4b, 0x8d, 0x26, 0xc0, 0x9a, 0x69, 0x39, 0xc8, 0x85, 0x14,
	0xfa, 0x0b, 0x4c, 0xb6, 0xff, 0x32, 0x33, 0x1c, 0xf6, 0x13, 0xf8, 0x05, 0x3c, 0xfb, 0x2d, 0xfc,
	0x7e, 0xa6, 0xd3, 0x42, 0x8b, 0xd9, 0x4d, 0x3c, 0x79, 0xeb, 0x7b, 0xbf, 0xc7, 0x9b, 0xf7, 0x7b,
	0x1d, 0x0a, 0x64, 0x9b, 0x88, 0x9d, 0xcf, 0x71, 0xb8, 0x17, 0xc8, 0xed, 0x94, 0x27, 0x32, 0x21,
	0x67, 0x05, 0x67, 0xfd, 0xd4, 0xa0, 0xbe, 0x10, 0xc8, 0x49, 0x17, 0x6a, 0x2c, 0x30, 0x34, 0x53,
	0x1b, 0xb4, 0x69, 0x8d, 0x05, 0x84, 0x40, 0x3d, 0xf6, 0x23, 0x34, 0x6a, 0x8a, 0x51, 0xcf, 0xe4,
	0x15, 0xd4, 0xe5, 0x7d, 0x8a, 0x86, 0x6e, 0x6a, 0x83, 0xee, 0xe8, 0x99, 0x5d, 0x98, 0xd8, 0x99,
	0x81, 0x77, 0x9f, 0x22, 0x55, 0x63, 0xf2, 0x06, 0x9a, 0x42, 0xfa, 0x72, 0x2f, 0x8c, 0xba, 0x12,
	0x5e, 0x9c, 0x08, 0x5d, 0x35, 0xa2, 0x85, 0x84, 0xf4, 0xa1, 0xc1, 0x93, 0x10, 0x85, 0xd1, 0x30,
	0xf5, 0x81, 0x4e, 0x73, 0x60, 0xbd, 0x83, 0x56, 0xa6, 0x9d, 0x32, 0x21, 0xc9, 0x4b, 0xa8, 0x87,
	0x4c, 0x48, 0x43, 0x33, 0xf5, 0x41, 0x67, 0x74, 0x7e, 0x62, 0x46, 0xd5, 0xc8, 0x5a, 0x02, 0x7c,
	0x43, 0x1e, 0x31, 0x21, 0x58, 0x12, 0x57, 0x56, 0xd1, 0x1f, 0x5d, 0xa5, 0x0f, 0x8d, 0x08, 0xe3,
	0xbd, 0x30, 0x74, 0x53, 0x1f, 0xb4, 0x69, 0x0e, 0x32, 0xa5, 0x9f, 0xb2, 0x2c, 0x77, 0x46, 0xaa,
	0x67, 0xeb, 0x97, 0x06, 0xa4, 0x34, 0xf7, 0x38, 0xe2, 0x3c, 0x09, 0x90, 0x5c, 0x01, 0xa4, 0x47,
	0x56, 0x1d, 0xd6, 0xa9, 0x2c, 0x5a, 0xfe, 0x80, 0x56, 0x64, 0xe4, 0x39, 0xb4, 0x04, 0x86, 0xb8,
	0x91, 0x18, 0xa8, 0x34, 0x2d, 0x7a, 0xc4, 0xe4, 0x03, 0xb4, 0x36, 0x3b, 0x16, 0x06, 0x1c, 0x63,
	0x15, 0xaa, 0x33, 0x7a, 0xf1, 0x80, 0xdd, 0xe1, 0x7c, 0x7a, 0x14, 0x5b, 0x1f, 0xa1, 0x5b, 0xce,
	0x55, 0x63, 0xaf, 0x4f, 0x1a, 0x7b, 0x30, 0x55, 0xde, 0xdb, 0x0f, 0x2d, 0xef, 0x99, 0x26, 0x21,
	0xfe, 0x53, 0x6d, 0x26, 0x74, 0x02, 0x26, 0x36, 0x9c, 0xa5, 0x32, 0x5b, 0x5b, 0x57, 0xa3, 0x2a,
	0x45, 0xde, 0x43, 0xa7, 0x5c, 0x38, 0x6f, 0xf2, 0x91, 0x08, 0x55, 0x9d, 0xf5, 0x5b, 0x83, 0x33,
	0x17, 0xff, 0x7e, 0x7f, 0xf9, 0x55, 0xec, 0x43, 0x43, 0x26, 0x77, 0x18, 0x17, 0x49, 0x72, 0xf0,
	0xff, 0x2e, 0x63, 0xf6, 0xd6, 0x02, 0xf4, 0x83, 0x90, 0xc5, 0x68, 0x34, 0x55, 0x3d, 0x47, 0x7c,
	0x79, 0x93, 0x17, 0x98, 0x1d, 0x48, 0xce, 0xa1, 0xbd, 0xf0, 0x56, 0xf3, 0x5b, 0x3a, 0x1b, 0x4f,
	0x7b, 0x4f, 0xc8, 0x05, 0x3c, 0x5d, 0x78, 0xab, 0x6b, 0x3a, 0x9e, 0xdf, 0x7c, 0x59, 0x8d, 0x27,
	0xb3, 0xaf, 0xf3, 0x9e, 0x56, 0x90, 0xee, 0x77, 0xd7, 0x73, 0x66, 0x05, 0x59, 0xbb, 0xfc, 0x04,
	0x50, 0x86, 0x51, 0x36, 0x6e, 0x69, 0x93, 0xc3, 0xcf, 0xf4, 0x76, 0xe9, 0x64, 0x06, 0x5d, 0x80,
	0x85, 0xbb, 0x9a, 0x38, 0x53, 0xc7, 0x73, 0x26, 0xbd, 0xda, 0xb5, 0xbd, 0x7c, 0xbb, 0x65, 0x72,
	0xb7, 0x5f, 0xdb, 0x9b, 0x24, 0x1a, 0x46, 0x2c, 0x46, 0xee, 0x87, 0x1c, 0xc5, 0xf0, 0xf0, 0xaf,
	0x4f, 0xef, 0xb6, 0xc3, 0x74, 0x7d, 0x80, 0xeb, 0xa6, 0xfa, 0x00, 0x5c, 0xfd, 0x09, 0x00, 0x00,
	0xff, 0xff, 0x7f, 0xb1, 0x9f, 0x31, 0x16, 0x04, 0x00, 0x00,
}
