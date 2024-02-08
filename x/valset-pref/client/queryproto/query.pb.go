// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/valsetpref/v1beta1/query.proto

package queryproto

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	types "github.com/osmosis-labs/osmosis/vv23/x/valset-pref/types"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Request type for UserValidatorPreferences.
type UserValidatorPreferencesRequest struct {
	// user account address
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *UserValidatorPreferencesRequest) Reset()         { *m = UserValidatorPreferencesRequest{} }
func (m *UserValidatorPreferencesRequest) String() string { return proto.CompactTextString(m) }
func (*UserValidatorPreferencesRequest) ProtoMessage()    {}
func (*UserValidatorPreferencesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e2d5b0777f607c6, []int{0}
}
func (m *UserValidatorPreferencesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserValidatorPreferencesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserValidatorPreferencesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UserValidatorPreferencesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserValidatorPreferencesRequest.Merge(m, src)
}
func (m *UserValidatorPreferencesRequest) XXX_Size() int {
	return m.Size()
}
func (m *UserValidatorPreferencesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserValidatorPreferencesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserValidatorPreferencesRequest proto.InternalMessageInfo

// Response type the QueryUserValidatorPreferences query request
type UserValidatorPreferencesResponse struct {
	Preferences []types.ValidatorPreference `protobuf:"bytes,1,rep,name=preferences,proto3" json:"preferences"`
}

func (m *UserValidatorPreferencesResponse) Reset()         { *m = UserValidatorPreferencesResponse{} }
func (m *UserValidatorPreferencesResponse) String() string { return proto.CompactTextString(m) }
func (*UserValidatorPreferencesResponse) ProtoMessage()    {}
func (*UserValidatorPreferencesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e2d5b0777f607c6, []int{1}
}
func (m *UserValidatorPreferencesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserValidatorPreferencesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserValidatorPreferencesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UserValidatorPreferencesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserValidatorPreferencesResponse.Merge(m, src)
}
func (m *UserValidatorPreferencesResponse) XXX_Size() int {
	return m.Size()
}
func (m *UserValidatorPreferencesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserValidatorPreferencesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserValidatorPreferencesResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*UserValidatorPreferencesRequest)(nil), "osmosis.valsetpref.v1beta1.UserValidatorPreferencesRequest")
	proto.RegisterType((*UserValidatorPreferencesResponse)(nil), "osmosis.valsetpref.v1beta1.UserValidatorPreferencesResponse")
}

func init() {
	proto.RegisterFile("osmosis/valsetpref/v1beta1/query.proto", fileDescriptor_6e2d5b0777f607c6)
}

