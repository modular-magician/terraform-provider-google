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

package google

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceBigqueryReservationReservationCapacityCommitment() *schema.Resource {
	return &schema.Resource{
		Create: resourceBigqueryReservationReservationCapacityCommitmentCreate,
		Read:   resourceBigqueryReservationReservationCapacityCommitmentRead,
		Update: resourceBigqueryReservationReservationCapacityCommitmentUpdate,
		Delete: resourceBigqueryReservationReservationCapacityCommitmentDelete,

		Importer: &schema.ResourceImporter{
			State: resourceBigqueryReservationReservationCapacityCommitmentImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"capacity_commitment_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The name of the capacity commitment. This field must only contain alphanumeric characters or dash.`,
			},
			"plan": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Capacity commitment commitment plan.`,
			},
			"slot_count": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: `Number of slots in this commitment.`,
			},
			"enforce_single_admin_project_per_org": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: `If true, fail the request if another project in the organization has a capacity commitment.`,
			},
			"location": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `The geographic location where the capacity commitment should reside.
Examples: US, EU, asia-northeast1.`,
			},
			"multi_region_auxiliary": {
				Type:     schema.TypeBool,
				Optional: true,
				Description: `Applicable only for commitments located within one of the BigQuery multi-regions (US or EU).
If set to true, this commitment is placed in the organization's secondary region which is designated for disaster recovery purposes. If false, this commitment is placed in the organization's default region.`,
			},
			"renewal_plan": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The plan this capacity commitment is converted to after commitmentEndTime passes. Once the plan is changed, committed period is extended according to commitment plan. Only applicable for ANNUAL and TRIAL commitments.`,
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

func resourceBigqueryReservationReservationCapacityCommitmentCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	slotCountProp, err := expandBigqueryReservationReservationCapacityCommitmentSlotCount(d.Get("slot_count"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("slot_count"); !isEmptyValue(reflect.ValueOf(slotCountProp)) && (ok || !reflect.DeepEqual(v, slotCountProp)) {
		obj["slotCount"] = slotCountProp
	}
	planProp, err := expandBigqueryReservationReservationCapacityCommitmentPlan(d.Get("plan"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("plan"); !isEmptyValue(reflect.ValueOf(planProp)) && (ok || !reflect.DeepEqual(v, planProp)) {
		obj["plan"] = planProp
	}
	renewalPlanProp, err := expandBigqueryReservationReservationCapacityCommitmentRenewalPlan(d.Get("renewal_plan"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("renewal_plan"); !isEmptyValue(reflect.ValueOf(renewalPlanProp)) && (ok || !reflect.DeepEqual(v, renewalPlanProp)) {
		obj["renewalPlan"] = renewalPlanProp
	}
	multiRegionAuxiliaryProp, err := expandBigqueryReservationReservationCapacityCommitmentMultiRegionAuxiliary(d.Get("multi_region_auxiliary"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("multi_region_auxiliary"); !isEmptyValue(reflect.ValueOf(multiRegionAuxiliaryProp)) && (ok || !reflect.DeepEqual(v, multiRegionAuxiliaryProp)) {
		obj["multiRegionAuxiliary"] = multiRegionAuxiliaryProp
	}

	url, err := replaceVars(d, config, "{{BigqueryReservationBasePath}}projects/{{project}}/locations/{{location}}/capacityCommitments?capacityCommitmentId={{capacityCommitmentId}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new ReservationCapacityCommitment: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ReservationCapacityCommitment: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating ReservationCapacityCommitment: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/locations/{{location}}/capacityCommitments/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating ReservationCapacityCommitment %q: %#v", d.Id(), res)

	return resourceBigqueryReservationReservationCapacityCommitmentRead(d, meta)
}

func resourceBigqueryReservationReservationCapacityCommitmentRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{BigqueryReservationBasePath}}projects/{{project}}/locations/{{location}}/capacityCommitments/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ReservationCapacityCommitment: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("BigqueryReservationReservationCapacityCommitment %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading ReservationCapacityCommitment: %s", err)
	}

	if err := d.Set("slot_count", flattenBigqueryReservationReservationCapacityCommitmentSlotCount(res["slotCount"], d, config)); err != nil {
		return fmt.Errorf("Error reading ReservationCapacityCommitment: %s", err)
	}
	if err := d.Set("plan", flattenBigqueryReservationReservationCapacityCommitmentPlan(res["plan"], d, config)); err != nil {
		return fmt.Errorf("Error reading ReservationCapacityCommitment: %s", err)
	}
	if err := d.Set("renewal_plan", flattenBigqueryReservationReservationCapacityCommitmentRenewalPlan(res["renewalPlan"], d, config)); err != nil {
		return fmt.Errorf("Error reading ReservationCapacityCommitment: %s", err)
	}
	if err := d.Set("multi_region_auxiliary", flattenBigqueryReservationReservationCapacityCommitmentMultiRegionAuxiliary(res["multiRegionAuxiliary"], d, config)); err != nil {
		return fmt.Errorf("Error reading ReservationCapacityCommitment: %s", err)
	}

	return nil
}

func resourceBigqueryReservationReservationCapacityCommitmentUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ReservationCapacityCommitment: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	slotCountProp, err := expandBigqueryReservationReservationCapacityCommitmentSlotCount(d.Get("slot_count"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("slot_count"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, slotCountProp)) {
		obj["slotCount"] = slotCountProp
	}
	planProp, err := expandBigqueryReservationReservationCapacityCommitmentPlan(d.Get("plan"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("plan"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, planProp)) {
		obj["plan"] = planProp
	}
	renewalPlanProp, err := expandBigqueryReservationReservationCapacityCommitmentRenewalPlan(d.Get("renewal_plan"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("renewal_plan"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, renewalPlanProp)) {
		obj["renewalPlan"] = renewalPlanProp
	}
	multiRegionAuxiliaryProp, err := expandBigqueryReservationReservationCapacityCommitmentMultiRegionAuxiliary(d.Get("multi_region_auxiliary"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("multi_region_auxiliary"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, multiRegionAuxiliaryProp)) {
		obj["multiRegionAuxiliary"] = multiRegionAuxiliaryProp
	}

	url, err := replaceVars(d, config, "{{BigqueryReservationBasePath}}projects/{{project}}/locations/{{location}}/capacityCommitments/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating ReservationCapacityCommitment %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("slot_count") {
		updateMask = append(updateMask, "slotCount")
	}

	if d.HasChange("plan") {
		updateMask = append(updateMask, "plan")
	}

	if d.HasChange("renewal_plan") {
		updateMask = append(updateMask, "renewalPlan")
	}

	if d.HasChange("multi_region_auxiliary") {
		updateMask = append(updateMask, "multiRegionAuxiliary")
	}
	// updateMask is a URL parameter but not present in the schema, so replaceVars
	// won't set it
	url, err = addQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating ReservationCapacityCommitment %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating ReservationCapacityCommitment %q: %#v", d.Id(), res)
	}

	return resourceBigqueryReservationReservationCapacityCommitmentRead(d, meta)
}

func resourceBigqueryReservationReservationCapacityCommitmentDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ReservationCapacityCommitment: %s", err)
	}
	billingProject = project

	url, err := replaceVars(d, config, "{{BigqueryReservationBasePath}}projects/{{project}}/locations/{{location}}/capacityCommitments/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting ReservationCapacityCommitment %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "ReservationCapacityCommitment")
	}

	log.Printf("[DEBUG] Finished deleting ReservationCapacityCommitment %q: %#v", d.Id(), res)
	return nil
}

func resourceBigqueryReservationReservationCapacityCommitmentImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/capacityCommitments/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<name>[^/]+)",
		"(?P<location>[^/]+)/(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/locations/{{location}}/capacityCommitments/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenBigqueryReservationReservationCapacityCommitmentSlotCount(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := stringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenBigqueryReservationReservationCapacityCommitmentPlan(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenBigqueryReservationReservationCapacityCommitmentRenewalPlan(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenBigqueryReservationReservationCapacityCommitmentMultiRegionAuxiliary(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandBigqueryReservationReservationCapacityCommitmentSlotCount(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryReservationReservationCapacityCommitmentPlan(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryReservationReservationCapacityCommitmentRenewalPlan(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryReservationReservationCapacityCommitmentMultiRegionAuxiliary(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
