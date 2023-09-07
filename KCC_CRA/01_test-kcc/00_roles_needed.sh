#!/bin/bash

######## No longer relevant, efforts to discover yakima service account, resolved by proper rolebindings #######


kubectl get configconnectorcontext.core.cnrm.cloud.google.com -n config-control -o yaml

# The instructions do not specify a service account at cluster creation. 
# Find what service account KCC is using...

k describe serviceaccounts -n cnrm-system cnrm-controller-manager-config-control

# The Annotations section displays which service-account KCC is using:

# Name:                cnrm-controller-manager-config-control
# Namespace:           cnrm-system
# Labels:              cnrm.cloud.google.com/scoped-namespace=config-control
#                      cnrm.cloud.google.com/system=true
#                      configconnectorcontext.cnrm.cloud.google.com/namespace=config-control
# Annotations:         cnrm.cloud.google.com/version: 1.108.0
#                      iam.gke.io/gcp-service-account: service-376994270622@gcp-sa-yakima.iam.gserviceaccount.com   <----------- service account
# Image pull secrets:  <none>
# Mountable secrets:   <none>
# Tokens:              <none>
# Events:              <none>

# That account will need:

# 1. "Owner" role or Service specific role for whatever services are being provisioned

# 2. Service Usage Consumer role (for fetching live state)
#     roles/serviceusage.serviceUsageConsumer role

######################## 
# Can't find the "yakima" role, so patching the serviceaccount to use another role

### Patching serviceaccount

# kubectl patch serviceaccount cnrm-controller-manager-config-control -n cnrm-system \
#   --type='strategic' \
#   -p='{"metadata": {"annotations": {"iam.gke.io/gcp-service-account": "service-376994270622@container-engine-robot.iam.gserviceaccount.com"}}}'

### Patching context

# kubectl patch configconnectorcontext.core.cnrm.cloud.google.com/configconnectorcontext.core.cnrm.cloud.google.com  -n config-control \
#   --type='strategic' \
#   -p='{"items": [{"spec": {"googleServiceAccount": "service-376994270622@container-engine-robot.iam.gserviceaccount.com"}}]}'

### Not working due to gatekeeper constraint
k describe k8sallowedresources.constraints.gatekeeper.sh/forbidden-namespaces

### Patching by re-applying with changed service account

kubectl get configconnectorcontext.core.cnrm.cloud.google.com -n config-control -o yaml

k apply -f CRA_KCC/01_test_kcc/cccontext.yaml