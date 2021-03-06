// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: queryfrontend/response.proto

package queryfrontend

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	labelpb "github.com/thanos-io/thanos/pkg/store/labelpb"
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

type ThanosLabelsResponse struct {
	Status    string   `protobuf:"bytes,1,opt,name=Status,proto3" json:"status"`
	Data      []string `protobuf:"bytes,2,rep,name=Data,proto3" json:"data"`
	ErrorType string   `protobuf:"bytes,3,opt,name=ErrorType,proto3" json:"errorType,omitempty"`
	Error     string   `protobuf:"bytes,4,opt,name=Error,proto3" json:"error,omitempty"`
}

func (m *ThanosLabelsResponse) Reset()         { *m = ThanosLabelsResponse{} }
func (m *ThanosLabelsResponse) String() string { return proto.CompactTextString(m) }
func (*ThanosLabelsResponse) ProtoMessage()    {}
func (*ThanosLabelsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b882fa7024d92f38, []int{0}
}
func (m *ThanosLabelsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ThanosLabelsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ThanosLabelsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ThanosLabelsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThanosLabelsResponse.Merge(m, src)
}
func (m *ThanosLabelsResponse) XXX_Size() int {
	return m.Size()
}
func (m *ThanosLabelsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ThanosLabelsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ThanosLabelsResponse proto.InternalMessageInfo

type ThanosSeriesResponse struct {
	Status string `protobuf:"bytes,1,opt,name=Status,proto3" json:"status"`
	// TODO(bwplotka): Experiment with ZLabelSet here.
	Data      []labelpb.LabelSet `protobuf:"bytes,2,rep,name=Data,proto3" json:"data"`
	ErrorType string             `protobuf:"bytes,3,opt,name=ErrorType,proto3" json:"errorType,omitempty"`
	Error     string             `protobuf:"bytes,4,opt,name=Error,proto3" json:"error,omitempty"`
}

func (m *ThanosSeriesResponse) Reset()         { *m = ThanosSeriesResponse{} }
func (m *ThanosSeriesResponse) String() string { return proto.CompactTextString(m) }
func (*ThanosSeriesResponse) ProtoMessage()    {}
func (*ThanosSeriesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b882fa7024d92f38, []int{1}
}
func (m *ThanosSeriesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ThanosSeriesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ThanosSeriesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ThanosSeriesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThanosSeriesResponse.Merge(m, src)
}
func (m *ThanosSeriesResponse) XXX_Size() int {
	return m.Size()
}
func (m *ThanosSeriesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ThanosSeriesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ThanosSeriesResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ThanosLabelsResponse)(nil), "queryfrontend.ThanosLabelsResponse")
	proto.RegisterType((*ThanosSeriesResponse)(nil), "queryfrontend.ThanosSeriesResponse")
}

func init() { proto.RegisterFile("queryfrontend/response.proto", fileDescriptor_b882fa7024d92f38) }

var fileDescriptor_b882fa7024d92f38 = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x91, 0x31, 0x4e, 0xc3, 0x30,
	0x14, 0x86, 0x63, 0x5a, 0x2a, 0x6a, 0x40, 0xa0, 0xb4, 0x12, 0xa1, 0xaa, 0x9c, 0xaa, 0x53, 0x91,
	0x50, 0x22, 0x15, 0x71, 0x81, 0x08, 0x36, 0xa6, 0xa4, 0x17, 0x70, 0xd4, 0x47, 0xa9, 0xd4, 0xc6,
	0xc6, 0x7e, 0x1d, 0x72, 0x0b, 0xce, 0xc3, 0x09, 0x32, 0x66, 0x64, 0x8a, 0x20, 0xd9, 0x72, 0x0a,
	0x54, 0xa7, 0x40, 0x19, 0x59, 0xd8, 0xec, 0xff, 0xfb, 0xdf, 0xaf, 0xdf, 0x7e, 0x74, 0xf8, 0xbc,
	0x01, 0x95, 0x3e, 0x2a, 0x91, 0x20, 0x24, 0x73, 0x5f, 0x81, 0x96, 0x22, 0xd1, 0xe0, 0x49, 0x25,
	0x50, 0xd8, 0xa7, 0xbf, 0xe8, 0xa0, 0xbf, 0x10, 0x0b, 0x61, 0x88, 0xbf, 0x3d, 0x35, 0xa6, 0xc1,
	0xa5, 0x46, 0xa1, 0xc0, 0x5f, 0xf1, 0x18, 0x56, 0x32, 0xf6, 0x31, 0x95, 0xa0, 0x1b, 0x34, 0x7e,
	0x25, 0xb4, 0x3f, 0x7b, 0xe2, 0x89, 0xd0, 0x0f, 0x5b, 0xaa, 0xc3, 0x5d, 0xbc, 0x3d, 0xa6, 0x9d,
	0x08, 0x39, 0x6e, 0xb4, 0x43, 0x46, 0x64, 0xd2, 0x0d, 0x68, 0x5d, 0xb8, 0x1d, 0x6d, 0x94, 0x70,
	0x47, 0xec, 0x21, 0x6d, 0xdf, 0x71, 0xe4, 0xce, 0xc1, 0xa8, 0x35, 0xe9, 0x06, 0x47, 0x75, 0xe1,
	0xb6, 0xe7, 0x1c, 0x79, 0x68, 0x54, 0xfb, 0x96, 0x76, 0xef, 0x95, 0x12, 0x6a, 0x96, 0x4a, 0x70,
	0x5a, 0x26, 0xe4, 0xa2, 0x2e, 0xdc, 0x1e, 0x7c, 0x89, 0xd7, 0x62, 0xbd, 0x44, 0x58, 0x4b, 0x4c,
	0xc3, 0x1f, 0xa7, 0x7d, 0x45, 0x0f, 0xcd, 0xc5, 0x69, 0x9b, 0x91, 0x5e, 0x5d, 0xb8, 0x67, 0x66,
	0x64, 0xcf, 0xde, 0x38, 0xc6, 0xf9, 0x77, 0xf9, 0x08, 0xd4, 0x12, 0xfe, 0x56, 0x7e, 0xba, 0x57,
	0xfe, 0x78, 0x7a, 0xee, 0xa1, 0x09, 0xf2, 0xcc, 0x37, 0x44, 0x80, 0xc1, 0x49, 0x56, 0xb8, 0xd6,
	0x7f, 0x3f, 0x29, 0x18, 0x66, 0x1f, 0xcc, 0xca, 0x4a, 0x46, 0xf2, 0x92, 0x91, 0xf7, 0x92, 0x91,
	0x97, 0x8a, 0x59, 0x79, 0xc5, 0xac, 0xb7, 0x8a, 0x59, 0x71, 0xc7, 0x2c, 0xed, 0xe6, 0x33, 0x00,
	0x00, 0xff, 0xff, 0xd2, 0x7d, 0x34, 0xcb, 0x14, 0x02, 0x00, 0x00,
}

