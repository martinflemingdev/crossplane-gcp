$ ssh -i /home/martinfleming/src/crossplane/gcp/crossplane-gcp/gcp.pem ec2-user@184.73.120.8

# list pools
gcloud iam workload-identity-pools list --location="global"

# list providers
gcloud iam workload-identity-pools providers list --location="global" --workload-identity-pool="youtube-demo"

# lint-condition
gcloud alpha iam policies lint-condition --resource-type=project --resource=workloadidfederation --member="principalSet://iam.googleapis.com/projects/1086849180150/locations/global/workloadIdentityPools/youtube-demo/providers/youtube-demo-provider/attribute.aws_role/arn:aws:sts::011882408883:assumed-role/gcp-youtube-demo" --role=roles/storage.admin

# apply IAM policy member binding
gcloud projects add-iam-policy-binding workloadidfederation \
    --member="principalSet://iam.googleapis.com/projects/1086849180150/locations/global/workloadIdentityPools/youtube-demo/attribute.aws_role/arn:aws:sts::011882408883:assumed-role/gcp-youtube-demo" \
    --role=roles/storage.admin
