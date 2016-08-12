// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/bigtable/admin/v2/table.proto
// DO NOT EDIT!

package google_bigtable_admin_v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/serviceconfig"
import google_protobuf4 "github.com/golang/protobuf/ptypes/duration"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Possible timestamp granularities to use when keeping multiple versions
// of data in a table.
type Table_TimestampGranularity int32

const (
	// The user did not specify a granularity. Should not be returned.
	// When specified during table creation, MILLIS will be used.
	Table_TIMESTAMP_GRANULARITY_UNSPECIFIED Table_TimestampGranularity = 0
	// The table keeps data versioned at a granularity of 1ms.
	Table_MILLIS Table_TimestampGranularity = 1
)

var Table_TimestampGranularity_name = map[int32]string{
	0: "TIMESTAMP_GRANULARITY_UNSPECIFIED",
	1: "MILLIS",
}
var Table_TimestampGranularity_value = map[string]int32{
	"TIMESTAMP_GRANULARITY_UNSPECIFIED": 0,
	"MILLIS": 1,
}

func (x Table_TimestampGranularity) String() string {
	return proto.EnumName(Table_TimestampGranularity_name, int32(x))
}
func (Table_TimestampGranularity) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor4, []int{0, 0}
}

// Defines a view over a table's fields.
type Table_View int32

const (
	// Uses the default view for each method as documented in its request.
	Table_VIEW_UNSPECIFIED Table_View = 0
	// Only populates `name`.
	Table_NAME_ONLY Table_View = 1
	// Only populates `name` and fields related to the table's schema.
	Table_SCHEMA_VIEW Table_View = 2
	// Populates all fields.
	Table_FULL Table_View = 4
)

var Table_View_name = map[int32]string{
	0: "VIEW_UNSPECIFIED",
	1: "NAME_ONLY",
	2: "SCHEMA_VIEW",
	4: "FULL",
}
var Table_View_value = map[string]int32{
	"VIEW_UNSPECIFIED": 0,
	"NAME_ONLY":        1,
	"SCHEMA_VIEW":      2,
	"FULL":             4,
}

func (x Table_View) String() string {
	return proto.EnumName(Table_View_name, int32(x))
}
func (Table_View) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{0, 1} }

