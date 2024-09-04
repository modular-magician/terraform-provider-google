// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package dataproc_test

import (
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/envvar"
	dataproctf "github.com/hashicorp/terraform-provider-google/google/services/dataproc"
	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"

	"google.golang.org/api/googleapi"

	"google.golang.org/api/dataproc/v1"
)

func TestAccDataprocCluster_missingZoneGlobalRegion1(t *testing.T) {
	t.Parallel()

	rnd := acctest.RandString(t, 10)
	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckDataproc_missingZoneGlobalRegion1(rnd),
				ExpectError: regexp.MustCompile("zone is mandatory when region is set to 'global'"),
			},
		},
	})
}

func TestAccDataprocCluster_missingZoneGlobalRegion2(t *testing.T) {
	t.Parallel()

	rnd := acctest.RandString(t, 10)
	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckDataproc_missingZoneGlobalRegion2(rnd),
				ExpectError: regexp.MustCompile("zone is mandatory when region is set to 'global'"),
			},
		},
	})
}

func TestAccDataprocCluster_basic(t *testing.T) {
	t.Parallel()

	var cluster dataproc.Cluster
	rnd := acctest.RandString(t, 10)
	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataprocClusterDestroy(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataprocCluster_basic(rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataprocClusterExists(t, "google_dataproc_cluster.basic", &cluster),

					// Default behaviour is for Dataproc to autogen or autodiscover a config bucket
					resource.TestCheckResourceAttrSet("google_dataproc_cluster.basic", "cluster_config.0.bucket"),

					// Default behavior is for Dataproc to not use only internal IP addresses
					resource.TestCheckResourceAttr("google_dataproc_cluster.basic", "cluster_config.0.gce_cluster_config.0.internal_ip_only", "false"),

					// Expect 1 master instances with computed values
					resource.TestCheckResourceAttr("google_dataproc_cluster.basic", "cluster_config.0.master_config.#", "1"),
					resource.TestCheckResourceAttr("google_dataproc_cluster.basic", "cluster_config.0.master_config.0.num_instances", "1"),
					resource.TestCheckResourceAttrSet("google_dataproc_cluster.basic", "cluster_config.0.master_config.0.disk_config.0.boot_disk_size_gb"),
					resource.TestCheckResourceAttrSet("google_dataproc_cluster.basic", "cluster_config.0.master_config.0.disk_config.0.num_local_ssds"),
					resource.TestCheckResourceAttrSet("google_dataproc_cluster.basic", "cluster_config.0.master_config.0.disk_config.0.boot_disk_type"),
					resource.TestCheckResourceAttrSet("google_dataproc_cluster.basic", "cluster_config.0.master_config.0.machine_type"),
					resource.TestCheckResourceAttr("google_dataproc_cluster.basic", "cluster_config.0.master_config.0.instance_names.#", "1"),

					// Expect 2 worker instances with computed values
					resource.TestCheckResourceAttr("google_dataproc_cluster.basic", "cluster_config.0.worker_config.#", "1"),
					resource.TestCheckResourceAttr("google_dataproc_cluster.basic", "cluster_config.0.worker_config.0.num_instances", "2"),
					resource.TestCheckResourceAttrSet("google_dataproc_cluster.basic", "cluster_config.0.worker_config.0.disk_config.0.boot_disk_size_gb"),
					resource.TestCheckResourceAttrSet("google_dataproc_cluster.basic", "cluster_config.0.worker_config.0.disk_config.0.num_local_ssds"),
					resource.TestCheckResourceAttrSet("google_dataproc_cluster.basic", "cluster_config.0.worker_config.0.disk_config.0.boot_disk_type"),
					resource.TestCheckResourceAttrSet("google_dataproc_cluster.basic", "cluster_config.0.worker_config.0.machine_type"),
					resource.TestCheckResourceAttr("google_dataproc_cluster.basic", "cluster_config.0.worker_config.0.instance_names.#", "2"),

					// Expect 0 preemptible worker instances
					resource.TestCheckResourceAttr("google_dataproc_cluster.basic", "cluster_config.0.preemptible_worker_config.#", "1"),
					resource.TestCheckResourceAttr("google_dataproc_cluster.basic", "cluster_config.0.preemptible_worker_config.0.num_instances", "0"),
					resource.TestCheckResourceAttr("google_dataproc_cluster.basic", "cluster_config.0.preemptible_worker_config.0.instance_names.#", "0"),
				),
			},
		},
	})
}

func TestAccDataprocVirtualCluster_basic(t *testing.T) {
	// Currently failing
	acctest.SkipIfVcr(t)
	t.Parallel()

	var cluster dataproc.Cluster
	rnd := acctest.RandString(t, 10)
	pid := envvar.GetTestProjectFromEnv()
	version := "3.1-dataproc-7"
	networkName := acctest.BootstrapSharedTestNetwork(t, "gke-cluster")
	subnetworkName := acctest.BootstrapSubnet(t, "gke-cluster", networkName)

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataprocClusterDestroy(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataprocVirtualCluster_basic(pid, rnd, networkName, subnetworkName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataprocClusterExists(t, "google_dataproc_cluster.virtual_cluster", &cluster),

					// Expect 1 dataproc on gke instances with computed values
					resource.TestCheckResourceAttr("google_dataproc_cluster.virtual_cluster", "virtual_cluster_config.#", "1"),
					resource.TestCheckResourceAttr("google_dataproc_cluster.virtual_cluster", "virtual_cluster_config.0.kubernetes_cluster_config.#", "1"),
					resource.TestCheckResourceAttrSet("google_dataproc_cluster.virtual_cluster", "virtual_cluster_config.0.kubernetes_cluster_config.0.kubernetes_namespace"),
					resource.TestCheckResourceAttr("google_dataproc_cluster.virtual_cluster", "virtual_cluster_config.0.kubernetes_cluster_config.0.kubernetes_software_config.#", "1"),
					resource.TestCheckResourceAttr("google_dataproc_cluster.virtual_cluster", "virtual_cluster_config.0.kubernetes_cluster_config.0.kubernetes_software_config.0.component_version.SPARK", version),

					resource.TestCheckResourceAttr("google_dataproc_cluster.virtual_cluster", "virtual_cluster_config.0.kubernetes_cluster_config.0.gke_cluster_config.#", "1"),
					resource.TestCheckResourceAttrSet("google_dataproc_cluster.virtual_cluster", "virtual_cluster_config.0.kubernetes_cluster_config.0.gke_cluster_config.0.gke_cluster_target"),
					resource.TestCheckResourceAttr("google_dataproc_cluster.virtual_cluster", "virtual_cluster_config.0.kubernetes_cluster_config.0.gke_cluster_config.0.node_pool_target.#", "1"),
					resource.TestCheckResourceAttrSet("google_dataproc_cluster.virtual_cluster", "virtual_cluster_config.0.kubernetes_cluster_config.0.gke_cluster_config.0.node_pool_target.0.node_pool"),
					resource.TestCheckResourceAttr("google_dataproc_cluster.virtual_cluster", "virtual_cluster_config.0.kubernetes_cluster_config.0.gke_cluster_config.0.node_pool_target.0.roles.#", "1"),
					testAccCheckDataprocGkeClusterNodePoolsHaveRoles(&cluster, "DEFAULT"),
				),
			},
		},
	})
}

func TestAccDataprocCluster_withAccelerators(t *testing.T) {
	t.Parallel()

	rnd := acctest.RandString(t, 10)
	var cluster dataproc.Cluster

	project := envvar.GetTestProjectFromEnv()
	acceleratorType := "nvidia-tesla-t4"
	zone := "us-central1-c"
	networkName := acctest.BootstrapSharedTestNetwork(t, "dataproc-cluster")
	subnetworkName := acctest.BootstrapSubnet(t, "dataproc-cluster", networkName)
	acctest.BootstrapFirewallForDataprocSharedNetwork(t, "dataproc-cluster", networkName)

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataprocClusterDestroy(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataprocCluster_withAccelerators(rnd, acceleratorType, zone, subnetworkName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataprocClusterExists(t, "google_dataproc_cluster.accelerated_cluster", &cluster),
					testAccCheckDataprocClusterAccelerator(&cluster, project, 1, 1),
				),
			},
		},
	})
}

func testAccCheckDataprocAuxiliaryNodeGroupAccelerator(cluster *dataproc.Cluster, project string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		expectedUri := fmt.Sprintf("projects/%s/zones/.*/acceleratorTypes/nvidia-tesla-t4", project)
		r := regexp.MustCompile(expectedUri)

		nodeGroup := cluster.Config.AuxiliaryNodeGroups[0].NodeGroup.NodeGroupConfig.Accelerators
		if len(nodeGroup) != 1 {
			return fmt.Errorf("Saw %d nodeGroup accelerator types instead of 1", len(nodeGroup))
		}

		matches := r.FindStringSubmatch(nodeGroup[0].AcceleratorTypeUri)
		if len(matches) != 1 {
			return fmt.Errorf("Saw %s master accelerator type instead of %s", nodeGroup[0].AcceleratorTypeUri, expectedUri)
		}
		return nil
	}
}

func testAccCheckDataprocClusterAccelerator(cluster *dataproc.Cluster, project string, masterCount int, workerCount int) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		expectedUri := fmt.Sprintf("projects/%s/zones/.*/acceleratorTypes/nvidia-tesla-t4", project)
		r := regexp.MustCompile(expectedUri)

		master := cluster.Config.MasterConfig.Accelerators
		if len(master) != 1 {
			return fmt.Errorf("Saw %d master accelerator types instead of 1", len(master))
		}

		if int(master[0].AcceleratorCount) != masterCount {
			return fmt.Errorf("Saw %d master accelerators instead of %d", int(master[0].AcceleratorCount), masterCount)
		}

		matches := r.FindStringSubmatch(master[0].AcceleratorTypeUri)
		if len(matches) != 1 {
			return fmt.Errorf("Saw %s master accelerator type instead of %s", master[0].AcceleratorTypeUri, expectedUri)
		}

		worker := cluster.Config.WorkerConfig.Accelerators
		if len(worker) != 1 {
			return fmt.Errorf("Saw %d worker accelerator types instead of 1", len(worker))
		}

		if int(worker[0].AcceleratorCount) != workerCount {
			return fmt.Errorf("Saw %d worker accelerators instead of %d", int(worker[0].AcceleratorCount), workerCount)
		}

		matches = r.FindStringSubmatch(worker[0].AcceleratorTypeUri)
		if len(matches) != 1 {
			return fmt.Errorf("Saw %s worker accelerator type instead of %s", worker[0].AcceleratorTypeUri, expectedUri)
		}

		return nil
	}
}

func TestAccDataprocCluster_withInternalIpOnlyTrueAndShieldedConfig(t *testing.T) {
	t.Parallel()

	var cluster dataproc.Cluster
	rnd := acctest.RandString(t, 10)
	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataprocClusterDestroy(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataprocCluster_withInternalIpOnlyTrueAndShieldedConfig(rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataprocClusterExists(t, "google_dataproc_cluster.basic", &cluster),

					// Testing behavior for Dataproc to use only internal IP addresses
					resource.TestCheckResourceAttr("google_dataproc_cluster.basic", "cluster_config.0.gce_cluster_config.0.internal_ip_only", "true"),
					resource.TestCheckResourceAttr("google_dataproc_cluster.basic", "cluster_config.0.gce_cluster_config.0.shielded_instance_config.0.enable_integrity_monitoring", "true"),
					resource.TestCheckResourceAttr("google_dataproc_cluster.basic", "cluster_config.0.gce_cluster_config.0.shielded_instance_config.0.enable_secure_boot", "true"),
					resource.TestCheckResourceAttr("google_dataproc_cluster.basic", "cluster_config.0.gce_cluster_config.0.shielded_instance_config.0.enable_vtpm", "true"),
				),
			},
		},
	})
}

