package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:path=stack,scope=Namespaced,categories=stack,singular=stack
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

// StackSpec specifies details of the Tinkerbell setup.
type StackSpec struct {
	// Version is the Tinkerbell CRD version.
	Version string `json:"version"`

	// Services contains all Tinkerbell Stack services.
	Services Services `json:"services"`

	// DNSResolverIP is indicative of the resolver IP utilized for setting up the nginx server responsible for proxying
	// to the Tinkerbell services and serving the Hook artifacts.
	// +optional
	DNSResolverIP *string `json:"dnsResolverIP,omitempty"`

	// Registry is the registry to use for all images. If this field is set, all tink service deployment images
	// will be prefixed with this value. For example if the value here was set to docker.io, then smee image will be
	// docker.io/tinkerbell/smee.
	// +optional
	Registry *string `json:"registry,omitempty"`

	// ImagePullSecrets the secret name containing the docker auth config which should exist in the same namespace where
	// the operator is deployed(typically tinkerbell)
	// +optional
	ImagePullSecrets []string `json:"imagePullSecrets,omitempty"`
}

// Services contains all Tinkerbell Stack services.
type Services struct {
	// Smee contains all the information and spec about smee.
	// +optional
	Smee *Smee `json:"smee,omitempty"`

	// Hegel contains all the information and spec about smee.
	// +optional
	Hegel *Hegel `json:"hegel,omitempty"`

	// Rufio contains all the information and spec about rufio.
	// +optional
	Rufio *Rufio `json:"rufio,omitempty"`

	// TinkServer contains all the information and spec about tink server.
	TinkServer TinkServer `json:"tinkServer"`

	// TinkController contains all the information and spec about tink controller.
	TinkController TinkController `json:"tinkController"`
}

// Smee specifies the deployment details of Tinkerbell service, Smee.
type Smee struct {
	// Image specifies the image repo and tag for Smee.
	Image Image `json:"image"`

	// BackendConfigs contains the configurations for smee backend.
	BackendConfigs BackendConfigs `json:"backendConfigs"`

	// SyslogConfigs contains the configurations of the syslog server.
	// +optional
	SyslogConfigs *SyslogConfigs `json:"syslogConfigs,omitempty"`

	// TFTPConfigs contains the configurations of Tinkerbell TFTP server.
	// +optional
	TFTPConfigs *TFTPConfigs `json:"tftpConfigs,omitempty"`

	// IPXEConfigs contains the iPXE configurations.
	// +optional
	IPXEConfigs *IPXEConfigs `json:"ipxeConfigs"`

	// DHCPConfigs contains the DHCP server configurations.
	// +optional
	DHCPConfigs *DHCPConfigs `json:"dhcpConfigs"`

	// LogLevel sets the debug level for smee.
	// +optional
	LogLevel *string `json:"logLevel,omitempty"`
}

// SyslogConfigs contains the configurations of the syslog server.
type SyslogConfigs struct {
	// IP is the local IP to listen on for syslog messages.
	IP string `json:"bindAddress"`

	// Port is the  local port to listen on for syslog messages.
	Port int `json:"port"`
}

// TFTPConfigs contains the configurations of Tinkerbell TFTP server.
type TFTPConfigs struct {
	// IP is the local IP to listen on to serve TFTP binaries.
	IP string `json:"ip"`

	// Port is the  local port to listen on to serve TFTP binaries.
	Port int `json:"port"`

	// TFTPTimeout specifies the iPXE tftp binary server requests timeout.
	// +optional
	TFTPTimeout *int `json:"tftpTimeout,omitempty"`

	// IPXEScriptPatch specifies the iPXE script fragment to patch into served iPXE binaries served via TFTP or HTTP.
	// +optional
	IPXEScriptPatch *string `json:"ipxeScriptPatch,omitempty"`
}

