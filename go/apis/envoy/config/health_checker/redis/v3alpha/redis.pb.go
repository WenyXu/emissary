// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/health_checker/redis/v3alpha/redis.proto

package envoy_config_health_checker_redis_v3alpha

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

type Redis struct {
	// If set, optionally perform ``EXISTS <key>`` instead of ``PING``. A return value
	// from Redis of 0 (does not exist) is considered a passing healthcheck. A return value other
	// than 0 is considered a failure. This allows the user to mark a Redis instance for maintenance
	// by setting the specified key to any value and waiting for traffic to drain.
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Redis) Reset()         { *m = Redis{} }
func (m *Redis) String() string { return proto.CompactTextString(m) }
func (*Redis) ProtoMessage()    {}
func (*Redis) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd8be85b91892539, []int{0}
}
func (m *Redis) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Redis) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Redis.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Redis) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Redis.Merge(m, src)
}
func (m *Redis) XXX_Size() int {
	return m.Size()
}
func (m *Redis) XXX_DiscardUnknown() {
	xxx_messageInfo_Redis.DiscardUnknown(m)
}

var xxx_messageInfo_Redis proto.InternalMessageInfo

func (m *Redis) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func init() {
	proto.RegisterType((*Redis)(nil), "envoy.config.health_checker.redis.v3alpha.Redis")
}

func init() {
	proto.RegisterFile("envoy/config/health_checker/redis/v3alpha/redis.proto", fileDescriptor_cd8be85b91892539)
}

var fileDescriptor_cd8be85b91892539 = []byte{
	// 158 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4d, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0xcf, 0x48, 0x4d, 0xcc, 0x29, 0xc9, 0x88,
	0x4f, 0xce, 0x48, 0x4d, 0xce, 0x4e, 0x2d, 0xd2, 0x2f, 0x4a, 0x4d, 0xc9, 0x2c, 0xd6, 0x2f, 0x33,
	0x4e, 0xcc, 0x29, 0xc8, 0x48, 0x84, 0xf0, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x34, 0xc1,
	0xda, 0xf4, 0x20, 0xda, 0xf4, 0x50, 0xb5, 0xe9, 0x41, 0x14, 0x42, 0xb5, 0x29, 0x49, 0x72, 0xb1,
	0x06, 0x81, 0x04, 0x84, 0x04, 0xb8, 0x98, 0xb3, 0x53, 0x2b, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38,
	0x83, 0x40, 0x4c, 0xa7, 0xe0, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48,
	0x8e, 0x91, 0xcb, 0x3c, 0x33, 0x5f, 0x0f, 0x6c, 0x6c, 0x41, 0x51, 0x7e, 0x45, 0xa5, 0x1e, 0xd1,
	0x36, 0x38, 0x71, 0x81, 0xcd, 0x0f, 0x00, 0x39, 0x2c, 0x80, 0x31, 0x89, 0x0d, 0xec, 0x42, 0x63,
	0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3d, 0x56, 0x36, 0xe3, 0xda, 0x00, 0x00, 0x00,
}

func (m *Redis) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Redis) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Redis) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintRedis(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRedis(dAtA []byte, offset int, v uint64) int {
	offset -= sovRedis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Redis) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovRedis(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovRedis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRedis(x uint64) (n int) {
	return sovRedis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Redis) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRedis
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
			return fmt.Errorf("proto: Redis: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Redis: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRedis
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
				return ErrInvalidLengthRedis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRedis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRedis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRedis
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRedis
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
func skipRedis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRedis
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
					return 0, ErrIntOverflowRedis
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
					return 0, ErrIntOverflowRedis
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
				return 0, ErrInvalidLengthRedis
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthRedis
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRedis
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
				next, err := skipRedis(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthRedis
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
	ErrInvalidLengthRedis = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRedis   = fmt.Errorf("proto: integer overflow")
)
