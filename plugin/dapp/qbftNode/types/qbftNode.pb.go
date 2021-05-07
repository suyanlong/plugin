// Code generated by protoc-gen-go. DO NOT EDIT.
// source: qbftNode.proto

package types

import (
	context "context"
	fmt "fmt"
	math "math"

	types "github.com/33cn/chain33/types"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type QbftNode struct {
	PubKey               string   `protobuf:"bytes,1,opt,name=pubKey,proto3" json:"pubKey,omitempty"`
	Power                int64    `protobuf:"varint,2,opt,name=power,proto3" json:"power,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QbftNode) Reset()         { *m = QbftNode{} }
func (m *QbftNode) String() string { return proto.CompactTextString(m) }
func (*QbftNode) ProtoMessage()    {}
func (*QbftNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_df7368f483391686, []int{0}
}

func (m *QbftNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QbftNode.Unmarshal(m, b)
}
func (m *QbftNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QbftNode.Marshal(b, m, deterministic)
}
func (m *QbftNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QbftNode.Merge(m, src)
}
func (m *QbftNode) XXX_Size() int {
	return xxx_messageInfo_QbftNode.Size(m)
}
func (m *QbftNode) XXX_DiscardUnknown() {
	xxx_messageInfo_QbftNode.DiscardUnknown(m)
}

var xxx_messageInfo_QbftNode proto.InternalMessageInfo

func (m *QbftNode) GetPubKey() string {
	if m != nil {
		return m.PubKey
	}
	return ""
}

func (m *QbftNode) GetPower() int64 {
	if m != nil {
		return m.Power
	}
	return 0
}

type QbftNodes struct {
	Nodes                []*QbftNode `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *QbftNodes) Reset()         { *m = QbftNodes{} }
func (m *QbftNodes) String() string { return proto.CompactTextString(m) }
func (*QbftNodes) ProtoMessage()    {}
func (*QbftNodes) Descriptor() ([]byte, []int) {
	return fileDescriptor_df7368f483391686, []int{1}
}

func (m *QbftNodes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QbftNodes.Unmarshal(m, b)
}
func (m *QbftNodes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QbftNodes.Marshal(b, m, deterministic)
}
func (m *QbftNodes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QbftNodes.Merge(m, src)
}
func (m *QbftNodes) XXX_Size() int {
	return xxx_messageInfo_QbftNodes.Size(m)
}
func (m *QbftNodes) XXX_DiscardUnknown() {
	xxx_messageInfo_QbftNodes.DiscardUnknown(m)
}

var xxx_messageInfo_QbftNodes proto.InternalMessageInfo

func (m *QbftNodes) GetNodes() []*QbftNode {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type QbftNodeAction struct {
	// Types that are valid to be assigned to Value:
	//	*QbftNodeAction_Node
	//	*QbftNodeAction_BlockInfo
	Value                isQbftNodeAction_Value `protobuf_oneof:"value"`
	Ty                   int32                  `protobuf:"varint,3,opt,name=Ty,proto3" json:"Ty,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *QbftNodeAction) Reset()         { *m = QbftNodeAction{} }
func (m *QbftNodeAction) String() string { return proto.CompactTextString(m) }
func (*QbftNodeAction) ProtoMessage()    {}
func (*QbftNodeAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_df7368f483391686, []int{2}
}

func (m *QbftNodeAction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QbftNodeAction.Unmarshal(m, b)
}
func (m *QbftNodeAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QbftNodeAction.Marshal(b, m, deterministic)
}
func (m *QbftNodeAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QbftNodeAction.Merge(m, src)
}
func (m *QbftNodeAction) XXX_Size() int {
	return xxx_messageInfo_QbftNodeAction.Size(m)
}
func (m *QbftNodeAction) XXX_DiscardUnknown() {
	xxx_messageInfo_QbftNodeAction.DiscardUnknown(m)
}

var xxx_messageInfo_QbftNodeAction proto.InternalMessageInfo

type isQbftNodeAction_Value interface {
	isQbftNodeAction_Value()
}

type QbftNodeAction_Node struct {
	Node *QbftNode `protobuf:"bytes,1,opt,name=node,proto3,oneof"`
}

type QbftNodeAction_BlockInfo struct {
	BlockInfo *QbftBlockInfo `protobuf:"bytes,2,opt,name=blockInfo,proto3,oneof"`
}

func (*QbftNodeAction_Node) isQbftNodeAction_Value() {}

func (*QbftNodeAction_BlockInfo) isQbftNodeAction_Value() {}

func (m *QbftNodeAction) GetValue() isQbftNodeAction_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *QbftNodeAction) GetNode() *QbftNode {
	if x, ok := m.GetValue().(*QbftNodeAction_Node); ok {
		return x.Node
	}
	return nil
}

func (m *QbftNodeAction) GetBlockInfo() *QbftBlockInfo {
	if x, ok := m.GetValue().(*QbftNodeAction_BlockInfo); ok {
		return x.BlockInfo
	}
	return nil
}

func (m *QbftNodeAction) GetTy() int32 {
	if m != nil {
		return m.Ty
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*QbftNodeAction) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*QbftNodeAction_Node)(nil),
		(*QbftNodeAction_BlockInfo)(nil),
	}
}

type ReqQbftNodes struct {
	Height               int64    `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqQbftNodes) Reset()         { *m = ReqQbftNodes{} }
func (m *ReqQbftNodes) String() string { return proto.CompactTextString(m) }
func (*ReqQbftNodes) ProtoMessage()    {}
func (*ReqQbftNodes) Descriptor() ([]byte, []int) {
	return fileDescriptor_df7368f483391686, []int{3}
}

func (m *ReqQbftNodes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqQbftNodes.Unmarshal(m, b)
}
func (m *ReqQbftNodes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqQbftNodes.Marshal(b, m, deterministic)
}
func (m *ReqQbftNodes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqQbftNodes.Merge(m, src)
}
func (m *ReqQbftNodes) XXX_Size() int {
	return xxx_messageInfo_ReqQbftNodes.Size(m)
}
func (m *ReqQbftNodes) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqQbftNodes.DiscardUnknown(m)
}

var xxx_messageInfo_ReqQbftNodes proto.InternalMessageInfo

func (m *ReqQbftNodes) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

type ReqQbftBlockInfo struct {
	Height               int64    `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqQbftBlockInfo) Reset()         { *m = ReqQbftBlockInfo{} }
func (m *ReqQbftBlockInfo) String() string { return proto.CompactTextString(m) }
func (*ReqQbftBlockInfo) ProtoMessage()    {}
func (*ReqQbftBlockInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_df7368f483391686, []int{4}
}

func (m *ReqQbftBlockInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqQbftBlockInfo.Unmarshal(m, b)
}
func (m *ReqQbftBlockInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqQbftBlockInfo.Marshal(b, m, deterministic)
}
func (m *ReqQbftBlockInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqQbftBlockInfo.Merge(m, src)
}
func (m *ReqQbftBlockInfo) XXX_Size() int {
	return xxx_messageInfo_ReqQbftBlockInfo.Size(m)
}
func (m *ReqQbftBlockInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqQbftBlockInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReqQbftBlockInfo proto.InternalMessageInfo

func (m *ReqQbftBlockInfo) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

type QbftNodeInfo struct {
	NodeIP               string   `protobuf:"bytes,1,opt,name=nodeIP,proto3" json:"nodeIP,omitempty"`
	NodeID               string   `protobuf:"bytes,2,opt,name=nodeID,proto3" json:"nodeID,omitempty"`
	Address              string   `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	PubKey               string   `protobuf:"bytes,4,opt,name=pubKey,proto3" json:"pubKey,omitempty"`
	VotingPower          int64    `protobuf:"varint,5,opt,name=votingPower,proto3" json:"votingPower,omitempty"`
	Accum                int64    `protobuf:"varint,6,opt,name=accum,proto3" json:"accum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QbftNodeInfo) Reset()         { *m = QbftNodeInfo{} }
func (m *QbftNodeInfo) String() string { return proto.CompactTextString(m) }
func (*QbftNodeInfo) ProtoMessage()    {}
func (*QbftNodeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_df7368f483391686, []int{5}
}

func (m *QbftNodeInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QbftNodeInfo.Unmarshal(m, b)
}
func (m *QbftNodeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QbftNodeInfo.Marshal(b, m, deterministic)
}
func (m *QbftNodeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QbftNodeInfo.Merge(m, src)
}
func (m *QbftNodeInfo) XXX_Size() int {
	return xxx_messageInfo_QbftNodeInfo.Size(m)
}
func (m *QbftNodeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_QbftNodeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_QbftNodeInfo proto.InternalMessageInfo

func (m *QbftNodeInfo) GetNodeIP() string {
	if m != nil {
		return m.NodeIP
	}
	return ""
}

func (m *QbftNodeInfo) GetNodeID() string {
	if m != nil {
		return m.NodeID
	}
	return ""
}

func (m *QbftNodeInfo) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *QbftNodeInfo) GetPubKey() string {
	if m != nil {
		return m.PubKey
	}
	return ""
}

func (m *QbftNodeInfo) GetVotingPower() int64 {
	if m != nil {
		return m.VotingPower
	}
	return 0
}

func (m *QbftNodeInfo) GetAccum() int64 {
	if m != nil {
		return m.Accum
	}
	return 0
}

type QbftNodeInfoSet struct {
	Nodes                []*QbftNodeInfo `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *QbftNodeInfoSet) Reset()         { *m = QbftNodeInfoSet{} }
func (m *QbftNodeInfoSet) String() string { return proto.CompactTextString(m) }
func (*QbftNodeInfoSet) ProtoMessage()    {}
func (*QbftNodeInfoSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_df7368f483391686, []int{6}
}

func (m *QbftNodeInfoSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QbftNodeInfoSet.Unmarshal(m, b)
}
func (m *QbftNodeInfoSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QbftNodeInfoSet.Marshal(b, m, deterministic)
}
func (m *QbftNodeInfoSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QbftNodeInfoSet.Merge(m, src)
}
func (m *QbftNodeInfoSet) XXX_Size() int {
	return xxx_messageInfo_QbftNodeInfoSet.Size(m)
}
func (m *QbftNodeInfoSet) XXX_DiscardUnknown() {
	xxx_messageInfo_QbftNodeInfoSet.DiscardUnknown(m)
}

var xxx_messageInfo_QbftNodeInfoSet proto.InternalMessageInfo

func (m *QbftNodeInfoSet) GetNodes() []*QbftNodeInfo {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type QbftPerfStat struct {
	TotalTx              int64    `protobuf:"varint,1,opt,name=totalTx,proto3" json:"totalTx,omitempty"`
	TotalBlock           int64    `protobuf:"varint,2,opt,name=totalBlock,proto3" json:"totalBlock,omitempty"`
	TxPerBlock           int64    `protobuf:"varint,3,opt,name=txPerBlock,proto3" json:"txPerBlock,omitempty"`
	TotalSecond          int64    `protobuf:"varint,4,opt,name=totalSecond,proto3" json:"totalSecond,omitempty"`
	TxPerSecond          int64    `protobuf:"varint,5,opt,name=txPerSecond,proto3" json:"txPerSecond,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QbftPerfStat) Reset()         { *m = QbftPerfStat{} }
func (m *QbftPerfStat) String() string { return proto.CompactTextString(m) }
func (*QbftPerfStat) ProtoMessage()    {}
func (*QbftPerfStat) Descriptor() ([]byte, []int) {
	return fileDescriptor_df7368f483391686, []int{7}
}

func (m *QbftPerfStat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QbftPerfStat.Unmarshal(m, b)
}
func (m *QbftPerfStat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QbftPerfStat.Marshal(b, m, deterministic)
}
func (m *QbftPerfStat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QbftPerfStat.Merge(m, src)
}
func (m *QbftPerfStat) XXX_Size() int {
	return xxx_messageInfo_QbftPerfStat.Size(m)
}
func (m *QbftPerfStat) XXX_DiscardUnknown() {
	xxx_messageInfo_QbftPerfStat.DiscardUnknown(m)
}

var xxx_messageInfo_QbftPerfStat proto.InternalMessageInfo

func (m *QbftPerfStat) GetTotalTx() int64 {
	if m != nil {
		return m.TotalTx
	}
	return 0
}

func (m *QbftPerfStat) GetTotalBlock() int64 {
	if m != nil {
		return m.TotalBlock
	}
	return 0
}

func (m *QbftPerfStat) GetTxPerBlock() int64 {
	if m != nil {
		return m.TxPerBlock
	}
	return 0
}

func (m *QbftPerfStat) GetTotalSecond() int64 {
	if m != nil {
		return m.TotalSecond
	}
	return 0
}

func (m *QbftPerfStat) GetTxPerSecond() int64 {
	if m != nil {
		return m.TxPerSecond
	}
	return 0
}

type ReqQbftPerfStat struct {
	Start                int64    `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"`
	End                  int64    `protobuf:"varint,2,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqQbftPerfStat) Reset()         { *m = ReqQbftPerfStat{} }
func (m *ReqQbftPerfStat) String() string { return proto.CompactTextString(m) }
func (*ReqQbftPerfStat) ProtoMessage()    {}
func (*ReqQbftPerfStat) Descriptor() ([]byte, []int) {
	return fileDescriptor_df7368f483391686, []int{8}
}

func (m *ReqQbftPerfStat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqQbftPerfStat.Unmarshal(m, b)
}
func (m *ReqQbftPerfStat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqQbftPerfStat.Marshal(b, m, deterministic)
}
func (m *ReqQbftPerfStat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqQbftPerfStat.Merge(m, src)
}
func (m *ReqQbftPerfStat) XXX_Size() int {
	return xxx_messageInfo_ReqQbftPerfStat.Size(m)
}
func (m *ReqQbftPerfStat) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqQbftPerfStat.DiscardUnknown(m)
}

var xxx_messageInfo_ReqQbftPerfStat proto.InternalMessageInfo

func (m *ReqQbftPerfStat) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *ReqQbftPerfStat) GetEnd() int64 {
	if m != nil {
		return m.End
	}
	return 0
}

func init() {
	proto.RegisterType((*QbftNode)(nil), "types.QbftNode")
	proto.RegisterType((*QbftNodes)(nil), "types.QbftNodes")
	proto.RegisterType((*QbftNodeAction)(nil), "types.QbftNodeAction")
	proto.RegisterType((*ReqQbftNodes)(nil), "types.ReqQbftNodes")
	proto.RegisterType((*ReqQbftBlockInfo)(nil), "types.ReqQbftBlockInfo")
	proto.RegisterType((*QbftNodeInfo)(nil), "types.QbftNodeInfo")
	proto.RegisterType((*QbftNodeInfoSet)(nil), "types.QbftNodeInfoSet")
	proto.RegisterType((*QbftPerfStat)(nil), "types.QbftPerfStat")
	proto.RegisterType((*ReqQbftPerfStat)(nil), "types.ReqQbftPerfStat")
}

func init() {
	proto.RegisterFile("qbftNode.proto", fileDescriptor_df7368f483391686)
}

var fileDescriptor_df7368f483391686 = []byte{
	// 488 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0xd1, 0x8e, 0xd2, 0x40,
	0x14, 0xa5, 0x74, 0xdb, 0x5d, 0x2e, 0x08, 0x9b, 0x91, 0x90, 0x86, 0x07, 0xd3, 0x34, 0x59, 0x83,
	0x3e, 0x60, 0x82, 0xc6, 0x68, 0xe2, 0x8b, 0x9b, 0x4d, 0x84, 0x98, 0x6c, 0x70, 0xca, 0x0f, 0x94,
	0xf6, 0xb2, 0x10, 0x4b, 0x07, 0xda, 0x61, 0xdd, 0xfe, 0x80, 0xdf, 0xe2, 0x83, 0x1f, 0x69, 0xe6,
	0x76, 0x06, 0xba, 0xab, 0xbe, 0xcd, 0x39, 0xf7, 0xdc, 0xce, 0xb9, 0xf7, 0x4c, 0xa1, 0xbb, 0x5f,
	0xae, 0xe4, 0xad, 0x48, 0x70, 0xbc, 0xcb, 0x85, 0x14, 0xcc, 0x91, 0xe5, 0x0e, 0x8b, 0x61, 0x27,
	0x16, 0xdb, 0xad, 0xc8, 0x2a, 0x72, 0x08, 0x4a, 0x54, 0x9d, 0x83, 0x0f, 0x70, 0xf1, 0x4d, 0xb7,
	0xb0, 0x01, 0xb8, 0xbb, 0xc3, 0xf2, 0x2b, 0x96, 0x9e, 0xe5, 0x5b, 0xa3, 0x16, 0xd7, 0x88, 0xf5,
	0xc1, 0xd9, 0x89, 0x1f, 0x98, 0x7b, 0x4d, 0xdf, 0x1a, 0xd9, 0xbc, 0x02, 0xc1, 0x04, 0x5a, 0xa6,
	0xb3, 0x60, 0x57, 0xe0, 0x64, 0xea, 0xe0, 0x59, 0xbe, 0x3d, 0x6a, 0x4f, 0x7a, 0x63, 0xba, 0x77,
	0x6c, 0x04, 0xbc, 0xaa, 0x06, 0x3f, 0x2d, 0xe8, 0x1a, 0xee, 0x73, 0x2c, 0x37, 0x22, 0x63, 0x57,
	0x70, 0xa6, 0x6a, 0x74, 0xe5, 0xdf, 0x8d, 0xd3, 0x06, 0xa7, 0x32, 0x7b, 0x07, 0xad, 0x65, 0x2a,
	0xe2, 0xef, 0xb3, 0x6c, 0x25, 0xc8, 0x47, 0x7b, 0xd2, 0xaf, 0x69, 0xaf, 0x4d, 0x6d, 0xda, 0xe0,
	0x27, 0x21, 0xeb, 0x42, 0x73, 0x51, 0x7a, 0xb6, 0x6f, 0x8d, 0x1c, 0xde, 0x5c, 0x94, 0xd7, 0xe7,
	0xe0, 0xdc, 0x47, 0xe9, 0x01, 0x83, 0x97, 0xd0, 0xe1, 0xb8, 0x3f, 0xf9, 0x1f, 0x80, 0xbb, 0xc6,
	0xcd, 0xdd, 0x5a, 0x92, 0x0f, 0x9b, 0x6b, 0x14, 0xbc, 0x86, 0x4b, 0xad, 0x3b, 0xde, 0xf0, 0x5f,
	0xed, 0x2f, 0x0b, 0x3a, 0xe6, 0x8b, 0x46, 0xa8, 0xbc, 0xcf, 0xe6, 0x66, 0x9f, 0x15, 0x3a, 0xf2,
	0x37, 0x34, 0x88, 0xe1, 0x6f, 0x98, 0x07, 0xe7, 0x51, 0x92, 0xe4, 0x58, 0x14, 0x64, 0xb9, 0xc5,
	0x0d, 0xac, 0x25, 0x73, 0xf6, 0x28, 0x19, 0x1f, 0xda, 0xf7, 0x42, 0x6e, 0xb2, 0xbb, 0x39, 0xe5,
	0xe3, 0x90, 0x9f, 0x3a, 0xa5, 0xb2, 0x8b, 0xe2, 0xf8, 0xb0, 0xf5, 0xdc, 0x2a, 0x3b, 0x02, 0xc1,
	0x27, 0xe8, 0xd5, 0x9d, 0x86, 0x28, 0xd9, 0xab, 0xc7, 0x09, 0x3e, 0x7f, 0x12, 0x84, 0x92, 0x99,
	0x14, 0x7f, 0xeb, 0x41, 0xe7, 0x98, 0xaf, 0x42, 0x19, 0x49, 0x65, 0x5c, 0x0a, 0x19, 0xa5, 0x8b,
	0x07, 0xbd, 0x12, 0x03, 0xd9, 0x0b, 0x00, 0x3a, 0xd2, 0xf6, 0xf4, 0xfb, 0xa9, 0x31, 0x54, 0x7f,
	0x98, 0x63, 0x5e, 0xd5, 0x6d, 0x5d, 0x3f, 0x32, 0x6a, 0x40, 0x52, 0x87, 0x18, 0x8b, 0x2c, 0xa1,
	0xe9, 0x6d, 0x5e, 0xa7, 0x48, 0xa1, 0xf4, 0x5a, 0xa1, 0x57, 0x50, 0xa3, 0x82, 0x8f, 0xd0, 0xd3,
	0x19, 0x1e, 0x0d, 0xf7, 0xc1, 0x29, 0x64, 0x94, 0x9b, 0x04, 0x2b, 0xc0, 0x2e, 0xc1, 0xc6, 0x2c,
	0xd1, 0x2e, 0xd5, 0x71, 0x52, 0xc0, 0x85, 0xf9, 0xa1, 0xd8, 0x1b, 0x70, 0x67, 0x45, 0x58, 0x66,
	0x31, 0x7b, 0xa6, 0x77, 0xc3, 0x71, 0x7f, 0xbb, 0x49, 0x87, 0xf5, 0x77, 0x38, 0x2b, 0xa6, 0x18,
	0xa5, 0x72, 0x5d, 0x06, 0x0d, 0xf6, 0x1e, 0xda, 0x5f, 0xf0, 0xf4, 0x1a, 0x9e, 0x74, 0x0d, 0xfe,
	0xb1, 0xe0, 0x10, 0x65, 0xd0, 0x58, 0xba, 0xf4, 0x67, 0xbe, 0xfd, 0x13, 0x00, 0x00, 0xff, 0xff,
	0x8a, 0xd1, 0x4b, 0x13, 0xcc, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// QbftNodeClient is the client API for QbftNode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QbftNodeClient interface {
	IsSync(ctx context.Context, in *types.ReqNil, opts ...grpc.CallOption) (*QbftIsHealthy, error)
	GetNodeInfo(ctx context.Context, in *types.ReqNil, opts ...grpc.CallOption) (*QbftNodeInfoSet, error)
}

type qbftNodeClient struct {
	cc grpc.ClientConnInterface
}

func NewQbftNodeClient(cc grpc.ClientConnInterface) QbftNodeClient {
	return &qbftNodeClient{cc}
}

func (c *qbftNodeClient) IsSync(ctx context.Context, in *types.ReqNil, opts ...grpc.CallOption) (*QbftIsHealthy, error) {
	out := new(QbftIsHealthy)
	err := c.cc.Invoke(ctx, "/types.qbftNode/IsSync", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qbftNodeClient) GetNodeInfo(ctx context.Context, in *types.ReqNil, opts ...grpc.CallOption) (*QbftNodeInfoSet, error) {
	out := new(QbftNodeInfoSet)
	err := c.cc.Invoke(ctx, "/types.qbftNode/GetNodeInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QbftNodeServer is the server API for QbftNode service.
type QbftNodeServer interface {
	IsSync(context.Context, *types.ReqNil) (*QbftIsHealthy, error)
	GetNodeInfo(context.Context, *types.ReqNil) (*QbftNodeInfoSet, error)
}

// UnimplementedQbftNodeServer can be embedded to have forward compatible implementations.
type UnimplementedQbftNodeServer struct {
}

func (*UnimplementedQbftNodeServer) IsSync(ctx context.Context, req *types.ReqNil) (*QbftIsHealthy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsSync not implemented")
}
func (*UnimplementedQbftNodeServer) GetNodeInfo(ctx context.Context, req *types.ReqNil) (*QbftNodeInfoSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNodeInfo not implemented")
}

func RegisterQbftNodeServer(s *grpc.Server, srv QbftNodeServer) {
	s.RegisterService(&_QbftNode_serviceDesc, srv)
}

func _QbftNode_IsSync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.ReqNil)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QbftNodeServer).IsSync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.qbftNode/IsSync",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QbftNodeServer).IsSync(ctx, req.(*types.ReqNil))
	}
	return interceptor(ctx, in, info, handler)
}

func _QbftNode_GetNodeInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.ReqNil)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QbftNodeServer).GetNodeInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.qbftNode/GetNodeInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QbftNodeServer).GetNodeInfo(ctx, req.(*types.ReqNil))
	}
	return interceptor(ctx, in, info, handler)
}

var _QbftNode_serviceDesc = grpc.ServiceDesc{
	ServiceName: "types.qbftNode",
	HandlerType: (*QbftNodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsSync",
			Handler:    _QbftNode_IsSync_Handler,
		},
		{
			MethodName: "GetNodeInfo",
			Handler:    _QbftNode_GetNodeInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "qbftNode.proto",
}
