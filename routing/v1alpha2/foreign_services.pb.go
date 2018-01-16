// Code generated by protoc-gen-go. DO NOT EDIT.
// source: routing/v1alpha2/foreign_services.proto

package istio_routing_v1alpha2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Different ways of discovering the IP addresses associated with the
// service.
type Service_Discovery int32

const (
	// If set to "none", the proxy will assume that incoming connections
	// have already been resolved (to a specific destination IP
	// address). Such connections are typically routed via the proxy using
	// mechanisms such as IP table REDIRECT/ eBPF. After performing any
	// routing related transformations, the proxy will forward the
	// connection to the IP address to which the connection was bound.
	Service_NONE Service_Discovery = 0
	// If set to "static", the proxy will use specified endpoints (See
	// below) as the backing nodes associated with the foreign service.
	Service_STATIC Service_Discovery = 1
	// If set to "dns", the proxy will attempt to resolve the DNS
	// address during request processing. Use this mode if the set of
	// resolved addresses change dynamically. The "dns" mode is applicable
	// only when the hosts use exact DNS names without any wildcards.
	Service_DNS Service_Discovery = 2
)

var Service_Discovery_name = map[int32]string{
	0: "NONE",
	1: "STATIC",
	2: "DNS",
}
var Service_Discovery_value = map[string]int32{
	"NONE":   0,
	"STATIC": 1,
	"DNS":    2,
}

func (x Service_Discovery) String() string {
	return proto.EnumName(Service_Discovery_name, int32(x))
}
func (Service_Discovery) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{1, 0} }

// Service registry entries describe describe the endpoints, ports and
// protocols of a white-listed set of mesh-external domains and IP blocks
// that services in the mesh are allowed to access.
//
// NOTE 1: If a foreign service has the same name as a service in the
// service registry, the foreign service's declaration will be given
// precedence.
//
// NOTE 2: There can be ONLY ONE ForeignServices configuration for the
// entire mesh.
//
// For example, the following foreign services configuration describes the
// set of services at https://example.com. Eventhough the services behind
// example.com have to be accessed via HTTPS, in order for the application to
// obtain metrics from Istio, a plain text port (HTTP over port 80) is
// declared in addition to a secure port (HTTPS over 443). Connections
// arriving at port 443 on the sidecar will be treated as opaque TCP
// connections and will be forwarded as is to the destination, with limited
// visibility into the application flow. Connections arriving on port 80 on
// the sidecar will be able to take advantage of Istio's advanced routing
// and policy enforcement features. The associated routing rule ensures
// that outbound connections from the sidecar to the destination service
// happen over HTTPS.
//
//     apiVersion: config.istio.io/v1alpha2
//     kind: ForeignServices
//     metadata:
//       name: foreign-svc
//     spec:
//       services:
//       - hosts:
//         - example.com
//         ports:
//         - number: 443
//           name: https
//           protocol: HTTPS #treated as opaque TCP
//         - number: 80
//           name: http
//           protocol: HTTP
//         discovery: none
//
// And the associated route rule
//
//     apiVersion: config.istio.io/v1alpha2
//     kind: RouteRule
//     metadata:
//       name: my-foreign-rule
//     spec:
//       hosts:
//       - example.com
//       http:
//       - match:
//         - port:
//             name: http
//         route:
//         - destination:
//             name: example.com
//             port:
//               name: https
//
// Route rules can also be applied to services described in the
// ForeignServices resource. The following sample route rule rewrites
// /foocatalog to /barcatalog before forwarding the call to the intended
// destination.
//
//     apiVersion: config.istio.io/v1alpha2
//     kind: RouteRule
//     metadata:
//       name: foo-rule
//     spec:
//       hosts:
//       - example.com
//       http:
//       - match:
//         - uri:
//             prefix: /foocatalog
//         rewrite:
//           uri: /barcatalog
//
type ForeignServices struct {
	// REQUIRED: A list of server specifications.
	Services []*Service `protobuf:"bytes,1,rep,name=services" json:"services,omitempty"`
}

func (m *ForeignServices) Reset()                    { *m = ForeignServices{} }
func (m *ForeignServices) String() string            { return proto.CompactTextString(m) }
func (*ForeignServices) ProtoMessage()               {}
func (*ForeignServices) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *ForeignServices) GetServices() []*Service {
	if m != nil {
		return m.Services
	}
	return nil
}

