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


def get_primitive_value(field_schema: dict) -> any:
    """
    Returns a primitive value based on the field's type and format.
    """
    field_type = field_schema.get('type', 'string')
    field_format = field_schema.get('format', '')
    
    if field_type == 'string':
        if field_format == 'date-time':
            return "2023-01-01T00:00:00Z"
        elif field_format == 'date':
            return "2023-01-01"
        elif field_format == 'time':
            return "00:00:00"
        else:
            return "string"
    elif field_type == 'integer':
        if field_format == 'int64':
            return 0
        else:
            return 0
    elif field_type == 'number':
        return 0.0
    elif field_type == 'boolean':
        return False
    else:
        return "string"  # fallback


def process_schema_properties(properties: dict, visited_refs: set = None) -> dict:
    """
    Recursively process schema properties to build nested structure with primitive values.
    """
    if visited_refs is None:
        visited_refs = set()
    
    result = {}
    
    for key, field_schema in properties.items():
        # Handle $ref references to avoid infinite recursion
        if '$ref' in field_schema:
            ref_path = field_schema['$ref']
            if ref_path in visited_refs:
                result[key] = "string"  # circular reference fallback
                continue
            visited_refs.add(ref_path)
        
        field_type = field_schema.get('type', 'string')
        
        if field_type == 'object':
            # Handle object type
            if 'properties' in field_schema:
                result[key] = process_schema_properties(field_schema['properties'], visited_refs.copy())
            elif 'additionalProperties' in field_schema:
                # Handle maps/dictionaries
                additional_props = field_schema['additionalProperties']
                if isinstance(additional_props, dict) and additional_props.get('type'):
                    result[key] = {
                        "key": get_primitive_value(additional_props)
                    }
                else:
                    result[key] = {"key": "string"}
            else:
                result[key] = {}
        elif field_type == 'array':
            # Handle array type
            items_schema = field_schema.get('items', {})
            items_type = items_schema.get('type', 'string')
            
            if items_type == 'object' and 'properties' in items_schema:
                # Array of objects
                result[key] = [process_schema_properties(items_schema['properties'], visited_refs.copy())]
            else:
                # Array of primitives
                result[key] = [get_primitive_value(items_schema)]
        else:
            # Handle primitive types
            result[key] = get_primitive_value(field_schema)
    
    return result


def stub_manifest(crd: dict) -> str:
    # Use newest version
    versions = []
    for v in crd['spec']['versions']:
        versions.append(v['name'])
    i = return_newest_version_index(versions)

    # Create manifest template
    manifest = {
        'apiVersion': f"{crd['spec']['group']}/{crd['spec']['versions'][i]['name']}",
        'kind': crd['spec']['names']['kind'],
        'metadata': {'name': f"stubbed-{crd['spec']['names']['singular']}.{crd['spec']['group']}"},
        'spec': {'forProvider': {}}
    }

    # Navigate to the forProvider fields
    schema_root = crd['spec']['versions'][i]['schema']['openAPIV3Schema']
    for_provider_schema = schema_root['properties']['spec']['properties']['forProvider']
    
    if 'properties' in for_provider_schema:
        manifest['spec']['forProvider'] = process_schema_properties(for_provider_schema['properties'])

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