func TestAccDataprocCluster_withMetadataAndTags(t *testing.T) {
	t.Parallel()

	var cluster dataproc.Cluster
	rnd := acctest.RandString(t, 10)
	networkName := acctest.BootstrapSharedTestNetwork(t, "dataproc-cluster")
	subnetworkName := acctest.BootstrapSubnet(t, "dataproc-cluster", networkName)
	acctest.BootstrapFirewallForDataprocSharedNetwork(t, "dataproc-cluster", networkName)

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataprocClusterDestroy(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataprocCluster_withMetadataAndTags(rnd, subnetworkName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataprocClusterExists(t, "google_dataproc_cluster.basic", &cluster),

					resource.TestCheckResourceAttr("google_dataproc_cluster.basic", "cluster_config.0.gce_cluster_config.0.metadata.foo", "bar"),
					resource.TestCheckResourceAttr("google_dataproc_cluster.basic", "cluster_config.0.gce_cluster_config.0.metadata.baz", "qux"),
					resource.TestCheckResourceAttr("google_dataproc_cluster.basic", "cluster_config.0.gce_cluster_config.0.tags.#", "4"),
				),
			},
		},
	})
}

func TestAccDataprocCluster_withMinNumInstances(t *testing.T) {
	t.Parallel()

	var cluster dataproc.Cluster
	rnd := acctest.RandString(t, 10)
	networkName := acctest.BootstrapSharedTestNetwork(t, "dataproc-cluster")
	subnetworkName := acctest.BootstrapSubnet(t, "dataproc-cluster", networkName)
	acctest.BootstrapFirewallForDataprocSharedNetwork(t, "dataproc-cluster", networkName)

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataprocClusterDestroy(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataprocCluster_withMinNumInstances(rnd, subnetworkName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataprocClusterExists(t, "google_dataproc_cluster.with_min_num_instances", &cluster),

					resource.TestCheckResourceAttr("google_dataproc_cluster.with_min_num_instances", "cluster_config.0.worker_config.0.min_num_instances", "2"),
				),
			},
		},
	})
}

func TestAccDataprocCluster_withReservationAffinity(t *testing.T) {
	t.Parallel()

	var cluster dataproc.Cluster
	rnd := acctest.RandString(t, 10)
	networkName := acctest.BootstrapSharedTestNetwork(t, "dataproc-cluster")
	subnetworkName := acctest.BootstrapSubnet(t, "dataproc-cluster", networkName)
	acctest.BootstrapFirewallForDataprocSharedNetwork(t, "dataproc-cluster", networkName)

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataprocClusterDestroy(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataprocCluster_withReservationAffinity(rnd, subnetworkName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataprocClusterExists(t, "google_dataproc_cluster.basic", &cluster),

					resource.TestCheckResourceAttr("google_dataproc_cluster.basic", "cluster_config.0.gce_cluster_config.0.reservation_affinity.0.consume_reservation_type", "SPECIFIC_RESERVATION"),
					resource.TestCheckResourceAttr("google_dataproc_cluster.basic", "cluster_config.0.gce_cluster_config.0.reservation_affinity.0.key", "compute.googleapis.com/reservation-name"),
					resource.TestCheckResourceAttr("google_dataproc_cluster.basic", "cluster_config.0.gce_cluster_config.0.reservation_affinity.0.values.#", "1"),
				),
			},
		},
	})
}

func TestAccDataprocCluster_withDataprocMetricConfig(t *testing.T) {
	t.Parallel()

	var cluster dataproc.Cluster
	rnd := acctest.RandString(t, 10)
	networkName := acctest.BootstrapSharedTestNetwork(t, "dataproc-cluster")
	subnetworkName := acctest.BootstrapSubnet(t, "dataproc-cluster", networkName)
	acctest.BootstrapFirewallForDataprocSharedNetwork(t, "dataproc-cluster", networkName)

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataprocClusterDestroy(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataprocCluster_withDataprocMetricConfig(rnd, subnetworkName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataprocClusterExists(t, "google_dataproc_cluster.basic", &cluster),

					resource.TestCheckResourceAttr("google_dataproc_cluster.basic", "cluster_config.0.dataproc_metric_config.0.metrics.#", "2"),

					resource.TestCheckResourceAttr("google_dataproc_cluster.basic", "cluster_config.0.dataproc_metric_config.0.metrics.0.metric_source", "HDFS"),
					resource.TestCheckResourceAttr("google_dataproc_cluster.basic", "cluster_config.0.dataproc_metric_config.0.metrics.0.metric_overrides.#", "1"),
				),
			},
		},
	})
}

func TestAccDataprocCluster_withNodeGroupAffinity(t *testing.T) {
	t.Parallel()

	var cluster dataproc.Cluster
	rnd := acctest.RandString(t, 10)
	networkName := acctest.BootstrapSharedTestNetwork(t, "dataproc-cluster")
	subnetworkName := acctest.BootstrapSubnet(t, "dataproc-cluster", networkName)
	acctest.BootstrapFirewallForDataprocSharedNetwork(t, "dataproc-cluster", networkName)

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataprocClusterDestroy(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataprocCluster_withNodeGroupAffinity(rnd, subnetworkName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataprocClusterExists(t, "google_dataproc_cluster.basic", &cluster),

					resource.TestMatchResourceAttr("google_dataproc_cluster.basic", "cluster_config.0.gce_cluster_config.0.node_group_affinity.0.node_group_uri", regexp.MustCompile("test-nodegroup")),
				),
			},
		},
	})
}

func TestAccDataprocCluster_singleNodeCluster(t *testing.T) {
	t.Parallel()

	rnd := acctest.RandString(t, 10)
	networkName := acctest.BootstrapSharedTestNetwork(t, "dataproc-cluster")
	subnetworkName := acctest.BootstrapSubnet(t, "dataproc-cluster", networkName)
	acctest.BootstrapFirewallForDataprocSharedNetwork(t, "dataproc-cluster", networkName)

	var cluster dataproc.Cluster
	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataprocClusterDestroy(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataprocCluster_singleNodeCluster(rnd, subnetworkName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataprocClusterExists(t, "google_dataproc_cluster.single_node_cluster", &cluster),
					resource.TestCheckResourceAttr("google_dataproc_cluster.single_node_cluster", "cluster_config.0.master_config.0.num_instances", "1"),
					resource.TestCheckResourceAttr("google_dataproc_cluster.single_node_cluster", "cluster_config.0.worker_config.0.num_instances", "0"),

					// We set the "dataproc:dataproc.allow.zero.workers" override property.
					// GCP should populate the 'properties' value with this value, as well as many others
					resource.TestCheckResourceAttrSet("google_dataproc_cluster.single_node_cluster", "cluster_config.0.software_config.0.properties.%"),
				),
			},
		},
	})
}

func TestAccDataprocCluster_updatable(t *testing.T) {
	t.Parallel()

	rnd := acctest.RandString(t, 10)
	var cluster dataproc.Cluster

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataprocClusterDestroy(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataprocCluster_updatable(rnd, 2, 1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataprocClusterExists(t, "google_dataproc_cluster.updatable", &cluster),
					resource.TestCheckResourceAttr("google_dataproc_cluster.updatable", "cluster_config.0.master_config.0.num_instances", "1"),
					resource.TestCheckResourceAttr("google_dataproc_cluster.updatable", "cluster_config.0.worker_config.0.num_instances", "2"),
					resource.TestCheckResourceAttr("google_dataproc_cluster.updatable", "cluster_config.0.preemptible_worker_config.0.num_instances", "1")),
			},
			{
				Config: testAccDataprocCluster_updatable(rnd, 2, 0),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataprocClusterExists(t, "google_dataproc_cluster.updatable", &cluster),
					resource.TestCheckResourceAttr("google_dataproc_cluster.updatable", "cluster_config.0.master_config.0.num_instances", "1"),
					resource.TestCheckResourceAttr("google_dataproc_cluster.updatable", "cluster_config.0.worker_config.0.num_instances", "2"),
					resource.TestCheckResourceAttr("google_dataproc_cluster.updatable", "cluster_config.0.preemptible_worker_config.0.num_instances", "0")),
			},
			{
				Config: testAccDataprocCluster_updatable(rnd, 3, 2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("google_dataproc_cluster.updatable", "cluster_config.0.master_config.0.num_instances", "1"),
					resource.TestCheckResourceAttr("google_dataproc_cluster.updatable", "cluster_config.0.worker_config.0.num_instances", "3"),
					resource.TestCheckResourceAttr("google_dataproc_cluster.updatable", "cluster_config.0.preemptible_worker_config.0.num_instances", "2")),
			},
		},
	})
}

func TestAccDataprocCluster_nonPreemptibleSecondary(t *testing.T) {
	t.Parallel()

	rnd := acctest.RandString(t, 10)
	networkName := acctest.BootstrapSharedTestNetwork(t, "dataproc-cluster")
	subnetworkName := acctest.BootstrapSubnet(t, "dataproc-cluster", networkName)
	acctest.BootstrapFirewallForDataprocSharedNetwork(t, "dataproc-cluster", networkName)
	var cluster dataproc.Cluster

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataprocClusterDestroy(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataprocCluster_nonPreemptibleSecondary(rnd, subnetworkName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataprocClusterExists(t, "google_dataproc_cluster.non_preemptible_secondary", &cluster),
					resource.TestCheckResourceAttr("google_dataproc_cluster.non_preemptible_secondary", "cluster_config.0.preemptible_worker_config.0.preemptibility", "NON_PREEMPTIBLE"),
				),
			},
		},
	})
}

func TestAccDataprocCluster_spotSecondary(t *testing.T) {
	t.Parallel()

	rnd := acctest.RandString(t, 10)
	networkName := acctest.BootstrapSharedTestNetwork(t, "dataproc-cluster")
	subnetworkName := acctest.BootstrapSubnet(t, "dataproc-cluster", networkName)
	acctest.BootstrapFirewallForDataprocSharedNetwork(t, "dataproc-cluster", networkName)
	var cluster dataproc.Cluster

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataprocClusterDestroy(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataprocCluster_spotSecondary(rnd, subnetworkName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataprocClusterExists(t, "google_dataproc_cluster.spot_secondary", &cluster),
					resource.TestCheckResourceAttr("google_dataproc_cluster.spot_secondary", "cluster_config.0.preemptible_worker_config.0.preemptibility", "SPOT"),
				),
			},
		},
	})
}

