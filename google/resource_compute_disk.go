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
	"log"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform/helper/customdiff"
	"github.com/hashicorp/terraform/helper/schema"
	compute "google.golang.org/api/compute/v1"
	"google.golang.org/api/googleapi"
)

const (
	computeDiskUserRegexString = "^(?:https://www.googleapis.com/compute/v1/projects/)?(" + ProjectRegex + ")/zones/([-_a-zA-Z0-9]*)/instances/([-_a-zA-Z0-9]*)$"
)

var (
	computeDiskUserRegex = regexp.MustCompile(computeDiskUserRegexString)
)

// Is the new disk size smaller than the old one?
func isDiskShrinkage(old, new, _ interface{}) bool {
	// It's okay to remove size entirely.
	if old == nil || new == nil {
		return false
	}
	return new.(int) < old.(int)
}

// We cannot suppress the diff for the case when family name is not part of the image name since we can't
// make a network call in a DiffSuppressFunc.
func diskImageDiffSuppress(_, old, new string, _ *schema.ResourceData) bool {
	// 'old' is read from the API.
	// It always has the format 'https://www.googleapis.com/compute/v1/projects/(%s)/global/images/(%s)'
	matches := resolveImageLink.FindStringSubmatch(old)
	if matches == nil {
		// Image read from the API doesn't have the expected format. In practice, it should never happen
		return false
	}
	oldProject := matches[1]
	oldName := matches[2]

	// Partial or full self link family
	if resolveImageProjectFamily.MatchString(new) {
		// Value matches pattern "projects/{project}/global/images/family/{family-name}$"
		matches := resolveImageProjectFamily.FindStringSubmatch(new)
		newProject := matches[1]
		newFamilyName := matches[2]

		return diskImageProjectNameEquals(oldProject, newProject) && diskImageFamilyEquals(oldName, newFamilyName)
	}

	// Partial or full self link image
	if resolveImageProjectImage.MatchString(new) {
		// Value matches pattern "projects/{project}/global/images/{image-name}$"
		matches := resolveImageProjectImage.FindStringSubmatch(new)
		newProject := matches[1]
		newImageName := matches[2]

		return diskImageProjectNameEquals(oldProject, newProject) && diskImageEquals(oldName, newImageName)
	}

	// Partial link without project family
	if resolveImageGlobalFamily.MatchString(new) {
		// Value is "global/images/family/{family-name}"
		matches := resolveImageGlobalFamily.FindStringSubmatch(new)
		familyName := matches[1]

		return diskImageFamilyEquals(oldName, familyName)
	}

	// Partial link without project image
	if resolveImageGlobalImage.MatchString(new) {
		// Value is "global/images/{image-name}"
		matches := resolveImageGlobalImage.FindStringSubmatch(new)
		imageName := matches[1]

		return diskImageEquals(oldName, imageName)
	}

	// Family shorthand
	if resolveImageFamilyFamily.MatchString(new) {
		// Value is "family/{family-name}"
		matches := resolveImageFamilyFamily.FindStringSubmatch(new)
		familyName := matches[1]

		return diskImageFamilyEquals(oldName, familyName)
	}

	// Shorthand for image or family
	if resolveImageProjectImageShorthand.MatchString(new) {
		// Value is "{project}/{image-name}" or "{project}/{family-name}"
		matches := resolveImageProjectImageShorthand.FindStringSubmatch(new)
		newProject := matches[1]
		newName := matches[2]

		return diskImageProjectNameEquals(oldProject, newProject) &&
			(diskImageEquals(oldName, newName) || diskImageFamilyEquals(oldName, newName))
	}

	// Image or family only
	if diskImageEquals(oldName, new) || diskImageFamilyEquals(oldName, new) {
		// Value is "{image-name}" or "{family-name}"
		return true
	}

	return false
}

func diskImageProjectNameEquals(project1, project2 string) bool {
	// Convert short project name to full name
	// For instance, centos => centos-cloud
	fullProjectName, ok := imageMap[project2]
	if ok {
		project2 = fullProjectName
	}

	return project1 == project2
}

func diskImageEquals(oldImageName, newImageName string) bool {
	return oldImageName == newImageName
}

