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

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceLoggingMetric() *schema.Resource {
	return &schema.Resource{
		Create: resourceLoggingMetricCreate,
		Read:   resourceLoggingMetricRead,
		Update: resourceLoggingMetricUpdate,
		Delete: resourceLoggingMetricDelete,

		Importer: &schema.ResourceImporter{
			State: resourceLoggingMetricImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"filter": {
				Type:     schema.TypeString,
				Required: true,
				Description: `An advanced logs filter (https://cloud.google.com/logging/docs/view/advanced-filters) which
is used to match log entries.`,
			},
			"metric_descriptor": {
				Type:        schema.TypeList,
				Required:    true,
				Description: `The metric descriptor associated with the logs-based metric.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"metric_kind": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringInSlice([]string{"DELTA", "GAUGE", "CUMULATIVE"}, false),
							Description: `Whether the metric records instantaneous values, changes to a value, etc.
Some combinations of metricKind and valueType might not be supported.
For counter metrics, set this to DELTA.`,
						},
						"value_type": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringInSlice([]string{"BOOL", "INT64", "DOUBLE", "STRING", "DISTRIBUTION", "MONEY"}, false),
							Description: `Whether the measurement is an integer, a floating-point number, etc.
Some combinations of metricKind and valueType might not be supported.
For counter metrics, set this to INT64.`,
						},
						"display_name": {
							Type:     schema.TypeString,
							Optional: true,
							Description: `A concise name for the metric, which can be displayed in user interfaces. Use sentence case 
without an ending period, for example "Request count". This field is optional but it is 
recommended to be set for any metrics associated with user-visible concepts, such as Quota.`,
						},
						"labels": {
							Type:     schema.TypeSet,
							Optional: true,
							Description: `The set of labels that can be used to describe a specific instance of this metric type. For
example, the appengine.googleapis.com/http/server/response_latencies metric type has a label
for the HTTP response code, response_code, so you can look at latencies for successful responses
or just for responses that failed.`,
							Elem: loggingMetricMetricDescriptorLabelsSchema(),
							// Default schema.HashSchema is used.
						},
						"unit": {
							Type:     schema.TypeString,
							Optional: true,
							Description: `The unit in which the metric value is reported. It is only applicable if the valueType is
'INT64', 'DOUBLE', or 'DISTRIBUTION'. The supported units are a subset of
[The Unified Code for Units of Measure](http://unitsofmeasure.org/ucum.html) standard`,
							Default: "1",
						},
					},
				},
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				Description: `The client-assigned metric identifier. Examples - "error_count", "nginx/requests".
Metric identifiers are limited to 100 characters and can include only the following
characters A-Z, a-z, 0-9, and the special characters _-.,+!*',()%/. The forward-slash
character (/) denotes a hierarchy of name pieces, and it cannot be the first character
of the name.`,
			},
			"bucket_options": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `The bucketOptions are required when the logs-based metric is using a DISTRIBUTION value type and it
describes the bucket boundaries used to create a histogram of the extracted values.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"explicit_buckets": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `Specifies a set of buckets with arbitrary widths.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"bounds": {
										Type:        schema.TypeList,
										Required:    true,
										Description: `The values must be monotonically increasing.`,
										Elem: &schema.Schema{
											Type: schema.TypeFloat,
										},
									},
								},
							},
							AtLeastOneOf: []string{"bucket_options.0.linear_buckets", "bucket_options.0.exponential_buckets", "bucket_options.0.explicit_buckets"},
						},
						"exponential_buckets": {
							Type:     schema.TypeList,
							Optional: true,
							Description: `Specifies an exponential sequence of buckets that have a width that is proportional to the value of
the lower bound. Each bucket represents a constant relative uncertainty on a specific value in the bucket.`,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"growth_factor": {
										Type:         schema.TypeFloat,
										Optional:     true,
										Description:  `Must be greater than 1.`,
										AtLeastOneOf: []string{"bucket_options.0.exponential_buckets.0.num_finite_buckets", "bucket_options.0.exponential_buckets.0.growth_factor", "bucket_options.0.exponential_buckets.0.scale"},
									},
									"num_finite_buckets": {
										Type:         schema.TypeInt,
										Optional:     true,
										Description:  `Must be greater than 0.`,
										AtLeastOneOf: []string{"bucket_options.0.exponential_buckets.0.num_finite_buckets", "bucket_options.0.exponential_buckets.0.growth_factor", "bucket_options.0.exponential_buckets.0.scale"},
									},
									"scale": {
										Type:         schema.TypeFloat,
										Optional:     true,
										Description:  `Must be greater than 0.`,
										AtLeastOneOf: []string{"bucket_options.0.exponential_buckets.0.num_finite_buckets", "bucket_options.0.exponential_buckets.0.growth_factor", "bucket_options.0.exponential_buckets.0.scale"},
									},
								},
							},
							AtLeastOneOf: []string{"bucket_options.0.linear_buckets", "bucket_options.0.exponential_buckets", "bucket_options.0.explicit_buckets"},
						},
						"linear_buckets": {
							Type:     schema.TypeList,
							Optional: true,
							Description: `Specifies a linear sequence of buckets that all have the same width (except overflow and underflow).
Each bucket represents a constant absolute uncertainty on the specific value in the bucket.`,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"num_finite_buckets": {
										Type:         schema.TypeInt,
										Optional:     true,
										Description:  `Must be greater than 0.`,
										AtLeastOneOf: []string{"bucket_options.0.linear_buckets.0.num_finite_buckets", "bucket_options.0.linear_buckets.0.width", "bucket_options.0.linear_buckets.0.offset"},
									},
									"offset": {
										Type:         schema.TypeFloat,
										Optional:     true,
										Description:  `Lower bound of the first bucket.`,
										AtLeastOneOf: []string{"bucket_options.0.linear_buckets.0.num_finite_buckets", "bucket_options.0.linear_buckets.0.width", "bucket_options.0.linear_buckets.0.offset"},
									},
									"width": {
										Type:         schema.TypeInt,
										Optional:     true,
										Description:  `Must be greater than 0.`,
										AtLeastOneOf: []string{"bucket_options.0.linear_buckets.0.num_finite_buckets", "bucket_options.0.linear_buckets.0.width", "bucket_options.0.linear_buckets.0.offset"},
									},
								},
							},
							AtLeastOneOf: []string{"bucket_options.0.linear_buckets", "bucket_options.0.exponential_buckets", "bucket_options.0.explicit_buckets"},
						},
					},
				},
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `A description of this metric, which is used in documentation. The maximum length of the
description is 8000 characters.`,
			},
			"label_extractors": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `A map from a label key string to an extractor expression which is used to extract data from a log
entry field and assign as the label value. Each label key specified in the LabelDescriptor must
have an associated extractor expression in this map. The syntax of the extractor expression is
the same as for the valueExtractor field.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"value_extractor": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `A valueExtractor is required when using a distribution logs-based metric to extract the values to
record from a log entry. Two functions are supported for value extraction - EXTRACT(field) or
REGEXP_EXTRACT(field, regex). The argument are 1. field - The name of the log entry field from which
the value is to be extracted. 2. regex - A regular expression using the Google RE2 syntax
(https://github.com/google/re2/wiki/Syntax) with a single capture group to extract data from the specified
log entry field. The value of the field is converted to a string before applying the regex. It is an
error to specify a regex that does not include exactly one capture group.`,
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

func loggingMetricMetricDescriptorLabelsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"key": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The label key.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `A human-readable description for the label.`,
			},
			"value_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"BOOL", "INT64", "STRING", ""}, false),
				Description:  `The type of data that can be assigned to the label.`,
				Default:      "STRING",
			},
		},
	}
}

func resourceLoggingMetricCreate(d *schema.ResourceData, meta interface{}) error {

	config := meta.(*Config)

	obj := make(map[string]interface{})
	nameProp, err := expandLoggingMetricName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandLoggingMetricDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	filterProp, err := expandLoggingMetricFilter(d.Get("filter"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("filter"); !isEmptyValue(reflect.ValueOf(filterProp)) && (ok || !reflect.DeepEqual(v, filterProp)) {
		obj["filter"] = filterProp
	}
	metricDescriptorProp, err := expandLoggingMetricMetricDescriptor(d.Get("metric_descriptor"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("metric_descriptor"); !isEmptyValue(reflect.ValueOf(metricDescriptorProp)) && (ok || !reflect.DeepEqual(v, metricDescriptorProp)) {
		obj["metricDescriptor"] = metricDescriptorProp
	}
	labelExtractorsProp, err := expandLoggingMetricLabelExtractors(d.Get("label_extractors"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("label_extractors"); !isEmptyValue(reflect.ValueOf(labelExtractorsProp)) && (ok || !reflect.DeepEqual(v, labelExtractorsProp)) {
		obj["labelExtractors"] = labelExtractorsProp
	}
	valueExtractorProp, err := expandLoggingMetricValueExtractor(d.Get("value_extractor"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("value_extractor"); !isEmptyValue(reflect.ValueOf(valueExtractorProp)) && (ok || !reflect.DeepEqual(v, valueExtractorProp)) {
		obj["valueExtractor"] = valueExtractorProp
	}
	bucketOptionsProp, err := expandLoggingMetricBucketOptions(d.Get("bucket_options"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("bucket_options"); !isEmptyValue(reflect.ValueOf(bucketOptionsProp)) && (ok || !reflect.DeepEqual(v, bucketOptionsProp)) {
		obj["bucketOptions"] = bucketOptionsProp
	}

	lockName, err := replaceVars(d, config, "customMetric/{{project}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := replaceVars(d, config, "{{LoggingBasePath}}projects/{{project}}/metrics")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Metric: %#v", obj)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "POST", project, url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Metric: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Metric %q: %#v", d.Id(), res)

	// `name` is autogenerated from the api so needs to be set post-create
	name, ok := res["name"]
	if !ok {
		return fmt.Errorf("Create response didn't contain critical fields. Create may not have succeeded.")
	}
	d.Set("name", name.(string))
	d.SetId(name.(string))

	return resourceLoggingMetricRead(d, meta)
}

func resourceLoggingMetricRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{LoggingBasePath}}projects/{{project}}/metrics/{{%name}}")
	if err != nil {
		return err
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequest(config, "GET", project, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("LoggingMetric %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Metric: %s", err)
	}

	if err := d.Set("name", flattenLoggingMetricName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Metric: %s", err)
	}
	if err := d.Set("description", flattenLoggingMetricDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading Metric: %s", err)
	}
	if err := d.Set("filter", flattenLoggingMetricFilter(res["filter"], d, config)); err != nil {
		return fmt.Errorf("Error reading Metric: %s", err)
	}
	if err := d.Set("metric_descriptor", flattenLoggingMetricMetricDescriptor(res["metricDescriptor"], d, config)); err != nil {
		return fmt.Errorf("Error reading Metric: %s", err)
	}
	if err := d.Set("label_extractors", flattenLoggingMetricLabelExtractors(res["labelExtractors"], d, config)); err != nil {
		return fmt.Errorf("Error reading Metric: %s", err)
	}
	if err := d.Set("value_extractor", flattenLoggingMetricValueExtractor(res["valueExtractor"], d, config)); err != nil {
		return fmt.Errorf("Error reading Metric: %s", err)
	}
	if err := d.Set("bucket_options", flattenLoggingMetricBucketOptions(res["bucketOptions"], d, config)); err != nil {
		return fmt.Errorf("Error reading Metric: %s", err)
	}

	return nil
}

func resourceLoggingMetricUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandLoggingMetricName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandLoggingMetricDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	filterProp, err := expandLoggingMetricFilter(d.Get("filter"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("filter"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, filterProp)) {
		obj["filter"] = filterProp
	}
	metricDescriptorProp, err := expandLoggingMetricMetricDescriptor(d.Get("metric_descriptor"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("metric_descriptor"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, metricDescriptorProp)) {
		obj["metricDescriptor"] = metricDescriptorProp
	}
	labelExtractorsProp, err := expandLoggingMetricLabelExtractors(d.Get("label_extractors"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("label_extractors"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelExtractorsProp)) {
		obj["labelExtractors"] = labelExtractorsProp
	}
	valueExtractorProp, err := expandLoggingMetricValueExtractor(d.Get("value_extractor"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("value_extractor"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, valueExtractorProp)) {
		obj["valueExtractor"] = valueExtractorProp
	}
	bucketOptionsProp, err := expandLoggingMetricBucketOptions(d.Get("bucket_options"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("bucket_options"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, bucketOptionsProp)) {
		obj["bucketOptions"] = bucketOptionsProp
	}

	lockName, err := replaceVars(d, config, "customMetric/{{project}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := replaceVars(d, config, "{{LoggingBasePath}}projects/{{project}}/metrics/{{%name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Metric %q: %#v", d.Id(), obj)
	_, err = sendRequestWithTimeout(config, "PUT", project, url, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating Metric %q: %s", d.Id(), err)
	}

	return resourceLoggingMetricRead(d, meta)
}

func resourceLoggingMetricDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	lockName, err := replaceVars(d, config, "customMetric/{{project}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := replaceVars(d, config, "{{LoggingBasePath}}projects/{{project}}/metrics/{{%name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Metric %q", d.Id())

	res, err := sendRequestWithTimeout(config, "DELETE", project, url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "Metric")
	}

	log.Printf("[DEBUG] Finished deleting Metric %q: %#v", d.Id(), res)
	return nil
}

func resourceLoggingMetricImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	config := meta.(*Config)

	// current import_formats can't import fields with forward slashes in their value
	if err := parseImportId([]string{"(?P<name>.+)"}, d, config); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}

func flattenLoggingMetricName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenLoggingMetricDescription(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenLoggingMetricFilter(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenLoggingMetricMetricDescriptor(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["unit"] =
		flattenLoggingMetricMetricDescriptorUnit(original["unit"], d, config)
	transformed["value_type"] =
		flattenLoggingMetricMetricDescriptorValueType(original["valueType"], d, config)
	transformed["metric_kind"] =
		flattenLoggingMetricMetricDescriptorMetricKind(original["metricKind"], d, config)
	transformed["labels"] =
		flattenLoggingMetricMetricDescriptorLabels(original["labels"], d, config)
	transformed["display_name"] =
		flattenLoggingMetricMetricDescriptorDisplayName(original["displayName"], d, config)
	return []interface{}{transformed}
}
func flattenLoggingMetricMetricDescriptorUnit(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenLoggingMetricMetricDescriptorValueType(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenLoggingMetricMetricDescriptorMetricKind(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenLoggingMetricMetricDescriptorLabels(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := schema.NewSet(schema.HashResource(loggingMetricMetricDescriptorLabelsSchema()), []interface{}{})
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed.Add(map[string]interface{}{
			"key":         flattenLoggingMetricMetricDescriptorLabelsKey(original["key"], d, config),
			"description": flattenLoggingMetricMetricDescriptorLabelsDescription(original["description"], d, config),
			"value_type":  flattenLoggingMetricMetricDescriptorLabelsValueType(original["valueType"], d, config),
		})
	}
	return transformed
}
func flattenLoggingMetricMetricDescriptorLabelsKey(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenLoggingMetricMetricDescriptorLabelsDescription(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenLoggingMetricMetricDescriptorLabelsValueType(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil || isEmptyValue(reflect.ValueOf(v)) {
		return "STRING"
	}

	return v
}

func flattenLoggingMetricMetricDescriptorDisplayName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenLoggingMetricLabelExtractors(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenLoggingMetricValueExtractor(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenLoggingMetricBucketOptions(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["linear_buckets"] =
		flattenLoggingMetricBucketOptionsLinearBuckets(original["linearBuckets"], d, config)
	transformed["exponential_buckets"] =
		flattenLoggingMetricBucketOptionsExponentialBuckets(original["exponentialBuckets"], d, config)
	transformed["explicit_buckets"] =
		flattenLoggingMetricBucketOptionsExplicitBuckets(original["explicitBuckets"], d, config)
	return []interface{}{transformed}
}
func flattenLoggingMetricBucketOptionsLinearBuckets(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["num_finite_buckets"] =
		flattenLoggingMetricBucketOptionsLinearBucketsNumFiniteBuckets(original["numFiniteBuckets"], d, config)
	transformed["width"] =
		flattenLoggingMetricBucketOptionsLinearBucketsWidth(original["width"], d, config)
	transformed["offset"] =
		flattenLoggingMetricBucketOptionsLinearBucketsOffset(original["offset"], d, config)
	return []interface{}{transformed}
}
func flattenLoggingMetricBucketOptionsLinearBucketsNumFiniteBuckets(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenLoggingMetricBucketOptionsLinearBucketsWidth(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenLoggingMetricBucketOptionsLinearBucketsOffset(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenLoggingMetricBucketOptionsExponentialBuckets(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["num_finite_buckets"] =
		flattenLoggingMetricBucketOptionsExponentialBucketsNumFiniteBuckets(original["numFiniteBuckets"], d, config)
	transformed["growth_factor"] =
		flattenLoggingMetricBucketOptionsExponentialBucketsGrowthFactor(original["growthFactor"], d, config)
	transformed["scale"] =
		flattenLoggingMetricBucketOptionsExponentialBucketsScale(original["scale"], d, config)
	return []interface{}{transformed}
}
func flattenLoggingMetricBucketOptionsExponentialBucketsNumFiniteBuckets(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenLoggingMetricBucketOptionsExponentialBucketsGrowthFactor(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenLoggingMetricBucketOptionsExponentialBucketsScale(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenLoggingMetricBucketOptionsExplicitBuckets(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["bounds"] =
		flattenLoggingMetricBucketOptionsExplicitBucketsBounds(original["bounds"], d, config)
	return []interface{}{transformed}
}
func flattenLoggingMetricBucketOptionsExplicitBucketsBounds(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandLoggingMetricName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandLoggingMetricDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandLoggingMetricFilter(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandLoggingMetricMetricDescriptor(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedUnit, err := expandLoggingMetricMetricDescriptorUnit(original["unit"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUnit); val.IsValid() && !isEmptyValue(val) {
		transformed["unit"] = transformedUnit
	}

	transformedValueType, err := expandLoggingMetricMetricDescriptorValueType(original["value_type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedValueType); val.IsValid() && !isEmptyValue(val) {
		transformed["valueType"] = transformedValueType
	}

	transformedMetricKind, err := expandLoggingMetricMetricDescriptorMetricKind(original["metric_kind"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMetricKind); val.IsValid() && !isEmptyValue(val) {
		transformed["metricKind"] = transformedMetricKind
	}

	transformedLabels, err := expandLoggingMetricMetricDescriptorLabels(original["labels"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLabels); val.IsValid() && !isEmptyValue(val) {
		transformed["labels"] = transformedLabels
	}

	transformedDisplayName, err := expandLoggingMetricMetricDescriptorDisplayName(original["display_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDisplayName); val.IsValid() && !isEmptyValue(val) {
		transformed["displayName"] = transformedDisplayName
	}

	return transformed, nil
}

func expandLoggingMetricMetricDescriptorUnit(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandLoggingMetricMetricDescriptorValueType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandLoggingMetricMetricDescriptorMetricKind(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandLoggingMetricMetricDescriptorLabels(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedKey, err := expandLoggingMetricMetricDescriptorLabelsKey(original["key"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedKey); val.IsValid() && !isEmptyValue(val) {
			transformed["key"] = transformedKey
		}

		transformedDescription, err := expandLoggingMetricMetricDescriptorLabelsDescription(original["description"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !isEmptyValue(val) {
			transformed["description"] = transformedDescription
		}

		transformedValueType, err := expandLoggingMetricMetricDescriptorLabelsValueType(original["value_type"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedValueType); val.IsValid() && !isEmptyValue(val) {
			transformed["valueType"] = transformedValueType
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandLoggingMetricMetricDescriptorLabelsKey(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandLoggingMetricMetricDescriptorLabelsDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandLoggingMetricMetricDescriptorLabelsValueType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandLoggingMetricMetricDescriptorDisplayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandLoggingMetricLabelExtractors(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandLoggingMetricValueExtractor(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandLoggingMetricBucketOptions(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedLinearBuckets, err := expandLoggingMetricBucketOptionsLinearBuckets(original["linear_buckets"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLinearBuckets); val.IsValid() && !isEmptyValue(val) {
		transformed["linearBuckets"] = transformedLinearBuckets
	}

	transformedExponentialBuckets, err := expandLoggingMetricBucketOptionsExponentialBuckets(original["exponential_buckets"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedExponentialBuckets); val.IsValid() && !isEmptyValue(val) {
		transformed["exponentialBuckets"] = transformedExponentialBuckets
	}

	transformedExplicitBuckets, err := expandLoggingMetricBucketOptionsExplicitBuckets(original["explicit_buckets"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedExplicitBuckets); val.IsValid() && !isEmptyValue(val) {
		transformed["explicitBuckets"] = transformedExplicitBuckets
	}

	return transformed, nil
}

func expandLoggingMetricBucketOptionsLinearBuckets(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedNumFiniteBuckets, err := expandLoggingMetricBucketOptionsLinearBucketsNumFiniteBuckets(original["num_finite_buckets"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNumFiniteBuckets); val.IsValid() && !isEmptyValue(val) {
		transformed["numFiniteBuckets"] = transformedNumFiniteBuckets
	}

	transformedWidth, err := expandLoggingMetricBucketOptionsLinearBucketsWidth(original["width"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedWidth); val.IsValid() && !isEmptyValue(val) {
		transformed["width"] = transformedWidth
	}

	transformedOffset, err := expandLoggingMetricBucketOptionsLinearBucketsOffset(original["offset"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedOffset); val.IsValid() && !isEmptyValue(val) {
		transformed["offset"] = transformedOffset
	}

	return transformed, nil
}

func expandLoggingMetricBucketOptionsLinearBucketsNumFiniteBuckets(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandLoggingMetricBucketOptionsLinearBucketsWidth(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandLoggingMetricBucketOptionsLinearBucketsOffset(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandLoggingMetricBucketOptionsExponentialBuckets(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedNumFiniteBuckets, err := expandLoggingMetricBucketOptionsExponentialBucketsNumFiniteBuckets(original["num_finite_buckets"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNumFiniteBuckets); val.IsValid() && !isEmptyValue(val) {
		transformed["numFiniteBuckets"] = transformedNumFiniteBuckets
	}

	transformedGrowthFactor, err := expandLoggingMetricBucketOptionsExponentialBucketsGrowthFactor(original["growth_factor"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGrowthFactor); val.IsValid() && !isEmptyValue(val) {
		transformed["growthFactor"] = transformedGrowthFactor
	}

	transformedScale, err := expandLoggingMetricBucketOptionsExponentialBucketsScale(original["scale"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedScale); val.IsValid() && !isEmptyValue(val) {
		transformed["scale"] = transformedScale
	}

	return transformed, nil
}

func expandLoggingMetricBucketOptionsExponentialBucketsNumFiniteBuckets(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandLoggingMetricBucketOptionsExponentialBucketsGrowthFactor(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandLoggingMetricBucketOptionsExponentialBucketsScale(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandLoggingMetricBucketOptionsExplicitBuckets(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedBounds, err := expandLoggingMetricBucketOptionsExplicitBucketsBounds(original["bounds"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedBounds); val.IsValid() && !isEmptyValue(val) {
		transformed["bounds"] = transformedBounds
	}

	return transformed, nil
}

func expandLoggingMetricBucketOptionsExplicitBucketsBounds(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