func TestAccDataprocCluster_spotWithInstanceFlexibilityPolicy(t *testing.T) {
	t.Parallel()

	rnd := acctest.RandString(t, 10)
	var cluster dataproc.Cluster
	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataprocClusterDestroy(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataprocCluster_spotWithInstanceFlexibilityPolicy(rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataprocClusterExists(t, "google_dataproc_cluster.spot_with_instance_flexibility_policy", &cluster),
					resource.TestCheckResourceAttr("google_dataproc_cluster.spot_with_instance_flexibility_policy", "cluster_config.0.preemptible_worker_config.0.preemptibility", "SPOT"),
					resource.TestCheckResourceAttr("google_dataproc_cluster.spot_with_instance_flexibility_policy", "cluster_config.0.preemptible_worker_config.0.instance_flexibility_policy.0.instance_selection_list.0.machine_types.0", "n2d-standard-2"),
					resource.TestCheckResourceAttr("google_dataproc_cluster.spot_with_instance_flexibility_policy", "cluster_config.0.preemptible_worker_config.0.instance_flexibility_policy.0.instance_selection_list.0.rank", "3"),
				),
			},
		},
	})
}

func TestAccDataprocCluster_spotWithAuxiliaryNodeGroups(t *testing.T) {
	t.Parallel()

	project := envvar.GetTestProjectFromEnv()
	rnd := acctest.RandString(t, 10)
	var cluster dataproc.Cluster
	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataprocClusterDestroy(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataprocCluster_withAuxiliaryNodeGroups(rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataprocClusterExists(t, "google_dataproc_cluster.with_auxiliary_node_groups", &cluster),
					resource.TestCheckResourceAttr("google_dataproc_cluster.with_auxiliary_node_groups", "cluster_config.0.auxiliary_node_groups.0.node_group.0.roles.0", "DRIVER"),
					resource.TestCheckResourceAttr("google_dataproc_cluster.with_auxiliary_node_groups", "cluster_config.0.auxiliary_node_groups.0.node_group.0.node_group_config.0.num_instances", "2"),
					resource.TestCheckResourceAttr("google_dataproc_cluster.with_auxiliary_node_groups", "cluster_config.0.auxiliary_node_groups.0.node_group.0.node_group_config.0.machine_type", "n1-standard-2"),
					resource.TestCheckResourceAttr("google_dataproc_cluster.with_auxiliary_node_groups", "cluster_config.0.auxiliary_node_groups.0.node_group.0.node_group_config.0.min_cpu_platform", "Intel Haswell"),
					resource.TestCheckResourceAttr("google_dataproc_cluster.with_auxiliary_node_groups", "cluster_config.0.auxiliary_node_groups.0.node_group.0.node_group_config.0.disk_config.0.boot_disk_size_gb", "35"),
					resource.TestCheckResourceAttr("google_dataproc_cluster.with_auxiliary_node_groups", "cluster_config.0.auxiliary_node_groups.0.node_group.0.node_group_config.0.disk_config.0.boot_disk_type", "pd-standard"),
					resource.TestCheckResourceAttr("google_dataproc_cluster.with_auxiliary_node_groups", "cluster_config.0.auxiliary_node_groups.0.node_group.0.node_group_config.0.disk_config.0.num_local_ssds", "1"),
					resource.TestCheckResourceAttr("google_dataproc_cluster.with_auxiliary_node_groups", "cluster_config.0.auxiliary_node_groups.0.node_group.0.node_group_config.0.disk_config.0.local_ssd_interface", "nvme"),
					resource.TestCheckResourceAttr("google_dataproc_cluster.with_auxiliary_node_groups", "cluster_config.0.auxiliary_node_groups.0.node_group.0.node_group_config.0.accelerators.0.accelerator_count", "1"),
					resource.TestCheckResourceAttr("google_dataproc_cluster.with_auxiliary_node_groups", "cluster_config.0.auxiliary_node_groups.0.node_group_id", "node-group-id"),
					testAccCheckDataprocAuxiliaryNodeGroupAccelerator(&cluster, project),
				),
			},
		},
	})
}

func TestAccDataprocCluster_withStagingBucket(t *testing.T) {
	t.Parallel()

	rnd := acctest.RandString(t, 10)
	var cluster dataproc.Cluster
	clusterName := fmt.Sprintf("tf-test-dproc-%s", rnd)
	bucketName := fmt.Sprintf("%s-bucket", clusterName)
	networkName := acctest.BootstrapSharedTestNetwork(t, "dataproc-cluster")
	subnetworkName := acctest.BootstrapSubnet(t, "dataproc-cluster", networkName)
	acctest.BootstrapFirewallForDataprocSharedNetwork(t, "dataproc-cluster", networkName)

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataprocClusterDestroy(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataprocCluster_withStagingBucketAndCluster(clusterName, bucketName, subnetworkName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataprocClusterExists(t, "google_dataproc_cluster.with_bucket", &cluster),
					resource.TestCheckResourceAttr("google_dataproc_cluster.with_bucket", "cluster_config.0.staging_bucket", bucketName),
					resource.TestCheckResourceAttr("google_dataproc_cluster.with_bucket", "cluster_config.0.bucket", bucketName)),
			},
			{
				// Simulate destroy of cluster by removing it from definition,
				// but leaving the storage bucket (should not be auto deleted)
				Config: testAccDataprocCluster_withStagingBucketOnly(bucketName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataprocStagingBucketExists(t, bucketName),
				),
			},
		},
	})
}

func TestAccDataprocCluster_withTempBucket(t *testing.T) {
	t.Parallel()

	rnd := acctest.RandString(t, 10)
	var cluster dataproc.Cluster
	clusterName := fmt.Sprintf("tf-test-dproc-%s", rnd)
	bucketName := fmt.Sprintf("%s-temp-bucket", clusterName)
	networkName := acctest.BootstrapSharedTestNetwork(t, "dataproc-cluster")
	subnetworkName := acctest.BootstrapSubnet(t, "dataproc-cluster", networkName)
	acctest.BootstrapFirewallForDataprocSharedNetwork(t, "dataproc-cluster", networkName)

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataprocClusterDestroy(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataprocCluster_withTempBucketAndCluster(clusterName, bucketName, subnetworkName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataprocClusterExists(t, "google_dataproc_cluster.with_bucket", &cluster),
					resource.TestCheckResourceAttr("google_dataproc_cluster.with_bucket", "cluster_config.0.temp_bucket", bucketName)),
			},
			{
				// Simulate destroy of cluster by removing it from definition,
				// but leaving the temp bucket (should not be auto deleted)
				Config: testAccDataprocCluster_withTempBucketOnly(bucketName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataprocTempBucketExists(t, bucketName),
				),
			},
		},
	})
}

func TestAccDataprocCluster_withInitAction(t *testing.T) {
	t.Parallel()

	rnd := acctest.RandString(t, 10)
	var cluster dataproc.Cluster
	bucketName := fmt.Sprintf("tf-test-dproc-%s-init-bucket", rnd)
	objectName := "msg.txt"
	networkName := acctest.BootstrapSharedTestNetwork(t, "dataproc-cluster")
	subnetworkName := acctest.BootstrapSubnet(t, "dataproc-cluster", networkName)
	acctest.BootstrapFirewallForDataprocSharedNetwork(t, "dataproc-cluster", networkName)

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataprocClusterDestroy(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataprocCluster_withInitAction(rnd, bucketName, objectName, subnetworkName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataprocClusterExists(t, "google_dataproc_cluster.with_init_action", &cluster),
					resource.TestCheckResourceAttr("google_dataproc_cluster.with_init_action", "cluster_config.0.initialization_action.#", "2"),
					resource.TestCheckResourceAttr("google_dataproc_cluster.with_init_action", "cluster_config.0.initialization_action.0.timeout_sec", "500"),
					testAccCheckDataprocClusterInitActionSucceeded(t, bucketName, objectName),
				),
			},
		},
	})
}

func TestAccDataprocCluster_withConfigOverrides(t *testing.T) {
	t.Parallel()

	rnd := acctest.RandString(t, 10)
	var cluster dataproc.Cluster
	networkName := acctest.BootstrapSharedTestNetwork(t, "dataproc-cluster")
	subnetworkName := acctest.BootstrapSubnet(t, "dataproc-cluster", networkName)
	acctest.BootstrapFirewallForDataprocSharedNetwork(t, "dataproc-cluster", networkName)

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataprocClusterDestroy(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataprocCluster_withConfigOverrides(rnd, subnetworkName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataprocClusterExists(t, "google_dataproc_cluster.with_config_overrides", &cluster),
					validateDataprocCluster_withConfigOverrides("google_dataproc_cluster.with_config_overrides", &cluster),
				),
			},
		},
	})
}

func TestAccDataprocCluster_withServiceAcc(t *testing.T) {
	t.Parallel()

	sa := "a" + acctest.RandString(t, 10)
	saEmail := fmt.Sprintf("%s@%s.iam.gserviceaccount.com", sa, envvar.GetTestProjectFromEnv())
	rnd := acctest.RandString(t, 10)
	networkName := acctest.BootstrapSharedTestNetwork(t, "dataproc-cluster")
	subnetworkName := acctest.BootstrapSubnet(t, "dataproc-cluster", networkName)
	acctest.BootstrapFirewallForDataprocSharedNetwork(t, "dataproc-cluster", networkName)

	var cluster dataproc.Cluster

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		CheckDestroy: testAccCheckDataprocClusterDestroy(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataprocCluster_withServiceAcc(sa, rnd, subnetworkName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataprocClusterExists(
						t, "google_dataproc_cluster.with_service_account", &cluster),
					testAccCheckDataprocClusterHasServiceScopes(t, &cluster,
						"https://www.googleapis.com/auth/cloud.useraccounts.readonly",
						"https://www.googleapis.com/auth/devstorage.read_write",
						"https://www.googleapis.com/auth/logging.write",
						"https://www.googleapis.com/auth/monitoring",
					),
					resource.TestCheckResourceAttr("google_dataproc_cluster.with_service_account", "cluster_config.0.gce_cluster_config.0.service_account", saEmail),
				),
			},
		},
	})
}

func TestAccDataprocCluster_withImageVersion(t *testing.T) {
	t.Parallel()

	rnd := acctest.RandString(t, 10)
	version := "2.0.35-debian10"
	networkName := acctest.BootstrapSharedTestNetwork(t, "dataproc-cluster")
	subnetworkName := acctest.BootstrapSubnet(t, "dataproc-cluster", networkName)
	acctest.BootstrapFirewallForDataprocSharedNetwork(t, "dataproc-cluster", networkName)

	var cluster dataproc.Cluster
	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataprocClusterDestroy(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataprocCluster_withImageVersion(rnd, version, subnetworkName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataprocClusterExists(t, "google_dataproc_cluster.with_image_version", &cluster),
					resource.TestCheckResourceAttr("google_dataproc_cluster.with_image_version", "cluster_config.0.software_config.0.image_version", version),
				),
			},
		},
	})
}

func TestAccDataprocCluster_withOptionalComponents(t *testing.T) {
	t.Parallel()

	rnd := acctest.RandString(t, 10)
	networkName := acctest.BootstrapSharedTestNetwork(t, "dataproc-cluster")
	subnetworkName := acctest.BootstrapSubnet(t, "dataproc-cluster", networkName)
	acctest.BootstrapFirewallForDataprocSharedNetwork(t, "dataproc-cluster", networkName)
	var cluster dataproc.Cluster

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataprocClusterDestroy(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataprocCluster_withOptionalComponents(rnd, subnetworkName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataprocClusterExists(t, "google_dataproc_cluster.with_opt_components", &cluster),
					testAccCheckDataprocClusterHasOptionalComponents(&cluster, "ZOOKEEPER", "DOCKER"),
				),
			},
		},
	})
}

