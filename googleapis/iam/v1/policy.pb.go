// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/iam/v1/policy.proto
// DO NOT EDIT!

package google_iam_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// # Overview
//
// The `Policy` defines an access control policy language. It is used to
// define policies that are attached to resources like files, folders, VMs,
// etc.
//
//
// # Policy structure
//
// A `Policy` consists of a list of bindings. A `Binding` binds a set of members
// to a role, where the members include user accounts, user groups, user
// domains, and service accounts. A 'role' is a named set of permissions,
// defined by IAM. The definition of a role is outside the policy.
//
// A permission check first determines the roles that include the specified
// permission, and then determines if the principal specified is a
// member of a binding to at least one of these roles. The membership check is
// recursive when a group is bound to a role.
//
// Policy examples:
//
// ```
// {
//   "bindings": [
//     {
//       "role": "roles/owner",
//       "members": [
//         "user:mike@example.com",
//         "group:admins@example.com",
//         "domain:google.com",
//         "serviceAccount:frontend@example.iam.gserviceaccounts.com"]
//     },
//     {
//       "role": "roles/viewer",
//       "members": ["user:sean@example.com"]
//     }
//   ]
// }
// ```
type Policy struct {
	// The policy language version. The version of the policy is
	// represented by the etag. The default version is 0.
	Version int32 `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	// It is an error to specify multiple bindings for the same role.
	// It is an error to specify a binding with no members.
	Bindings []*Binding `protobuf:"bytes,4,rep,name=bindings" json:"bindings,omitempty"`
	// Can be used to perform a read-modify-write.
	Etag []byte `protobuf:"bytes,3,opt,name=etag,proto3" json:"etag,omitempty"`
}

func (m *Policy) Reset()                    { *m = Policy{} }
func (m *Policy) String() string            { return proto.CompactTextString(m) }
func (*Policy) ProtoMessage()               {}
func (*Policy) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Policy) GetBindings() []*Binding {
	if m != nil {
		return m.Bindings
	}
	return nil
}

// Associates members with roles. See below for allowed
// formats of members.
type Binding struct {
	// The name of the role to which the members should be bound.
	// Examples: "roles/viewer", "roles/editor", "roles/owner".
	// Required
	Role string `protobuf:"bytes,1,opt,name=role" json:"role,omitempty"`
	// Format of member entries:
	// 1. allUsers
	//    Matches any requesting principal (users, service accounts or anonymous).
	//
	// 2. allAuthenticatedUsers
	//    Matches any requesting authenticated principal (users or service
	//    accounts).
	//
	// 3. user:{emailid}
	//    A google user account using an email address.
	//    For example alice@gmail.com, joe@example.com
	//
	// 4. serviceAccount:{emailid}
	//    An service account email.
	//
	// 5. group:{emailid}
	//    A google group with an email address. For example
	//    auth-ti-cloud@google.com
	//
	// 6. domain:{domain}
	//    A Google Apps domain name.
	//    For example google.com, example.com
	Members []string `protobuf:"bytes,2,rep,name=members" json:"members,omitempty"`
}

func (m *Binding) Reset()                    { *m = Binding{} }
func (m *Binding) String() string            { return proto.CompactTextString(m) }
func (*Binding) ProtoMessage()               {}
func (*Binding) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func init() {
	proto.RegisterType((*Policy)(nil), "google.iam.v1.Policy")
	proto.RegisterType((*Binding)(nil), "google.iam.v1.Binding")
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/iam/v1/policy.proto", fileDescriptor1)
}

var fileDescriptor1 = []byte{
	// 210 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x8f, 0xc1, 0x6a, 0x84, 0x30,
	0x10, 0x86, 0xb1, 0xbb, 0xdd, 0xed, 0x66, 0xdb, 0x43, 0x73, 0x28, 0x39, 0x8a, 0x27, 0x4f, 0x09,
	0xda, 0x43, 0xe9, 0xd5, 0x27, 0x90, 0xbc, 0x41, 0xb4, 0x61, 0x48, 0x49, 0x32, 0x12, 0x45, 0xe8,
	0xdb, 0x37, 0x46, 0x5d, 0xf0, 0x36, 0xc3, 0x37, 0x33, 0xdf, 0xfc, 0xe4, 0x1b, 0x10, 0xc1, 0x6a,
	0x0e, 0x68, 0x95, 0x07, 0x8e, 0x01, 0x04, 0x68, 0x3f, 0x04, 0x9c, 0x50, 0xac, 0x48, 0x0d, 0x66,
	0x14, 0x46, 0x39, 0x31, 0x57, 0x62, 0x40, 0x6b, 0xfa, 0x3f, 0x9e, 0x30, 0x7d, 0xdb, 0x56, 0x23,
	0xe3, 0x73, 0x55, 0xfc, 0x92, 0x4b, 0x9b, 0x30, 0x65, 0xe4, 0x3a, 0xeb, 0x30, 0x1a, 0xf4, 0x2c,
	0xcb, 0xb3, 0xf2, 0x59, 0xee, 0x2d, 0xad, 0xc9, 0x4b, 0x67, 0xfc, 0x8f, 0xf1, 0x30, 0xb2, 0x73,
	0x7e, 0x2a, 0xef, 0xf5, 0x07, 0x3f, 0x5c, 0xe1, 0xcd, 0x8a, 0xe5, 0x63, 0x8e, 0x52, 0x72, 0xd6,
	0x93, 0x02, 0x76, 0x8a, 0xa7, 0x5e, 0x65, 0xaa, 0x8b, 0x2f, 0x72, 0xdd, 0x06, 0x17, 0x1c, 0xd0,
	0xea, 0x64, 0xba, 0xc9, 0x54, 0x2f, 0x0f, 0x38, 0xed, 0xba, 0x28, 0x65, 0x4f, 0xd1, 0x72, 0x93,
	0x7b, 0xdb, 0x14, 0xe4, 0xbd, 0x47, 0x77, 0x74, 0x36, 0xf7, 0xf5, 0xef, 0x76, 0x49, 0xd5, 0x66,
	0xdd, 0x25, 0xc5, 0xfb, 0xfc, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x17, 0x1d, 0x81, 0xd6, 0x1b, 0x01,
	0x00, 0x00,
}
