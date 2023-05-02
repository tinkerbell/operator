# Tinkerbell Operator CRD


## Overview
Currently, the tinkerbell services are deployed by the operator using the manifest file (tinkerbell.yaml) found in the
deploy directory during deployment. However, our goal is to deploy the tinkerbell services only after a
CustomResourceDefinition (CRD) named Tinkerbell or TinkerbellInstance is created and to uninstall them
when the CRD is removed.

## Goals/Non-goals

**Goals**

- Define the CRD body and it's underlying specs
- How the CRD influences the operators' behaviour(setting up TB services and cleaning them up)
- Focus on Tinkerbell deployment specs and how they can be customized.

**Non-goals**

- Define a relationship between Tinkerbell core services or objects with this CRD
- Address advanced use cases such as migrations and upgrades as this is too early at this stage
- Extend the CRD spec to the `tink-stack` deployment as it is not yet clear how to handle it in the future
- Customize Kubernetes native fields such as deployment resources, image pull policies, etc...

## proposal

### Custom Resource Definitions

It is not yet clear about the name of the CRD at this point, however, I will use the name Tinkerbell but of course
open for a change.


#### Tinkerbell
```go
// Tinkerbell represents the tinkerbell stack that is being deployed in the kubernetes where the operator is deployed.
// Tinkerbell operator watches for different resources such as deployment, services, serviceAccounts, etc. One of those 
// CRs is Tinkerbell which the operator will install the tink-stack based on its specs. Once the CR is deleted, 
// the operator will delete all tinkerbell resources. 
type Tinkerbell struct {
    metav1.TypeMeta   `json:",inline"`
    metav1.ObjectMeta `json:"metadata,omitempty"`
	
    // Spec describes the desired tinkerbell stack state.
    Spec TinkerbellSpec `json:"spec,omitempty"`

    // Status contains information about the reconciliation status.
    Status TinkerbellStatus `json:"status,omitempty"`
}
```

#### TinkerbellSpec
```go
// TinkerbellSpec specifies details of an Tinkerbell setup.
type TinkerbellSpec struct {
    // TinkerbellVersion is the tinkerbell crd version.
    TinkerbellVersion string `json:"tinkerbellVersion"`
    
    // ClusterDNS is the ip address of the cluster dns resolver.
    ClusterDNS string `json:"clusterDNS"`
	
    // OverwriteRegistry is the registry to use for all images. If this field is set, all tink service deployment images
    // will be prefixed with this value. For example if the value here was set to docker.io, then boots image will be 
    // docker.io/tinkerbell/boots.
    OverwriteRegistry string `json:"overwriteRegistry"`
    
    // DockerPullConfig the secret name containing the docker auth config which should exist in the same namespace where 
    // the operator is deployed(typically tinkerbell)
    DockerPullComfig string `json:"dockerPullComfig"`
	
    // Boots contains all the information and spec about boots.
    Boots BootsSpec `json:"boots"`
    // Hegel contains all the information and spec about boots.
    Hegel HegelSpec `json:"hegel"`
    // Rufio contains all the information and spec about rufio.
    Rufio RufioSpec `json:"rufio"` 	
    // Tink contains all the information and spec about tink.
    Tink  TinkSpec  `json:"tink"` 
}
```

#### ImageSpec
```go
// ImageSpec specifies the details of a tinkerbell services images.
type ImageSpec struct {
    // ImageRepository is used to set the BootsSpec image repository.
    ImageRepository string `json:"imageRepository,omitempty"`
    
    // ImageTag is used to set the BootsSpec image tag.
    ImageTag string `json:"imageTag,omitempty"`
}
```

#### BootsSpec
```go
// BootsSpec specifies the details of tinkerbell service boots.
type BootsSpec struct {
    // Image specifies the details of a tinkerbell services images
    Image ImageSpec `json:"image"`
    
    // DHCPAddress set the ip and port to listen on for DHCP.
    DHCPAddress string `json:"dhcpAddress"`
	
    // TrustedProxies comma separated allowed CIDRs subnets to be used as trusted proxies
    TrustedProxies string `json:"trustedProxies"`

    // FacilityCode represents the facility in use.
    FacilityCode string `json:"facilityCode"`
	
    // HTTPBind is the port to listen on for the serving iPXE binaries and files via HTTP.
    HTTPBind int `json:"httpBind"`
	
    // MirrorBaseURL the URL from where the "OSIE" or Hook kernel(s) and initrd(s) will be downloaded by netboot clients
    MirrorBaseURL string `json:"mirrorBaseURL"`
	
    // OSIEPathOverride override the URL where OSIE/Hook images are located
    OSIEPathOverride string `json:"osiePathOverride"`    	
	
    // PublicIP is the IP that netboot clients and/or DHCP relay's will use to reach Boots
    PublicIP string `json:"publicIP"`
	
    // PublicSyslogFQDN is the IP that syslog clients will use to send messages
    PublicSyslogFQDN string `json:"publicSyslogFQDN"`
    
    // SyslogBind is the port that syslog clients will use to send messages
    SyslogBind int `json:"syslogBind"`
	
    // TinkerbellGRPCAuthority is the IP:Port that a Tink worker will use for communicated with the Tink server	
    TinkerbellGRPCAuthority	string `json:"tinkerbellGRPCAuthority"`
	
    // TinkerbellTLS sets if the boots should run with TLS or not.	
    TinkerbellTLS bool `json:"tinkerbellTLS"`

    // LogLevel sets the debug level for boots.
    logLevel string `json:"logLevel"`	
}
```

#### HegelSpec
```go
// HegelSpec specifies the details of tinkerbell service hegel.
type HegelSpec struct {
    // Image specifies the details of a tinkerbell services images
    Image ImageSpec `json:"image"`
    
    // TrustedProxies comma separated allowed CIDRs subnets to be used as trusted proxies
    TrustedProxies string `json:"trustedProxies"`
}
```

#### RufioSpec
```go
// RufioSpec specifies the details of tinkerbell service rufio.
type RufioSpec struct {
    // Image specifies the details of a tinkerbell services images
    Image ImageSpec `json:"image"`

    // TrustedProxies comma separated allowed CIDRs subnets to be used as trusted proxies
    TrustedProxies string `json:"trustedProxies"`
}
```

#### TinkSpec
```go
// TinkSpec specifies the details of tinkerbell service tink server.
type TinkSpec struct {
    // Image specifies the details of a tinkerbell services images
    Image ImageSpec `json:"image"`

    // TinkerbellTLS sets if the tink server should run with TLS or not.	
    TinkerbellTLS bool `json:"tinkerbellTLS"`
}
```
