// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
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

package google

import (
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strconv"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
)

const datasetIdRegexp = `[0-9A-Za-z_]+`

func validateDatasetId(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)
	if !regexp.MustCompile(datasetIdRegexp).MatchString(value) {
		errors = append(errors, fmt.Errorf(
			"%q must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_)", k))
	}

	if len(value) > 1024 {
		errors = append(errors, fmt.Errorf(
			"%q cannot be greater than 1,024 characters", k))
	}

	return
}

func validateDefaultTableExpirationMs(v interface{}, k string) (ws []string, errors []error) {
	value := v.(int)
	if value < 3600000 {
		errors = append(errors, fmt.Errorf("%q cannot be shorter than 3600000 milliseconds (one hour)", k))
	}

	return
}

func resourceBigQueryDataset() *schema.Resource {
	return &schema.Resource{
		Create: resourceBigQueryDatasetCreate,
		Read:   resourceBigQueryDatasetRead,
		Update: resourceBigQueryDatasetUpdate,
		Delete: resourceBigQueryDatasetDelete,

		Importer: &schema.ResourceImporter{
			State: resourceBigQueryDatasetImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"dataset_id": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateDatasetId,
			},

			"access": {
				Type:     schema.TypeSet,
				Computed: true,
				Optional: true,
				Elem:     bigqueryDatasetAccessSchema(),
				// Default schema.HashSchema is used.
			},
			"default_partition_expiration_ms": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"default_table_expiration_ms": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validateDefaultTableExpirationMs,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"friendly_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"location": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Default:  "US",
			},
			"creation_time": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_modified_time": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"delete_contents_on_destroy": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"self_link": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func bigqueryDatasetAccessSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"domain": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"group_by_email": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"role": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"special_group": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_by_email": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"view": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dataset_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"project_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"table_id": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func resourceBigQueryDatasetCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	accessProp, err := expandBigQueryDatasetAccess(d.Get("access"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("access"); !isEmptyValue(reflect.ValueOf(accessProp)) && (ok || !reflect.DeepEqual(v, accessProp)) {
		obj["access"] = accessProp
	}
	datasetReferenceProp, err := expandBigQueryDatasetDatasetReference(nil, d, config)
	if err != nil {
		return err
	} else if !isEmptyValue(reflect.ValueOf(datasetReferenceProp)) {
		obj["datasetReference"] = datasetReferenceProp
	}
	defaultTableExpirationMsProp, err := expandBigQueryDatasetDefaultTableExpirationMs(d.Get("default_table_expiration_ms"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("default_table_expiration_ms"); !isEmptyValue(reflect.ValueOf(defaultTableExpirationMsProp)) && (ok || !reflect.DeepEqual(v, defaultTableExpirationMsProp)) {
		obj["defaultTableExpirationMs"] = defaultTableExpirationMsProp
	}
	defaultPartitionExpirationMsProp, err := expandBigQueryDatasetDefaultPartitionExpirationMs(d.Get("default_partition_expiration_ms"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("default_partition_expiration_ms"); !isEmptyValue(reflect.ValueOf(defaultPartitionExpirationMsProp)) && (ok || !reflect.DeepEqual(v, defaultPartitionExpirationMsProp)) {
		obj["defaultPartitionExpirationMs"] = defaultPartitionExpirationMsProp
	}
	descriptionProp, err := expandBigQueryDatasetDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	friendlyNameProp, err := expandBigQueryDatasetFriendlyName(d.Get("friendly_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("friendly_name"); !isEmptyValue(reflect.ValueOf(friendlyNameProp)) && (ok || !reflect.DeepEqual(v, friendlyNameProp)) {
		obj["friendlyName"] = friendlyNameProp
	}
	labelsProp, err := expandBigQueryDatasetLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	locationProp, err := expandBigQueryDatasetLocation(d.Get("location"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("location"); !isEmptyValue(reflect.ValueOf(locationProp)) && (ok || !reflect.DeepEqual(v, locationProp)) {
		obj["location"] = locationProp
	}

	url, err := replaceVars(d, config, "{{BigQueryBasePath}}projects/{{project}}/datasets")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Dataset: %#v", obj)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "POST", project, url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Dataset: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{project}}:{{dataset_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Dataset %q: %#v", d.Id(), res)

	return resourceBigQueryDatasetRead(d, meta)
}

func resourceBigQueryDatasetRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{BigQueryBasePath}}projects/{{project}}/datasets/{{dataset_id}}")
	if err != nil {
		return err
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequest(config, "GET", project, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("BigQueryDataset %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Dataset: %s", err)
	}

	if err := d.Set("access", flattenBigQueryDatasetAccess(res["access"], d)); err != nil {
		return fmt.Errorf("Error reading Dataset: %s", err)
	}
	if err := d.Set("creation_time", flattenBigQueryDatasetCreationTime(res["creationTime"], d)); err != nil {
		return fmt.Errorf("Error reading Dataset: %s", err)
	}
	// Terraform must set the top level schema field, but since this object contains collapsed properties
	// it's difficult to know what the top level should be. Instead we just loop over the map returned from flatten.
	if flattenedProp := flattenBigQueryDatasetDatasetReference(res["datasetReference"], d); flattenedProp != nil {
		casted := flattenedProp.([]interface{})[0]
		if casted != nil {
			for k, v := range casted.(map[string]interface{}) {
				d.Set(k, v)
			}
		}
	}
	if err := d.Set("default_table_expiration_ms", flattenBigQueryDatasetDefaultTableExpirationMs(res["defaultTableExpirationMs"], d)); err != nil {
		return fmt.Errorf("Error reading Dataset: %s", err)
	}
	if err := d.Set("default_partition_expiration_ms", flattenBigQueryDatasetDefaultPartitionExpirationMs(res["defaultPartitionExpirationMs"], d)); err != nil {
		return fmt.Errorf("Error reading Dataset: %s", err)
	}
	if err := d.Set("description", flattenBigQueryDatasetDescription(res["description"], d)); err != nil {
		return fmt.Errorf("Error reading Dataset: %s", err)
	}
	if err := d.Set("etag", flattenBigQueryDatasetEtag(res["etag"], d)); err != nil {
		return fmt.Errorf("Error reading Dataset: %s", err)
	}
	if err := d.Set("friendly_name", flattenBigQueryDatasetFriendlyName(res["friendlyName"], d)); err != nil {
		return fmt.Errorf("Error reading Dataset: %s", err)
	}
	if err := d.Set("labels", flattenBigQueryDatasetLabels(res["labels"], d)); err != nil {
		return fmt.Errorf("Error reading Dataset: %s", err)
	}
	if err := d.Set("last_modified_time", flattenBigQueryDatasetLastModifiedTime(res["lastModifiedTime"], d)); err != nil {
		return fmt.Errorf("Error reading Dataset: %s", err)
	}
	if err := d.Set("location", flattenBigQueryDatasetLocation(res["location"], d)); err != nil {
		return fmt.Errorf("Error reading Dataset: %s", err)
	}
	if err := d.Set("self_link", ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading Dataset: %s", err)
	}

	return nil
}

func resourceBigQueryDatasetUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	accessProp, err := expandBigQueryDatasetAccess(d.Get("access"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("access"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, accessProp)) {
		obj["access"] = accessProp
	}
	datasetReferenceProp, err := expandBigQueryDatasetDatasetReference(nil, d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("dataset_reference"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, datasetReferenceProp)) {
		obj["datasetReference"] = datasetReferenceProp
	}
	defaultTableExpirationMsProp, err := expandBigQueryDatasetDefaultTableExpirationMs(d.Get("default_table_expiration_ms"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("default_table_expiration_ms"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, defaultTableExpirationMsProp)) {
		obj["defaultTableExpirationMs"] = defaultTableExpirationMsProp
	}
	defaultPartitionExpirationMsProp, err := expandBigQueryDatasetDefaultPartitionExpirationMs(d.Get("default_partition_expiration_ms"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("default_partition_expiration_ms"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, defaultPartitionExpirationMsProp)) {
		obj["defaultPartitionExpirationMs"] = defaultPartitionExpirationMsProp
	}
	descriptionProp, err := expandBigQueryDatasetDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	friendlyNameProp, err := expandBigQueryDatasetFriendlyName(d.Get("friendly_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("friendly_name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, friendlyNameProp)) {
		obj["friendlyName"] = friendlyNameProp
	}
	labelsProp, err := expandBigQueryDatasetLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	locationProp, err := expandBigQueryDatasetLocation(d.Get("location"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("location"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, locationProp)) {
		obj["location"] = locationProp
	}

	url, err := replaceVars(d, config, "{{BigQueryBasePath}}projects/{{project}}/datasets/{{dataset_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Dataset %q: %#v", d.Id(), obj)
	_, err = sendRequestWithTimeout(config, "PUT", project, url, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating Dataset %q: %s", d.Id(), err)
	}

	return resourceBigQueryDatasetRead(d, meta)
}

func resourceBigQueryDatasetDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{BigQueryBasePath}}projects/{{project}}/datasets/{{dataset_id}}?deleteContents={{delete_contents_on_destroy}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Dataset %q", d.Id())

	res, err := sendRequestWithTimeout(config, "DELETE", project, url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "Dataset")
	}

	log.Printf("[DEBUG] Finished deleting Dataset %q: %#v", d.Id(), res)
	return nil
}

func resourceBigQueryDatasetImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	// Explicitly set to default as a workaround for `ImportStateVerify` tests, and so that users
	// don't see a diff immediately after import.
	d.Set("delete_contents_on_destroy", false)

	// Auto generated importers do not support neither ':' or '/' as Id separators
	// so we customize it here.
	id := d.Id()

	pd := fmt.Sprintf("(%s)[:/](%s)", ProjectRegex, datasetIdRegexp)
	re := regexp.MustCompile(pd)
	if parts := re.FindStringSubmatch(id); parts != nil {
		d.Set("project", parts[1])
		d.Set("dataset_id", parts[2])
		return []*schema.ResourceData{d}, nil
	}

	pd = fmt.Sprintf("(%s)", datasetIdRegexp)
	re = regexp.MustCompile(pd)
	if parts := re.FindStringSubmatch(id); parts != nil {
		d.Set("dataset_id", parts[1])
		return []*schema.ResourceData{d}, nil
	}

	return nil, fmt.Errorf("cannot parse import ID: %s", id)
}

func flattenBigQueryDatasetAccess(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := schema.NewSet(schema.HashResource(bigqueryDatasetAccessSchema()), []interface{}{})
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed.Add(map[string]interface{}{
			"domain":         flattenBigQueryDatasetAccessDomain(original["domain"], d),
			"group_by_email": flattenBigQueryDatasetAccessGroupByEmail(original["groupByEmail"], d),
			"role":           flattenBigQueryDatasetAccessRole(original["role"], d),
			"special_group":  flattenBigQueryDatasetAccessSpecialGroup(original["specialGroup"], d),
			"user_by_email":  flattenBigQueryDatasetAccessUserByEmail(original["userByEmail"], d),
			"view":           flattenBigQueryDatasetAccessView(original["view"], d),
		})
	}
	return transformed
}
func flattenBigQueryDatasetAccessDomain(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenBigQueryDatasetAccessGroupByEmail(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenBigQueryDatasetAccessRole(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenBigQueryDatasetAccessSpecialGroup(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenBigQueryDatasetAccessUserByEmail(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenBigQueryDatasetAccessView(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["dataset_id"] =
		flattenBigQueryDatasetAccessViewDatasetId(original["datasetId"], d)
	transformed["project_id"] =
		flattenBigQueryDatasetAccessViewProjectId(original["projectId"], d)
	transformed["table_id"] =
		flattenBigQueryDatasetAccessViewTableId(original["tableId"], d)
	return []interface{}{transformed}
}
func flattenBigQueryDatasetAccessViewDatasetId(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenBigQueryDatasetAccessViewProjectId(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenBigQueryDatasetAccessViewTableId(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenBigQueryDatasetCreationTime(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenBigQueryDatasetDatasetReference(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["dataset_id"] =
		flattenBigQueryDatasetDatasetReferenceDatasetId(original["datasetId"], d)
	return []interface{}{transformed}
}
func flattenBigQueryDatasetDatasetReferenceDatasetId(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenBigQueryDatasetDefaultTableExpirationMs(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenBigQueryDatasetDefaultPartitionExpirationMs(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenBigQueryDatasetDescription(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenBigQueryDatasetEtag(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenBigQueryDatasetFriendlyName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenBigQueryDatasetLabels(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenBigQueryDatasetLastModifiedTime(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

// Older Datasets in BigQuery have no Location set in the API response. This may be an issue when importing
// datasets created before BigQuery was available in multiple zones. We can safely assume that these datasets
// are in the US, as this was the default at the time.
func flattenBigQueryDatasetLocation(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return "US"
	}
	return v
}

func expandBigQueryDatasetAccess(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedDomain, err := expandBigQueryDatasetAccessDomain(original["domain"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDomain); val.IsValid() && !isEmptyValue(val) {
			transformed["domain"] = transformedDomain
		}

		transformedGroupByEmail, err := expandBigQueryDatasetAccessGroupByEmail(original["group_by_email"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedGroupByEmail); val.IsValid() && !isEmptyValue(val) {
			transformed["groupByEmail"] = transformedGroupByEmail
		}

		transformedRole, err := expandBigQueryDatasetAccessRole(original["role"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedRole); val.IsValid() && !isEmptyValue(val) {
			transformed["role"] = transformedRole
		}

		transformedSpecialGroup, err := expandBigQueryDatasetAccessSpecialGroup(original["special_group"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedSpecialGroup); val.IsValid() && !isEmptyValue(val) {
			transformed["specialGroup"] = transformedSpecialGroup
		}

		transformedUserByEmail, err := expandBigQueryDatasetAccessUserByEmail(original["user_by_email"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedUserByEmail); val.IsValid() && !isEmptyValue(val) {
			transformed["userByEmail"] = transformedUserByEmail
		}

		transformedView, err := expandBigQueryDatasetAccessView(original["view"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedView); val.IsValid() && !isEmptyValue(val) {
			transformed["view"] = transformedView
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandBigQueryDatasetAccessDomain(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetAccessGroupByEmail(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetAccessRole(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetAccessSpecialGroup(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetAccessUserByEmail(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetAccessView(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDatasetId, err := expandBigQueryDatasetAccessViewDatasetId(original["dataset_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDatasetId); val.IsValid() && !isEmptyValue(val) {
		transformed["datasetId"] = transformedDatasetId
	}

	transformedProjectId, err := expandBigQueryDatasetAccessViewProjectId(original["project_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProjectId); val.IsValid() && !isEmptyValue(val) {
		transformed["projectId"] = transformedProjectId
	}

	transformedTableId, err := expandBigQueryDatasetAccessViewTableId(original["table_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTableId); val.IsValid() && !isEmptyValue(val) {
		transformed["tableId"] = transformedTableId
	}

	return transformed, nil
}

func expandBigQueryDatasetAccessViewDatasetId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetAccessViewProjectId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetAccessViewTableId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetDatasetReference(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	transformed := make(map[string]interface{})
	transformedDatasetId, err := expandBigQueryDatasetDatasetReferenceDatasetId(d.Get("dataset_id"), d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDatasetId); val.IsValid() && !isEmptyValue(val) {
		transformed["datasetId"] = transformedDatasetId
	}

	return transformed, nil
}

func expandBigQueryDatasetDatasetReferenceDatasetId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetDefaultTableExpirationMs(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetDefaultPartitionExpirationMs(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetFriendlyName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandBigQueryDatasetLocation(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