func diskImageFamilyEquals(imageName, familyName string) bool {
	// Handles the case when the image name includes the family name
	// e.g. image name: debian-9-drawfork-v20180109, family name: debian-9
	if strings.Contains(imageName, familyName) {
		return true
	}

	if suppressCanonicalFamilyDiff(imageName, familyName) {
		return true
	}

	if suppressWindowsSqlFamilyDiff(imageName, familyName) {
		return true
	}

	if suppressWindowsFamilyDiff(imageName, familyName) {
		return true
	}

	return false
}

// e.g. image: ubuntu-1404-trusty-v20180122, family: ubuntu-1404-lts
func suppressCanonicalFamilyDiff(imageName, familyName string) bool {
	parts := canonicalUbuntuLtsImage.FindStringSubmatch(imageName)
	if len(parts) == 2 {
		f := fmt.Sprintf("ubuntu-%s-lts", parts[1])
		if f == familyName {
			return true
		}
	}

	return false
}

// e.g. image: sql-2017-standard-windows-2016-dc-v20180109, family: sql-std-2017-win-2016
// e.g. image: sql-2017-express-windows-2012-r2-dc-v20180109, family: sql-exp-2017-win-2012-r2
func suppressWindowsSqlFamilyDiff(imageName, familyName string) bool {
	parts := windowsSqlImage.FindStringSubmatch(imageName)
	if len(parts) == 5 {
		edition := parts[2] // enterprise, standard or web.
		sqlVersion := parts[1]
		windowsVersion := parts[3]

		// Translate edition
		switch edition {
		case "enterprise":
			edition = "ent"
		case "standard":
			edition = "std"
		case "express":
			edition = "exp"
		}

		var f string
		if revision := parts[4]; revision != "" {
			// With revision
			f = fmt.Sprintf("sql-%s-%s-win-%s-r%s", edition, sqlVersion, windowsVersion, revision)
		} else {
			// No revision
			f = fmt.Sprintf("sql-%s-%s-win-%s", edition, sqlVersion, windowsVersion)
		}

		if f == familyName {
			return true
		}
	}

	return false
}

// e.g. image: windows-server-1709-dc-core-v20180109, family: windows-1709-core
// e.g. image: windows-server-1709-dc-core-for-containers-v20180109, family: "windows-1709-core-for-containers
func suppressWindowsFamilyDiff(imageName, familyName string) bool {
	updatedFamilyString := strings.Replace(familyName, "windows-", "windows-server-", 1)
	updatedFamilyString = strings.Replace(updatedFamilyString, "-core", "-dc-core", 1)

	if strings.Contains(imageName, updatedFamilyString) {
		return true
	}

	return false
}

func diskEncryptionKeyDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	if strings.HasSuffix(k, "#") {
		if old == "1" && new == "0" {
			// If we have a disk_encryption_key_raw, we can trust that the diff will be handled there
			// and we don't need to worry about it here.
			return d.Get("disk_encryption_key_raw").(string) != ""
		} else if new == "1" && old == "0" {
			// This will be handled by diffing the 'raw_key' attribute.
			return true
		}
	} else if strings.HasSuffix(k, "raw_key") {
		disk_key := d.Get("disk_encryption_key_raw").(string)
		return disk_key == old && old != "" && new == ""
	} else if k == "disk_encryption_key_raw" {
		disk_key := d.Get("disk_encryption_key.0.raw_key").(string)
		return disk_key == old && old != "" && new == ""
	}
	return false
}

