// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: node.proto

package node

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// //message google {
//
//	//  message protobuf {
//	   message Timestamp {
//	     int64 seconds = 1;
//	     int32 nanos = 2;
//	   }
//
// //  }
// //}
type NodeStatus int32

const (
	NodeStatus_UNKNOWN   NodeStatus = 0
	NodeStatus_CLIENT    NodeStatus = 1
	NodeStatus_LISTENER  NodeStatus = 2
	NodeStatus_LEADER    NodeStatus = 3
	NodeStatus_FOLLOWER  NodeStatus = 4
	NodeStatus_CANDIDATE NodeStatus = 5
)

// Enum value maps for NodeStatus.
var (
	NodeStatus_name = map[int32]string{
		0: "UNKNOWN",
		1: "CLIENT",
		2: "LISTENER",
		3: "LEADER",
		4: "FOLLOWER",
		5: "CANDIDATE",
	}
	NodeStatus_value = map[string]int32{
		"UNKNOWN":   0,
		"CLIENT":    1,
		"LISTENER":  2,
		"LEADER":    3,
		"FOLLOWER":  4,
		"CANDIDATE": 5,
	}
)

func (x NodeStatus) Enum() *NodeStatus {
	p := new(NodeStatus)
	*p = x
	return p
}

func (x NodeStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NodeStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_node_proto_enumTypes[0].Descriptor()
}

func (NodeStatus) Type() protoreflect.EnumType {
	return &file_node_proto_enumTypes[0]
}

func (x NodeStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NodeStatus.Descriptor instead.
func (NodeStatus) EnumDescriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{0}
}

type RespStatus int32

const (
	RespStatus_OK    RespStatus = 0
	RespStatus_NO    RespStatus = 1
	RespStatus_ERROR RespStatus = 2
)

// Enum value maps for RespStatus.
var (
	RespStatus_name = map[int32]string{
		0: "OK",
		1: "NO",
		2: "ERROR",
	}
	RespStatus_value = map[string]int32{
		"OK":    0,
		"NO":    1,
		"ERROR": 2,
	}
)

func (x RespStatus) Enum() *RespStatus {
	p := new(RespStatus)
	*p = x
	return p
}

func (x RespStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RespStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_node_proto_enumTypes[1].Descriptor()
}

func (RespStatus) Type() protoreflect.EnumType {
	return &file_node_proto_enumTypes[1]
}

func (x RespStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RespStatus.Descriptor instead.
func (RespStatus) EnumDescriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{1}
}

type DataNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
}

func (x *DataNode) Reset() {
	*x = DataNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataNode) ProtoMessage() {}

func (x *DataNode) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataNode.ProtoReflect.Descriptor instead.
func (*DataNode) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{0}
}

func (x *DataNode) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

