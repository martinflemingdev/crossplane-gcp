#!/bin/bash

source "$(dirname "$0")/00_variables.sh"

gcloud anthos config controller list
echo "Should say STATE:RUNNING"

gcloud anthos config controller get-credentials ${CC_NAME} --location=${LOCATION}
