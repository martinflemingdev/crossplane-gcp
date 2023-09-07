#!/bin/bash

source "$(dirname "$0")/00_variables.sh"

yes | gcloud config set project ${PROJECT_ID}

gcloud services enable krmapihosting.googleapis.com \
    container.googleapis.com \
    cloudresourcemanager.googleapis.com \
    serviceusage.googleapis.com \
    anthos.googleapis.com 

# changing location to us-east1 to avoid resource issues
gcloud alpha anthos config controller create ${CC_NAME} \
    --location ${LOCATION} \
    --master-ipv4-cidr-block="172.16.10.0/28" \
    --experimental-features CRA
