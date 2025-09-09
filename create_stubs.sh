#!/usr/bin/env bash
set -euo pipefail

if [ "$#" -ne 1 ]; then
  echo "Usage: $0 <api.group>  (e.g. secretmanager.gcp.upbound.io)"
  exit 2
fi

# Resolve script directory so paths are reliable regardless of cwd
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
GROUP="$1"
CRD_DIR="${SCRIPT_DIR}/Crossplane_GCP-Upbound/02_crds/family"
PY_SCRIPT="${SCRIPT_DIR}/Crossplane_GCP-Upbound/03_manifest-stubs/create_stubbed_manifest_with_required.py"

mkdir -p "$CRD_DIR"

# Dependencies
command -v kubectl >/dev/null 2>&1 || { echo "kubectl not found in PATH"; exit 3; }
command -v python3 >/dev/null 2>&1 || { echo "python3 not found in PATH"; exit 4; }
if [ ! -f "$PY_SCRIPT" ]; then
  echo "stub generator not found at $PY_SCRIPT"
  exit 5
fi

echo "Finding CRDs for API group: $GROUP"

mapfile -t CRDS < <(kubectl get crds -o jsonpath='{range .items[*]}{.metadata.name}{"\n"}{end}' | grep -F -- "$GROUP" || true)

if [ "${#CRDS[@]}" -eq 0 ]; then
  echo "No CRDs found for group '$GROUP'."
  exit 0
fi

echo "Found ${#CRDS[@]} CRD(s). Saving to $CRD_DIR"

for crd in "${CRDS[@]}"; do
  out="$CRD_DIR/${crd}.yaml"
  echo " - Saving CRD: $crd -> $out"
  kubectl get crd "$crd" -o yaml > "$out"
done

echo "Generating stubs using $PY_SCRIPT (stubs will be created in $CRD_DIR)"

pushd "$CRD_DIR" >/dev/null
for crd in "${CRDS[@]}"; do
  crdfile="${crd}.yaml"
  if [ -f "$crdfile" ]; then
    echo " - Generating stub for $crdfile"
    python3 "$PY_SCRIPT" "$crdfile"
  else
    echo " - Skipping missing file $crdfile"
  fi
done
popd >/dev/null

echo "Done. CRDs saved to: $CRD_DIR"
echo "Stub manifests (stubbed-*.yaml) created in: