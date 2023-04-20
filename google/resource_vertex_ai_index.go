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

func ResourceVertexAIIndex() *schema.Resource {
	return &schema.Resource{
		Create: resourceVertexAIIndexCreate,
		Read:   resourceVertexAIIndexRead,
		Update: resourceVertexAIIndexUpdate,
		Delete: resourceVertexAIIndexDelete,

		Importer: &schema.ResourceImporter{
			State: resourceVertexAIIndexImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(180 * time.Minute),
			Update: schema.DefaultTimeout(180 * time.Minute),
			Delete: schema.DefaultTimeout(180 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The display name of the Index. The name can be up to 128 characters long and can consist of any UTF-8 characters.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The description of the Index.`,
			},
			"index_update_method": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `The update method to use with this Index. The value must be the followings. If not set, BATCH_UPDATE will be used by default.
* BATCH_UPDATE: user can call indexes.patch with files on Cloud Storage of datapoints to update.
* STREAM_UPDATE: user can call indexes.upsertDatapoints/DeleteDatapoints to update the Index and the updates will be applied in corresponding DeployedIndexes in nearly real-time.`,
				Default: "BATCH_UPDATE",
			},
			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: `The labels with user-defined metadata to organize your Indexes.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"metadata": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `An additional information about the Index`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"config": {
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Description: `The configuration of the Matching Engine Index.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dimensions": {
										Type:        schema.TypeInt,
										Required:    true,
										Description: `The number of dimensions of the input vectors.`,
									},
									"algorithm_config": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: `The configuration with regard to the algorithms used for efficient search.`,
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"brute_force_config": {
													Type:     schema.TypeList,
													Optional: true,
													Description: `Configuration options for using brute force search, which simply implements the
standard linear search in the database for each query.`,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{},
													},
													ExactlyOneOf: []string{},
												},
												"tree_ah_config": {
													Type:     schema.TypeList,
													Optional: true,
													Description: `Configuration options for using the tree-AH algorithm (Shallow tree + Asymmetric Hashing).
Please refer to this paper for more details: https://arxiv.org/abs/1908.10396`,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"leaf_node_embedding_count": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: `Number of embeddings on each leaf node. The default value is 1000 if not set.`,
																Default:     1000,
															},
															"leaf_nodes_to_search_percent": {
																Type:     schema.TypeInt,
																Optional: true,
																Description: `The default percentage of leaf nodes that any query may be searched. Must be in
range 1-100, inclusive. The default value is 10 (means 10%) if not set.`,
																Default: 10,
															},
														},
													},
													ExactlyOneOf: []string{},
												},
											},
										},
									},
									"approximate_neighbors_count": {
										Type:     schema.TypeInt,
										Optional: true,
										Description: `The default number of neighbors to find via approximate search before exact reordering is
performed. Exact reordering is a procedure where results returned by an
approximate search algorithm are reordered via a more expensive distance computation.
Required if tree-AH algorithm is used.`,
									},
									"distance_measure_type": {
										Type:     schema.TypeString,
										Optional: true,
										Description: `The distance measure used in nearest neighbor search. The value must be one of the followings:
* SQUARED_L2_DISTANCE: Euclidean (L_2) Distance
* L1_DISTANCE: Manhattan (L_1) Distance
* COSINE_DISTANCE: Cosine Distance. Defined as 1 - cosine similarity.
* DOT_PRODUCT_DISTANCE: Dot Product Distance. Defined as a negative of the dot product`,
										Default: "DOT_PRODUCT_DISTANCE",
									},
									"feature_norm_type": {
										Type:     schema.TypeString,
										Optional: true,
										Description: `Type of normalization to be carried out on each vector. The value must be one of the followings:
* UNIT_L2_NORM: Unit L2 normalization type
* NONE: No normalization type is specified.`,
										Default: "NONE",
									},
								},
							},
						},
						"contents_delta_uri": {
							Type:     schema.TypeString,
							Optional: true,
							Description: `Allows inserting, updating  or deleting the contents of the Matching Engine Index.
The string must be a valid Cloud Storage directory path. If this
field is set when calling IndexService.UpdateIndex, then no other
Index field can be also updated as part of the same call.
The expected structure and format of the files this URI points to is
described at https://cloud.google.com/vertex-ai/docs/matching-engine/using-matching-engine#input-data-format`,
						},
						"is_complete_overwrite": {
							Type:     schema.TypeBool,
							Optional: true,
							Description: `If this field is set together with contentsDeltaUri when calling IndexService.UpdateIndex,
then existing content of the Index will be replaced by the data from the contentsDeltaUri.`,
							Default: false,
						},
					},
				},
			},
			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `The region of the index. eg us-central1`,
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The timestamp of when the Index was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.`,
			},
			"deployed_indexes": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `The pointers to DeployedIndexes created from this Index. An Index can be only deleted if all its DeployedIndexes had been undeployed first.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"deployed_index_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The ID of the DeployedIndex in the above IndexEndpoint.`,
						},
						"index_endpoint": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `A resource name of the IndexEndpoint.`,
						},
					},
				},
			},
			"etag": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Used to perform consistent read-modify-write updates.`,
			},
			"index_stats": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `Stats of the index resource.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"shards_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: `The number of shards in the Index.`,
						},
						"vectors_count": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The number of vectors in the Index.`,
						},
					},
				},
			},
			"metadata_schema_uri": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Points to a YAML file stored on Google Cloud Storage describing additional information about the Index, that is specific to it. Unset if the Index does not have any additional information.`,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The resource name of the Index.`,
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The timestamp of when the Index was last updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.`,
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

func resourceVertexAIIndexCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	displayNameProp, err := expandVertexAIIndexDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	descriptionProp, err := expandVertexAIIndexDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	metadataProp, err := expandVertexAIIndexMetadata(d.Get("metadata"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("metadata"); !isEmptyValue(reflect.ValueOf(metadataProp)) && (ok || !reflect.DeepEqual(v, metadataProp)) {
		obj["metadata"] = metadataProp
	}
	labelsProp, err := expandVertexAIIndexLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	indexUpdateMethodProp, err := expandVertexAIIndexIndexUpdateMethod(d.Get("index_update_method"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("index_update_method"); !isEmptyValue(reflect.ValueOf(indexUpdateMethodProp)) && (ok || !reflect.DeepEqual(v, indexUpdateMethodProp)) {
		obj["indexUpdateMethod"] = indexUpdateMethodProp
	}

	url, err := ReplaceVars(d, config, "{{VertexAIBasePath}}projects/{{project}}/locations/{{region}}/indexes")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Index: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Index: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Index: %s", err)
	}

	// Store the ID now
	id, err := ReplaceVars(d, config, "projects/{{project}}/locations/{{region}}/indexes/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = VertexAIOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating Index", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create Index: %s", err)
	}

	if err := d.Set("name", flattenVertexAIIndexName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = ReplaceVars(d, config, "projects/{{project}}/locations/{{region}}/indexes/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Index %q: %#v", d.Id(), res)

	return resourceVertexAIIndexRead(d, meta)
}

func resourceVertexAIIndexRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{VertexAIBasePath}}projects/{{project}}/locations/{{region}}/indexes/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Index: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("VertexAIIndex %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Index: %s", err)
	}

	if err := d.Set("name", flattenVertexAIIndexName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Index: %s", err)
	}
	if err := d.Set("display_name", flattenVertexAIIndexDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading Index: %s", err)
	}
	if err := d.Set("description", flattenVertexAIIndexDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading Index: %s", err)
	}
	if err := d.Set("metadata", flattenVertexAIIndexMetadata(res["metadata"], d, config)); err != nil {
		return fmt.Errorf("Error reading Index: %s", err)
	}
	if err := d.Set("metadata_schema_uri", flattenVertexAIIndexMetadataSchemaUri(res["metadataSchemaUri"], d, config)); err != nil {
		return fmt.Errorf("Error reading Index: %s", err)
	}
	if err := d.Set("deployed_indexes", flattenVertexAIIndexDeployedIndexes(res["deployedIndexes"], d, config)); err != nil {
		return fmt.Errorf("Error reading Index: %s", err)
	}
	if err := d.Set("labels", flattenVertexAIIndexLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Index: %s", err)
	}
	if err := d.Set("create_time", flattenVertexAIIndexCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Index: %s", err)
	}
	if err := d.Set("update_time", flattenVertexAIIndexUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Index: %s", err)
	}
	if err := d.Set("index_stats", flattenVertexAIIndexIndexStats(res["indexStats"], d, config)); err != nil {
		return fmt.Errorf("Error reading Index: %s", err)
	}
	if err := d.Set("index_update_method", flattenVertexAIIndexIndexUpdateMethod(res["indexUpdateMethod"], d, config)); err != nil {
		return fmt.Errorf("Error reading Index: %s", err)
	}

	return nil
}

func resourceVertexAIIndexUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Index: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	displayNameProp, err := expandVertexAIIndexDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	descriptionProp, err := expandVertexAIIndexDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	metadataProp, err := expandVertexAIIndexMetadata(d.Get("metadata"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("metadata"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, metadataProp)) {
		obj["metadata"] = metadataProp
	}
	labelsProp, err := expandVertexAIIndexLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := ReplaceVars(d, config, "{{VertexAIBasePath}}projects/{{project}}/locations/{{region}}/indexes/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Index %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
	}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}

	if d.HasChange("metadata") {
		updateMask = append(updateMask, "metadata")
	}

	if d.HasChange("labels") {
		updateMask = append(updateMask, "labels")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}
	newUpdateMask := []string{}

	if d.HasChange("metadata.0.contents_delta_uri") {
		// Use the current value of isCompleteOverwrite when updating contentsDeltaUri
		newUpdateMask = append(newUpdateMask, "metadata.contentsDeltaUri")
		newUpdateMask = append(newUpdateMask, "metadata.isCompleteOverwrite")
	}

	for _, mask := range updateMask {
		// Use granular update masks instead of 'metadata' to avoid the following error:
		// 'If `contents_delta_gcs_uri` is set as part of `index.metadata`, then no other Index fields can be also updated as part of the same update call.'
		if mask == "metadata" {
			continue
		}
		newUpdateMask = append(newUpdateMask, mask)
	}

	// Refreshing updateMask after adding extra schema entries
	url, err = AddQueryParams(url, map[string]string{"updateMask": strings.Join(newUpdateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating Index %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Index %q: %#v", d.Id(), res)
	}

	err = VertexAIOperationWaitTime(
		config, res, project, "Updating Index", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceVertexAIIndexRead(d, meta)
}

func resourceVertexAIIndexDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Index: %s", err)
	}
	billingProject = project

	url, err := ReplaceVars(d, config, "{{VertexAIBasePath}}projects/{{project}}/locations/{{region}}/indexes/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Index %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "Index")
	}

	err = VertexAIOperationWaitTime(
		config, res, project, "Deleting Index", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Index %q: %#v", d.Id(), res)
	return nil
}

func resourceVertexAIIndexImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := ParseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<region>[^/]+)/indexes/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := ReplaceVars(d, config, "projects/{{project}}/locations/{{region}}/indexes/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenVertexAIIndexName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return NameFromSelfLinkStateFunc(v)
}