// A collection of user data indexed by row, column, and timestamp.
// Each table is served using the resources of its parent cluster.
type Table struct {
	// The unique name of the table. Values are of the form
	// projects/<project>/instances/<instance>/tables/[_a-zA-Z0-9][-_.a-zA-Z0-9]*
	// Views: NAME_ONLY, SCHEMA_VIEW, REPLICATION_VIEW, FULL
	// @OutputOnly
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The column families configured for this table, mapped by column family ID.
	// Views: SCHEMA_VIEW, FULL
	// @CreationOnly
	ColumnFamilies map[string]*ColumnFamily `protobuf:"bytes,3,rep,name=column_families,json=columnFamilies" json:"column_families,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// The granularity (e.g. MILLIS, MICROS) at which timestamps are stored in
	// this table. Timestamps not matching the granularity will be rejected.
	// If unspecified at creation time, the value will be set to MILLIS.
	// Views: SCHEMA_VIEW, FULL
	// @CreationOnly
	Granularity Table_TimestampGranularity `protobuf:"varint,4,opt,name=granularity,enum=google.bigtable.admin.v2.Table_TimestampGranularity" json:"granularity,omitempty"`
}

func (m *Table) Reset()                    { *m = Table{} }
func (m *Table) String() string            { return proto.CompactTextString(m) }
func (*Table) ProtoMessage()               {}
func (*Table) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *Table) GetColumnFamilies() map[string]*ColumnFamily {
	if m != nil {
		return m.ColumnFamilies
	}
	return nil
}

// A set of columns within a table which share a common configuration.
type ColumnFamily struct {
	// Garbage collection rule specified as a protobuf.
	// Must serialize to at most 500 bytes.
	//
	// NOTE: Garbage collection executes opportunistically in the background, and
	// so it's possible for reads to return a cell even if it matches the active
	// GC expression for its family.
	GcRule *GcRule `protobuf:"bytes,1,opt,name=gc_rule,json=gcRule" json:"gc_rule,omitempty"`
}

func (m *ColumnFamily) Reset()                    { *m = ColumnFamily{} }
func (m *ColumnFamily) String() string            { return proto.CompactTextString(m) }
func (*ColumnFamily) ProtoMessage()               {}
func (*ColumnFamily) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *ColumnFamily) GetGcRule() *GcRule {
	if m != nil {
		return m.GcRule
	}
	return nil
}

// Rule for determining which cells to delete during garbage collection.
type GcRule struct {
	// Types that are valid to be assigned to Rule:
	//	*GcRule_MaxNumVersions
	//	*GcRule_MaxAge
	//	*GcRule_Intersection_
	//	*GcRule_Union_
	Rule isGcRule_Rule `protobuf_oneof:"rule"`
}

func (m *GcRule) Reset()                    { *m = GcRule{} }
func (m *GcRule) String() string            { return proto.CompactTextString(m) }
func (*GcRule) ProtoMessage()               {}
func (*GcRule) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

type isGcRule_Rule interface {
	isGcRule_Rule()
}

type GcRule_MaxNumVersions struct {
	MaxNumVersions int32 `protobuf:"varint,1,opt,name=max_num_versions,json=maxNumVersions,oneof"`
}
type GcRule_MaxAge struct {
	MaxAge *google_protobuf4.Duration `protobuf:"bytes,2,opt,name=max_age,json=maxAge,oneof"`
}
type GcRule_Intersection_ struct {
	Intersection *GcRule_Intersection `protobuf:"bytes,3,opt,name=intersection,oneof"`
}
type GcRule_Union_ struct {
	Union *GcRule_Union `protobuf:"bytes,4,opt,name=union,oneof"`
}

func (*GcRule_MaxNumVersions) isGcRule_Rule() {}
func (*GcRule_MaxAge) isGcRule_Rule()         {}
func (*GcRule_Intersection_) isGcRule_Rule()  {}
func (*GcRule_Union_) isGcRule_Rule()         {}

func (m *GcRule) GetRule() isGcRule_Rule {
	if m != nil {
		return m.Rule
	}
	return nil
}

func (m *GcRule) GetMaxNumVersions() int32 {
	if x, ok := m.GetRule().(*GcRule_MaxNumVersions); ok {
		return x.MaxNumVersions
	}
	return 0
}

func (m *GcRule) GetMaxAge() *google_protobuf4.Duration {
	if x, ok := m.GetRule().(*GcRule_MaxAge); ok {
		return x.MaxAge
	}
	return nil
}

func (m *GcRule) GetIntersection() *GcRule_Intersection {
	if x, ok := m.GetRule().(*GcRule_Intersection_); ok {
		return x.Intersection
	}
	return nil
}

func (m *GcRule) GetUnion() *GcRule_Union {
	if x, ok := m.GetRule().(*GcRule_Union_); ok {
		return x.Union
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*GcRule) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _GcRule_OneofMarshaler, _GcRule_OneofUnmarshaler, _GcRule_OneofSizer, []interface{}{
		(*GcRule_MaxNumVersions)(nil),
		(*GcRule_MaxAge)(nil),
		(*GcRule_Intersection_)(nil),
		(*GcRule_Union_)(nil),
	}
}

func _GcRule_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*GcRule)
	// rule
	switch x := m.Rule.(type) {
	case *GcRule_MaxNumVersions:
		b.EncodeVarint(1<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.MaxNumVersions))
	case *GcRule_MaxAge:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.MaxAge); err != nil {
			return err
		}
	case *GcRule_Intersection_:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Intersection); err != nil {
			return err
		}
	case *GcRule_Union_:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Union); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("GcRule.Rule has unexpected type %T", x)
	}
	return nil
}

func _GcRule_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*GcRule)
	switch tag {
	case 1: // rule.max_num_versions
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Rule = &GcRule_MaxNumVersions{int32(x)}
		return true, err
	case 2: // rule.max_age
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(google_protobuf4.Duration)
		err := b.DecodeMessage(msg)
		m.Rule = &GcRule_MaxAge{msg}
		return true, err
	case 3: // rule.intersection
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(GcRule_Intersection)
		err := b.DecodeMessage(msg)
		m.Rule = &GcRule_Intersection_{msg}
		return true, err
	case 4: // rule.union
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(GcRule_Union)
		err := b.DecodeMessage(msg)
		m.Rule = &GcRule_Union_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _GcRule_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*GcRule)
	// rule
	switch x := m.Rule.(type) {
	case *GcRule_MaxNumVersions:
		n += proto.SizeVarint(1<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.MaxNumVersions))
	case *GcRule_MaxAge:
		s := proto.Size(x.MaxAge)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *GcRule_Intersection_:
		s := proto.Size(x.Intersection)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *GcRule_Union_:
		s := proto.Size(x.Union)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// A GcRule which deletes cells matching all of the given rules.
type GcRule_Intersection struct {
	// Only delete cells which would be deleted by every element of `rules`.
	Rules []*GcRule `protobuf:"bytes,1,rep,name=rules" json:"rules,omitempty"`
}

func (m *GcRule_Intersection) Reset()                    { *m = GcRule_Intersection{} }
func (m *GcRule_Intersection) String() string            { return proto.CompactTextString(m) }
func (*GcRule_Intersection) ProtoMessage()               {}
func (*GcRule_Intersection) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2, 0} }

func (m *GcRule_Intersection) GetRules() []*GcRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

// A GcRule which deletes cells matching any of the given rules.
type GcRule_Union struct {
	// Delete cells which would be deleted by any element of `rules`.
	Rules []*GcRule `protobuf:"bytes,1,rep,name=rules" json:"rules,omitempty"`
}

func (m *GcRule_Union) Reset()                    { *m = GcRule_Union{} }
func (m *GcRule_Union) String() string            { return proto.CompactTextString(m) }
func (*GcRule_Union) ProtoMessage()               {}
func (*GcRule_Union) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2, 1} }

func (m *GcRule_Union) GetRules() []*GcRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

func init() {
	proto.RegisterType((*Table)(nil), "google.bigtable.admin.v2.Table")
	proto.RegisterType((*ColumnFamily)(nil), "google.bigtable.admin.v2.ColumnFamily")
	proto.RegisterType((*GcRule)(nil), "google.bigtable.admin.v2.GcRule")
	proto.RegisterType((*GcRule_Intersection)(nil), "google.bigtable.admin.v2.GcRule.Intersection")
	proto.RegisterType((*GcRule_Union)(nil), "google.bigtable.admin.v2.GcRule.Union")
	proto.RegisterEnum("google.bigtable.admin.v2.Table_TimestampGranularity", Table_TimestampGranularity_name, Table_TimestampGranularity_value)
	proto.RegisterEnum("google.bigtable.admin.v2.Table_View", Table_View_name, Table_View_value)
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/bigtable/admin/v2/table.proto", fileDescriptor4)
}

var fileDescriptor4 = []byte{
	// 606 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x93, 0xed, 0x6e, 0x12, 0x4d,
	0x14, 0xc7, 0x4b, 0x59, 0xe8, 0xd3, 0x43, 0x9f, 0x76, 0x33, 0xf6, 0x03, 0x12, 0x3f, 0x54, 0x12,
	0x4d, 0x63, 0x74, 0x27, 0xa1, 0x8d, 0xb1, 0xc6, 0x68, 0x28, 0x5d, 0xda, 0x4d, 0x00, 0xc9, 0xf2,
	0x62, 0x9a, 0x98, 0x6c, 0x86, 0xed, 0x74, 0x9c, 0xb8, 0x3b, 0x4b, 0xf6, 0x05, 0xe5, 0x2e, 0xbc,
	0x33, 0x6f, 0xc9, 0x99, 0x59, 0x50, 0x6a, 0x4a, 0x68, 0xfc, 0x02, 0xb3, 0x73, 0xfe, 0xff, 0xdf,
	0x79, 0xd9, 0xb3, 0xd0, 0x62, 0x51, 0xc4, 0x02, 0x6a, 0xb1, 0x28, 0x20, 0x82, 0x59, 0x51, 0xcc,
	0x30, 0xa3, 0x62, 0x1a, 0x47, 0x69, 0x84, 0xf3, 0x10, 0x99, 0xf2, 0x04, 0x4f, 0x38, 0x4b, 0xc9,
	0x24, 0xa0, 0x98, 0xdc, 0x84, 0x5c, 0xe0, 0x59, 0x03, 0xeb, 0x47, 0x4b, 0x0b, 0x51, 0x75, 0x01,
	0x59, 0xaa, 0x2c, 0xad, 0xb2, 0x66, 0x8d, 0x9a, 0xf3, 0x30, 0xbc, 0xfc, 0xc1, 0x09, 0x8d, 0x67,
	0xdc, 0xa7, 0x7e, 0x24, 0x6e, 0x39, 0xc3, 0x44, 0x88, 0x28, 0x25, 0x29, 0x8f, 0x44, 0x92, 0x27,
	0xa9, 0x9d, 0x31, 0x9e, 0x7e, 0xc9, 0x26, 0x96, 0x1f, 0x85, 0x38, 0xc7, 0x61, 0x1d, 0x98, 0x64,
	0xb7, 0x78, 0x9a, 0xce, 0xa7, 0x34, 0xc1, 0x37, 0x59, 0xac, 0x2d, 0xbf, 0x0f, 0xb9, 0xb5, 0xfe,
	0xb3, 0x08, 0xa5, 0xa1, 0x2a, 0x0c, 0x21, 0x30, 0x04, 0x09, 0x69, 0xb5, 0x70, 0x54, 0x38, 0xde,
	0x75, 0xf5, 0x19, 0x7d, 0x86, 0x03, 0x3f, 0x0a, 0xb2, 0x50, 0x78, 0xb7, 0x24, 0xe4, 0x01, 0xa7,
	0x49, 0xb5, 0x78, 0x54, 0x3c, 0xae, 0x34, 0x4e, 0xac, 0x75, 0x7d, 0x59, 0x9a, 0x66, 0xb5, 0xb4,
	0xad, 0xbd, 0x70, 0xd9, 0x22, 0x8d, 0xe7, 0xee, 0xbe, 0x7f, 0xe7, 0x12, 0x8d, 0xa1, 0xc2, 0x62,
	0x22, 0xb2, 0x80, 0xc4, 0x3c, 0x9d, 0x57, 0x0d, 0x99, 0x78, 0xbf, 0x71, 0xba, 0x89, 0x3c, 0xe4,
	0x21, 0x4d, 0x52, 0x12, 0x4e, 0x2f, 0xff, 0x78, 0xdd, 0x55, 0x50, 0x8d, 0xc3, 0xa3, 0x7b, 0xd2,
	0x23, 0x13, 0x8a, 0x5f, 0xe9, 0x7c, 0xd1, 0x9f, 0x3a, 0xa2, 0x77, 0x50, 0x9a, 0x91, 0x20, 0xa3,
	0xd5, 0x6d, 0x79, 0x57, 0x69, 0x3c, 0x5f, 0x9f, 0x7a, 0x85, 0x37, 0x77, 0x73, 0xd3, 0xdb, 0xed,
	0x37, 0x85, 0xba, 0x03, 0x87, 0xf7, 0xd5, 0x83, 0x9e, 0xc1, 0xd3, 0xa1, 0xd3, 0xb5, 0x07, 0xc3,
	0x66, 0xb7, 0xef, 0x5d, 0xba, 0xcd, 0xde, 0xa8, 0xd3, 0x74, 0x9d, 0xe1, 0xb5, 0x37, 0xea, 0x0d,
	0xfa, 0x76, 0xcb, 0x69, 0x3b, 0xf6, 0x85, 0xb9, 0x85, 0x00, 0xca, 0x5d, 0xa7, 0xd3, 0x71, 0x06,
	0x66, 0xa1, 0xde, 0x06, 0x63, 0xcc, 0xe9, 0x37, 0x74, 0x08, 0xe6, 0xd8, 0xb1, 0x3f, 0xfd, 0xa5,
	0xfc, 0x1f, 0x76, 0x7b, 0xcd, 0xae, 0xed, 0x7d, 0xec, 0x75, 0xae, 0xcd, 0x02, 0x3a, 0x80, 0xca,
	0xa0, 0x75, 0x65, 0x77, 0x9b, 0x9e, 0xd2, 0x9a, 0xdb, 0xe8, 0x3f, 0x30, 0xda, 0xa3, 0x4e, 0xc7,
	0x34, 0x64, 0x49, 0x7b, 0xab, 0xd5, 0xa2, 0x33, 0xd8, 0x61, 0xbe, 0x17, 0x67, 0x41, 0xfe, 0x6a,
	0x2b, 0x8d, 0xa3, 0xf5, 0x6d, 0x5e, 0xfa, 0xae, 0xd4, 0xb9, 0x65, 0xa6, 0xff, 0xeb, 0x3f, 0x8a,
	0x50, 0xce, 0xaf, 0xd0, 0x0b, 0x30, 0x43, 0xf2, 0xdd, 0x13, 0x59, 0xe8, 0xcd, 0x68, 0x9c, 0xa8,
	0xe5, 0xd3, 0xb8, 0xd2, 0xd5, 0x96, 0xbb, 0x2f, 0x23, 0xbd, 0x2c, 0x1c, 0x2f, 0xee, 0xd1, 0x29,
	0xec, 0x28, 0x2d, 0x61, 0xcb, 0xc1, 0x3e, 0x5e, 0x66, 0x5c, 0x6e, 0xa5, 0x75, 0xb1, 0xd8, 0x42,
	0xe9, 0x2e, 0x4b, 0x6d, 0x93, 0x51, 0x34, 0x80, 0x3d, 0x2e, 0x52, 0xc9, 0xa0, 0xbe, 0x8a, 0xc8,
	0x45, 0x53, 0xd6, 0x57, 0x9b, 0x8a, 0xb5, 0x9c, 0x15, 0x93, 0xc4, 0xdd, 0x81, 0xa0, 0xf7, 0x50,
	0xca, 0x84, 0xa2, 0x19, 0x9b, 0xde, 0xf0, 0x82, 0x36, 0x12, 0x39, 0x26, 0xb7, 0xd5, 0xda, 0xb0,
	0xb7, 0xca, 0x47, 0xaf, 0xa1, 0xa4, 0x26, 0xa9, 0x7a, 0x2f, 0x3e, 0x68, 0x94, 0xb9, 0xbc, 0xf6,
	0x01, 0x4a, 0x9a, 0xfc, 0xaf, 0x80, 0xf3, 0x32, 0x18, 0xea, 0x70, 0xfe, 0x12, 0x9e, 0xc8, 0xaf,
	0x7c, 0xad, 0xeb, 0x1c, 0xf4, 0x47, 0xd2, 0x57, 0x73, 0xee, 0x17, 0x26, 0x65, 0x3d, 0xf0, 0x93,
	0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xaf, 0x0f, 0xf7, 0x8b, 0xcb, 0x04, 0x00, 0x00,
}
