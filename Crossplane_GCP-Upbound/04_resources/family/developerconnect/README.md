# Developer Connect Dataform Crossplane manifests

This folder contains the Crossplane resources used to prove the Dataform + Developer Connect GitHub POC.

Apply everything from this directory:

```bash
kubectl apply -k .
```

Check status:

```bash
kubectl get connectconnections.developerconnect.gcp.upbound.io,connectgitrepositorylinks.developerconnect.gcp.upbound.io,repositories.dataform.gcp.upbound.io,projectiammembers.cloudplatform.gcp.upbound.io -l poc=dataform-developer-connect -o wide
```

The high-level setup notes live in the Dataform repo setup guide. This README explains what the manifests in this folder do.

## Manifest map

| File | Purpose |
| --- | --- |
| `00_ConnectConnection-github-app.yaml` | Creates the Crossplane-managed Developer Connect connection in platform project `axial-life-395119`, region `australia-southeast1`. |
| `01_ConnectGitRepositoryLink.yaml` | Creates the GitRepositoryLink child resource for `martinflemingdev/gcp-dataform-developer-connect` under the Crossplane-managed connection. |
| `02_ProjectIAMMember-dataform-service-agent.yaml` | Grants Dataform service agents permission on the platform project so they can use the central Developer Connect link. This includes both the same-project service agent and the outside `cdmc-data` service agent. |
| `03_DataformRepository-same-project.yaml` | Creates a Dataform repository in `axial-life-395119` that uses the Crossplane-managed GitRepositoryLink. |
| `04_DataformRepository-cross-project.yaml` | Creates a Dataform repository in `cdmc-data` that uses the GitRepositoryLink from platform project `axial-life-395119`. This proves cross-project consumption works through the API/Crossplane. |
| `05_ProjectIAMMember-crossplane-provider.yaml` | Grants the Crossplane provider identity permission to create/manage Dataform repositories in `cdmc-data`. |
| `06_DataformExecutionServiceAccount-and-IAM.yaml` | Creates the custom `dataform-e2e-runner` workflow identity, grants its BigQuery permissions, and lets the default Dataform service agent impersonate it under strict act-as mode. |
| `90_observe-console-created-connection-and-link.yaml` | Observes the console-created Developer Connect connection/link without managing updates or deletion. |
| `kustomization.yaml` | Applies the POC resources together. |

## Managed Developer Connect resources

`00_ConnectConnection-github-app.yaml` creates:

```text
projects/axial-life-395119/locations/australia-southeast1/connections/syd-data-res-crossplane
```

The connection reuses the GitHub App created during the console test:

```text
App ID:          5062914
Installation ID: 164497631
Host URI:        https://github.com
```

The private key and webhook secret are not stored in Kubernetes manifests. The connection points at the regional Secret Manager secret versions created by the console flow:

```text
projects/376994270622/locations/australia-southeast1/secrets/syd-data-res-ghe-private-key-e12264/versions/latest
projects/376994270622/locations/australia-southeast1/secrets/syd-data-res-ghe-webhook-secret-82755d/versions/latest
```

`01_ConnectGitRepositoryLink.yaml` creates:

```text
projects/axial-life-395119/locations/australia-southeast1/connections/syd-data-res-crossplane/gitRepositoryLinks/martinflemingdev-gcp-dataform-developer-connect-xp
```

## Dataform repository tests

The GA Dataform provider `dataform.gcp.upbound.io/v1beta1` exposes `gitRemoteSettings.gitRepositoryLink`, so Crossplane can create Dataform repositories that use Developer Connect machine credentials.

Same-project test:

```text
projects/axial-life-395119/locations/australia-southeast1/repositories/dataform-devconnect-same-project
```

Cross-project test:

```text
projects/cdmc-data/locations/australia-southeast1/repositories/dataform-devconnect-cross-project
```

Both repositories point at the central GitRepositoryLink in `axial-life-395119`:

```text
projects/axial-life-395119/locations/australia-southeast1/connections/syd-data-res-crossplane/gitRepositoryLinks/martinflemingdev-gcp-dataform-developer-connect-xp
```

This proves that the API accepts a full cross-project Developer Connect GitRepositoryLink even though the console picker did not expose the central connection/link cleanly.

## `02_ProjectIAMMember-dataform-service-agent.yaml`

This file is the service-agent grant for Dataform to use the central Developer Connect link.

The important detail is that `spec.forProvider.project` is the platform project:

```yaml
project: axial-life-395119
```

For a consuming project, the `member` is the consuming project's Dataform service agent:

