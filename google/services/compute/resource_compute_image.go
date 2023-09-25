// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

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

package compute

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
	"github.com/hashicorp/terraform-provider-google/google/verify"
)

func ResourceComputeImage() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeImageCreate,
		Read:   resourceComputeImageRead,
		Update: resourceComputeImageUpdate,
		Delete: resourceComputeImageDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeImageImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.SetLabelsDiff,
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `Name of the resource; provided by the client when the resource is
created. The name must be 1-63 characters long, and comply with
RFC1035. Specifically, the name must be 1-63 characters long and
match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means
the first character must be a lowercase letter, and all following
characters must be a dash, lowercase letter, or digit, except the
last character, which cannot be a dash.`,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `An optional description of this resource. Provide this property when
you create the resource.`,
			},
			"disk_size_gb": {
				Type:        schema.TypeInt,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: `Size of the image when restored onto a persistent disk (in GB).`,
			},
			"family": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `The name of the image family to which this image belongs. You can
create disks by specifying an image family instead of a specific
image name. The image family always returns its latest image that is
not deprecated. The name of the image family must comply with
RFC1035.`,
			},
			"guest_os_features": {
				Type:     schema.TypeSet,
				Computed: true,
				Optional: true,
				ForceNew: true,
				Description: `A list of features to enable on the guest operating system.
Applicable only for bootable images.`,
				Elem: computeImageGuestOsFeaturesSchema(),
				// Default schema.HashSchema is used.
			},
			"image_encryption_key": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Description: `Encrypts the image using a customer-supplied encryption key.

After you encrypt an image with a customer-supplied key, you must
provide the same key if you use the image later (e.g. to create a
disk from the image)`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"kms_key_self_link": {
							Type:             schema.TypeString,
							Optional:         true,
							ForceNew:         true,
							DiffSuppressFunc: tpgresource.CompareSelfLinkRelativePaths,
							Description: `The self link of the encryption key that is stored in Google Cloud
KMS.`,
						},
						"kms_key_service_account": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Description: `The service account being used for the encryption request for the
given KMS key. If absent, the Compute Engine default service
account is used.`,
						},
					},
				},
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `Labels to apply to this Image.

**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
Please refer to the field 'effective_labels' for all of the labels present on the resource.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"licenses": {
				Type:        schema.TypeList,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: `Any applicable license URI.`,
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				},
			},
			"raw_disk": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: `The parameters of the raw disk image.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"source": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
							Description: `The full Google Cloud Storage URL where disk storage is stored
You must provide either this property or the sourceDisk property
but not both.`,
						},
						"container_type": {
							Type:         schema.TypeString,
							Optional:     true,
							ForceNew:     true,
							ValidateFunc: verify.ValidateEnum([]string{"TAR", ""}),
							Description: `The format used to encode and transmit the block device, which
should be TAR. This is just a container and transmission format
and not a runtime format. Provided by the client when the disk
image is created. Default value: "TAR" Possible values: ["TAR"]`,
							Default: "TAR",
						},
						"sha1": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Description: `An optional SHA1 checksum of the disk image before unpackaging.
This is provided by the client when the disk image is created.`,
						},
					},
				},
			},
			"source_disk": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description: `The source disk to create this image based on.
You must provide either this property or the
rawDisk.source property but not both to create an image.`,
			},
			"source_image": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description: `URL of the source image used to create this image. In order to create an image, you must provide the full or partial
URL of one of the following:

* The selfLink URL
* This property
* The rawDisk.source URL
* The sourceDisk URL`,
			},
			"source_snapshot": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description: `URL of the source snapshot used to create this image.

In order to create an image, you must provide the full or partial URL of one of the following:

* The selfLink URL
* This property
* The sourceImage URL
* The rawDisk.source URL
* The sourceDisk URL`,
			},
			"storage_locations": {
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				ForceNew: true,
				Description: `Cloud Storage bucket storage location of the image
(regional or multi-regional).
Reference link: https://cloud.google.com/compute/docs/reference/rest/v1/images`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"archive_size_bytes": {
				Type:     schema.TypeInt,
				Computed: true,
				Description: `Size of the image tar.gz archive stored in Google Cloud Storage (in
bytes).`,
			},
			"creation_timestamp": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Creation timestamp in RFC3339 text format.`,
			},
			"effective_labels": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: `All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"label_fingerprint": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The fingerprint used for optimistic locking of this resource. Used
internally during updates.`,
			},
			"terraform_labels": {
				Type:     schema.TypeMap,
				Computed: true,
				Description: `The combination of labels configured directly on the resource
 and default labels configured on the provider.`,
				Elem: &schema.Schema{Type: schema.TypeString},
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
		UseJSONNumber: true,
	}
}

func computeImageGuestOsFeaturesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"type": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateEnum([]string{"MULTI_IP_SUBNET", "SECURE_BOOT", "SEV_CAPABLE", "UEFI_COMPATIBLE", "VIRTIO_SCSI_MULTIQUEUE", "WINDOWS", "GVNIC", "SEV_LIVE_MIGRATABLE", "SEV_SNP_CAPABLE", "SUSPEND_RESUME_COMPATIBLE", "TDX_CAPABLE"}),
				Description:  `The type of supported feature. Read [Enabling guest operating system features](https://cloud.google.com/compute/docs/images/create-delete-deprecate-private-images#guest-os-features) to see a list of available options. Possible values: ["MULTI_IP_SUBNET", "SECURE_BOOT", "SEV_CAPABLE", "UEFI_COMPATIBLE", "VIRTIO_SCSI_MULTIQUEUE", "WINDOWS", "GVNIC", "SEV_LIVE_MIGRATABLE", "SEV_SNP_CAPABLE", "SUSPEND_RESUME_COMPATIBLE", "TDX_CAPABLE"]`,
			},
		},
	}
}

func resourceComputeImageCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	descriptionProp, err := expandComputeImageDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	storageLocationsProp, err := expandComputeImageStorageLocations(d.Get("storage_locations"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("storage_locations"); !tpgresource.IsEmptyValue(reflect.ValueOf(storageLocationsProp)) && (ok || !reflect.DeepEqual(v, storageLocationsProp)) {
		obj["storageLocations"] = storageLocationsProp
	}
	diskSizeGbProp, err := expandComputeImageDiskSizeGb(d.Get("disk_size_gb"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("disk_size_gb"); !tpgresource.IsEmptyValue(reflect.ValueOf(diskSizeGbProp)) && (ok || !reflect.DeepEqual(v, diskSizeGbProp)) {
		obj["diskSizeGb"] = diskSizeGbProp
	}
	familyProp, err := expandComputeImageFamily(d.Get("family"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("family"); !tpgresource.IsEmptyValue(reflect.ValueOf(familyProp)) && (ok || !reflect.DeepEqual(v, familyProp)) {
		obj["family"] = familyProp
	}
	guestOsFeaturesProp, err := expandComputeImageGuestOsFeatures(d.Get("guest_os_features"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("guest_os_features"); !tpgresource.IsEmptyValue(reflect.ValueOf(guestOsFeaturesProp)) && (ok || !reflect.DeepEqual(v, guestOsFeaturesProp)) {
		obj["guestOsFeatures"] = guestOsFeaturesProp
	}
	imageEncryptionKeyProp, err := expandComputeImageImageEncryptionKey(d.Get("image_encryption_key"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("image_encryption_key"); !tpgresource.IsEmptyValue(reflect.ValueOf(imageEncryptionKeyProp)) && (ok || !reflect.DeepEqual(v, imageEncryptionKeyProp)) {
		obj["imageEncryptionKey"] = imageEncryptionKeyProp
	}
	labelFingerprintProp, err := expandComputeImageLabelFingerprint(d.Get("label_fingerprint"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("label_fingerprint"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelFingerprintProp)) && (ok || !reflect.DeepEqual(v, labelFingerprintProp)) {
		obj["labelFingerprint"] = labelFingerprintProp
	}
	licensesProp, err := expandComputeImageLicenses(d.Get("licenses"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("licenses"); !tpgresource.IsEmptyValue(reflect.ValueOf(licensesProp)) && (ok || !reflect.DeepEqual(v, licensesProp)) {
		obj["licenses"] = licensesProp
	}
	nameProp, err := expandComputeImageName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	rawDiskProp, err := expandComputeImageRawDisk(d.Get("raw_disk"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("raw_disk"); !tpgresource.IsEmptyValue(reflect.ValueOf(rawDiskProp)) && (ok || !reflect.DeepEqual(v, rawDiskProp)) {
		obj["rawDisk"] = rawDiskProp
	}
	sourceDiskProp, err := expandComputeImageSourceDisk(d.Get("source_disk"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("source_disk"); !tpgresource.IsEmptyValue(reflect.ValueOf(sourceDiskProp)) && (ok || !reflect.DeepEqual(v, sourceDiskProp)) {
		obj["sourceDisk"] = sourceDiskProp
	}
	sourceImageProp, err := expandComputeImageSourceImage(d.Get("source_image"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("source_image"); !tpgresource.IsEmptyValue(reflect.ValueOf(sourceImageProp)) && (ok || !reflect.DeepEqual(v, sourceImageProp)) {
		obj["sourceImage"] = sourceImageProp
	}
	sourceSnapshotProp, err := expandComputeImageSourceSnapshot(d.Get("source_snapshot"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("source_snapshot"); !tpgresource.IsEmptyValue(reflect.ValueOf(sourceSnapshotProp)) && (ok || !reflect.DeepEqual(v, sourceSnapshotProp)) {
		obj["sourceSnapshot"] = sourceSnapshotProp
	}
	labelsProp, err := expandComputeImageEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/images")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Image: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Image: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
	})
	if err != nil {
		return fmt.Errorf("Error creating Image: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/global/images/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = ComputeOperationWaitTime(
		config, res, project, "Creating Image", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create Image: %s", err)
	}

	log.Printf("[DEBUG] Finished creating Image %q: %#v", d.Id(), res)

	return resourceComputeImageRead(d, meta)
}

func resourceComputeImageRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/images/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Image: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("ComputeImage %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Image: %s", err)
	}

	if err := d.Set("archive_size_bytes", flattenComputeImageArchiveSizeBytes(res["archiveSizeBytes"], d, config)); err != nil {
		return fmt.Errorf("Error reading Image: %s", err)
	}
	if err := d.Set("creation_timestamp", flattenComputeImageCreationTimestamp(res["creationTimestamp"], d, config)); err != nil {
		return fmt.Errorf("Error reading Image: %s", err)
	}
	if err := d.Set("description", flattenComputeImageDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading Image: %s", err)
	}
	if err := d.Set("storage_locations", flattenComputeImageStorageLocations(res["storageLocations"], d, config)); err != nil {
		return fmt.Errorf("Error reading Image: %s", err)
	}
	if err := d.Set("disk_size_gb", flattenComputeImageDiskSizeGb(res["diskSizeGb"], d, config)); err != nil {
		return fmt.Errorf("Error reading Image: %s", err)
	}
	if err := d.Set("family", flattenComputeImageFamily(res["family"], d, config)); err != nil {
		return fmt.Errorf("Error reading Image: %s", err)
	}
	if err := d.Set("guest_os_features", flattenComputeImageGuestOsFeatures(res["guestOsFeatures"], d, config)); err != nil {
		return fmt.Errorf("Error reading Image: %s", err)
	}
	if err := d.Set("image_encryption_key", flattenComputeImageImageEncryptionKey(res["imageEncryptionKey"], d, config)); err != nil {
		return fmt.Errorf("Error reading Image: %s", err)
	}
	if err := d.Set("labels", flattenComputeImageLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Image: %s", err)
	}
	if err := d.Set("label_fingerprint", flattenComputeImageLabelFingerprint(res["labelFingerprint"], d, config)); err != nil {
		return fmt.Errorf("Error reading Image: %s", err)
	}
	if err := d.Set("licenses", flattenComputeImageLicenses(res["licenses"], d, config)); err != nil {
		return fmt.Errorf("Error reading Image: %s", err)
	}
	if err := d.Set("name", flattenComputeImageName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Image: %s", err)
	}
	if err := d.Set("source_disk", flattenComputeImageSourceDisk(res["sourceDisk"], d, config)); err != nil {
		return fmt.Errorf("Error reading Image: %s", err)
	}
	if err := d.Set("source_image", flattenComputeImageSourceImage(res["sourceImage"], d, config)); err != nil {
		return fmt.Errorf("Error reading Image: %s", err)
	}
	if err := d.Set("source_snapshot", flattenComputeImageSourceSnapshot(res["sourceSnapshot"], d, config)); err != nil {
		return fmt.Errorf("Error reading Image: %s", err)
	}
	if err := d.Set("terraform_labels", flattenComputeImageTerraformLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Image: %s", err)
	}
	if err := d.Set("effective_labels", flattenComputeImageEffectiveLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Image: %s", err)
	}
	if err := d.Set("self_link", tpgresource.ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading Image: %s", err)
	}

	return nil
}

func resourceComputeImageUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Image: %s", err)
	}
	billingProject = project

	d.Partial(true)

	if d.HasChange("label_fingerprint") || d.HasChange("effective_labels") {
		obj := make(map[string]interface{})

		labelFingerprintProp, err := expandComputeImageLabelFingerprint(d.Get("label_fingerprint"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("label_fingerprint"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelFingerprintProp)) {
			obj["labelFingerprint"] = labelFingerprintProp
		}
		labelsProp, err := expandComputeImageEffectiveLabels(d.Get("effective_labels"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
			obj["labels"] = labelsProp
		}

		url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/images/{{name}}/setLabels")
		if err != nil {
			return err
		}

		// err == nil indicates that the billing_project value was found
		if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
			billingProject = bp
		}

		res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "POST",
			Project:   billingProject,
			RawURL:    url,
			UserAgent: userAgent,
			Body:      obj,
			Timeout:   d.Timeout(schema.TimeoutUpdate),
		})
		if err != nil {
			return fmt.Errorf("Error updating Image %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating Image %q: %#v", d.Id(), res)
		}

		err = ComputeOperationWaitTime(
			config, res, project, "Updating Image", userAgent,
			d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return err
		}
	}

	d.Partial(false)

	return resourceComputeImageRead(d, meta)
}

func resourceComputeImageDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Image: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/images/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Image %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "Image")
	}

	err = ComputeOperationWaitTime(
		config, res, project, "Deleting Image", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Image %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeImageImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"projects/(?P<project>[^/]+)/global/images/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/global/images/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeImageArchiveSizeBytes(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
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

func flattenComputeImageCreationTimestamp(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeImageDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeImageStorageLocations(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeImageDiskSizeGb(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
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

func flattenComputeImageFamily(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeImageGuestOsFeatures(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := schema.NewSet(schema.HashResource(computeImageGuestOsFeaturesSchema()), []interface{}{})
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed.Add(map[string]interface{}{
			"type": flattenComputeImageGuestOsFeaturesType(original["type"], d, config),
		})
	}
	return transformed
}
func flattenComputeImageGuestOsFeaturesType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeImageImageEncryptionKey(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["kms_key_self_link"] =
		flattenComputeImageImageEncryptionKeyKmsKeySelfLink(original["kmsKeyName"], d, config)
	transformed["kms_key_service_account"] =
		flattenComputeImageImageEncryptionKeyKmsKeyServiceAccount(original["kmsKeyServiceAccount"], d, config)
	return []interface{}{transformed}
}
func flattenComputeImageImageEncryptionKeyKmsKeySelfLink(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	vStr := v.(string)
	return strings.Split(vStr, "/cryptoKeyVersions/")[0]
}

func flattenComputeImageImageEncryptionKeyKmsKeyServiceAccount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeImageLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("labels"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenComputeImageLabelFingerprint(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeImageLicenses(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.ConvertAndMapStringArr(v.([]interface{}), tpgresource.ConvertSelfLinkToV1)
}

func flattenComputeImageName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeImageSourceDisk(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.ConvertSelfLinkToV1(v.(string))
}

func flattenComputeImageSourceImage(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.ConvertSelfLinkToV1(v.(string))
}

func flattenComputeImageSourceSnapshot(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.ConvertSelfLinkToV1(v.(string))
}

func flattenComputeImageTerraformLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("terraform_labels"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenComputeImageEffectiveLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandComputeImageDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeImageStorageLocations(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeImageDiskSizeGb(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeImageFamily(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeImageGuestOsFeatures(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedType, err := expandComputeImageGuestOsFeaturesType(original["type"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["type"] = transformedType
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeImageGuestOsFeaturesType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeImageImageEncryptionKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedKmsKeySelfLink, err := expandComputeImageImageEncryptionKeyKmsKeySelfLink(original["kms_key_self_link"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsKeySelfLink); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["kmsKeyName"] = transformedKmsKeySelfLink
	}

	transformedKmsKeyServiceAccount, err := expandComputeImageImageEncryptionKeyKmsKeyServiceAccount(original["kms_key_service_account"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsKeyServiceAccount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["kmsKeyServiceAccount"] = transformedKmsKeyServiceAccount
	}

	return transformed, nil
}

func expandComputeImageImageEncryptionKeyKmsKeySelfLink(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeImageImageEncryptionKeyKmsKeyServiceAccount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeImageLabelFingerprint(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeImageLicenses(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			return nil, fmt.Errorf("Invalid value for licenses: nil")
		}
		f, err := tpgresource.ParseGlobalFieldValue("licenses", raw.(string), "project", d, config, true)
		if err != nil {
			return nil, fmt.Errorf("Invalid value for licenses: %s", err)
		}
		req = append(req, f.RelativeLink())
	}
	return req, nil
}

func expandComputeImageName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeImageRawDisk(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedContainerType, err := expandComputeImageRawDiskContainerType(original["container_type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedContainerType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["containerType"] = transformedContainerType
	}

	transformedSha1, err := expandComputeImageRawDiskSha1(original["sha1"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSha1); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["sha1Checksum"] = transformedSha1
	}

	transformedSource, err := expandComputeImageRawDiskSource(original["source"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSource); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["source"] = transformedSource
	}

	return transformed, nil
}

func expandComputeImageRawDiskContainerType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeImageRawDiskSha1(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeImageRawDiskSource(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeImageSourceDisk(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseZonalFieldValue("disks", v.(string), "project", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for source_disk: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeImageSourceImage(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseGlobalFieldValue("images", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for source_image: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeImageSourceSnapshot(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseGlobalFieldValue("snapshots", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for source_snapshot: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeImageEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
