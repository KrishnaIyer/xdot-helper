// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/command.proto

package pbapi // import "github.com/KrishnaIyer/xdot-helper/pkg/pbapi"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Result_ResCode int32

const (
	Result_NONE  Result_ResCode = 0
	Result_OK    Result_ResCode = 1
	Result_ERROR Result_ResCode = 2
)

var Result_ResCode_name = map[int32]string{
	0: "NONE",
	1: "OK",
	2: "ERROR",
}
var Result_ResCode_value = map[string]int32{
	"NONE":  0,
	"OK":    1,
	"ERROR": 2,
}

func (x Result_ResCode) String() string {
	return proto.EnumName(Result_ResCode_name, int32(x))
}
func (Result_ResCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_command_9ee6ba3dcbd6590a, []int{1, 0}
}

// Command is equivalent to a single AT command as defined in the ATcommand spec.
type Command struct {
	// name is the user defined name of the AT command.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// description is a simple function description.
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	// request is the actual request sent to the device.
	Request string `protobuf:"bytes,3,opt,name=request" json:"request,omitempty"`
	// arguments are optional arguments used for some commands.
	Arguments string `protobuf:"bytes,4,opt,name=arguments" json:"arguments,omitempty"`
	// wait_period is the amount of time to wait for a response.
	WaitPeriod int32 `protobuf:"varint,5,opt,name=wait_period,json=waitPeriod" json:"wait_period,omitempty"`
	// lines_in_response is the expected lines in the response for each  command.
	LinesInResponse      int32    `protobuf:"varint,6,opt,name=lines_in_response,json=linesInResponse" json:"lines_in_response,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Command) Reset()         { *m = Command{} }
func (m *Command) String() string { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()    {}
func (*Command) Descriptor() ([]byte, []int) {
	return fileDescriptor_command_9ee6ba3dcbd6590a, []int{0}
}
func (m *Command) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Command.Unmarshal(m, b)
}
func (m *Command) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Command.Marshal(b, m, deterministic)
}
func (dst *Command) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Command.Merge(dst, src)
}
func (m *Command) XXX_Size() int {
	return xxx_messageInfo_Command.Size(m)
}
func (m *Command) XXX_DiscardUnknown() {
	xxx_messageInfo_Command.DiscardUnknown(m)
}

var xxx_messageInfo_Command proto.InternalMessageInfo

func (m *Command) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Command) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Command) GetRequest() string {
	if m != nil {
		return m.Request
	}
	return ""
}

func (m *Command) GetArguments() string {
	if m != nil {
		return m.Arguments
	}
	return ""
}

func (m *Command) GetWaitPeriod() int32 {
	if m != nil {
		return m.WaitPeriod
	}
	return 0
}

func (m *Command) GetLinesInResponse() int32 {
	if m != nil {
		return m.LinesInResponse
	}
	return 0
}

// Result is the structure containing information on the command execution result.
type Result struct {
	// request is the command that was executed.
	Request string `protobuf:"bytes,1,opt,name=request" json:"request,omitempty"`
	// response_code is either `OK` or `ERROR`, which is received from the device upon executing the command.
	ResponseCode Result_ResCode `protobuf:"varint,2,opt,name=response_code,json=responseCode,enum=api.Result_ResCode" json:"response_code,omitempty"`
	// response is the string recevied from the device.
	Response             string   `protobuf:"bytes,3,opt,name=response" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_command_9ee6ba3dcbd6590a, []int{1}
}
func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (dst *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(dst, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetRequest() string {
	if m != nil {
		return m.Request
	}
	return ""
}

func (m *Result) GetResponseCode() Result_ResCode {
	if m != nil {
		return m.ResponseCode
	}
	return Result_NONE
}

func (m *Result) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

func init() {
	proto.RegisterType((*Command)(nil), "api.Command")
	proto.RegisterType((*Result)(nil), "api.Result")
	proto.RegisterEnum("api.Result_ResCode", Result_ResCode_name, Result_ResCode_value)
}

func init() { proto.RegisterFile("api/command.proto", fileDescriptor_command_9ee6ba3dcbd6590a) }

var fileDescriptor_command_9ee6ba3dcbd6590a = []byte{
	// 317 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xcf, 0x4e, 0x32, 0x31,
	0x14, 0xc5, 0xbf, 0xf2, 0x67, 0x80, 0xcb, 0xa7, 0x42, 0xdd, 0x34, 0xc6, 0x44, 0xc2, 0xc2, 0x10,
	0xa3, 0x33, 0x89, 0x6e, 0x5c, 0x4b, 0x58, 0x10, 0x12, 0x30, 0x5d, 0xba, 0x99, 0x94, 0x99, 0x1b,
	0x68, 0x64, 0xda, 0xda, 0x76, 0xa2, 0xbe, 0x8d, 0x0f, 0xe3, 0x83, 0x19, 0x0a, 0x83, 0xb8, 0x6a,
	0xcf, 0xef, 0xdc, 0xdb, 0x9c, 0x93, 0x42, 0x5f, 0x18, 0x99, 0x64, 0xba, 0x28, 0x84, 0xca, 0x63,
	0x63, 0xb5, 0xd7, 0xb4, 0x2e, 0x8c, 0x1c, 0x7e, 0x13, 0x68, 0x8d, 0x77, 0x98, 0x52, 0x68, 0x28,
	0x51, 0x20, 0x23, 0x03, 0x32, 0xea, 0xf0, 0x70, 0xa7, 0x03, 0xe8, 0xe6, 0xe8, 0x32, 0x2b, 0x8d,
	0x97, 0x5a, 0xb1, 0x5a, 0xb0, 0x8e, 0x11, 0x65, 0xd0, 0xb2, 0xf8, 0x56, 0xa2, 0xf3, 0xac, 0x1e,
	0xdc, 0x4a, 0xd2, 0x4b, 0xe8, 0x08, 0xbb, 0x2a, 0x0b, 0x54, 0xde, 0xb1, 0x46, 0xf0, 0x7e, 0x01,
	0xbd, 0x82, 0xee, 0xbb, 0x90, 0x3e, 0x35, 0x68, 0xa5, 0xce, 0x59, 0x73, 0x40, 0x46, 0x4d, 0x0e,
	0x5b, 0xf4, 0x1c, 0x08, 0xbd, 0x81, 0xfe, 0x46, 0x2a, 0x74, 0xa9, 0x54, 0xa9, 0x45, 0x67, 0xb4,
	0x72, 0xc8, 0xa2, 0x30, 0x76, 0x16, 0x8c, 0xa9, 0xe2, 0x7b, 0x3c, 0xfc, 0x22, 0x10, 0x71, 0x74,
	0xe5, 0xc6, 0x1f, 0xe7, 0x21, 0x7f, 0xf3, 0x3c, 0xc2, 0x49, 0xf5, 0x4e, 0x9a, 0xe9, 0x1c, 0x43,
	0x9b, 0xd3, 0xfb, 0xf3, 0x58, 0x18, 0x19, 0xef, 0xb6, 0xb7, 0xc7, 0x58, 0xe7, 0xc8, 0xff, 0x57,
	0x93, 0x5b, 0x45, 0x2f, 0xa0, 0x7d, 0x48, 0xb0, 0x2b, 0x79, 0xd0, 0xc3, 0x6b, 0x68, 0xed, 0x97,
	0x68, 0x1b, 0x1a, 0xf3, 0xc5, 0x7c, 0xd2, 0xfb, 0x47, 0x23, 0xa8, 0x2d, 0x66, 0x3d, 0x42, 0x3b,
	0xd0, 0x9c, 0x70, 0xbe, 0xe0, 0xbd, 0xda, 0x53, 0xfc, 0x72, 0xbb, 0x92, 0x7e, 0x5d, 0x2e, 0xe3,
	0x4c, 0x17, 0xc9, 0xcc, 0x4a, 0xb7, 0x56, 0x62, 0xfa, 0x89, 0x36, 0xf9, 0xc8, 0xb5, 0xbf, 0x5b,
	0xe3, 0xc6, 0xa0, 0x4d, 0xcc, 0xeb, 0x2a, 0x31, 0x4b, 0x61, 0xe4, 0x32, 0x0a, 0xbf, 0xf4, 0xf0,
	0x13, 0x00, 0x00, 0xff, 0xff, 0x84, 0x15, 0x0d, 0xfd, 0xba, 0x01, 0x00, 0x00,
}
