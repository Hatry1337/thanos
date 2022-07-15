// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: delete_plan.proto

package purger

import (
	fmt "fmt"
	_ "github.com/thanos-io/thanos/internal/cortex/cortexpb"
	github_com_cortexproject_cortex_pkg_cortexpb "github.com/thanos-io/thanos/internal/cortex/cortexpb"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
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

// DeletePlan holds all the chunks that are supposed to be deleted within an interval(usually a day)
// This Proto file is used just for storing Delete Plans in proto format.
type DeletePlan struct {
	PlanInterval *Interval     `protobuf:"bytes,1,opt,name=plan_interval,json=planInterval,proto3" json:"plan_interval,omitempty"`
	ChunksGroup  []ChunksGroup `protobuf:"bytes,2,rep,name=chunks_group,json=chunksGroup,proto3" json:"chunks_group"`
}

func (m *DeletePlan) Reset()      { *m = DeletePlan{} }
func (*DeletePlan) ProtoMessage() {}
func (*DeletePlan) Descriptor() ([]byte, []int) {
	return fileDescriptor_c38868cf63b27372, []int{0}
}
func (m *DeletePlan) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DeletePlan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DeletePlan.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DeletePlan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePlan.Merge(m, src)
}
func (m *DeletePlan) XXX_Size() int {
	return m.Size()
}
func (m *DeletePlan) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePlan.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePlan proto.InternalMessageInfo

func (m *DeletePlan) GetPlanInterval() *Interval {
	if m != nil {
		return m.PlanInterval
	}
	return nil
}

func (m *DeletePlan) GetChunksGroup() []ChunksGroup {
	if m != nil {
		return m.ChunksGroup
	}
	return nil
}

// ChunksGroup holds ChunkDetails and Labels for a group of chunks which have same series ID
type ChunksGroup struct {
	Labels []github_com_cortexproject_cortex_pkg_cortexpb.LabelAdapter `protobuf:"bytes,1,rep,name=labels,proto3,customtype=github.com/thanos-io/thanos/internal/cortex/cortexpb.LabelAdapter" json:"labels"`
	Chunks []ChunkDetails                                              `protobuf:"bytes,2,rep,name=chunks,proto3" json:"chunks"`
}

func (m *ChunksGroup) Reset()      { *m = ChunksGroup{} }
func (*ChunksGroup) ProtoMessage() {}
func (*ChunksGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_c38868cf63b27372, []int{1}
}
func (m *ChunksGroup) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChunksGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChunksGroup.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ChunksGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChunksGroup.Merge(m, src)
}
func (m *ChunksGroup) XXX_Size() int {
	return m.Size()
}
func (m *ChunksGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_ChunksGroup.DiscardUnknown(m)
}

var xxx_messageInfo_ChunksGroup proto.InternalMessageInfo

func (m *ChunksGroup) GetChunks() []ChunkDetails {
	if m != nil {
		return m.Chunks
	}
	return nil
}

