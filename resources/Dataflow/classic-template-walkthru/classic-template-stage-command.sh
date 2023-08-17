#!/bin/bash

YOUR_PROJECT_ID=axial-life-395119
YOUR_BUCKET_NAME=dataflow-classic-template-example
YOUR_TEMPLATE_NAME=classic-template
YOUR_REGION=us-east-1

python3 classic-template.py \
  --runner DataflowRunner \
  --project $YOUR_PROJECT_ID \
  --staging_location gs://$YOUR_BUCKET_NAME/staging/ \
  --temp_location gs://$YOUR_BUCKET_NAME/temp/ \
  --template_location gs://$YOUR_BUCKET_NAME/templates/$YOUR_TEMPLATE_NAME \
  --inputFile gs://$YOUR_BUCKET_NAME/input.txt \
  --outputFile gs://$YOUR_BUCKET_NAME/output.txt \
  --region $YOUR_REGION
