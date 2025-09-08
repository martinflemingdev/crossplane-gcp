#!/usr/bin/python3

import yaml, sys, os


class RequiredFieldTracker:
    """Helper class to track required fields and add comments."""
    def __init__(self):
        self.required_fields = {}
    
    def add_required_fields(self, path: str, required_list: list):
        """Add required fields for a given path."""
        self.required_fields[path] = set(required_list) if required_list else set()
    
    def is_required(self, path: str, field: str) -> bool:
        """Check if a field is required at the given path."""
        return field in self.required_fields.get(path, set())


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


def process_schema_properties(properties: dict, required_tracker: RequiredFieldTracker, current_path: str = "", visited_refs: set = None) -> dict:
    """
    Recursively process schema properties to build nested structure with primitive values.
    """
    if visited_refs is None:
        visited_refs = set()
    
    result = {}
    
    for key, field_schema in properties.items():
        field_path = f"{current_path}.{key}" if current_path else key
        
        # Handle $ref references to avoid infinite recursion
        if '$ref' in field_schema:
            ref_path = field_schema['$ref']
            if ref_path in visited_refs:
                result[key] = "string"  # circular reference fallback
                continue
            visited_refs.add(ref_path)
        
        field_type = field_schema.get('type', 'string')
        
        # Track required fields for this level
        if 'required' in field_schema:
            required_tracker.add_required_fields(field_path, field_schema['required'])
        
        if field_type == 'object':
            # Handle object type
            if 'properties' in field_schema:
                # Track required fields for the object
                if 'required' in field_schema:
                    required_tracker.add_required_fields(field_path, field_schema['required'])
                result[key] = process_schema_properties(field_schema['properties'], required_tracker, field_path, visited_refs.copy())
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
                array_path = f"{field_path}[0]"
                if 'required' in items_schema:
                    required_tracker.add_required_fields(array_path, items_schema['required'])
                result[key] = [process_schema_properties(items_schema['properties'], required_tracker, array_path, visited_refs.copy())]
            else:
                # Array of primitives
                result[key] = [get_primitive_value(items_schema)]
        else:
            # Handle primitive types
            result[key] = get_primitive_value(field_schema)
    
    return result


def add_required_comments_to_yaml(yaml_content: str, required_tracker: RequiredFieldTracker) -> str:
    """
    Add '# required' comments to YAML content for required fields.
    """
    lines = yaml_content.split('\n')
    result_lines = []
    
    for line in lines:
        # Check if this line contains a field definition
        if ':' in line and not line.strip().startswith('#'):
            # Extract indentation and field name
            stripped = line.lstrip()
            indent = line[:len(line) - len(stripped)]
            
            if ':' in stripped:
                field_name = stripped.split(':')[0].strip()
                
                # Calculate the path based on indentation level
                indent_level = len(indent) // 2  # Assuming 2 spaces per indent
                
                # Build path from context (this is simplified - for complex nested structures
                # you might need a more sophisticated path tracking)
                if indent_level == 1 and field_name == 'forProvider':
                    line += '  # required'
                elif indent_level == 2:  # Fields under forProvider
                    if required_tracker.is_required('forProvider', field_name):
                        line += '  # required'
                # Add more path tracking as needed for deeper nesting
        
        result_lines.append(line)
    
    return '\n'.join(result_lines)


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
    
    # Check if spec is required
    required_tracker = RequiredFieldTracker()
    spec_schema = schema_root['properties']['spec']
    if 'required' in spec_schema:
        required_tracker.add_required_fields('spec', spec_schema['required'])
    
    for_provider_schema = spec_schema['properties']['forProvider']
    if 'required' in for_provider_schema:
        required_tracker.add_required_fields('forProvider', for_provider_schema['required'])
    
    if 'properties' in for_provider_schema:
        manifest['spec']['forProvider'] = process_schema_properties(for_provider_schema['properties'], required_tracker, 'forProvider')

    # Convert to YAML and add required comments
    yaml_content = yaml.dump(manifest, default_flow_style=False)
    return add_required_comments_to_yaml(yaml_content, required_tracker)


def write_to_file(filename: str, content: str):
    with open(filename, 'w') as file:
        file.write(content)


def run():
    if len(sys.argv) != 2:
        print("Please provider exactly one argument\nUsage: python3 create_stubbed_manifest_with_required.py <crd dir or file path>")
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
