// Code generated by protoc-gen-go. DO NOT EDIT.
// source: variflight.proto

package protobuf

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type VariFlightData struct {
	Fcategory             string   `protobuf:"bytes,1,opt,name=fcategory,proto3" json:"fcategory,omitempty"`
	FlightNo              string   `protobuf:"bytes,2,opt,name=FlightNo,proto3" json:"FlightNo,omitempty"`
	FlightCompany         string   `protobuf:"bytes,3,opt,name=FlightCompany,proto3" json:"FlightCompany,omitempty"`
	FlightDepcode         string   `protobuf:"bytes,4,opt,name=FlightDepcode,proto3" json:"FlightDepcode,omitempty"`
	FlightArrcode         string   `protobuf:"bytes,5,opt,name=FlightArrcode,proto3" json:"FlightArrcode,omitempty"`
	FlightDeptimePlanDate string   `protobuf:"bytes,6,opt,name=FlightDeptimePlanDate,proto3" json:"FlightDeptimePlanDate,omitempty"`
	FlightArrtimePlanDate string   `protobuf:"bytes,7,opt,name=FlightArrtimePlanDate,proto3" json:"FlightArrtimePlanDate,omitempty"`
	FlightDeptimeDate     string   `protobuf:"bytes,8,opt,name=FlightDeptimeDate,proto3" json:"FlightDeptimeDate,omitempty"`
	FlightArrtimeDate     string   `protobuf:"bytes,9,opt,name=FlightArrtimeDate,proto3" json:"FlightArrtimeDate,omitempty"`
	FlightState           string   `protobuf:"bytes,10,opt,name=FlightState,proto3" json:"FlightState,omitempty"`
	FlightHTerminal       string   `protobuf:"bytes,11,opt,name=FlightHTerminal,proto3" json:"FlightHTerminal,omitempty"`
	FlightTerminal        string   `protobuf:"bytes,12,opt,name=FlightTerminal,proto3" json:"FlightTerminal,omitempty"`
	OrgTimezone           string   `protobuf:"bytes,13,opt,name=org_timezone,json=orgTimezone,proto3" json:"org_timezone,omitempty"`
	DstTimezone           string   `protobuf:"bytes,14,opt,name=dst_timezone,json=dstTimezone,proto3" json:"dst_timezone,omitempty"`
	ShareFlightNo         string   `protobuf:"bytes,15,opt,name=ShareFlightNo,proto3" json:"ShareFlightNo,omitempty"`
	StopFlag              string   `protobuf:"bytes,16,opt,name=StopFlag,proto3" json:"StopFlag,omitempty"`
	ShareFlag             string   `protobuf:"bytes,17,opt,name=ShareFlag,proto3" json:"ShareFlag,omitempty"`
	VirtualFlag           string   `protobuf:"bytes,18,opt,name=VirtualFlag,proto3" json:"VirtualFlag,omitempty"`
	LegFlag               string   `protobuf:"bytes,19,opt,name=LegFlag,proto3" json:"LegFlag,omitempty"`
	FlightDep             string   `protobuf:"bytes,20,opt,name=FlightDep,proto3" json:"FlightDep,omitempty"`
	FlightArr             string   `protobuf:"bytes,21,opt,name=FlightArr,proto3" json:"FlightArr,omitempty"`
	FlightDepAirport      string   `protobuf:"bytes,22,opt,name=FlightDepAirport,proto3" json:"FlightDepAirport,omitempty"`
	FlightArrAirport      string   `protobuf:"bytes,23,opt,name=FlightArrAirport,proto3" json:"FlightArrAirport,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *VariFlightData) Reset()         { *m = VariFlightData{} }
func (m *VariFlightData) String() string { return proto.CompactTextString(m) }
func (*VariFlightData) ProtoMessage()    {}
func (*VariFlightData) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b7cf0666aebe7ef, []int{0}
}

func (m *VariFlightData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VariFlightData.Unmarshal(m, b)
}
func (m *VariFlightData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VariFlightData.Marshal(b, m, deterministic)
}
func (m *VariFlightData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VariFlightData.Merge(m, src)
}
func (m *VariFlightData) XXX_Size() int {
	return xxx_messageInfo_VariFlightData.Size(m)
}
func (m *VariFlightData) XXX_DiscardUnknown() {
	xxx_messageInfo_VariFlightData.DiscardUnknown(m)
}

var xxx_messageInfo_VariFlightData proto.InternalMessageInfo

func (m *VariFlightData) GetFcategory() string {
	if m != nil {
		return m.Fcategory
	}
	return ""
}

func (m *VariFlightData) GetFlightNo() string {
	if m != nil {
		return m.FlightNo
	}
	return ""
}

func (m *VariFlightData) GetFlightCompany() string {
	if m != nil {
		return m.FlightCompany
	}
	return ""
}

func (m *VariFlightData) GetFlightDepcode() string {
	if m != nil {
		return m.FlightDepcode
	}
	return ""
}

func (m *VariFlightData) GetFlightArrcode() string {
	if m != nil {
		return m.FlightArrcode
	}
	return ""
}

func (m *VariFlightData) GetFlightDeptimePlanDate() string {
	if m != nil {
		return m.FlightDeptimePlanDate
	}
	return ""
}

func (m *VariFlightData) GetFlightArrtimePlanDate() string {
	if m != nil {
		return m.FlightArrtimePlanDate
	}
	return ""
}

func (m *VariFlightData) GetFlightDeptimeDate() string {
	if m != nil {
		return m.FlightDeptimeDate
	}
	return ""
}

func (m *VariFlightData) GetFlightArrtimeDate() string {
	if m != nil {
		return m.FlightArrtimeDate
	}
	return ""
}

func (m *VariFlightData) GetFlightState() string {
	if m != nil {
		return m.FlightState
	}
	return ""
}

func (m *VariFlightData) GetFlightHTerminal() string {
	if m != nil {
		return m.FlightHTerminal
	}
	return ""
}

func (m *VariFlightData) GetFlightTerminal() string {
	if m != nil {
		return m.FlightTerminal
	}
	return ""
}

func (m *VariFlightData) GetOrgTimezone() string {
	if m != nil {
		return m.OrgTimezone
	}
	return ""
}

func (m *VariFlightData) GetDstTimezone() string {
	if m != nil {
		return m.DstTimezone
	}
	return ""
}

func (m *VariFlightData) GetShareFlightNo() string {
	if m != nil {
		return m.ShareFlightNo
	}
	return ""
}

func (m *VariFlightData) GetStopFlag() string {
	if m != nil {
		return m.StopFlag
	}
	return ""
}

func (m *VariFlightData) GetShareFlag() string {
	if m != nil {
		return m.ShareFlag
	}
	return ""
}

func (m *VariFlightData) GetVirtualFlag() string {
	if m != nil {
		return m.VirtualFlag
	}
	return ""
}

func (m *VariFlightData) GetLegFlag() string {
	if m != nil {
		return m.LegFlag
	}
	return ""
}

func (m *VariFlightData) GetFlightDep() string {
	if m != nil {
		return m.FlightDep
	}
	return ""
}

func (m *VariFlightData) GetFlightArr() string {
	if m != nil {
		return m.FlightArr
	}
	return ""
}

func (m *VariFlightData) GetFlightDepAirport() string {
	if m != nil {
		return m.FlightDepAirport
	}
	return ""
}

func (m *VariFlightData) GetFlightArrAirport() string {
	if m != nil {
		return m.FlightArrAirport
	}
	return ""
}

type GetFlightDataByFlightNumberRequest struct {
	FlightNumber         string   `protobuf:"bytes,1,opt,name=flightNumber,proto3" json:"flightNumber,omitempty"`
	Date                 string   `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFlightDataByFlightNumberRequest) Reset()         { *m = GetFlightDataByFlightNumberRequest{} }
