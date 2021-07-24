# Terraform Provider Placeos

Run the following command to build the provider

```shell
go build -o terraform-provider-placeos
```

## Test sample configuration

First, build and install the provider.

```shell
make install
go build -o terraform-provider-hashicups
mv terraform-provider-hashicups ~/.terraform.d/plugins/hashicorp.com/edu/hashicups/0.2/darwin_amd64
```

Then, run the following command to initialize the workspace and apply the sample configuration.

```shell
terraform init && terraform apply
```