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
)

func resourceSqlDatabase() *schema.Resource {
	return &schema.Resource{
		Create: resourceSqlDatabaseCreate,
		Read:   resourceSqlDatabaseRead,
		Update: resourceSqlDatabaseUpdate,
		Delete: resourceSqlDatabaseDelete,

		Importer: &schema.ResourceImporter{
			State: resourceSqlDatabaseImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(15 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"instance": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"charset": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"collation": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
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

func resourceSqlDatabaseCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	charsetProp, err := expandSqlDatabaseCharset(d.Get("charset"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("charset"); !isEmptyValue(reflect.ValueOf(charsetProp)) && (ok || !reflect.DeepEqual(v, charsetProp)) {
		obj["charset"] = charsetProp
	}
	collationProp, err := expandSqlDatabaseCollation(d.Get("collation"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("collation"); !isEmptyValue(reflect.ValueOf(collationProp)) && (ok || !reflect.DeepEqual(v, collationProp)) {
		obj["collation"] = collationProp
	}
	nameProp, err := expandSqlDatabaseName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	instanceProp, err := expandSqlDatabaseInstance(d.Get("instance"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("instance"); !isEmptyValue(reflect.ValueOf(instanceProp)) && (ok || !reflect.DeepEqual(v, instanceProp)) {
		obj["instance"] = instanceProp
	}

	lockName, err := replaceVars(d, config, "google-sql-database-instance-{{project}}-{{instance}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := replaceVars(d, config, "{{SqlBasePath}}projects/{{project}}/instances/{{instance}}/databases")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Database: %#v", obj)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "POST", project, url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Database: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{instance}}:{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	op := &sqladmin.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	waitErr := sqlAdminOperationWaitTime(
		config.clientSqlAdmin, op, project, "Creating Database",
		int(d.Timeout(schema.TimeoutCreate).Minutes()))

	if waitErr != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create Database: %s", waitErr)
	}

	log.Printf("[DEBUG] Finished creating Database %q: %#v", d.Id(), res)

	return resourceSqlDatabaseRead(d, meta)
}

func resourceSqlDatabaseRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{SqlBasePath}}projects/{{project}}/instances/{{instance}}/databases/{{name}}")
	if err != nil {
		return err
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequest(config, "GET", project, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("SqlDatabase %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Database: %s", err)
	}

	if err := d.Set("charset", flattenSqlDatabaseCharset(res["charset"], d)); err != nil {
		return fmt.Errorf("Error reading Database: %s", err)
	}
	if err := d.Set("collation", flattenSqlDatabaseCollation(res["collation"], d)); err != nil {
		return fmt.Errorf("Error reading Database: %s", err)
	}
	if err := d.Set("name", flattenSqlDatabaseName(res["name"], d)); err != nil {
		return fmt.Errorf("Error reading Database: %s", err)
	}
	if err := d.Set("instance", flattenSqlDatabaseInstance(res["instance"], d)); err != nil {
		return fmt.Errorf("Error reading Database: %s", err)
	}
	if err := d.Set("self_link", ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading Database: %s", err)
	}

	return nil
}

func resourceSqlDatabaseUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	charsetProp, err := expandSqlDatabaseCharset(d.Get("charset"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("charset"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, charsetProp)) {
		obj["charset"] = charsetProp
	}
	collationProp, err := expandSqlDatabaseCollation(d.Get("collation"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("collation"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, collationProp)) {
		obj["collation"] = collationProp
	}
	nameProp, err := expandSqlDatabaseName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	instanceProp, err := expandSqlDatabaseInstance(d.Get("instance"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("instance"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, instanceProp)) {
		obj["instance"] = instanceProp
	}

	lockName, err := replaceVars(d, config, "google-sql-database-instance-{{project}}-{{instance}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := replaceVars(d, config, "{{SqlBasePath}}projects/{{project}}/instances/{{instance}}/databases/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Database %q: %#v", d.Id(), obj)
	res, err := sendRequestWithTimeout(config, "PUT", project, url, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating Database %q: %s", d.Id(), err)
	}

	op := &sqladmin.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	err = sqlAdminOperationWaitTime(
		config.clientSqlAdmin, op, project, "Updating Database",
		int(d.Timeout(schema.TimeoutUpdate).Minutes()))

	if err != nil {
		return err
	}

	return resourceSqlDatabaseRead(d, meta)
}

func resourceSqlDatabaseDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	lockName, err := replaceVars(d, config, "google-sql-database-instance-{{project}}-{{instance}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := replaceVars(d, config, "{{SqlBasePath}}projects/{{project}}/instances/{{instance}}/databases/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Database %q", d.Id())

	res, err := sendRequestWithTimeout(config, "DELETE", project, url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "Database")
	}

	op := &sqladmin.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	err = sqlAdminOperationWaitTime(
		config.clientSqlAdmin, op, project, "Deleting Database",
		int(d.Timeout(schema.TimeoutDelete).Minutes()))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Database %q: %#v", d.Id(), res)
	return nil
}

func resourceSqlDatabaseImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/instances/(?P<instance>[^/]+)/databases/(?P<name>[^/]+)",
		"instances/(?P<instance>[^/]+)/databases/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<instance>[^/]+)/(?P<name>[^/]+)",
		"(?P<instance>[^/]+)/(?P<name>[^/]+)",
		"(?P<instance>[^/]+):(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "{{instance}}:{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenSqlDatabaseCharset(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenSqlDatabaseCollation(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenSqlDatabaseName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenSqlDatabaseInstance(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func expandSqlDatabaseCharset(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandSqlDatabaseCollation(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandSqlDatabaseName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandSqlDatabaseInstance(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
