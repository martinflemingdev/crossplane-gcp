// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package datastream

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func extractError(d *schema.ResourceData) error {
	// Casts are not safe since the logic that populate it is type deterministic.
	error := d.Get("error").([]interface{})[0].(map[string]interface{})
	message := error["message"].(string)
	details := error["details"].(map[string]interface{})
	detailsJSON, _ := json.Marshal(details)
	return fmt.Errorf("Failed to create PrivateConnection. %s details = %s", message, string(detailsJSON))
}

// waitForPrivateConnectionReady waits for a private connection state to become
// CREATED, if the state is FAILED propegate the error to the user.
func waitForPrivateConnectionReady(d *schema.ResourceData, config *transport_tpg.Config, timeout time.Duration) error {
	return resource.Retry(timeout, func() *resource.RetryError {
		if err := resourceDatastreamPrivateConnectionRead(d, config); err != nil {
			return resource.NonRetryableError(err)
		}

		name := d.Get("name").(string)
		state := d.Get("state").(string)
		if state == "CREATING" {
			return resource.RetryableError(fmt.Errorf("PrivateConnection %q has state %q.", name, state))
		} else if state == "CREATED" {
			log.Printf("[DEBUG] PrivateConnection %q has state %q.", name, state)
			return nil
		} else if state == "FAILED" {
			return resource.NonRetryableError(extractError(d))
		} else {
			return resource.NonRetryableError(fmt.Errorf("PrivateConnection %q has state %q.", name, state))
		}
	})
}