var fileDescriptor_6e2d5b0777f607c6 = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0x4d, 0x4b, 0xc3, 0x30,
	0x18, 0x6e, 0xfc, 0xc4, 0xec, 0x56, 0x3c, 0x94, 0x22, 0xd9, 0xd8, 0x61, 0xec, 0xb2, 0x84, 0xd5,
	0xe3, 0x3c, 0xcd, 0x3f, 0xa0, 0x03, 0x15, 0xbc, 0xa5, 0xdb, 0xbb, 0x5a, 0xe8, 0x9a, 0x2e, 0x6f,
	0x36, 0x94, 0xe1, 0xc5, 0x5f, 0x20, 0xf8, 0x9b, 0x84, 0xdd, 0x1c, 0x78, 0xf1, 0x24, 0xba, 0xf9,
	0x43, 0x64, 0x6b, 0xc7, 0x1c, 0xd8, 0x09, 0x9e, 0xda, 0x24, 0xcf, 0x47, 0x9e, 0xf7, 0x09, 0xad,
	0x28, 0xec, 0x29, 0x0c, 0x51, 0x0c, 0x65, 0x84, 0x60, 0x12, 0x0d, 0x5d, 0x31, 0xac, 0xfb, 0x60,
	0x64, 0x5d, 0xf4, 0x07, 0xa0, 0xef, 0x78, 0xa2, 0x95, 0x51, 0xb6, 0x9b, 0xe1, 0xf8, 0x0a, 0xc7,
	0x33, 0x9c, 0x7b, 0x18, 0xa8, 0x40, 0x2d, 0x60, 0x62, 0xfe, 0x97, 0x32, 0xdc, 0xa3, 0x40, 0xa9,
	0x20, 0x02, 0x21, 0x93, 0x50, 0xc8, 0x38, 0x56, 0x46, 0x9a, 0x50, 0xc5, 0x98, 0x9d, 0x6e, 0xf2,
	0x45, 0x23, 0x0d, 0xa4, 0xb8, 0x72, 0x83, 0x16, 0x2f, 0x10, 0xf4, 0xa5, 0x8c, 0xc2, 0x8e, 0x34,
	0x4a, 0x9f, 0x69, 0xe8, 0x82, 0x86, 0xb8, 0x0d, 0xd8, 0x82, 0xfe, 0x00, 0xd0, 0xd8, 0x0e, 0xdd,
	0x97, 0x9d, 0x8e, 0x06, 0x44, 0x87, 0x94, 0x48, 0xf5, 0xa0, 0xb5, 0x5c, 0x96, 0x47, 0xb4, 0x94,
	0x4f, 0xc6, 0x44, 0xc5, 0x08, 0xf6, 0x15, 0x2d, 0x24, 0xab, 0x6d, 0x87, 0x94, 0xb6, 0xab, 0x05,
	0x4f, 0xf0, 0xfc, 0xb8, 0xfc, 0x17, 0xb9, 0xe6, 0xce, 0xf8, 0xbd, 0x68, 0xb5, 0x7e, 0x2a, 0x79,
	0x2f, 0x84, 0xee, 0x9e, 0xcf, 0x27, 0x68, 0x3f, 0x13, 0xea, 0xe4, 0xdd, 0xc3, 0x6e, 0x6c, 0xb2,
	0xfa, 0x23, 0xba, 0x7b, 0xf2, 0x3f, 0x72, 0x1a, 0xbd, 0xcc, 0x1f, 0x5e, 0xbf, 0x9e, 0xb6, 0xaa,
	0x76, 0x45, 0xac, 0x97, 0x51, 0x5b, 0x6b, 0x63, 0x94, 0x4d, 0xf3, 0xbe, 0x29, 0xc7, 0x9f, 0xcc,
	0x1a, 0x4f, 0x19, 0x99, 0x4c, 0x19, 0xf9, 0x98, 0x32, 0xf2, 0x38, 0x63, 0xd6, 0x64, 0xc6, 0xac,
	0xb7, 0x19, 0xb3, 0xae, 0x4f, 0x83, 0xd0, 0xdc, 0x0c, 0x7c, 0xde, 0x56, 0xbd, 0xa5, 0x5e, 0x2d,
	0x92, 0x3e, 0xae, 0xc4, 0x3d, 0x4f, 0xdc, 0xae, 0x59, 0xb4, 0xa3, 0x10, 0x62, 0x93, 0xbe, 0xb3,
	0x45, 0xdd, 0xfe, 0xde, 0xe2, 0x73, 0xfc, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x53, 0x3b, 0x84, 0x33,
	0x97, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Returns the list of ValidatorPreferences for the user.
	UserValidatorPreferences(ctx context.Context, in *UserValidatorPreferencesRequest, opts ...grpc.CallOption) (*UserValidatorPreferencesResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) UserValidatorPreferences(ctx context.Context, in *UserValidatorPreferencesRequest, opts ...grpc.CallOption) (*UserValidatorPreferencesResponse, error) {
	out := new(UserValidatorPreferencesResponse)
	err := c.cc.Invoke(ctx, "/osmosis.valsetpref.v1beta1.Query/UserValidatorPreferences", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Returns the list of ValidatorPreferences for the user.
	UserValidatorPreferences(context.Context, *UserValidatorPreferencesRequest) (*UserValidatorPreferencesResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) UserValidatorPreferences(ctx context.Context, req *UserValidatorPreferencesRequest) (*UserValidatorPreferencesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserValidatorPreferences not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_UserValidatorPreferences_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserValidatorPreferencesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).UserValidatorPreferences(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/osmosis.valsetpref.v1beta1.Query/UserValidatorPreferences",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).UserValidatorPreferences(ctx, req.(*UserValidatorPreferencesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "osmosis.valsetpref.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserValidatorPreferences",
			Handler:    _Query_UserValidatorPreferences_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "osmosis/valsetpref/v1beta1/query.proto",
}

func (m *UserValidatorPreferencesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserValidatorPreferencesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UserValidatorPreferencesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *UserValidatorPreferencesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserValidatorPreferencesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UserValidatorPreferencesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Preferences) > 0 {
		for iNdEx := len(m.Preferences) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Preferences[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *UserValidatorPreferencesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *UserValidatorPreferencesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Preferences) > 0 {
		for _, e := range m.Preferences {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *UserValidatorPreferencesRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: UserValidatorPreferencesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserValidatorPreferencesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *UserValidatorPreferencesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: UserValidatorPreferencesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserValidatorPreferencesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Preferences", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Preferences = append(m.Preferences, types.ValidatorPreference{})
			if err := m.Preferences[len(m.Preferences)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
