// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/filter/network/zookeeper_proxy/v1alpha1/zookeeper_proxy.proto

package envoy_config_filter_network_zookeeper_proxy_v1alpha1

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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

// [#protodoc-title: ZooKeeper proxy]
// ZooKeeper Proxy :ref:`configuration overview <config_network_filters_zookeeper_proxy>`.
type ZooKeeperProxy struct {
	// The human readable prefix to use when emitting :ref:`statistics
	// <config_network_filters_zookeeper_proxy_stats>`.
	StatPrefix string `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	// [#not-implemented-hide:] The optional path to use for writing ZooKeeper access logs.
	// If the access log field is empty, access logs will not be written.
	AccessLog string `protobuf:"bytes,2,opt,name=access_log,json=accessLog,proto3" json:"access_log,omitempty"`
	// Messages — requests, responses and events — that are bigger than this value will
	// be ignored. If it is not set, the default value is 1Mb.
	//
	// The value here should match the jute.maxbuffer property in your cluster configuration:
	//
	// https://zookeeper.apache.org/doc/r3.4.10/zookeeperAdmin.html#Unsafe+Options
	//
	// if that is set. If it isn't, ZooKeeper's default is also 1Mb.
	MaxPacketBytes       *types.UInt32Value `protobuf:"bytes,3,opt,name=max_packet_bytes,json=maxPacketBytes,proto3" json:"max_packet_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ZooKeeperProxy) Reset()         { *m = ZooKeeperProxy{} }
func (m *ZooKeeperProxy) String() string { return proto.CompactTextString(m) }
func (*ZooKeeperProxy) ProtoMessage()    {}
func (*ZooKeeperProxy) Descriptor() ([]byte, []int) {
	return fileDescriptor_05247350458709ad, []int{0}
}
func (m *ZooKeeperProxy) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ZooKeeperProxy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ZooKeeperProxy.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ZooKeeperProxy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ZooKeeperProxy.Merge(m, src)
}
func (m *ZooKeeperProxy) XXX_Size() int {
	return m.Size()
}
func (m *ZooKeeperProxy) XXX_DiscardUnknown() {
	xxx_messageInfo_ZooKeeperProxy.DiscardUnknown(m)
}

var xxx_messageInfo_ZooKeeperProxy proto.InternalMessageInfo

func (m *ZooKeeperProxy) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func (m *ZooKeeperProxy) GetAccessLog() string {
	if m != nil {
		return m.AccessLog
	}
	return ""
}

func (m *ZooKeeperProxy) GetMaxPacketBytes() *types.UInt32Value {
	if m != nil {
		return m.MaxPacketBytes
	}
	return nil
}

func init() {
	proto.RegisterType((*ZooKeeperProxy)(nil), "envoy.config.filter.network.zookeeper_proxy.v1alpha1.ZooKeeperProxy")
}

func init() {
	proto.RegisterFile("envoy/config/filter/network/zookeeper_proxy/v1alpha1/zookeeper_proxy.proto", fileDescriptor_05247350458709ad)
}

