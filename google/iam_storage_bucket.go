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

var StorageBucketIamSchema = map[string]*schema.Schema{
	"bucket": {
		Type:             schema.TypeString,
		Required:         true,
		ForceNew:         true,
		DiffSuppressFunc: compareSelfLinkOrResourceName,
	},
}

type StorageBucketIamUpdater struct {
	bucket string
	d      *schema.ResourceData
	Config *Config
}

func StorageBucketIamUpdaterProducer(d *schema.ResourceData, config *Config) (ResourceIamUpdater, error) {
	values := make(map[string]string)

	if v, ok := d.GetOk("bucket"); ok {
		values["bucket"] = v.(string)
	}

	// We may have gotten either a long or short name, so attempt to parse long name if possible
	m, err := getImportIdQualifiers([]string{"b/(?P<bucket>[^/]+)", "(?P<bucket>[^/]+)"}, d, config, d.Get("bucket").(string))
	if err != nil {
		return nil, err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &StorageBucketIamUpdater{
		bucket: values["bucket"],
		d:      d,
		Config: config,
	}

	d.Set("bucket", u.GetResourceId())

	return u, nil
}

func StorageBucketIdParseFunc(d *schema.ResourceData, config *Config) error {
	values := make(map[string]string)

	m, err := getImportIdQualifiers([]string{"b/(?P<bucket>[^/]+)", "(?P<bucket>[^/]+)"}, d, config, d.Id())
	if err != nil {
		return err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &StorageBucketIamUpdater{
		bucket: values["bucket"],
		d:      d,
		Config: config,
	}
	d.Set("bucket", u.GetResourceId())
	d.SetId(u.GetResourceId())
	return nil
}

func (u *StorageBucketIamUpdater) GetResourceIamPolicy() (*cloudresourcemanager.Policy, error) {
	url, err := u.qualifyBucketUrl("iam")
	if err != nil {
		return nil, err
	}

	var obj map[string]interface{}

	policy, err := sendRequest(u.Config, "GET", "", url, obj)
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

func (u *StorageBucketIamUpdater) SetResourceIamPolicy(policy *cloudresourcemanager.Policy) error {
	json, err := ConvertToMap(policy)
	if err != nil {
		return err
	}

	obj := json

	url, err := u.qualifyBucketUrl("iam")
	if err != nil {
		return err
	}

	_, err = sendRequestWithTimeout(u.Config, "PUT", "", url, obj, u.d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return errwrap.Wrapf(fmt.Sprintf("Error setting IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	return nil
}

func (u *StorageBucketIamUpdater) qualifyBucketUrl(methodIdentifier string) (string, error) {
	urlTemplate := fmt.Sprintf("{{StorageBasePath}}%s/%s", fmt.Sprintf("b/%s", u.bucket), methodIdentifier)
	url, err := replaceVars(u.d, u.Config, urlTemplate)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (u *StorageBucketIamUpdater) GetResourceId() string {
	return fmt.Sprintf("b/%s", u.bucket)
}

func (u *StorageBucketIamUpdater) GetMutexKey() string {
	return fmt.Sprintf("iam-storage-bucket-%s", u.GetResourceId())
}

func (u *StorageBucketIamUpdater) DescribeResource() string {
	return fmt.Sprintf("storage bucket %q", u.GetResourceId())
}
