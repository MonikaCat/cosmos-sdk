// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/auth/types/codec.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type BaseAccount struct {
	Address       github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=address,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"address,omitempty"`
	Coins         []types.Coin                                  `protobuf:"bytes,2,rep,name=coins,proto3" json:"coins"`
	Pubkey        []byte                                        `protobuf:"bytes,3,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	AccountNumber uint64                                        `protobuf:"varint,4,opt,name=account_number,json=accountNumber,proto3" json:"account_number,omitempty"`
	Sequence      uint64                                        `protobuf:"varint,5,opt,name=sequence,proto3" json:"sequence,omitempty"`
}

func (m *BaseAccount) Reset()      { *m = BaseAccount{} }
func (*BaseAccount) ProtoMessage() {}
func (*BaseAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_e44d2a11716720ad, []int{0}
}
func (m *BaseAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BaseAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BaseAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BaseAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseAccount.Merge(m, src)
}
func (m *BaseAccount) XXX_Size() int {
	return m.Size()
}
func (m *BaseAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseAccount.DiscardUnknown(m)
}

var xxx_messageInfo_BaseAccount proto.InternalMessageInfo

type ContinuousVestingAccount struct {
}

func (m *ContinuousVestingAccount) Reset()      { *m = ContinuousVestingAccount{} }
func (*ContinuousVestingAccount) ProtoMessage() {}
func (*ContinuousVestingAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_e44d2a11716720ad, []int{1}
}
func (m *ContinuousVestingAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ContinuousVestingAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ContinuousVestingAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ContinuousVestingAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContinuousVestingAccount.Merge(m, src)
}
func (m *ContinuousVestingAccount) XXX_Size() int {
	return m.Size()
}
func (m *ContinuousVestingAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_ContinuousVestingAccount.DiscardUnknown(m)
}

var xxx_messageInfo_ContinuousVestingAccount proto.InternalMessageInfo

func init() {
	proto.RegisterType((*BaseAccount)(nil), "cosmos_sdk.x.auth.BaseAccount")
	proto.RegisterType((*ContinuousVestingAccount)(nil), "cosmos_sdk.x.auth.ContinuousVestingAccount")
}

func init() { proto.RegisterFile("x/auth/types/codec.proto", fileDescriptor_e44d2a11716720ad) }

var fileDescriptor_e44d2a11716720ad = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x3f, 0x4f, 0xeb, 0x30,
	0x14, 0xc5, 0xed, 0xd7, 0x3f, 0xaf, 0x72, 0xdf, 0x43, 0x34, 0x42, 0xc8, 0xca, 0xe0, 0x46, 0x95,
	0x90, 0x3a, 0xd0, 0x44, 0xc0, 0x06, 0x53, 0xd3, 0x11, 0x89, 0x21, 0x03, 0x03, 0x4b, 0x95, 0x38,
	0x56, 0x1a, 0x55, 0xb5, 0x4b, 0x6d, 0x4b, 0xed, 0xd6, 0x91, 0x91, 0x8f, 0xc0, 0xc8, 0x47, 0xe9,
	0xd8, 0xb1, 0x03, 0xaa, 0xa8, 0xbb, 0xb0, 0x20, 0x75, 0x66, 0x42, 0x4d, 0x02, 0x42, 0x62, 0xb2,
	0xef, 0xf5, 0xef, 0xfa, 0x9c, 0x73, 0x11, 0x9e, 0x7a, 0xa1, 0x56, 0x03, 0x4f, 0xcd, 0xc6, 0x4c,
	0x7a, 0x54, 0xc4, 0x8c, 0xba, 0xe3, 0x89, 0x50, 0xc2, 0x6a, 0x50, 0x21, 0x47, 0x42, 0xf6, 0x65,
	0x3c, 0x74, 0xa7, 0xee, 0x1e, 0xb2, 0x8f, 0x12, 0x91, 0x88, 0xec, 0xd5, 0xdb, 0xdf, 0x72, 0xd0,
	0x6e, 0xfc, 0x9a, 0x6d, 0xbd, 0x43, 0x54, 0xf7, 0x43, 0xc9, 0xba, 0x94, 0x0a, 0xcd, 0x95, 0x75,
	0x8d, 0xfe, 0x86, 0x71, 0x3c, 0x61, 0x52, 0x62, 0xe8, 0xc0, 0xf6, 0x3f, 0xff, 0xec, 0x63, 0xdd,
	0xec, 0x24, 0xa9, 0x1a, 0xe8, 0xc8, 0xa5, 0x62, 0xe4, 0xe5, 0x5a, 0xc5, 0xd1, 0x91, 0xf1, 0x30,
	0x37, 0xe4, 0x76, 0x29, 0xed, 0xe6, 0x83, 0xc1, 0xd7, 0x0f, 0xd6, 0x29, 0xaa, 0x50, 0x91, 0x72,
	0x89, 0xff, 0x38, 0xa5, 0x76, 0xfd, 0xfc, 0xd0, 0xfd, 0x61, 0xb4, 0x27, 0x52, 0xee, 0x97, 0x17,
	0xeb, 0x26, 0x08, 0x72, 0xc8, 0x3a, 0x46, 0xd5, 0xb1, 0x8e, 0x86, 0x6c, 0x86, 0x4b, 0x7b, 0xe5,
	0xa0, 0xa8, 0xac, 0x13, 0x74, 0x10, 0xe6, 0xee, 0xfa, 0x5c, 0x8f, 0x22, 0x36, 0xc1, 0x65, 0x07,
	0xb6, 0xcb, 0xc1, 0xff, 0xa2, 0x7b, 0x93, 0x35, 0x2d, 0x1b, 0xd5, 0x24, 0xbb, 0xd7, 0x8c, 0x53,
	0x86, 0x2b, 0x19, 0xf0, 0x5d, 0x5f, 0xd6, 0x1e, 0x9e, 0x9a, 0x60, 0xfe, 0xe2, 0x80, 0x96, 0x8d,
	0x70, 0x4f, 0x70, 0x95, 0x72, 0x2d, 0xb4, 0xbc, 0x65, 0x52, 0xa5, 0x3c, 0x29, 0xb2, 0xfb, 0x57,
	0xcb, 0x0d, 0x01, 0xab, 0x0d, 0x01, 0xbb, 0x0d, 0x81, 0x73, 0x43, 0xe0, 0xb3, 0x21, 0x70, 0x61,
	0x08, 0x5c, 0x1a, 0x02, 0x5f, 0x0d, 0x81, 0x6f, 0x86, 0x80, 0x9d, 0x21, 0xf0, 0x71, 0x4b, 0xc0,
	0x72, 0x4b, 0xc0, 0x6a, 0x4b, 0xc0, 0x5d, 0x25, 0xcb, 0x1f, 0x55, 0xb3, 0x7d, 0x5e, 0x7c, 0x06,
	0x00, 0x00, 0xff, 0xff, 0xac, 0x2d, 0xe0, 0x5c, 0xa7, 0x01, 0x00, 0x00,
}

