#!/bin/bash

source "$(dirname "$0")/00_variables.sh"

export SA_EMAIL="$(kubectl get ConfigConnectorContext -n config-control \
    -o jsonpath='{.items[0].spec.googleServiceAccount}' 2> /dev/null)"

gcloud projects add-iam-policy-binding ${PROJECT_ID} \
    --member "serviceAccount:${SA_EMAIL}" \
    --role "roles/owner" \
    --project ${PROJECT_ID}