func (m *GetFlightDataByFlightNumberRequest) String() string { return proto.CompactTextString(m) }
func (*GetFlightDataByFlightNumberRequest) ProtoMessage()    {}
func (*GetFlightDataByFlightNumberRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b7cf0666aebe7ef, []int{1}
}

func (m *GetFlightDataByFlightNumberRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFlightDataByFlightNumberRequest.Unmarshal(m, b)
}
func (m *GetFlightDataByFlightNumberRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFlightDataByFlightNumberRequest.Marshal(b, m, deterministic)
}
func (m *GetFlightDataByFlightNumberRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFlightDataByFlightNumberRequest.Merge(m, src)
}
func (m *GetFlightDataByFlightNumberRequest) XXX_Size() int {
	return xxx_messageInfo_GetFlightDataByFlightNumberRequest.Size(m)
}
func (m *GetFlightDataByFlightNumberRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFlightDataByFlightNumberRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetFlightDataByFlightNumberRequest proto.InternalMessageInfo

func (m *GetFlightDataByFlightNumberRequest) GetFlightNumber() string {
	if m != nil {
		return m.FlightNumber
	}
	return ""
}

func (m *GetFlightDataByFlightNumberRequest) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

type GetFlightDataBetweenTwoAirportsRequest struct {
	DepartureAirport     string   `protobuf:"bytes,1,opt,name=departureAirport,proto3" json:"departureAirport,omitempty"`
	ArrivalAirport       string   `protobuf:"bytes,2,opt,name=arrivalAirport,proto3" json:"arrivalAirport,omitempty"`
	Date                 string   `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFlightDataBetweenTwoAirportsRequest) Reset() {
	*m = GetFlightDataBetweenTwoAirportsRequest{}
}
func (m *GetFlightDataBetweenTwoAirportsRequest) String() string { return proto.CompactTextString(m) }
func (*GetFlightDataBetweenTwoAirportsRequest) ProtoMessage()    {}
func (*GetFlightDataBetweenTwoAirportsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b7cf0666aebe7ef, []int{2}
}

func (m *GetFlightDataBetweenTwoAirportsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFlightDataBetweenTwoAirportsRequest.Unmarshal(m, b)
}
func (m *GetFlightDataBetweenTwoAirportsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFlightDataBetweenTwoAirportsRequest.Marshal(b, m, deterministic)
}
func (m *GetFlightDataBetweenTwoAirportsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFlightDataBetweenTwoAirportsRequest.Merge(m, src)
}
func (m *GetFlightDataBetweenTwoAirportsRequest) XXX_Size() int {
	return xxx_messageInfo_GetFlightDataBetweenTwoAirportsRequest.Size(m)
}
func (m *GetFlightDataBetweenTwoAirportsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFlightDataBetweenTwoAirportsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetFlightDataBetweenTwoAirportsRequest proto.InternalMessageInfo

func (m *GetFlightDataBetweenTwoAirportsRequest) GetDepartureAirport() string {
	if m != nil {
		return m.DepartureAirport
	}
	return ""
}

func (m *GetFlightDataBetweenTwoAirportsRequest) GetArrivalAirport() string {
	if m != nil {
		return m.ArrivalAirport
	}
	return ""
}

func (m *GetFlightDataBetweenTwoAirportsRequest) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

type GetFlightDataBetweenTwoCitiesRequest struct {
	DepartureCity        string   `protobuf:"bytes,1,opt,name=departureCity,proto3" json:"departureCity,omitempty"`
	ArrivalCity          string   `protobuf:"bytes,2,opt,name=arrivalCity,proto3" json:"arrivalCity,omitempty"`
	Date                 string   `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFlightDataBetweenTwoCitiesRequest) Reset()         { *m = GetFlightDataBetweenTwoCitiesRequest{} }
func (m *GetFlightDataBetweenTwoCitiesRequest) String() string { return proto.CompactTextString(m) }
func (*GetFlightDataBetweenTwoCitiesRequest) ProtoMessage()    {}
func (*GetFlightDataBetweenTwoCitiesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b7cf0666aebe7ef, []int{3}
}

func (m *GetFlightDataBetweenTwoCitiesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFlightDataBetweenTwoCitiesRequest.Unmarshal(m, b)
}
func (m *GetFlightDataBetweenTwoCitiesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFlightDataBetweenTwoCitiesRequest.Marshal(b, m, deterministic)
}
func (m *GetFlightDataBetweenTwoCitiesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFlightDataBetweenTwoCitiesRequest.Merge(m, src)
}
func (m *GetFlightDataBetweenTwoCitiesRequest) XXX_Size() int {
	return xxx_messageInfo_GetFlightDataBetweenTwoCitiesRequest.Size(m)
}
func (m *GetFlightDataBetweenTwoCitiesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFlightDataBetweenTwoCitiesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetFlightDataBetweenTwoCitiesRequest proto.InternalMessageInfo

func (m *GetFlightDataBetweenTwoCitiesRequest) GetDepartureCity() string {
	if m != nil {
		return m.DepartureCity
	}
	return ""
}

func (m *GetFlightDataBetweenTwoCitiesRequest) GetArrivalCity() string {
	if m != nil {
		return m.ArrivalCity
	}
	return ""
}

func (m *GetFlightDataBetweenTwoCitiesRequest) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

type GetFlightDataAtOneAirportByStatusRequest struct {
	Airport              string   `protobuf:"bytes,1,opt,name=airport,proto3" json:"airport,omitempty"`
	Status               string   `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Date                 string   `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFlightDataAtOneAirportByStatusRequest) Reset() {
	*m = GetFlightDataAtOneAirportByStatusRequest{}
}
func (m *GetFlightDataAtOneAirportByStatusRequest) String() string { return proto.CompactTextString(m) }
func (*GetFlightDataAtOneAirportByStatusRequest) ProtoMessage()    {}
func (*GetFlightDataAtOneAirportByStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b7cf0666aebe7ef, []int{4}
}

func (m *GetFlightDataAtOneAirportByStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFlightDataAtOneAirportByStatusRequest.Unmarshal(m, b)
}
func (m *GetFlightDataAtOneAirportByStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFlightDataAtOneAirportByStatusRequest.Marshal(b, m, deterministic)
}
func (m *GetFlightDataAtOneAirportByStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFlightDataAtOneAirportByStatusRequest.Merge(m, src)
}
func (m *GetFlightDataAtOneAirportByStatusRequest) XXX_Size() int {
	return xxx_messageInfo_GetFlightDataAtOneAirportByStatusRequest.Size(m)
}
func (m *GetFlightDataAtOneAirportByStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFlightDataAtOneAirportByStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetFlightDataAtOneAirportByStatusRequest proto.InternalMessageInfo

func (m *GetFlightDataAtOneAirportByStatusRequest) GetAirport() string {
	if m != nil {
		return m.Airport
	}
	return ""
}

func (m *GetFlightDataAtOneAirportByStatusRequest) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *GetFlightDataAtOneAirportByStatusRequest) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

func init() {
	proto.RegisterType((*VariFlightData)(nil), "protobuf.VariFlightData")
	proto.RegisterType((*GetFlightDataByFlightNumberRequest)(nil), "protobuf.GetFlightDataByFlightNumberRequest")
	proto.RegisterType((*GetFlightDataBetweenTwoAirportsRequest)(nil), "protobuf.GetFlightDataBetweenTwoAirportsRequest")
	proto.RegisterType((*GetFlightDataBetweenTwoCitiesRequest)(nil), "protobuf.GetFlightDataBetweenTwoCitiesRequest")
	proto.RegisterType((*GetFlightDataAtOneAirportByStatusRequest)(nil), "protobuf.GetFlightDataAtOneAirportByStatusRequest")
}

func init() { proto.RegisterFile("variflight.proto", fileDescriptor_0b7cf0666aebe7ef) }

var fileDescriptor_0b7cf0666aebe7ef = []byte{
	// 652 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0xd1, 0x4e, 0xdb, 0x30,
	0x14, 0x86, 0xd7, 0xc1, 0x0a, 0x1c, 0x28, 0x14, 0x6f, 0x30, 0x8b, 0x6d, 0x1a, 0xab, 0x2a, 0x54,
	0x21, 0x54, 0x21, 0xb6, 0x17, 0x08, 0x54, 0x6c, 0x17, 0xd3, 0x36, 0x51, 0xc4, 0xd5, 0xa4, 0xc9,
	0x50, 0x37, 0x58, 0x4b, 0xe3, 0xec, 0xc4, 0x01, 0x75, 0x17, 0xbb, 0xd8, 0x13, 0xec, 0x39, 0x77,
	0xb1, 0x67, 0x98, 0x62, 0x27, 0x8e, 0x5d, 0x28, 0xb9, 0x6a, 0xfc, 0x9f, 0xef, 0xf8, 0x3f, 0x8d,
	0x7f, 0x07, 0xda, 0x37, 0x0c, 0xc5, 0x38, 0x12, 0xe1, 0xb5, 0xea, 0x27, 0x28, 0x95, 0x24, 0xcb,
	0xfa, 0xe7, 0x32, 0x1b, 0x77, 0xfe, 0x35, 0x61, 0xfd, 0x82, 0xa1, 0x38, 0xd5, 0xe5, 0x01, 0x53,
	0x8c, 0xbc, 0x84, 0x95, 0xf1, 0x15, 0x53, 0x3c, 0x94, 0x38, 0xa5, 0x8d, 0xdd, 0x46, 0x6f, 0xe5,
	0xac, 0x12, 0xc8, 0x0e, 0x2c, 0x1b, 0xf6, 0x93, 0xa4, 0x8f, 0x75, 0xd1, 0xae, 0x49, 0x17, 0x5a,
	0xe6, 0xf9, 0x44, 0x4e, 0x12, 0x16, 0x4f, 0xe9, 0x82, 0x06, 0x7c, 0xb1, 0xa2, 0x06, 0x3c, 0xb9,
	0x92, 0x23, 0x4e, 0x17, 0x5d, 0xaa, 0x10, 0x2b, 0x2a, 0x40, 0xd4, 0xd4, 0x13, 0x97, 0x2a, 0x44,
	0xf2, 0x0e, 0xb6, 0x6c, 0x9b, 0x12, 0x13, 0xfe, 0x25, 0x62, 0xf1, 0x80, 0x29, 0x4e, 0x9b, 0x9a,
	0xbe, 0xbf, 0x58, 0x75, 0x05, 0x88, 0x5e, 0xd7, 0x92, 0xdb, 0x35, 0x53, 0x24, 0x07, 0xb0, 0xe9,
	0x6d, 0xa7, 0x3b, 0x96, 0x75, 0xc7, 0xdd, 0x42, 0x45, 0x17, 0xdb, 0x68, 0x7a, 0xc5, 0xa5, 0x9d,
	0x02, 0xd9, 0x85, 0x55, 0x23, 0x0e, 0x55, 0xce, 0x81, 0xe6, 0x5c, 0x89, 0xf4, 0x60, 0xc3, 0x2c,
	0x3f, 0x9c, 0x73, 0x9c, 0x88, 0x98, 0x45, 0x74, 0x55, 0x53, 0xb3, 0x32, 0xd9, 0x83, 0x75, 0x23,
	0x59, 0x70, 0x4d, 0x83, 0x33, 0x2a, 0x79, 0x03, 0x6b, 0x12, 0xc3, 0x6f, 0xf9, 0x0c, 0x3f, 0x65,
	0xcc, 0x69, 0xcb, 0x98, 0x4a, 0x0c, 0xcf, 0x0b, 0x29, 0x47, 0x46, 0xa9, 0xaa, 0x90, 0x75, 0x83,
	0x8c, 0x52, 0x65, 0x91, 0x2e, 0xb4, 0x86, 0xd7, 0x0c, 0xb9, 0x0d, 0xc5, 0x86, 0x39, 0x27, 0x4f,
	0xcc, 0x53, 0x33, 0x54, 0x32, 0x39, 0x8d, 0x58, 0x48, 0xdb, 0x26, 0x35, 0xe5, 0x3a, 0xcf, 0x5b,
	0x01, 0xb3, 0x90, 0x6e, 0x9a, 0xbc, 0x59, 0x21, 0x7f, 0x33, 0x17, 0x02, 0x55, 0xc6, 0x22, 0x5d,
	0x27, 0x66, 0x02, 0x47, 0x22, 0x14, 0x96, 0x3e, 0xf2, 0x50, 0x57, 0x9f, 0xea, 0x6a, 0xb9, 0xcc,
	0x77, 0xb6, 0x07, 0x43, 0x9f, 0x99, 0x9d, 0xad, 0x50, 0x55, 0x03, 0x44, 0xba, 0xe5, 0x56, 0x03,
	0x44, 0xb2, 0x0f, 0x6d, 0x8b, 0x06, 0x02, 0x13, 0x89, 0x8a, 0x6e, 0x6b, 0xe8, 0x8e, 0x5e, 0xb1,
	0x01, 0x62, 0xc9, 0x3e, 0x77, 0xd9, 0x4a, 0xef, 0x7c, 0x85, 0xce, 0x7b, 0xae, 0xaa, 0xeb, 0x76,
	0x3c, 0x2d, 0x5e, 0x52, 0x36, 0xb9, 0xe4, 0x78, 0xc6, 0x7f, 0x64, 0x3c, 0x55, 0xa4, 0x03, 0x6b,
	0x63, 0x47, 0x2e, 0xae, 0xa1, 0xa7, 0x11, 0x02, 0x8b, 0xa3, 0x3c, 0x2c, 0xe6, 0x16, 0xea, 0xe7,
	0xce, 0x9f, 0x06, 0xec, 0xf9, 0xdb, 0x73, 0x75, 0xcb, 0x79, 0x7c, 0x7e, 0x2b, 0x8b, 0x01, 0xd2,
	0xd2, 0x62, 0x1f, 0xda, 0x23, 0x9e, 0x30, 0x54, 0x19, 0xf2, 0x72, 0x68, 0x63, 0x73, 0x47, 0xcf,
	0x23, 0xc5, 0x10, 0xc5, 0x0d, 0x8b, 0x4a, 0xd2, 0x98, 0xce, 0xa8, 0x76, 0xa4, 0x05, 0x67, 0xa4,
	0xdf, 0x0d, 0xe8, 0xce, 0x19, 0xe9, 0x44, 0x28, 0xc1, 0xed, 0x40, 0x5d, 0x68, 0x59, 0xe3, 0x13,
	0xa1, 0xca, 0x6f, 0x8f, 0x2f, 0xe6, 0x79, 0x28, 0x4c, 0x35, 0x63, 0xe6, 0x70, 0xa5, 0x7b, 0x87,
	0x48, 0xa0, 0xe7, 0xcd, 0x10, 0xa8, 0xcf, 0x71, 0xf9, 0xef, 0x8e, 0xa7, 0xf9, 0x0d, 0xcb, 0xec,
	0x1c, 0x14, 0x96, 0x98, 0xf7, 0x3e, 0xca, 0x25, 0xd9, 0x86, 0x66, 0xaa, 0xd1, 0xc2, 0xb6, 0x58,
	0xdd, 0xe7, 0x78, 0xf4, 0x77, 0x01, 0xb6, 0xfc, 0x0f, 0xeb, 0x90, 0xe3, 0x8d, 0xb8, 0xe2, 0xe4,
	0x3b, 0xbc, 0x78, 0x20, 0x01, 0xe4, 0xa0, 0x5f, 0x7e, 0x9c, 0xfb, 0xf5, 0x41, 0xd9, 0xa1, 0x15,
	0xed, 0xbb, 0x75, 0x1e, 0xf5, 0x1a, 0x87, 0x0d, 0x92, 0xc2, 0xeb, 0x9a, 0x3c, 0x90, 0xc3, 0x79,
	0x86, 0xf3, 0xa2, 0x53, 0x6b, 0x2a, 0xe1, 0xd5, 0x83, 0x27, 0x4e, 0xfa, 0xb5, 0x96, 0x5e, 0x34,
	0x6a, 0x0d, 0x7f, 0xcd, 0x1c, 0xef, 0xf1, 0x74, 0x60, 0x23, 0x1c, 0x8f, 0x02, 0x13, 0x0e, 0x73,
	0xca, 0xe4, 0x68, 0x8e, 0xf7, 0x03, 0x91, 0xa8, 0xf3, 0xbf, 0x6c, 0xea, 0xf2, 0xdb, 0xff, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x03, 0x7e, 0x9f, 0xf7, 0x6a, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// VariFlightDataServiceClient is the client API for VariFlightDataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VariFlightDataServiceClient interface {
	GetFlightDataByFlightNumber(ctx context.Context, opts ...grpc.CallOption) (VariFlightDataService_GetFlightDataByFlightNumberClient, error)
	GetFlightDataBetweenTwoAirports(ctx context.Context, opts ...grpc.CallOption) (VariFlightDataService_GetFlightDataBetweenTwoAirportsClient, error)
	GetFlightDataBetweenTwoCities(ctx context.Context, opts ...grpc.CallOption) (VariFlightDataService_GetFlightDataBetweenTwoCitiesClient, error)
	GetFlightDataByDepartureAndArrivalStatus(ctx context.Context, opts ...grpc.CallOption) (VariFlightDataService_GetFlightDataByDepartureAndArrivalStatusClient, error)
}

type variFlightDataServiceClient struct {
	cc *grpc.ClientConn
}

func NewVariFlightDataServiceClient(cc *grpc.ClientConn) VariFlightDataServiceClient {
	return &variFlightDataServiceClient{cc}
}

func (c *variFlightDataServiceClient) GetFlightDataByFlightNumber(ctx context.Context, opts ...grpc.CallOption) (VariFlightDataService_GetFlightDataByFlightNumberClient, error) {
	stream, err := c.cc.NewStream(ctx, &_VariFlightDataService_serviceDesc.Streams[0], "/protobuf.VariFlightDataService/GetFlightDataByFlightNumber", opts...)
	if err != nil {
		return nil, err
	}
	x := &variFlightDataServiceGetFlightDataByFlightNumberClient{stream}
	return x, nil
}

type VariFlightDataService_GetFlightDataByFlightNumberClient interface {
	Send(*GetFlightDataByFlightNumberRequest) error
	Recv() (*VariFlightData, error)
	grpc.ClientStream
}

type variFlightDataServiceGetFlightDataByFlightNumberClient struct {
	grpc.ClientStream
}

func (x *variFlightDataServiceGetFlightDataByFlightNumberClient) Send(m *GetFlightDataByFlightNumberRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *variFlightDataServiceGetFlightDataByFlightNumberClient) Recv() (*VariFlightData, error) {
	m := new(VariFlightData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *variFlightDataServiceClient) GetFlightDataBetweenTwoAirports(ctx context.Context, opts ...grpc.CallOption) (VariFlightDataService_GetFlightDataBetweenTwoAirportsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_VariFlightDataService_serviceDesc.Streams[1], "/protobuf.VariFlightDataService/GetFlightDataBetweenTwoAirports", opts...)
	if err != nil {
		return nil, err
	}
	x := &variFlightDataServiceGetFlightDataBetweenTwoAirportsClient{stream}
	return x, nil
}

type VariFlightDataService_GetFlightDataBetweenTwoAirportsClient interface {
	Send(*GetFlightDataBetweenTwoAirportsRequest) error
	Recv() (*VariFlightData, error)
	grpc.ClientStream
}

type variFlightDataServiceGetFlightDataBetweenTwoAirportsClient struct {
	grpc.ClientStream
}

func (x *variFlightDataServiceGetFlightDataBetweenTwoAirportsClient) Send(m *GetFlightDataBetweenTwoAirportsRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *variFlightDataServiceGetFlightDataBetweenTwoAirportsClient) Recv() (*VariFlightData, error) {
	m := new(VariFlightData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *variFlightDataServiceClient) GetFlightDataBetweenTwoCities(ctx context.Context, opts ...grpc.CallOption) (VariFlightDataService_GetFlightDataBetweenTwoCitiesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_VariFlightDataService_serviceDesc.Streams[2], "/protobuf.VariFlightDataService/GetFlightDataBetweenTwoCities", opts...)
	if err != nil {
		return nil, err
	}
	x := &variFlightDataServiceGetFlightDataBetweenTwoCitiesClient{stream}
	return x, nil
}

type VariFlightDataService_GetFlightDataBetweenTwoCitiesClient interface {
	Send(*GetFlightDataBetweenTwoCitiesRequest) error
	Recv() (*VariFlightData, error)
	grpc.ClientStream
}

type variFlightDataServiceGetFlightDataBetweenTwoCitiesClient struct {
	grpc.ClientStream
}

func (x *variFlightDataServiceGetFlightDataBetweenTwoCitiesClient) Send(m *GetFlightDataBetweenTwoCitiesRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *variFlightDataServiceGetFlightDataBetweenTwoCitiesClient) Recv() (*VariFlightData, error) {
	m := new(VariFlightData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *variFlightDataServiceClient) GetFlightDataByDepartureAndArrivalStatus(ctx context.Context, opts ...grpc.CallOption) (VariFlightDataService_GetFlightDataByDepartureAndArrivalStatusClient, error) {
	stream, err := c.cc.NewStream(ctx, &_VariFlightDataService_serviceDesc.Streams[3], "/protobuf.VariFlightDataService/GetFlightDataByDepartureAndArrivalStatus", opts...)
	if err != nil {
		return nil, err
	}
	x := &variFlightDataServiceGetFlightDataByDepartureAndArrivalStatusClient{stream}
	return x, nil
}

type VariFlightDataService_GetFlightDataByDepartureAndArrivalStatusClient interface {
	Send(*GetFlightDataAtOneAirportByStatusRequest) error
	Recv() (*VariFlightData, error)
	grpc.ClientStream
}

type variFlightDataServiceGetFlightDataByDepartureAndArrivalStatusClient struct {
	grpc.ClientStream
}

func (x *variFlightDataServiceGetFlightDataByDepartureAndArrivalStatusClient) Send(m *GetFlightDataAtOneAirportByStatusRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *variFlightDataServiceGetFlightDataByDepartureAndArrivalStatusClient) Recv() (*VariFlightData, error) {
	m := new(VariFlightData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// VariFlightDataServiceServer is the server API for VariFlightDataService service.
type VariFlightDataServiceServer interface {
	GetFlightDataByFlightNumber(VariFlightDataService_GetFlightDataByFlightNumberServer) error
	GetFlightDataBetweenTwoAirports(VariFlightDataService_GetFlightDataBetweenTwoAirportsServer) error
	GetFlightDataBetweenTwoCities(VariFlightDataService_GetFlightDataBetweenTwoCitiesServer) error
	GetFlightDataByDepartureAndArrivalStatus(VariFlightDataService_GetFlightDataByDepartureAndArrivalStatusServer) error
}

// UnimplementedVariFlightDataServiceServer can be embedded to have forward compatible implementations.
type UnimplementedVariFlightDataServiceServer struct {
}

func (*UnimplementedVariFlightDataServiceServer) GetFlightDataByFlightNumber(srv VariFlightDataService_GetFlightDataByFlightNumberServer) error {
	return status.Errorf(codes.Unimplemented, "method GetFlightDataByFlightNumber not implemented")
}
func (*UnimplementedVariFlightDataServiceServer) GetFlightDataBetweenTwoAirports(srv VariFlightDataService_GetFlightDataBetweenTwoAirportsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetFlightDataBetweenTwoAirports not implemented")
}
func (*UnimplementedVariFlightDataServiceServer) GetFlightDataBetweenTwoCities(srv VariFlightDataService_GetFlightDataBetweenTwoCitiesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetFlightDataBetweenTwoCities not implemented")
}
func (*UnimplementedVariFlightDataServiceServer) GetFlightDataByDepartureAndArrivalStatus(srv VariFlightDataService_GetFlightDataByDepartureAndArrivalStatusServer) error {
	return status.Errorf(codes.Unimplemented, "method GetFlightDataByDepartureAndArrivalStatus not implemented")
}

func RegisterVariFlightDataServiceServer(s *grpc.Server, srv VariFlightDataServiceServer) {
	s.RegisterService(&_VariFlightDataService_serviceDesc, srv)
}

func _VariFlightDataService_GetFlightDataByFlightNumber_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(VariFlightDataServiceServer).GetFlightDataByFlightNumber(&variFlightDataServiceGetFlightDataByFlightNumberServer{stream})
}

type VariFlightDataService_GetFlightDataByFlightNumberServer interface {
	Send(*VariFlightData) error
	Recv() (*GetFlightDataByFlightNumberRequest, error)
	grpc.ServerStream
}

type variFlightDataServiceGetFlightDataByFlightNumberServer struct {
	grpc.ServerStream
}

func (x *variFlightDataServiceGetFlightDataByFlightNumberServer) Send(m *VariFlightData) error {
	return x.ServerStream.SendMsg(m)
}

func (x *variFlightDataServiceGetFlightDataByFlightNumberServer) Recv() (*GetFlightDataByFlightNumberRequest, error) {
	m := new(GetFlightDataByFlightNumberRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _VariFlightDataService_GetFlightDataBetweenTwoAirports_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(VariFlightDataServiceServer).GetFlightDataBetweenTwoAirports(&variFlightDataServiceGetFlightDataBetweenTwoAirportsServer{stream})
}

type VariFlightDataService_GetFlightDataBetweenTwoAirportsServer interface {
	Send(*VariFlightData) error
	Recv() (*GetFlightDataBetweenTwoAirportsRequest, error)
	grpc.ServerStream
}

type variFlightDataServiceGetFlightDataBetweenTwoAirportsServer struct {
	grpc.ServerStream
}

func (x *variFlightDataServiceGetFlightDataBetweenTwoAirportsServer) Send(m *VariFlightData) error {
	return x.ServerStream.SendMsg(m)
}

func (x *variFlightDataServiceGetFlightDataBetweenTwoAirportsServer) Recv() (*GetFlightDataBetweenTwoAirportsRequest, error) {
	m := new(GetFlightDataBetweenTwoAirportsRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _VariFlightDataService_GetFlightDataBetweenTwoCities_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(VariFlightDataServiceServer).GetFlightDataBetweenTwoCities(&variFlightDataServiceGetFlightDataBetweenTwoCitiesServer{stream})
}

type VariFlightDataService_GetFlightDataBetweenTwoCitiesServer interface {
	Send(*VariFlightData) error
	Recv() (*GetFlightDataBetweenTwoCitiesRequest, error)
	grpc.ServerStream
}

type variFlightDataServiceGetFlightDataBetweenTwoCitiesServer struct {
	grpc.ServerStream
}

func (x *variFlightDataServiceGetFlightDataBetweenTwoCitiesServer) Send(m *VariFlightData) error {
	return x.ServerStream.SendMsg(m)
}

func (x *variFlightDataServiceGetFlightDataBetweenTwoCitiesServer) Recv() (*GetFlightDataBetweenTwoCitiesRequest, error) {
	m := new(GetFlightDataBetweenTwoCitiesRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _VariFlightDataService_GetFlightDataByDepartureAndArrivalStatus_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(VariFlightDataServiceServer).GetFlightDataByDepartureAndArrivalStatus(&variFlightDataServiceGetFlightDataByDepartureAndArrivalStatusServer{stream})
}

type VariFlightDataService_GetFlightDataByDepartureAndArrivalStatusServer interface {
	Send(*VariFlightData) error
	Recv() (*GetFlightDataAtOneAirportByStatusRequest, error)
	grpc.ServerStream
}

type variFlightDataServiceGetFlightDataByDepartureAndArrivalStatusServer struct {
	grpc.ServerStream
}

func (x *variFlightDataServiceGetFlightDataByDepartureAndArrivalStatusServer) Send(m *VariFlightData) error {
	return x.ServerStream.SendMsg(m)
}

func (x *variFlightDataServiceGetFlightDataByDepartureAndArrivalStatusServer) Recv() (*GetFlightDataAtOneAirportByStatusRequest, error) {
	m := new(GetFlightDataAtOneAirportByStatusRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _VariFlightDataService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.VariFlightDataService",
	HandlerType: (*VariFlightDataServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetFlightDataByFlightNumber",
			Handler:       _VariFlightDataService_GetFlightDataByFlightNumber_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "GetFlightDataBetweenTwoAirports",
			Handler:       _VariFlightDataService_GetFlightDataBetweenTwoAirports_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "GetFlightDataBetweenTwoCities",
			Handler:       _VariFlightDataService_GetFlightDataBetweenTwoCities_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "GetFlightDataByDepartureAndArrivalStatus",
			Handler:       _VariFlightDataService_GetFlightDataByDepartureAndArrivalStatus_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "variflight.proto",
}
