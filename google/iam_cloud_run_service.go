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

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"google.golang.org/api/cloudresourcemanager/v1"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

var CloudRunServiceIamSchema = map[string]*schema.Schema{
	"project": {
		Type:     schema.TypeString,
		Computed: true,
		Optional: true,
		ForceNew: true,
	},
	"location": {
		Type:     schema.TypeString,
		Computed: true,
		Optional: true,
		ForceNew: true,
	},
	"service": {
		Type:             schema.TypeString,
		Required:         true,
		ForceNew:         true,
		DiffSuppressFunc: compareSelfLinkOrResourceName,
	},
}

type CloudRunServiceIamUpdater struct {
	project  string
	location string
	service  string
	d        tpgresource.TerraformResourceData
	Config   *transport_tpg.Config
}

func CloudRunServiceIamUpdaterProducer(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (ResourceIamUpdater, error) {
	values := make(map[string]string)

	project, _ := getProject(d, config)
	if project != "" {
		if err := d.Set("project", project); err != nil {
			return nil, fmt.Errorf("Error setting project: %s", err)
		}
	}
	values["project"] = project
	location, _ := getLocation(d, config)
	if location != "" {
		if err := d.Set("location", location); err != nil {
			return nil, fmt.Errorf("Error setting location: %s", err)
		}
	}
	values["location"] = location
	if v, ok := d.GetOk("service"); ok {
		values["service"] = v.(string)
	}

	// We may have gotten either a long or short name, so attempt to parse long name if possible
	m, err := tpgresource.GetImportIdQualifiers([]string{"projects/(?P<project>.+)/locations/(?P<location>.+)/services/(?P<service>.+)", "(?P<project>.+)/(?P<location>.+)/(?P<service>.+)", "(?P<location>.+)/(?P<service>.+)", "(?P<service>.+)"}, d, config, d.Get("service").(string))
	if err != nil {
		return nil, err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &CloudRunServiceIamUpdater{
		project:  values["project"],
		location: values["location"],
		service:  values["service"],
		d:        d,
		Config:   config,
	}

	if err := d.Set("project", u.project); err != nil {
		return nil, fmt.Errorf("Error setting project: %s", err)
	}
	if err := d.Set("location", u.location); err != nil {
		return nil, fmt.Errorf("Error setting location: %s", err)
	}
	if err := d.Set("service", u.GetResourceId()); err != nil {
		return nil, fmt.Errorf("Error setting service: %s", err)
	}

	return u, nil
}

func CloudRunServiceIdParseFunc(d *schema.ResourceData, config *transport_tpg.Config) error {
	values := make(map[string]string)

	project, _ := getProject(d, config)
	if project != "" {
		values["project"] = project
	}

	location, _ := getLocation(d, config)
	if location != "" {
		values["location"] = location
	}

	m, err := tpgresource.GetImportIdQualifiers([]string{"projects/(?P<project>.+)/locations/(?P<location>.+)/services/(?P<service>.+)", "(?P<project>.+)/(?P<location>.+)/(?P<service>.+)", "(?P<location>.+)/(?P<service>.+)", "(?P<service>.+)"}, d, config, d.Id())
	if err != nil {
		return err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &CloudRunServiceIamUpdater{
		project:  values["project"],
		location: values["location"],
		service:  values["service"],
		d:        d,
		Config:   config,
	}
	if err := d.Set("service", u.GetResourceId()); err != nil {
		return fmt.Errorf("Error setting service: %s", err)
	}
	d.SetId(u.GetResourceId())
	return nil
}

func (u *CloudRunServiceIamUpdater) GetResourceIamPolicy() (*cloudresourcemanager.Policy, error) {
	url, err := u.qualifyServiceUrl("getIamPolicy")
	if err != nil {
		return nil, err
	}

	project, err := tpgresource.GetProject(u.d, u.Config)
	if err != nil {
		return nil, err
	}
	var obj map[string]interface{}

	userAgent, err := tpgresource.GenerateUserAgentString(u.d, u.Config.UserAgent)
	if err != nil {
		return nil, err
	}

	policy, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:               u.Config,
		Method:               "GET",
		Project:              project,
		RawURL:               url,
		UserAgent:            userAgent,
		Body:                 obj,
		ErrorRetryPredicates: []transport_tpg.RetryErrorPredicateFunc{transport_tpg.IsCloudRunCreationConflict},
	})
	if err != nil {
		return nil, errwrap.Wrapf(fmt.Sprintf("Error retrieving IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	out := &cloudresourcemanager.Policy{}
	err = tpgresource.Convert(policy, out)
	if err != nil {
		return nil, errwrap.Wrapf("Cannot convert a policy to a resource manager policy: {{err}}", err)
	}

	return out, nil
}

func (u *CloudRunServiceIamUpdater) SetResourceIamPolicy(policy *cloudresourcemanager.Policy) error {
	json, err := tpgresource.ConvertToMap(policy)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	obj["policy"] = json

	url, err := u.qualifyServiceUrl("setIamPolicy")
	if err != nil {
		return err
	}
	project, err := tpgresource.GetProject(u.d, u.Config)
	if err != nil {
		return err
	}

	userAgent, err := tpgresource.GenerateUserAgentString(u.d, u.Config.UserAgent)
	if err != nil {
		return err
	}

	_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:               u.Config,
		Method:               "POST",
		Project:              project,
		RawURL:               url,
		UserAgent:            userAgent,
		Body:                 obj,
		Timeout:              u.d.Timeout(schema.TimeoutCreate),
		ErrorRetryPredicates: []transport_tpg.RetryErrorPredicateFunc{transport_tpg.IsCloudRunCreationConflict},
	})
	if err != nil {
		return errwrap.Wrapf(fmt.Sprintf("Error setting IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	return nil
}

func (u *CloudRunServiceIamUpdater) qualifyServiceUrl(methodIdentifier string) (string, error) {
	urlTemplate := fmt.Sprintf("{{CloudRunBasePath}}%s:%s", fmt.Sprintf("v1/projects/%s/locations/%s/services/%s", u.project, u.location, u.service), methodIdentifier)
	url, err := tpgresource.ReplaceVars(u.d, u.Config, urlTemplate)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (u *CloudRunServiceIamUpdater) GetResourceId() string {
	return fmt.Sprintf("v1/projects/%s/locations/%s/services/%s", u.project, u.location, u.service)
}

func (u *CloudRunServiceIamUpdater) GetMutexKey() string {
	return fmt.Sprintf("iam-cloudrun-service-%s", u.GetResourceId())
}

func (u *CloudRunServiceIamUpdater) DescribeResource() string {
	return fmt.Sprintf("cloudrun service %q", u.GetResourceId())
}