```yaml
member: serviceAccount:service-<CONSUMING_PROJECT_NUMBER>@gcp-sa-dataform.iam.gserviceaccount.com
```

That means the POC cross-project grant is intentionally:

```yaml
project: axial-life-395119
member: serviceAccount:service-370318638050@gcp-sa-dataform.iam.gserviceaccount.com
```

Roles granted:

```text
roles/developerconnect.tokenAccessor
roles/developerconnect.gitProxyUser
```

Those roles let the `cdmc-data` Dataform service agent use the Developer Connect GitRepositoryLink that lives in `axial-life-395119`. They do not grant BigQuery execution access.

The same file also grants the same roles to the `axial-life-395119` Dataform service agent for the same-project test:

```text
service-376994270622@gcp-sa-dataform.iam.gserviceaccount.com
```

## `05_ProjectIAMMember-crossplane-provider.yaml`

This file captures the provider-side permission needed to create Dataform repositories in `cdmc-data`:

```text
serviceAccount:crossplane@axial-life-395119.iam.gserviceaccount.com
roles/dataform.admin
```

Without this, the cross-project repository failed with `dataform.repositories.create` denied in `cdmc-data`.

Keep this separate from `02_ProjectIAMMember-dataform-service-agent.yaml`:

- `02_ProjectIAMMember-dataform-service-agent.yaml` authorizes the Dataform service agent to use the platform Git link.
- `05_ProjectIAMMember-crossplane-provider.yaml` authorizes the Crossplane provider identity to create the Dataform repository in the consuming project.

## Dataform workflow execution IAM

`06_DataformExecutionServiceAccount-and-IAM.yaml` creates this custom execution identity in the consuming project:

```text
dataform-e2e-runner@cdmc-data.iam.gserviceaccount.com
```

The two `ProjectIAMMember` resources grant it `roles/bigquery.jobUser` and, for this POC, project-wide `roles/bigquery.dataEditor`. The latter should normally be replaced with dataset-scoped `DatasetIAMMember` grants in production.

The two `ServiceAccountIAMMember` resources grant `roles/iam.serviceAccountUser` and `roles/iam.serviceAccountTokenCreator` on the custom execution account to the default Dataform service agent:

```text
service-370318638050@gcp-sa-dataform.iam.gserviceaccount.com
```

Those bindings cannot be represented correctly as `ProjectIAMMember` resources because strict act-as authorization is evaluated on the custom service account resource. `04_DataformRepository-cross-project.yaml` selects this custom account through `spec.forProvider.serviceAccount`.

## Observed console resources

`90_observe-console-created-connection-and-link.yaml` observes these console-created resources without managing updates or deletion:

```text
projects/axial-life-395119/locations/australia-southeast1/connections/syd-data-res
projects/axial-life-395119/locations/australia-southeast1/connections/syd-data-res/gitRepositoryLinks/martinflemingdev-gcp-dataform-developer-connect
```

Use this path when testing Dataform against the connection created by the Google Cloud console.

## Current provider gap

The Developer Connect Crossplane CRD used here does not expose `gitProxyConfig`. Current `hashicorp/google` Terraform provider schema also does not expose a proxy field for `google_developer_connect_connection`, so this is not just an older Crossplane package issue.

The console-created connection has proxy enabled by default:

```json
"gitProxyConfig": {
  "enabled": true,
  "httpProxyBaseUri": "https://7pllpw2lcnfc5lw2h2zqux2un4aaaacxy2oa3hq-c-h-au-se1.developerconnect.dev"
}
```

The console-created GitRepositoryLink also exposes a proxy URL:

```json
"gitProxyUri": "https://australia-southeast1-git.developerconnect.dev/376994270622/syd-data-res/martinflemingdev-gcp-dataform-developer-connect"
```

The Crossplane-created connection did not include `gitProxyConfig`, and its GitRepositoryLink did not include `gitProxyUri`. Treat the Crossplane-created connection as proxy disabled unless a separate API/gcloud step enables it.

Workaround:

```bash
gcloud alpha developer-connect connections update syd-data-res-crossplane \
  --project=axial-life-395119 \
  --location=australia-southeast1 \
  --git-proxy-config-enabled
```

The older beta Dataform provider installed in this cluster exposes `gitRemoteSettings.url`, `defaultBranch`, and legacy HTTPS/SSH auth fields, but not `gitRepositoryLink`. Use the GA `dataform.gcp.upbound.io/v1beta1` provider from the Terraform v8.4.0 package for this POC.
