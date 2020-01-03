// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/gov/types/codec.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// VoteOption defines a vote option
type VoteOption int32

const (
	OptionEmpty      VoteOption = 0
	OptionYes        VoteOption = 1
	OptionAbstain    VoteOption = 2
	OptionNo         VoteOption = 3
	OptionNoWithVeto VoteOption = 4
)

var VoteOption_name = map[int32]string{
	0: "EMPTY",
	1: "YES",
	2: "ABSTAIN",
	3: "NO",
	4: "NO_WITH_VETO",
}

var VoteOption_value = map[string]int32{
	"EMPTY":        0,
	"YES":          1,
	"ABSTAIN":      2,
	"NO":           3,
	"NO_WITH_VETO": 4,
}

func (VoteOption) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4ed4b2a5a30fa918, []int{0}
}

type MsgBase struct {
	// Types that are valid to be assigned to Sum:
	//	*MsgBase_GovDeposit
	//	*MsgBase_GovVote
	Sum isMsgBase_Sum `protobuf_oneof:"sum"`
}

func (m *MsgBase) Reset()      { *m = MsgBase{} }
func (*MsgBase) ProtoMessage() {}
func (*MsgBase) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ed4b2a5a30fa918, []int{0}
}
func (m *MsgBase) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgBase) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgBase.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgBase) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgBase.Merge(m, src)
}
func (m *MsgBase) XXX_Size() int {
	return m.Size()
}
func (m *MsgBase) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgBase.DiscardUnknown(m)
}

var xxx_messageInfo_MsgBase proto.InternalMessageInfo

type isMsgBase_Sum interface {
	isMsgBase_Sum()
	MarshalTo([]byte) (int, error)
	Size() int
}

type MsgBase_GovDeposit struct {
	GovDeposit *MsgDeposit `protobuf:"bytes,1,opt,name=gov_deposit,json=govDeposit,proto3,oneof" json:"gov_deposit,omitempty"`
}
type MsgBase_GovVote struct {
	GovVote *MsgVote `protobuf:"bytes,2,opt,name=gov_vote,json=govVote,proto3,oneof" json:"gov_vote,omitempty"`
}

func (*MsgBase_GovDeposit) isMsgBase_Sum() {}
func (*MsgBase_GovVote) isMsgBase_Sum()    {}

func (m *MsgBase) GetSum() isMsgBase_Sum {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (m *MsgBase) GetGovDeposit() *MsgDeposit {
	if x, ok := m.GetSum().(*MsgBase_GovDeposit); ok {
		return x.GovDeposit
	}
	return nil
}

func (m *MsgBase) GetGovVote() *MsgVote {
	if x, ok := m.GetSum().(*MsgBase_GovVote); ok {
		return x.GovVote
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MsgBase) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MsgBase_GovDeposit)(nil),
		(*MsgBase_GovVote)(nil),
	}
}

type MsgSubmitProposalBase struct {
	IntialDeposit []types.Coin                                  `protobuf:"bytes,1,rep,name=intial_deposit,json=intialDeposit,proto3" json:"intial_deposit"`
	Proposer      github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,2,opt,name=proposer,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"proposer,omitempty"`
}

func (m *MsgSubmitProposalBase) Reset()      { *m = MsgSubmitProposalBase{} }
func (*MsgSubmitProposalBase) ProtoMessage() {}
func (*MsgSubmitProposalBase) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ed4b2a5a30fa918, []int{1}
}
func (m *MsgSubmitProposalBase) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitProposalBase) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitProposalBase.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitProposalBase) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitProposalBase.Merge(m, src)
}
func (m *MsgSubmitProposalBase) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitProposalBase) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitProposalBase.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitProposalBase proto.InternalMessageInfo

// MsgDeposit defines a message to submit a deposit to an existing proposal
type MsgDeposit struct {
	ProposalID uint64                                        `protobuf:"varint,1,opt,name=proposal_id,json=proposalId,proto3" json:"proposal_id,omitempty"`
	Depositor  github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,2,opt,name=depositor,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"depositor,omitempty"`
	Amount     []types.Coin                                  `protobuf:"bytes,3,rep,name=amount,proto3" json:"amount"`
}

func (m *MsgDeposit) Reset()      { *m = MsgDeposit{} }
func (*MsgDeposit) ProtoMessage() {}
func (*MsgDeposit) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ed4b2a5a30fa918, []int{2}
}
func (m *MsgDeposit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeposit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeposit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeposit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeposit.Merge(m, src)
}
func (m *MsgDeposit) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeposit) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeposit.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeposit proto.InternalMessageInfo

