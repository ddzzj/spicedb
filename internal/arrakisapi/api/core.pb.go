// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: REDACTEDapi/api/core.proto

package REDACTEDapi

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type RelationTupleUpdate_Operation int32

const (
	RelationTupleUpdate_UNKNOWN RelationTupleUpdate_Operation = 0
	RelationTupleUpdate_CREATE  RelationTupleUpdate_Operation = 1
	RelationTupleUpdate_TOUCH   RelationTupleUpdate_Operation = 2
	RelationTupleUpdate_DELETE  RelationTupleUpdate_Operation = 3
)

// Enum value maps for RelationTupleUpdate_Operation.
var (
	RelationTupleUpdate_Operation_name = map[int32]string{
		0: "UNKNOWN",
		1: "CREATE",
		2: "TOUCH",
		3: "DELETE",
	}
	RelationTupleUpdate_Operation_value = map[string]int32{
		"UNKNOWN": 0,
		"CREATE":  1,
		"TOUCH":   2,
		"DELETE":  3,
	}
)

func (x RelationTupleUpdate_Operation) Enum() *RelationTupleUpdate_Operation {
	p := new(RelationTupleUpdate_Operation)
	*p = x
	return p
}

func (x RelationTupleUpdate_Operation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RelationTupleUpdate_Operation) Descriptor() protoreflect.EnumDescriptor {
	return file_REDACTEDapi_api_core_proto_enumTypes[0].Descriptor()
}

func (RelationTupleUpdate_Operation) Type() protoreflect.EnumType {
	return &file_REDACTEDapi_api_core_proto_enumTypes[0]
}

