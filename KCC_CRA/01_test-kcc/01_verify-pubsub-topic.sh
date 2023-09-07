#!/bin/bash

gcloud pubsub topics list
gcloud pubsub topics list --format="table(name, labels, kmsKeyName)"
gcloud pubsub topics list --format="json"

# pubsubtopic.pubsub.cnrm.cloud.google.com/kcc-example-topic