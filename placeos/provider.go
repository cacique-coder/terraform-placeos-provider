package placeos

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("PLACEOS_USERNAME", ""),
			},

			"password": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("PLACEOS_PASSWORD", ""),
			},

			"host": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("PLACEOS_HOST", ""),
			},

			"client_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				// DefaultFunc: schema.EnvDefaultFunc("PLACEOS_CLIENT_ID", ""),
			},
			"client_secret": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("PLACEOS_CLIENT_SECRET", "secret"),
			},

			"insecure_ssl": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("PLACEOS_CLIENT_INSECURE_SSL", false),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"placeos_repository": resourceRepository(),
			"placeos_driver":     resourceDriver(),
			"placeos_module":     resourceModule(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"placeos_repositories": dataSourceRepository(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	username := d.Get("username").(string)
	password := d.Get("password").(string)
	host := d.Get("host").(string)
	insecureSsl := d.Get("insecure_ssl").(bool)
	clientId := d.Get("client_id").(string)
	clientSecret := d.Get("client_secret").(string)
	// Warning or errors can be collected in a slice type

	fmt.Println("insecure value")
	fmt.Println(clientId)

	fmt.Println("insecure value")
	fmt.Println(insecureSsl)

	var diags diag.Diagnostics

	client := NewBasicAuthClient(username, password, host, insecureSsl, clientId, clientSecret)
	return client, diags
}