func (x RelationTupleUpdate_Operation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RelationTupleUpdate_Operation.Descriptor instead.
func (RelationTupleUpdate_Operation) EnumDescriptor() ([]byte, []int) {
	return file_REDACTEDapi_api_core_proto_rawDescGZIP(), []int{4, 0}
}

type SetOperationUserset_Operation int32

const (
	SetOperationUserset_INVALID      SetOperationUserset_Operation = 0
	SetOperationUserset_UNION        SetOperationUserset_Operation = 1
	SetOperationUserset_INTERSECTION SetOperationUserset_Operation = 2
	SetOperationUserset_EXCLUSION    SetOperationUserset_Operation = 3
)

// Enum value maps for SetOperationUserset_Operation.
var (
	SetOperationUserset_Operation_name = map[int32]string{
		0: "INVALID",
		1: "UNION",
		2: "INTERSECTION",
		3: "EXCLUSION",
	}
	SetOperationUserset_Operation_value = map[string]int32{
		"INVALID":      0,
		"UNION":        1,
		"INTERSECTION": 2,
		"EXCLUSION":    3,
	}
)

func (x SetOperationUserset_Operation) Enum() *SetOperationUserset_Operation {
	p := new(SetOperationUserset_Operation)
	*p = x
	return p
}

func (x SetOperationUserset_Operation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SetOperationUserset_Operation) Descriptor() protoreflect.EnumDescriptor {
	return file_REDACTEDapi_api_core_proto_enumTypes[1].Descriptor()
}

func (SetOperationUserset_Operation) Type() protoreflect.EnumType {
	return &file_REDACTEDapi_api_core_proto_enumTypes[1]
}

func (x SetOperationUserset_Operation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SetOperationUserset_Operation.Descriptor instead.
func (SetOperationUserset_Operation) EnumDescriptor() ([]byte, []int) {
	return file_REDACTEDapi_api_core_proto_rawDescGZIP(), []int{6, 0}
}

type RelationTuple struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Each tupleset specifies keys of a set of relation tuples. The set can
	// include a single tuple key, or all tuples with a given object ID or
	// userset in a namespace, optionally constrained by a relation name.
	//
	// examples:
	// doc:readme#viewer@group:eng#member (fully specified)
	// doc:*#*#group:eng#member (all tuples that this userset relates to)
	// doc:12345#*#* (all tuples with a direct relationship to a document)
	// doc:12345#writer#* (all tuples with direct write relationship with the
	// document) doc:#writer#group:eng#member (all tuples that eng group has write
	// relationship)
	ObjectAndRelation *ObjectAndRelation `protobuf:"bytes,1,opt,name=object_and_relation,json=objectAndRelation,proto3" json:"object_and_relation,omitempty"`
	User              *User              `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *RelationTuple) Reset() {
	*x = RelationTuple{}
	if protoimpl.UnsafeEnabled {
		mi := &file_REDACTEDapi_api_core_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationTuple) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationTuple) ProtoMessage() {}

func (x *RelationTuple) ProtoReflect() protoreflect.Message {
	mi := &file_REDACTEDapi_api_core_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationTuple.ProtoReflect.Descriptor instead.
func (*RelationTuple) Descriptor() ([]byte, []int) {
	return file_REDACTEDapi_api_core_proto_rawDescGZIP(), []int{0}
}

func (x *RelationTuple) GetObjectAndRelation() *ObjectAndRelation {
	if x != nil {
		return x.ObjectAndRelation
	}
	return nil
}

func (x *RelationTuple) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type ObjectAndRelation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	ObjectId  string `protobuf:"bytes,2,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
	Relation  string `protobuf:"bytes,3,opt,name=relation,proto3" json:"relation,omitempty"`
}

func (x *ObjectAndRelation) Reset() {
	*x = ObjectAndRelation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_REDACTEDapi_api_core_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectAndRelation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectAndRelation) ProtoMessage() {}

func (x *ObjectAndRelation) ProtoReflect() protoreflect.Message {
	mi := &file_REDACTEDapi_api_core_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectAndRelation.ProtoReflect.Descriptor instead.
func (*ObjectAndRelation) Descriptor() ([]byte, []int) {
	return file_REDACTEDapi_api_core_proto_rawDescGZIP(), []int{1}
}

func (x *ObjectAndRelation) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ObjectAndRelation) GetObjectId() string {
	if x != nil {
		return x.ObjectId
	}
	return ""
}

func (x *ObjectAndRelation) GetRelation() string {
	if x != nil {
		return x.Relation
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to UserOneof:
	//	*User_Userset
	UserOneof isUser_UserOneof `protobuf_oneof:"user_oneof"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_REDACTEDapi_api_core_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_REDACTEDapi_api_core_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_REDACTEDapi_api_core_proto_rawDescGZIP(), []int{2}
}

func (m *User) GetUserOneof() isUser_UserOneof {
	if m != nil {
		return m.UserOneof
	}
	return nil
}

func (x *User) GetUserset() *ObjectAndRelation {
	if x, ok := x.GetUserOneof().(*User_Userset); ok {
		return x.Userset
	}
	return nil
}

type isUser_UserOneof interface {
	isUser_UserOneof()
}

type User_Userset struct {
	Userset *ObjectAndRelation `protobuf:"bytes,2,opt,name=userset,proto3,oneof"`
}

func (*User_Userset) isUser_UserOneof() {}

type Zookie struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *Zookie) Reset() {
	*x = Zookie{}
	if protoimpl.UnsafeEnabled {
		mi := &file_REDACTEDapi_api_core_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Zookie) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Zookie) ProtoMessage() {}

func (x *Zookie) ProtoReflect() protoreflect.Message {
	mi := &file_REDACTEDapi_api_core_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Zookie.ProtoReflect.Descriptor instead.
func (*Zookie) Descriptor() ([]byte, []int) {
	return file_REDACTEDapi_api_core_proto_rawDescGZIP(), []int{3}
}

func (x *Zookie) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type RelationTupleUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operation RelationTupleUpdate_Operation `protobuf:"varint,1,opt,name=operation,proto3,enum=RelationTupleUpdate_Operation" json:"operation,omitempty"`
	Tuple     *RelationTuple                `protobuf:"bytes,2,opt,name=tuple,proto3" json:"tuple,omitempty"`
}

func (x *RelationTupleUpdate) Reset() {
	*x = RelationTupleUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_REDACTEDapi_api_core_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationTupleUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationTupleUpdate) ProtoMessage() {}

func (x *RelationTupleUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_REDACTEDapi_api_core_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationTupleUpdate.ProtoReflect.Descriptor instead.
func (*RelationTupleUpdate) Descriptor() ([]byte, []int) {
	return file_REDACTEDapi_api_core_proto_rawDescGZIP(), []int{4}
}

func (x *RelationTupleUpdate) GetOperation() RelationTupleUpdate_Operation {
	if x != nil {
		return x.Operation
	}
	return RelationTupleUpdate_UNKNOWN
}

func (x *RelationTupleUpdate) GetTuple() *RelationTuple {
	if x != nil {
		return x.Tuple
	}
	return nil
}

type RelationTupleTreeNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to NodeType:
	//	*RelationTupleTreeNode_IntermediateNode
	//	*RelationTupleTreeNode_LeafNode
	NodeType isRelationTupleTreeNode_NodeType `protobuf_oneof:"node_type"`
	Expanded *ObjectAndRelation               `protobuf:"bytes,3,opt,name=expanded,proto3" json:"expanded,omitempty"`
}

func (x *RelationTupleTreeNode) Reset() {
	*x = RelationTupleTreeNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_REDACTEDapi_api_core_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationTupleTreeNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationTupleTreeNode) ProtoMessage() {}

func (x *RelationTupleTreeNode) ProtoReflect() protoreflect.Message {
	mi := &file_REDACTEDapi_api_core_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationTupleTreeNode.ProtoReflect.Descriptor instead.
func (*RelationTupleTreeNode) Descriptor() ([]byte, []int) {
	return file_REDACTEDapi_api_core_proto_rawDescGZIP(), []int{5}
}

func (m *RelationTupleTreeNode) GetNodeType() isRelationTupleTreeNode_NodeType {
	if m != nil {
		return m.NodeType
	}
	return nil
}

func (x *RelationTupleTreeNode) GetIntermediateNode() *SetOperationUserset {
	if x, ok := x.GetNodeType().(*RelationTupleTreeNode_IntermediateNode); ok {
		return x.IntermediateNode
	}
	return nil
}

func (x *RelationTupleTreeNode) GetLeafNode() *DirectUserset {
	if x, ok := x.GetNodeType().(*RelationTupleTreeNode_LeafNode); ok {
		return x.LeafNode
	}
	return nil
}

func (x *RelationTupleTreeNode) GetExpanded() *ObjectAndRelation {
	if x != nil {
		return x.Expanded
	}
	return nil
}

type isRelationTupleTreeNode_NodeType interface {
	isRelationTupleTreeNode_NodeType()
}

type RelationTupleTreeNode_IntermediateNode struct {
	IntermediateNode *SetOperationUserset `protobuf:"bytes,1,opt,name=intermediate_node,json=intermediateNode,proto3,oneof"`
}

type RelationTupleTreeNode_LeafNode struct {
	LeafNode *DirectUserset `protobuf:"bytes,2,opt,name=leaf_node,json=leafNode,proto3,oneof"`
}

func (*RelationTupleTreeNode_IntermediateNode) isRelationTupleTreeNode_NodeType() {}

func (*RelationTupleTreeNode_LeafNode) isRelationTupleTreeNode_NodeType() {}

type SetOperationUserset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operation  SetOperationUserset_Operation `protobuf:"varint,1,opt,name=operation,proto3,enum=SetOperationUserset_Operation" json:"operation,omitempty"`
	ChildNodes []*RelationTupleTreeNode      `protobuf:"bytes,2,rep,name=child_nodes,json=childNodes,proto3" json:"child_nodes,omitempty"`
}

func (x *SetOperationUserset) Reset() {
	*x = SetOperationUserset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_REDACTEDapi_api_core_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetOperationUserset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetOperationUserset) ProtoMessage() {}

func (x *SetOperationUserset) ProtoReflect() protoreflect.Message {
	mi := &file_REDACTEDapi_api_core_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetOperationUserset.ProtoReflect.Descriptor instead.
func (*SetOperationUserset) Descriptor() ([]byte, []int) {
	return file_REDACTEDapi_api_core_proto_rawDescGZIP(), []int{6}
}

func (x *SetOperationUserset) GetOperation() SetOperationUserset_Operation {
	if x != nil {
		return x.Operation
	}
	return SetOperationUserset_INVALID
}

func (x *SetOperationUserset) GetChildNodes() []*RelationTupleTreeNode {
	if x != nil {
		return x.ChildNodes
	}
	return nil
}

type DirectUserset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *DirectUserset) Reset() {
	*x = DirectUserset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_REDACTEDapi_api_core_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DirectUserset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DirectUserset) ProtoMessage() {}

func (x *DirectUserset) ProtoReflect() protoreflect.Message {
	mi := &file_REDACTEDapi_api_core_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DirectUserset.ProtoReflect.Descriptor instead.
func (*DirectUserset) Descriptor() ([]byte, []int) {
	return file_REDACTEDapi_api_core_proto_rawDescGZIP(), []int{7}
}

func (x *DirectUserset) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

var File_REDACTEDapi_api_core_proto protoreflect.FileDescriptor

var file_REDACTEDapi_api_core_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x72, 0x72, 0x61, 0x6b, 0x69, 0x73, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6e, 0x0a, 0x0d, 0x52,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x12, 0x42, 0x0a, 0x13,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x61, 0x6e, 0x64, 0x5f, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x41, 0x6e, 0x64, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x41, 0x6e, 0x64, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x19, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x6a, 0x0a, 0x11, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x41, 0x6e, 0x64, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x44, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x2e, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x41, 0x6e, 0x64, 0x52, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x73, 0x65, 0x74, 0x42,
	0x0c, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x22, 0x1e, 0x0a,
	0x06, 0x5a, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xb6, 0x01,
	0x0a, 0x13, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x3c, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x05, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x75, 0x70,
	0x6c, 0x65, 0x52, 0x05, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x22, 0x3b, 0x0a, 0x09, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x10, 0x01, 0x12,
	0x09, 0x0a, 0x05, 0x54, 0x4f, 0x55, 0x43, 0x48, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x45,
	0x4c, 0x45, 0x54, 0x45, 0x10, 0x03, 0x22, 0xc8, 0x01, 0x0a, 0x15, 0x52, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x54, 0x72, 0x65, 0x65, 0x4e, 0x6f, 0x64, 0x65,
	0x12, 0x43, 0x0a, 0x11, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x65,
	0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x53, 0x65,
	0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x65,
	0x74, 0x48, 0x00, 0x52, 0x10, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74,
	0x65, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x2d, 0x0a, 0x09, 0x6c, 0x65, 0x61, 0x66, 0x5f, 0x6e, 0x6f,
	0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x65, 0x74, 0x48, 0x00, 0x52, 0x08, 0x6c, 0x65, 0x61, 0x66,
	0x4e, 0x6f, 0x64, 0x65, 0x12, 0x2e, 0x0a, 0x08, 0x65, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x65, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x41,
	0x6e, 0x64, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x65, 0x78, 0x70, 0x61,
	0x6e, 0x64, 0x65, 0x64, 0x42, 0x0b, 0x0a, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x22, 0xd2, 0x01, 0x0a, 0x13, 0x53, 0x65, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x65, 0x74, 0x12, 0x3c, 0x0a, 0x09, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x53,
	0x65, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x73,
	0x65, 0x74, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x37, 0x0a, 0x0b, 0x63, 0x68, 0x69, 0x6c, 0x64,
	0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x52,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x54, 0x72, 0x65, 0x65,
	0x4e, 0x6f, 0x64, 0x65, 0x52, 0x0a, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x73,
	0x22, 0x44, 0x0a, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0b, 0x0a,
	0x07, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x55, 0x4e,
	0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x53, 0x45,
	0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x45, 0x58, 0x43, 0x4c, 0x55,
	0x53, 0x49, 0x4f, 0x4e, 0x10, 0x03, 0x22, 0x2c, 0x0a, 0x0d, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x65, 0x74, 0x12, 0x1b, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x70, 0x65, 0x74, 0x72, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x63, 0x6f, 0x64,
	0x65, 0x2f, 0x61, 0x72, 0x72, 0x61, 0x6b, 0x69, 0x73, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_REDACTEDapi_api_core_proto_rawDescOnce sync.Once
	file_REDACTEDapi_api_core_proto_rawDescData = file_REDACTEDapi_api_core_proto_rawDesc
)

func file_REDACTEDapi_api_core_proto_rawDescGZIP() []byte {
	file_REDACTEDapi_api_core_proto_rawDescOnce.Do(func() {
		file_REDACTEDapi_api_core_proto_rawDescData = protoimpl.X.CompressGZIP(file_REDACTEDapi_api_core_proto_rawDescData)
	})
	return file_REDACTEDapi_api_core_proto_rawDescData
}

var file_REDACTEDapi_api_core_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_REDACTEDapi_api_core_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_REDACTEDapi_api_core_proto_goTypes = []interface{}{
	(RelationTupleUpdate_Operation)(0), // 0: RelationTupleUpdate.Operation
	(SetOperationUserset_Operation)(0), // 1: SetOperationUserset.Operation
	(*RelationTuple)(nil),              // 2: RelationTuple
	(*ObjectAndRelation)(nil),          // 3: ObjectAndRelation
	(*User)(nil),                       // 4: User
	(*Zookie)(nil),                     // 5: Zookie
	(*RelationTupleUpdate)(nil),        // 6: RelationTupleUpdate
	(*RelationTupleTreeNode)(nil),      // 7: RelationTupleTreeNode
	(*SetOperationUserset)(nil),        // 8: SetOperationUserset
	(*DirectUserset)(nil),              // 9: DirectUserset
}
var file_REDACTEDapi_api_core_proto_depIdxs = []int32{
	3,  // 0: RelationTuple.object_and_relation:type_name -> ObjectAndRelation
	4,  // 1: RelationTuple.user:type_name -> User
	3,  // 2: User.userset:type_name -> ObjectAndRelation
	0,  // 3: RelationTupleUpdate.operation:type_name -> RelationTupleUpdate.Operation
	2,  // 4: RelationTupleUpdate.tuple:type_name -> RelationTuple
	8,  // 5: RelationTupleTreeNode.intermediate_node:type_name -> SetOperationUserset
	9,  // 6: RelationTupleTreeNode.leaf_node:type_name -> DirectUserset
	3,  // 7: RelationTupleTreeNode.expanded:type_name -> ObjectAndRelation
	1,  // 8: SetOperationUserset.operation:type_name -> SetOperationUserset.Operation
	7,  // 9: SetOperationUserset.child_nodes:type_name -> RelationTupleTreeNode
	4,  // 10: DirectUserset.users:type_name -> User
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_REDACTEDapi_api_core_proto_init() }
func file_REDACTEDapi_api_core_proto_init() {
	if File_REDACTEDapi_api_core_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_REDACTEDapi_api_core_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationTuple); i {
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
		file_REDACTEDapi_api_core_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectAndRelation); i {
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
		file_REDACTEDapi_api_core_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_REDACTEDapi_api_core_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Zookie); i {
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
		file_REDACTEDapi_api_core_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationTupleUpdate); i {
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
		file_REDACTEDapi_api_core_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationTupleTreeNode); i {
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
		file_REDACTEDapi_api_core_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetOperationUserset); i {
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
		file_REDACTEDapi_api_core_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DirectUserset); i {
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
	file_REDACTEDapi_api_core_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*User_Userset)(nil),
	}
	file_REDACTEDapi_api_core_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*RelationTupleTreeNode_IntermediateNode)(nil),
		(*RelationTupleTreeNode_LeafNode)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_REDACTEDapi_api_core_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_REDACTEDapi_api_core_proto_goTypes,
		DependencyIndexes: file_REDACTEDapi_api_core_proto_depIdxs,
		EnumInfos:         file_REDACTEDapi_api_core_proto_enumTypes,
		MessageInfos:      file_REDACTEDapi_api_core_proto_msgTypes,
	}.Build()
	File_REDACTEDapi_api_core_proto = out.File
	file_REDACTEDapi_api_core_proto_rawDesc = nil
	file_REDACTEDapi_api_core_proto_goTypes = nil
	file_REDACTEDapi_api_core_proto_depIdxs = nil
}