// MsgVote defines a message to cast a vote
type MsgVote struct {
	ProposalID uint64                                        `protobuf:"varint,1,opt,name=proposal_id,json=proposalId,proto3" json:"proposal_id,omitempty"`
	Voter      github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,2,opt,name=voter,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"voter,omitempty"`
	Option     VoteOption                                    `protobuf:"varint,3,opt,name=option,proto3,enum=cosmos_sdk.x.gov.v1.VoteOption" json:"option,omitempty"`
}

func (m *MsgVote) Reset()      { *m = MsgVote{} }
func (*MsgVote) ProtoMessage() {}
func (*MsgVote) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ed4b2a5a30fa918, []int{3}
}
func (m *MsgVote) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgVote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgVote.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgVote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgVote.Merge(m, src)
}
func (m *MsgVote) XXX_Size() int {
	return m.Size()
}
func (m *MsgVote) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgVote.DiscardUnknown(m)
}

var xxx_messageInfo_MsgVote proto.InternalMessageInfo

// TextProposal defines a standard text proposal whose changes need to be
// manually updated in case of approval
type TextProposal struct {
	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (m *TextProposal) Reset()      { *m = TextProposal{} }
func (*TextProposal) ProtoMessage() {}
func (*TextProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ed4b2a5a30fa918, []int{4}
}
func (m *TextProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TextProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TextProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TextProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TextProposal.Merge(m, src)
}
func (m *TextProposal) XXX_Size() int {
	return m.Size()
}
func (m *TextProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_TextProposal.DiscardUnknown(m)
}

var xxx_messageInfo_TextProposal proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("cosmos_sdk.x.gov.v1.VoteOption", VoteOption_name, VoteOption_value)
	proto.RegisterType((*MsgBase)(nil), "cosmos_sdk.x.gov.v1.MsgBase")
	proto.RegisterType((*MsgSubmitProposalBase)(nil), "cosmos_sdk.x.gov.v1.MsgSubmitProposalBase")
	proto.RegisterType((*MsgDeposit)(nil), "cosmos_sdk.x.gov.v1.MsgDeposit")
	proto.RegisterType((*MsgVote)(nil), "cosmos_sdk.x.gov.v1.MsgVote")
	proto.RegisterType((*TextProposal)(nil), "cosmos_sdk.x.gov.v1.TextProposal")
}

func init() { proto.RegisterFile("x/gov/types/codec.proto", fileDescriptor_4ed4b2a5a30fa918) }

var fileDescriptor_4ed4b2a5a30fa918 = []byte{
	// 678 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xbd, 0x4f, 0xdb, 0x5e,
	0x14, 0xb5, 0xf3, 0xc1, 0xc7, 0x4d, 0xe0, 0x17, 0x1e, 0xfc, 0xda, 0xc8, 0xaa, 0x1c, 0x2b, 0x03,
	0x8a, 0x2a, 0xc5, 0x56, 0x68, 0xa5, 0x0a, 0xa6, 0xc6, 0x25, 0x2d, 0x19, 0x92, 0xa0, 0x10, 0x81,
	0xe8, 0x12, 0x25, 0xf6, 0xab, 0xb1, 0x20, 0x7e, 0x96, 0xdf, 0x8b, 0x0b, 0x1b, 0x63, 0x95, 0xa9,
	0xff, 0x40, 0xa6, 0x32, 0x74, 0xa8, 0xaa, 0x0e, 0x1d, 0x2b, 0x55, 0xdd, 0x50, 0x27, 0x46, 0x86,
	0x0a, 0x95, 0xf0, 0x5f, 0x74, 0xaa, 0xec, 0x97, 0x40, 0x24, 0x4a, 0x3f, 0xc4, 0xe6, 0xeb, 0x7b,
	0xce, 0x79, 0xe7, 0xdc, 0x77, 0xf5, 0xe0, 0xee, 0xbe, 0x66, 0x11, 0x5f, 0x63, 0x07, 0x2e, 0xa6,
	0x9a, 0x41, 0x4c, 0x6c, 0xa8, 0xae, 0x47, 0x18, 0x41, 0xf3, 0x06, 0xa1, 0x1d, 0x42, 0x9b, 0xd4,
	0xdc, 0x55, 0xf7, 0x55, 0x8b, 0xf8, 0xaa, 0x5f, 0x90, 0xe6, 0xae, 0xe1, 0xa4, 0x87, 0x3e, 0x76,
	0x4c, 0xe2, 0x69, 0x96, 0xcd, 0x76, 0xba, 0x6d, 0xd5, 0x20, 0x1d, 0xcd, 0x22, 0x16, 0xd1, 0xc2,
	0x6e, 0xbb, 0xfb, 0x22, 0xac, 0xc2, 0x22, 0xfc, 0x1a, 0xb2, 0x96, 0xaf, 0xb3, 0x3c, 0x6c, 0x61,
	0x27, 0xef, 0x60, 0xf6, 0x92, 0x78, 0xbb, 0x1a, 0x3f, 0x3d, 0xcf, 0x89, 0xbc, 0xe0, 0xd4, 0xec,
	0x67, 0x11, 0x26, 0x2b, 0xd4, 0xd2, 0x5b, 0x14, 0x23, 0x1d, 0x12, 0x16, 0xf1, 0x9b, 0x26, 0x76,
	0x09, 0xb5, 0x59, 0x5a, 0x54, 0xc4, 0x5c, 0x62, 0x29, 0xa3, 0xfe, 0xc2, 0xba, 0x5a, 0xa1, 0xd6,
	0x2a, 0x87, 0xad, 0x09, 0x75, 0xb0, 0x88, 0x3f, 0xac, 0xd0, 0x32, 0x4c, 0x05, 0x1a, 0x3e, 0x61,
	0x38, 0x1d, 0x09, 0x05, 0xee, 0xdd, 0x24, 0xb0, 0x49, 0x18, 0x5e, 0x13, 0xea, 0x93, 0x16, 0xf1,
	0x83, 0xcf, 0x15, 0xf5, 0xf0, 0x9b, 0x22, 0x7e, 0xfd, 0x98, 0x5f, 0x1c, 0x8b, 0xc1, 0xa9, 0x23,
	0xff, 0xd4, 0xdc, 0xe5, 0xb3, 0x0d, 0xe8, 0x7a, 0x1c, 0xa2, 0xb4, 0xdb, 0xc9, 0xbe, 0x17, 0xe1,
	0xff, 0x0a, 0xb5, 0x36, 0xba, 0xed, 0x8e, 0xcd, 0xd6, 0x3d, 0xe2, 0x12, 0xda, 0xda, 0x0b, 0xf3,
	0x3c, 0x86, 0x59, 0xdb, 0x61, 0x76, 0x6b, 0x6f, 0x2c, 0x52, 0x34, 0x97, 0x58, 0x9a, 0x1f, 0x77,
	0xe4, 0x17, 0xd4, 0x27, 0xc4, 0x76, 0xf4, 0xd8, 0xf1, 0x59, 0x46, 0xa8, 0xcf, 0x70, 0xc2, 0x28,
	0x4d, 0x05, 0xa6, 0xdc, 0x50, 0x11, 0x7b, 0x61, 0x9a, 0xa4, 0x5e, 0xf8, 0x71, 0x96, 0xc9, 0xff,
	0x85, 0xc1, 0xa2, 0x61, 0x14, 0x4d, 0xd3, 0xc3, 0x94, 0xd6, 0x2f, 0x25, 0x56, 0x62, 0x41, 0xc2,
	0xec, 0x17, 0x11, 0xe0, 0x6a, 0x7e, 0x48, 0x83, 0x84, 0x3b, 0x74, 0xdd, 0xb4, 0xcd, 0x70, 0xea,
	0x31, 0x7d, 0x76, 0x70, 0x96, 0x81, 0x51, 0x98, 0xf2, 0x6a, 0x1d, 0x46, 0x90, 0xb2, 0x89, 0x6a,
	0x30, 0x3d, 0xcc, 0x43, 0x6e, 0xe1, 0xea, 0x4a, 0x03, 0x15, 0x60, 0xa2, 0xd5, 0x21, 0x5d, 0x87,
	0xa5, 0xa3, 0x7f, 0x9a, 0xcf, 0x10, 0x98, 0xfd, 0xc4, 0xd7, 0x26, 0xb8, 0xb7, 0x7f, 0x0f, 0xf0,
	0x0c, 0xe2, 0xc1, 0x7e, 0xdc, 0xc2, 0x3c, 0xe7, 0xa3, 0x47, 0x30, 0x41, 0x5c, 0x66, 0x13, 0x27,
	0x1d, 0x55, 0xc4, 0xdc, 0xec, 0x0d, 0xbb, 0x1a, 0x98, 0xac, 0x85, 0xb0, 0xfa, 0x10, 0x9e, 0x7d,
	0x0a, 0xc9, 0x06, 0xde, 0xbf, 0xdc, 0x16, 0xb4, 0x00, 0x71, 0x66, 0xb3, 0x3d, 0x1c, 0x9a, 0x9f,
	0xae, 0xf3, 0x02, 0x29, 0x90, 0x30, 0x31, 0x35, 0x3c, 0x9b, 0x9f, 0x11, 0x09, 0x7b, 0xe3, 0xbf,
	0xee, 0xbf, 0x13, 0x01, 0xae, 0xe4, 0x91, 0x04, 0xf1, 0x52, 0x65, 0xbd, 0xb1, 0x9d, 0x12, 0xa4,
	0xff, 0x7a, 0x7d, 0x25, 0xc1, 0x7f, 0x97, 0x3a, 0x2e, 0x3b, 0x40, 0x77, 0x20, 0xba, 0x5d, 0xda,
	0x48, 0x89, 0xd2, 0x4c, 0xaf, 0xaf, 0x4c, 0xf3, 0xce, 0x36, 0xa6, 0x48, 0x86, 0xc9, 0xa2, 0xbe,
	0xd1, 0x28, 0x96, 0xab, 0xa9, 0x88, 0x34, 0xd7, 0xeb, 0x2b, 0x33, 0xbc, 0x57, 0x6c, 0x53, 0xd6,
	0xb2, 0x1d, 0xb4, 0x00, 0x91, 0x6a, 0x2d, 0x15, 0x95, 0x92, 0xbd, 0xbe, 0x32, 0xc5, 0x5b, 0x55,
	0x82, 0x16, 0x21, 0x59, 0xad, 0x35, 0xb7, 0xca, 0x8d, 0xb5, 0xe6, 0x66, 0xa9, 0x51, 0x4b, 0xc5,
	0xa4, 0x85, 0x5e, 0x5f, 0x49, 0x8d, 0xfa, 0x5b, 0x36, 0xdb, 0xd9, 0xc4, 0x8c, 0x48, 0xc9, 0x57,
	0x6f, 0x64, 0xe1, 0xed, 0x91, 0x2c, 0x7c, 0x38, 0x92, 0x05, 0xbd, 0x7a, 0x7c, 0x2e, 0x0b, 0xa7,
	0xe7, 0xb2, 0x70, 0x38, 0x90, 0x85, 0xe3, 0x81, 0x2c, 0x9e, 0x0c, 0x64, 0xf1, 0xfb, 0x40, 0x16,
	0x5f, 0x5f, 0xc8, 0xc2, 0xc9, 0x85, 0x2c, 0x9c, 0x5e, 0xc8, 0xc2, 0xf3, 0xdc, 0x6f, 0xef, 0x62,
	0xec, 0x85, 0x6b, 0x4f, 0x84, 0x6f, 0xc8, 0x83, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x69, 0x53,
	0xbf, 0xb1, 0xf7, 0x04, 0x00, 0x00,
}

func (this *MsgBase) GetMsg() github_com_cosmos_cosmos_sdk_types.Msg {
	if x := this.GetGovDeposit(); x != nil {
		return x
	}
	if x := this.GetGovVote(); x != nil {
		return x
	}
	return nil
}

func (this *MsgBase) SetMsg(value github_com_cosmos_cosmos_sdk_types.Msg) error {
	switch vt := value.(type) {
	case *MsgDeposit:
		this.Sum = &MsgBase_GovDeposit{vt}
		return nil
	case *MsgVote:
		this.Sum = &MsgBase_GovVote{vt}
		return nil
	}
	return fmt.Errorf("can't encode value of type %T as message MsgBase", value)
}

func (m *MsgBase) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgBase) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgBase) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Sum != nil {
		{
			size := m.Sum.Size()
			i -= size
			if _, err := m.Sum.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *MsgBase_GovDeposit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgBase_GovDeposit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.GovDeposit != nil {
		{
			size, err := m.GovDeposit.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCodec(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *MsgBase_GovVote) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgBase_GovVote) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.GovVote != nil {
		{
			size, err := m.GovVote.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCodec(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *MsgSubmitProposalBase) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitProposalBase) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitProposalBase) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Proposer) > 0 {
		i -= len(m.Proposer)
		copy(dAtA[i:], m.Proposer)
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Proposer)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.IntialDeposit) > 0 {
		for iNdEx := len(m.IntialDeposit) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.IntialDeposit[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCodec(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *MsgDeposit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeposit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeposit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCodec(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Depositor) > 0 {
		i -= len(m.Depositor)
		copy(dAtA[i:], m.Depositor)
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Depositor)))
		i--
		dAtA[i] = 0x12
	}
	if m.ProposalID != 0 {
		i = encodeVarintCodec(dAtA, i, uint64(m.ProposalID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgVote) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgVote) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgVote) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Option != 0 {
		i = encodeVarintCodec(dAtA, i, uint64(m.Option))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Voter) > 0 {
		i -= len(m.Voter)
		copy(dAtA[i:], m.Voter)
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Voter)))
		i--
		dAtA[i] = 0x12
	}
	if m.ProposalID != 0 {
		i = encodeVarintCodec(dAtA, i, uint64(m.ProposalID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *TextProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TextProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TextProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCodec(dAtA []byte, offset int, v uint64) int {
	offset -= sovCodec(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgBase) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sum != nil {
		n += m.Sum.Size()
	}
	return n
}

func (m *MsgBase_GovDeposit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.GovDeposit != nil {
		l = m.GovDeposit.Size()
		n += 1 + l + sovCodec(uint64(l))
	}
	return n
}
func (m *MsgBase_GovVote) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.GovVote != nil {
		l = m.GovVote.Size()
		n += 1 + l + sovCodec(uint64(l))
	}
	return n
}
func (m *MsgSubmitProposalBase) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.IntialDeposit) > 0 {
		for _, e := range m.IntialDeposit {
			l = e.Size()
			n += 1 + l + sovCodec(uint64(l))
		}
	}
	l = len(m.Proposer)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	return n
}

func (m *MsgDeposit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ProposalID != 0 {
		n += 1 + sovCodec(uint64(m.ProposalID))
	}
	l = len(m.Depositor)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovCodec(uint64(l))
		}
	}
	return n
}

func (m *MsgVote) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ProposalID != 0 {
		n += 1 + sovCodec(uint64(m.ProposalID))
	}
	l = len(m.Voter)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	if m.Option != 0 {
		n += 1 + sovCodec(uint64(m.Option))
	}
	return n
}

func (m *TextProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	return n
}

func sovCodec(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCodec(x uint64) (n int) {
	return sovCodec(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *MsgBase) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&MsgBase{`,
		`Sum:` + fmt.Sprintf("%v", this.Sum) + `,`,
		`}`,
	}, "")
	return s
}
func (this *MsgBase_GovDeposit) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&MsgBase_GovDeposit{`,
		`GovDeposit:` + strings.Replace(fmt.Sprintf("%v", this.GovDeposit), "MsgDeposit", "MsgDeposit", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *MsgBase_GovVote) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&MsgBase_GovVote{`,
		`GovVote:` + strings.Replace(fmt.Sprintf("%v", this.GovVote), "MsgVote", "MsgVote", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *MsgSubmitProposalBase) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForIntialDeposit := "[]Coin{"
	for _, f := range this.IntialDeposit {
		repeatedStringForIntialDeposit += fmt.Sprintf("%v", f) + ","
	}
	repeatedStringForIntialDeposit += "}"
	s := strings.Join([]string{`&MsgSubmitProposalBase{`,
		`IntialDeposit:` + repeatedStringForIntialDeposit + `,`,
		`Proposer:` + fmt.Sprintf("%v", this.Proposer) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringCodec(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *MsgBase) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCodec
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgBase: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgBase: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GovDeposit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &MsgDeposit{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &MsgBase_GovDeposit{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GovVote", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &MsgVote{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &MsgBase_GovVote{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCodec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCodec
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCodec
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgSubmitProposalBase) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCodec
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSubmitProposalBase: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitProposalBase: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IntialDeposit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IntialDeposit = append(m.IntialDeposit, types.Coin{})
			if err := m.IntialDeposit[len(m.IntialDeposit)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proposer", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Proposer = append(m.Proposer[:0], dAtA[iNdEx:postIndex]...)
			if m.Proposer == nil {
				m.Proposer = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCodec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCodec
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCodec
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgDeposit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCodec
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgDeposit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeposit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposalID", wireType)
			}
			m.ProposalID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProposalID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Depositor", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Depositor = append(m.Depositor[:0], dAtA[iNdEx:postIndex]...)
			if m.Depositor == nil {
				m.Depositor = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = append(m.Amount, types.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCodec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCodec
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCodec
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgVote) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCodec
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgVote: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgVote: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposalID", wireType)
			}
			m.ProposalID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProposalID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Voter", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Voter = append(m.Voter[:0], dAtA[iNdEx:postIndex]...)
			if m.Voter == nil {
				m.Voter = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Option", wireType)
			}
			m.Option = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Option |= VoteOption(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCodec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCodec
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCodec
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TextProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCodec
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TextProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TextProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCodec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCodec
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCodec
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipCodec(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCodec
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthCodec
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCodec
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCodec
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCodec        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCodec          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCodec = fmt.Errorf("proto: unexpected end of group")
)
