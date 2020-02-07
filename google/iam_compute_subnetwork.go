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

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"google.golang.org/api/cloudresourcemanager/v1"
)

var ComputeSubnetworkIamSchema = map[string]*schema.Schema{
	"project": {
		Type:     schema.TypeString,
		Computed: true,
		Optional: true,
		ForceNew: true,
	},
	"region": {
		Type:     schema.TypeString,
		Computed: true,
		Optional: true,
		ForceNew: true,
	},
	"subnetwork": {
		Type:             schema.TypeString,
		Required:         true,
		ForceNew:         true,
		DiffSuppressFunc: compareSelfLinkOrResourceName,
	},
}

type ComputeSubnetworkIamUpdater struct {
	project    string
	region     string
	subnetwork string
	d          *schema.ResourceData
	Config     *Config
}

func ComputeSubnetworkIamUpdaterProducer(d *schema.ResourceData, config *Config) (ResourceIamUpdater, error) {
	values := make(map[string]string)

	project, _ := getProject(d, config)
	if project != "" {
		values["project"] = project
	}
	values["project"] = project
	region, _ := getRegion(d, config)
	if region != "" {
		values["region"] = region
	}
	values["region"] = region
	if v, ok := d.GetOk("subnetwork"); ok {
		values["subnetwork"] = v.(string)
	}

	// We may have gotten either a long or short name, so attempt to parse long name if possible
	m, err := getImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/regions/(?P<region>[^/]+)/subnetworks/(?P<subnetwork>[^/]+)", "(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<subnetwork>[^/]+)", "(?P<region>[^/]+)/(?P<subnetwork>[^/]+)", "(?P<subnetwork>[^/]+)"}, d, config, d.Get("subnetwork").(string))
	if err != nil {
		return nil, err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &ComputeSubnetworkIamUpdater{
		project:    values["project"],
		region:     values["region"],
		subnetwork: values["subnetwork"],
		d:          d,
		Config:     config,
	}

	d.Set("project", u.project)
	d.Set("region", u.region)
	d.Set("subnetwork", u.GetResourceId())

	return u, nil
}

func ComputeSubnetworkIdParseFunc(d *schema.ResourceData, config *Config) error {
	values := make(map[string]string)

	project, _ := getProject(d, config)
	if project != "" {
		values["project"] = project
	}

	region, _ := getRegion(d, config)
	if region != "" {
		values["region"] = region
	}

	m, err := getImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/regions/(?P<region>[^/]+)/subnetworks/(?P<subnetwork>[^/]+)", "(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<subnetwork>[^/]+)", "(?P<region>[^/]+)/(?P<subnetwork>[^/]+)", "(?P<subnetwork>[^/]+)"}, d, config, d.Id())
	if err != nil {
		return err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &ComputeSubnetworkIamUpdater{
		project:    values["project"],
		region:     values["region"],
		subnetwork: values["subnetwork"],
		d:          d,
		Config:     config,
	}
	d.Set("subnetwork", u.GetResourceId())
	d.SetId(u.GetResourceId())
	return nil
}

func (u *ComputeSubnetworkIamUpdater) GetResourceIamPolicy() (*cloudresourcemanager.Policy, error) {
	url, err := u.qualifySubnetworkUrl("getIamPolicy")
	if err != nil {
		return nil, err
	}

	project, err := getProject(u.d, u.Config)
	if err != nil {
		return nil, err
	}
	var obj map[string]interface{}

	policy, err := sendRequest(u.Config, "GET", project, url, obj)
	if err != nil {
		return nil, errwrap.Wrapf(fmt.Sprintf("Error retrieving IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	out := &cloudresourcemanager.Policy{}
	err = Convert(policy, out)
	if err != nil {
		return nil, errwrap.Wrapf("Cannot convert a policy to a resource manager policy: {{err}}", err)
	}

	return out, nil
}

func (u *ComputeSubnetworkIamUpdater) SetResourceIamPolicy(policy *cloudresourcemanager.Policy) error {
	json, err := ConvertToMap(policy)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	obj["policy"] = json

	url, err := u.qualifySubnetworkUrl("setIamPolicy")
	if err != nil {
		return err
	}
	project, err := getProject(u.d, u.Config)
	if err != nil {
		return err
	}

	_, err = sendRequestWithTimeout(u.Config, "POST", project, url, obj, u.d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return errwrap.Wrapf(fmt.Sprintf("Error setting IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	return nil
}

func (u *ComputeSubnetworkIamUpdater) qualifySubnetworkUrl(methodIdentifier string) (string, error) {
	urlTemplate := fmt.Sprintf("{{ComputeBasePath}}%s/%s", fmt.Sprintf("projects/%s/regions/%s/subnetworks/%s", u.project, u.region, u.subnetwork), methodIdentifier)
	url, err := replaceVars(u.d, u.Config, urlTemplate)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (u *ComputeSubnetworkIamUpdater) GetResourceId() string {
	return fmt.Sprintf("projects/%s/regions/%s/subnetworks/%s", u.project, u.region, u.subnetwork)
}

func (u *ComputeSubnetworkIamUpdater) GetMutexKey() string {
	return fmt.Sprintf("iam-compute-subnetwork-%s", u.GetResourceId())
}

func (u *ComputeSubnetworkIamUpdater) DescribeResource() string {
	return fmt.Sprintf("compute subnetwork %q", u.GetResourceId())
}