type KnownNodes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nodes map[string]*DataNode `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *KnownNodes) Reset() {
	*x = KnownNodes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KnownNodes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KnownNodes) ProtoMessage() {}

func (x *KnownNodes) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KnownNodes.ProtoReflect.Descriptor instead.
func (*KnownNodes) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{1}
}

func (x *KnownNodes) GetNodes() map[string]*DataNode {
	if x != nil {
		return x.Nodes
	}
	return nil
}

type PingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PingRequest) Reset() {
	*x = PingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRequest) ProtoMessage() {}

func (x *PingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRequest.ProtoReflect.Descriptor instead.
func (*PingRequest) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{2}
}

type PingResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Res bool `protobuf:"varint,1,opt,name=res,proto3" json:"res,omitempty"`
}

func (x *PingResult) Reset() {
	*x = PingResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingResult) ProtoMessage() {}

func (x *PingResult) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingResult.ProtoReflect.Descriptor instead.
func (*PingResult) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{3}
}

func (x *PingResult) GetRes() bool {
	if x != nil {
		return x.Res
	}
	return false
}

type ArtResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Art *Artefact  `protobuf:"bytes,1,opt,name=art,proto3" json:"art,omitempty"`
	Res RespStatus `protobuf:"varint,2,opt,name=res,proto3,enum=RespStatus" json:"res,omitempty"`
}

func (x *ArtResp) Reset() {
	*x = ArtResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArtResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArtResp) ProtoMessage() {}

func (x *ArtResp) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArtResp.ProtoReflect.Descriptor instead.
func (*ArtResp) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{4}
}

func (x *ArtResp) GetArt() *Artefact {
	if x != nil {
		return x.Art
	}
	return nil
}

func (x *ArtResp) GetRes() RespStatus {
	if x != nil {
		return x.Res
	}
	return RespStatus_OK
}

type Artefact struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	HashCluster string `protobuf:"bytes,2,opt,name=hashCluster,proto3" json:"hashCluster,omitempty"`
}

func (x *Artefact) Reset() {
	*x = Artefact{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Artefact) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Artefact) ProtoMessage() {}

func (x *Artefact) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Artefact.ProtoReflect.Descriptor instead.
func (*Artefact) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{5}
}

func (x *Artefact) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Artefact) GetHashCluster() string {
	if x != nil {
		return x.HashCluster
	}
	return ""
}

type CliReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Req       *Request `protobuf:"bytes,1,opt,name=req,proto3" json:"req,omitempty"`
	Addresses []string `protobuf:"bytes,2,rep,name=addresses,proto3" json:"addresses,omitempty"`
}

func (x *CliReq) Reset() {
	*x = CliReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CliReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CliReq) ProtoMessage() {}

func (x *CliReq) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CliReq.ProtoReflect.Descriptor instead.
func (*CliReq) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{6}
}

func (x *CliReq) GetReq() *Request {
	if x != nil {
		return x.Req
	}
	return nil
}

func (x *CliReq) GetAddresses() []string {
	if x != nil {
		return x.Addresses
	}
	return nil
}

type NodeResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value       string                `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Code        uint64                `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	CountResult uint64                `protobuf:"varint,3,opt,name=countResult,proto3" json:"countResult,omitempty"`
	Result      map[string]RespStatus `protobuf:"bytes,4,rep,name=result,proto3" json:"result,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=RespStatus"`
	HashCluster string                `protobuf:"bytes,5,opt,name=hashCluster,proto3" json:"hashCluster,omitempty"`
}

func (x *NodeResp) Reset() {
	*x = NodeResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeResp) ProtoMessage() {}

func (x *NodeResp) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeResp.ProtoReflect.Descriptor instead.
func (*NodeResp) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{7}
}

func (x *NodeResp) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *NodeResp) GetCode() uint64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *NodeResp) GetCountResult() uint64 {
	if x != nil {
		return x.CountResult
	}
	return 0
}

func (x *NodeResp) GetResult() map[string]RespStatus {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *NodeResp) GetHashCluster() string {
	if x != nil {
		return x.HashCluster
	}
	return ""
}

type Info struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address   string      `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Repl      int32       `protobuf:"varint,2,opt,name=repl,proto3" json:"repl,omitempty"`
	SizeVault int32       `protobuf:"varint,3,opt,name=sizeVault,proto3" json:"sizeVault,omitempty"`
	Status    NodeStatus  `protobuf:"varint,4,opt,name=status,proto3,enum=NodeStatus" json:"status,omitempty"`
	Env       *KnownNodes `protobuf:"bytes,5,opt,name=env,proto3" json:"env,omitempty"`
}

func (x *Info) Reset() {
	*x = Info{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Info) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Info) ProtoMessage() {}

func (x *Info) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Info.ProtoReflect.Descriptor instead.
func (*Info) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{8}
}

func (x *Info) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Info) GetRepl() int32 {
	if x != nil {
		return x.Repl
	}
	return 0
}

func (x *Info) GetSizeVault() int32 {
	if x != nil {
		return x.SizeVault
	}
	return 0
}

func (x *Info) GetStatus() NodeStatus {
	if x != nil {
		return x.Status
	}
	return NodeStatus_UNKNOWN
}