type ChunkDetails struct {
	ID                       string    `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	PartiallyDeletedInterval *Interval `protobuf:"bytes,2,opt,name=partially_deleted_interval,json=partiallyDeletedInterval,proto3" json:"partially_deleted_interval,omitempty"`
}

func (m *ChunkDetails) Reset()      { *m = ChunkDetails{} }
func (*ChunkDetails) ProtoMessage() {}
func (*ChunkDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_c38868cf63b27372, []int{2}
}
func (m *ChunkDetails) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChunkDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChunkDetails.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ChunkDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChunkDetails.Merge(m, src)
}
func (m *ChunkDetails) XXX_Size() int {
	return m.Size()
}
func (m *ChunkDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_ChunkDetails.DiscardUnknown(m)
}

var xxx_messageInfo_ChunkDetails proto.InternalMessageInfo

func (m *ChunkDetails) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *ChunkDetails) GetPartiallyDeletedInterval() *Interval {
	if m != nil {
		return m.PartiallyDeletedInterval
	}
	return nil
}

type Interval struct {
	StartTimestampMs int64 `protobuf:"varint,1,opt,name=start_timestamp_ms,json=startTimestampMs,proto3" json:"start_timestamp_ms,omitempty"`
	EndTimestampMs   int64 `protobuf:"varint,2,opt,name=end_timestamp_ms,json=endTimestampMs,proto3" json:"end_timestamp_ms,omitempty"`
}

func (m *Interval) Reset()      { *m = Interval{} }
func (*Interval) ProtoMessage() {}
func (*Interval) Descriptor() ([]byte, []int) {
	return fileDescriptor_c38868cf63b27372, []int{3}
}
func (m *Interval) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Interval) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Interval.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Interval) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Interval.Merge(m, src)
}
func (m *Interval) XXX_Size() int {
	return m.Size()
}
func (m *Interval) XXX_DiscardUnknown() {
	xxx_messageInfo_Interval.DiscardUnknown(m)
}

var xxx_messageInfo_Interval proto.InternalMessageInfo

func (m *Interval) GetStartTimestampMs() int64 {
	if m != nil {
		return m.StartTimestampMs
	}
	return 0
}

func (m *Interval) GetEndTimestampMs() int64 {
	if m != nil {
		return m.EndTimestampMs
	}
	return 0
}

func init() {
	proto.RegisterType((*DeletePlan)(nil), "purgeplan.DeletePlan")
	proto.RegisterType((*ChunksGroup)(nil), "purgeplan.ChunksGroup")
	proto.RegisterType((*ChunkDetails)(nil), "purgeplan.ChunkDetails")
	proto.RegisterType((*Interval)(nil), "purgeplan.Interval")
}

func init() { proto.RegisterFile("delete_plan.proto", fileDescriptor_c38868cf63b27372) }

var fileDescriptor_c38868cf63b27372 = []byte{
	// 446 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x41, 0x8b, 0xd4, 0x30,
	0x18, 0x6d, 0xba, 0x52, 0xdc, 0x74, 0x5c, 0xd6, 0x2c, 0x68, 0x99, 0x43, 0x76, 0xe9, 0x69, 0x0e,
	0xda, 0x81, 0x15, 0x41, 0x41, 0x90, 0x1d, 0x0b, 0x32, 0xa0, 0xb0, 0x16, 0x4f, 0x5e, 0x4a, 0xda,
	0xc6, 0x6e, 0xdd, 0xb4, 0x89, 0x69, 0x2a, 0x7a, 0xf3, 0xe6, 0xd5, 0x9f, 0xe1, 0x0f, 0xf0, 0x47,
	0xec, 0x71, 0x8e, 0x8b, 0x87, 0xc1, 0xe9, 0x5c, 0x3c, 0xce, 0x4f, 0x90, 0xa6, 0xed, 0x4c, 0x15,
	0x3c, 0x78, 0xcb, 0xfb, 0xde, 0x7b, 0xc9, 0xcb, 0x4b, 0xe0, 0xed, 0x84, 0x32, 0xaa, 0x68, 0x28,
	0x18, 0x29, 0x3c, 0x21, 0xb9, 0xe2, 0x68, 0x5f, 0x54, 0x32, 0xa5, 0xcd, 0x60, 0x7c, 0x3f, 0xcd,
	0xd4, 0x45, 0x15, 0x79, 0x31, 0xcf, 0xa7, 0x29, 0x4f, 0xf9, 0x54, 0x2b, 0xa2, 0xea, 0xad, 0x46,
	0x1a, 0xe8, 0x55, 0xeb, 0x1c, 0x3f, 0x1e, 0xc8, 0x63, 0x2e, 0x15, 0xfd, 0x28, 0x24, 0x7f, 0x47,
	0x63, 0xd5, 0xa1, 0xa9, 0xb8, 0x4c, 0x7b, 0x22, 0xea, 0x16, 0xad, 0xd5, 0xfd, 0x02, 0x20, 0xf4,
	0x75, 0x94, 0x73, 0x46, 0x0a, 0xf4, 0x08, 0xde, 0x6a, 0x02, 0x84, 0x59, 0xa1, 0xa8, 0xfc, 0x40,
	0x98, 0x03, 0x4e, 0xc0, 0xc4, 0x3e, 0x3d, 0xf2, 0xb6, 0xd9, 0xbc, 0x79, 0x47, 0x05, 0xa3, 0x06,
	0xf6, 0x08, 0x3d, 0x85, 0xa3, 0xf8, 0xa2, 0x2a, 0x2e, 0xcb, 0x30, 0x95, 0xbc, 0x12, 0x8e, 0x79,
	0xb2, 0x37, 0xb1, 0x4f, 0xef, 0x0c, 0x8c, 0xcf, 0x34, 0xfd, 0xbc, 0x61, 0x67, 0x37, 0xae, 0x96,
	0xc7, 0x46, 0x60, 0xc7, 0xbb, 0x91, 0xfb, 0x1d, 0x40, 0x7b, 0x20, 0x41, 0x05, 0xb4, 0x18, 0x89,
	0x28, 0x2b, 0x1d, 0xa0, 0xb7, 0x3a, 0xf2, 0xfa, 0x1b, 0x78, 0x2f, 0x9a, 0xf9, 0x39, 0xc9, 0xe4,
	0xec, 0xac, 0xd9, 0xe7, 0xc7, 0xf2, 0xf8, 0xbf, 0x1a, 0x68, 0xfd, 0x67, 0x09, 0x11, 0x8a, 0xca,
	0xa0, 0x3b, 0x05, 0x3d, 0x84, 0x56, 0x1b, 0xa7, 0x8b, 0x7e, 0xf7, 0xef, 0xe8, 0x3e, 0x55, 0x24,
	0x63, 0x65, 0x97, 0xbd, 0x13, 0xbb, 0xef, 0xe1, 0x68, 0xc8, 0xa2, 0x03, 0x68, 0xce, 0x7d, 0x5d,
	0xdb, 0x7e, 0x60, 0xce, 0x7d, 0xf4, 0x0a, 0x8e, 0x05, 0x91, 0x2a, 0x23, 0x8c, 0x7d, 0x0a, 0xdb,
	0x47, 0x4f, 0x76, 0xf5, 0x9a, 0xff, 0xae, 0xd7, 0xd9, 0xda, 0xda, 0xf7, 0x49, 0x7a, 0xc6, 0x8d,
	0xe0, 0xcd, 0x6d, 0xed, 0xf7, 0x20, 0x2a, 0x15, 0x91, 0x2a, 0x54, 0x59, 0x4e, 0x4b, 0x45, 0x72,
	0x11, 0xe6, 0xa5, 0x3e, 0x7e, 0x2f, 0x38, 0xd4, 0xcc, 0xeb, 0x9e, 0x78, 0x59, 0xa2, 0x09, 0x3c,
	0xa4, 0x45, 0xf2, 0xa7, 0xd6, 0xd4, 0xda, 0x03, 0x5a, 0x24, 0x03, 0xe5, 0xec, 0xc9, 0x62, 0x85,
	0x8d, 0xeb, 0x15, 0x36, 0x36, 0x2b, 0x0c, 0x3e, 0xd7, 0x18, 0x7c, 0xab, 0x31, 0xb8, 0xaa, 0x31,
	0x58, 0xd4, 0x18, 0xfc, 0xac, 0x31, 0xf8, 0x55, 0x63, 0x63, 0x53, 0x63, 0xf0, 0x75, 0x8d, 0x8d,
	0xc5, 0x1a, 0x1b, 0xd7, 0x6b, 0x6c, 0xbc, 0xb1, 0xf4, 0x3d, 0x64, 0x64, 0xe9, 0xcf, 0xf5, 0xe0,
	0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf5, 0x46, 0x96, 0xf6, 0xe6, 0x02, 0x00, 0x00,
}

func (this *DeletePlan) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DeletePlan)
	if !ok {
		that2, ok := that.(DeletePlan)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.PlanInterval.Equal(that1.PlanInterval) {
		return false
	}
	if len(this.ChunksGroup) != len(that1.ChunksGroup) {
		return false
	}
	for i := range this.ChunksGroup {
		if !this.ChunksGroup[i].Equal(&that1.ChunksGroup[i]) {
			return false
		}
	}
	return true
}
func (this *ChunksGroup) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ChunksGroup)
	if !ok {
		that2, ok := that.(ChunksGroup)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Labels) != len(that1.Labels) {
		return false
	}
	for i := range this.Labels {
		if !this.Labels[i].Equal(that1.Labels[i]) {
			return false
		}
	}
	if len(this.Chunks) != len(that1.Chunks) {
		return false
	}
	for i := range this.Chunks {
		if !this.Chunks[i].Equal(&that1.Chunks[i]) {
			return false
		}
	}
	return true
}
func (this *ChunkDetails) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ChunkDetails)
	if !ok {
		that2, ok := that.(ChunkDetails)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ID != that1.ID {
		return false
	}
	if !this.PartiallyDeletedInterval.Equal(that1.PartiallyDeletedInterval) {
		return false
	}
	return true
}
func (this *Interval) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Interval)
	if !ok {
		that2, ok := that.(Interval)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.StartTimestampMs != that1.StartTimestampMs {
		return false
	}
	if this.EndTimestampMs != that1.EndTimestampMs {
		return false
	}
	return true
}
func (this *DeletePlan) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&purger.DeletePlan{")
	if this.PlanInterval != nil {
		s = append(s, "PlanInterval: "+fmt.Sprintf("%#v", this.PlanInterval)+",\n")
	}
	if this.ChunksGroup != nil {
		vs := make([]*ChunksGroup, len(this.ChunksGroup))
		for i := range vs {
			vs[i] = &this.ChunksGroup[i]
		}
		s = append(s, "ChunksGroup: "+fmt.Sprintf("%#v", vs)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ChunksGroup) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&purger.ChunksGroup{")
	s = append(s, "Labels: "+fmt.Sprintf("%#v", this.Labels)+",\n")
	if this.Chunks != nil {
		vs := make([]*ChunkDetails, len(this.Chunks))
		for i := range vs {
			vs[i] = &this.Chunks[i]
		}
		s = append(s, "Chunks: "+fmt.Sprintf("%#v", vs)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ChunkDetails) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&purger.ChunkDetails{")
	s = append(s, "ID: "+fmt.Sprintf("%#v", this.ID)+",\n")
	if this.PartiallyDeletedInterval != nil {
		s = append(s, "PartiallyDeletedInterval: "+fmt.Sprintf("%#v", this.PartiallyDeletedInterval)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Interval) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&purger.Interval{")
	s = append(s, "StartTimestampMs: "+fmt.Sprintf("%#v", this.StartTimestampMs)+",\n")
	s = append(s, "EndTimestampMs: "+fmt.Sprintf("%#v", this.EndTimestampMs)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringDeletePlan(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *DeletePlan) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeletePlan) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DeletePlan) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ChunksGroup) > 0 {
		for iNdEx := len(m.ChunksGroup) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ChunksGroup[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDeletePlan(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.PlanInterval != nil {
		{
			size, err := m.PlanInterval.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDeletePlan(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ChunksGroup) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChunksGroup) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChunksGroup) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Chunks) > 0 {
		for iNdEx := len(m.Chunks) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Chunks[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDeletePlan(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Labels) > 0 {
		for iNdEx := len(m.Labels) - 1; iNdEx >= 0; iNdEx-- {
			{
				size := m.Labels[iNdEx].Size()
				i -= size
				if _, err := m.Labels[iNdEx].MarshalTo(dAtA[i:]); err != nil {
					return 0, err
				}
				i = encodeVarintDeletePlan(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ChunkDetails) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChunkDetails) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChunkDetails) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PartiallyDeletedInterval != nil {
		{
			size, err := m.PartiallyDeletedInterval.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDeletePlan(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.ID) > 0 {
		i -= len(m.ID)
		copy(dAtA[i:], m.ID)
		i = encodeVarintDeletePlan(dAtA, i, uint64(len(m.ID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Interval) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Interval) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Interval) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.EndTimestampMs != 0 {
		i = encodeVarintDeletePlan(dAtA, i, uint64(m.EndTimestampMs))
		i--
		dAtA[i] = 0x10
	}
	if m.StartTimestampMs != 0 {
		i = encodeVarintDeletePlan(dAtA, i, uint64(m.StartTimestampMs))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintDeletePlan(dAtA []byte, offset int, v uint64) int {
	offset -= sovDeletePlan(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DeletePlan) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PlanInterval != nil {
		l = m.PlanInterval.Size()
		n += 1 + l + sovDeletePlan(uint64(l))
	}
	if len(m.ChunksGroup) > 0 {
		for _, e := range m.ChunksGroup {
			l = e.Size()
			n += 1 + l + sovDeletePlan(uint64(l))
		}
	}
	return n
}

func (m *ChunksGroup) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Labels) > 0 {
		for _, e := range m.Labels {
			l = e.Size()
			n += 1 + l + sovDeletePlan(uint64(l))
		}
	}
	if len(m.Chunks) > 0 {
		for _, e := range m.Chunks {
			l = e.Size()
			n += 1 + l + sovDeletePlan(uint64(l))
		}
	}
	return n
}

func (m *ChunkDetails) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovDeletePlan(uint64(l))
	}
	if m.PartiallyDeletedInterval != nil {
		l = m.PartiallyDeletedInterval.Size()
		n += 1 + l + sovDeletePlan(uint64(l))
	}
	return n
}

func (m *Interval) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.StartTimestampMs != 0 {
		n += 1 + sovDeletePlan(uint64(m.StartTimestampMs))
	}
	if m.EndTimestampMs != 0 {
		n += 1 + sovDeletePlan(uint64(m.EndTimestampMs))
	}
	return n
}

func sovDeletePlan(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDeletePlan(x uint64) (n int) {
	return sovDeletePlan(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *DeletePlan) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForChunksGroup := "[]ChunksGroup{"
	for _, f := range this.ChunksGroup {
		repeatedStringForChunksGroup += strings.Replace(strings.Replace(f.String(), "ChunksGroup", "ChunksGroup", 1), `&`, ``, 1) + ","
	}
	repeatedStringForChunksGroup += "}"
	s := strings.Join([]string{`&DeletePlan{`,
		`PlanInterval:` + strings.Replace(this.PlanInterval.String(), "Interval", "Interval", 1) + `,`,
		`ChunksGroup:` + repeatedStringForChunksGroup + `,`,
		`}`,
	}, "")
	return s
}
func (this *ChunksGroup) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForChunks := "[]ChunkDetails{"
	for _, f := range this.Chunks {
		repeatedStringForChunks += strings.Replace(strings.Replace(f.String(), "ChunkDetails", "ChunkDetails", 1), `&`, ``, 1) + ","
	}
	repeatedStringForChunks += "}"
	s := strings.Join([]string{`&ChunksGroup{`,
		`Labels:` + fmt.Sprintf("%v", this.Labels) + `,`,
		`Chunks:` + repeatedStringForChunks + `,`,
		`}`,
	}, "")
	return s
}
func (this *ChunkDetails) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ChunkDetails{`,
		`ID:` + fmt.Sprintf("%v", this.ID) + `,`,
		`PartiallyDeletedInterval:` + strings.Replace(this.PartiallyDeletedInterval.String(), "Interval", "Interval", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Interval) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Interval{`,
		`StartTimestampMs:` + fmt.Sprintf("%v", this.StartTimestampMs) + `,`,
		`EndTimestampMs:` + fmt.Sprintf("%v", this.EndTimestampMs) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringDeletePlan(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *DeletePlan) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeletePlan
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
			return fmt.Errorf("proto: DeletePlan: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeletePlan: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlanInterval", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeletePlan
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
				return ErrInvalidLengthDeletePlan
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDeletePlan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PlanInterval == nil {
				m.PlanInterval = &Interval{}
			}
			if err := m.PlanInterval.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChunksGroup", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeletePlan
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
				return ErrInvalidLengthDeletePlan
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDeletePlan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChunksGroup = append(m.ChunksGroup, ChunksGroup{})
			if err := m.ChunksGroup[len(m.ChunksGroup)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDeletePlan(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDeletePlan
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDeletePlan
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
func (m *ChunksGroup) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeletePlan
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
			return fmt.Errorf("proto: ChunksGroup: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChunksGroup: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Labels", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeletePlan
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
				return ErrInvalidLengthDeletePlan
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDeletePlan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Labels = append(m.Labels, github_com_cortexproject_cortex_pkg_cortexpb.LabelAdapter{})
			if err := m.Labels[len(m.Labels)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chunks", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeletePlan
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
				return ErrInvalidLengthDeletePlan
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDeletePlan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chunks = append(m.Chunks, ChunkDetails{})
			if err := m.Chunks[len(m.Chunks)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDeletePlan(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDeletePlan
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDeletePlan
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
func (m *ChunkDetails) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeletePlan
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
			return fmt.Errorf("proto: ChunkDetails: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChunkDetails: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeletePlan
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
				return ErrInvalidLengthDeletePlan
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeletePlan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PartiallyDeletedInterval", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeletePlan
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
				return ErrInvalidLengthDeletePlan
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDeletePlan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PartiallyDeletedInterval == nil {
				m.PartiallyDeletedInterval = &Interval{}
			}
			if err := m.PartiallyDeletedInterval.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDeletePlan(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDeletePlan
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDeletePlan
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
func (m *Interval) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeletePlan
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
			return fmt.Errorf("proto: Interval: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Interval: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTimestampMs", wireType)
			}
			m.StartTimestampMs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeletePlan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartTimestampMs |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTimestampMs", wireType)
			}
			m.EndTimestampMs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeletePlan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndTimestampMs |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDeletePlan(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDeletePlan
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDeletePlan
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
func skipDeletePlan(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDeletePlan
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
					return 0, ErrIntOverflowDeletePlan
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
					return 0, ErrIntOverflowDeletePlan
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
				return 0, ErrInvalidLengthDeletePlan
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthDeletePlan
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowDeletePlan
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
				next, err := skipDeletePlan(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthDeletePlan
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
	ErrInvalidLengthDeletePlan = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDeletePlan   = fmt.Errorf("proto: integer overflow")
)
