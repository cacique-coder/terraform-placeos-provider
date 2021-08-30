package placeos

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystem() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSystemCreate,
		ReadContext:   resourceSystemRead,
		UpdateContext: resourceSystemUpdate,
		DeleteContext: resourceSystemDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"email": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"code": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"timezone": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"support_url": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"map_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"bookable": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"version": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"installed_ui_devices": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"capacity": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"images": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},

			"zones": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: true,
			},
			"modules": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
			"features": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_at": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"updated_at": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func resourceSystemCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	c := m.(*Client)
	var diags diag.Diagnostics

	//  "description" 	"email" 	"display_name" 	"code" 	"timezone" 	"support_url" 	"map_id" 	"bookable" 	"version" 	"installed_ui_devices" "capacity"
	name := d.Get("name").(string)
	description := d.Get("description").(string)
	email := d.Get("email").(string)
	displayName := d.Get("display_name").(string)
	code := d.Get("code").(string)
	timezone := d.Get("timezone").(string)
	supportUrl := d.Get("support_url").(string)
	mapId := d.Get("map_id").(string)
	bookable := d.Get("bookable").(bool)
	version := d.Get("version").(int)
	installedUiDevices := d.Get("installed_ui_devices").(int)
	capacity := d.Get("capacity").(int)

	// "images" "zones" "modules" "features"
	images_interface := d.Get("images")

	images := make([]string, len(images_interface.([]interface{})))
	for i, v := range images_interface.([]interface{}) {
		images[i] = v.(string)
	}

	zones_interface := d.Get("zones")
	zoneIds := make([]string, len(zones_interface.([]interface{})))
	for i, v := range zones_interface.([]interface{}) {
		zoneIds[i] = v.(string)
	}

	modules_interface := d.Get("modules")
	moduleIds := make([]string, len(modules_interface.([]interface{})))
	for i, v := range modules_interface.([]interface{}) {
		moduleIds[i] = v.(string)
	}

	features_interface := d.Get("features")
	features := make([]string, len(features_interface.([]interface{})))
	for i, v := range features_interface.([]interface{}) {
		features[i] = v.(string)
	}

	system, err := c.CreateSystem(name, zoneIds, email, displayName, supportUrl, int64(installedUiDevices), int64(capacity), bookable, description, features, mapId, moduleIds, timezone, code, int64(version), images)

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(system.Id)
	resourceSystemRead(ctx, d, m)
	return diags
}

func resourceSystemRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	c := m.(*Client)
	var diags diag.Diagnostics

	systemId := d.Id()

	system, err := c.GetSystem(systemId)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("name", system.Name)
	d.Set("description", system.Description)
	d.Set("email", system.Email)
	d.Set("display_name", system.DisplayName)
	d.Set("code", system.Code)
	d.Set("timezone", system.Timezone)
	d.Set("support_url", system.SupportUrl)
	d.Set("map_id", system.MapId)
	d.Set("bookable", system.Bookable)
	d.Set("version", system.Version)
	d.Set("installed_ui_devices", system.InstalledUiDevices)
	d.Set("capacity", system.Capacity)
	d.Set("images", system.Images)
	d.Set("zones", system.Zones)
	d.Set("modules", system.Modules)
	d.Set("features", system.Features)
	d.Set("id", system.Id)
	d.Set("created_at", system.CreatedAt)
	d.Set("updated_at", system.UpdatedAt)

	return diags
}

func resourceSystemUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	system, err := c.GetSystem(d.Id())

	if err != nil {
		return diag.FromErr(err)
	}

	if d.HasChange("name") {
		name := d.Get("name").(string)
		system.Name = name
	}
	if d.HasChange("description") {
		description := d.Get("description").(string)
		system.Description = description
	}
	if d.HasChange("email") {
		email := d.Get("email").(string)
		system.Email = email
	}
	if d.HasChange("display_name") {
		displayName := d.Get("display_name").(string)
		system.DisplayName = displayName
	}
	if d.HasChange("code") {
		code := d.Get("code").(string)
		system.Code = code
	}
	if d.HasChange("timezone") {
		timezone := d.Get("timezone").(string)
		system.Timezone = timezone
	}
	if d.HasChange("support_url") {
		supportUrl := d.Get("support_url").(string)
		system.SupportUrl = supportUrl
	}
	if d.HasChange("map_id") {
		mapId := d.Get("map_id").(string)
		system.MapId = mapId
	}
	if d.HasChange("bookable") {
		bookable := d.Get("bookable").(bool)
		system.Bookable = bookable
	}
	if d.HasChange("version") {
		version := d.Get("version").(int)
		system.Version = int64(version)
	}
	if d.HasChange("installed_ui_devices") {
		installedUiDevices := d.Get("installed_ui_devices").(int)
		system.InstalledUiDevices = int64(installedUiDevices)
	}
	if d.HasChange("capacity") {
		capacity := d.Get("capacity").(int)
		system.Capacity = int64(capacity)
	}
	if d.HasChange("zones") {
		zones_interface := d.Get("zones")
		zoneIds := make([]string, len(zones_interface.([]interface{})))
		for i, v := range zones_interface.([]interface{}) {
			zoneIds[i] = v.(string)
		}
		system.Zones = zoneIds
	}
	if d.HasChange("modules") {
		modules_interface := d.Get("modules")
		moduleIds := make([]string, len(modules_interface.([]interface{})))
		for i, v := range modules_interface.([]interface{}) {
			moduleIds[i] = v.(string)
		}
		system.Modules = moduleIds
	}
	if d.HasChange("features") {
		features_interface := d.Get("features")
		features := make([]string, len(features_interface.([]interface{})))
		for i, v := range features_interface.([]interface{}) {
			features[i] = v.(string)
		}
		system.Features = features
	}

	if d.HasChange("images") {
		images_interface := d.Get("images")
		images := make([]string, len(images_interface.([]interface{})))
		for i, v := range images_interface.([]interface{}) {
			images[i] = v.(string)
		}
		system.Images = images
	}

	d.SetId(system.Id)

	c.UpdateSystem(system)

	return resourceSystemRead(ctx, d, m)
}

func resourceSystemDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	c := m.(*Client)
	var diags diag.Diagnostics

	id := d.Get("id").(string)

	err := c.DeleteSystem(id)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return diags
}