func TestAccDataprocCluster_withLifecycleConfigIdleDeleteTtl(t *testing.T) {
	t.Parallel()

	rnd := acctest.RandString(t, 10)
	networkName := acctest.BootstrapSharedTestNetwork(t, "dataproc-cluster")
	subnetworkName := acctest.BootstrapSubnet(t, "dataproc-cluster", networkName)
	acctest.BootstrapFirewallForDataprocSharedNetwork(t, "dataproc-cluster", networkName)
	var cluster dataproc.Cluster

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataprocClusterDestroy(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataprocCluster_withLifecycleConfigIdleDeleteTtl(rnd, "600s", subnetworkName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataprocClusterExists(t, "google_dataproc_cluster.with_lifecycle_config", &cluster),
				),
			},
			{
				Config: testAccDataprocCluster_withLifecycleConfigIdleDeleteTtl(rnd, "610s", subnetworkName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataprocClusterExists(t, "google_dataproc_cluster.with_lifecycle_config", &cluster),
				),
			},
		},
	})
}

func TestAccDataprocCluster_withLifecycleConfigAutoDeletion(t *testing.T) {
	// Uses time.Now
	acctest.SkipIfVcr(t)
	t.Parallel()

	rnd := acctest.RandString(t, 10)
	now := time.Now()
	fmtString := "2006-01-02T15:04:05.072Z"
	networkName := acctest.BootstrapSharedTestNetwork(t, "dataproc-cluster")
	subnetworkName := acctest.BootstrapSubnet(t, "dataproc-cluster", networkName)
	acctest.BootstrapFirewallForDataprocSharedNetwork(t, "dataproc-cluster", networkName)

	var cluster dataproc.Cluster
	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataprocClusterDestroy(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataprocCluster_withLifecycleConfigAutoDeletionTime(rnd, now.Add(time.Hour*10).Format(fmtString), subnetworkName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataprocClusterExists(t, "google_dataproc_cluster.with_lifecycle_config", &cluster),
				),
			},
			{
				Config: testAccDataprocCluster_withLifecycleConfigAutoDeletionTime(rnd, now.Add(time.Hour*20).Format(fmtString), subnetworkName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataprocClusterExists(t, "google_dataproc_cluster.with_lifecycle_config", &cluster),
				),
			},
		},
	})
}

func TestAccDataprocCluster_withLabels(t *testing.T) {
	t.Parallel()

	rnd := acctest.RandString(t, 10)
	networkName := acctest.BootstrapSharedTestNetwork(t, "dataproc-cluster")
	subnetworkName := acctest.BootstrapSubnet(t, "dataproc-cluster", networkName)
	acctest.BootstrapFirewallForDataprocSharedNetwork(t, "dataproc-cluster", networkName)
	var cluster dataproc.Cluster

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataprocClusterDestroy(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataprocCluster_withoutLabels(rnd, subnetworkName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataprocClusterExists(t, "google_dataproc_cluster.with_labels", &cluster),

					resource.TestCheckNoResourceAttr("google_dataproc_cluster.with_labels", "labels.%"),
					// We don't provide any, but GCP adds three and goog-dataproc-autozone is added internally, so expect 4.
					resource.TestCheckResourceAttr("google_dataproc_cluster.with_labels", "effective_labels.%", "4"),
				),
			},
			{
				Config: testAccDataprocCluster_withLabels(rnd, subnetworkName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataprocClusterExists(t, "google_dataproc_cluster.with_labels", &cluster),

					resource.TestCheckResourceAttr("google_dataproc_cluster.with_labels", "labels.%", "1"),
					resource.TestCheckResourceAttr("google_dataproc_cluster.with_labels", "labels.key1", "value1"),
					// We only provide one, but GCP adds three and goog-dataproc-autozone is added internally, so expect 5.
					resource.TestCheckResourceAttr("google_dataproc_cluster.with_labels", "effective_labels.%", "5"),
					resource.TestCheckResourceAttr("google_dataproc_cluster.with_labels", "effective_labels.key1", "value1"),
				),
			},
			{
				Config: testAccDataprocCluster_withLabelsUpdate(rnd, subnetworkName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataprocClusterExists(t, "google_dataproc_cluster.with_labels", &cluster),

					// We only provide two, so expect 2.
					resource.TestCheckResourceAttr("google_dataproc_cluster.with_labels", "labels.%", "1"),
					resource.TestCheckResourceAttr("google_dataproc_cluster.with_labels", "labels.key2", "value2"),
				),
			},
			{
				Config: testAccDataprocCluster_withoutLabels(rnd, subnetworkName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataprocClusterExists(t, "google_dataproc_cluster.with_labels", &cluster),

					resource.TestCheckNoResourceAttr("google_dataproc_cluster.with_labels", "labels.%"),
					// We don't provide any, but GCP adds three and goog-dataproc-autozone is added internally, so expect 4.
					resource.TestCheckResourceAttr("google_dataproc_cluster.with_labels", "effective_labels.%", "4"),
				),
			},
		},
	})
}

func TestAccDataprocCluster_withNetworkRefs(t *testing.T) {
	// Multiple fine-grained resources
	acctest.SkipIfVcr(t)
	t.Parallel()

	var c1, c2 dataproc.Cluster
	rnd := acctest.RandString(t, 10)
	netName := fmt.Sprintf(`dproc-cluster-test-%s-net`, rnd)
	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataprocClusterDestroy(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataprocCluster_withNetworkRefs(rnd, netName),
				Check: resource.ComposeTestCheckFunc(
					// successful creation of the clusters is good enough to assess it worked
					testAccCheckDataprocClusterExists(t, "google_dataproc_cluster.with_net_ref_by_url", &c1),
					testAccCheckDataprocClusterExists(t, "google_dataproc_cluster.with_net_ref_by_name", &c2),
				),
			},
		},
	})
}

func TestAccDataprocCluster_withEndpointConfig(t *testing.T) {
	t.Parallel()

	var cluster dataproc.Cluster
	rnd := acctest.RandString(t, 10)
	networkName := acctest.BootstrapSharedTestNetwork(t, "dataproc-cluster")
	subnetworkName := acctest.BootstrapSubnet(t, "dataproc-cluster", networkName)
	acctest.BootstrapFirewallForDataprocSharedNetwork(t, "dataproc-cluster", networkName)

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataprocClusterDestroy(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataprocCluster_withEndpointConfig(rnd, subnetworkName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataprocClusterExists(t, "google_dataproc_cluster.with_endpoint_config", &cluster),
					resource.TestCheckResourceAttr("google_dataproc_cluster.with_endpoint_config", "cluster_config.0.endpoint_config.0.enable_http_port_access", "true"),
				),
			},
		},
	})
}

func TestAccDataprocCluster_KMS(t *testing.T) {
	t.Parallel()

	rnd := acctest.RandString(t, 10)
	kms := acctest.BootstrapKMSKey(t)
	networkName := acctest.BootstrapSharedTestNetwork(t, "dataproc-cluster")
	subnetworkName := acctest.BootstrapSubnet(t, "dataproc-cluster", networkName)
	acctest.BootstrapFirewallForDataprocSharedNetwork(t, "dataproc-cluster", networkName)

	if acctest.BootstrapPSARole(t, "service-", "compute-system", "roles/cloudkms.cryptoKeyEncrypterDecrypter") {
		t.Fatal("Stopping the test because a role was added to the policy.")
	}

	var cluster dataproc.Cluster
	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataprocClusterDestroy(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataprocCluster_KMS(rnd, kms.CryptoKey.Name, subnetworkName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataprocClusterExists(t, "google_dataproc_cluster.kms", &cluster),
				),
			},
		},
	})
}

func TestFlattenSecurityConfig(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		sc       *dataproc.SecurityConfig
		expected []map[string]interface{}
	}{
		"nil": {
			sc:       nil,
			expected: nil,
		},
		"empty": {
			sc: &dataproc.SecurityConfig{},
			expected: []map[string]interface{}{
				{},
			},
		},
		"with kerberos": {
			sc: &dataproc.SecurityConfig{
				KerberosConfig: &dataproc.KerberosConfig{},
			},
			expected: []map[string]interface{}{
				{
					"kerberos_config": []map[string]interface{}{
						{
							"cross_realm_trust_admin_server":        "",
							"cross_realm_trust_kdc":                 "",
							"cross_realm_trust_realm":               "",
							"cross_realm_trust_shared_password_uri": "",
							"enable_kerberos":                       false,
							"kdc_db_key_uri":                        "",
							"key_password_uri":                      "",
							"keystore_password_uri":                 "",
							"keystore_uri":                          "",
							"kms_key_uri":                           "",
							"realm":                                 "",
							"root_principal_password_uri":           "",
							"tgt_lifetime_hours":                    int64(0),
							"truststore_password_uri":               "",
							"truststore_uri":                        "",
						},
					},
				},
			},
		},
	}
	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			actual := dataproctf.FlattenSecurityConfig(nil, tc.sc)
			if diff := cmp.Diff(tc.expected, actual); diff != "" {
				t.Errorf("Unexpected flattened security config. Diff (-want +got): %s", diff)
			}
		})
	}
}

func TestAccDataprocCluster_withKerberos(t *testing.T) {
	t.Parallel()

	rnd := acctest.RandString(t, 10)
	kms := acctest.BootstrapKMSKey(t)
	networkName := acctest.BootstrapSharedTestNetwork(t, "dataproc-cluster")
	subnetworkName := acctest.BootstrapSubnet(t, "dataproc-cluster", networkName)
	acctest.BootstrapFirewallForDataprocSharedNetwork(t, "dataproc-cluster", networkName)

	var cluster dataproc.Cluster
	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataprocClusterDestroy(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataprocCluster_withKerberos(rnd, kms.CryptoKey.Name, subnetworkName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataprocClusterExists(t, "google_dataproc_cluster.kerb", &cluster),
				),
			},
		},
	})
}

func TestAccDataprocCluster_withSecurityConfig_withoutKerberosConfig(t *testing.T) {
	t.Parallel()

	rnd := acctest.RandString(t, 10)
	networkName := acctest.BootstrapSharedTestNetwork(t, "dataproc-cluster")
	subnetworkName := acctest.BootstrapSubnet(t, "dataproc-cluster", networkName)
	acctest.BootstrapFirewallForDataprocSharedNetwork(t, "dataproc-cluster", networkName)

	var cluster dataproc.Cluster
	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataprocClusterDestroy(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataprocCluster_withSecurityConfig_withoutKerberosConfig(rnd, subnetworkName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataprocClusterExists(t, "google_dataproc_cluster.nokerb", &cluster),
				),
			},
		},
	})
}

