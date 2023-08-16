# set CLI to use service account credentials
gcloud auth activate-service-account --key-file=PATH_TO_YOUR_KEY_FILE.json
gcloud config set project PROJECT_ID

gcloud auth application-default login

# storage
gcloud storage buckets list
gcloud storage buckets describe gs://crossplane-bucket

gsutil cp [LOCAL_FILE_PATH] gs://[BUCKET_NAME]/[OPTIONAL_DESTINATION_PATH]
gsutil cp -r [LOCAL_DIRECTORY_PATH] gs://[BUCKET_NAME]/[OPTIONAL_DESTINATION_PATH]

# bigquery
gcloud bigquery tables list --dataset=<DATASET_NAME> --project=<PROJECT_ID>
gcloud bigquery tables list --dataset=exampledataset --project=axial-life-395119