// IPXEConfigs contains the iPXE configurations.
type IPXEConfigs struct {
	// IP is the local IP to listen on to serve TFTP binaries.
	IP string `json:"ip"`

	// Port is the  local port to listen on to serve TFTP binaries.
	Port int `json:"port"`

	// TinkServerAddress specifies the IP:Port of the tink server.
	// +optional
	TinkServerAddress *string `json:"tinkServerAddress"`

	// EnableHTTPBinary enable iPXE HTTP binary server.
	// +optional
	EnableHTTPBinary *bool `json:"enableHTTPBinary,omitempty"`

	// EnableTLS sets if the smee should run with TLS or not.
	// +optional
	EnableTLS *bool `json:"enableTLS,omitempty"`

	// ExtraKernelArgs specifies extra set of kernel args (k=v k=v) that are appended to the kernel cmdline iPXE script.
	// +optional
	ExtraKernelArgs *string `json:"extraKernelArgs,omitempty"`

	// HookURL specifies the URL where OSIE(Hook) images are located.
	// +optional
	HookURL *string `json:"hookURL,omitempty"`

	// TrustedProxies comma separated allowed CIDRs subnets to be used as trusted proxies.
	// +optional
	TrustedProxies []string `json:"trustedProxies,omitempty"`
}

// DHCPConfigs contains the DHCP server configurations.
type DHCPConfigs struct {
	// IP is the local IP to listen on to serve TFTP binaries.
	IP string `json:"ip"`

	// Port is the  local port to listen on to serve TFTP binaries.
	Port int `json:"port"`

	// IPForPacket IP address to use in DHCP packets
	// +optional
	IPForPacket *string `json:"IPForPacket,omitempty"`

	// SyslogIP specifies the syslog server IP address to use in DHCP packets.
	// +optional
	SyslogIP *string `json:"syslogIP,omitempty"`

	// TFTPAddress specifies the tftp server address to use in DHCP packets.
	// +optional
	TFTPAddress *string `json:"tftpAddress,omitempty"`

	// HTTPIPXEBinaryAddress specifies the http ipxe binary server address (IP:Port) to use in DHCP packets.
	// +optional
	HTTPIPXEBinaryAddress *string `json:"httpIPXEBinaryAddress,omitempty"`

	// HTTPIPXEBinaryURI specifies the http ipxe script server URL to use in DHCP packets.
	// +optional
	HTTPIPXEScriptURI *string `json:"httpIPXEBinaryURI"`
}

// BackendConfigs contains the configurations for smee backend. The Backend has two modes, BackendKubeMode and BackendFileMode.
// Those modes are mutually exclusive with the BackendKubeMode being the default one. Users must choose one of the two modes.
type BackendConfigs struct {
	// BackendKubeMode contains the Kubernetes backend configurations for DHCP and the HTTP iPXE script.
	// +optional
	BackendKubeMode *BackendKubeMode `json:"backendKubeMode,omitempty"`

	// BackendFileMode contains the file backend configurations for DHCP and the HTTP iPXE script.
	// +optional
	BackendFileMode *BackendFileMode `json:"backendFileMode,omitempty"`
}

// BackendKubeMode contains the Kubernetes backend configurations for DHCP and the HTTP iPXE script.
type BackendKubeMode struct {
	// ConfigFilePath specifies the Kubernetes config file location.
	// +optional
	KubeConfigFilePath *string `json:"configFilePath,omitempty"`

	// KubeAPIURL specifies the Kubernetes API URL, used for in-cluster client construction.
	// +optional
	KubeAPIURL *string `json:"kubeAPIURL,omitempty"`

	// KubeNamespace specifies an optional Kubernetes namespace override to query hardware data from.
	// +optional
	KubeNamespace *string `json:"kubeNamespace,omitempty"`
}

// BackendFileMode contains the file backend configurations for DHCP and the HTTP iPXE script
type BackendFileMode struct {
	// FilePath specifies the hardware yaml file path for the file backend.
	FilePath string `json:"filePath"`
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

	// EnableTLS sets if the tink server should run with TLS or not.
	EnableTLS bool `json:"enableTLS,omitempty"`
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
