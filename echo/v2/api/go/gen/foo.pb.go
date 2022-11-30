// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: foo.proto

// Revision 1

package foo_v2

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgDoSomething struct {
	Sender string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	// amount must be a non-zero integer
	Amount uint64 `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	// condition is an optional condition on doing the thing
	//
	// Since: Revision 1
	Condition string `protobuf:"bytes,3,opt,name=condition,proto3" json:"condition,omitempty"`
}

func (m *MsgDoSomething) Reset()         { *m = MsgDoSomething{} }
func (m *MsgDoSomething) String() string { return proto.CompactTextString(m) }
func (*MsgDoSomething) ProtoMessage()    {}
func (*MsgDoSomething) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ce1e2eec643ca48, []int{0}
}
func (m *MsgDoSomething) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDoSomething) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDoSomething.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDoSomething) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDoSomething.Merge(m, src)
}
func (m *MsgDoSomething) XXX_Size() int {
	return m.Size()
}
func (m *MsgDoSomething) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDoSomething.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDoSomething proto.InternalMessageInfo

func (m *MsgDoSomething) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *MsgDoSomething) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *MsgDoSomething) GetCondition() string {
	if m != nil {
		return m.Condition
	}
	return ""
}

type MsgDoSomethingResponse struct {
	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (m *MsgDoSomethingResponse) Reset()         { *m = MsgDoSomethingResponse{} }
func (m *MsgDoSomethingResponse) String() string { return proto.CompactTextString(m) }
func (*MsgDoSomethingResponse) ProtoMessage()    {}
func (*MsgDoSomethingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ce1e2eec643ca48, []int{1}
}
func (m *MsgDoSomethingResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDoSomethingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDoSomethingResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDoSomethingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDoSomethingResponse.Merge(m, src)
}
func (m *MsgDoSomethingResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgDoSomethingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDoSomethingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDoSomethingResponse proto.InternalMessageInfo

func (m *MsgDoSomethingResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*MsgDoSomething)(nil), "foo.v2.MsgDoSomething")
	proto.RegisterType((*MsgDoSomethingResponse)(nil), "foo.v2.MsgDoSomethingResponse")
}

func init() { proto.RegisterFile("foo.proto", fileDescriptor_7ce1e2eec643ca48) }

var fileDescriptor_7ce1e2eec643ca48 = []byte{
	// 203 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0xcb, 0xcf, 0xd7,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x03, 0x31, 0xcb, 0x8c, 0x94, 0xe2, 0xb8, 0xf8, 0x7c,
	0x8b, 0xd3, 0x5d, 0xf2, 0x83, 0xf3, 0x73, 0x53, 0x4b, 0x32, 0x32, 0xf3, 0xd2, 0x85, 0xc4, 0xb8,
	0xd8, 0x8a, 0x53, 0xf3, 0x52, 0x52, 0x8b, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xa0, 0x3c,
	0x90, 0x78, 0x62, 0x6e, 0x7e, 0x69, 0x5e, 0x89, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x4b, 0x10, 0x94,
	0x27, 0x24, 0xc3, 0xc5, 0x99, 0x9c, 0x9f, 0x97, 0x92, 0x59, 0x92, 0x99, 0x9f, 0x27, 0xc1, 0x0c,
	0xd6, 0x82, 0x10, 0x50, 0x32, 0xe2, 0x12, 0x43, 0x35, 0x3f, 0x28, 0xb5, 0xb8, 0x20, 0x3f, 0xaf,
	0x38, 0x55, 0x48, 0x82, 0x8b, 0xbd, 0xb8, 0x34, 0x39, 0x39, 0xb5, 0xb8, 0x18, 0x6c, 0x11, 0x47,
	0x10, 0x8c, 0x6b, 0xe4, 0xc3, 0xc5, 0xec, 0x5b, 0x9c, 0x2e, 0xe4, 0xca, 0xc5, 0x8d, 0xe2, 0x2e,
	0x3d, 0x88, 0x93, 0xf5, 0x50, 0xcd, 0x93, 0x92, 0xc3, 0x2e, 0x0e, 0xb3, 0xc7, 0x49, 0xe2, 0xc4,
	0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1,
	0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0x92, 0xd8, 0xc0, 0x41, 0x61, 0x0c, 0x08, 0x00,
	0x00, 0xff, 0xff, 0x9b, 0x7e, 0x02, 0xbe, 0x17, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	DoSomething(ctx context.Context, in *MsgDoSomething, opts ...grpc.CallOption) (*MsgDoSomethingResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) DoSomething(ctx context.Context, in *MsgDoSomething, opts ...grpc.CallOption) (*MsgDoSomethingResponse, error) {
	out := new(MsgDoSomethingResponse)
	err := c.cc.Invoke(ctx, "/foo.v2.Msg/DoSomething", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	DoSomething(context.Context, *MsgDoSomething) (*MsgDoSomethingResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) DoSomething(ctx context.Context, req *MsgDoSomething) (*MsgDoSomethingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoSomething not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_DoSomething_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDoSomething)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DoSomething(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/foo.v2.Msg/DoSomething",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DoSomething(ctx, req.(*MsgDoSomething))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "foo.v2.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DoSomething",
			Handler:    _Msg_DoSomething_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "foo.proto",
}

func (m *MsgDoSomething) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDoSomething) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDoSomething) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Condition) > 0 {
		i -= len(m.Condition)
		copy(dAtA[i:], m.Condition)
		i = encodeVarintFoo(dAtA, i, uint64(len(m.Condition)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Amount != 0 {
		i = encodeVarintFoo(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintFoo(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgDoSomethingResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDoSomethingResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDoSomethingResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Success {
		i--
		if m.Success {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintFoo(dAtA []byte, offset int, v uint64) int {
	offset -= sovFoo(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgDoSomething) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovFoo(uint64(l))
	}
	if m.Amount != 0 {
		n += 1 + sovFoo(uint64(m.Amount))
	}
	l = len(m.Condition)
	if l > 0 {
		n += 1 + l + sovFoo(uint64(l))
	}
	return n
}

func (m *MsgDoSomethingResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Success {
		n += 2
	}
	return n
}

func sovFoo(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFoo(x uint64) (n int) {
	return sovFoo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgDoSomething) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFoo
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
			return fmt.Errorf("proto: MsgDoSomething: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDoSomething: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFoo
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
				return ErrInvalidLengthFoo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFoo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFoo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Condition", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFoo
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
				return ErrInvalidLengthFoo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFoo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Condition = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFoo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFoo
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
func (m *MsgDoSomethingResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFoo
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
			return fmt.Errorf("proto: MsgDoSomethingResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDoSomethingResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Success", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFoo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Success = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipFoo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFoo
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
func skipFoo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFoo
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
					return 0, ErrIntOverflowFoo
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
					return 0, ErrIntOverflowFoo
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
				return 0, ErrInvalidLengthFoo
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFoo
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFoo
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFoo        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFoo          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFoo = fmt.Errorf("proto: unexpected end of group")
)
