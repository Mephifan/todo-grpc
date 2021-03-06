// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: protos/rest.proto

package todo

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

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

// https://tools.ietf.org/html/rfc7231#section-4
type Link_Method int32

const (
	Link_PUT     Link_Method = 0
	Link_GET     Link_Method = 1
	Link_HEAD    Link_Method = 2
	Link_POST    Link_Method = 3
	Link_PATCH   Link_Method = 4
	Link_DELETE  Link_Method = 5
	Link_TRACE   Link_Method = 6
	Link_OPTIONS Link_Method = 7
)

var Link_Method_name = map[int32]string{
	0: "PUT",
	1: "GET",
	2: "HEAD",
	3: "POST",
	4: "PATCH",
	5: "DELETE",
	6: "TRACE",
	7: "OPTIONS",
}
var Link_Method_value = map[string]int32{
	"PUT":     0,
	"GET":     1,
	"HEAD":    2,
	"POST":    3,
	"PATCH":   4,
	"DELETE":  5,
	"TRACE":   6,
	"OPTIONS": 7,
}

func (x Link_Method) String() string {
	return proto.EnumName(Link_Method_name, int32(x))
}
func (Link_Method) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_rest_8747127e184342fa, []int{0, 0}
}

// /-------------------------------------------------------------------
// / List of official link rels:
// / http://www.iana.org/assignments/link-relations/link-relations.xhtml
// /-------------------------------------------------------------------
type Link struct {
	Rel                  string      `protobuf:"bytes,1,opt,name=rel,proto3" json:"rel,omitempty"`
	Method               Link_Method `protobuf:"varint,2,opt,name=method,proto3,enum=todo.v1.Link_Method" json:"method,omitempty"`
	Href                 string      `protobuf:"bytes,3,opt,name=href,proto3" json:"href,omitempty"`
	Type                 string      `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Link) Reset()         { *m = Link{} }
func (m *Link) String() string { return proto.CompactTextString(m) }
func (*Link) ProtoMessage()    {}
func (*Link) Descriptor() ([]byte, []int) {
	return fileDescriptor_rest_8747127e184342fa, []int{0}
}
func (m *Link) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Link) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Link.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Link) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Link.Merge(dst, src)
}
func (m *Link) XXX_Size() int {
	return m.Size()
}
func (m *Link) XXX_DiscardUnknown() {
	xxx_messageInfo_Link.DiscardUnknown(m)
}

var xxx_messageInfo_Link proto.InternalMessageInfo

func (m *Link) GetRel() string {
	if m != nil {
		return m.Rel
	}
	return ""
}

func (m *Link) GetMethod() Link_Method {
	if m != nil {
		return m.Method
	}
	return Link_PUT
}

func (m *Link) GetHref() string {
	if m != nil {
		return m.Href
	}
	return ""
}

func (m *Link) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type Hateoas struct {
	Links                []*Link  `protobuf:"bytes,1,rep,name=Links" json:"Links,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Hateoas) Reset()         { *m = Hateoas{} }
func (m *Hateoas) String() string { return proto.CompactTextString(m) }
func (*Hateoas) ProtoMessage()    {}
func (*Hateoas) Descriptor() ([]byte, []int) {
	return fileDescriptor_rest_8747127e184342fa, []int{1}
}
func (m *Hateoas) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Hateoas) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Hateoas.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Hateoas) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Hateoas.Merge(dst, src)
}
func (m *Hateoas) XXX_Size() int {
	return m.Size()
}
func (m *Hateoas) XXX_DiscardUnknown() {
	xxx_messageInfo_Hateoas.DiscardUnknown(m)
}

var xxx_messageInfo_Hateoas proto.InternalMessageInfo

func (m *Hateoas) GetLinks() []*Link {
	if m != nil {
		return m.Links
	}
	return nil
}