func (this *BaseAccount) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*BaseAccount)
	if !ok {
		that2, ok := that.(BaseAccount)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.Address, that1.Address) {
		return false
	}
	if len(this.Coins) != len(that1.Coins) {
		return false
	}
	for i := range this.Coins {
		if !this.Coins[i].Equal(&that1.Coins[i]) {
			return false
		}
	}
	if !bytes.Equal(this.Pubkey, that1.Pubkey) {
		return false
	}
	if this.AccountNumber != that1.AccountNumber {
		return false
	}
	if this.Sequence != that1.Sequence {
		return false
	}
	return true
}
func (this *ContinuousVestingAccount) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ContinuousVestingAccount)
	if !ok {
		that2, ok := that.(ContinuousVestingAccount)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	return true
}
func (this *BaseAccount) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 9)
	s = append(s, "&types.BaseAccount{")
	s = append(s, "Address: "+fmt.Sprintf("%#v", this.Address)+",\n")
	if this.Coins != nil {
		vs := make([]types.Coin, len(this.Coins))
		for i := range vs {
			vs[i] = this.Coins[i]
		}
		s = append(s, "Coins: "+fmt.Sprintf("%#v", vs)+",\n")
	}
	s = append(s, "Pubkey: "+fmt.Sprintf("%#v", this.Pubkey)+",\n")
	s = append(s, "AccountNumber: "+fmt.Sprintf("%#v", this.AccountNumber)+",\n")
	s = append(s, "Sequence: "+fmt.Sprintf("%#v", this.Sequence)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ContinuousVestingAccount) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&types.ContinuousVestingAccount{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringCodec(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *BaseAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BaseAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BaseAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Sequence != 0 {
		i = encodeVarintCodec(dAtA, i, uint64(m.Sequence))
		i--
		dAtA[i] = 0x28
	}
	if m.AccountNumber != 0 {
		i = encodeVarintCodec(dAtA, i, uint64(m.AccountNumber))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Pubkey) > 0 {
		i -= len(m.Pubkey)
		copy(dAtA[i:], m.Pubkey)
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Pubkey)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Coins) > 0 {
		for iNdEx := len(m.Coins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Coins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCodec(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ContinuousVestingAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContinuousVestingAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ContinuousVestingAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
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
func (m *BaseAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	if len(m.Coins) > 0 {
		for _, e := range m.Coins {
			l = e.Size()
			n += 1 + l + sovCodec(uint64(l))
		}
	}
	l = len(m.Pubkey)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	if m.AccountNumber != 0 {
		n += 1 + sovCodec(uint64(m.AccountNumber))
	}
	if m.Sequence != 0 {
		n += 1 + sovCodec(uint64(m.Sequence))
	}
	return n
}

func (m *ContinuousVestingAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovCodec(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCodec(x uint64) (n int) {
	return sovCodec(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ContinuousVestingAccount) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ContinuousVestingAccount{`,
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
func (m *BaseAccount) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: BaseAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BaseAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coins", wireType)
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
			m.Coins = append(m.Coins, types.Coin{})
			if err := m.Coins[len(m.Coins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pubkey", wireType)
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
			m.Pubkey = append(m.Pubkey[:0], dAtA[iNdEx:postIndex]...)
			if m.Pubkey == nil {
				m.Pubkey = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccountNumber", wireType)
			}
			m.AccountNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AccountNumber |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
			}
			m.Sequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sequence |= uint64(b&0x7F) << shift
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
func (m *ContinuousVestingAccount) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ContinuousVestingAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContinuousVestingAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
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
