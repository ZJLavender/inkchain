// Code generated by protoc-gen-go. DO NOT EDIT.
// source: peer/query.proto

package peer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// ChaincodeQueryResponse returns information about each chaincode that pertains
// to a query in lscc.go, such as GetChaincodes (returns all chaincodes
// instantiated on a channel), and GetInstalledChaincodes (returns all chaincodes
// installed on a peer)
type ChaincodeQueryResponse struct {
	Chaincodes []*ChaincodeInfo `protobuf:"bytes,1,rep,name=chaincodes" json:"chaincodes,omitempty"`
}

func (m *ChaincodeQueryResponse) Reset()                    { *m = ChaincodeQueryResponse{} }
func (m *ChaincodeQueryResponse) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeQueryResponse) ProtoMessage()               {}
func (*ChaincodeQueryResponse) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{0} }

func (m *ChaincodeQueryResponse) GetChaincodes() []*ChaincodeInfo {
	if m != nil {
		return m.Chaincodes
	}
	return nil
}

// ChaincodeInfo contains general information about an installed/instantiated
// chaincode
type ChaincodeInfo struct {
	Name    string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Version string `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
	// the path as specified by the install/instantiate transaction
	Path string `protobuf:"bytes,3,opt,name=path" json:"path,omitempty"`
	// the chaincode function upon instantiation and its arguments. This will be
	// blank if the query is returning information about installed chaincodes.
	Input string `protobuf:"bytes,4,opt,name=input" json:"input,omitempty"`
	// the name of the ESCC for this chaincode. This will be
	// blank if the query is returning information about installed chaincodes.
	Escc string `protobuf:"bytes,5,opt,name=escc" json:"escc,omitempty"`
	// the name of the VSCC for this chaincode. This will be
	// blank if the query is returning information about installed chaincodes.
	Vscc string `protobuf:"bytes,6,opt,name=vscc" json:"vscc,omitempty"`
}

func (m *ChaincodeInfo) Reset()                    { *m = ChaincodeInfo{} }
func (m *ChaincodeInfo) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeInfo) ProtoMessage()               {}
func (*ChaincodeInfo) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{1} }

func (m *ChaincodeInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ChaincodeInfo) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ChaincodeInfo) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *ChaincodeInfo) GetInput() string {
	if m != nil {
		return m.Input
	}
	return ""
}

func (m *ChaincodeInfo) GetEscc() string {
	if m != nil {
		return m.Escc
	}
	return ""
}

func (m *ChaincodeInfo) GetVscc() string {
	if m != nil {
		return m.Vscc
	}
	return ""
}

// ChannelQueryResponse returns information about each channel that pertains
// to a query in lscc.go, such as GetChannels (returns all channels for a
// given peer)
type ChannelQueryResponse struct {
	Channels []*ChannelInfo `protobuf:"bytes,1,rep,name=channels" json:"channels,omitempty"`
}

func (m *ChannelQueryResponse) Reset()                    { *m = ChannelQueryResponse{} }
func (m *ChannelQueryResponse) String() string            { return proto.CompactTextString(m) }
func (*ChannelQueryResponse) ProtoMessage()               {}
func (*ChannelQueryResponse) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{2} }

func (m *ChannelQueryResponse) GetChannels() []*ChannelInfo {
	if m != nil {
		return m.Channels
	}
	return nil
}

// ChannelInfo contains general information about channels
type ChannelInfo struct {
	ChannelId string `protobuf:"bytes,1,opt,name=channel_id,json=channelId" json:"channel_id,omitempty"`
}

func (m *ChannelInfo) Reset()                    { *m = ChannelInfo{} }
func (m *ChannelInfo) String() string            { return proto.CompactTextString(m) }
func (*ChannelInfo) ProtoMessage()               {}
func (*ChannelInfo) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{3} }

func (m *ChannelInfo) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func init() {
	proto.RegisterType((*ChaincodeQueryResponse)(nil), "protos.ChaincodeQueryResponse")
	proto.RegisterType((*ChaincodeInfo)(nil), "protos.ChaincodeInfo")
	proto.RegisterType((*ChannelQueryResponse)(nil), "protos.ChannelQueryResponse")
	proto.RegisterType((*ChannelInfo)(nil), "protos.ChannelInfo")
}

func init() { proto.RegisterFile("peer/query.proto", fileDescriptor9) }

var fileDescriptor9 = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0x5f, 0x4b, 0xc3, 0x30,
	0x14, 0xc5, 0xa9, 0xfb, 0xa3, 0xbb, 0x43, 0x90, 0x38, 0x25, 0x2f, 0xc2, 0xe8, 0xd3, 0x10, 0x69,
	0x50, 0xf1, 0x0b, 0xb8, 0x07, 0xd9, 0x93, 0xb8, 0x47, 0x5f, 0x46, 0x9a, 0x66, 0x6b, 0x70, 0xbb,
	0xa9, 0x4d, 0x3b, 0xf0, 0x53, 0xf8, 0x95, 0xe5, 0x26, 0xcb, 0xe8, 0xf0, 0xa9, 0xf7, 0xfe, 0xce,
	0x39, 0x94, 0x93, 0x0b, 0x57, 0x95, 0xd6, 0xb5, 0xf8, 0x6e, 0x75, 0xfd, 0x93, 0x55, 0xb5, 0x6d,
	0x2c, 0x1b, 0xfa, 0x8f, 0x4b, 0xdf, 0xe1, 0x76, 0x5e, 0x4a, 0x83, 0xca, 0x16, 0xfa, 0x83, 0xf4,
	0xa5, 0x76, 0x95, 0x45, 0xa7, 0xd9, 0x0b, 0x80, 0x8a, 0x8a, 0xe3, 0xc9, 0xb4, 0x37, 0x1b, 0x3f,
	0xdd, 0x84, 0xb4, 0xcb, 0x8e, 0x99, 0x05, 0xae, 0xed, 0xb2, 0x63, 0x4c, 0x7f, 0x13, 0xb8, 0x3c,
	0x51, 0x19, 0x83, 0x3e, 0xca, 0x9d, 0xe6, 0xc9, 0x34, 0x99, 0x8d, 0x96, 0x7e, 0x66, 0x1c, 0xce,
	0xf7, 0xba, 0x76, 0xc6, 0x22, 0x3f, 0xf3, 0x38, 0xae, 0xe4, 0xae, 0x64, 0x53, 0xf2, 0x5e, 0x70,
	0xd3, 0xcc, 0x26, 0x30, 0x30, 0x58, 0xb5, 0x0d, 0xef, 0x7b, 0x18, 0x16, 0x72, 0x6a, 0xa7, 0x14,
	0x1f, 0x04, 0x27, 0xcd, 0xc4, 0xf6, 0xc4, 0x86, 0x81, 0xd1, 0x9c, 0xbe, 0xc1, 0x64, 0x5e, 0x4a,
	0x44, 0xbd, 0x3d, 0x2d, 0x28, 0xe0, 0x42, 0x05, 0x1e, 0xeb, 0x5d, 0x77, 0xea, 0x11, 0xf7, 0xe5,
	0x8e, 0xa6, 0xf4, 0x01, 0xc6, 0x1d, 0x81, 0xdd, 0xf9, 0x07, 0xa2, 0x75, 0x65, 0x8a, 0x43, 0xbb,
	0xd1, 0x81, 0x2c, 0x8a, 0xd7, 0x15, 0xdc, 0xdb, 0x7a, 0x93, 0x19, 0xfc, 0xda, 0xca, 0xdc, 0xad,
	0x6d, 0x8b, 0x85, 0x6c, 0x8c, 0x45, 0x22, 0xfe, 0xbd, 0xe2, 0xcf, 0xe8, 0x36, 0x9f, 0x8f, 0x1b,
	0xd3, 0x94, 0x6d, 0x9e, 0x29, 0xbb, 0x13, 0xff, 0x22, 0x22, 0x46, 0x44, 0x88, 0x08, 0x8a, 0xe4,
	0xe1, 0x84, 0xcf, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xfe, 0x61, 0x50, 0xe3, 0xdd, 0x01, 0x00,
	0x00,
}