func TestAccDataprocCluster_withAutoscalingPolicy(t *testing.T) {
	t.Parallel()

	rnd := acctest.RandString(t, 10)
	networkName := acctest.BootstrapSharedTestNetwork(t, "dataproc-cluster")
	subnetworkName := acctest.BootstrapSubnet(t, "dataproc-cluster", networkName)
	acctest.BootstrapFirewallForDataprocSharedNetwork(t, "dataproc-cluster", networkName)

	var cluster dataproc.Cluster
	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataprocClusterDestroy(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataprocCluster_withAutoscalingPolicy(rnd, subnetworkName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataprocClusterExists(t, "google_dataproc_cluster.basic", &cluster),
					testAccCheckDataprocClusterAutoscaling(t, &cluster, true),
				),
			},
			{
				Config: testAccDataprocCluster_removeAutoscalingPolicy(rnd, subnetworkName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataprocClusterExists(t, "google_dataproc_cluster.basic", &cluster),
					testAccCheckDataprocClusterAutoscaling(t, &cluster, false),
				),
			},
		},
	})
}

func TestAccDataprocCluster_withMetastoreConfig(t *testing.T) {
	t.Parallel()

	pid := envvar.GetTestProjectFromEnv()
	basicServiceId := "tf-test-metastore-srv-" + acctest.RandString(t, 10)
	updateServiceId := "tf-test-metastore-srv-update-" + acctest.RandString(t, 10)
	msName_basic := fmt.Sprintf("projects/%s/locations/us-central1/services/%s", pid, basicServiceId)
	msName_update := fmt.Sprintf("projects/%s/locations/us-central1/services/%s", pid, updateServiceId)

	var cluster dataproc.Cluster
	clusterName := "tf-test-" + acctest.RandString(t, 10)
	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataprocClusterDestroy(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataprocCluster_withMetastoreConfig(clusterName, basicServiceId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataprocClusterExists(t, "google_dataproc_cluster.with_metastore_config", &cluster),
					resource.TestCheckResourceAttr("google_dataproc_cluster.with_metastore_config", "cluster_config.0.metastore_config.0.dataproc_metastore_service", msName_basic),
				),
			},
			{
				Config: testAccDataprocCluster_withMetastoreConfig_update(clusterName, updateServiceId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataprocClusterExists(t, "google_dataproc_cluster.with_metastore_config", &cluster),
					resource.TestCheckResourceAttr("google_dataproc_cluster.with_metastore_config", "cluster_config.0.metastore_config.0.dataproc_metastore_service", msName_update),
				),
			},
		},
	})
}

func testAccCheckDataprocClusterDestroy(t *testing.T) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		config := acctest.GoogleProviderConfig(t)

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "google_dataproc_cluster" {
				continue
			}

			if rs.Primary.ID == "" {
				return fmt.Errorf("Unable to verify delete of dataproc cluster, ID is empty")
			}

			attributes := rs.Primary.Attributes
			project, err := acctest.GetTestProject(rs.Primary, config)
			if err != nil {
				return err
			}

			parts := strings.Split(rs.Primary.ID, "/")
			clusterId := parts[len(parts)-1]
			_, err = config.NewDataprocClient(config.UserAgent).Projects.Regions.Clusters.Get(
				project, attributes["region"], clusterId).Do()

			if err != nil {
				if gerr, ok := err.(*googleapi.Error); ok && gerr.Code == http.StatusNotFound {
					return nil
				} else if ok {
					return fmt.Errorf("Error validating cluster deleted. Code: %d. Message: %s", gerr.Code, gerr.Message)
				}
				return fmt.Errorf("Error validating cluster deleted. %s", err.Error())
			}
			return fmt.Errorf("Dataproc cluster still exists")
		}

		return nil
	}
}

func testAccCheckDataprocClusterHasServiceScopes(t *testing.T, cluster *dataproc.Cluster, scopes ...string) func(s *terraform.State) error {
	return func(s *terraform.State) error {

		if !reflect.DeepEqual(scopes, cluster.Config.GceClusterConfig.ServiceAccountScopes) {
			return fmt.Errorf("Cluster does not contain expected set of service account scopes : %v : instead %v",
				scopes, cluster.Config.GceClusterConfig.ServiceAccountScopes)
		}
		return nil
	}
}

func testAccCheckDataprocClusterAutoscaling(t *testing.T, cluster *dataproc.Cluster, expectAutoscaling bool) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		if cluster.Config.AutoscalingConfig == nil && expectAutoscaling {
			return fmt.Errorf("Cluster does not contain AutoscalingConfig, expected it would")
		} else if cluster.Config.AutoscalingConfig != nil && !expectAutoscaling {
			return fmt.Errorf("Cluster contains AutoscalingConfig, expected it not to")
		}

		return nil
	}
}

func validateBucketExists(bucket string, config *transport_tpg.Config) (bool, error) {
	_, err := config.NewStorageClient(config.UserAgent).Buckets.Get(bucket).Do()

	if err != nil {
		if gerr, ok := err.(*googleapi.Error); ok && gerr.Code == http.StatusNotFound {
			return false, nil
		} else if ok {
			return false, fmt.Errorf("Error validating bucket exists: http code error : %d, http message error: %s", gerr.Code, gerr.Message)
		}
		return false, fmt.Errorf("Error validating bucket exists: %s", err.Error())
	}
	return true, nil
}

func testAccCheckDataprocStagingBucketExists(t *testing.T, bucketName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		config := acctest.GoogleProviderConfig(t)

		exists, err := validateBucketExists(bucketName, config)
		if err != nil {
			return err
		}
		if !exists {
			return fmt.Errorf("Staging Bucket %s does not exist", bucketName)
		}
		return nil
	}
}

func testAccCheckDataprocTempBucketExists(t *testing.T, bucketName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		config := acctest.GoogleProviderConfig(t)

		exists, err := validateBucketExists(bucketName, config)
		if err != nil {
			return err
		}
		if !exists {
			return fmt.Errorf("Temp Bucket %s does not exist", bucketName)
		}
		return nil
	}
}

func testAccCheckDataprocClusterHasOptionalComponents(cluster *dataproc.Cluster, components ...string) func(s *terraform.State) error {
	return func(s *terraform.State) error {

		if !reflect.DeepEqual(components, cluster.Config.SoftwareConfig.OptionalComponents) {
			return fmt.Errorf("Cluster does not contain expected optional components : %v : instead %v",
				components, cluster.Config.SoftwareConfig.OptionalComponents)
		}
		return nil
	}
}

func testAccCheckDataprocClusterInitActionSucceeded(t *testing.T, bucket, object string) resource.TestCheckFunc {

	// The init script will have created an object in the specified bucket.
	// Ensure it exists
	return func(s *terraform.State) error {
		config := acctest.GoogleProviderConfig(t)
		_, err := config.NewStorageClient(config.UserAgent).Objects.Get(bucket, object).Do()
		if err != nil {
			return fmt.Errorf("Unable to verify init action success: Error reading object %s in bucket %s: %v", object, bucket, err)
		}

		return nil
	}
}

func validateDataprocCluster_withConfigOverrides(n string, cluster *dataproc.Cluster) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		type tfAndGCPTestField struct {
			tfAttr       string
			expectedVal  string
			actualGCPVal string
		}

		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Terraform resource Not found: %s", n)
		}

		if cluster.Config.MasterConfig == nil || cluster.Config.WorkerConfig == nil || cluster.Config.SecondaryWorkerConfig == nil {
			return fmt.Errorf("Master/Worker/SecondaryConfig values not set in GCP, expecting values")
		}

		clusterTests := []tfAndGCPTestField{
			{"cluster_config.0.master_config.0.num_instances", "3", strconv.Itoa(int(cluster.Config.MasterConfig.NumInstances))},
			{"cluster_config.0.master_config.0.disk_config.0.boot_disk_size_gb", "35", strconv.Itoa(int(cluster.Config.MasterConfig.DiskConfig.BootDiskSizeGb))},
			{"cluster_config.0.master_config.0.disk_config.0.num_local_ssds", "0", strconv.Itoa(int(cluster.Config.MasterConfig.DiskConfig.NumLocalSsds))},
			{"cluster_config.0.master_config.0.disk_config.0.boot_disk_type", "pd-ssd", cluster.Config.MasterConfig.DiskConfig.BootDiskType},
			{"cluster_config.0.master_config.0.disk_config.0.local_ssd_interface", "nvme", cluster.Config.MasterConfig.DiskConfig.LocalSsdInterface},
			{"cluster_config.0.master_config.0.machine_type", "n1-standard-2", tpgresource.GetResourceNameFromSelfLink(cluster.Config.MasterConfig.MachineTypeUri)},
			{"cluster_config.0.master_config.0.instance_names.#", "3", strconv.Itoa(len(cluster.Config.MasterConfig.InstanceNames))},
			{"cluster_config.0.master_config.0.min_cpu_platform", "Intel Skylake", cluster.Config.MasterConfig.MinCpuPlatform},

			{"cluster_config.0.worker_config.0.num_instances", "3", strconv.Itoa(int(cluster.Config.WorkerConfig.NumInstances))},
			{"cluster_config.0.worker_config.0.disk_config.0.boot_disk_size_gb", "35", strconv.Itoa(int(cluster.Config.WorkerConfig.DiskConfig.BootDiskSizeGb))},
			{"cluster_config.0.worker_config.0.disk_config.0.num_local_ssds", "1", strconv.Itoa(int(cluster.Config.WorkerConfig.DiskConfig.NumLocalSsds))},
			{"cluster_config.0.worker_config.0.disk_config.0.boot_disk_type", "pd-standard", cluster.Config.WorkerConfig.DiskConfig.BootDiskType},
			{"cluster_config.0.worker_config.0.disk_config.0.local_ssd_interface", "scsi", cluster.Config.WorkerConfig.DiskConfig.LocalSsdInterface},
			{"cluster_config.0.worker_config.0.machine_type", "n1-standard-2", tpgresource.GetResourceNameFromSelfLink(cluster.Config.WorkerConfig.MachineTypeUri)},
			{"cluster_config.0.worker_config.0.instance_names.#", "3", strconv.Itoa(len(cluster.Config.WorkerConfig.InstanceNames))},
			{"cluster_config.0.worker_config.0.min_cpu_platform", "Intel Broadwell", cluster.Config.WorkerConfig.MinCpuPlatform},

			{"cluster_config.0.preemptible_worker_config.0.num_instances", "1", strconv.Itoa(int(cluster.Config.SecondaryWorkerConfig.NumInstances))},
			{"cluster_config.0.preemptible_worker_config.0.disk_config.0.boot_disk_size_gb", "35", strconv.Itoa(int(cluster.Config.SecondaryWorkerConfig.DiskConfig.BootDiskSizeGb))},
			{"cluster_config.0.preemptible_worker_config.0.disk_config.0.num_local_ssds", "1", strconv.Itoa(int(cluster.Config.SecondaryWorkerConfig.DiskConfig.NumLocalSsds))},
			{"cluster_config.0.preemptible_worker_config.0.disk_config.0.boot_disk_type", "pd-ssd", cluster.Config.SecondaryWorkerConfig.DiskConfig.BootDiskType},
			{"cluster_config.0.preemptible_worker_config.0.disk_config.0.local_ssd_interface", "nvme", cluster.Config.SecondaryWorkerConfig.DiskConfig.LocalSsdInterface},
			{"cluster_config.0.preemptible_worker_config.0.instance_names.#", "1", strconv.Itoa(len(cluster.Config.SecondaryWorkerConfig.InstanceNames))},
		}

		for _, attrs := range clusterTests {
			tfVal := rs.Primary.Attributes[attrs.tfAttr]
			if tfVal != attrs.expectedVal {
				return fmt.Errorf("%s: Terraform Attribute value '%s' is not as expected '%s' ", attrs.tfAttr, tfVal, attrs.expectedVal)
			}
			if attrs.actualGCPVal != tfVal {
				return fmt.Errorf("%s: Terraform Attribute value '%s' is not aligned with that in GCP '%s' ", attrs.tfAttr, tfVal, attrs.actualGCPVal)
			}
		}

		return nil
	}
}

