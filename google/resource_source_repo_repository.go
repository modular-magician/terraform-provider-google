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
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSourceRepoRepository() *schema.Resource {
	return &schema.Resource{
		Create: resourceSourceRepoRepositoryCreate,
		Read:   resourceSourceRepoRepositoryRead,
		Delete: resourceSourceRepoRepositoryDelete,

		Importer: &schema.ResourceImporter{
			State: resourceSourceRepoRepositoryImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"size": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"url": {
				Type:     schema.TypeString,
				Computed: true,
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

func resourceSourceRepoRepositoryCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	nameProp, err := expandSourceRepoRepositoryName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}

	url, err := replaceVars(d, config, "{{SourceRepoBasePath}}projects/{{project}}/repos")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Repository: %#v", obj)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "POST", project, url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Repository: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/repos/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Repository %q: %#v", d.Id(), res)

	return resourceSourceRepoRepositoryRead(d, meta)
}

func resourceSourceRepoRepositoryRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{SourceRepoBasePath}}projects/{{project}}/repos/{{name}}")
	if err != nil {
		return err
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequest(config, "GET", project, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("SourceRepoRepository %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Repository: %s", err)
	}

	if err := d.Set("name", flattenSourceRepoRepositoryName(res["name"], d)); err != nil {
		return fmt.Errorf("Error reading Repository: %s", err)
	}
	if err := d.Set("url", flattenSourceRepoRepositoryUrl(res["url"], d)); err != nil {
		return fmt.Errorf("Error reading Repository: %s", err)
	}
	if err := d.Set("size", flattenSourceRepoRepositorySize(res["size"], d)); err != nil {
		return fmt.Errorf("Error reading Repository: %s", err)
	}

	return nil
}

func resourceSourceRepoRepositoryDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{SourceRepoBasePath}}projects/{{project}}/repos/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Repository %q", d.Id())

	res, err := sendRequestWithTimeout(config, "DELETE", project, url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "Repository")
	}

	log.Printf("[DEBUG] Finished deleting Repository %q: %#v", d.Id(), res)
	return nil
}

func resourceSourceRepoRepositoryImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/repos/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/repos/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenSourceRepoRepositoryName(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}

	// We can't use a standard name_from_self_link because the name can include /'s
	parts := strings.SplitAfterN(v.(string), "/", 4)
	return parts[3]
}

func flattenSourceRepoRepositoryUrl(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenSourceRepoRepositorySize(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func expandSourceRepoRepositoryName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return replaceVars(d, config, "projects/{{project}}/repos/{{name}}")
}
