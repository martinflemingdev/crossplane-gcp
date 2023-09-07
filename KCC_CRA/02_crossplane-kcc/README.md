# Create a Crossplane Composite Resource with KCC resource reference

To use the sample you need to have a Config Controller instance with the
crossplane feature enabled.

```
gcloud alpha anthos config controller create ${INSTANCE_NAME} \
  --location us-central1 \
  --experimental-features Crossplane --async
```

Also make sure you have followed Config Controller
[instructions](https://cloud.google.com/anthos-config-management/docs/tutorials/manage-resources-config-controller#grant-permission)
to grant Config Controller required permission in your GCP project.

This sample demonstrates how we can correctly specify a KCC resource reference
when using Crossplane composite to compose multiple KCC resources with
dependencies. PubSubSubscription is a dependent resource of PubSubTopic, and
in order to specify the correct name of the PubSubTopic resource for the
PubSubSubscription resource in the Composition, we use a status field
`topicRefName` to save the name of the PubSubTopic resource, then further
patch the dependent resource using the value from this status field.

Please follow the steps below in the order provided:

```
kubectl apply -f CompositeResourceDefinition.yaml
kubectl apply -f Composition_KCC.yaml
kubectl apply -f CompositeResource_KCC.yaml
```

You can use `kubectl get -f {FILENAME}` to check if the apply is successful.
If everything goes through, you will be able to find one KCC PubSubTopic
resource and one KCC PubSubSubscription resource created as a result of
`CompositionResource_KCC.yaml`. Example:

```
$ kubectl get pubsubtopics.pubsub.cnrm.cloud.google.com -n config-control
NAME                AGE    READY   STATUS     STATUS AGE
pubsubcmp01-pvrpn   4m3s   True    UpToDate   4m1s
$ kubectl get pubsubsubscriptions.pubsub.cnrm.cloud.google.com -n config-control
NAME                AGE     READY   STATUS     STATUS AGE
pubsubcmp01-ktcl2   4m10s   True    UpToDate   3m52s
```

