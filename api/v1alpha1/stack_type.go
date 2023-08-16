package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:path=stack,scope=Namespaced,categories=stack,singular=stack,shortName=tb
// +kubebuilder:storageversion

// Stack represents the tinkerbell stack that is being deployed in the kubernetes where the operator is deployed.
// Tinkerbell operator watches for different resources such as deployment, services, serviceAccounts, etc. One of those
// CRs is Stack which the operator will install the tink-stack based on its specs. Once the CR is deleted,
// the operator will delete all tinkerbell resources.
type Stack struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec describes the desired tinkerbell stack state.
	Spec StackSpec `json:"spec"`
}

// StackSpec specifies details of an Tinkerbell setup.
type StackSpec struct {
	// Version is the Tinkerbell CRD version.
	Version string `json:"version"`

	// DNSResolverIP is indicative of the resolver IP utilized for setting up the nginx server responsible for proxying
	// to the Tinkerbell service and serving the Hook artifacts.
	DNSResolverIP string `json:"dnsResolverIP,omitempty"`

	// Registry is the registry to use for all images. If this field is set, all tink service deployment images
	// will be prefixed with this value. For example if the value here was set to docker.io, then boots image will be
	// docker.io/tinkerbell/boots.
	Registry string `json:"registry,omitempty"`

	// ImagePullSecret the secret name containing the docker auth config which should exist in the same namespace where
	// the operator is deployed(typically tinkerbell)
	ImagePullSecret []string `json:"imagePullSecret,omitempty"`

	// Services contains all Tinkerbell Stack services.
	Services Services `json:"services,omitempty"`
}

// Services contains all Tinkerbell Stack services.
type Services struct {
	// Boots contains all the information and spec about boots.
	Boots Boots `json:"boots,omitempty"`

	// Hegel contains all the information and spec about boots.
	Hegel Hegel `json:"hegel,omitempty"`

	// Rufio contains all the information and spec about rufio.
	Rufio Rufio `json:"rufio,omitempty"`

	// TinkServer contains all the information and spec about tink server.
	TinkServer TinkServer `json:"tinkServer,omitempty"`

	// TinkController contains all the information and spec about tink controller.
	TinkController TinkController `json:"tinkController,omitempty"`
}

// Boots specifies the deployment details of Tinkerbell service, Boots.
type Boots struct {
	// Image specifies the details of a tinkerbell services images
	Image Image `json:"image,omitempty"`

	// DHCPAddressListener set the ip and port to listen on for DHCP.
	DHCPAddressListener string `json:"dhcpAddressListener,omitempty"`

	// TrustedProxies comma separated allowed CIDRs subnets to be used as trusted proxies
	TrustedProxies []string `json:"trustedProxies,omitempty"`

	// HTTPAddress is the address to listen on for the serving iPXE binaries and files via HTTP.
	HTTPAddress string `json:"httpAddress,omitempty"`

	// OSIEURL override the URL where OSIE/Hook images are located
	OSIEURL string `json:"osieURL,omitempty"`

	// PublicIP is the IP that netboot clients and/or DHCP relay's will use to reach Boots
	PublicIP string `json:"publicIP,omitempty"`

	// PublicSyslogFQDN is the IP that syslog clients will use to send messages
	PublicSyslogFQDN string `json:"publicSyslogFQDN,omitempty"`

	// SyslogAddress is the IP and port that syslog clients will use to send messages
	SyslogAddress string `json:"syslogAddress,omitempty"`

	// TinkerbellGRPCAuthority IP and port to listen on for syslog messages.
	TinkerbellGRPCAuthority string `json:"tinkerbellGRPCAuthority,omitempty"`

	// TinkerbellTLS sets if the boots should run with TLS or not.
	TinkerbellTLS bool `json:"tinkerbellTLS,omitempty"`

	// LogLevel sets the debug level for boots.
	LogLevel string `json:"logLevel,omitempty"`
}

// Hegel specifies the details of tinkerbell service hegel.
type Hegel struct {
	// Image specifies the details of a tinkerbell services images
	Image Image `json:"image,omitempty"`

	// TrustedProxies comma separated allowed CIDRs subnets to be used as trusted proxies
	TrustedProxies []string `json:"trustedProxies,omitempty"`
}

// Rufio specifies the details of tinkerbell service rufio.
type Rufio struct {
	// Image specifies the details of a tinkerbell services images
	Image Image `json:"image,omitempty"`
}

// TinkServer specifies the details of tinkerbell service tink server.
type TinkServer struct {
	// Image specifies the details of a tinkerbell services images
	Image Image `json:"image,omitempty"`

	// TinkerbellTLS sets if the tink server should run with TLS or not.
	TinkerbellTLS bool `json:"tinkerbellTLS,omitempty"`
}

// TinkController specifies the details of tinkerbell service tink controller.
type TinkController struct {
	// Image specifies the details of a tinkerbell services images
	Image Image `json:"image,omitempty"`
}

// Image specifies the details of a tinkerbell services images.
type Image struct {
	// Repository is used to set the image repository for tinkerbell services.
	Repository string `json:"repository,omitempty"`

	// Tag is used to set the image tag for tinkerbell services.
	Tag string `json:"tag,omitempty"`
}
