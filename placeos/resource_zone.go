package placeos

import (
	"context"
	"encoding/json"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceZone() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceZoneCreate,
		ReadContext:   resourceZoneRead,
		UpdateContext: resourceZoneUpdate,
		DeleteContext: resourceZoneDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  nil,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  nil,
			},
			"code": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  nil,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  nil,
			},
			"location": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  nil,
			},
			"count_field": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  nil,
			},
			"capacity": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  nil,
			},
			"map_id": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  nil,
			},
			"parent_id": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  nil,
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

// "name" "description" "display_name" "code" "type" "count" "capacity" "map_id" "parent_id" "tags"

func resourceZoneCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	c := m.(*Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	display_name := d.Get("display_name").(string)
	code := d.Get("code").(string)
	type_ := d.Get("type").(string)
	count := d.Get("count_field").(int)
	capacity := d.Get("capacity").(int)
	location := d.Get("location").(string)
	map_id := d.Get("map_id").(string)
	parent_id := d.Get("parent_id").(string)
	tags_interface := d.Get("tags")

	tags := make([]string, len(tags_interface.([]interface{})))
	for i, v := range tags_interface.([]interface{}) {
		tags[i] = v.(string)
	}

	setting, err := c.CreateZone(name, description, tags, location, display_name, code, type_, count, capacity, map_id, parent_id)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(setting.Id)
	return resourceZoneRead(ctx, d, m)
}

// add a resource driver read function with diagnostic tool for each field
func resourceZoneRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	c := m.(*Client)
	var diags diag.Diagnostics
	id := d.Get("id").(string)
	zone, err := c.GetZone(id)
	if err != nil {
		return diag.FromErr(err)
	}
	d.Set("name", zone.Name)
	d.Set("description", zone.Description)
	d.Set("display_name", zone.DisplayName)
	d.Set("code", zone.Code)
	d.Set("type", zone.Type)
	d.Set("count_field", zone.Count)
	d.Set("capacity", zone.Capacity)
	d.Set("location", zone.Location)
	d.Set("map_id", zone.MapId)
	d.Set("parent_id", zone.ParentId)
	d.Set("tags", zone.Tags)
	d.Set("id", zone.Id)
	d.Set("created_at", zone.CreatedAt)
	d.Set("updated_at", zone.UpdatedAt)

	return diags
}

// add a resource driver update function with diagnostic tool for each field
func resourceZoneUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	c := m.(*Client)
	id := d.Get("id").(string)
	zone, err := c.GetZone(id)
	if err != nil {
		return diag.FromErr(err)
	}
	file, err := os.Create("/tmp/zone.json")
	if err != nil {
		return diag.FromErr(err)
	}

	text, err := json.Marshal(zone)

	if err != nil {
		return diag.FromErr(err)
	}

	file.Write(text)
	file.Close()

	if err != nil {
		return diag.FromErr(err)
	}
	// check each field has change and replace it if it has

	zone.Id = id
	if d.HasChange("name") {
		zone.Name = d.Get("name").(string)
	}
	if d.HasChange("description") {
		zone.Description = d.Get("description").(string)
	}
	if d.HasChange("display_name") {
		zone.DisplayName = d.Get("display_name").(string)
	}
	if d.HasChange("code") {
		zone.Code = d.Get("code").(string)
	}
	if d.HasChange("type") {
		zone.Type = d.Get("type").(string)
	}
	if d.HasChange("count") {
		zone.Count = d.Get("count").(int)
	}
	if d.HasChange("capacity") {
		zone.Capacity = d.Get("capacity").(int)
	}
	if d.HasChange("location") {
		zone.Location = d.Get("location").(string)
	}
	if d.HasChange("map_id") {
		zone.MapId = d.Get("map_id").(string)
	}
	if d.HasChange("parent_id") {
		zone.ParentId = d.Get("parent_id").(string)
	}
	if d.HasChange("tags") {
		tags := make([]string, len(d.Get("tags").([]interface{})))
		for i, v := range d.Get("tags").([]interface{}) {
			tags[i] = v.(string)
		}
		zone.Tags = tags
	}

	zone, err = c.UpdateZone(zone)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceZoneRead(ctx, d, m)
}

// add a resource driver delete function with diagnostic as return
func resourceZoneDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	var diags diag.Diagnostics

	id := d.Get("id").(string)
	err := c.deleteZone(id)

	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}
