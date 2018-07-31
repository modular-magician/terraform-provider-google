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



func resourceComputeRouter() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeRouterCreate,
		Read: resourceComputeRouterRead,
				Update: resourceComputeRouterUpdate,
				Delete: resourceComputeRouterDelete,
        
		Importer: &schema.ResourceImporter{
			State: resourceComputeRouterImport,
		},

				Timeouts: &schema.ResourceTimeout {
			Create: schema.DefaultTimeout(240 * time.Second),
						Update: schema.DefaultTimeout(240 * time.Second),
						Delete: schema.DefaultTimeout(240 * time.Second),
		},
		
		Schema: map[string]*schema.Schema{
	"name": {
    Type: schema.TypeString,
    Required: true,
  ForceNew: true,
		ValidateFunc: validateGCPName,
	},
	"network": {
    Type: schema.TypeString,
    Required: true,
  ForceNew: true,
  DiffSuppressFunc: compareSelfLinkOrResourceName,
},
	"description": {
    Type: schema.TypeString,
    Optional: true,
},
	"bgp": {
    Type: schema.TypeList,
    Optional: true,
  MaxItems: 1,
  Elem: &schema.Resource{
    Schema: map[string]*schema.Schema{
              "asn": {
    Type: schema.TypeInt,
    Required: true,
},
              "advertise_mode": {
    Type: schema.TypeString,
    Optional: true,
	ValidateFunc: validation.StringInSlice([]string{"DEFAULT","CUSTOM",""}, false),
    Default: "DEFAULT",
},
              "advertised_groups": {
    Type: schema.TypeList,
    Optional: true,
           Elem: &schema.Schema{
      Type: schema.TypeString,
            },
    },
              "advertised_ip_ranges": {
    Type: schema.TypeList,
    Optional: true,
           Elem: &schema.Resource{
        Schema: map[string]*schema.Schema{
                      "range": {
    Type: schema.TypeString,
    Optional: true,
},
                      "description": {
    Type: schema.TypeString,
    Optional: true,
},
                  },
      },
    },
          },
  },
},
	"region": {
    Type: schema.TypeString,
  	Computed: true,
	Optional: true,
  ForceNew: true,
  DiffSuppressFunc: compareSelfLinkOrResourceName,
},
	"creation_timestamp": {
    Type: schema.TypeString,
    Computed: true,
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

func resourceComputeRouterCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

  	project, err := getProject(d, config)
	if err != nil {
		return err
  }
  
	obj := make(map[string]interface{})
	nameProp, err := expandComputeRouterName(d.Get("name"), d, config)
	if err != nil {
		return err
		} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
			obj["name"] = nameProp
	}
	descriptionProp, err := expandComputeRouterDescription(d.Get("description"), d, config)
	if err != nil {
		return err
		} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
			obj["description"] = descriptionProp
	}
	networkProp, err := expandComputeRouterNetwork(d.Get("network"), d, config)
	if err != nil {
		return err
		} else if v, ok := d.GetOkExists("network"); !isEmptyValue(reflect.ValueOf(networkProp)) && (ok || !reflect.DeepEqual(v, networkProp)) {
			obj["network"] = networkProp
	}
	bgpProp, err := expandComputeRouterBgp(d.Get("bgp"), d, config)
	if err != nil {
		return err
		} else if v, ok := d.GetOkExists("bgp"); !isEmptyValue(reflect.ValueOf(bgpProp)) && (ok || !reflect.DeepEqual(v, bgpProp)) {
			obj["bgp"] = bgpProp
	}
	regionProp, err := expandComputeRouterRegion(d.Get("region"), d, config)
	if err != nil {
		return err
		} else if v, ok := d.GetOkExists("region"); !isEmptyValue(reflect.ValueOf(regionProp)) && (ok || !reflect.DeepEqual(v, regionProp)) {
			obj["region"] = regionProp
	}

  
	lockName, err := replaceVars(d, config, "router/{{region}}/{{name}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/regions/{{region}}/routers")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Router: %#v", obj)
	res, err := Post(config, url, obj)
	if err != nil {
		return fmt.Errorf("Error creating Router: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{region}}/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

		op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	waitErr := computeOperationWaitTime(
	config.clientCompute, op, project,  "Creating Router",
		int(d.Timeout(schema.TimeoutCreate).Minutes()))

	if waitErr != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create Router: %s", waitErr)
	}
	
	log.Printf("[DEBUG] Finished creating Router %q: %#v", d.Id(), res)


	return resourceComputeRouterRead(d, meta)
}

func resourceComputeRouterRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

  	project, err := getProject(d, config)
	if err != nil {
		return err
  }
  
	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/regions/{{region}}/routers/{{name}}")
	if err != nil {
		return err
	}

	res, err := Get(config, url)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeRouter %q", d.Id()))
	}





  if err := d.Set("creation_timestamp", flattenComputeRouterCreationTimestamp(res["creationTimestamp"])); err != nil {
		return fmt.Errorf("Error reading Router: %s", err)
	}
  if err := d.Set("name", flattenComputeRouterName(res["name"])); err != nil {
		return fmt.Errorf("Error reading Router: %s", err)
	}
  if err := d.Set("description", flattenComputeRouterDescription(res["description"])); err != nil {
		return fmt.Errorf("Error reading Router: %s", err)
	}
  if err := d.Set("network", flattenComputeRouterNetwork(res["network"])); err != nil {
		return fmt.Errorf("Error reading Router: %s", err)
	}
  if err := d.Set("bgp", flattenComputeRouterBgp(res["bgp"])); err != nil {
		return fmt.Errorf("Error reading Router: %s", err)
	}
  if err := d.Set("region", flattenComputeRouterRegion(res["region"])); err != nil {
		return fmt.Errorf("Error reading Router: %s", err)
	}
	if err := d.Set("self_link", ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading Router: %s", err)
	}
	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Router: %s", err)
	}

	return nil
}

func resourceComputeRouterUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

  	project, err := getProject(d, config)
	if err != nil {
		return err
	}
  
		obj := make(map[string]interface{})
			nameProp, err := expandComputeRouterName(d.Get("name"), d, config)
		if err != nil {
				return err
				} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, nameProp)) {
					obj["name"] = nameProp
		}
			descriptionProp, err := expandComputeRouterDescription(d.Get("description"), d, config)
		if err != nil {
				return err
				} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
					obj["description"] = descriptionProp
		}
			networkProp, err := expandComputeRouterNetwork(d.Get("network"), d, config)
		if err != nil {
				return err
				} else if v, ok := d.GetOkExists("network"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, networkProp)) {
					obj["network"] = networkProp
		}
			bgpProp, err := expandComputeRouterBgp(d.Get("bgp"), d, config)
		if err != nil {
				return err
				} else if v, ok := d.GetOkExists("bgp"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, bgpProp)) {
					obj["bgp"] = bgpProp
		}
			regionProp, err := expandComputeRouterRegion(d.Get("region"), d, config)
		if err != nil {
				return err
				} else if v, ok := d.GetOkExists("region"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, regionProp)) {
					obj["region"] = regionProp
		}
	
		
	lockName, err := replaceVars(d, config, "router/{{region}}/{{name}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := replaceVars(d, config, "[["https://www.googleapis.com/compute/v1/"], "projects/{{project}}/regions/{{region}}/routers/{{name}}"]")
	if err != nil {
	return err
	}

	log.Printf("[DEBUG] Updating Router %q: %#v", d.Id(), obj)
	res, err := sendRequest(config, "PATCH", url, obj)

	if err != nil {
	return fmt.Errorf("Error updating Router %q: %s", d.Id(), err)
	}

		op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
	return err
	}

	err = computeOperationWaitTime(
	config.clientCompute, op,  project,  "Updating Router",
	int(d.Timeout(schema.TimeoutUpdate).Minutes()))

	if err != nil {
	return err
	}
		
	return resourceComputeRouterRead(d, meta)
}

func resourceComputeRouterDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

  	project, err := getProject(d, config)
	if err != nil {
		return err
  }
  
	lockName, err := replaceVars(d, config, "router/{{region}}/{{name}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/regions/{{region}}/routers/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Deleting Router %q", d.Id())
	res, err := Delete(config, url)
	if err != nil {
		return handleNotFoundError(err, d, "Router")
	}

		op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	err = computeOperationWaitTime(
		config.clientCompute, op,  project,  "Deleting Router",
		int(d.Timeout(schema.TimeoutDelete).Minutes()))

	if err != nil {
		return err
	}
	
	log.Printf("[DEBUG] Finished deleting Router %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeRouterImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	parseImportId([]string{"projects/(?P<project>[^/]+)/regions/(?P<region>[^/]+)/routers/(?P<name>[^/]+)","(?P<region>[^/]+)/(?P<name>[^/]+)","(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<name>[^/]+)","(?P<name>[^/]+)"}, d, config)

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "{{region}}/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)
  
	return []*schema.ResourceData{d}, nil
}

func flattenComputeRouterCreationTimestamp(v interface{}) interface{} {
  return v
}

func flattenComputeRouterName(v interface{}) interface{} {
  return v
}

func flattenComputeRouterDescription(v interface{}) interface{} {
  return v
}

func flattenComputeRouterNetwork(v interface{}) interface{} {
  if v == nil {
    return v
  }
  return ConvertSelfLinkToV1(v.(string))
}