// Service describes the endpoints, ports and protocols of the external
// service to be made accessible from within the mesh. For example,
//
//     apiVersion: config.istio.io/v1alpha2
//     kind: ForeignServices
//     metadata:
//       name: foreign-svc
//     spec:
//       services:
//       - hosts:
//         - *.foo.com
//         ports:
//         - number: 80
//           protocol: HTTP2
//           name: http2
//         resolution: none
//       - hosts:
//         - 192.192.33.33/16
//         ports:
//         - number: 27018
//           protocol: MONGO
//           name: mongo
//         resolution: none
//
type Service struct {
	// REQUIRED. The hosts associated with the external service. Could be a
	// DNS name with wildcard prefix or a CIDR prefix. Note that the hosts
	// field applies to all protocols.
	Hosts []string `protobuf:"bytes,1,rep,name=hosts" json:"hosts,omitempty"`
	// REQUIRED: The Ports associated with the external services.
	Ports []*Port `protobuf:"bytes,2,rep,name=ports" json:"ports,omitempty"`
	// Service discovery mode for the hosts.
	Resolution Service_Discovery `protobuf:"varint,3,opt,name=resolution,enum=istio.routing.v1alpha2.Service_Discovery" json:"resolution,omitempty"`
	// One or more endpoints associated with the service. Endpoints are valid
	// only when the discovery mode is set to "static".
	Endpoints []*Service_Endpoint `protobuf:"bytes,4,rep,name=endpoints" json:"endpoints,omitempty"`
}

func (m *Service) Reset()                    { *m = Service{} }
func (m *Service) String() string            { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()               {}
func (*Service) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *Service) GetHosts() []string {
	if m != nil {
		return m.Hosts
	}
	return nil
}

func (m *Service) GetPorts() []*Port {
	if m != nil {
		return m.Ports
	}
	return nil
}

func (m *Service) GetResolution() Service_Discovery {
	if m != nil {
		return m.Resolution
	}
	return Service_NONE
}

