#!/usr/bin/bash

# CROSSPLANE APP
kubectl create namespace crossplane-system
helm repo add crossplane-stable https://charts.crossplane.io/stable
helm repo update

helm install crossplane --namespace crossplane-system crossplane-stable/crossplane
helm list -n crossplane-system
kubectl get all -n crossplane-system

# CROSSPLANE CLI
echo installing crossplane CLI to /usr/local/bin
curl -sL https://raw.githubusercontent.com/crossplane/crossplane/master/install.sh | sh
sudo mv kubectl-crossplane /usr/local/bin