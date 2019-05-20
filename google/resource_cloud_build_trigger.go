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
)

func resourceCloudBuildTrigger() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudBuildTriggerCreate,
		Read:   resourceCloudBuildTriggerRead,
		Update: resourceCloudBuildTriggerUpdate,
		Delete: resourceCloudBuildTriggerDelete,

		Importer: &schema.ResourceImporter{
			State: resourceCloudBuildTriggerImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(240 * time.Second),
			Update: schema.DefaultTimeout(240 * time.Second),
			Delete: schema.DefaultTimeout(240 * time.Second),
		},

		SchemaVersion: 1,

		Schema: map[string]*schema.Schema{
			"build": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"images": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"step": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"args": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"dir": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"entrypoint": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"env": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"id": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"secret_env": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"timeout": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"timing": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"volumes": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"path": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"wait_for": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"tags": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
				ConflictsWith: []string{"filename"},
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"disabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"filename": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"build"},
			},
			"ignored_files": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"included_files": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"substitutions": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"trigger_template": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"branch_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"commit_sha": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"dir": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"project_id": {
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
						},
						"repo_name": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "default",
						},
						"tag_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"trigger_id": {
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

func resourceCloudBuildTriggerCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	descriptionProp, err := expandCloudBuildTriggerDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	disabledProp, err := expandCloudBuildTriggerDisabled(d.Get("disabled"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("disabled"); !isEmptyValue(reflect.ValueOf(disabledProp)) && (ok || !reflect.DeepEqual(v, disabledProp)) {
		obj["disabled"] = disabledProp
	}
	substitutionsProp, err := expandCloudBuildTriggerSubstitutions(d.Get("substitutions"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("substitutions"); !isEmptyValue(reflect.ValueOf(substitutionsProp)) && (ok || !reflect.DeepEqual(v, substitutionsProp)) {
		obj["substitutions"] = substitutionsProp
	}
	filenameProp, err := expandCloudBuildTriggerFilename(d.Get("filename"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("filename"); !isEmptyValue(reflect.ValueOf(filenameProp)) && (ok || !reflect.DeepEqual(v, filenameProp)) {
		obj["filename"] = filenameProp
	}
	ignoredFilesProp, err := expandCloudBuildTriggerIgnoredFiles(d.Get("ignored_files"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ignored_files"); !isEmptyValue(reflect.ValueOf(ignoredFilesProp)) && (ok || !reflect.DeepEqual(v, ignoredFilesProp)) {
		obj["ignoredFiles"] = ignoredFilesProp
	}
	includedFilesProp, err := expandCloudBuildTriggerIncludedFiles(d.Get("included_files"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("included_files"); !isEmptyValue(reflect.ValueOf(includedFilesProp)) && (ok || !reflect.DeepEqual(v, includedFilesProp)) {
		obj["includedFiles"] = includedFilesProp
	}
	triggerTemplateProp, err := expandCloudBuildTriggerTriggerTemplate(d.Get("trigger_template"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("trigger_template"); !isEmptyValue(reflect.ValueOf(triggerTemplateProp)) && (ok || !reflect.DeepEqual(v, triggerTemplateProp)) {
		obj["triggerTemplate"] = triggerTemplateProp
	}
	buildProp, err := expandCloudBuildTriggerBuild(d.Get("build"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("build"); !isEmptyValue(reflect.ValueOf(buildProp)) && (ok || !reflect.DeepEqual(v, buildProp)) {
		obj["build"] = buildProp
	}

	url, err := replaceVars(d, config, "https://cloudbuild.googleapis.com/v1/projects/{{project}}/triggers")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Trigger: %#v", obj)
	res, err := sendRequestWithTimeout(config, "POST", url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Trigger: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{project}}/{{trigger_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Trigger %q: %#v", d.Id(), res)

	// `name` is autogenerated from the api so needs to be set post-create
	triggerId, ok := res["id"]
	if !ok {
		return fmt.Errorf("Create response didn't contain id. Create may not have succeeded.")
	}
	d.Set("trigger_id", triggerId.(string))

	// Store the ID now. We tried to set it before and it failed because
	// trigger_id didn't exist yet.
	id, err = replaceVars(d, config, "{{project}}/{{trigger_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return resourceCloudBuildTriggerRead(d, meta)
}

func resourceCloudBuildTriggerRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "https://cloudbuild.googleapis.com/v1/projects/{{project}}/triggers/{{trigger_id}}")
	if err != nil {
		return err
	}

	res, err := sendRequest(config, "GET", url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("CloudBuildTrigger %q", d.Id()))
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Trigger: %s", err)
	}

	// Terraform must set the top level schema field, but since this object contains collapsed properties
	// it's difficult to know what the top level should be. Instead we just loop over the map returned from flatten.
	if err := d.Set("trigger_id", flattenCloudBuildTriggerTrigger_id(res["id"], d)); err != nil {
		return fmt.Errorf("Error reading Trigger: %s", err)
	}
	// Terraform must set the top level schema field, but since this object contains collapsed properties
	// it's difficult to know what the top level should be. Instead we just loop over the map returned from flatten.
	if err := d.Set("description", flattenCloudBuildTriggerDescription(res["description"], d)); err != nil {
		return fmt.Errorf("Error reading Trigger: %s", err)
	}
	// Terraform must set the top level schema field, but since this object contains collapsed properties
	// it's difficult to know what the top level should be. Instead we just loop over the map returned from flatten.
	if err := d.Set("disabled", flattenCloudBuildTriggerDisabled(res["disabled"], d)); err != nil {
		return fmt.Errorf("Error reading Trigger: %s", err)
	}
	// Terraform must set the top level schema field, but since this object contains collapsed properties
	// it's difficult to know what the top level should be. Instead we just loop over the map returned from flatten.
	if err := d.Set("create_time", flattenCloudBuildTriggerCreateTime(res["createTime"], d)); err != nil {
		return fmt.Errorf("Error reading Trigger: %s", err)
	}
	// Terraform must set the top level schema field, but since this object contains collapsed properties
	// it's difficult to know what the top level should be. Instead we just loop over the map returned from flatten.
	if err := d.Set("substitutions", flattenCloudBuildTriggerSubstitutions(res["substitutions"], d)); err != nil {
		return fmt.Errorf("Error reading Trigger: %s", err)
	}
	// Terraform must set the top level schema field, but since this object contains collapsed properties
	// it's difficult to know what the top level should be. Instead we just loop over the map returned from flatten.
	if err := d.Set("filename", flattenCloudBuildTriggerFilename(res["filename"], d)); err != nil {
		return fmt.Errorf("Error reading Trigger: %s", err)
	}
	// Terraform must set the top level schema field, but since this object contains collapsed properties
	// it's difficult to know what the top level should be. Instead we just loop over the map returned from flatten.
	if err := d.Set("ignored_files", flattenCloudBuildTriggerIgnoredFiles(res["ignoredFiles"], d)); err != nil {
		return fmt.Errorf("Error reading Trigger: %s", err)
	}
	// Terraform must set the top level schema field, but since this object contains collapsed properties
	// it's difficult to know what the top level should be. Instead we just loop over the map returned from flatten.
	if err := d.Set("included_files", flattenCloudBuildTriggerIncludedFiles(res["includedFiles"], d)); err != nil {
		return fmt.Errorf("Error reading Trigger: %s", err)
	}
	// Terraform must set the top level schema field, but since this object contains collapsed properties
	// it's difficult to know what the top level should be. Instead we just loop over the map returned from flatten.
	if err := d.Set("trigger_template", flattenCloudBuildTriggerTriggerTemplate(res["triggerTemplate"], d)); err != nil {
		return fmt.Errorf("Error reading Trigger: %s", err)
	}
	// Terraform must set the top level schema field, but since this object contains collapsed properties
	// it's difficult to know what the top level should be. Instead we just loop over the map returned from flatten.
	if err := d.Set("build", flattenCloudBuildTriggerBuild(res["build"], d)); err != nil {
		return fmt.Errorf("Error reading Trigger: %s", err)
	}

	return nil
}

func resourceCloudBuildTriggerUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	descriptionProp, err := expandCloudBuildTriggerDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	disabledProp, err := expandCloudBuildTriggerDisabled(d.Get("disabled"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("disabled"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, disabledProp)) {
		obj["disabled"] = disabledProp
	}
	substitutionsProp, err := expandCloudBuildTriggerSubstitutions(d.Get("substitutions"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("substitutions"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, substitutionsProp)) {
		obj["substitutions"] = substitutionsProp
	}
	filenameProp, err := expandCloudBuildTriggerFilename(d.Get("filename"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("filename"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, filenameProp)) {
		obj["filename"] = filenameProp
	}
	ignoredFilesProp, err := expandCloudBuildTriggerIgnoredFiles(d.Get("ignored_files"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ignored_files"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, ignoredFilesProp)) {
		obj["ignoredFiles"] = ignoredFilesProp
	}
	includedFilesProp, err := expandCloudBuildTriggerIncludedFiles(d.Get("included_files"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("included_files"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, includedFilesProp)) {
		obj["includedFiles"] = includedFilesProp
	}
	triggerTemplateProp, err := expandCloudBuildTriggerTriggerTemplate(d.Get("trigger_template"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("trigger_template"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, triggerTemplateProp)) {
		obj["triggerTemplate"] = triggerTemplateProp
	}
	buildProp, err := expandCloudBuildTriggerBuild(d.Get("build"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("build"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, buildProp)) {
		obj["build"] = buildProp
	}

	url, err := replaceVars(d, config, "https://cloudbuild.googleapis.com/v1/projects/{{project}}/triggers/{{trigger_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Trigger %q: %#v", d.Id(), obj)
	obj["id"] = d.Get("trigger_id")
	_, err = sendRequestWithTimeout(config, "PATCH", url, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating Trigger %q: %s", d.Id(), err)
	}

	return resourceCloudBuildTriggerRead(d, meta)
}

func resourceCloudBuildTriggerDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "https://cloudbuild.googleapis.com/v1/projects/{{project}}/triggers/{{trigger_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Trigger %q", d.Id())
	res, err := sendRequestWithTimeout(config, "DELETE", url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "Trigger")
	}

	log.Printf("[DEBUG] Finished deleting Trigger %q: %#v", d.Id(), res)
	return nil
}

func resourceCloudBuildTriggerImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{"projects/(?P<project>[^/]+)/triggers/(?P<trigger_id>[^/]+)", "(?P<project>[^/]+)/(?P<trigger_id>[^/]+)", "(?P<trigger_id>[^/]+)"}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "{{project}}/{{trigger_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenCloudBuildTriggerTrigger_id(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudBuildTriggerDescription(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudBuildTriggerDisabled(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudBuildTriggerCreateTime(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudBuildTriggerSubstitutions(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudBuildTriggerFilename(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudBuildTriggerIgnoredFiles(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudBuildTriggerIncludedFiles(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudBuildTriggerTriggerTemplate(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["project_id"] =
		flattenCloudBuildTriggerTriggerTemplateProjectId(original["projectId"], d)
	transformed["repo_name"] =
		flattenCloudBuildTriggerTriggerTemplateRepoName(original["repoName"], d)
	transformed["dir"] =
		flattenCloudBuildTriggerTriggerTemplateDir(original["dir"], d)
	transformed["branch_name"] =
		flattenCloudBuildTriggerTriggerTemplateBranchName(original["branchName"], d)
	transformed["tag_name"] =
		flattenCloudBuildTriggerTriggerTemplateTagName(original["tagName"], d)
	transformed["commit_sha"] =
		flattenCloudBuildTriggerTriggerTemplateCommitSha(original["commitSha"], d)
	return []interface{}{transformed}
}
func flattenCloudBuildTriggerTriggerTemplateProjectId(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudBuildTriggerTriggerTemplateRepoName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudBuildTriggerTriggerTemplateDir(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudBuildTriggerTriggerTemplateBranchName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudBuildTriggerTriggerTemplateTagName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudBuildTriggerTriggerTemplateCommitSha(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudBuildTriggerBuild(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["tags"] =
		flattenCloudBuildTriggerBuildTags(original["tags"], d)
	transformed["images"] =
		flattenCloudBuildTriggerBuildImages(original["images"], d)
	transformed["step"] =
		flattenCloudBuildTriggerBuildStep(original["steps"], d)
	return []interface{}{transformed}
}
func flattenCloudBuildTriggerBuildTags(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudBuildTriggerBuildImages(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudBuildTriggerBuildStep(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"name":       flattenCloudBuildTriggerBuildStepName(original["name"], d),
			"args":       flattenCloudBuildTriggerBuildStepArgs(original["args"], d),
			"env":        flattenCloudBuildTriggerBuildStepEnv(original["env"], d),
			"id":         flattenCloudBuildTriggerBuildStepId(original["id"], d),
			"entrypoint": flattenCloudBuildTriggerBuildStepEntrypoint(original["entrypoint"], d),
			"dir":        flattenCloudBuildTriggerBuildStepDir(original["dir"], d),
			"secret_env": flattenCloudBuildTriggerBuildStepSecretEnv(original["secretEnv"], d),
			"timeout":    flattenCloudBuildTriggerBuildStepTimeout(original["timeout"], d),
			"timing":     flattenCloudBuildTriggerBuildStepTiming(original["timing"], d),
			"volumes":    flattenCloudBuildTriggerBuildStepVolumes(original["volumes"], d),
			"wait_for":   flattenCloudBuildTriggerBuildStepWaitFor(original["waitFor"], d),
		})
	}
	return transformed
}
func flattenCloudBuildTriggerBuildStepName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudBuildTriggerBuildStepArgs(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudBuildTriggerBuildStepEnv(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudBuildTriggerBuildStepId(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudBuildTriggerBuildStepEntrypoint(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudBuildTriggerBuildStepDir(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudBuildTriggerBuildStepSecretEnv(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudBuildTriggerBuildStepTimeout(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudBuildTriggerBuildStepTiming(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudBuildTriggerBuildStepVolumes(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"name": flattenCloudBuildTriggerBuildStepVolumesName(original["name"], d),
			"path": flattenCloudBuildTriggerBuildStepVolumesPath(original["path"], d),
		})
	}
	return transformed
}
func flattenCloudBuildTriggerBuildStepVolumesName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudBuildTriggerBuildStepVolumesPath(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenCloudBuildTriggerBuildStepWaitFor(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func expandCloudBuildTriggerDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerDisabled(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerSubstitutions(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandCloudBuildTriggerFilename(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerIgnoredFiles(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerIncludedFiles(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerTriggerTemplate(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedProjectId, err := expandCloudBuildTriggerTriggerTemplateProjectId(original["project_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProjectId); val.IsValid() && !isEmptyValue(val) {
		transformed["projectId"] = transformedProjectId
	}

	transformedRepoName, err := expandCloudBuildTriggerTriggerTemplateRepoName(original["repo_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRepoName); val.IsValid() && !isEmptyValue(val) {
		transformed["repoName"] = transformedRepoName
	}

	transformedDir, err := expandCloudBuildTriggerTriggerTemplateDir(original["dir"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDir); val.IsValid() && !isEmptyValue(val) {
		transformed["dir"] = transformedDir
	}

	transformedBranchName, err := expandCloudBuildTriggerTriggerTemplateBranchName(original["branch_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedBranchName); val.IsValid() && !isEmptyValue(val) {
		transformed["branchName"] = transformedBranchName
	}

	transformedTagName, err := expandCloudBuildTriggerTriggerTemplateTagName(original["tag_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTagName); val.IsValid() && !isEmptyValue(val) {
		transformed["tagName"] = transformedTagName
	}

	transformedCommitSha, err := expandCloudBuildTriggerTriggerTemplateCommitSha(original["commit_sha"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCommitSha); val.IsValid() && !isEmptyValue(val) {
		transformed["commitSha"] = transformedCommitSha
	}

	return transformed, nil
}

func expandCloudBuildTriggerTriggerTemplateProjectId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerTriggerTemplateRepoName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerTriggerTemplateDir(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerTriggerTemplateBranchName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerTriggerTemplateTagName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerTriggerTemplateCommitSha(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerBuild(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedTags, err := expandCloudBuildTriggerBuildTags(original["tags"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTags); val.IsValid() && !isEmptyValue(val) {
		transformed["tags"] = transformedTags
	}

	transformedImages, err := expandCloudBuildTriggerBuildImages(original["images"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedImages); val.IsValid() && !isEmptyValue(val) {
		transformed["images"] = transformedImages
	}

	transformedStep, err := expandCloudBuildTriggerBuildStep(original["step"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedStep); val.IsValid() && !isEmptyValue(val) {
		transformed["steps"] = transformedStep
	}

	return transformed, nil
}

func expandCloudBuildTriggerBuildTags(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerBuildImages(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerBuildStep(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandCloudBuildTriggerBuildStepName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !isEmptyValue(val) {
			transformed["name"] = transformedName
		}

		transformedArgs, err := expandCloudBuildTriggerBuildStepArgs(original["args"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedArgs); val.IsValid() && !isEmptyValue(val) {
			transformed["args"] = transformedArgs
		}

		transformedEnv, err := expandCloudBuildTriggerBuildStepEnv(original["env"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedEnv); val.IsValid() && !isEmptyValue(val) {
			transformed["env"] = transformedEnv
		}

		transformedId, err := expandCloudBuildTriggerBuildStepId(original["id"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedId); val.IsValid() && !isEmptyValue(val) {
			transformed["id"] = transformedId
		}

		transformedEntrypoint, err := expandCloudBuildTriggerBuildStepEntrypoint(original["entrypoint"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedEntrypoint); val.IsValid() && !isEmptyValue(val) {
			transformed["entrypoint"] = transformedEntrypoint
		}

		transformedDir, err := expandCloudBuildTriggerBuildStepDir(original["dir"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDir); val.IsValid() && !isEmptyValue(val) {
			transformed["dir"] = transformedDir
		}

		transformedSecretEnv, err := expandCloudBuildTriggerBuildStepSecretEnv(original["secret_env"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedSecretEnv); val.IsValid() && !isEmptyValue(val) {
			transformed["secretEnv"] = transformedSecretEnv
		}

		transformedTimeout, err := expandCloudBuildTriggerBuildStepTimeout(original["timeout"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedTimeout); val.IsValid() && !isEmptyValue(val) {
			transformed["timeout"] = transformedTimeout
		}

		transformedTiming, err := expandCloudBuildTriggerBuildStepTiming(original["timing"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedTiming); val.IsValid() && !isEmptyValue(val) {
			transformed["timing"] = transformedTiming
		}

		transformedVolumes, err := expandCloudBuildTriggerBuildStepVolumes(original["volumes"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedVolumes); val.IsValid() && !isEmptyValue(val) {
			transformed["volumes"] = transformedVolumes
		}

		transformedWaitFor, err := expandCloudBuildTriggerBuildStepWaitFor(original["wait_for"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedWaitFor); val.IsValid() && !isEmptyValue(val) {
			transformed["waitFor"] = transformedWaitFor
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandCloudBuildTriggerBuildStepName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerBuildStepArgs(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerBuildStepEnv(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerBuildStepId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerBuildStepEntrypoint(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerBuildStepDir(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerBuildStepSecretEnv(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerBuildStepTimeout(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerBuildStepTiming(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerBuildStepVolumes(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandCloudBuildTriggerBuildStepVolumesName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !isEmptyValue(val) {
			transformed["name"] = transformedName
		}

		transformedPath, err := expandCloudBuildTriggerBuildStepVolumesPath(original["path"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPath); val.IsValid() && !isEmptyValue(val) {
			transformed["path"] = transformedPath
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandCloudBuildTriggerBuildStepVolumesName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerBuildStepVolumesPath(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudBuildTriggerBuildStepWaitFor(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
