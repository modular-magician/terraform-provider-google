// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package sweeper_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	_ "github.com/hashicorp/terraform-provider-google/google/services/accessapproval"
	_ "github.com/hashicorp/terraform-provider-google/google/services/accesscontextmanager"
	_ "github.com/hashicorp/terraform-provider-google/google/services/activedirectory"
	_ "github.com/hashicorp/terraform-provider-google/google/services/alloydb"
	_ "github.com/hashicorp/terraform-provider-google/google/services/apigee"
	_ "github.com/hashicorp/terraform-provider-google/google/services/appengine"
	_ "github.com/hashicorp/terraform-provider-google/google/services/artifactregistry"
	_ "github.com/hashicorp/terraform-provider-google/google/services/beyondcorp"
	_ "github.com/hashicorp/terraform-provider-google/google/services/biglake"
	_ "github.com/hashicorp/terraform-provider-google/google/services/bigquery"
	_ "github.com/hashicorp/terraform-provider-google/google/services/bigqueryanalyticshub"
	_ "github.com/hashicorp/terraform-provider-google/google/services/bigqueryconnection"
	_ "github.com/hashicorp/terraform-provider-google/google/services/bigquerydatapolicy"
	_ "github.com/hashicorp/terraform-provider-google/google/services/bigquerydatatransfer"
	_ "github.com/hashicorp/terraform-provider-google/google/services/bigqueryreservation"
	_ "github.com/hashicorp/terraform-provider-google/google/services/bigtable"
	_ "github.com/hashicorp/terraform-provider-google/google/services/billing"
	_ "github.com/hashicorp/terraform-provider-google/google/services/binaryauthorization"
	_ "github.com/hashicorp/terraform-provider-google/google/services/certificatemanager"
	_ "github.com/hashicorp/terraform-provider-google/google/services/cloudasset"
	_ "github.com/hashicorp/terraform-provider-google/google/services/cloudbuild"
	_ "github.com/hashicorp/terraform-provider-google/google/services/cloudbuildv2"
	_ "github.com/hashicorp/terraform-provider-google/google/services/cloudfunctions"
	_ "github.com/hashicorp/terraform-provider-google/google/services/cloudfunctions2"
	_ "github.com/hashicorp/terraform-provider-google/google/services/cloudidentity"
	_ "github.com/hashicorp/terraform-provider-google/google/services/cloudids"
	_ "github.com/hashicorp/terraform-provider-google/google/services/cloudrun"
	_ "github.com/hashicorp/terraform-provider-google/google/services/cloudrunv2"
	_ "github.com/hashicorp/terraform-provider-google/google/services/cloudscheduler"
	_ "github.com/hashicorp/terraform-provider-google/google/services/cloudtasks"
	_ "github.com/hashicorp/terraform-provider-google/google/services/compute"
	_ "github.com/hashicorp/terraform-provider-google/google/services/containeranalysis"
	_ "github.com/hashicorp/terraform-provider-google/google/services/containerattached"
	_ "github.com/hashicorp/terraform-provider-google/google/services/corebilling"
	_ "github.com/hashicorp/terraform-provider-google/google/services/databasemigrationservice"
	_ "github.com/hashicorp/terraform-provider-google/google/services/datacatalog"
	_ "github.com/hashicorp/terraform-provider-google/google/services/datafusion"
	_ "github.com/hashicorp/terraform-provider-google/google/services/datalossprevention"
	_ "github.com/hashicorp/terraform-provider-google/google/services/datapipeline"
	_ "github.com/hashicorp/terraform-provider-google/google/services/dataplex"
	_ "github.com/hashicorp/terraform-provider-google/google/services/dataproc"
	_ "github.com/hashicorp/terraform-provider-google/google/services/dataprocmetastore"
	_ "github.com/hashicorp/terraform-provider-google/google/services/datastore"
	_ "github.com/hashicorp/terraform-provider-google/google/services/datastream"
	_ "github.com/hashicorp/terraform-provider-google/google/services/deploymentmanager"
	_ "github.com/hashicorp/terraform-provider-google/google/services/dialogflow"
	_ "github.com/hashicorp/terraform-provider-google/google/services/dialogflowcx"
	_ "github.com/hashicorp/terraform-provider-google/google/services/dns"
	_ "github.com/hashicorp/terraform-provider-google/google/services/documentai"
	_ "github.com/hashicorp/terraform-provider-google/google/services/documentaiwarehouse"
	_ "github.com/hashicorp/terraform-provider-google/google/services/edgecontainer"
	_ "github.com/hashicorp/terraform-provider-google/google/services/edgenetwork"
	_ "github.com/hashicorp/terraform-provider-google/google/services/essentialcontacts"
	_ "github.com/hashicorp/terraform-provider-google/google/services/filestore"
	_ "github.com/hashicorp/terraform-provider-google/google/services/firestore"
	_ "github.com/hashicorp/terraform-provider-google/google/services/gkebackup"
	_ "github.com/hashicorp/terraform-provider-google/google/services/gkehub"
	_ "github.com/hashicorp/terraform-provider-google/google/services/gkehub2"
	_ "github.com/hashicorp/terraform-provider-google/google/services/healthcare"
	_ "github.com/hashicorp/terraform-provider-google/google/services/iam2"
	_ "github.com/hashicorp/terraform-provider-google/google/services/iambeta"
	_ "github.com/hashicorp/terraform-provider-google/google/services/iamworkforcepool"
	_ "github.com/hashicorp/terraform-provider-google/google/services/iap"
	_ "github.com/hashicorp/terraform-provider-google/google/services/identityplatform"
	_ "github.com/hashicorp/terraform-provider-google/google/services/kms"
	_ "github.com/hashicorp/terraform-provider-google/google/services/logging"
	_ "github.com/hashicorp/terraform-provider-google/google/services/looker"
	_ "github.com/hashicorp/terraform-provider-google/google/services/memcache"
	_ "github.com/hashicorp/terraform-provider-google/google/services/mlengine"
	_ "github.com/hashicorp/terraform-provider-google/google/services/monitoring"
	_ "github.com/hashicorp/terraform-provider-google/google/services/netapp"
	_ "github.com/hashicorp/terraform-provider-google/google/services/networkconnectivity"
	_ "github.com/hashicorp/terraform-provider-google/google/services/networkmanagement"
	_ "github.com/hashicorp/terraform-provider-google/google/services/networksecurity"
	_ "github.com/hashicorp/terraform-provider-google/google/services/networkservices"
	_ "github.com/hashicorp/terraform-provider-google/google/services/notebooks"
	_ "github.com/hashicorp/terraform-provider-google/google/services/osconfig"
	_ "github.com/hashicorp/terraform-provider-google/google/services/oslogin"
	_ "github.com/hashicorp/terraform-provider-google/google/services/privateca"
	_ "github.com/hashicorp/terraform-provider-google/google/services/publicca"
	_ "github.com/hashicorp/terraform-provider-google/google/services/pubsub"
	_ "github.com/hashicorp/terraform-provider-google/google/services/pubsublite"
	_ "github.com/hashicorp/terraform-provider-google/google/services/redis"
	_ "github.com/hashicorp/terraform-provider-google/google/services/resourcemanager"
	_ "github.com/hashicorp/terraform-provider-google/google/services/secretmanager"
	_ "github.com/hashicorp/terraform-provider-google/google/services/securitycenter"
	_ "github.com/hashicorp/terraform-provider-google/google/services/servicemanagement"
	_ "github.com/hashicorp/terraform-provider-google/google/services/serviceusage"
	_ "github.com/hashicorp/terraform-provider-google/google/services/sourcerepo"
	_ "github.com/hashicorp/terraform-provider-google/google/services/spanner"
	_ "github.com/hashicorp/terraform-provider-google/google/services/sql"
	_ "github.com/hashicorp/terraform-provider-google/google/services/storage"
	_ "github.com/hashicorp/terraform-provider-google/google/services/storageinsights"
	_ "github.com/hashicorp/terraform-provider-google/google/services/storagetransfer"
	_ "github.com/hashicorp/terraform-provider-google/google/services/tags"
	_ "github.com/hashicorp/terraform-provider-google/google/services/tpu"
	_ "github.com/hashicorp/terraform-provider-google/google/services/vertexai"
	_ "github.com/hashicorp/terraform-provider-google/google/services/vpcaccess"
	_ "github.com/hashicorp/terraform-provider-google/google/services/workflows"

	// Manually add the services for DCL resource and handwritten resource sweepers if they are not in the above list
	_ "github.com/hashicorp/terraform-provider-google/google/services/apikeys"
	_ "github.com/hashicorp/terraform-provider-google/google/services/clouddeploy"
	_ "github.com/hashicorp/terraform-provider-google/google/services/composer"
	_ "github.com/hashicorp/terraform-provider-google/google/services/container"
	_ "github.com/hashicorp/terraform-provider-google/google/services/containeraws"
	_ "github.com/hashicorp/terraform-provider-google/google/services/containerazure"
	_ "github.com/hashicorp/terraform-provider-google/google/services/dataflow"
	_ "github.com/hashicorp/terraform-provider-google/google/services/eventarc"
	_ "github.com/hashicorp/terraform-provider-google/google/services/firebase"
	_ "github.com/hashicorp/terraform-provider-google/google/services/firebaserules"
	_ "github.com/hashicorp/terraform-provider-google/google/services/networkconnectivity"
	_ "github.com/hashicorp/terraform-provider-google/google/services/recaptchaenterprise"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}
