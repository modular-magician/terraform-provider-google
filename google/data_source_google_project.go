package google

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceGoogleProject() *schema.Resource {
	// Generate datasource schema from resource
	dsSchema := datasourceSchemaFromResourceSchema(resourceGoogleProject().Schema)

	addOptionalFieldsToSchema(dsSchema, "project_id")

	return &schema.Resource{
		Read:   datasourceGoogleProjectRead,
		Schema: dsSchema,
	}
}

func datasourceGoogleProjectRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	if v, ok := d.GetOk("project_id"); ok {
		project := v.(string)
		d.SetId(fmt.Sprintf("projects/%s", project))
	} else {
		project, err := getProject(d, config)
		if err != nil {
			return err
		}
		d.SetId(fmt.Sprintf("projects/%s", project))
	}

	return resourceGoogleProjectRead(d, meta)
}