func testAccCheckDataprocClusterExists(t *testing.T, n string, cluster *dataproc.Cluster) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Terraform resource Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set for Dataproc cluster")
		}

		config := acctest.GoogleProviderConfig(t)
		project, err := acctest.GetTestProject(rs.Primary, config)
		if err != nil {
			return err
		}

		parts := strings.Split(rs.Primary.ID, "/")
		clusterId := parts[len(parts)-1]
		found, err := config.NewDataprocClient(config.UserAgent).Projects.Regions.Clusters.Get(
			project, rs.Primary.Attributes["region"], clusterId).Do()
		if err != nil {
			return err
		}

		if found.ClusterName != clusterId {
			return fmt.Errorf("Dataproc cluster %s not found, found %s instead", clusterId, cluster.ClusterName)
		}

		*cluster = *found

		return nil
	}
}

func testAccCheckDataproc_missingZoneGlobalRegion1(rnd string) string {
	return fmt.Sprintf(`
resource "google_dataproc_cluster" "basic" {
  name   = "tf-test-dproc-%s"
  region = "global"
}
`, rnd)
}

func testAccCheckDataproc_missingZoneGlobalRegion2(rnd string) string {
	return fmt.Sprintf(`
resource "google_dataproc_cluster" "basic" {
  name   = "tf-test-dproc-%s"
  region = "global"

  cluster_config {
    gce_cluster_config {
      network = "default"
    }
  }
}
`, rnd)
}

func testAccDataprocCluster_basic(rnd string) string {
	return fmt.Sprintf(`
resource "google_dataproc_cluster" "basic" {
  name   = "tf-test-dproc-%s"
  region = "us-central1"
}
`, rnd)
}

func testAccDataprocVirtualCluster_basic(projectID, rnd, networkName, subnetworkName string) string {
	return fmt.Sprintf(`
data "google_project" "project" {
  project_id = "%s"
}

resource "google_container_cluster" "primary" {
  name     = "tf-test-gke-%s"
  location = "us-central1-a"
  network    = "%s"
  subnetwork    = "%s"

  initial_node_count = 1

  workload_identity_config {
    workload_pool = "${data.google_project.project.project_id}.svc.id.goog"
  }
  deletion_protection = false
}

resource "google_project_iam_binding" "workloadidentity" {
  project = "%s"
  role    = "roles/iam.workloadIdentityUser"

  members = [
    "serviceAccount:${data.google_project.project.project_id}.svc.id.goog[tf-test-dproc-%s/agent]",
    "serviceAccount:${data.google_project.project.project_id}.svc.id.goog[tf-test-dproc-%s/spark-driver]",
    "serviceAccount:${data.google_project.project.project_id}.svc.id.goog[tf-test-dproc-%s/spark-executor]",
  ]
}

resource "google_dataproc_cluster" "virtual_cluster" {
	depends_on = [
	  google_project_iam_binding.workloadidentity
	]
  
	name   	= "tf-test-dproc-%s"
	region  = "us-central1"
  
	virtual_cluster_config {
	  kubernetes_cluster_config {
		kubernetes_namespace = "tf-test-dproc-%s"
		kubernetes_software_config {
		  component_version = {
			"SPARK": "3.1-dataproc-7",
		  }
		}
		gke_cluster_config {
		  gke_cluster_target = google_container_cluster.primary.id
		  node_pool_target {
			node_pool = "tf-test-gke-np-%s"
			roles = [
			  "DEFAULT"
			]
		  }
		} 
	  }
	}
  }
`, projectID, rnd, networkName, subnetworkName, projectID, rnd, rnd, rnd, rnd, rnd, rnd)
}

func testAccCheckDataprocGkeClusterNodePoolsHaveRoles(cluster *dataproc.Cluster, roles ...string) func(s *terraform.State) error {
	return func(s *terraform.State) error {

		for _, nodePool := range cluster.VirtualClusterConfig.KubernetesClusterConfig.GkeClusterConfig.NodePoolTarget {
			if reflect.DeepEqual(roles, nodePool.Roles) {
				return nil
			}
		}

		return fmt.Errorf("Cluster NodePools does not contain expected roles : %v", roles)
	}
}

func testAccDataprocCluster_withAccelerators(rnd, acceleratorType, zone, subnetworkName string) string {
	return fmt.Sprintf(`
resource "google_dataproc_cluster" "accelerated_cluster" {
  name   = "tf-test-dproc-%s"
  region = "us-central1"

  cluster_config {
    gce_cluster_config {
      subnetwork = "%s"
      zone = "%s"
    }

    master_config {
      accelerators {
        accelerator_type  = "%s"
        accelerator_count = "1"
      }
    }

    worker_config {
      accelerators {
        accelerator_type  = "%s"
        accelerator_count = "1"
      }
    }
  }
}
`, rnd, subnetworkName, zone, acceleratorType, acceleratorType)
}

func testAccDataprocCluster_withInternalIpOnlyTrueAndShieldedConfig(rnd string) string {
	return fmt.Sprintf(`
variable "subnetwork_cidr" {
  default = "10.0.0.0/16"
}

resource "google_compute_network" "dataproc_network" {
  name                    = "tf-test-dproc-net-%s"
  auto_create_subnetworks = false
}

#
# Create a subnet with Private IP Access enabled to test
# deploying a Dataproc cluster with Internal IP Only enabled.
#
resource "google_compute_subnetwork" "dataproc_subnetwork" {
  name                     = "tf-test-dproc-subnet-%s"
  ip_cidr_range            = var.subnetwork_cidr
  network                  = google_compute_network.dataproc_network.self_link
  region                   = "us-central1"
  private_ip_google_access = true
}

#
# The default network within GCP already comes pre configured with
# certain firewall rules open to allow internal communication. As we
# are creating a new one here for this test, we need to additionally
# open up similar rules to allow the nodes to talk to each other
# internally as part of their configuration or this will just hang.
#
resource "google_compute_firewall" "dataproc_network_firewall" {
  name        = "tf-test-dproc-firewall-%s"
  description = "Firewall rules for dataproc Terraform acceptance testing"
  network     = google_compute_network.dataproc_network.name

  allow {
    protocol = "icmp"
  }

  allow {
    protocol = "tcp"
    ports    = ["0-65535"]
  }

  allow {
    protocol = "udp"
    ports    = ["0-65535"]
  }

  source_ranges = [var.subnetwork_cidr]
}

resource "google_dataproc_cluster" "basic" {
  name       = "tf-test-dproc-%s"
  region     = "us-central1"
  depends_on = [google_compute_firewall.dataproc_network_firewall]

  cluster_config {
    gce_cluster_config {
      subnetwork       = google_compute_subnetwork.dataproc_subnetwork.name
      internal_ip_only = true
      shielded_instance_config{
        enable_integrity_monitoring = true
        enable_secure_boot          = true
        enable_vtpm                 = true
      }
    }
  }
}
`, rnd, rnd, rnd, rnd)
}

func testAccDataprocCluster_withMetadataAndTags(rnd, subnetworkName string) string {
	return fmt.Sprintf(`
resource "google_dataproc_cluster" "basic" {
  name   = "tf-test-dproc-%s"
  region = "us-central1"

  cluster_config {
    gce_cluster_config {
      subnetwork = "%s"
      metadata = {
        foo = "bar"
        baz = "qux"
      }
      tags = ["my-tag", "your-tag", "our-tag", "their-tag"]
    }
  }
}
`, rnd, subnetworkName)
}

func testAccDataprocCluster_withMinNumInstances(rnd, subnetworkName string) string {
	return fmt.Sprintf(`
resource "google_dataproc_cluster" "with_min_num_instances" {
  name   = "tf-test-dproc-%s"
  region = "us-central1"
 
  cluster_config {
    gce_cluster_config {
      subnetwork = "%s"
    }
    master_config{
      num_instances=1
    }
    worker_config{
      num_instances = 3
      min_num_instances = 2
    }
  }
}
`, rnd, subnetworkName)
}

func testAccDataprocCluster_withReservationAffinity(rnd, subnetworkName string) string {
	return fmt.Sprintf(`

resource "google_compute_reservation" "reservation" {
  name = "tf-test-dproc-reservation-%s"
  zone = "us-central1-f"

  specific_reservation {
    count = 10
    instance_properties {
      machine_type = "n1-standard-2"
    }
  }
  specific_reservation_required = true
}

resource "google_dataproc_cluster" "basic" {
  name   = "tf-test-dproc-%s"
  region = "us-central1"

  cluster_config {
    master_config {
      machine_type  = "n1-standard-2"
    }

    worker_config {
      machine_type  = "n1-standard-2"
    }

    gce_cluster_config {
      subnetwork = "%s"
      zone = "us-central1-f"
      reservation_affinity {
        consume_reservation_type = "SPECIFIC_RESERVATION"
        key = "compute.googleapis.com/reservation-name"
        values = [google_compute_reservation.reservation.name]
      }
    }
  }
}
`, rnd, rnd, subnetworkName)
}

func testAccDataprocCluster_withDataprocMetricConfig(rnd, subnetworkName string) string {
	return fmt.Sprintf(`
resource "google_dataproc_cluster" "basic" {
  name   = "tf-test-dproc-%s"
  region = "us-central1"

  cluster_config {
    gce_cluster_config {
      subnetwork = "%s"
    }
    dataproc_metric_config {
      metrics {
        metric_source = "HDFS"
        metric_overrides = ["yarn:ResourceManager:QueueMetrics:AppsCompleted"]
      }

      metrics {
        metric_source = "SPARK"
        metric_overrides = ["spark:driver:DAGScheduler:job.allJobs"]
      }
    }
  }
}
`, rnd, subnetworkName)
}

func testAccDataprocCluster_withNodeGroupAffinity(rnd, subnetworkName string) string {
	return fmt.Sprintf(`

resource "google_compute_node_template" "nodetmpl" {
  name   = "test-nodetmpl-%s"
  region = "us-central1"

  node_affinity_labels = {
    tfacc = "test"
  }

  node_type = "n1-node-96-624"

  cpu_overcommit_type = "ENABLED"
}

resource "google_compute_node_group" "nodes" {
  name = "test-nodegroup-%s"
  zone = "us-central1-f"

  initial_size	= 3
  node_template = google_compute_node_template.nodetmpl.self_link
}

resource "google_dataproc_cluster" "basic" {
  name   = "tf-test-dproc-%s"
  region = "us-central1"

  cluster_config {
    gce_cluster_config {
      subnetwork = "%s"
      zone = "us-central1-f"
      node_group_affinity {
        node_group_uri = google_compute_node_group.nodes.name
      }
    }
  }
}
`, rnd, rnd, rnd, subnetworkName)
}

