import requests
import json

API_KEY = "API KEY HERE"
BASE_URL = "https://cloudbilling.googleapis.com/v1/services"

def fetch_all_services():
    """Fetch all GCP billing services with pagination"""
    all_services = []
    page_token = None
    
    while True:
        params = {
            "key": API_KEY,
            "pageSize": 5000  # Maximum page size
        }
        
        if page_token:
            params["pageToken"] = page_token
        
        response = requests.get(BASE_URL, params=params)
        response.raise_for_status()
        
        data = response.json()
        services = data.get("services", [])
        all_services.extend(services)
        
        print(f"Fetched {len(services)} services (Total: {len(all_services)})")
        
        page_token = data.get("nextPageToken")
        if not page_token:
            break
    
    return all_services

def create_markdown_table(services):
    """Create a markdown table from services list"""
    markdown = "# GCP Billing Services\n\n"
    markdown += f"**Total Services:** {len(services)}\n\n"
    markdown += "| Service ID | Display Name | Business Entity | Full Name |\n"
    markdown += "|------------|--------------|-----------------|------------|\n"
    
    for service in services:
        service_id = service.get("serviceId", "")
        display_name = service.get("displayName", "")
        business_entity = service.get("businessEntityName", "").replace("businessEntities/", "")
        full_name = service.get("name", "")
        
        # Escape pipe characters in names
        display_name = display_name.replace("|", "\\|")
        
        markdown += f"| {service_id} | {display_name} | {business_entity} | {full_name} |\n"
    
    return markdown

def main():
    print("Fetching all GCP billing services...")
    services = fetch_all_services()
    
    # Sort services alphabetically by display name (strip whitespace first)
    services.sort(key=lambda s: s.get("displayName", "").strip().lower())
    
    # Create markdown table
    markdown_output = create_markdown_table(services)
    
    # Save to file
    output_file = "/home/martinfleming/src/local-testing/GCP-ServiceIDs/services_list.md"
    with open(output_file, "w") as f:
        f.write(markdown_output)
    
    print(f"\n✅ Successfully saved {len(services)} services to {output_file}")
    
    # Also save raw JSON for reference
    json_file = "/home/martinfleming/src/local-testing/GCP-ServiceIDs/services_raw.json"
    with open(json_file, "w") as f:
        json.dump(services, f, indent=2)
    
    print(f"✅ Raw JSON saved to {json_file}")

if __name__ == "__main__":
    main()