func resourceComputeDisk() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeDiskCreate,
		Read:   resourceComputeDiskRead,
		Update: resourceComputeDiskUpdate,
		Delete: resourceComputeDiskDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeDiskImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(300 * time.Second),
			Update: schema.DefaultTimeout(240 * time.Second),
			Delete: schema.DefaultTimeout(240 * time.Second),
		},
		CustomizeDiff: customdiff.All(
			customdiff.ForceNewIfChange("size", isDiskShrinkage)),

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"size": {
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"image": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: diskImageDiffSuppress,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Default:  "pd-standard",
			},
			"zone": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
			},
			"disk_encryption_key": {
				Type:             schema.TypeList,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: diskEncryptionKeyDiffSuppress,
				MaxItems:         1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"raw_key": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"sha256": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"source_image_encryption_key": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"raw_key": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"sha256": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"snapshot": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: linkDiffSuppress,
			},
			"source_snapshot_encryption_key": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"raw_key": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"sha256": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"creation_timestamp": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"label_fingerprint": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_attach_timestamp": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_detach_timestamp": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"source_image_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"source_snapshot_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"users": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"disk_encryption_key_raw": &schema.Schema{
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				Sensitive:        true,
				DiffSuppressFunc: diskEncryptionKeyDiffSuppress,
				Deprecated:       "Use disk_encryption_key.raw_key instead.",
			},

			"disk_encryption_key_sha256": &schema.Schema{
				Type:       schema.TypeString,
				Computed:   true,
				Deprecated: "Use disk_encryption_key.sha256 instead.",
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

func resourceComputeDiskCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	descriptionProp, err := expandComputeDiskDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); ok || !reflect.DeepEqual(v, descriptionProp) {
		obj["description"] = descriptionProp
	}
	labelsProp, err := expandComputeDiskLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); ok || !reflect.DeepEqual(v, labelsProp) {
		obj["labels"] = labelsProp
	}
	nameProp, err := expandComputeDiskName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); ok || !reflect.DeepEqual(v, nameProp) {
		obj["name"] = nameProp
	}
	sizeGbProp, err := expandComputeDiskSize(d.Get("size"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("size"); ok || !reflect.DeepEqual(v, sizeGbProp) {
		obj["sizeGb"] = sizeGbProp
	}
	sourceImageProp, err := expandComputeDiskImage(d.Get("image"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("image"); ok || !reflect.DeepEqual(v, sourceImageProp) {
		obj["sourceImage"] = sourceImageProp
	}
	typeProp, err := expandComputeDiskType(d.Get("type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("type"); ok || !reflect.DeepEqual(v, typeProp) {
		obj["type"] = typeProp
	}
	zoneProp, err := expandComputeDiskZone(d.Get("zone"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("zone"); ok || !reflect.DeepEqual(v, zoneProp) {
		obj["zone"] = zoneProp
	}
	diskEncryptionKeyProp, err := expandComputeDiskDiskEncryptionKey(d.Get("disk_encryption_key"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("disk_encryption_key"); ok || !reflect.DeepEqual(v, diskEncryptionKeyProp) {
		obj["diskEncryptionKey"] = diskEncryptionKeyProp
	}
	sourceImageEncryptionKeyProp, err := expandComputeDiskSourceImageEncryptionKey(d.Get("source_image_encryption_key"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("source_image_encryption_key"); ok || !reflect.DeepEqual(v, sourceImageEncryptionKeyProp) {
		obj["sourceImageEncryptionKey"] = sourceImageEncryptionKeyProp
	}
	sourceSnapshotProp, err := expandComputeDiskSnapshot(d.Get("snapshot"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("snapshot"); ok || !reflect.DeepEqual(v, sourceSnapshotProp) {
		obj["sourceSnapshot"] = sourceSnapshotProp
	}
	sourceSnapshotEncryptionKeyProp, err := expandComputeDiskSourceSnapshotEncryptionKey(d.Get("source_snapshot_encryption_key"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("source_snapshot_encryption_key"); ok || !reflect.DeepEqual(v, sourceSnapshotEncryptionKeyProp) {
		obj["sourceSnapshotEncryptionKey"] = sourceSnapshotEncryptionKeyProp
	}

	obj, err = resourceComputeDiskEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/zones/{{zone}}/disks")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Disk: %#v", obj)
	res, err := Post(config, url, obj)
	if err != nil {
		return fmt.Errorf("Error creating Disk: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{name}}")
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
		config.clientCompute, op, project, "Creating Disk",
		int(d.Timeout(schema.TimeoutCreate).Minutes()))

	if waitErr != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create Disk: %s", waitErr)
	}

	log.Printf("[DEBUG] Finished creating Disk %q: %#v", d.Id(), res)

	return resourceComputeDiskRead(d, meta)
}

func resourceComputeDiskRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/zones/{{zone}}/disks/{{name}}")
	if err != nil {
		return err
	}

	res, err := Get(config, url)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeDisk %q", d.Id()))
	}

	res, err = resourceComputeDiskDecoder(d, meta, res)
	if err != nil {
		return err
	}

	if err := d.Set("label_fingerprint", flattenComputeDiskLabelFingerprint(res["labelFingerprint"])); err != nil {
		return fmt.Errorf("Error reading Disk: %s", err)
	}
	if err := d.Set("creation_timestamp", flattenComputeDiskCreationTimestamp(res["creationTimestamp"])); err != nil {
		return fmt.Errorf("Error reading Disk: %s", err)
	}
	if err := d.Set("description", flattenComputeDiskDescription(res["description"])); err != nil {
		return fmt.Errorf("Error reading Disk: %s", err)
	}
	if err := d.Set("last_attach_timestamp", flattenComputeDiskLastAttachTimestamp(res["lastAttachTimestamp"])); err != nil {
		return fmt.Errorf("Error reading Disk: %s", err)
	}
	if err := d.Set("last_detach_timestamp", flattenComputeDiskLastDetachTimestamp(res["lastDetachTimestamp"])); err != nil {
		return fmt.Errorf("Error reading Disk: %s", err)
	}
	if err := d.Set("labels", flattenComputeDiskLabels(res["labels"])); err != nil {
		return fmt.Errorf("Error reading Disk: %s", err)
	}
	if err := d.Set("name", flattenComputeDiskName(res["name"])); err != nil {
		return fmt.Errorf("Error reading Disk: %s", err)
	}
	if err := d.Set("size", flattenComputeDiskSize(res["sizeGb"])); err != nil {
		return fmt.Errorf("Error reading Disk: %s", err)
	}
	if err := d.Set("image", flattenComputeDiskImage(res["sourceImage"])); err != nil {
		return fmt.Errorf("Error reading Disk: %s", err)
	}
	if err := d.Set("type", flattenComputeDiskType(res["type"])); err != nil {
		return fmt.Errorf("Error reading Disk: %s", err)
	}
	if err := d.Set("users", flattenComputeDiskUsers(res["users"])); err != nil {
		return fmt.Errorf("Error reading Disk: %s", err)
	}
	if err := d.Set("zone", flattenComputeDiskZone(res["zone"])); err != nil {
		return fmt.Errorf("Error reading Disk: %s", err)
	}
	if err := d.Set("disk_encryption_key", flattenComputeDiskDiskEncryptionKey(res["diskEncryptionKey"])); err != nil {
		return fmt.Errorf("Error reading Disk: %s", err)
	}
	if err := d.Set("source_image_encryption_key", flattenComputeDiskSourceImageEncryptionKey(res["sourceImageEncryptionKey"])); err != nil {
		return fmt.Errorf("Error reading Disk: %s", err)
	}
	if err := d.Set("source_image_id", flattenComputeDiskSourceImageId(res["sourceImageId"])); err != nil {
		return fmt.Errorf("Error reading Disk: %s", err)
	}
	if err := d.Set("snapshot", flattenComputeDiskSnapshot(res["sourceSnapshot"])); err != nil {
		return fmt.Errorf("Error reading Disk: %s", err)
	}
	if err := d.Set("source_snapshot_encryption_key", flattenComputeDiskSourceSnapshotEncryptionKey(res["sourceSnapshotEncryptionKey"])); err != nil {
		return fmt.Errorf("Error reading Disk: %s", err)
	}
	if err := d.Set("source_snapshot_id", flattenComputeDiskSourceSnapshotId(res["sourceSnapshotId"])); err != nil {
		return fmt.Errorf("Error reading Disk: %s", err)
	}
	if err := d.Set("self_link", ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading Disk: %s", err)
	}
	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Disk: %s", err)
	}

	return nil
}

func resourceComputeDiskUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	var url string
	var res map[string]interface{}
	op := &compute.Operation{}

	d.Partial(true)

	if d.HasChange("label_fingerprint") || d.HasChange("labels") {
		obj := make(map[string]interface{})
		labelFingerprintProp := d.Get("label_fingerprint")
		obj["labelFingerprint"] = labelFingerprintProp
		labelsProp, err := expandComputeDiskLabels(d.Get("labels"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("labels"); ok || !reflect.DeepEqual(v, labelsProp) || d.HasChange("labels") {
			obj["labels"] = labelsProp
		}

		url, err = replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/zones/{{zone}}/disks/{{name}}/setLabels")
		if err != nil {
			return err
		}
		res, err = sendRequest(config, "POST", url, obj)
		if err != nil {
			return fmt.Errorf("Error updating Disk %q: %s", d.Id(), err)
		}

		err = Convert(res, op)
		if err != nil {
			return err
		}

		err = computeOperationWaitTime(
			config.clientCompute, op, project, "Updating Disk",
			int(d.Timeout(schema.TimeoutUpdate).Minutes()))

		if err != nil {
			return err
		}

		d.SetPartial("label_fingerprint")
		d.SetPartial("labels")
	}
	if d.HasChange("size") {
		obj := make(map[string]interface{})
		sizeGbProp, err := expandComputeDiskSize(d.Get("size"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("size"); ok || !reflect.DeepEqual(v, sizeGbProp) || d.HasChange("size") {
			obj["sizeGb"] = sizeGbProp
		}

		url, err = replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/zones/{{zone}}/disks/{{name}}/resize")
		if err != nil {
			return err
		}
		res, err = sendRequest(config, "POST", url, obj)
		if err != nil {
			return fmt.Errorf("Error updating Disk %q: %s", d.Id(), err)
		}

		err = Convert(res, op)
		if err != nil {
			return err
		}

		err = computeOperationWaitTime(
			config.clientCompute, op, project, "Updating Disk",
			int(d.Timeout(schema.TimeoutUpdate).Minutes()))

		if err != nil {
			return err
		}

		d.SetPartial("size")
	}

	d.Partial(false)

	return resourceComputeDiskRead(d, meta)
}

func resourceComputeDiskDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/zones/{{zone}}/disks/{{name}}")
	if err != nil {
		return err
	}

	// if disks are attached, they must be detached before the disk can be deleted
	if instances, ok := d.Get("users").([]interface{}); ok {
		type detachArgs struct{ project, zone, instance, deviceName string }
		var detachCalls []detachArgs
		self := d.Get("self_link").(string)
		for _, instance := range instances {
			if !computeDiskUserRegex.MatchString(instance.(string)) {
				return fmt.Errorf("Unknown user %q of disk %q", instance, self)
			}
			matches := computeDiskUserRegex.FindStringSubmatch(instance.(string))
			instanceProject := matches[1]
			instanceZone := matches[2]
			instanceName := matches[3]
			i, err := config.clientCompute.Instances.Get(instanceProject, instanceZone, instanceName).Do()
			if err != nil {
				if gerr, ok := err.(*googleapi.Error); ok && gerr.Code == 404 {
					log.Printf("[WARN] instance %q not found, not bothering to detach disks", instance.(string))
					continue
				}
				return fmt.Errorf("Error retrieving instance %s: %s", instance.(string), err.Error())
			}
			for _, disk := range i.Disks {
				if disk.Source == self {
					detachCalls = append(detachCalls, detachArgs{
						project:    project,
						zone:       GetResourceNameFromSelfLink(i.Zone),
						instance:   i.Name,
						deviceName: disk.DeviceName,
					})
				}
			}
		}
		for _, call := range detachCalls {
			op, err := config.clientCompute.Instances.DetachDisk(call.project, call.zone, call.instance, call.deviceName).Do()
			if err != nil {
				return fmt.Errorf("Error detaching disk %s from instance %s/%s/%s: %s", call.deviceName, call.project,
					call.zone, call.instance, err.Error())
			}
			err = computeOperationWait(config.clientCompute, op, call.project,
				fmt.Sprintf("Detaching disk from %s/%s/%s", call.project, call.zone, call.instance))
			if err != nil {
				if opErr, ok := err.(ComputeOperationError); ok && len(opErr.Errors) == 1 && opErr.Errors[0].Code == "RESOURCE_NOT_FOUND" {
					log.Printf("[WARN] instance %q was deleted while awaiting detach", call.instance)
					continue
				}
				return err
			}
		}
	}
	log.Printf("[DEBUG] Deleting Disk %q", d.Id())
	res, err := Delete(config, url)
	if err != nil {
		return handleNotFoundError(err, d, "Disk")
	}

	op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	err = computeOperationWaitTime(
		config.clientCompute, op, project, "Deleting Disk",
		int(d.Timeout(schema.TimeoutDelete).Minutes()))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Disk %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeDiskImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	parseImportId([]string{"projects/(?P<project>[^/]+)/zones/(?P<zone>[^/]+)/disks/(?P<name>[^/]+)", "(?P<project>[^/]+)/(?P<zone>[^/]+)/(?P<name>[^/]+)", "(?P<name>[^/]+)"}, d, config)

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)
	// In the end, it's possible that someone has tried to import
	// a disk using only the region.  To find out what zone the
	// disk is in, we need to check every zone in the region, to
	// see if we can find a disk with the same name.  This will
	// find the first disk in the specified region with a matching
	// name.  There might be multiple matching disks - we're not
	// considering that an error case here.  We don't check for it.
	if zone, err := getZone(d, config); err != nil || zone == "" {
		project, err := getProject(d, config)
		if err != nil {
			return nil, err
		}
		region, err := getRegion(d, config)
		if err != nil {
			return nil, err
		}

		getDisk := func(zone string) (interface{}, error) {
			return config.clientCompute.Disks.Get(project, zone, d.Id()).Do()
		}
		resource, err := getZonalResourceFromRegion(getDisk, region, config.clientCompute, project)
		if err != nil {
			return nil, err
		}
		d.Set("zone", resource.(*compute.Disk).Zone)
	}

	return []*schema.ResourceData{d}, nil
}

func flattenComputeDiskLabelFingerprint(v interface{}) interface{} {
	return v
}

func flattenComputeDiskCreationTimestamp(v interface{}) interface{} {
	return v
}

func flattenComputeDiskDescription(v interface{}) interface{} {
	return v
}

func flattenComputeDiskLastAttachTimestamp(v interface{}) interface{} {
	return v
}

func flattenComputeDiskLastDetachTimestamp(v interface{}) interface{} {
	return v
}

func flattenComputeDiskLabels(v interface{}) interface{} {
	return v
}

func flattenComputeDiskName(v interface{}) interface{} {
	return v
}

func flattenComputeDiskSize(v interface{}) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenComputeDiskImage(v interface{}) interface{} {
	return v
}

func flattenComputeDiskType(v interface{}) interface{} {
	return NameFromSelfLinkStateFunc(v)
}

func flattenComputeDiskUsers(v interface{}) interface{} {
	return v
}

func flattenComputeDiskZone(v interface{}) interface{} {
	return NameFromSelfLinkStateFunc(v)
}

func flattenComputeDiskDiskEncryptionKey(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	transformed := make(map[string]interface{})
	transformed["raw_key"] =
		flattenComputeDiskDiskEncryptionKeyRawKey(original["rawKey"])
	transformed["sha256"] =
		flattenComputeDiskDiskEncryptionKeySha256(original["sha256"])
	return []interface{}{transformed}
}
func flattenComputeDiskDiskEncryptionKeyRawKey(v interface{}) interface{} {
	return v
}

func flattenComputeDiskDiskEncryptionKeySha256(v interface{}) interface{} {
	return v
}

func flattenComputeDiskSourceImageEncryptionKey(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	transformed := make(map[string]interface{})
	transformed["raw_key"] =
		flattenComputeDiskSourceImageEncryptionKeyRawKey(original["rawKey"])
	transformed["sha256"] =
		flattenComputeDiskSourceImageEncryptionKeySha256(original["sha256"])
	return []interface{}{transformed}
}
func flattenComputeDiskSourceImageEncryptionKeyRawKey(v interface{}) interface{} {
	return v
}

func flattenComputeDiskSourceImageEncryptionKeySha256(v interface{}) interface{} {
	return v
}

func flattenComputeDiskSourceImageId(v interface{}) interface{} {
	return v
}

func flattenComputeDiskSnapshot(v interface{}) interface{} {
	return v
}

func flattenComputeDiskSourceSnapshotEncryptionKey(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	transformed := make(map[string]interface{})
	transformed["raw_key"] =
		flattenComputeDiskSourceSnapshotEncryptionKeyRawKey(original["rawKey"])
	transformed["sha256"] =
		flattenComputeDiskSourceSnapshotEncryptionKeySha256(original["sha256"])
	return []interface{}{transformed}
}
func flattenComputeDiskSourceSnapshotEncryptionKeyRawKey(v interface{}) interface{} {
	return v
}

func flattenComputeDiskSourceSnapshotEncryptionKeySha256(v interface{}) interface{} {
	return v
}

func flattenComputeDiskSourceSnapshotId(v interface{}) interface{} {
	return v
}

func expandComputeDiskDescription(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeDiskLabels(v interface{}, d *schema.ResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandComputeDiskName(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeDiskSize(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeDiskImage(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeDiskType(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeDiskZone(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("zones", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for zone: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeDiskDiskEncryptionKey(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, 1)
	if len(l) == 1 {
		// There is a value
		outMap := make(map[string]interface{})
		outMap["rawKey"] = l[0].(map[string]interface{})["raw_key"]
		req = append(req, outMap)
	} else {
		// Check alternative setting?
		if altV, ok := d.GetOk("disk_encryption_key_raw"); ok && altV != "" {
			outMap := make(map[string]interface{})
			outMap["rawKey"] = altV
			req = append(req, outMap)
		}
	}
	return req, nil
}

func expandComputeDiskSourceImageEncryptionKey(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, 1)
	if len(l) == 1 {
		// There is a value
		outMap := make(map[string]interface{})
		outMap["rawKey"] = l[0].(map[string]interface{})["raw_key"]
		req = append(req, outMap)
	}
	return req, nil
}

func expandComputeDiskSnapshot(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeDiskSourceSnapshotEncryptionKey(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, 1)
	if len(l) == 1 {
		// There is a value
		outMap := make(map[string]interface{})
		outMap["rawKey"] = l[0].(map[string]interface{})["raw_key"]
		req = append(req, outMap)
	}
	return req, nil
}

func resourceComputeDiskEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return nil, err
	}

	// Get the zone
	z, err := getZone(d, config)
	if err != nil {
		return nil, err
	}

	zone, err := config.clientCompute.Zones.Get(project, z).Do()
	if err != nil {
		return nil, err
	}

	if v, ok := d.GetOk("image"); ok {
		log.Printf("[DEBUG] Resolving image name: %s", v.(string))
		imageUrl, err := resolveImage(config, project, v.(string))
		if err != nil {
			return nil, fmt.Errorf(
				"Error resolving image name '%s': %s",
				v.(string), err)
		}

		obj["sourceImage"] = imageUrl
		log.Printf("[DEBUG] Image name resolved to: %s", imageUrl)
	}

	if v, ok := d.GetOk("type"); ok {
		log.Printf("[DEBUG] Loading disk type: %s", v.(string))
		diskType, err := readDiskType(config, zone, project, v.(string))
		if err != nil {
			return nil, fmt.Errorf(
				"Error loading disk type '%s': %s",
				v.(string), err)
		}

		obj["type"] = diskType.SelfLink
	}

	if v, ok := d.GetOk("snapshot"); ok {
		snapshotName := v.(string)
		match, _ := regexp.MatchString("^https://www.googleapis.com/compute", snapshotName)
		if match {
			obj["sourceSnapshot"] = snapshotName
		} else {
			log.Printf("[DEBUG] Loading snapshot: %s", snapshotName)
			snapshotData, err := config.clientCompute.Snapshots.Get(
				project, snapshotName).Do()

			if err != nil {
				return nil, fmt.Errorf(
					"Error loading snapshot '%s': %s",
					snapshotName, err)
			}
			obj["sourceSnapshot"] = snapshotData.SelfLink
		}
	}

	return obj, nil
}

func resourceComputeDiskDecoder(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	if v, ok := res["diskEncryptionKey"]; ok {
		original := v.(map[string]interface{})
		transformed := make(map[string]interface{})
		// The raw key won't be returned, so we need to use the original.
		transformed["rawKey"] = d.Get("disk_encryption_key.0.raw_key")
		transformed["sha256"] = original["sha256"]
		if v, ok := d.GetOk("disk_encryption_key_raw"); ok {
			transformed["rawKey"] = v
		}
		d.Set("disk_encryption_key_sha256", original["sha256"])
		res["diskEncryptionKey"] = transformed
	}

	if v, ok := res["sourceImageEncryptionKey"]; ok {
		original := v.(map[string]interface{})
		transformed := make(map[string]interface{})
		// The raw key won't be returned, so we need to use the original.
		transformed["rawKey"] = d.Get("source_image_encryption_key.0.raw_key")
		transformed["sha256"] = original["sha256"]
		res["sourceImageEncryptionKey"] = transformed
	}

	if v, ok := res["sourceSnapshotEncryptionKey"]; ok {
		original := v.(map[string]interface{})
		transformed := make(map[string]interface{})
		// The raw key won't be returned, so we need to use the original.
		transformed["rawKey"] = d.Get("source_snapshot_encryption_key.0.raw_key")
		transformed["sha256"] = original["sha256"]
		res["sourceSnapshotEncryptionKey"] = transformed
	}

	return res, nil
}