func (m *Service) GetEndpoints() []*Service_Endpoint {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

// Endpoint defines a network address (IP:port or hostname:port)
// associated with the foreign service.
type Service_Endpoint struct {
	// REQUIRED: Address associated with the network endpoint ( IP or fully
	// qualified domain name without wildcards).
	Address string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	// The port on which the endpoint is listening for network connections.
	Port *Port `protobuf:"bytes,2,opt,name=port" json:"port,omitempty"`
	// One or more labels associated with the endpoint.
	Labels map[string]string `protobuf:"bytes,3,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Service_Endpoint) Reset()                    { *m = Service_Endpoint{} }
func (m *Service_Endpoint) String() string            { return proto.CompactTextString(m) }
func (*Service_Endpoint) ProtoMessage()               {}
func (*Service_Endpoint) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1, 0} }

func (m *Service_Endpoint) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Service_Endpoint) GetPort() *Port {
	if m != nil {
		return m.Port
	}
	return nil
}

func (m *Service_Endpoint) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func init() {
	proto.RegisterType((*ForeignServices)(nil), "istio.routing.v1alpha2.ForeignServices")
	proto.RegisterType((*Service)(nil), "istio.routing.v1alpha2.Service")
	proto.RegisterType((*Service_Endpoint)(nil), "istio.routing.v1alpha2.Service.Endpoint")
	proto.RegisterEnum("istio.routing.v1alpha2.Service_Discovery", Service_Discovery_name, Service_Discovery_value)
}

func init() { proto.RegisterFile("routing/v1alpha2/foreign_services.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 363 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4d, 0x4b, 0xeb, 0x40,
	0x14, 0x86, 0x6f, 0x92, 0x7e, 0xe5, 0x14, 0xee, 0x0d, 0xc3, 0xe5, 0x12, 0xca, 0x45, 0x4b, 0x37,
	0x46, 0x17, 0xa9, 0x46, 0x17, 0x7e, 0xac, 0xc4, 0xb6, 0x50, 0x28, 0x51, 0xa6, 0xdd, 0x4b, 0xda,
	0x8e, 0xed, 0x60, 0xc8, 0x84, 0x99, 0x69, 0x24, 0xbf, 0xd6, 0x3f, 0xe1, 0x0f, 0x90, 0x4c, 0x26,
	0xb5, 0xa0, 0x52, 0x77, 0x39, 0x87, 0xf7, 0x79, 0x78, 0x4f, 0x12, 0x38, 0xe2, 0x6c, 0x23, 0x69,
	0xb2, 0xea, 0x67, 0x67, 0x51, 0x9c, 0xae, 0xa3, 0xa0, 0xff, 0xc4, 0x38, 0xa1, 0xab, 0xe4, 0x51,
	0x10, 0x9e, 0xd1, 0x05, 0x11, 0x7e, 0xca, 0x99, 0x64, 0xe8, 0x1f, 0x15, 0x92, 0x32, 0x5f, 0xc7,
	0xfd, 0x2a, 0xde, 0x39, 0xf8, 0x24, 0x58, 0x45, 0x92, 0xbc, 0x44, 0x79, 0xc9, 0xf5, 0x42, 0xf8,
	0x33, 0x2a, 0x8d, 0x53, 0x2d, 0x44, 0x37, 0xd0, 0xaa, 0xe4, 0xae, 0xd1, 0xb5, 0xbc, 0x76, 0x70,
	0xe8, 0x7f, 0x6d, 0xf7, 0x35, 0x83, 0xb7, 0x40, 0xef, 0xcd, 0x82, 0xa6, 0xde, 0xa2, 0xbf, 0x50,
	0x5f, 0x33, 0x21, 0x4b, 0x8b, 0x8d, 0xcb, 0x01, 0x05, 0x50, 0x4f, 0x19, 0x97, 0xc2, 0x35, 0x95,
	0xfb, 0xff, 0x77, 0xee, 0x07, 0xc6, 0x25, 0x2e, 0xa3, 0x68, 0x0c, 0xc0, 0x89, 0x60, 0xf1, 0x46,
	0x52, 0x96, 0xb8, 0x56, 0xd7, 0xf0, 0x7e, 0x07, 0xc7, 0x7b, 0x4a, 0xf9, 0x03, 0x2a, 0x16, 0x2c,
	0x23, 0x3c, 0xc7, 0x3b, 0x30, 0x1a, 0x81, 0x4d, 0x92, 0x65, 0xca, 0x68, 0x22, 0x85, 0x5b, 0x53,
	0x15, 0xbc, 0x7d, 0xa6, 0xa1, 0x06, 0xf0, 0x07, 0xda, 0x79, 0x35, 0xa0, 0x55, 0xed, 0x91, 0x0b,
	0xcd, 0x68, 0xb9, 0xe4, 0x44, 0x14, 0xb7, 0x1a, 0x9e, 0x8d, 0xab, 0x11, 0x9d, 0x42, 0xad, 0x38,
	0xc1, 0x35, 0xbb, 0xc6, 0xde, 0x63, 0x55, 0x12, 0x4d, 0xa0, 0x11, 0x47, 0x73, 0x12, 0x0b, 0xd7,
	0x52, 0xed, 0x2e, 0x7e, 0xda, 0xce, 0x9f, 0x28, 0x6c, 0x98, 0x48, 0x9e, 0x63, 0xed, 0xe8, 0x5c,
	0x41, 0x7b, 0x67, 0x8d, 0x1c, 0xb0, 0x9e, 0x49, 0xae, 0x4b, 0x16, 0x8f, 0xc5, 0x47, 0xca, 0xa2,
	0x78, 0x43, 0x54, 0x43, 0x1b, 0x97, 0xc3, 0xb5, 0x79, 0x69, 0xf4, 0x4e, 0xc0, 0xde, 0xbe, 0x42,
	0xd4, 0x82, 0x5a, 0x78, 0x1f, 0x0e, 0x9d, 0x5f, 0x08, 0xa0, 0x31, 0x9d, 0xdd, 0xce, 0xc6, 0x77,
	0x8e, 0x81, 0x9a, 0x60, 0x0d, 0xc2, 0xa9, 0x63, 0xce, 0x1b, 0xea, 0x6f, 0x3a, 0x7f, 0x0f, 0x00,
	0x00, 0xff, 0xff, 0x74, 0xe7, 0x07, 0x46, 0xb0, 0x02, 0x00, 0x00,
}