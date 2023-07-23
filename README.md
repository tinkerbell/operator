# Tinkerbell Operator

[![Experimental](https://img.shields.io/badge/Stability-Experimental-red.svg)](https://github.com/equinix-labs/standards#about-uniform-standards)

> This project is experimental and a work in progress. Use at your own risk!

Tinkerbell operator is a Kubernetes operator that deploys Tinkerbell components in a Kubernetes cluster. The operator takes care of 
the deployment and lifecycle of these Tinkerbell services: 

- **Boots**: The DHCP and iPXE server for Tinkerbell
- **Hegel**: An instance metadata service for Tinkerbell
- **Rufio**: Rufio is a declarative state manager for BMCs
- **Tink**: A workflow engine for provisioning bare metal

> **_NOTE:_** The operator presently deploys the tinkerbell services by utilizing the manifest file (tinkerbell.yaml) 
> located in the deploy directory upon deployment. However, our objective is to deploy the tinkerbell services only
> after the creation of a CustomResourceDefinition (CRD) named Tinkerbell or TinkerbellInstance and to remove them 
> when the CRD is deleted.

> **_NOTE:_** the operator is a tech preview project thus it shouldn't be used in production environments. 

## Motivation
The Tinkerbell organization offers the possibility of installing the Tinkerbell stack using Helm, which can be found in
this [repository](https://operatorcharts). This should be the simplest way to install the Tinkerbell service. However, we have identified different 
factors that have led us to start looking for a solution that introduces more observability and robustness to maintain 
different complex configurations and manage Tinkerbell service lifecycle.

Here are some of these factors:

- **Complex Configurations**: Different Tinkerbell services might need fine-tuned configurations (e.g., configurations that 
are derived from machine state). These can be very difficult to implement in a YAML format that Helm offers.
- **Upgrades and Migration**: While it is possible to upgrade Tinkerbell services using Helm, it is not possible to react to 
this upgrade, such as migrating the existing CRs to the new CRDs.
- **Day-2 Operations**: Taking care of various operations that are considered as Day-2 operations, such as maintaining and 
monitoring the running services and deployments.
- **Observability**: Observing and reflecting on the changes in the Tinkerbell setup ecosystem (e.g., picking up and deploying 
new configurations without the need for manual intervention).
- **Building up the Standards**: Identifying and focusing on Tinkerbell’s essential and standard components (e.g., Boots and Hegel) 
without confusing these components with utilities (e.g., KubeVip and Nginx) to introduce a clearer path of what we support 
out of the box as a community and what we do not.

## Installation
The operator can be installed on a Kubernetes cluster using `kubectl`:

```shell
kubectl apply -f ./deploy/tinkerbell.yaml
```

The tinkerbell.yaml file comprises all the necessary Kubernetes resources required by the operator, including the definition
of the tinkerbell namespace, tinkerbell CustomResourceDefinitions (CRDs), the tinkerbell service account, RBAC definitions, 
and the deployment definition of the operator. Once applied, it will create a new namespace named tinkerbell and deploy 
all the required resources there.

## Current Stage
The operator only deploys tinkerbell provisioning components, and it doesn't take care of any other utilities and network plumbings
(e.g: it doesn't install network services to expose boots). However, we are considering of adding some of these utilities in
the future as Addons.
