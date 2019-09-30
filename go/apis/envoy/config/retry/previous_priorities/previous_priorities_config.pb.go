// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/retry/previous_priorities/previous_priorities_config.proto

package envoy_config_retry_previous_priorities

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// A retry host selector that attempts to spread retries between priorities, even if certain
// priorities would not normally be attempted due to higher priorities being available.
//
// As priorities get excluded, load will be distributed amongst the remaining healthy priorities
// based on the relative health of the priorities, matching how load is distributed during regular
// host selection. For example, given priority healths of {100, 50, 50}, the original load will be
// {100, 0, 0} (since P0 has capacity to handle 100% of the traffic). If P0 is excluded, the load
// changes to {0, 50, 50}, because P1 is only able to handle 50% of the traffic, causing the
// remaining to spill over to P2.
//
// Each priority attempted will be excluded until there are no healthy priorities left, at which
// point the list of attempted priorities will be reset, essentially starting from the beginning.
// For example, given three priorities P0, P1, P2 with healthy % of 100, 0 and 50 respectively, the
// following sequence of priorities would be selected (assuming update_frequency = 1):
// Attempt 1: P0 (P0 is 100% healthy)
// Attempt 2: P2 (P0 already attempted, P2 only healthy priority)
// Attempt 3: P0 (no healthy priorities, reset)
// Attempt 4: P2
//
// In the case of all upstream hosts being unhealthy, no adjustments will be made to the original
// priority load, so behavior should be identical to not using this plugin.
//
// Using this PriorityFilter requires rebuilding the priority load, which runs in O(# of
// priorities), which might incur significant overhead for clusters with many priorities.
type PreviousPrioritiesConfig struct {
	// How often the priority load should be updated based on previously attempted priorities. Useful
	// to allow each priorities to receive more than one request before being excluded or to reduce
	// the number of times that the priority load has to be recomputed.
	//
	// For example, by setting this to 2, then the first two attempts (initial attempt and first
	// retry) will use the unmodified priority load. The third and fourth attempt will use priority
	// load which excludes the priorities routed to with the first two attempts, and the fifth and
	// sixth attempt will use the priority load excluding the priorities used for the first four
	// attempts.
	UpdateFrequency      int32    `protobuf:"varint,1,opt,name=update_frequency,json=updateFrequency,proto3" json:"update_frequency,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PreviousPrioritiesConfig) Reset()         { *m = PreviousPrioritiesConfig{} }
func (m *PreviousPrioritiesConfig) String() string { return proto.CompactTextString(m) }
func (*PreviousPrioritiesConfig) ProtoMessage()    {}
func (*PreviousPrioritiesConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_763b66cb0fac7d2f, []int{0}
}
func (m *PreviousPrioritiesConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PreviousPrioritiesConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PreviousPrioritiesConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PreviousPrioritiesConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PreviousPrioritiesConfig.Merge(m, src)
}
func (m *PreviousPrioritiesConfig) XXX_Size() int {
	return m.Size()
}
func (m *PreviousPrioritiesConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_PreviousPrioritiesConfig.DiscardUnknown(m)
}

var xxx_messageInfo_PreviousPrioritiesConfig proto.InternalMessageInfo

func (m *PreviousPrioritiesConfig) GetUpdateFrequency() int32 {
	if m != nil {
		return m.UpdateFrequency
	}
	return 0
}

func init() {
	proto.RegisterType((*PreviousPrioritiesConfig)(nil), "envoy.config.retry.previous_priorities.PreviousPrioritiesConfig")
}

func init() {
	proto.RegisterFile("envoy/config/retry/previous_priorities/previous_priorities_config.proto", fileDescriptor_763b66cb0fac7d2f)
}

var fileDescriptor_763b66cb0fac7d2f = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x4f, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x2f, 0x4a, 0x2d, 0x29, 0xaa, 0xd4, 0x2f,
	0x28, 0x4a, 0x2d, 0xcb, 0xcc, 0x2f, 0x2d, 0x8e, 0x2f, 0x28, 0xca, 0xcc, 0x2f, 0xca, 0x2c, 0xc9,
	0x4c, 0x2d, 0xc6, 0x26, 0x16, 0x0f, 0xd1, 0xa4, 0x57, 0x50, 0x94, 0x5f, 0x92, 0x2f, 0xa4, 0x06,
	0x36, 0x48, 0x0f, 0x2a, 0x06, 0x36, 0x48, 0x0f, 0x8b, 0x26, 0x25, 0x57, 0x2e, 0x89, 0x00, 0xa8,
	0x70, 0x00, 0x5c, 0xd4, 0x19, 0xac, 0x4b, 0x48, 0x93, 0x4b, 0xa0, 0xb4, 0x20, 0x25, 0xb1, 0x24,
	0x35, 0x3e, 0xad, 0x28, 0xb5, 0xb0, 0x34, 0x35, 0x2f, 0xb9, 0x52, 0x82, 0x51, 0x81, 0x51, 0x83,
	0x35, 0x88, 0x1f, 0x22, 0xee, 0x06, 0x13, 0x76, 0x4a, 0x3e, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23,
	0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0xb9, 0x4c, 0x32, 0xf3, 0xf5, 0xc0, 0xf6, 0x17, 0x14, 0xe5,
	0x57, 0x54, 0xea, 0x11, 0xe7, 0x14, 0x27, 0x59, 0x5c, 0x0e, 0x09, 0x00, 0xf9, 0x28, 0x80, 0x31,
	0x89, 0x0d, 0xec, 0x35, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x42, 0x02, 0x93, 0x87, 0x25,
	0x01, 0x00, 0x00,
}

func (m *PreviousPrioritiesConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PreviousPrioritiesConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PreviousPrioritiesConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.UpdateFrequency != 0 {
		i = encodeVarintPreviousPrioritiesConfig(dAtA, i, uint64(m.UpdateFrequency))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPreviousPrioritiesConfig(dAtA []byte, offset int, v uint64) int {
	offset -= sovPreviousPrioritiesConfig(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PreviousPrioritiesConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.UpdateFrequency != 0 {
		n += 1 + sovPreviousPrioritiesConfig(uint64(m.UpdateFrequency))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovPreviousPrioritiesConfig(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPreviousPrioritiesConfig(x uint64) (n int) {
	return sovPreviousPrioritiesConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PreviousPrioritiesConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPreviousPrioritiesConfig
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
			return fmt.Errorf("proto: PreviousPrioritiesConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PreviousPrioritiesConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdateFrequency", wireType)
			}
			m.UpdateFrequency = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPreviousPrioritiesConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UpdateFrequency |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPreviousPrioritiesConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPreviousPrioritiesConfig
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPreviousPrioritiesConfig
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
func skipPreviousPrioritiesConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPreviousPrioritiesConfig
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
					return 0, ErrIntOverflowPreviousPrioritiesConfig
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
					return 0, ErrIntOverflowPreviousPrioritiesConfig
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
				return 0, ErrInvalidLengthPreviousPrioritiesConfig
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthPreviousPrioritiesConfig
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPreviousPrioritiesConfig
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
				next, err := skipPreviousPrioritiesConfig(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthPreviousPrioritiesConfig
				}
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
	ErrInvalidLengthPreviousPrioritiesConfig = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPreviousPrioritiesConfig   = fmt.Errorf("proto: integer overflow")
)
