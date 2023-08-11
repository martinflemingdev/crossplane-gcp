#!/usr/bin/python3

import yaml, sys, os


def load_yaml_to_dict(file_path: str) -> dict:
    with open(file_path, 'r') as file:
        data = yaml.safe_load(file)
    return data


def return_newest_version_index(versions: list) -> int:
    latest_version_index = 0
    version_check = max(versions)
    for i in range(len(versions)):
        if len(versions[i]) == 2 and versions[i][1] >= version_check[1]:
            version_check = versions[i]
            latest_version_index = i
        elif versions[i] == version_check:
           latest_version_index = i
    return latest_version_index


def stub_manifest(crd: dict) -> str:
    # Use newest version
    i = return_newest_version_index(crd['spec']['versions'])

    # Create manifest template
    manifest = {
        'apiVersion': f"{crd['spec']['group']}/{crd['spec']['versions'][i]['name']}",
        'kind': crd['spec']['names']['kind'],
        'metadata': {'name': f"stubbed-{crd['spec']['names']['singular']}.{crd['spec']['group']}"},
        'spec': {'forProvider': {}}
    }

    # Navigate to the forProvider fields
    for_provider_fields = crd['spec']['versions'][i]['schema']['openAPIV3Schema']['properties']['spec']['properties']['forProvider']['properties']

    # Set each field to None
    for key in for_provider_fields:
        manifest['spec']['forProvider'][key] = None

    return yaml.dump(manifest, default_flow_style=False)


def write_to_file(filename: str, content: str):
    with open(filename, 'w') as file:
        file.write(content)


def run():
    if len(sys.argv) != 2:
        print("Please provider exactly one argument\nUsage: python3 create_stubbed_manifest.py <crd dir or file path>")
        sys.exit(1)
    path = sys.argv[1]

    try:
        # if dir
        files = os.listdir(path)
        for file in files:
            filename="stubbed-" + file
            write_to_file(filename,stub_manifest(load_yaml_to_dict(f"{path}/{file}")))
    except:
        # if file
        filename="stubbed-" + path.split('/')[-1]
        write_to_file(filename,stub_manifest(load_yaml_to_dict(path)))

if __name__ == '__main__':
    
    # test
    # print(stub_manifest(load_yaml_to_dict('/home/martinfleming/src/crossplane/gcp/crossplane-gcp/crds/jobs.dataflow.gcp.upbound.io.yaml')))

    run()