func (m *ThanosLabelsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ThanosLabelsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ThanosLabelsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Error) > 0 {
		i -= len(m.Error)
		copy(dAtA[i:], m.Error)
		i = encodeVarintResponse(dAtA, i, uint64(len(m.Error)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ErrorType) > 0 {
		i -= len(m.ErrorType)
		copy(dAtA[i:], m.ErrorType)
		i = encodeVarintResponse(dAtA, i, uint64(len(m.ErrorType)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Data) > 0 {
		for iNdEx := len(m.Data) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Data[iNdEx])
			copy(dAtA[i:], m.Data[iNdEx])
			i = encodeVarintResponse(dAtA, i, uint64(len(m.Data[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintResponse(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ThanosSeriesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ThanosSeriesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ThanosSeriesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Error) > 0 {
		i -= len(m.Error)
		copy(dAtA[i:], m.Error)
		i = encodeVarintResponse(dAtA, i, uint64(len(m.Error)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ErrorType) > 0 {
		i -= len(m.ErrorType)
		copy(dAtA[i:], m.ErrorType)
		i = encodeVarintResponse(dAtA, i, uint64(len(m.ErrorType)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Data) > 0 {
		for iNdEx := len(m.Data) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Data[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintResponse(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintResponse(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintResponse(dAtA []byte, offset int, v uint64) int {
	offset -= sovResponse(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ThanosLabelsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovResponse(uint64(l))
	}
	if len(m.Data) > 0 {
		for _, s := range m.Data {
			l = len(s)
			n += 1 + l + sovResponse(uint64(l))
		}
	}
	l = len(m.ErrorType)
	if l > 0 {
		n += 1 + l + sovResponse(uint64(l))
	}
	l = len(m.Error)
	if l > 0 {
		n += 1 + l + sovResponse(uint64(l))
	}
	return n
}

func (m *ThanosSeriesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovResponse(uint64(l))
	}
	if len(m.Data) > 0 {
		for _, e := range m.Data {
			l = e.Size()
			n += 1 + l + sovResponse(uint64(l))
		}
	}
	l = len(m.ErrorType)
	if l > 0 {
		n += 1 + l + sovResponse(uint64(l))
	}
	l = len(m.Error)
	if l > 0 {
		n += 1 + l + sovResponse(uint64(l))
	}
	return n
}

func sovResponse(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozResponse(x uint64) (n int) {
	return sovResponse(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ThanosLabelsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowResponse
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
			return fmt.Errorf("proto: ThanosLabelsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ThanosLabelsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResponse
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
				return ErrInvalidLengthResponse
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthResponse
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResponse
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
				return ErrInvalidLengthResponse
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthResponse
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ErrorType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResponse
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
				return ErrInvalidLengthResponse
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthResponse
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ErrorType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResponse
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
				return ErrInvalidLengthResponse
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthResponse
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Error = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipResponse(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthResponse
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthResponse
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
func (m *ThanosSeriesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowResponse
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
			return fmt.Errorf("proto: ThanosSeriesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ThanosSeriesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResponse
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
				return ErrInvalidLengthResponse
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthResponse
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResponse
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
				return ErrInvalidLengthResponse
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthResponse
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data, labelpb.LabelSet{})
			if err := m.Data[len(m.Data)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ErrorType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResponse
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
				return ErrInvalidLengthResponse
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthResponse
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ErrorType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResponse
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
				return ErrInvalidLengthResponse
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthResponse
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Error = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipResponse(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthResponse
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthResponse
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
func skipResponse(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowResponse
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
					return 0, ErrIntOverflowResponse
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
					return 0, ErrIntOverflowResponse
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
				return 0, ErrInvalidLengthResponse
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupResponse
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthResponse
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthResponse        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowResponse          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupResponse = fmt.Errorf("proto: unexpected end of group")
)
