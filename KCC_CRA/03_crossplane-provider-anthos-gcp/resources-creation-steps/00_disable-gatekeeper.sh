#!/bin/bash

# getting this error when trying to create a secret:

# error: failed to create secret admission webhook "validation.gatekeeper.sh" 
# denied the request: [forbidden-namespaces] username <crossplane@axial-life-395119.iam.gserviceaccount.com>
# is not allowed to access this resource </Secret/crossplane-system/gcp-secret>

############# To Disable Gatekeeper ##############

# Look for a webhook related to Gatekeeper, often named something like gatekeeper-validating-webhook-configuration.
kubectl get validatingwebhookconfigurations

# write backup config to file:
kubectl get validatingwebhookconfigurations gatekeeper-validating-webhook-configuration -o yaml > gatekeeper-validating-webhook-configuration-backup.yaml

# patch the webhook to disable it: (didn't work)
# kubectl patch validatingwebhookconfigurations gatekeeper-validating-webhook-configuration \
#     -p '{"webhooks":[{"name":"validation.gatekeeper.sh","clientConfig":{"service":{} }}]}'

# delete the config for now (can re-apply from backup)
kubectl delete validatingwebhookconfigurations gatekeeper-validating-webhook-configuration