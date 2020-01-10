// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//     ***     DIFF TEST DIFF TEST    ***
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

var IapWebIamSchema = map[string]*schema.Schema{
	"project": {
		Type:             schema.TypeString,
		Computed:         true,
		Optional:         true,
		ForceNew:         true,
		DiffSuppressFunc: compareSelfLinkOrResourceName,
	},
}

type IapWebIamUpdater struct {
	project string
	d       *schema.ResourceData
	Config  *Config
}

func IapWebIamUpdaterProducer(d *schema.ResourceData, config *Config) (ResourceIamUpdater, error) {
	values := make(map[string]string)

	project, err := getProject(d, config)
	if err != nil {
		return nil, err
	}
	values["project"] = project

	// We may have gotten either a long or short name, so attempt to parse long name if possible
	m, err := getImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/iap_web", "(?P<project>[^/]+)"}, d, config, d.Get("project").(string))
	if err != nil {
		return nil, err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &IapWebIamUpdater{
		project: values["project"],
		d:       d,
		Config:  config,
	}

	d.Set("project", u.project)

	return u, nil
}

func IapWebIdParseFunc(d *schema.ResourceData, config *Config) error {
	values := make(map[string]string)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	values["project"] = project

	m, err := getImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/iap_web", "(?P<project>[^/]+)"}, d, config, d.Id())
	if err != nil {
		return err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &IapWebIamUpdater{
		project: values["project"],
		d:       d,
		Config:  config,
	}
	d.Set("project", u.project)
	d.SetId(u.GetResourceId())
	return nil
}

func (u *IapWebIamUpdater) GetResourceIamPolicy() (*cloudresourcemanager.Policy, error) {
	url, err := u.qualifyWebUrl("getIamPolicy")
	if err != nil {
		return nil, err
	}

	project, err := getProject(u.d, u.Config)
	if err != nil {
		return nil, err
	}
	var obj map[string]interface{}

	policy, err := sendRequest(u.Config, "POST", project, url, obj)
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

func (u *IapWebIamUpdater) SetResourceIamPolicy(policy *cloudresourcemanager.Policy) error {
	json, err := ConvertToMap(policy)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	obj["policy"] = json

	url, err := u.qualifyWebUrl("setIamPolicy")
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

func (u *IapWebIamUpdater) qualifyWebUrl(methodIdentifier string) (string, error) {
	urlTemplate := fmt.Sprintf("{{IapBasePath}}%s:%s", fmt.Sprintf("projects/%s/iap_web", u.project), methodIdentifier)
	url, err := replaceVars(u.d, u.Config, urlTemplate)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (u *IapWebIamUpdater) GetResourceId() string {
	return fmt.Sprintf("projects/%s/iap_web", u.project)
}

func (u *IapWebIamUpdater) GetMutexKey() string {
	return fmt.Sprintf("iam-iap-web-%s", u.GetResourceId())
}

func (u *IapWebIamUpdater) DescribeResource() string {
	return fmt.Sprintf("iap web %q", u.GetResourceId())
}
