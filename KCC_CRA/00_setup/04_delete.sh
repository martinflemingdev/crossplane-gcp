#!/bin/bash 
source "$(dirname "$0")/00_variables.sh"

gcloud alpha anthos config controller delete ${CC_NAME} \
    --location ${LOCATION}
