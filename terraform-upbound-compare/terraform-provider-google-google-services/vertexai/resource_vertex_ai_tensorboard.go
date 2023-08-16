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

package vertexai

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func ResourceVertexAITensorboard() *schema.Resource {
	return &schema.Resource{
		Create: resourceVertexAITensorboardCreate,
		Read:   resourceVertexAITensorboardRead,
		Update: resourceVertexAITensorboardUpdate,
		Delete: resourceVertexAITensorboardDelete,

		Importer: &schema.ResourceImporter{
			State: resourceVertexAITensorboardImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `User provided name of this Tensorboard.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Description of this Tensorboard.`,
			},
			"encryption_spec": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: `Customer-managed encryption key spec for a Tensorboard. If set, this Tensorboard and all sub-resources of this Tensorboard will be secured by this key.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"kms_key_name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
							Description: `The Cloud KMS resource identifier of the customer managed encryption key used to protect a resource.
Has the form: projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key. The key needs to be in the same region as where the resource is created.`,
						},
					},
				},
			},
			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: `The labels with user-defined metadata to organize your Tensorboards.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"region": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: `The region of the tensorboard. eg us-central1`,
			},
			"blob_storage_path_prefix": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Consumer project Cloud Storage path prefix used to store blob data, which can either be a bucket or directory. Does not end with a '/'.`,
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The timestamp of when the Tensorboard was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.`,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Name of the Tensorboard.`,
			},
			"run_count": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The number of Runs stored in this Tensorboard.`,
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The timestamp of when the Tensorboard was last updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.`,
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

func resourceVertexAITensorboardCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	displayNameProp, err := expandVertexAITensorboardDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	descriptionProp, err := expandVertexAITensorboardDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	encryptionSpecProp, err := expandVertexAITensorboardEncryptionSpec(d.Get("encryption_spec"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("encryption_spec"); !tpgresource.IsEmptyValue(reflect.ValueOf(encryptionSpecProp)) && (ok || !reflect.DeepEqual(v, encryptionSpecProp)) {
		obj["encryptionSpec"] = encryptionSpecProp
	}
	labelsProp, err := expandVertexAITensorboardLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{VertexAIBasePath}}projects/{{project}}/locations/{{region}}/tensorboards")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Tensorboard: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Tensorboard: %s", err)
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
		return fmt.Errorf("Error creating Tensorboard: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = VertexAIOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating Tensorboard", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create Tensorboard: %s", err)
	}

	if err := d.Set("name", flattenVertexAITensorboardName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = tpgresource.ReplaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Tensorboard %q: %#v", d.Id(), res)

	return resourceVertexAITensorboardRead(d, meta)
}

func resourceVertexAITensorboardRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{VertexAIBasePath}}{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Tensorboard: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("VertexAITensorboard %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Tensorboard: %s", err)
	}

	if err := d.Set("name", flattenVertexAITensorboardName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Tensorboard: %s", err)
	}
	if err := d.Set("display_name", flattenVertexAITensorboardDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading Tensorboard: %s", err)
	}
	if err := d.Set("description", flattenVertexAITensorboardDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading Tensorboard: %s", err)
	}
	if err := d.Set("encryption_spec", flattenVertexAITensorboardEncryptionSpec(res["encryptionSpec"], d, config)); err != nil {
		return fmt.Errorf("Error reading Tensorboard: %s", err)
	}
	if err := d.Set("blob_storage_path_prefix", flattenVertexAITensorboardBlobStoragePathPrefix(res["blobStoragePathPrefix"], d, config)); err != nil {
		return fmt.Errorf("Error reading Tensorboard: %s", err)
	}
	if err := d.Set("run_count", flattenVertexAITensorboardRunCount(res["runCount"], d, config)); err != nil {
		return fmt.Errorf("Error reading Tensorboard: %s", err)
	}
	if err := d.Set("create_time", flattenVertexAITensorboardCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Tensorboard: %s", err)
	}
	if err := d.Set("update_time", flattenVertexAITensorboardUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Tensorboard: %s", err)
	}
	if err := d.Set("labels", flattenVertexAITensorboardLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Tensorboard: %s", err)
	}

	return nil
}

func resourceVertexAITensorboardUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Tensorboard: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	displayNameProp, err := expandVertexAITensorboardDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	descriptionProp, err := expandVertexAITensorboardDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	labelsProp, err := expandVertexAITensorboardLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{VertexAIBasePath}}{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Tensorboard %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
	}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}

	if d.HasChange("labels") {
		updateMask = append(updateMask, "labels")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "PATCH",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutUpdate),
	})

	if err != nil {
		return fmt.Errorf("Error updating Tensorboard %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Tensorboard %q: %#v", d.Id(), res)
	}

	err = VertexAIOperationWaitTime(
		config, res, project, "Updating Tensorboard", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceVertexAITensorboardRead(d, meta)
}

func resourceVertexAITensorboardDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Tensorboard: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{VertexAIBasePath}}{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Tensorboard %q", d.Id())

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
		return transport_tpg.HandleNotFoundError(err, d, "Tensorboard")
	}

	err = VertexAIOperationWaitTime(
		config, res, project, "Deleting Tensorboard", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Tensorboard %q: %#v", d.Id(), res)
	return nil
}

func resourceVertexAITensorboardImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<region>[^/]+)/tensorboards/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{region}}/tensorboards/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	if err := d.Set("name", id); err != nil {
		return nil, fmt.Errorf("Error setting name for import: %s", err)
	}

	return []*schema.ResourceData{d}, nil
}

func flattenVertexAITensorboardName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVertexAITensorboardDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVertexAITensorboardDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVertexAITensorboardEncryptionSpec(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["kms_key_name"] =
		flattenVertexAITensorboardEncryptionSpecKmsKeyName(original["kmsKeyName"], d, config)
	return []interface{}{transformed}
}
func flattenVertexAITensorboardEncryptionSpecKmsKeyName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVertexAITensorboardBlobStoragePathPrefix(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVertexAITensorboardRunCount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVertexAITensorboardCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVertexAITensorboardUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVertexAITensorboardLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandVertexAITensorboardDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVertexAITensorboardDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVertexAITensorboardEncryptionSpec(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedKmsKeyName, err := expandVertexAITensorboardEncryptionSpecKmsKeyName(original["kms_key_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsKeyName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["kmsKeyName"] = transformedKmsKeyName
	}

	return transformed, nil
}

func expandVertexAITensorboardEncryptionSpecKmsKeyName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVertexAITensorboardLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
