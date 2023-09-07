#!/bin/bash

PATH_TO_YOUR_JSON_KEY_FILE='/home/martinfleming/src/crossplane/gcp/crossplane-gcp/CRA_KCC/axial-life-395119-e0b9a3b71fc0-crossplane.json'
YOUR_PROJECT_ID='axial-life-395119'

gcloud auth activate-service-account --key-file=${PATH_TO_YOUR_JSON_KEY_FILE}
gcloud config set project ${YOUR_PROJECT_ID}