func flattenComputeRouterBgp(v interface{}) interface{} {
  if v == nil {
    return nil
  }
  original := v.(map[string]interface{})
  transformed := make(map[string]interface{})
      transformed["asn"] =
    flattenComputeRouterBgpAsn(original["asn"])
      transformed["advertise_mode"] =
    flattenComputeRouterBgpAdvertiseMode(original["advertiseMode"])
      transformed["advertised_groups"] =
    flattenComputeRouterBgpAdvertisedGroups(original["advertisedGroups"])
      transformed["advertised_ip_ranges"] =
    flattenComputeRouterBgpAdvertisedIpRanges(original["advertisedIpRanges"])
    return []interface{}{transformed}
}
      func flattenComputeRouterBgpAsn(v interface{}) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

      func flattenComputeRouterBgpAdvertiseMode(v interface{}) interface{} {
  return v
}

      func flattenComputeRouterBgpAdvertisedGroups(v interface{}) interface{} {
  return v
}

      func flattenComputeRouterBgpAdvertisedIpRanges(v interface{}) interface{} {
  if v == nil {
    return v
  }
  l := v.([]interface{})
  transformed := make([]interface{}, 0, len(l))
  for _, raw := range l {
    original := raw.(map[string]interface{})
    transformed = append(transformed, map[string]interface{}{
          "range": flattenComputeRouterBgpAdvertisedIpRangesRange(original["range"]),
          "description": flattenComputeRouterBgpAdvertisedIpRangesDescription(original["description"]),
        })
  }
  return transformed
}
      func flattenComputeRouterBgpAdvertisedIpRangesRange(v interface{}) interface{} {
  return v
}

      func flattenComputeRouterBgpAdvertisedIpRangesDescription(v interface{}) interface{} {
  return v
}

  

  

func flattenComputeRouterRegion(v interface{}) interface{} {
	return NameFromSelfLinkStateFunc(v)
}


func expandComputeRouterName(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
  return v, nil
}

func expandComputeRouterDescription(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
  return v, nil
}

func expandComputeRouterNetwork(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
  f, err := parseGlobalFieldValue("networks", v.(string), "project", d, config, true)
  if err != nil {
    return nil, fmt.Errorf("Invalid value for network: %s", err)
  }
  return f.RelativeLink(), nil
}

func expandComputeRouterBgp(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
  l := v.([]interface{})
  if len(l) == 0 {
    return nil, nil
  }
  raw := l[0]
    original := raw.(map[string]interface{})
    transformed := make(map[string]interface{})

      transformedAsn, err := expandComputeRouterBgpAsn(original["asn"], d, config)
      if err != nil {
        return nil, err
      }
      transformed["asn"] = transformedAsn
      transformedAdvertiseMode, err := expandComputeRouterBgpAdvertiseMode(original["advertise_mode"], d, config)
      if err != nil {
        return nil, err
      }
      transformed["advertiseMode"] = transformedAdvertiseMode
      transformedAdvertisedGroups, err := expandComputeRouterBgpAdvertisedGroups(original["advertised_groups"], d, config)
      if err != nil {
        return nil, err
      }
      transformed["advertisedGroups"] = transformedAdvertisedGroups
      transformedAdvertisedIpRanges, err := expandComputeRouterBgpAdvertisedIpRanges(original["advertised_ip_ranges"], d, config)
      if err != nil {
        return nil, err
      }
      transformed["advertisedIpRanges"] = transformedAdvertisedIpRanges
  return transformed, nil
}

func expandComputeRouterBgpAsn(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
  return v, nil
}

func expandComputeRouterBgpAdvertiseMode(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
  return v, nil
}

func expandComputeRouterBgpAdvertisedGroups(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
  return v, nil
}

func expandComputeRouterBgpAdvertisedIpRanges(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
  l := v.([]interface{})
  req := make([]interface{}, 0, len(l))
  for _, raw := range l {
    original := raw.(map[string]interface{})
    transformed := make(map[string]interface{})

      transformedRange, err := expandComputeRouterBgpAdvertisedIpRangesRange(original["range"], d, config)
      if err != nil {
        return nil, err
      }
      transformed["range"] = transformedRange
      transformedDescription, err := expandComputeRouterBgpAdvertisedIpRangesDescription(original["description"], d, config)
      if err != nil {
        return nil, err
      }
      transformed["description"] = transformedDescription
    req = append(req, transformed)
  }
  return req, nil
}

func expandComputeRouterBgpAdvertisedIpRangesRange(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
  return v, nil
}

func expandComputeRouterBgpAdvertisedIpRangesDescription(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
  return v, nil
}

func expandComputeRouterRegion(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
  f, err := parseGlobalFieldValue("regions", v.(string), "project", d, config, true)
  if err != nil {
    return nil, fmt.Errorf("Invalid value for region: %s", err)
  }
  return f.RelativeLink(), nil
}