func flattenVertexAIIndexDisplayName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenVertexAIIndexDescription(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenVertexAIIndexMetadata(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["contents_delta_uri"] =
		flattenVertexAIIndexMetadataContentsDeltaUri(original["contentsDeltaUri"], d, config)
	transformed["is_complete_overwrite"] =
		flattenVertexAIIndexMetadataIsCompleteOverwrite(original["isCompleteOverwrite"], d, config)
	transformed["config"] =
		flattenVertexAIIndexMetadataConfig(original["config"], d, config)
	return []interface{}{transformed}
}
func flattenVertexAIIndexMetadataContentsDeltaUri(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// We want to ignore read on this field, but cannot because it is nested
	return d.Get("metadata.0.contents_delta_uri")
}

func flattenVertexAIIndexMetadataIsCompleteOverwrite(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// We want to ignore read on this field, but cannot because it is nested
	return d.Get("metadata.0.is_complete_overwrite")
}

func flattenVertexAIIndexMetadataConfig(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["dimensions"] =
		flattenVertexAIIndexMetadataConfigDimensions(original["dimensions"], d, config)
	transformed["approximate_neighbors_count"] =
		flattenVertexAIIndexMetadataConfigApproximateNeighborsCount(original["approximateNeighborsCount"], d, config)
	transformed["distance_measure_type"] =
		flattenVertexAIIndexMetadataConfigDistanceMeasureType(original["distanceMeasureType"], d, config)
	transformed["feature_norm_type"] =
		flattenVertexAIIndexMetadataConfigFeatureNormType(original["featureNormType"], d, config)
	transformed["algorithm_config"] =
		flattenVertexAIIndexMetadataConfigAlgorithmConfig(original["algorithmConfig"], d, config)
	return []interface{}{transformed}
}
func flattenVertexAIIndexMetadataConfigDimensions(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := StringToFixed64(strVal); err == nil {
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

func flattenVertexAIIndexMetadataConfigApproximateNeighborsCount(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := StringToFixed64(strVal); err == nil {
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

func flattenVertexAIIndexMetadataConfigDistanceMeasureType(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenVertexAIIndexMetadataConfigFeatureNormType(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenVertexAIIndexMetadataConfigAlgorithmConfig(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["tree_ah_config"] =
		flattenVertexAIIndexMetadataConfigAlgorithmConfigTreeAhConfig(original["treeAhConfig"], d, config)
	transformed["brute_force_config"] =
		flattenVertexAIIndexMetadataConfigAlgorithmConfigBruteForceConfig(original["bruteForceConfig"], d, config)
	return []interface{}{transformed}
}
func flattenVertexAIIndexMetadataConfigAlgorithmConfigTreeAhConfig(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["leaf_node_embedding_count"] =
		flattenVertexAIIndexMetadataConfigAlgorithmConfigTreeAhConfigLeafNodeEmbeddingCount(original["leafNodeEmbeddingCount"], d, config)
	transformed["leaf_nodes_to_search_percent"] =
		flattenVertexAIIndexMetadataConfigAlgorithmConfigTreeAhConfigLeafNodesToSearchPercent(original["leafNodesToSearchPercent"], d, config)
	return []interface{}{transformed}
}
func flattenVertexAIIndexMetadataConfigAlgorithmConfigTreeAhConfigLeafNodeEmbeddingCount(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := StringToFixed64(strVal); err == nil {
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

func flattenVertexAIIndexMetadataConfigAlgorithmConfigTreeAhConfigLeafNodesToSearchPercent(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := StringToFixed64(strVal); err == nil {
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

func flattenVertexAIIndexMetadataConfigAlgorithmConfigBruteForceConfig(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	transformed := make(map[string]interface{})
	return []interface{}{transformed}
}

func flattenVertexAIIndexMetadataSchemaUri(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenVertexAIIndexDeployedIndexes(v interface{}, d *schema.ResourceData, config *Config) interface{} {
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
			"index_endpoint":    flattenVertexAIIndexDeployedIndexesIndexEndpoint(original["indexEndpoint"], d, config),
			"deployed_index_id": flattenVertexAIIndexDeployedIndexesDeployedIndexId(original["deployedIndexId"], d, config),
		})
	}
	return transformed
}
func flattenVertexAIIndexDeployedIndexesIndexEndpoint(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenVertexAIIndexDeployedIndexesDeployedIndexId(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenVertexAIIndexLabels(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenVertexAIIndexCreateTime(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenVertexAIIndexUpdateTime(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenVertexAIIndexIndexStats(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["vectors_count"] =
		flattenVertexAIIndexIndexStatsVectorsCount(original["vectorsCount"], d, config)
	transformed["shards_count"] =
		flattenVertexAIIndexIndexStatsShardsCount(original["shardsCount"], d, config)
	return []interface{}{transformed}
}
func flattenVertexAIIndexIndexStatsVectorsCount(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenVertexAIIndexIndexStatsShardsCount(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := StringToFixed64(strVal); err == nil {
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

func flattenVertexAIIndexIndexUpdateMethod(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandVertexAIIndexDisplayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIIndexDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIIndexMetadata(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedContentsDeltaUri, err := expandVertexAIIndexMetadataContentsDeltaUri(original["contents_delta_uri"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedContentsDeltaUri); val.IsValid() && !isEmptyValue(val) {
		transformed["contentsDeltaUri"] = transformedContentsDeltaUri
	}

	transformedIsCompleteOverwrite, err := expandVertexAIIndexMetadataIsCompleteOverwrite(original["is_complete_overwrite"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIsCompleteOverwrite); val.IsValid() && !isEmptyValue(val) {
		transformed["isCompleteOverwrite"] = transformedIsCompleteOverwrite
	}

	transformedConfig, err := expandVertexAIIndexMetadataConfig(original["config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedConfig); val.IsValid() && !isEmptyValue(val) {
		transformed["config"] = transformedConfig
	}

	return transformed, nil
}

func expandVertexAIIndexMetadataContentsDeltaUri(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIIndexMetadataIsCompleteOverwrite(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIIndexMetadataConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDimensions, err := expandVertexAIIndexMetadataConfigDimensions(original["dimensions"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDimensions); val.IsValid() && !isEmptyValue(val) {
		transformed["dimensions"] = transformedDimensions
	}

	transformedApproximateNeighborsCount, err := expandVertexAIIndexMetadataConfigApproximateNeighborsCount(original["approximate_neighbors_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedApproximateNeighborsCount); val.IsValid() && !isEmptyValue(val) {
		transformed["approximateNeighborsCount"] = transformedApproximateNeighborsCount
	}

	transformedDistanceMeasureType, err := expandVertexAIIndexMetadataConfigDistanceMeasureType(original["distance_measure_type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDistanceMeasureType); val.IsValid() && !isEmptyValue(val) {
		transformed["distanceMeasureType"] = transformedDistanceMeasureType
	}

	transformedFeatureNormType, err := expandVertexAIIndexMetadataConfigFeatureNormType(original["feature_norm_type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFeatureNormType); val.IsValid() && !isEmptyValue(val) {
		transformed["featureNormType"] = transformedFeatureNormType
	}

	transformedAlgorithmConfig, err := expandVertexAIIndexMetadataConfigAlgorithmConfig(original["algorithm_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAlgorithmConfig); val.IsValid() && !isEmptyValue(val) {
		transformed["algorithmConfig"] = transformedAlgorithmConfig
	}

	return transformed, nil
}

func expandVertexAIIndexMetadataConfigDimensions(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIIndexMetadataConfigApproximateNeighborsCount(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIIndexMetadataConfigDistanceMeasureType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIIndexMetadataConfigFeatureNormType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIIndexMetadataConfigAlgorithmConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedTreeAhConfig, err := expandVertexAIIndexMetadataConfigAlgorithmConfigTreeAhConfig(original["tree_ah_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTreeAhConfig); val.IsValid() && !isEmptyValue(val) {
		transformed["treeAhConfig"] = transformedTreeAhConfig
	}

	transformedBruteForceConfig, err := expandVertexAIIndexMetadataConfigAlgorithmConfigBruteForceConfig(original["brute_force_config"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["bruteForceConfig"] = transformedBruteForceConfig
	}

	return transformed, nil
}

func expandVertexAIIndexMetadataConfigAlgorithmConfigTreeAhConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedLeafNodeEmbeddingCount, err := expandVertexAIIndexMetadataConfigAlgorithmConfigTreeAhConfigLeafNodeEmbeddingCount(original["leaf_node_embedding_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLeafNodeEmbeddingCount); val.IsValid() && !isEmptyValue(val) {
		transformed["leafNodeEmbeddingCount"] = transformedLeafNodeEmbeddingCount
	}

	transformedLeafNodesToSearchPercent, err := expandVertexAIIndexMetadataConfigAlgorithmConfigTreeAhConfigLeafNodesToSearchPercent(original["leaf_nodes_to_search_percent"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLeafNodesToSearchPercent); val.IsValid() && !isEmptyValue(val) {
		transformed["leafNodesToSearchPercent"] = transformedLeafNodesToSearchPercent
	}

	return transformed, nil
}

func expandVertexAIIndexMetadataConfigAlgorithmConfigTreeAhConfigLeafNodeEmbeddingCount(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIIndexMetadataConfigAlgorithmConfigTreeAhConfigLeafNodesToSearchPercent(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIIndexMetadataConfigAlgorithmConfigBruteForceConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 {
		return nil, nil
	}

	if l[0] == nil {
		transformed := make(map[string]interface{})
		return transformed, nil
	}
	transformed := make(map[string]interface{})

	return transformed, nil
}

func expandVertexAIIndexLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandVertexAIIndexIndexUpdateMethod(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
