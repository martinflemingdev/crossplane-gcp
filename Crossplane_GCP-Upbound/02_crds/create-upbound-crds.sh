#!/bin/bash

# Check if an argument is provided
if [ -z "$1" ]; then
  echo "Usage: $0 <output_directory>"
  exit 1
fi

# Assign the output directory from the first argument
output_dir="$1"

# Change to the specified directory
if ! cd "$output_dir"; then
  echo "Error: Could not change to directory $output_dir"
  exit 1
fi

# Get the CRDs filtered by 'upbound'
crds=$(kubectl get crds | grep upbound)

# Loop through each line in the result
while IFS= read -r line; do
  # Extract the first column (CRD name)
  crd_name=$(echo "$line" | awk '{print $1}')
  
  # Get crd and save the output to a YAML file
  kubectl get crd "$crd_name" -o yaml > "${crd_name}.yaml"
done <<< "$crds"
