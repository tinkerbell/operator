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