func ResourceDatastreamPrivateConnection() *schema.Resource {
	return &schema.Resource{
		Create: resourceDatastreamPrivateConnectionCreate,
		Read:   resourceDatastreamPrivateConnectionRead,
		Delete: resourceDatastreamPrivateConnectionDelete,

		Importer: &schema.ResourceImporter{
			State: resourceDatastreamPrivateConnectionImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Display name.`,
			},
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The name of the location this private connection is located in.`,
			},
			"private_connection_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The private connectivity identifier.`,
			},
			"vpc_peering_config": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Description: `The VPC Peering configuration is used to create VPC peering
between Datastream and the consumer's VPC.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"subnet": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: `A free subnet for peering. (CIDR of /29)`,
						},
						"vpc": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
							Description: `Fully qualified name of the VPC that Datastream will peer to.
Format: projects/{project}/global/{networks}/{name}`,
						},
					},
				},
			},
			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: `Labels.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"error": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `The PrivateConnection error in case of failure.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"details": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: `A list of messages that carry the error details.`,
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"message": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `A message containing more information about the error that occurred.`,
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The resource's name.`,
			},
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `State of the PrivateConnection.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceDatastreamPrivateConnectionCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	labelsProp, err := expandDatastreamPrivateConnectionLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	displayNameProp, err := expandDatastreamPrivateConnectionDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	vpcPeeringConfigProp, err := expandDatastreamPrivateConnectionVpcPeeringConfig(d.Get("vpc_peering_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("vpc_peering_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(vpcPeeringConfigProp)) && (ok || !reflect.DeepEqual(v, vpcPeeringConfigProp)) {
		obj["vpcPeeringConfig"] = vpcPeeringConfigProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DatastreamBasePath}}projects/{{project}}/locations/{{location}}/privateConnections?privateConnectionId={{private_connection_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new PrivateConnection: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for PrivateConnection: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
	})
	if err != nil {
		return fmt.Errorf("Error creating PrivateConnection: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/privateConnections/{{private_connection_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = DatastreamOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating PrivateConnection", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create PrivateConnection: %s", err)
	}

	if err := d.Set("name", flattenDatastreamPrivateConnectionName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/privateConnections/{{private_connection_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	if err := waitForPrivateConnectionReady(d, config, d.Timeout(schema.TimeoutCreate)-time.Minute); err != nil {
		return fmt.Errorf("Error waiting for PrivateConnection %q to be CREATED. %q", d.Get("name").(string), err)
	}

	log.Printf("[DEBUG] Finished creating PrivateConnection %q: %#v", d.Id(), res)

	return resourceDatastreamPrivateConnectionRead(d, meta)
}

func resourceDatastreamPrivateConnectionRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DatastreamBasePath}}projects/{{project}}/locations/{{location}}/privateConnections/{{private_connection_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for PrivateConnection: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("DatastreamPrivateConnection %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading PrivateConnection: %s", err)
	}

	if err := d.Set("name", flattenDatastreamPrivateConnectionName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading PrivateConnection: %s", err)
	}
	if err := d.Set("labels", flattenDatastreamPrivateConnectionLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading PrivateConnection: %s", err)
	}
	if err := d.Set("display_name", flattenDatastreamPrivateConnectionDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading PrivateConnection: %s", err)
	}
	if err := d.Set("state", flattenDatastreamPrivateConnectionState(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading PrivateConnection: %s", err)
	}
	if err := d.Set("error", flattenDatastreamPrivateConnectionError(res["error"], d, config)); err != nil {
		return fmt.Errorf("Error reading PrivateConnection: %s", err)
	}
	if err := d.Set("vpc_peering_config", flattenDatastreamPrivateConnectionVpcPeeringConfig(res["vpcPeeringConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading PrivateConnection: %s", err)
	}

	return nil
}

func resourceDatastreamPrivateConnectionDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for PrivateConnection: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{DatastreamBasePath}}projects/{{project}}/locations/{{location}}/privateConnections/{{private_connection_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting PrivateConnection %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "PrivateConnection")
	}

	err = DatastreamOperationWaitTime(
		config, res, project, "Deleting PrivateConnection", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting PrivateConnection %q: %#v", d.Id(), res)
	return nil
}

func resourceDatastreamPrivateConnectionImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/privateConnections/(?P<private_connection_id>[^/]+)",
		"(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<private_connection_id>[^/]+)",
		"(?P<location>[^/]+)/(?P<private_connection_id>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/privateConnections/{{private_connection_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	if err := waitForPrivateConnectionReady(d, config, d.Timeout(schema.TimeoutCreate)-time.Minute); err != nil {
		return nil, fmt.Errorf("Error waiting for PrivateConnection %q to be CREATED during importing: %q", d.Get("name").(string), err)
	}

	return []*schema.ResourceData{d}, nil
}

func flattenDatastreamPrivateConnectionName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDatastreamPrivateConnectionLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDatastreamPrivateConnectionDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDatastreamPrivateConnectionState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDatastreamPrivateConnectionError(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["message"] =
		flattenDatastreamPrivateConnectionErrorMessage(original["message"], d, config)
	transformed["details"] =
		flattenDatastreamPrivateConnectionErrorDetails(original["details"], d, config)
	return []interface{}{transformed}
}
func flattenDatastreamPrivateConnectionErrorMessage(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDatastreamPrivateConnectionErrorDetails(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDatastreamPrivateConnectionVpcPeeringConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["vpc"] =
		flattenDatastreamPrivateConnectionVpcPeeringConfigVpc(original["vpc"], d, config)
	transformed["subnet"] =
		flattenDatastreamPrivateConnectionVpcPeeringConfigSubnet(original["subnet"], d, config)
	return []interface{}{transformed}
}
func flattenDatastreamPrivateConnectionVpcPeeringConfigVpc(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDatastreamPrivateConnectionVpcPeeringConfigSubnet(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandDatastreamPrivateConnectionLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandDatastreamPrivateConnectionDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDatastreamPrivateConnectionVpcPeeringConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedVpc, err := expandDatastreamPrivateConnectionVpcPeeringConfigVpc(original["vpc"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedVpc); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["vpc"] = transformedVpc
	}

	transformedSubnet, err := expandDatastreamPrivateConnectionVpcPeeringConfigSubnet(original["subnet"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSubnet); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["subnet"] = transformedSubnet
	}

	return transformed, nil
}

func expandDatastreamPrivateConnectionVpcPeeringConfigVpc(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDatastreamPrivateConnectionVpcPeeringConfigSubnet(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
