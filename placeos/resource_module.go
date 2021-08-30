package placeos

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceModule() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceModuleCreate,
		ReadContext:   resourceModuleRead,
		UpdateContext: resourceModuleUpdate,
		DeleteContext: resourceModuleDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"custom_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"driver_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"uri": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"notes": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"makebreak": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"ignore_connected": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"ignore_starstop": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"tls": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"udp": {
				Type:     schema.TypeBool,
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

func resourceModuleCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	c := m.(*Client)
	var diags diag.Diagnostics

	ip := d.Get("ip").(string)
	port := d.Get("port").(int)
	driverId := d.Get("driver_id").(string)
	tls := d.Get("tls").(bool)
	udp := d.Get("udp").(bool)
	makebreak := d.Get("makebreak").(bool)
	uri := d.Get("uri").(string)
	customName := d.Get("custom_name").(string)
	notes := d.Get("notes").(string)
	ignore_connected := d.Get("ignore_connected").(bool)

	module, err := c.createModule(ip, driverId, uri, port, tls, udp, makebreak, customName, notes, ignore_connected)

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(module.Id)
	resourceModuleRead(ctx, d, m)
	return diags
}

func resourceModuleRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	c := m.(*Client)
	var diags diag.Diagnostics

	moduleId := d.Id()

	module, err := c.getModule(moduleId)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("name", module.Name)
	d.Set("custom_name", module.CustomName)
	d.Set("driver_id", module.DriverId)
	d.Set("uri", module.Uri)
	d.Set("notes", module.Notes)
	d.Set("ip", module.Ip)
	d.Set("port", module.Port)
	d.Set("tls", module.Tls)
	d.Set("udp", module.Udp)
	d.Set("makebreak", module.Makebreak)
	d.Set("ignore_connected", module.IgnoreConnected)
	d.Set("ignore_startstop", module.IgnoreStartStop)
	d.Set("id", module.Id)
	d.Set("created_at", module.CreatedAt)
	d.Set("updated_at", module.UpdatedAt)

	return diags
}

func resourceModuleUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	module, err := c.getModule(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	module.Id = d.Id()

	if d.HasChange("custom_name") {
		custom_name := d.Get("custom_name").(string)
		module.CustomName = custom_name
	}

	if d.HasChange("ip") {
		ip := d.Get("ip").(string)
		module.Ip = ip
	}
	if d.HasChange("port") {
		port := d.Get("port").(int)
		module.Port = port
	}
	if d.HasChange("tls") {
		tls := d.Get("tls").(bool)
		module.Tls = tls
	}
	if d.HasChange("udp") {
		udp := d.Get("udp").(bool)
		module.Udp = udp
	}
	if d.HasChange("makebreak") {
		makebreak := d.Get("makebreak").(bool)
		module.Makebreak = makebreak
	}
	if d.HasChange("notes") {
		notes := d.Get("notes").(string)
		module.Notes = notes
	}
	if d.HasChange("ignore_connected") {
		ignore_connected := d.Get("ignore_connected").(bool)
		module.IgnoreConnected = ignore_connected
	}
	if d.HasChange("ignore_start_stop") {
		ignore_start_stop := d.Get("ignore_start_stop").(bool)
		module.IgnoreStartStop = ignore_start_stop
	}
	if d.HasChange("notes") {
		notes := d.Get("notes").(string)
		module.Notes = notes
	}
	if d.HasChange("ignore_connected") {
		ignore_connected := d.Get("ignore_connected").(bool)
		module.IgnoreConnected = ignore_connected
	}

	module2, err := c.updateModule(module)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(module2.Id)

	return resourceModuleRead(ctx, d, m)
}

func resourceModuleDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	c := m.(*Client)
	var diags diag.Diagnostics

	id := d.Get("id").(string)

	err := c.deleteModule(id)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return diags
}
