package placeos

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceRepository() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceRepositoriesRead,
		Schema: map[string]*schema.Schema{
			"repositories": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"created_at": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"updated_at": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"folder_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"uri": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"commit_hash": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"branch": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"repo_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"username": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"password": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceRepositoriesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/engine/v2/repositories", "https://daniel-mac:8443"), nil)
	if err != nil {
		return diag.FromErr(err)
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	req.Header.Add("Authorization", "Bearer eyJhbGciOiJSUzI1NiJ9.eyJpc3MiOiJQT1MiLCJpYXQiOjE2MjQyNzUyMDYsImV4cCI6MTYyNDI4MjQwNiwianRpIjoiMDI4NDk3YTUtMmE5YS00MWFhLTlmOWMtMjliZWNiZDI1ZGRiIiwiYXVkIjoibG9jYWxob3N0Iiwic2NvcGUiOlsicHVibGljIl0sInN1YiI6InVzZXItSEF2R3Z4ZnRVNU0iLCJ1Ijp7Im4iOiJQbGFjZSBTdXBwb3J0IChsb2NhbGhvc3Q6ODQ0MykiLCJlIjoic3VwcG9ydEBwbGFjZS50ZWNoIiwicCI6MiwiciI6W119fQ.Nh6G3ljZXsAsIXTcEeviDhug_lMFqY_laRCs5z8NUN-J9FV4ki4pYYkAR6ldFPtWSJaWGiEVT-PM5lBrH59zM86hXX8F1G5eXNe2SxxZkPFDGJND-aKxcfVRVGqMkWL5xEObCRbwaCA7q1Lz_Lro9dwq54utqkLi-cnX7L4gw0nAYpQkzCY2iNAyIoCoVrP8qB5oVeg55COuEq49wmE9Nlk1ybLt8VBfrAks8pSfFzUp8NiaKB1mfskoABiAXY-j5dAixibPZcvcf3O-FO8q4xoQXnDMf6C35t61c22t9TvfQFye1Bz_2cWC9ADE67QcPMhn29JShPf3bvHrx1e0tA3S7tfYXPFLDIiC3haGhhQNwDP02q9p0CEHwOJDmQb4-2PWWbxTzXsbH2p77yvWRcGcXIf4ZBoVipHe7lDJDMBg0qdMriT54Woby6cSglxBayWJd72gzTzPvCSQ1xiMm8JNrR35Y9NX_oGfHJxjF4DV9yAkB3TljvgqSJbeK6g5V7bEVn5VPFpNIm_Oa1rkD1yEpkRWSBOB-HXsWmfLWakUdt4z5kh9Am9zrPkGXBNMYmjYxxzetTBuoNX8S2IcVMJOQwxlYPbBpqw4QazoFrC5hDOywPTLVapRU3L2uGwUCvCy42NaAiKhkCIsmQM3P3SJIxOkuRom5fXHFhYie1k")

	client := &http.Client{Timeout: 10 * time.Second, Transport: tr}

	r, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer r.Body.Close()

	repositories := make([]map[string]interface{}, 0)

	err = json.NewDecoder(r.Body).Decode(&repositories)

	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("repositories", repositories); err != nil {
		return diag.FromErr(err)
	}
	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