var fileDescriptor_05247350458709ad = []byte{
	// 315 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x90, 0x3f, 0x4e, 0xc3, 0x30,
	0x18, 0xc5, 0xe5, 0x16, 0x21, 0xd5, 0x95, 0x2a, 0x14, 0x06, 0xaa, 0x0a, 0xa2, 0x8a, 0xa9, 0x62,
	0xb0, 0xd5, 0x96, 0x13, 0x64, 0x40, 0xe2, 0xcf, 0x10, 0x55, 0x82, 0xa1, 0x4b, 0xe4, 0x86, 0x2f,
	0x26, 0xaa, 0x9b, 0xcf, 0x72, 0xdc, 0x36, 0xe5, 0x3a, 0xdc, 0x82, 0x89, 0x91, 0x91, 0x23, 0xa0,
	0x6e, 0xdc, 0x02, 0xc5, 0x26, 0x4b, 0x47, 0xb6, 0xc4, 0x3f, 0xbf, 0x9f, 0x9e, 0x1f, 0xbd, 0x83,
	0x62, 0x83, 0x3b, 0x9e, 0x62, 0x91, 0xe5, 0x92, 0x67, 0xb9, 0xb2, 0x60, 0x78, 0x01, 0x76, 0x8b,
	0x66, 0xc9, 0x5f, 0x11, 0x97, 0x00, 0x1a, 0x4c, 0xa2, 0x0d, 0x56, 0x3b, 0xbe, 0x19, 0x0b, 0xa5,
	0x5f, 0xc4, 0xf8, 0x10, 0x30, 0x6d, 0xd0, 0x62, 0x70, 0xed, 0x5c, 0xcc, 0xbb, 0x98, 0x77, 0xb1,
	0x3f, 0x17, 0x3b, 0x8c, 0x34, 0xae, 0xc1, 0xd9, 0x46, 0xa8, 0xfc, 0x59, 0x58, 0xe0, 0xcd, 0x87,
	0xd7, 0x0d, 0x42, 0x89, 0x28, 0x15, 0x70, 0xf7, 0xb7, 0x58, 0x67, 0x7c, 0x6b, 0x84, 0xd6, 0x60,
	0x4a, 0xcf, 0x2f, 0xdf, 0x08, 0xed, 0xcd, 0x11, 0xef, 0x9d, 0x35, 0xae, 0xa5, 0xc1, 0x15, 0xed,
	0x96, 0x56, 0xd8, 0x44, 0x1b, 0xc8, 0xf2, 0xaa, 0x4f, 0x86, 0x64, 0xd4, 0x89, 0x3a, 0xef, 0x3f,
	0x1f, 0xed, 0x23, 0xd3, 0x1a, 0x92, 0x19, 0xad, 0x69, 0xec, 0x60, 0x70, 0x41, 0xa9, 0x48, 0x53,
	0x28, 0xcb, 0x44, 0xa1, 0xec, 0xb7, 0xea, 0xab, 0xb3, 0x8e, 0x3f, 0x79, 0x40, 0x19, 0xdc, 0xd0,
	0x93, 0x95, 0xa8, 0x12, 0x2d, 0xd2, 0x25, 0xd8, 0x64, 0xb1, 0xb3, 0x50, 0xf6, 0xdb, 0x43, 0x32,
	0xea, 0x4e, 0xce, 0x99, 0x2f, 0xc6, 0x9a, 0x62, 0xec, 0xf1, 0xb6, 0xb0, 0xd3, 0xc9, 0x93, 0x50,
	0x6b, 0x98, 0xf5, 0x56, 0xa2, 0x8a, 0x5d, 0x28, 0xaa, 0x33, 0x91, 0xfc, 0xdc, 0x87, 0xe4, 0x6b,
	0x1f, 0x92, 0xef, 0x7d, 0x48, 0x68, 0x94, 0x23, 0x73, 0x2b, 0xf9, 0x11, 0xfe, 0x33, 0x58, 0x74,
	0x3a, 0x6f, 0x88, 0x7b, 0x74, 0x5c, 0xb7, 0x88, 0xc9, 0xe2, 0xd8, 0xd5, 0x99, 0xfe, 0x06, 0x00,
	0x00, 0xff, 0xff, 0xf9, 0xfb, 0xec, 0xa1, 0xd2, 0x01, 0x00, 0x00,
}

func (m *ZooKeeperProxy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ZooKeeperProxy) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ZooKeeperProxy) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.MaxPacketBytes != nil {
		{
			size, err := m.MaxPacketBytes.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintZookeeperProxy(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.AccessLog) > 0 {
		i -= len(m.AccessLog)
		copy(dAtA[i:], m.AccessLog)
		i = encodeVarintZookeeperProxy(dAtA, i, uint64(len(m.AccessLog)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.StatPrefix) > 0 {
		i -= len(m.StatPrefix)
		copy(dAtA[i:], m.StatPrefix)
		i = encodeVarintZookeeperProxy(dAtA, i, uint64(len(m.StatPrefix)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintZookeeperProxy(dAtA []byte, offset int, v uint64) int {
	offset -= sovZookeeperProxy(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ZooKeeperProxy) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.StatPrefix)
	if l > 0 {
		n += 1 + l + sovZookeeperProxy(uint64(l))
	}
	l = len(m.AccessLog)
	if l > 0 {
		n += 1 + l + sovZookeeperProxy(uint64(l))
	}
	if m.MaxPacketBytes != nil {
		l = m.MaxPacketBytes.Size()
		n += 1 + l + sovZookeeperProxy(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovZookeeperProxy(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozZookeeperProxy(x uint64) (n int) {
	return sovZookeeperProxy(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ZooKeeperProxy) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowZookeeperProxy
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
			return fmt.Errorf("proto: ZooKeeperProxy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ZooKeeperProxy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StatPrefix", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowZookeeperProxy
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
				return ErrInvalidLengthZookeeperProxy
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthZookeeperProxy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StatPrefix = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccessLog", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowZookeeperProxy
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
				return ErrInvalidLengthZookeeperProxy
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthZookeeperProxy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccessLog = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxPacketBytes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowZookeeperProxy
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
				return ErrInvalidLengthZookeeperProxy
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthZookeeperProxy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MaxPacketBytes == nil {
				m.MaxPacketBytes = &types.UInt32Value{}
			}
			if err := m.MaxPacketBytes.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipZookeeperProxy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthZookeeperProxy
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthZookeeperProxy
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
func skipZookeeperProxy(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowZookeeperProxy
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
					return 0, ErrIntOverflowZookeeperProxy
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
					return 0, ErrIntOverflowZookeeperProxy
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
				return 0, ErrInvalidLengthZookeeperProxy
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthZookeeperProxy
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowZookeeperProxy
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
				next, err := skipZookeeperProxy(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthZookeeperProxy
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
	ErrInvalidLengthZookeeperProxy = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowZookeeperProxy   = fmt.Errorf("proto: integer overflow")
)
