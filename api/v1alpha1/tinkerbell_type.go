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
	Spec TinkerbellSpec `json:"spec,omitempty"`
}

// TinkerbellSpec specifies details of an Tinkerbell setup.
type TinkerbellSpec struct {
	// TinkerbellVersion is the Tinkerbell CRD version.
	TinkerbellVersion string `json:"tinkerbellVersion"`

	// ClusterDNS is the IP address of the cluster DNS resolver.
	ClusterDNS string `json:"clusterDNS"`

	// Registry is the registry to use for all images. If this field is set, all tink service deployment images
	// will be prefixed with this value. For example if the value here was set to docker.io, then boots image will be
	// docker.io/tinkerbell/boots.
	Registry string `json:"registry"`

	// ImagePullSecret the secret name containing the docker auth config which should exist in the same namespace where
	// the operator is deployed(typically tinkerbell)
	ImagePullSecret []string `json:"imagePullSecret"`

	// Boots contains all the information and spec about boots.
	Boots BootsSpec `json:"boots"`

	// Hegel contains all the information and spec about boots.
	Hegel HegelSpec `json:"hegel"`

	// Rufio contains all the information and spec about rufio.
	Rufio RufioSpec `json:"rufio"`

	// Tink contains all the information and spec about tink.
	Tink TinkSpec `json:"tink"`
}

// ImageSpec specifies the details of a tinkerbell services images.
type ImageSpec struct {
	// Repository is used to set the image repository for tinkerbell services.
	Repository string `json:"repository,omitempty"`

	// Tag is used to set the image tag for tinkerbell services.
	Tag string `json:"tag,omitempty"`
}

// BootsSpec specifies the deployment details of Tinkerbell service, Boots.
type BootsSpec struct {
	// Image specifies the details of a tinkerbell services images
	Image ImageSpec `json:"image"`

	// DHCPAddressListener set the ip and port to listen on for DHCP.
	DHCPAddressListener string `json:"dhcpAddress"`

	// TrustedProxies comma separated allowed CIDRs subnets to be used as trusted proxies
	TrustedProxies string `json:"trustedProxies"`

	// HTTPBind is the port to listen on for the serving iPXE binaries and files via HTTP.
	HTTPBind int `json:"httpBind"`

	// MirrorBaseURL is the URL for downloading an in-memory os such as Hook.
	MirrorBaseURL string `json:"mirrorBaseURL"`

	// OSIEPathOverride override the URL where OSIE/Hook images are located
	OSIEPathOverride string `json:"osiePathOverride"`

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

// HegelSpec specifies the details of tinkerbell service hegel.
type HegelSpec struct {
	// Image specifies the details of a tinkerbell services images
	Image ImageSpec `json:"image"`

	// TrustedProxies comma separated allowed CIDRs subnets to be used as trusted proxies
	TrustedProxies string `json:"trustedProxies"`
}

// RufioSpec specifies the details of tinkerbell service rufio.
type RufioSpec struct {
	// Image specifies the details of a tinkerbell services images
	Image ImageSpec `json:"image"`

	// TrustedProxies comma separated allowed CIDRs subnets to be used as trusted proxies
	TrustedProxies string `json:"trustedProxies"`
}

// TinkSpec specifies the details of tinkerbell service tink server.
type TinkSpec struct {
	// Image specifies the details of a tinkerbell services images
	Image ImageSpec `json:"image"`

	// TinkerbellTLS sets if the tink server should run with TLS or not.
	TinkerbellTLS bool `json:"tinkerbellTLS"`
}
