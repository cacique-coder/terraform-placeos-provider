package placeos

import (
	"context"
	"encoding/json"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSetting() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSettingCreate,
		ReadContext:   resourceSettingRead,
		UpdateContext: resourceSettingUpdate,
		DeleteContext: resourceSettingDelete,
		Schema: map[string]*schema.Schema{
			"parent_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"parent_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"keys": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: true,
			},
			"settings_string": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"encryption_level": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_at": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"updated_at": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func resourceSettingCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	c := m.(*Client)

	setting_string := d.Get("settings_string").(string)
	encryption_level := d.Get("encryption_level").(int)
	parent_type := d.Get("parent_type").(string)
	parent_id := d.Get("parent_id").(string)
	keys_interface := d.Get("keys")

	keys := make([]string, len(keys_interface.([]interface{})))
	for i, v := range keys_interface.([]interface{}) {
		keys[i] = v.(string)
	}

	setting, err := c.createSetting("", parent_id, parent_type, setting_string, encryption_level, keys)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(setting.Id)
	return resourceSettingRead(ctx, d, m)
}

// add a resource driver read function with diagnostic tool for each field
func resourceSettingRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	c := m.(*Client)
	var diags diag.Diagnostics
	id := d.Get("id").(string)
	setting, err := c.getSetting(id)
	if err != nil {
		return diag.FromErr(err)
	}
	d.Set("name", setting.Name)
	d.Set("parent_type", setting.ParentType)
	d.Set("parent_id", setting.ParentId)
	d.Set("keys", setting.Keys)
	d.Set("settings_string", setting.SettingsString)
	d.Set("encryption_level", setting.EncryptionLevel)

	return diags
}

// add a resource driver update function with diagnostic tool for each field
func resourceSettingUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	c := m.(*Client)
	id := d.Get("id").(string)
	setting, err := c.getSetting(id)
	if err != nil {
		return diag.FromErr(err)
	}
	file, err := os.Create("/tmp/setting.json")
	if err != nil {
		return diag.FromErr(err)
	}

	text, err := json.Marshal(setting)

	if err != nil {
		return diag.FromErr(err)
	}

	file.Write(text)
	file.Close()

	if err != nil {
		return diag.FromErr(err)
	}
	// check each field has change and replace it if it has

	setting.Id = id
	if d.HasChange("name") {
		setting.Name = d.Get("name").(string)
	}
	if d.HasChange("parent_type") {
		setting.ParentType = d.Get("parent_type").(string)
	}
	if d.HasChange("parent_id") {
		setting.ParentId = d.Get("parent_id").(string)
	}
	if d.HasChange("keys") {
		setting.Keys = d.Get("keys").([]string)
	}
	if d.HasChange("settings_string") {
		setting.SettingsString = d.Get("settings_string").(string)
	}
	if d.HasChange("encryption_level") {
		setting.EncryptionLevel = d.Get("encryption_level").(int)
	}
	setting, err = c.updateSetting(setting)

	if err != nil {
		return diag.FromErr(err)
	}

	setting, err = c.updateSetting(setting)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceSettingRead(ctx, d, m)
}

// add a resource driver delete function with diagnostic as return
func resourceSettingDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	var diags diag.Diagnostics

	id := d.Get("id").(string)
	err := c.deleteSetting(id)

	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}
