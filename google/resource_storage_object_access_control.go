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
	"strconv"
	"time"
)

func resourceStorageObjectAccessControl() *schema.Resource {
	return &schema.Resource{
		Create: resourceStorageObjectAccessControlCreate,
		Read:   resourceStorageObjectAccessControlRead,
		Update: resourceStorageObjectAccessControlUpdate,
		Delete: resourceStorageObjectAccessControlDelete,

		Importer: &schema.ResourceImporter{
			State: resourceStorageObjectAccessControlImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(240 * time.Second),
			Update: schema.DefaultTimeout(240 * time.Second),
			Delete: schema.DefaultTimeout(240 * time.Second),
		},

		Schema: map[string]*schema.Schema{
			"bucket": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
			},
			"entity": {
				Type:     schema.TypeString,
				Required: true,
			},
			"object": {
				Type:     schema.TypeString,
				Required: true,
			},
			"role": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"OWNER", "READER"}, false),
			},
			"domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"email": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"generation": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"project_team": {
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"project_number": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"team": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice([]string{"editors", "owners", "viewers", ""}, false),
						},
					},
				},
			},
		},
	}
}

func resourceStorageObjectAccessControlCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	bucketProp, err := expandStorageObjectAccessControlBucket(d.Get("bucket"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("bucket"); !isEmptyValue(reflect.ValueOf(bucketProp)) && (ok || !reflect.DeepEqual(v, bucketProp)) {
		obj["bucket"] = bucketProp
	}
	entityProp, err := expandStorageObjectAccessControlEntity(d.Get("entity"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("entity"); !isEmptyValue(reflect.ValueOf(entityProp)) && (ok || !reflect.DeepEqual(v, entityProp)) {
		obj["entity"] = entityProp
	}
	objectProp, err := expandStorageObjectAccessControlObject(d.Get("object"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("object"); !isEmptyValue(reflect.ValueOf(objectProp)) && (ok || !reflect.DeepEqual(v, objectProp)) {
		obj["object"] = objectProp
	}
	roleProp, err := expandStorageObjectAccessControlRole(d.Get("role"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("role"); !isEmptyValue(reflect.ValueOf(roleProp)) && (ok || !reflect.DeepEqual(v, roleProp)) {
		obj["role"] = roleProp
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/storage/v1/b/{{bucket}}/o/{{object}}/acl")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new ObjectAccessControl: %#v", obj)
	res, err := sendRequestWithTimeout(config, "POST", url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating ObjectAccessControl: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{bucket}}/{{object}}/{{entity}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating ObjectAccessControl %q: %#v", d.Id(), res)

	return resourceStorageObjectAccessControlRead(d, meta)
}

func resourceStorageObjectAccessControlRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "https://www.googleapis.com/storage/v1/b/{{bucket}}/o/{{object}}/acl/{{entity}}")
	if err != nil {
		return err
	}

	res, err := sendRequest(config, "GET", url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("StorageObjectAccessControl %q", d.Id()))
	}

	if err := d.Set("bucket", flattenStorageObjectAccessControlBucket(res["bucket"], d)); err != nil {
		return fmt.Errorf("Error reading ObjectAccessControl: %s", err)
	}
	if err := d.Set("domain", flattenStorageObjectAccessControlDomain(res["domain"], d)); err != nil {
		return fmt.Errorf("Error reading ObjectAccessControl: %s", err)
	}
	if err := d.Set("email", flattenStorageObjectAccessControlEmail(res["email"], d)); err != nil {
		return fmt.Errorf("Error reading ObjectAccessControl: %s", err)
	}
	if err := d.Set("entity", flattenStorageObjectAccessControlEntity(res["entity"], d)); err != nil {
		return fmt.Errorf("Error reading ObjectAccessControl: %s", err)
	}
	if err := d.Set("entity_id", flattenStorageObjectAccessControlEntityId(res["entityId"], d)); err != nil {
		return fmt.Errorf("Error reading ObjectAccessControl: %s", err)
	}
	if err := d.Set("generation", flattenStorageObjectAccessControlGeneration(res["generation"], d)); err != nil {
		return fmt.Errorf("Error reading ObjectAccessControl: %s", err)
	}
	if err := d.Set("object", flattenStorageObjectAccessControlObject(res["object"], d)); err != nil {
		return fmt.Errorf("Error reading ObjectAccessControl: %s", err)
	}
	if err := d.Set("project_team", flattenStorageObjectAccessControlProjectTeam(res["projectTeam"], d)); err != nil {
		return fmt.Errorf("Error reading ObjectAccessControl: %s", err)
	}
	if err := d.Set("role", flattenStorageObjectAccessControlRole(res["role"], d)); err != nil {
		return fmt.Errorf("Error reading ObjectAccessControl: %s", err)
	}

	return nil
}

func resourceStorageObjectAccessControlUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	bucketProp, err := expandStorageObjectAccessControlBucket(d.Get("bucket"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("bucket"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, bucketProp)) {
		obj["bucket"] = bucketProp
	}
	entityProp, err := expandStorageObjectAccessControlEntity(d.Get("entity"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("entity"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, entityProp)) {
		obj["entity"] = entityProp
	}
	objectProp, err := expandStorageObjectAccessControlObject(d.Get("object"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("object"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, objectProp)) {
		obj["object"] = objectProp
	}
	roleProp, err := expandStorageObjectAccessControlRole(d.Get("role"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("role"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, roleProp)) {
		obj["role"] = roleProp
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/storage/v1/b/{{bucket}}/o/{{object}}/acl/{{entity}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating ObjectAccessControl %q: %#v", d.Id(), obj)
	_, err = sendRequestWithTimeout(config, "PUT", url, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating ObjectAccessControl %q: %s", d.Id(), err)
	}

	return resourceStorageObjectAccessControlRead(d, meta)
}

func resourceStorageObjectAccessControlDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "https://www.googleapis.com/storage/v1/b/{{bucket}}/o/{{object}}/acl/{{entity}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting ObjectAccessControl %q", d.Id())
	res, err := sendRequestWithTimeout(config, "DELETE", url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "ObjectAccessControl")
	}

	log.Printf("[DEBUG] Finished deleting ObjectAccessControl %q: %#v", d.Id(), res)
	return nil
}

func resourceStorageObjectAccessControlImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{"(?P<bucket>[^/]+)/(?P<object>[^/]+)/(?P<entity>[^/]+)"}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "{{bucket}}/{{object}}/{{entity}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenStorageObjectAccessControlBucket(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func flattenStorageObjectAccessControlDomain(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenStorageObjectAccessControlEmail(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenStorageObjectAccessControlEntity(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenStorageObjectAccessControlEntityId(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenStorageObjectAccessControlGeneration(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenStorageObjectAccessControlObject(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenStorageObjectAccessControlProjectTeam(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["project_number"] =
		flattenStorageObjectAccessControlProjectTeamProjectNumber(original["projectNumber"], d)
	transformed["team"] =
		flattenStorageObjectAccessControlProjectTeamTeam(original["team"], d)
	return []interface{}{transformed}
}
func flattenStorageObjectAccessControlProjectTeamProjectNumber(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenStorageObjectAccessControlProjectTeamTeam(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenStorageObjectAccessControlRole(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func expandStorageObjectAccessControlBucket(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandStorageObjectAccessControlEntity(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandStorageObjectAccessControlObject(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandStorageObjectAccessControlRole(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}
