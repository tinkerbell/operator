# Kubetink

Kubetink is a Kubernetes operator that deploys Tinkerbell components in a Kubernetes cluster. Kubetink takes care of 
the deployment and lifecycle of these Tinkerbell services: 

- **Boots**: The DHCP and iPXE server for Tinkerbell
- **Hegel**: An instance metadata service for Tinkerbell
- **Rufio**: Rufio is a declarative state manager for BMCs
- **Tink**: A workflow engine for provisioning bare metal

> **_NOTE:_** Kubetink is a tech preview project thus it shouldn't be used in production environments. 

## Usage
Currently, the operator doesn't take care of deploying the needed k8s crds to start the controller, thus, the crds must be 
created first in the k8s cluster, then lunch the operator:

```shell
kubectl apply -f ./pkg/crd/tinkerbell.org
```

The operator will create a new namespace called `tinkerbell` abd deploys all the needed resources over there. 

## Current Stage
Kubetink only deploys tinkerbell provisioning components, and it doesn't take care of any other utilities and network plumbings
(e.g: it doesn't install network services to expose boots). However, we are considering of adding some of these utilities in
the future as Addons.