func testAccDataprocCluster_singleNodeCluster(rnd, subnetworkName string) string {
	return fmt.Sprintf(`
resource "google_dataproc_cluster" "single_node_cluster" {
  name   = "tf-test-dproc-%s"
  region = "us-central1"

  cluster_config {
    gce_cluster_config {
      subnetwork = "%s"
    }

    # Keep the costs down with smallest config we can get away with
    software_config {
      override_properties = {
        "dataproc:dataproc.allow.zero.workers" = "true"
      }
    }
  }
}
`, rnd, subnetworkName)
}

func testAccDataprocCluster_withConfigOverrides(rnd, subnetworkName string) string {
	return fmt.Sprintf(`
resource "google_dataproc_cluster" "with_config_overrides" {
  name     = "tf-test-dproc-%s"
  region   = "us-central1"

  cluster_config {
    gce_cluster_config {
      subnetwork = "%s"
    }
    master_config {
      num_instances = 3
      machine_type  = "n1-standard-2"  // can't be e2 because of min_cpu_platform
      disk_config {
        boot_disk_type    = "pd-ssd"
        boot_disk_size_gb = 35
        local_ssd_interface = "nvme"
      }
      min_cpu_platform = "Intel Skylake"
    }

    worker_config {
      num_instances = 3
      machine_type  = "n1-standard-2"  // can't be e2 because of min_cpu_platform
      disk_config {
        boot_disk_type    = "pd-standard"
        boot_disk_size_gb = 35
        num_local_ssds    = 1
        local_ssd_interface = "scsi"
      }

      min_cpu_platform = "Intel Broadwell"
    }

    preemptible_worker_config {
      num_instances = 1
      disk_config {
        boot_disk_type    = "pd-ssd"
        boot_disk_size_gb = 35
        num_local_ssds    = 1
        local_ssd_interface = "nvme"
      }
    }
  }
}
`, rnd, subnetworkName)
}

func testAccDataprocCluster_withInitAction(rnd, bucket, objName, subnetworkName string) string {
	return fmt.Sprintf(`
resource "google_storage_bucket" "init_bucket" {
  name          = "%s"
  location      = "US"
  force_destroy = "true"
}

resource "google_storage_bucket_object" "init_script" {
  name    = "dproc-cluster-test-%s-init-script.sh"
  bucket  = google_storage_bucket.init_bucket.name
  content = <<EOL
#!/bin/bash
echo "init action success" >> /tmp/%s
gsutil cp /tmp/%s ${google_storage_bucket.init_bucket.url}
EOL

}

resource "google_dataproc_cluster" "with_init_action" {
  name   = "tf-test-dproc-%s"
  region = "us-central1"

  cluster_config {
    gce_cluster_config {
      subnetwork = "%s"
    }

    # Keep the costs down with smallest config we can get away with
    software_config {
      override_properties = {
        "dataproc:dataproc.allow.zero.workers" = "true"
      }
    }

    master_config {
      machine_type = "e2-medium"
      disk_config {
        boot_disk_size_gb = 35
      }
    }

    initialization_action {
      script      = "${google_storage_bucket.init_bucket.url}/${google_storage_bucket_object.init_script.name}"
      timeout_sec = 500
    }
    initialization_action {
      script = "${google_storage_bucket.init_bucket.url}/${google_storage_bucket_object.init_script.name}"
    }
  }
}
`, bucket, rnd, objName, objName, rnd, subnetworkName)
}

func testAccDataprocCluster_updatable(rnd string, w, p int) string {
	return fmt.Sprintf(`
resource "google_dataproc_cluster" "updatable" {
  name   = "tf-test-dproc-%s"
  region = "us-central1"
  graceful_decommission_timeout = "0.2s"

  cluster_config {
    master_config {
      num_instances = "1"
      machine_type  = "e2-medium"
      disk_config {
        boot_disk_size_gb = 35
      }
    }

    worker_config {
      num_instances = "%d"
      machine_type  = "e2-medium"
      disk_config {
        boot_disk_size_gb = 35
      }
    }

    preemptible_worker_config {
      num_instances = "%d"
      disk_config {
        boot_disk_size_gb = 35
      }
    }
  }
}
`, rnd, w, p)
}

func testAccDataprocCluster_nonPreemptibleSecondary(rnd, subnetworkName string) string {
	return fmt.Sprintf(`
resource "google_dataproc_cluster" "non_preemptible_secondary" {
  name   = "tf-test-dproc-%s"
  region = "us-central1"

  cluster_config {
    gce_cluster_config {
      subnetwork = "%s"
    }

    master_config {
      num_instances = "1"
      machine_type  = "e2-medium"
      disk_config {
        boot_disk_size_gb = 35
      }
    }
  
    worker_config {
      num_instances = "2"
      machine_type  = "e2-medium"
      disk_config {
        boot_disk_size_gb = 35
      }
    }
  
    preemptible_worker_config {
      num_instances = "1"
      preemptibility = "NON_PREEMPTIBLE"
      disk_config {
        boot_disk_size_gb = 35
      }
    }
  }
}
	`, rnd, subnetworkName)
}

func testAccDataprocCluster_spotSecondary(rnd, subnetworkName string) string {
	return fmt.Sprintf(`
resource "google_dataproc_cluster" "spot_secondary" {
  name   = "tf-test-dproc-%s"
  region = "us-central1"

  cluster_config {
    gce_cluster_config {
      subnetwork = "%s"
    }

    master_config {
      num_instances = "1"
      machine_type  = "e2-medium"
      disk_config {
        boot_disk_size_gb = 35
      }
    }

    worker_config {
      num_instances = "2"
      machine_type  = "e2-medium"
      disk_config {
        boot_disk_size_gb = 35
      }
    }

    preemptible_worker_config {
      num_instances = "1"
      preemptibility = "SPOT"
      disk_config {
        boot_disk_size_gb = 35
      }
    }
  }
}
	`, rnd, subnetworkName)
}

func testAccDataprocCluster_spotWithInstanceFlexibilityPolicy(rnd string) string {
	return fmt.Sprintf(`
resource "google_dataproc_cluster" "spot_with_instance_flexibility_policy" {
  name   = "tf-test-dproc-%s"
  region = "us-central1"

  cluster_config {
    master_config {
      num_instances = "1"
      machine_type  = "e2-medium"
      disk_config {
        boot_disk_size_gb = 35
      }
    }

    worker_config {
      num_instances = "2"
      machine_type  = "e2-medium"
      disk_config {
        boot_disk_size_gb = 35
      }
    }

    preemptible_worker_config {
      num_instances = "3"
      preemptibility = "SPOT"
      disk_config {
        boot_disk_size_gb = 35
      }
	  instance_flexibility_policy {
        instance_selection_list {
          machine_types = ["n2d-standard-2"]
          rank          = 3
        }
      }
    }
  }
}
	`, rnd)
}

func testAccDataprocCluster_withAuxiliaryNodeGroups(rnd string) string {
	return fmt.Sprintf(`
resource "google_dataproc_cluster" "with_auxiliary_node_groups" {
  name   = "tf-test-dproc-%s"
  region = "us-central1"

  cluster_config {
    master_config {
      num_instances = "1"
      machine_type  = "e2-medium"
      disk_config {
        boot_disk_size_gb = 35
      }
    }

    worker_config {
      num_instances = "2"
      machine_type  = "e2-medium"
      disk_config {
        boot_disk_size_gb = 35
      }
    }

    auxiliary_node_groups{
      node_group_id="node-group-id"
      node_group {
        roles = ["DRIVER"]
        node_group_config{
          num_instances=2
          machine_type="n1-standard-2"
          min_cpu_platform = "Intel Haswell"
          disk_config {
            boot_disk_size_gb = 35
            boot_disk_type = "pd-standard"
            num_local_ssds = 1
            local_ssd_interface = "nvme"
          }
          accelerators {
            accelerator_count = 1
            accelerator_type  = "nvidia-tesla-t4"
          }
        }
      }
    }
  }
}
	`, rnd)
}

func testAccDataprocCluster_withStagingBucketOnly(bucketName string) string {
	return fmt.Sprintf(`
resource "google_storage_bucket" "bucket" {
  name          = "%s"
  location      = "US"
  force_destroy = "true"
}
`, bucketName)
}

func testAccDataprocCluster_withTempBucketOnly(bucketName string) string {
	return fmt.Sprintf(`
resource "google_storage_bucket" "bucket" {
  name          = "%s"
  location      = "US"
  force_destroy = "true"
}
`, bucketName)
}

func testAccDataprocCluster_withStagingBucketAndCluster(clusterName, bucketName, subnetworkName string) string {
	return fmt.Sprintf(`
%s

resource "google_dataproc_cluster" "with_bucket" {
  name   = "%s"
  region = "us-central1"

  cluster_config {
    staging_bucket = google_storage_bucket.bucket.name

    gce_cluster_config {
      subnetwork = "%s"
    }

    # Keep the costs down with smallest config we can get away with
    software_config {
      override_properties = {
        "dataproc:dataproc.allow.zero.workers" = "true"
      }
    }

    master_config {
      machine_type = "e2-medium"
      disk_config {
        boot_disk_size_gb = 35
      }
    }
  }
}
`, testAccDataprocCluster_withStagingBucketOnly(bucketName), clusterName, subnetworkName)
}

func testAccDataprocCluster_withTempBucketAndCluster(clusterName, bucketName, subnetworkName string) string {
	return fmt.Sprintf(`
%s

resource "google_dataproc_cluster" "with_bucket" {
  name   = "%s"
  region = "us-central1"

  cluster_config {
    temp_bucket = google_storage_bucket.bucket.name

    gce_cluster_config {
      subnetwork = "%s"
    }

    # Keep the costs down with smallest config we can get away with
    software_config {
      override_properties = {
        "dataproc:dataproc.allow.zero.workers" = "true"
      }
    }

    master_config {
      machine_type = "e2-medium"
      disk_config {
        boot_disk_size_gb = 35
      }
    }
  }
}
`, testAccDataprocCluster_withTempBucketOnly(bucketName), clusterName, subnetworkName)
}

func testAccDataprocCluster_withLabels(rnd, subnetworkName string) string {
	return fmt.Sprintf(`
resource "google_dataproc_cluster" "with_labels" {
  name   = "tf-test-dproc-%s"
  region = "us-central1"
  cluster_config {
    gce_cluster_config {
      subnetwork = "%s"
    }
  }

  labels = {
    key1 = "value1"
  }
}
`, rnd, subnetworkName)
}

func testAccDataprocCluster_withLabelsUpdate(rnd, subnetworkName string) string {
	return fmt.Sprintf(`
resource "google_dataproc_cluster" "with_labels" {
  name   = "tf-test-dproc-%s"
  region = "us-central1"
  cluster_config {
    gce_cluster_config {
      subnetwork = "%s"
    }
  }

  labels = {
    key2 = "value2"
  }
}
`, rnd, subnetworkName)
}

func testAccDataprocCluster_withoutLabels(rnd, subnetworkName string) string {
	return fmt.Sprintf(`
resource "google_dataproc_cluster" "with_labels" {
  name   = "tf-test-dproc-%s"
  region = "us-central1"
  cluster_config {
    gce_cluster_config {
      subnetwork = "%s"
    }
  }
}
`, rnd, subnetworkName)
}

func testAccDataprocCluster_withEndpointConfig(rnd, subnetworkName string) string {
	return fmt.Sprintf(`
resource "google_dataproc_cluster" "with_endpoint_config" {
	name                  = "tf-test-%s"
	region                = "us-central1"

	cluster_config {
    gce_cluster_config {
      subnetwork = "%s"
    }

		endpoint_config {
			enable_http_port_access = true
		}
	}
}
`, rnd, subnetworkName)
}

