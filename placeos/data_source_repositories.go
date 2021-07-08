package placeos

import (
	"context"
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
	c := m.(*Client)

	var diags diag.Diagnostics

	repositories, err := c.getRepositories()

	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("repositories", repoTerraform(&repositories)); err != nil {
		return diag.FromErr(err)
	}
	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func repoTerraform(repositories *[]Repository) []interface{} {
	if repositories != nil {
		j_repositories := make([]interface{}, len(*repositories))

		for i, repository := range *repositories {
			j_repo := make(map[string]interface{})

			j_repo["created_at"] = repository.CreatedAt
			j_repo["updated_at"] = repository.UpdatedAt
			j_repo["name"] = repository.Name
			j_repo["description"] = repository.Description
			j_repo["folder_name"] = repository.FolderName
			j_repo["uri"] = repository.Uri
			j_repo["commit_hash"] = repository.CommitHash
			j_repo["branch"] = repository.CommitHash
			j_repo["branch"] = repository.CommitHash
			j_repo["repo_type"] = repository.RepoType
			j_repo["id"] = repository.Id
			j_repo["username"] = repository.Username
			j_repo["password"] = repository.Password

			j_repositories[i] = j_repo

		}
		return j_repositories
	}

	return make([]interface{}, 0)
}
