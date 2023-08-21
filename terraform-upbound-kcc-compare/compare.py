import os

def list_files_recursive(folder_path):
    """List all files recursively in the given folder."""
    all_files = []

    for root, dirs, files in os.walk(folder_path):
        for file in files:
            all_files.append(os.path.join(root, file))

    return all_files

def compare_lists(listA, listB):
    """
    Compare two lists of strings and return three lists:
    1. Strings that are present in both lists.
    2. Strings that are in listA but not in listB.
    3. Strings that are in listB but not in listA.
    """
    matches = [item for item in listA if item in listB]
    only_in_listA = [item for item in listA if item not in listB]
    only_in_listB = [item for item in listB if item not in listA]

    return matches, only_in_listA, only_in_listB

def kcc_remove_service(s): 
    parts = s.split('_')
    parts[-1] = parts[-1].replace(parts[0], '', 1)
    return '_'.join(parts)

# Paths
terraform_path="terraform-provider-google-google-services"
upbound_path="provider-gcp-package-crds"
kcc_path="k8s-config-connector-crds"

# Files
terraform_files=list_files_recursive(terraform_path)
upbound_files=list_files_recursive(upbound_path)
kcc_files=list_files_recursive(kcc_path)

# Objects 
tf_resources =  [s.split('/')[-1] for s in terraform_files if "resource" in s and "test" not in s and "sweeper" not in s]
up_resources = [s.split('/')[-1] for s in upbound_files if "iammember" not in s and 'iampolicy' not in s and 'iambinding' not in s]
kcc_resources = [s.split('/')[-1] for s in kcc_files]

print(f"tf_resources: {len(tf_resources)}")
print(f"up_resources: {len(up_resources)}")
print(f"kcc_resources: {len(kcc_resources)}")

# Transforms

tf_transformed = [s.replace('_','').replace('.go','').replace('resource','') for s in tf_resources]

up_transformed = [s.replace('.gcp.upbound.io_','').replace('s.yaml','').replace('.yaml','') for s in up_resources]
up_transformed = [s.replace('policie','policy').replace('accesse','access').replace('indice','index').replace('entrie','entry').replace('repositorie','repository').replace('authoritie','authority').replace('addresse','address').replace('registrie','registry').replace('proxie','proxy') for s in up_transformed]

kcc_transformed = [s.replace('_v1alpha1','').replace('_v1beta1','').replace('.yaml','') for s in kcc_resources]
kcc_transformed = [kcc_remove_service(s).replace('_','') for s in kcc_transformed]


# matches, only_in_tf_transformed, only_in_up_transformed = compare_lists(tf_transformed, up_transformed)
matches, only_in_tf_transformed, only_in_kcc_transformed = compare_lists(tf_transformed, kcc_transformed)


print(f"matches: {len(matches)}")
print(f"only_in_tf_transformed: {len(only_in_tf_transformed)}")
# print(f"only_in_up_transformed: {len(only_in_up_transformed)}")
print(f"only_in_kcc_transformed: {len(only_in_kcc_transformed)}")


# with open('only_in_up_transformed.txt','w') as file:
#     for i in sorted(only_in_up_transformed):
#         file.write(f"{i}\n")

with open('only_in_kcc_transformed.txt','w') as file:
    for i in sorted(only_in_kcc_transformed):
        if i != '':
            file.write(f"{i}\n")

with open('only_in_tf_transformed.txt','w') as file:
    for i in sorted(only_in_tf_transformed):
        file.write(f"{i}\n")

with open('matches.txt','w') as file:
    for i in sorted(matches):
        file.write(f"{i}\n")

with open('outcome-summary.txt','w') as file:
    file.write(f"matches: {len(matches)}\n")
    file.write(f"only_in_tf_transformed: {len(only_in_tf_transformed)}\n")
    # file.write(f"only_in_up_transformed: {len(only_in_up_transformed)}")
    file.write(f"only_in_up_transformed: {len(only_in_kcc_transformed)}")