func testAccDataprocCluster_withImageVersion(rnd, version, subnetworkName string) string {
	return fmt.Sprintf(`
resource "google_dataproc_cluster" "with_image_version" {
  name   = "tf-test-dproc-%s"
  region = "us-central1"

  cluster_config {
    gce_cluster_config {
      subnetwork = "%s"
    }

    software_config {
      image_version = "%s"
    }
  }
}
`, rnd, subnetworkName, version)
}

func testAccDataprocCluster_withOptionalComponents(rnd, subnetworkName string) string {
	return fmt.Sprintf(`
resource "google_dataproc_cluster" "with_opt_components" {
  name   = "tf-test-dproc-%s"
  region = "us-central1"

  cluster_config {
    gce_cluster_config {
      subnetwork = "%s"
    }

    software_config {
      optional_components = ["DOCKER", "ZOOKEEPER"]
    }
  }
}
`, rnd, subnetworkName)
}

func testAccDataprocCluster_withLifecycleConfigIdleDeleteTtl(rnd, tm, subnetworkName string) string {
	return fmt.Sprintf(`
resource "google_dataproc_cluster" "with_lifecycle_config" {
  name   = "tf-test-dproc-%s"
  region = "us-central1"

  cluster_config {
    gce_cluster_config {
      subnetwork = "%s"
    }

    lifecycle_config {
      idle_delete_ttl = "%s"
    }
  }
}
`, rnd, subnetworkName, tm)
}

func testAccDataprocCluster_withLifecycleConfigAutoDeletionTime(rnd, tm, subnetworkName string) string {
	return fmt.Sprintf(`
resource "google_dataproc_cluster" "with_lifecycle_config" {
 name   = "tf-test-dproc-%s"
 region = "us-central1"

 cluster_config {
  gce_cluster_config {
      subnetwork = "%s"
    }

   lifecycle_config {
     auto_delete_time = "%s"
   }
 }
}
`, rnd, subnetworkName, tm)
}

func testAccDataprocCluster_withServiceAcc(sa, rnd, subnetworkName string) string {
	return fmt.Sprintf(`
data "google_project" "project" {}

resource "google_service_account" "service_account" {
  account_id = "%s"
}

resource "google_project_iam_member" "service_account" {
  project = data.google_project.project.project_id
  role   = "roles/dataproc.worker"
  member = "serviceAccount:${google_service_account.service_account.email}"
}

# Wait for IAM propagation
resource "time_sleep" "wait_120_seconds" {
  depends_on = [google_project_iam_member.service_account]

  create_duration = "120s"
}

resource "google_dataproc_cluster" "with_service_account" {
  name   = "dproc-cluster-test-%s"
  region = "us-central1"

  cluster_config {
    # Keep the costs down with smallest config we can get away with
    software_config {
      override_properties = {
        "dataproc:dataproc.allow.zero.workers" = "true"
      }
    }

    master_config {
      machine_type = "e2-medium"
      disk_config {
        boot_disk_size_gb = 35
      }
    }

    gce_cluster_config {
      subnetwork = "%s"
      service_account = google_service_account.service_account.email
      service_account_scopes = [
		#	User supplied scopes
        "https://www.googleapis.com/auth/monitoring",
		#	The following scopes necessary for the cluster to function properly are
		#	always added, even if not explicitly specified:
		#		useraccounts-ro: https://www.googleapis.com/auth/cloud.useraccounts.readonly
		#		storage-rw:      https://www.googleapis.com/auth/devstorage.read_write
		#		logging-write:   https://www.googleapis.com/auth/logging.write
        "useraccounts-ro",
        "storage-rw",
        "logging-write",
      ]
    }
  }

  depends_on = [time_sleep.wait_120_seconds]
}
`, sa, rnd, subnetworkName)
}

func testAccDataprocCluster_withNetworkRefs(rnd, netName string) string {
	return fmt.Sprintf(`
resource "google_compute_network" "dataproc_network" {
  name                    = "%s"
  auto_create_subnetworks = true
}

#
# The default network within GCP already comes pre configured with
# certain firewall rules open to allow internal communication. As we
# are creating a new one here for this test, we need to additionally
# open up similar rules to allow the nodes to talk to each other
# internally as part of their configuration or this will just hang.
#
resource "google_compute_firewall" "dataproc_network_firewall" {
  name          = "tf-test-dproc-%s"
  description   = "Firewall rules for dataproc Terraform acceptance testing"
  network       = google_compute_network.dataproc_network.name
  source_ranges = ["192.168.0.0/16"]

  allow {
    protocol = "icmp"
  }

  allow {
    protocol = "tcp"
    ports    = ["0-65535"]
  }

  allow {
    protocol = "udp"
    ports    = ["0-65535"]
  }
}

resource "google_dataproc_cluster" "with_net_ref_by_name" {
  name       = "tf-test-dproc-net-%s"
  region     = "us-central1"
  depends_on = [google_compute_firewall.dataproc_network_firewall]

  cluster_config {

    # Keep the costs down with smallest config we can get away with
    software_config {
      override_properties = {
        "dataproc:dataproc.allow.zero.workers" = "true"
      }
    }

    master_config {
      machine_type = "e2-medium"
      disk_config {
        boot_disk_size_gb = 35
      }
    }

    gce_cluster_config {
      network = google_compute_network.dataproc_network.name
    }
  }
}

resource "google_dataproc_cluster" "with_net_ref_by_url" {
  name       = "tf-test-dproc-url-%s"
  region     = "us-central1"
  depends_on = [google_compute_firewall.dataproc_network_firewall]

  cluster_config {

    # Keep the costs down with smallest config we can get away with
    software_config {
      override_properties = {
        "dataproc:dataproc.allow.zero.workers" = "true"
      }
    }

    master_config {
      machine_type = "e2-medium"
      disk_config {
        boot_disk_size_gb = 35
      }
    }

    gce_cluster_config {
      network = google_compute_network.dataproc_network.self_link
    }
  }
}
`, netName, rnd, rnd, rnd)
}

func testAccDataprocCluster_KMS(rnd, kmsKey, subnetworkName string) string {
	return fmt.Sprintf(`
resource "google_dataproc_cluster" "kms" {
  name   = "tf-test-dproc-%s"
  region = "us-central1"

  cluster_config {
    gce_cluster_config {
      subnetwork = "%s"
    }

    encryption_config {
      kms_key_name = "%s"
    }
  }
}
`, rnd, subnetworkName, kmsKey)
}

func testAccDataprocCluster_withSecurityConfig_withoutKerberosConfig(rnd, subnetworkName string) string {
	return fmt.Sprintf(`
resource "google_storage_bucket" "bucket" {
  name     = "tf-test-dproc-%s"
  location = "US"
}
resource "google_storage_bucket_object" "password" {
  name = "dataproc-password-%s"
  bucket = google_storage_bucket.bucket.name
  content = "hunter2"
}

resource "google_dataproc_cluster" "nokerb" {
  name   = "tf-test-dproc-%s"
  region = "us-central1"

  cluster_config {
    gce_cluster_config {
      subnetwork = "%s"
    }

    security_config {}
  }
}
`, rnd, rnd, rnd, subnetworkName)
}

func testAccDataprocCluster_withKerberos(rnd, kmsKey, subnetworkName string) string {
	return fmt.Sprintf(`
resource "google_storage_bucket" "bucket" {
  name     = "tf-test-dproc-%s"
  location = "US"
}
resource "google_storage_bucket_object" "password" {
  name = "dataproc-password-%s"
  bucket = google_storage_bucket.bucket.name
  content = "hunter2"
}

resource "google_dataproc_cluster" "kerb" {
  name   = "tf-test-dproc-%s"
  region = "us-central1"

  cluster_config {
    gce_cluster_config {
      subnetwork = "%s"
    }

    security_config {
      kerberos_config {
        root_principal_password_uri = google_storage_bucket_object.password.self_link
        kms_key_uri = "%s"
      }
    }
  }
}
`, rnd, rnd, rnd, subnetworkName, kmsKey)
}

func testAccDataprocCluster_withAutoscalingPolicy(rnd, subnetworkName string) string {
	return fmt.Sprintf(`
resource "google_dataproc_cluster" "basic" {
  name     = "tf-test-dataproc-policy-%s"
  region   = "us-central1"

  cluster_config {
    gce_cluster_config {
      subnetwork = "%s"
    }

    autoscaling_config {
      policy_uri = google_dataproc_autoscaling_policy.asp.id
    }
  }
}

resource "google_dataproc_autoscaling_policy" "asp" {
  policy_id = "tf-test-dataproc-policy-%s"
  location  = "us-central1"

  worker_config {
    max_instances = 3
  }

  basic_algorithm {
    yarn_config {
      graceful_decommission_timeout = "30s"
      scale_up_factor   = 0.5
      scale_down_factor = 0.5
    }
  }
}
`, rnd, subnetworkName, rnd)
}

func testAccDataprocCluster_removeAutoscalingPolicy(rnd, subnetworkName string) string {
	return fmt.Sprintf(`
resource "google_dataproc_cluster" "basic" {
  name     = "tf-test-dataproc-policy-%s"
  region   = "us-central1"

  cluster_config {
    gce_cluster_config {
      subnetwork = "%s"
    }

    autoscaling_config {
      policy_uri = ""
    }
  }
}

resource "google_dataproc_autoscaling_policy" "asp" {
  policy_id = "tf-test-dataproc-policy-%s"
  location  = "us-central1"

  worker_config {
    max_instances = 3
  }

  basic_algorithm {
    yarn_config {
      graceful_decommission_timeout = "30s"
      scale_up_factor   = 0.5
      scale_down_factor = 0.5
    }
  }
}
`, rnd, subnetworkName, rnd)
}

func testAccDataprocCluster_withMetastoreConfig(clusterName, serviceId string) string {
	return fmt.Sprintf(`
resource "google_dataproc_cluster" "with_metastore_config" {
  name                  = "%s"
  region                = "us-central1"

  cluster_config {
    metastore_config {
      dataproc_metastore_service = google_dataproc_metastore_service.ms.name
    }
  }
}

resource "google_dataproc_metastore_service" "ms" {
  service_id = "%s"
  location   = "us-central1"
  port       = 9080
  tier       = "DEVELOPER"

  maintenance_window {
    hour_of_day = 2
    day_of_week = "SUNDAY"
  }

  hive_metastore_config {
    version = "3.1.2"
  }
}
`, clusterName, serviceId)
}

func testAccDataprocCluster_withMetastoreConfig_update(clusterName, serviceId string) string {
	return fmt.Sprintf(`
resource "google_dataproc_cluster" "with_metastore_config" {
  name                  = "%s"
  region                = "us-central1"

  cluster_config {
    metastore_config {
      dataproc_metastore_service = google_dataproc_metastore_service.ms.name
    }
  }
}

resource "google_dataproc_metastore_service" "ms" {
  service_id = "%s"
  location   = "us-central1"
  port       = 9080
  tier       = "DEVELOPER"

  maintenance_window {
    hour_of_day = 2
    day_of_week = "SUNDAY"
  }

  hive_metastore_config {
    version = "3.1.2"
  }
}
`, clusterName, serviceId)
}
