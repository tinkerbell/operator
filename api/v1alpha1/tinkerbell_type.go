package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:path=tinkerbell,scope=Namespaced,categories=tinkerbell,singular=tinkerbell,shortName=tb
// +kubebuilder:storageversion

// Tinkerbell represents the tinkerbell stack that is being deployed in the kubernetes where the operator is deployed.
// Tinkerbell operator watches for different resources such as deployment, services, serviceAccounts, etc. One of those
// CRs is Tinkerbell which the operator will install the tink-stack based on its specs. Once the CR is deleted,
// the operator will delete all tinkerbell resources.
type Tinkerbell struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec describes the desired tinkerbell stack state.
	Spec TinkerbellStackSpec `json:"spec,omitempty"`
}

// TinkerbellStackSpec specifies details of an Tinkerbell setup.
type TinkerbellStackSpec struct {
	// Version is the Tinkerbell CRD version.
	Version string `json:"version"`

	// DNSResolverIP is indicative of the resolver IP utilized for setting up the nginx server responsible for proxying
	// to the Tinkerbell service and serving the Hook artifacts.
	DNSResolverIP string `json:"dnsResolverIP"`

	// Registry is the registry to use for all images. If this field is set, all tink service deployment images
	// will be prefixed with this value. For example if the value here was set to docker.io, then boots image will be
	// docker.io/tinkerbell/boots.
	Registry string `json:"registry"`

	// ImagePullSecret the secret name containing the docker auth config which should exist in the same namespace where
	// the operator is deployed(typically tinkerbell)
	ImagePullSecret []string `json:"imagePullSecret"`

	// Services contains all Tinkerbell Stack services.
	Services Services `json:"services"`
}

// Services contains all Tinkerbell Stack services.
type Services struct {
	// Boots contains all the information and spec about boots.
	Boots Boots `json:"boots"`

	// Hegel contains all the information and spec about boots.
	Hegel Hegel `json:"hegel"`

	// Rufio contains all the information and spec about rufio.
	Rufio Rufio `json:"rufio"`

	// TinkServer contains all the information and spec about tink server.
	TinkServer TinkServer `json:"tinkServer"`

	// TinkController contains all the information and spec about tink controller.
	TinkController Image `json:"tinkController"`
}

// Boots specifies the deployment details of Tinkerbell service, Boots.
type Boots struct {
	// Image specifies the details of a tinkerbell services images
	Image Image `json:"image"`

	// DHCPAddressListener set the ip and port to listen on for DHCP.
	DHCPAddressListener string `json:"dhcpAddress"`

	// TrustedProxies comma separated allowed CIDRs subnets to be used as trusted proxies
	TrustedProxies string `json:"trustedProxies"`

	// HTTPAddress is the address to listen on for the serving iPXE binaries and files via HTTP.
	HTTPAddress string `json:"httpAddress"`

	// OSIEURL override the URL where OSIE/Hook images are located
	OSIEURL string `json:"osieURL"`

	// PublicIP is the IP that netboot clients and/or DHCP relay's will use to reach Boots
	PublicIP string `json:"publicIP"`

	// PublicSyslogFQDN is the IP that syslog clients will use to send messages
	PublicSyslogFQDN string `json:"publicSyslogFQDN"`

	// SyslogAddress is the IP and port that syslog clients will use to send messages
	SyslogAddress string `json:"syslogAddress"`

	// TinkerbellGRPCAuthority IP and port to listen on for syslog messages.
	TinkerbellGRPCAuthority string `json:"tinkerbellGRPCAuthority"`

	// TinkerbellTLS sets if the boots should run with TLS or not.
	TinkerbellTLS bool `json:"tinkerbellTLS"`

	// LogLevel sets the debug level for boots.
	LogLevel string `json:"logLevel"`
}

// Hegel specifies the details of tinkerbell service hegel.
type Hegel struct {
	// Image specifies the details of a tinkerbell services images
	Image Image `json:"image"`

	// TrustedProxies comma separated allowed CIDRs subnets to be used as trusted proxies
	TrustedProxies string `json:"trustedProxies"`
}

// Rufio specifies the details of tinkerbell service rufio.
type Rufio struct {
	// Image specifies the details of a tinkerbell services images
	Image Image `json:"image"`

	// TrustedProxies comma separated allowed CIDRs subnets to be used as trusted proxies
	TrustedProxies string `json:"trustedProxies"`
}

// TinkServer specifies the details of tinkerbell service tink server.
type TinkServer struct {
	// Image specifies the details of a tinkerbell services images
	Image Image `json:"image"`

	// TinkerbellTLS sets if the tink server should run with TLS or not.
	TinkerbellTLS bool `json:"tinkerbellTLS"`
}

// Image specifies the details of a tinkerbell services images.
type Image struct {
	// Repository is used to set the image repository for tinkerbell services.
	Repository string `json:"repository,omitempty"`

	// Tag is used to set the image tag for tinkerbell services.
	Tag string `json:"tag,omitempty"`
}
