package placeos

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRepository() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceRepositoryCreate,
		ReadContext:   resourceRepositoryRead,
		UpdateContext: resourceRepositoryUpdate,
		DeleteContext: resourceRepositoryDelete,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"folder_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"uri": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"repo_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"branch": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "master",
			},
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"password": &schema.Schema{
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_at": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"commit_hash": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"updated_at": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func resourceRepositoryCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	c := m.(*Client)
	var diags diag.Diagnostics

	name := d.Get("name").(string)
	folder_name := d.Get("folder_name").(string)
	uri := d.Get("uri").(string)
	repo_type := d.Get("repo_type").(string)
	description := d.Get("description").(string)
	branch := d.Get("branch").(string)
	username := d.Get("username").(string)
	password := d.Get("password").(string)

	repository, err := c.createRepository(name, folder_name, uri, repo_type, description, branch, username, password)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(repository.Id)
	resourceRepositoryRead(ctx, d, m)
	return diags
}

func resourceRepositoryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	c := m.(*Client)
	var diags diag.Diagnostics

	repositoryId := d.Id()

	repository, err := c.getRepository(repositoryId)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("created_at", repository.CreatedAt); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("updated_at", repository.UpdatedAt); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("name", repository.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("description", repository.Description); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("folder_name", repository.FolderName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("uri", repository.Uri); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("commit_hash", repository.CommitHash); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("branch", repository.Branch); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("username", repository.Username); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("password", repository.Password); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("repo_type", repository.RepoType); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceRepositoryUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	repository, err := c.getRepository(d.Id())

	if d.HasChange("name") {
		name := d.Get("name").(string)
		repository.Name = name
	}

	if d.HasChange("folder_name") {
		folder_name := d.Get("folder_name").(string)
		repository.FolderName = folder_name
	}

	if d.HasChange("uri") {
		uri := d.Get("uri").(string)
		repository.Uri = uri
	}

	if d.HasChange("description") {
		description := d.Get("description").(string)
		repository.Description = description
	}

	if d.HasChange("branch") {
		branch := d.Get("branch").(string)
		repository.Branch = branch
	}

	if d.HasChange("username") {
		username := d.Get("username").(string)
		repository.Username = username
	}

	if d.HasChange("password") {
		password := d.Get("password").(string)
		repository.Password = password
	}

	if d.HasChange("repo_type") {
		repo_type := d.Get("repo_type").(string)
		repository.RepoType = repo_type
	}

	repository2, err := c.updateRepository(repository)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(repository2.Id)

	return resourceRepositoryRead(ctx, d, m)
}

func resourceRepositoryDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	c := m.(*Client)
	var diags diag.Diagnostics

	id := d.Get("id").(string)

	err := c.deleteRepository(id)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return diags
}
