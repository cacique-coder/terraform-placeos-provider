package placeos

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDriver() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceDriverCreate,
		ReadContext:   resourceDriverRead,
		UpdateContext: resourceDriverUpdate,
		DeleteContext: resourceDriverDelete,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"file_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"default_uri": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"module_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"repository_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"commit": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"role": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_at": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ignored_connected": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
				Optional: true,
			},
			"updated_at": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func resourceDriverCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	c := m.(*Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	file_name := d.Get("file_name").(string)
	default_uri := d.Get("default_uri").(string)
	commit := d.Get("commit").(string)
	role := d.Get("role").(int)
	module_name := d.Get("module_name").(string)
	repository_id := d.Get("repository_id").(string)
	ignore_connected := d.Get("ignored_connected").(bool)

	driver, err := c.createDriver(name, description, file_name, default_uri, module_name, repository_id, commit, role, ignore_connected)

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(driver.Id)
	return resourceDriverRead(ctx, d, m)
}

// add a resource driver read function with diagnostic tool for each field
func resourceDriverRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	c := m.(*Client)
	var diags diag.Diagnostics
	id := d.Get("id").(string)
	driver, err := c.getDriver(id)
	if err != nil {
		return diag.FromErr(err)
	}
	d.Set("name", driver.Name)
	d.Set("file_name", driver.FileName)
	d.Set("default_uri", driver.DefaultUri)
	d.Set("module_name", driver.ModuleName)
	d.Set("description", driver.Description)
	d.Set("repository_id", driver.RepositoryId)
	d.Set("commit", driver.Commit)
	d.Set("role", driver.Role)
	d.Set("created_at", driver.CreatedAt)
	d.Set("ignored_connected", driver.IgnoredConnected)
	d.Set("updated_at", driver.UpdatedAt)

	return diags
}

// add a resource driver update function with diagnostic tool for each field
func resourceDriverUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	c := m.(*Client)
	id := d.Get("id").(string)
	driver, err := c.getDriver(id)

	if err != nil {
		return diag.FromErr(err)
	}
	// check each field has change and replace it if it has
	if d.HasChange("name") {
		driver.Name = d.Get("name").(string)
	}
	if d.HasChange("file_name") {
		driver.FileName = d.Get("file_name").(string)
	}
	if d.HasChange("default_uri") {
		driver.DefaultUri = d.Get("default_uri").(string)
	}
	if d.HasChange("module_name") {
		driver.ModuleName = d.Get("module_name").(string)
	}
	if d.HasChange("description") {
		driver.Description = d.Get("description").(string)
	}
	if d.HasChange("repository_id") {
		driver.RepositoryId = d.Get("repository_id").(string)
	}
	if d.HasChange("commit") {
		driver.Commit = d.Get("commit").(string)
	}
	if d.HasChange("role") {
		driver.Role = d.Get("role").(int)
	}
	if d.HasChange("ignored_connected") {
		driver.IgnoredConnected = d.Get("ignored_connected").(bool)
	}

	// update driver
	err = c.updateDriver(driver)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceDriverRead(ctx, d, m)
}

// add a resource driver delete function with diagnostic as return
func resourceDriverDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	var diags diag.Diagnostics

	id := d.Get("id").(string)
	err := c.deleteDriver(id)

	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}
