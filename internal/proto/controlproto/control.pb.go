// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: control.proto

/*
	Package controlproto is a generated protocol buffer package.

	It is generated from these files:
		control.proto

	It has these top-level messages:
		Command
		Node
		Unsubscribe
		Disconnect
*/
package controlproto

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import github_com_centrifugal_centrifuge_internal_proto "github.com/centrifugal/centrifuge/internal/proto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type MethodType int32

const (
	MethodTypeNode        MethodType = 0
	MethodTypeUnsubscribe MethodType = 1
	MethodTypeDisconnect  MethodType = 2
)

var MethodType_name = map[int32]string{
	0: "NODE",
	1: "UNSUBSCRIBE",
	2: "DISCONNECT",
}
var MethodType_value = map[string]int32{
	"NODE":        0,
	"UNSUBSCRIBE": 1,
	"DISCONNECT":  2,
}

func (x MethodType) String() string {
	return proto.EnumName(MethodType_name, int32(x))
}
func (MethodType) EnumDescriptor() ([]byte, []int) { return fileDescriptorControl, []int{0} }

type Command struct {
	UID    string                                               `protobuf:"bytes,1,opt,name=UID,proto3" json:"UID,omitempty"`
	Method MethodType                                           `protobuf:"varint,2,opt,name=Method,proto3,enum=controlproto.MethodType" json:"Method,omitempty"`
	Params github_com_centrifugal_centrifuge_internal_proto.Raw `protobuf:"bytes,3,opt,name=Params,proto3,customtype=github.com/centrifugal/centrifuge/internal/proto.Raw" json:"Params"`
}

