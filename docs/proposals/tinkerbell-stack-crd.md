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