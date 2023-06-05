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

	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

var GKEHubMembershipIamSchema = map[string]*schema.Schema{
	"project": {
		Type:     schema.TypeString,
		Computed: true,
		Optional: true,
		ForceNew: true,
	},
	"membership_id": {
		Type:             schema.TypeString,
		Required:         true,
		ForceNew:         true,
		DiffSuppressFunc: compareSelfLinkOrResourceName,
	},
}

type GKEHubMembershipIamUpdater struct {
	project      string
	membershipId string
	d            TerraformResourceData
	Config       *transport_tpg.Config
}

func GKEHubMembershipIamUpdaterProducer(d TerraformResourceData, config *transport_tpg.Config) (ResourceIamUpdater, error) {
	values := make(map[string]string)

	project, _ := getProject(d, config)
	if project != "" {
		if err := d.Set("project", project); err != nil {
			return nil, fmt.Errorf("Error setting project: %s", err)
		}
	}
	values["project"] = project
	if v, ok := d.GetOk("membership_id"); ok {
		values["membership_id"] = v.(string)
	}

	// We may have gotten either a long or short name, so attempt to parse long name if possible
	m, err := getImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/memberships/(?P<membership_id>[^/]+)", "(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<membership_id>[^/]+)", "(?P<location>[^/]+)/(?P<membership_id>[^/]+)", "(?P<membership_id>[^/]+)"}, d, config, d.Get("membership_id").(string))
	if err != nil {
		return nil, err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &GKEHubMembershipIamUpdater{
		project:      values["project"],
		membershipId: values["membership_id"],
		d:            d,
		Config:       config,
	}

	if err := d.Set("project", u.project); err != nil {
		return nil, fmt.Errorf("Error setting project: %s", err)
	}
	if err := d.Set("membership_id", u.GetResourceId()); err != nil {
		return nil, fmt.Errorf("Error setting membership_id: %s", err)
	}

	return u, nil
}

func GKEHubMembershipIdParseFunc(d *schema.ResourceData, config *transport_tpg.Config) error {
	values := make(map[string]string)

	project, _ := getProject(d, config)
	if project != "" {
		values["project"] = project
	}

	m, err := getImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/memberships/(?P<membership_id>[^/]+)", "(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<membership_id>[^/]+)", "(?P<location>[^/]+)/(?P<membership_id>[^/]+)", "(?P<membership_id>[^/]+)"}, d, config, d.Id())
	if err != nil {
		return err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &GKEHubMembershipIamUpdater{
		project:      values["project"],
		membershipId: values["membership_id"],
		d:            d,
		Config:       config,
	}
	if err := d.Set("membership_id", u.GetResourceId()); err != nil {
		return fmt.Errorf("Error setting membership_id: %s", err)
	}
	d.SetId(u.GetResourceId())
	return nil
}

func (u *GKEHubMembershipIamUpdater) GetResourceIamPolicy() (*cloudresourcemanager.Policy, error) {
	url, err := u.qualifyMembershipUrl("getIamPolicy")
	if err != nil {
		return nil, err
	}

	project, err := getProject(u.d, u.Config)
	if err != nil {
		return nil, err
	}
	var obj map[string]interface{}

	userAgent, err := generateUserAgentString(u.d, u.Config.UserAgent)
	if err != nil {
		return nil, err
	}

	policy, err := transport_tpg.SendRequest(u.Config, "GET", project, url, userAgent, obj)
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

func (u *GKEHubMembershipIamUpdater) SetResourceIamPolicy(policy *cloudresourcemanager.Policy) error {
	json, err := ConvertToMap(policy)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	obj["policy"] = json

	url, err := u.qualifyMembershipUrl("setIamPolicy")
	if err != nil {
		return err
	}
	project, err := getProject(u.d, u.Config)
	if err != nil {
		return err
	}

	userAgent, err := generateUserAgentString(u.d, u.Config.UserAgent)
	if err != nil {
		return err
	}

	_, err = transport_tpg.SendRequestWithTimeout(u.Config, "POST", project, url, userAgent, obj, u.d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return errwrap.Wrapf(fmt.Sprintf("Error setting IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	return nil
}

func (u *GKEHubMembershipIamUpdater) qualifyMembershipUrl(methodIdentifier string) (string, error) {
	urlTemplate := fmt.Sprintf("{{GKEHubBasePath}}%s:%s", fmt.Sprintf("projects/%s/locations/global/memberships/%s", u.project, u.membershipId), methodIdentifier)
	url, err := ReplaceVars(u.d, u.Config, urlTemplate)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (u *GKEHubMembershipIamUpdater) GetResourceId() string {
	return fmt.Sprintf("projects/%s/locations/global/memberships/%s", u.project, u.membershipId)
}

func (u *GKEHubMembershipIamUpdater) GetMutexKey() string {
	return fmt.Sprintf("iam-gkehub-membership-%s", u.GetResourceId())
}

func (u *GKEHubMembershipIamUpdater) DescribeResource() string {
	return fmt.Sprintf("gkehub membership %q", u.GetResourceId())
}
