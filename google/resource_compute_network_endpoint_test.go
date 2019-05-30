package google

import (
	"testing"

	"fmt"

	"log"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

// TODO(emilyye): Fix this test once network endpoint and group are GA.
// Terraform will add an implicit provider dependency on the GA provider
// instead of the non-GA provider if the resource no longer exists in the
// test step config (and no longer has the explicit provider declaration).
func TestAccComputeNetworkEndpoint_networkEndpointsBasic(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
	}

	initialEndpt := &networkEndptTestId{}
	modifiedEndpt := &networkEndptTestId{}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProvidersOiCS,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeNetworkEndpoint_networkEndpointsBasic(context),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeNetworkEndpointExists("google_compute_network_endpoint.default-endpoint", initialEndpt),
				),
			},
			{
				Config: testAccComputeNetworkEndpoint_networkEndpointsModified(context),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeNetworkEndpointExists("google_compute_network_endpoint.default-endpoint", modifiedEndpt),
					testAccCheckComputeNetworkEndpointDestroyed("google_compute_network_endpoint.default-endpoint", initialEndpt),
				),
			},
		},
	})
}

func testAccComputeNetworkEndpoint_networkEndpointsBasic(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_network_endpoint" "default-endpoint" {
 provider               = "google-beta"

 zone                   = "us-central1-a"
 network_endpoint_group = "${google_compute_network_endpoint_group.neg.name}"

 instance    = "${google_compute_instance.default.name}"
 ip_address  = "${google_compute_instance.default.network_interface.0.network_ip}"
 port        = "${google_compute_network_endpoint_group.neg.default_port}"
}
`, context) + testAccComputeNetworkEndpoint_withoutNetworkEndpointsConfig(context)
}

func testAccComputeNetworkEndpoint_networkEndpointsModified(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_network_endpoint" "default-endpoint" {
 provider               = "google-beta"

 zone                   = "us-central1-a"
 network_endpoint_group = "${google_compute_network_endpoint_group.neg.name}"

 instance    = "${google_compute_instance.default.name}"
 ip_address  = "${google_compute_instance.default.network_interface.0.network_ip}"
 port        = "100"
}
`, context) + testAccComputeNetworkEndpoint_withoutNetworkEndpointsConfig(context)
}

func testAccComputeNetworkEndpoint_withoutNetworkEndpointsConfig(context map[string]interface{}) string {
	return Nprintf(`
provider "google-beta"{
 region = "us-central1"
 zone   = "us-central1-a"
}

resource "google_compute_network_endpoint_group" "neg" {
 provider = "google-beta"

 name         = "neg-%{random_suffix}"
 zone         = "us-central1-a"
 network      = "${google_compute_network.default.self_link}"
 subnetwork   = "${google_compute_subnetwork.default.self_link}"
 default_port = "90"
}

resource "google_compute_network" "default" {
 provider = "google-beta"

 name = "neg-network-%{random_suffix}"
 auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "default" {
 provider = "google-beta"

 name          = "neg-subnetwork-%{random_suffix}"
 ip_cidr_range = "10.0.0.0/16"
 region        = "us-central1"
 network       = "${google_compute_network.default.self_link}"
}

resource "google_compute_instance" "default" {
 provider = "google-beta"

 name         =  "neg-instance1-%{random_suffix}"
 machine_type = "n1-standard-1"

 boot_disk {
   initialize_params{
     image = "${data.google_compute_image.my_image.self_link}"
   }
 }

 network_interface {
   subnetwork = "${google_compute_subnetwork.default.self_link}"
   access_config { }
 }
}

data "google_compute_image" "my_image" {
 provider = "google-beta"

 family  = "debian-9"
 project = "debian-cloud"
}
`, context)
}

type networkEndptTestId struct {
	instance,
	ipAddress,
	port,
	getUrl string
}

// testAccCheckComputeNetworkEndpointExists makes sure the resource with given
// (Terraform) name exists, and returns identifying information about the
// existing endpoint
func testAccCheckComputeNetworkEndpointExists(name string, endpt *networkEndptTestId) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("resource %q not in path %q", name, s.RootModule().Path)
		}

		if rs.Type != "google_compute_network_endpoint" {
			return fmt.Errorf("resource %q has unexpected type %q", name, rs.Type)
		}

		url, err := replaceVarsForTest(rs, "https://www.googleapis.com/compute/beta/projects/{{project}}/zones/{{zone}}/networkEndpointGroups/{{network_endpoint_group}}/listNetworkEndpoints")
		if err != nil {
			return fmt.Errorf("creating URL for getting network endpoint %q failed: %v", name, err)
		}

		testEndpt := &networkEndptTestId{
			instance:  rs.Primary.Attributes["instance"],
			ipAddress: rs.Primary.Attributes["ip_address"],
			port:      rs.Primary.Attributes["port"],
			getUrl:    url,
		}

		found, err := getComputeNetworkEndpoint(name, testEndpt)
		if err != nil {
			return fmt.Errorf("unable to confirm existance of %q: %v", name, err)
		}
		if found == nil {
			return fmt.Errorf("network endpoint for %q not found", name)
		}
		*endpt = *testEndpt
		return nil
	}
}

// testAccCheckComputeNetworkEndpointDestroyed makes sure the endpoint with
// given Terraform resource name and previous information (obtained from Exists)
// was destroyed properly.
func testAccCheckComputeNetworkEndpointDestroyed(name string, destroyed *networkEndptTestId) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		found, err := getComputeNetworkEndpoint(name, destroyed)
		if err != nil {
			return fmt.Errorf("unable to confirm %q was destroyed: %v", name, err)
		}
		if found != nil {
			return fmt.Errorf("network endpoint for %q still exists", name)
		}
		return nil
	}
}

// getComputeNetworkEndpoint takes in a test-only representation of the network
// endpoint so we don't rely on ResourceState for generating the URL of the
// network endpoint after resource state has been removed in Destroy step
func getComputeNetworkEndpoint(name string, endpoint *networkEndptTestId) (map[string]interface{}, error) {
	log.Printf("[DEBUG] getting compute network endpoint %#v", endpoint)
	config := testAccProvider.Meta().(*Config)

	res, err := sendRequest(config, "POST", endpoint.getUrl, nil)
	if err != nil {
		return nil, err
	}

	v, ok := res["items"]
	if !ok || v == nil {
		return nil, nil
	}

	items := v.([]interface{})
	for _, item := range items {
		endptWithHealth := item.(map[string]interface{})
		v, ok := endptWithHealth["networkEndpoint"]
		if !ok || v == nil {
			continue
		}

		endpt := v.(map[string]interface{})
		if endpoint.instance != endpt["instance"] {
			continue
		}
		if endpoint.port != fmt.Sprintf("%v", endpt["port"]) {
			continue
		}
		if endpoint.ipAddress != endpt["ipAddress"] {
			continue
		}

		return endpt, nil
	}
	return nil, nil
}
