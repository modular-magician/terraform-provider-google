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
	"encoding/base64"
	"fmt"
	"log"
	"reflect"
	"regexp"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func ResourceKMSSecretCiphertext() *schema.Resource {
	return &schema.Resource{
		Create: resourceKMSSecretCiphertextCreate,
		Read:   resourceKMSSecretCiphertextRead,
		Delete: resourceKMSSecretCiphertextDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"crypto_key": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The full name of the CryptoKey that will be used to encrypt the provided plaintext.
Format: ''projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}/cryptoKeys/{{cryptoKey}}''`,
			},
			"plaintext": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The plaintext to be encrypted.`,
				Sensitive:   true,
			},
			"additional_authenticated_data": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `The additional authenticated data used for integrity checks during encryption and decryption.`,
				Sensitive:   true,
			},
			"ciphertext": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Contains the result of encrypting the provided plaintext, encoded in base64.`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceKMSSecretCiphertextCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	plaintextProp, err := expandKMSSecretCiphertextPlaintext(d.Get("plaintext"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("plaintext"); !isEmptyValue(reflect.ValueOf(plaintextProp)) && (ok || !reflect.DeepEqual(v, plaintextProp)) {
		obj["plaintext"] = plaintextProp
	}
	additionalAuthenticatedDataProp, err := expandKMSSecretCiphertextAdditionalAuthenticatedData(d.Get("additional_authenticated_data"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("additional_authenticated_data"); !isEmptyValue(reflect.ValueOf(additionalAuthenticatedDataProp)) && (ok || !reflect.DeepEqual(v, additionalAuthenticatedDataProp)) {
		obj["additionalAuthenticatedData"] = additionalAuthenticatedDataProp
	}

	url, err := ReplaceVars(d, config, "{{KMSBasePath}}{{crypto_key}}:encrypt")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new SecretCiphertext: %#v", obj)
	billingProject := ""

	if parts := regexp.MustCompile(`projects\/([^\/]+)\/`).FindStringSubmatch(url); parts != nil {
		billingProject = parts[1]
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating SecretCiphertext: %s", err)
	}

	// Store the ID now
	id, err := ReplaceVars(d, config, "{{crypto_key}}/{{ciphertext}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// we don't set anything on read and instead do it all in create
	ciphertext, ok := res["ciphertext"]
	if !ok {
		return fmt.Errorf("Create response didn't contain critical fields. Create may not have succeeded.")
	}
	if err := d.Set("ciphertext", ciphertext.(string)); err != nil {
		return fmt.Errorf("Error setting ciphertext: %s", err)
	}

	id, err = ReplaceVars(d, config, "{{crypto_key}}/{{ciphertext}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating SecretCiphertext %q: %#v", d.Id(), res)

	return resourceKMSSecretCiphertextRead(d, meta)
}

func resourceKMSSecretCiphertextRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{KMSBasePath}}{{crypto_key}}")
	if err != nil {
		return err
	}

	billingProject := ""

	if parts := regexp.MustCompile(`projects\/([^\/]+)\/`).FindStringSubmatch(url); parts != nil {
		billingProject = parts[1]
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("KMSSecretCiphertext %q", d.Id()))
	}

	res, err = resourceKMSSecretCiphertextDecoder(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Decoding the object has resulted in it being gone. It may be marked deleted
		log.Printf("[DEBUG] Removing KMSSecretCiphertext because it no longer exists.")
		d.SetId("")
		return nil
	}

	return nil
}

func resourceKMSSecretCiphertextDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[WARNING] KMS SecretCiphertext resources"+
		" cannot be deleted from Google Cloud. The resource %s will be removed from Terraform"+
		" state, but will still be present on Google Cloud.", d.Id())
	d.SetId("")

	return nil
}

func expandKMSSecretCiphertextPlaintext(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	if v == nil {
		return nil, nil
	}

	return base64.StdEncoding.EncodeToString([]byte(v.(string))), nil
}

func expandKMSSecretCiphertextAdditionalAuthenticatedData(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	if v == nil {
		return nil, nil
	}

	return base64.StdEncoding.EncodeToString([]byte(v.(string))), nil
}

func resourceKMSSecretCiphertextDecoder(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	return res, nil
}