func (m *Command) Reset()                    { *m = Command{} }
func (m *Command) String() string            { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()               {}
func (*Command) Descriptor() ([]byte, []int) { return fileDescriptorControl, []int{0} }

func (m *Command) GetUID() string {
	if m != nil {
		return m.UID
	}
	return ""
}

func (m *Command) GetMethod() MethodType {
	if m != nil {
		return m.Method
	}
	return MethodTypeNode
}

type Node struct {
	UID         string `protobuf:"bytes,1,opt,name=UID,proto3" json:"UID,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Version     string `protobuf:"bytes,3,opt,name=Version,proto3" json:"Version,omitempty"`
	NumClients  uint32 `protobuf:"varint,4,opt,name=NumClients,proto3" json:"NumClients,omitempty"`
	NumUsers    uint32 `protobuf:"varint,5,opt,name=NumUsers,proto3" json:"NumUsers,omitempty"`
	NumChannels uint32 `protobuf:"varint,6,opt,name=NumChannels,proto3" json:"NumChannels,omitempty"`
	Uptime      uint32 `protobuf:"varint,7,opt,name=Uptime,proto3" json:"Uptime,omitempty"`
}

func (m *Node) Reset()                    { *m = Node{} }
func (m *Node) String() string            { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()               {}
func (*Node) Descriptor() ([]byte, []int) { return fileDescriptorControl, []int{1} }

func (m *Node) GetUID() string {
	if m != nil {
		return m.UID
	}
	return ""
}

func (m *Node) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Node) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Node) GetNumClients() uint32 {
	if m != nil {
		return m.NumClients
	}
	return 0
}

func (m *Node) GetNumUsers() uint32 {
	if m != nil {
		return m.NumUsers
	}
	return 0
}

func (m *Node) GetNumChannels() uint32 {
	if m != nil {
		return m.NumChannels
	}
	return 0
}

func (m *Node) GetUptime() uint32 {
	if m != nil {
		return m.Uptime
	}
	return 0
}

type Unsubscribe struct {
	Channel string `protobuf:"bytes,1,opt,name=Channel,proto3" json:"Channel,omitempty"`
	User    string `protobuf:"bytes,2,opt,name=User,proto3" json:"User,omitempty"`
}

func (m *Unsubscribe) Reset()                    { *m = Unsubscribe{} }
func (m *Unsubscribe) String() string            { return proto.CompactTextString(m) }
func (*Unsubscribe) ProtoMessage()               {}
func (*Unsubscribe) Descriptor() ([]byte, []int) { return fileDescriptorControl, []int{2} }

func (m *Unsubscribe) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *Unsubscribe) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

type Disconnect struct {
	User string `protobuf:"bytes,1,opt,name=User,proto3" json:"User,omitempty"`
}

func (m *Disconnect) Reset()                    { *m = Disconnect{} }
func (m *Disconnect) String() string            { return proto.CompactTextString(m) }
func (*Disconnect) ProtoMessage()               {}
func (*Disconnect) Descriptor() ([]byte, []int) { return fileDescriptorControl, []int{3} }

func (m *Disconnect) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func init() {
	proto.RegisterType((*Command)(nil), "controlproto.Command")
	proto.RegisterType((*Node)(nil), "controlproto.Node")
	proto.RegisterType((*Unsubscribe)(nil), "controlproto.Unsubscribe")
	proto.RegisterType((*Disconnect)(nil), "controlproto.Disconnect")
	proto.RegisterEnum("controlproto.MethodType", MethodType_name, MethodType_value)
}
func (this *Command) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Command)
	if !ok {
		that2, ok := that.(Command)
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
	if this.UID != that1.UID {
		return false
	}
	if this.Method != that1.Method {
		return false
	}
	if !this.Params.Equal(that1.Params) {
		return false
	}
	return true
}
func (this *Node) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Node)
	if !ok {
		that2, ok := that.(Node)
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
	if this.UID != that1.UID {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Version != that1.Version {
		return false
	}
	if this.NumClients != that1.NumClients {
		return false
	}
	if this.NumUsers != that1.NumUsers {
		return false
	}
	if this.NumChannels != that1.NumChannels {
		return false
	}
	if this.Uptime != that1.Uptime {
		return false
	}
	return true
}
func (this *Unsubscribe) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Unsubscribe)
	if !ok {
		that2, ok := that.(Unsubscribe)
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
	if this.Channel != that1.Channel {
		return false
	}
	if this.User != that1.User {
		return false
	}
	return true
}
func (this *Disconnect) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Disconnect)
	if !ok {
		that2, ok := that.(Disconnect)
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
	if this.User != that1.User {
		return false
	}
	return true
}
func (m *Command) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Command) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.UID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintControl(dAtA, i, uint64(len(m.UID)))
		i += copy(dAtA[i:], m.UID)
	}
	if m.Method != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintControl(dAtA, i, uint64(m.Method))
	}
	dAtA[i] = 0x1a
	i++
	i = encodeVarintControl(dAtA, i, uint64(m.Params.Size()))
	n1, err := m.Params.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	return i, nil
}

func (m *Node) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Node) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.UID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintControl(dAtA, i, uint64(len(m.UID)))
		i += copy(dAtA[i:], m.UID)
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintControl(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Version) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintControl(dAtA, i, uint64(len(m.Version)))
		i += copy(dAtA[i:], m.Version)
	}
	if m.NumClients != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintControl(dAtA, i, uint64(m.NumClients))
	}
	if m.NumUsers != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintControl(dAtA, i, uint64(m.NumUsers))
	}
	if m.NumChannels != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintControl(dAtA, i, uint64(m.NumChannels))
	}
	if m.Uptime != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintControl(dAtA, i, uint64(m.Uptime))
	}
	return i, nil
}

func (m *Unsubscribe) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Unsubscribe) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Channel) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintControl(dAtA, i, uint64(len(m.Channel)))
		i += copy(dAtA[i:], m.Channel)
	}
	if len(m.User) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintControl(dAtA, i, uint64(len(m.User)))
		i += copy(dAtA[i:], m.User)
	}
	return i, nil
}

func (m *Disconnect) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Disconnect) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.User) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintControl(dAtA, i, uint64(len(m.User)))
		i += copy(dAtA[i:], m.User)
	}
	return i, nil
}

func encodeVarintControl(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedCommand(r randyControl, easy bool) *Command {
	this := &Command{}
	this.UID = string(randStringControl(r))
	this.Method = MethodType([]int32{0, 1, 2}[r.Intn(3)])
	v1 := github_com_centrifugal_centrifuge_internal_proto.NewPopulatedRaw(r)
	this.Params = *v1
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedNode(r randyControl, easy bool) *Node {
	this := &Node{}
	this.UID = string(randStringControl(r))
	this.Name = string(randStringControl(r))
	this.Version = string(randStringControl(r))
	this.NumClients = uint32(r.Uint32())
	this.NumUsers = uint32(r.Uint32())
	this.NumChannels = uint32(r.Uint32())
	this.Uptime = uint32(r.Uint32())
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedUnsubscribe(r randyControl, easy bool) *Unsubscribe {
	this := &Unsubscribe{}
	this.Channel = string(randStringControl(r))
	this.User = string(randStringControl(r))
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedDisconnect(r randyControl, easy bool) *Disconnect {
	this := &Disconnect{}
	this.User = string(randStringControl(r))
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyControl interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneControl(r randyControl) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringControl(r randyControl) string {
	v2 := r.Intn(100)
	tmps := make([]rune, v2)
	for i := 0; i < v2; i++ {
		tmps[i] = randUTF8RuneControl(r)
	}
	return string(tmps)
}
func randUnrecognizedControl(r randyControl, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldControl(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldControl(dAtA []byte, r randyControl, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateControl(dAtA, uint64(key))
		v3 := r.Int63()
		if r.Intn(2) == 0 {
			v3 *= -1
		}
		dAtA = encodeVarintPopulateControl(dAtA, uint64(v3))
	case 1:
		dAtA = encodeVarintPopulateControl(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateControl(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateControl(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateControl(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateControl(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *Command) Size() (n int) {
	var l int
	_ = l
	l = len(m.UID)
	if l > 0 {
		n += 1 + l + sovControl(uint64(l))
	}
	if m.Method != 0 {
		n += 1 + sovControl(uint64(m.Method))
	}
	l = m.Params.Size()
	n += 1 + l + sovControl(uint64(l))
	return n
}

func (m *Node) Size() (n int) {
	var l int
	_ = l
	l = len(m.UID)
	if l > 0 {
		n += 1 + l + sovControl(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovControl(uint64(l))
	}
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovControl(uint64(l))
	}
	if m.NumClients != 0 {
		n += 1 + sovControl(uint64(m.NumClients))
	}
	if m.NumUsers != 0 {
		n += 1 + sovControl(uint64(m.NumUsers))
	}
	if m.NumChannels != 0 {
		n += 1 + sovControl(uint64(m.NumChannels))
	}
	if m.Uptime != 0 {
		n += 1 + sovControl(uint64(m.Uptime))
	}
	return n
}

func (m *Unsubscribe) Size() (n int) {
	var l int
	_ = l
	l = len(m.Channel)
	if l > 0 {
		n += 1 + l + sovControl(uint64(l))
	}
	l = len(m.User)
	if l > 0 {
		n += 1 + l + sovControl(uint64(l))
	}
	return n
}

func (m *Disconnect) Size() (n int) {
	var l int
	_ = l
	l = len(m.User)
	if l > 0 {
		n += 1 + l + sovControl(uint64(l))
	}
	return n
}

func sovControl(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozControl(x uint64) (n int) {
	return sovControl(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Command) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowControl
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Command: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Command: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthControl
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Method", wireType)
			}
			m.Method = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Method |= (MethodType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthControl
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipControl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthControl
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
func (m *Node) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowControl
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Node: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Node: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthControl
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthControl
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthControl
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumClients", wireType)
			}
			m.NumClients = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumClients |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumUsers", wireType)
			}
			m.NumUsers = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumUsers |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumChannels", wireType)
			}
			m.NumChannels = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumChannels |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uptime", wireType)
			}
			m.Uptime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Uptime |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipControl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthControl
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
func (m *Unsubscribe) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowControl
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Unsubscribe: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Unsubscribe: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Channel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthControl
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Channel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthControl
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.User = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipControl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthControl
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
func (m *Disconnect) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowControl
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Disconnect: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Disconnect: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthControl
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.User = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipControl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthControl
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
func skipControl(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowControl
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
					return 0, ErrIntOverflowControl
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowControl
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthControl
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowControl
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipControl(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthControl = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowControl   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("control.proto", fileDescriptorControl) }

var fileDescriptorControl = []byte{
	// 462 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xb1, 0x6e, 0xd3, 0x40,
	0x1c, 0xc6, 0x73, 0xad, 0x71, 0xc8, 0x3f, 0x6d, 0x15, 0x9d, 0x00, 0x1d, 0x56, 0xe5, 0x5a, 0x99,
	0xac, 0x4a, 0x38, 0x08, 0xd8, 0x60, 0x8a, 0x93, 0x21, 0x03, 0x2e, 0x72, 0x62, 0x76, 0xdb, 0xb9,
	0x26, 0x96, 0xec, 0xbb, 0xc8, 0x77, 0x16, 0xe2, 0x05, 0x10, 0xca, 0xc4, 0x0b, 0x64, 0xea, 0xc2,
	0x23, 0x30, 0xf0, 0x00, 0x1d, 0x99, 0x19, 0x2a, 0x30, 0x2f, 0xc1, 0x88, 0x7c, 0x71, 0x6a, 0x0f,
	0x6c, 0xdf, 0xf7, 0xff, 0xbe, 0xf3, 0xfd, 0xfe, 0x3e, 0x38, 0x8d, 0x39, 0x93, 0x39, 0x4f, 0x9d,
	0x4d, 0xce, 0x25, 0xc7, 0x27, 0xb5, 0x55, 0xce, 0x78, 0xb6, 0x4a, 0xe4, 0xba, 0x88, 0x9c, 0x98,
	0x67, 0xa3, 0x15, 0x5f, 0xf1, 0x91, 0x1a, 0x47, 0xc5, 0xb5, 0x72, 0xca, 0x28, 0xb5, 0x3f, 0x3c,
	0xbc, 0x41, 0xd0, 0x75, 0x79, 0x96, 0x85, 0x6c, 0x89, 0x07, 0x70, 0x1c, 0xcc, 0x26, 0x04, 0x59,
	0xc8, 0xee, 0xf9, 0x95, 0xc4, 0xcf, 0x41, 0x7f, 0x4b, 0xe5, 0x9a, 0x2f, 0xc9, 0x91, 0x85, 0xec,
	0xb3, 0x17, 0xc4, 0x69, 0xdf, 0xe5, 0xec, 0xb3, 0xc5, 0xc7, 0x0d, 0xf5, 0xeb, 0x1e, 0x5e, 0x80,
	0xfe, 0x2e, 0xcc, 0xc3, 0x4c, 0x90, 0x63, 0x0b, 0xd9, 0x27, 0xe3, 0x37, 0xb7, 0x77, 0x17, 0x9d,
	0x9f, 0x77, 0x17, 0xaf, 0x5a, 0x58, 0x31, 0x65, 0x32, 0x4f, 0xae, 0x8b, 0x55, 0x98, 0x36, 0x9a,
	0x8e, 0x12, 0x26, 0x69, 0xce, 0xc2, 0x74, 0x4f, 0xec, 0xf8, 0xe1, 0x07, 0xbf, 0xfe, 0xd6, 0xf0,
	0x3b, 0x02, 0xcd, 0xe3, 0x4b, 0xfa, 0x1f, 0x44, 0x0c, 0x9a, 0x17, 0x66, 0x54, 0x01, 0xf6, 0x7c,
	0xa5, 0x31, 0x81, 0xee, 0x7b, 0x9a, 0x8b, 0x84, 0x33, 0x45, 0xd1, 0xf3, 0x0f, 0x16, 0x9b, 0x00,
	0x5e, 0x91, 0xb9, 0x69, 0x42, 0x99, 0x14, 0x44, 0xb3, 0x90, 0x7d, 0xea, 0xb7, 0x26, 0xd8, 0x80,
	0x87, 0x5e, 0x91, 0x05, 0x82, 0xe6, 0x82, 0x3c, 0x50, 0xe9, 0xbd, 0xc7, 0x16, 0xf4, 0xab, 0xe6,
	0x3a, 0x64, 0x8c, 0xa6, 0x82, 0xe8, 0x2a, 0x6e, 0x8f, 0xf0, 0x13, 0xd0, 0x83, 0x8d, 0x4c, 0x32,
	0x4a, 0xba, 0x2a, 0xac, 0xdd, 0xf0, 0x35, 0xf4, 0x03, 0x26, 0x8a, 0x48, 0xc4, 0x79, 0x12, 0x29,
	0xbc, 0xfa, 0x48, 0xbd, 0xc8, 0xc1, 0x56, 0xcb, 0x54, 0x77, 0x1d, 0x96, 0xa9, 0xf4, 0xd0, 0x02,
	0x98, 0x24, 0x22, 0xe6, 0x8c, 0xd1, 0x58, 0xde, 0x37, 0x50, 0xd3, 0xb8, 0xfc, 0x84, 0x00, 0x9a,
	0xa7, 0xc0, 0xe7, 0xa0, 0x79, 0x57, 0x93, 0xe9, 0xa0, 0x63, 0xe0, 0xed, 0xce, 0x3a, 0x6b, 0x12,
	0xf5, 0x07, 0x2f, 0xa1, 0x1f, 0x78, 0xf3, 0x60, 0x3c, 0x77, 0xfd, 0xd9, 0x78, 0x3a, 0x40, 0xc6,
	0xd3, 0xed, 0xce, 0x7a, 0xdc, 0x94, 0xda, 0xa0, 0x36, 0xc0, 0x64, 0x36, 0x77, 0xaf, 0x3c, 0x6f,
	0xea, 0x2e, 0x06, 0x47, 0x06, 0xd9, 0xee, 0xac, 0x47, 0x4d, 0xb5, 0xc1, 0x32, 0xb4, 0xcf, 0x37,
	0x66, 0x67, 0x7c, 0xfe, 0xf7, 0xb7, 0x89, 0xbe, 0x96, 0x26, 0xfa, 0x56, 0x9a, 0xe8, 0xb6, 0x34,
	0xd1, 0x8f, 0xd2, 0x44, 0xbf, 0x4a, 0x13, 0x7d, 0xf9, 0x63, 0x76, 0x22, 0x5d, 0xbd, 0xeb, 0xcb,
	0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x8d, 0xe9, 0x7e, 0xfd, 0xbf, 0x02, 0x00, 0x00,
}