func (x *Info) GetEnv() *KnownNodes {
	if x != nil {
		return x.Env
	}
	return nil
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comm        string `protobuf:"bytes,1,opt,name=comm,proto3" json:"comm,omitempty"`
	Uuid        string `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Value       string `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	HashCluster string `protobuf:"bytes,4,opt,name=hashCluster,proto3" json:"hashCluster,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{9}
}

func (x *Request) GetComm() string {
	if x != nil {
		return x.Comm
	}
	return ""
}

func (x *Request) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Request) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Request) GetHashCluster() string {
	if x != nil {
		return x.HashCluster
	}
	return ""
}

var File_node_proto protoreflect.FileDescriptor

var file_node_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x36, 0x0a,
	0x08, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x02, 0x74, 0x73, 0x22, 0x7f, 0x0a, 0x0a, 0x4b, 0x6e, 0x6f, 0x77, 0x6e, 0x4e, 0x6f,
	0x64, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x4b, 0x6e, 0x6f, 0x77, 0x6e, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x2e,
	0x4e, 0x6f, 0x64, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65,
	0x73, 0x1a, 0x43, 0x0a, 0x0a, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x1f, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x09, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x0d, 0x0a, 0x0b, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x1e, 0x0a, 0x0a, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x03, 0x72, 0x65, 0x73, 0x22, 0x45, 0x0a, 0x07, 0x41, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x1b, 0x0a, 0x03, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e,
	0x41, 0x72, 0x74, 0x65, 0x66, 0x61, 0x63, 0x74, 0x52, 0x03, 0x61, 0x72, 0x74, 0x12, 0x1d, 0x0a,
	0x03, 0x72, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x03, 0x72, 0x65, 0x73, 0x22, 0x40, 0x0a, 0x08,
	0x41, 0x72, 0x74, 0x65, 0x66, 0x61, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x68, 0x61, 0x73, 0x68, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x68, 0x61, 0x73, 0x68, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x22, 0x42,
	0x0a, 0x06, 0x43, 0x6c, 0x69, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x03, 0x72, 0x65, 0x71, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52,
	0x03, 0x72, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x65, 0x73, 0x22, 0xef, 0x01, 0x0a, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x2d, 0x0a, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x4e, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x68, 0x61,
	0x73, 0x68, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x68, 0x61, 0x73, 0x68, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x1a, 0x46, 0x0a, 0x0b,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x21, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x96, 0x01, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x65, 0x70, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x72, 0x65, 0x70, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x73,
	0x69, 0x7a, 0x65, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x73, 0x69, 0x7a, 0x65, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x23, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x4e, 0x6f, 0x64, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d,
	0x0a, 0x03, 0x65, 0x6e, 0x76, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x4b, 0x6e,
	0x6f, 0x77, 0x6e, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x03, 0x65, 0x6e, 0x76, 0x22, 0x69, 0x0a,
	0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x6d, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x6d, 0x6d, 0x12, 0x12, 0x0a, 0x04,
	0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x68, 0x61, 0x73, 0x68, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x68, 0x61, 0x73,
	0x68, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2a, 0x5c, 0x0a, 0x0a, 0x4e, 0x6f, 0x64, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x12,
	0x0c, 0x0a, 0x08, 0x4c, 0x49, 0x53, 0x54, 0x45, 0x4e, 0x45, 0x52, 0x10, 0x02, 0x12, 0x0a, 0x0a,
	0x06, 0x4c, 0x45, 0x41, 0x44, 0x45, 0x52, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x46, 0x4f, 0x4c,
	0x4c, 0x4f, 0x57, 0x45, 0x52, 0x10, 0x04, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x41, 0x4e, 0x44, 0x49,
	0x44, 0x41, 0x54, 0x45, 0x10, 0x05, 0x2a, 0x27, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x70, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02,
	0x4e, 0x4f, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x32,
	0x89, 0x02, 0x0a, 0x11, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x0c, 0x2e,
	0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x50, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x05, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x05, 0x2e, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x20, 0x0a, 0x0a, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x32, 0x50, 0x12,
	0x07, 0x2e, 0x43, 0x6c, 0x69, 0x52, 0x65, 0x71, 0x1a, 0x09, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x19, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x08, 0x2e, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x08, 0x2e, 0x41, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x19,
	0x0a, 0x03, 0x53, 0x65, 0x74, 0x12, 0x08, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x08, 0x2e, 0x41, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1c, 0x0a, 0x06, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x08, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x08, 0x2e,
	0x41, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1e, 0x0a, 0x04, 0x52, 0x65, 0x70, 0x6c, 0x12,
	0x09, 0x2e, 0x41, 0x72, 0x74, 0x65, 0x66, 0x61, 0x63, 0x74, 0x1a, 0x0b, 0x2e, 0x50, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x78, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x07, 0x2e, 0x43, 0x6c, 0x69, 0x52, 0x65, 0x71,
	0x1a, 0x09, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x32, 0x58, 0x0a, 0x13, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x4e, 0x6f, 0x64,
	0x65, 0x12, 0x0c, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x05, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x07, 0x2e, 0x43, 0x6c, 0x69, 0x52, 0x65, 0x71, 0x1a, 0x09, 0x2e, 0x4e, 0x6f, 0x64,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x42, 0x15, 0x5a, 0x13, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_node_proto_rawDescOnce sync.Once
	file_node_proto_rawDescData = file_node_proto_rawDesc
)

func file_node_proto_rawDescGZIP() []byte {
	file_node_proto_rawDescOnce.Do(func() {
		file_node_proto_rawDescData = protoimpl.X.CompressGZIP(file_node_proto_rawDescData)
	})
	return file_node_proto_rawDescData
}

var file_node_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_node_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_node_proto_goTypes = []interface{}{
	(NodeStatus)(0),               // 0: NodeStatus
	(RespStatus)(0),               // 1: RespStatus
	(*DataNode)(nil),              // 2: DataNode
	(*KnownNodes)(nil),            // 3: KnownNodes
	(*PingRequest)(nil),           // 4: PingRequest
	(*PingResult)(nil),            // 5: PingResult
	(*ArtResp)(nil),               // 6: ArtResp
	(*Artefact)(nil),              // 7: Artefact
	(*CliReq)(nil),                // 8: CliReq
	(*NodeResp)(nil),              // 9: NodeResp
	(*Info)(nil),                  // 10: Info
	(*Request)(nil),               // 11: request
	nil,                           // 12: KnownNodes.NodesEntry
	nil,                           // 13: NodeResp.ResultEntry
	(*timestamppb.Timestamp)(nil), // 14: google.protobuf.Timestamp
}
var file_node_proto_depIdxs = []int32{
	14, // 0: DataNode.ts:type_name -> google.protobuf.Timestamp
	12, // 1: KnownNodes.nodes:type_name -> KnownNodes.NodesEntry
	7,  // 2: ArtResp.art:type_name -> Artefact
	1,  // 3: ArtResp.res:type_name -> RespStatus
	11, // 4: CliReq.req:type_name -> request
	13, // 5: NodeResp.result:type_name -> NodeResp.ResultEntry
	0,  // 6: Info.status:type_name -> NodeStatus
	3,  // 7: Info.env:type_name -> KnownNodes
	2,  // 8: KnownNodes.NodesEntry.value:type_name -> DataNode
	1,  // 9: NodeResp.ResultEntry.value:type_name -> RespStatus
	4,  // 10: NodeCommunication.Ping:input_type -> PingRequest
	10, // 11: NodeCommunication.GetInfo:input_type -> Info
	8,  // 12: NodeCommunication.RequestP2P:input_type -> CliReq
	11, // 13: NodeCommunication.Get:input_type -> request
	11, // 14: NodeCommunication.Set:input_type -> request
	11, // 15: NodeCommunication.Delete:input_type -> request
	7,  // 16: NodeCommunication.Repl:input_type -> Artefact
	8,  // 17: NodeCommunication.ProxyRequest:input_type -> CliReq
	4,  // 18: ClientCommunication.GetInfoNode:input_type -> PingRequest
	8,  // 19: ClientCommunication.Request:input_type -> CliReq
	5,  // 20: NodeCommunication.Ping:output_type -> PingResult
	10, // 21: NodeCommunication.GetInfo:output_type -> Info
	9,  // 22: NodeCommunication.RequestP2P:output_type -> NodeResp
	6,  // 23: NodeCommunication.Get:output_type -> ArtResp
	6,  // 24: NodeCommunication.Set:output_type -> ArtResp
	6,  // 25: NodeCommunication.Delete:output_type -> ArtResp
	5,  // 26: NodeCommunication.Repl:output_type -> PingResult
	9,  // 27: NodeCommunication.ProxyRequest:output_type -> NodeResp
	10, // 28: ClientCommunication.GetInfoNode:output_type -> Info
	9,  // 29: ClientCommunication.Request:output_type -> NodeResp
	20, // [20:30] is the sub-list for method output_type
	10, // [10:20] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_node_proto_init() }
func file_node_proto_init() {
	if File_node_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_node_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataNode); i {
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
		file_node_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KnownNodes); i {
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
		file_node_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingRequest); i {
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
		file_node_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingResult); i {
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
		file_node_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArtResp); i {
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
		file_node_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Artefact); i {
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
		file_node_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CliReq); i {
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
		file_node_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeResp); i {
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
		file_node_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Info); i {
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
		file_node_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
			RawDescriptor: file_node_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_node_proto_goTypes,
		DependencyIndexes: file_node_proto_depIdxs,
		EnumInfos:         file_node_proto_enumTypes,
		MessageInfos:      file_node_proto_msgTypes,
	}.Build()
	File_node_proto = out.File
	file_node_proto_rawDesc = nil
	file_node_proto_goTypes = nil
	file_node_proto_depIdxs = nil
}
