#!/bin/bash

if [ -z "$1" ]
then
  echo "Please provide path to GCP service-account JSON key file"
  exit 1
fi

kubectl create secret \
generic gcp-secret \
-n crossplane-system \
--from-file=creds="$1"