func init() {
	proto.RegisterType((*Link)(nil), "todo.v1.Link")
	proto.RegisterType((*Hateoas)(nil), "todo.v1.Hateoas")
	proto.RegisterEnum("todo.v1.Link_Method", Link_Method_name, Link_Method_value)
}
func (m *Link) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Link) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Rel) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRest(dAtA, i, uint64(len(m.Rel)))
		i += copy(dAtA[i:], m.Rel)
	}
	if m.Method != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintRest(dAtA, i, uint64(m.Method))
	}
	if len(m.Href) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintRest(dAtA, i, uint64(len(m.Href)))
		i += copy(dAtA[i:], m.Href)
	}
	if len(m.Type) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintRest(dAtA, i, uint64(len(m.Type)))
		i += copy(dAtA[i:], m.Type)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Hateoas) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Hateoas) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Links) > 0 {
		for _, msg := range m.Links {
			dAtA[i] = 0xa
			i++
			i = encodeVarintRest(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintRest(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Link) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Rel)
	if l > 0 {
		n += 1 + l + sovRest(uint64(l))
	}
	if m.Method != 0 {
		n += 1 + sovRest(uint64(m.Method))
	}
	l = len(m.Href)
	if l > 0 {
		n += 1 + l + sovRest(uint64(l))
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovRest(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Hateoas) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Links) > 0 {
		for _, e := range m.Links {
			l = e.Size()
			n += 1 + l + sovRest(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovRest(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRest(x uint64) (n int) {
	return sovRest(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Link) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRest
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
			return fmt.Errorf("proto: Link: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Link: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRest
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
				return ErrInvalidLengthRest
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Rel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Method", wireType)
			}
			m.Method = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Method |= (Link_Method(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Href", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRest
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
				return ErrInvalidLengthRest
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Href = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRest
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
				return ErrInvalidLengthRest
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRest
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Hateoas) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRest
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
			return fmt.Errorf("proto: Hateoas: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Hateoas: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Links", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRest
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Links = append(m.Links, &Link{})
			if err := m.Links[len(m.Links)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRest
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipRest(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRest
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
					return 0, ErrIntOverflowRest
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
					return 0, ErrIntOverflowRest
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
				return 0, ErrInvalidLengthRest
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRest
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
				next, err := skipRest(dAtA[start:])
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
	ErrInvalidLengthRest = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRest   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("protos/rest.proto", fileDescriptor_rest_8747127e184342fa) }

var fileDescriptor_rest_8747127e184342fa = []byte{
	// 265 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xcb, 0x4a, 0xc3, 0x40,
	0x14, 0x86, 0x9d, 0x66, 0x92, 0xd8, 0x53, 0x94, 0xf1, 0xe0, 0x62, 0x56, 0x21, 0xc4, 0x4d, 0x16,
	0x32, 0x62, 0x7d, 0x82, 0xd8, 0x0e, 0x46, 0xa8, 0x26, 0xa4, 0xe3, 0x46, 0x70, 0x51, 0xe9, 0x48,
	0xc5, 0xcb, 0x94, 0x64, 0x10, 0x7c, 0x43, 0x37, 0x82, 0x8f, 0x20, 0x79, 0x12, 0x99, 0x49, 0x37,
	0xee, 0x3e, 0xfe, 0xcb, 0xe1, 0xe7, 0xc0, 0xd1, 0xb6, 0x35, 0xd6, 0x74, 0x67, 0xad, 0xee, 0xac,
	0xf0, 0x8c, 0xb1, 0x35, 0x6b, 0x23, 0x3e, 0xce, 0xb3, 0x6f, 0x02, 0x74, 0xf1, 0xfc, 0xfe, 0x82,
	0x0c, 0x82, 0x56, 0xbf, 0x72, 0x92, 0x92, 0x7c, 0xdc, 0x38, 0xc4, 0x53, 0x88, 0xde, 0xb4, 0xdd,
	0x98, 0x35, 0x1f, 0xa5, 0x24, 0x3f, 0x9c, 0x1e, 0x8b, 0x5d, 0x49, 0xb8, 0x82, 0xb8, 0xf1, 0x5e,
	0xb3, 0xcb, 0x20, 0x02, 0xdd, 0xb4, 0xfa, 0x89, 0x07, 0xfe, 0x80, 0x67, 0xa7, 0xd9, 0xcf, 0xad,
	0xe6, 0x74, 0xd0, 0x1c, 0x67, 0x0f, 0x10, 0x0d, 0x4d, 0x8c, 0x21, 0xa8, 0xef, 0x14, 0xdb, 0x73,
	0x70, 0x25, 0x15, 0x23, 0xb8, 0x0f, 0xb4, 0x94, 0xc5, 0x9c, 0x8d, 0x1c, 0xd5, 0xd5, 0x52, 0xb1,
	0x00, 0xc7, 0x10, 0xd6, 0x85, 0x9a, 0x95, 0x8c, 0x22, 0x40, 0x34, 0x97, 0x0b, 0xa9, 0x24, 0x0b,
	0x9d, 0xac, 0x9a, 0x62, 0x26, 0x59, 0x84, 0x13, 0x88, 0xab, 0x5a, 0x5d, 0x57, 0xb7, 0x4b, 0x16,
	0x67, 0x02, 0xe2, 0x72, 0x65, 0xb5, 0x59, 0x75, 0x78, 0x02, 0xa1, 0x1b, 0xda, 0x71, 0x92, 0x06,
	0xf9, 0x64, 0x7a, 0xf0, 0x6f, 0x7e, 0x33, 0x78, 0x97, 0xf8, 0xd5, 0x27, 0xe4, 0xa7, 0x4f, 0xc8,
	0x6f, 0x9f, 0x90, 0x7b, 0xea, 0x22, 0x8f, 0x91, 0xff, 0xd1, 0xc5, 0x5f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x8a, 0xd3, 0x01, 0x8e, 0x38, 0x01, 0x00, 0x00,
}
