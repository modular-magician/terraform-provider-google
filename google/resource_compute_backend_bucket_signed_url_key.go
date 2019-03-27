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
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"google.golang.org/api/compute/v1"
)

func resourceComputeBackendBucketSignedUrlKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeBackendBucketSignedUrlKeyCreate,
		Read:   resourceComputeBackendBucketSignedUrlKeyRead,
		Delete: resourceComputeBackendBucketSignedUrlKeyDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(240 * time.Second),
			Delete: schema.DefaultTimeout(240 * time.Second),
		},

		Schema: map[string]*schema.Schema{
			"backend_bucket": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
			},
			"key_value": {
				Type:      schema.TypeString,
				Required:  true,
				ForceNew:  true,
				Sensitive: true,
			},
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateRegexp(`^(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?)$`),
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func resourceComputeBackendBucketSignedUrlKeyCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	keyNameProp, err := expandComputeBackendBucketSignedUrlKeyName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(keyNameProp)) && (ok || !reflect.DeepEqual(v, keyNameProp)) {
		obj["keyName"] = keyNameProp
	}
	keyValueProp, err := expandComputeBackendBucketSignedUrlKeyKeyValue(d.Get("key_value"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("key_value"); !isEmptyValue(reflect.ValueOf(keyValueProp)) && (ok || !reflect.DeepEqual(v, keyValueProp)) {
		obj["keyValue"] = keyValueProp
	}
	backendBucketProp, err := expandComputeBackendBucketSignedUrlKeyBackendBucket(d.Get("backend_bucket"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("backend_bucket"); !isEmptyValue(reflect.ValueOf(backendBucketProp)) && (ok || !reflect.DeepEqual(v, backendBucketProp)) {
		obj["backendBucket"] = backendBucketProp
	}

	lockName, err := replaceVars(d, config, "signedUrlKey/{{project}}/backendBuckets/{{backend_bucket}}/")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/global/backendBuckets/{{backend_bucket}}/addSignedUrlKey")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new BackendBucketSignedUrlKey: %#v", obj)
	res, err := sendRequestWithTimeout(config, "POST", url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating BackendBucketSignedUrlKey: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	waitErr := computeOperationWaitTime(
		config.clientCompute, op, project, "Creating BackendBucketSignedUrlKey",
		int(d.Timeout(schema.TimeoutCreate).Minutes()))

	if waitErr != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create BackendBucketSignedUrlKey: %s", waitErr)
	}

	log.Printf("[DEBUG] Finished creating BackendBucketSignedUrlKey %q: %#v", d.Id(), res)

	return resourceComputeBackendBucketSignedUrlKeyRead(d, meta)
}

func resourceComputeBackendBucketSignedUrlKeyRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/global/backendBuckets/{{backend_bucket}}")
	if err != nil {
		return err
	}

	res, err := sendRequest(config, "GET", url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeBackendBucketSignedUrlKey %q", d.Id()))
	}

	res, err = getNestedResourceComputeBackendBucketSignedUrlKey(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Object isn't there any more - remove it from the state.
		log.Printf("[DEBUG] Removing ComputeBackendBucketSignedUrlKey because it couldn't be matched.")
		d.SetId("")
		return nil
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading BackendBucketSignedUrlKey: %s", err)
	}

	if err := d.Set("name", flattenComputeBackendBucketSignedUrlKeyName(res["keyName"], d)); err != nil {
		return fmt.Errorf("Error reading BackendBucketSignedUrlKey: %s", err)
	}

	return nil
}

func resourceComputeBackendBucketSignedUrlKeyDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	lockName, err := replaceVars(d, config, "signedUrlKey/{{project}}/backendBuckets/{{backend_bucket}}/")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/global/backendBuckets/{{backend_bucket}}/deleteSignedUrlKey?keyName={{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting BackendBucketSignedUrlKey %q", d.Id())
	res, err := sendRequestWithTimeout(config, "POST", url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "BackendBucketSignedUrlKey")
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	err = computeOperationWaitTime(
		config.clientCompute, op, project, "Deleting BackendBucketSignedUrlKey",
		int(d.Timeout(schema.TimeoutDelete).Minutes()))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting BackendBucketSignedUrlKey %q: %#v", d.Id(), res)
	return nil
}

func flattenComputeBackendBucketSignedUrlKeyName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func expandComputeBackendBucketSignedUrlKeyName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendBucketSignedUrlKeyKeyValue(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeBackendBucketSignedUrlKeyBackendBucket(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("backendBuckets", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for backend_bucket: %s", err)
	}
	return f.RelativeLink(), nil
}

func getNestedResourceComputeBackendBucketSignedUrlKey(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	config := meta.(*Config)

	var v interface{}
	var ok bool
	v, ok = res["cdnPolicy"]
	if !ok || v == nil {
		return nil, nil
	}
	res = v.(map[string]interface{})
	v, ok = res["signedUrlKeyNames"]
	if !ok || v == nil {
		return nil, nil
	}
	// Final nested resource is either a list of resources we need to filter
	// or just the resource itself, which we return.
	switch v.(type) {
	case []interface{}:
		break
	case map[string]interface{}:
		return v.(map[string]interface{}), nil
	default:
		return nil, fmt.Errorf("invalid value for cdnPolicy.signedUrlKeyNames: %v", v)
	}

	items := v.([]interface{})
	for _, vRaw := range items {
		// If only an id is given in parent resource,
		// construct a resource map for that id KV pair.
		item := map[string]interface{}{"keyName": vRaw}

		id, err := expandComputeBackendBucketSignedUrlKeyName(d.Get("name"), d, config)
		if err != nil {
			return nil, err
		}

		itemId := flattenComputeBackendBucketSignedUrlKeyName(item["keyName"], d)
		log.Printf("[DEBUG] Checking equality of %#v, %#v", itemId, id)
		if !reflect.DeepEqual(itemId, id) {
			continue
		}
		return item, nil
	}
	return nil, nil
}
