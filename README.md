# Terraform Provider Placeos

## Warning

This is a hobby project at the moment, I can't guarantee no breaking changes on futures releases.

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


## Roadmap

version 0.0.1 in progress
  - A basic Proof of concept. We need to be able to create Systems, drivers, modules, settings, zones and triggers. It is not expected to be usable, for instance roles should be "ssh", "services", etc. Instead we are using a plain int number.

Version 0.1.0
  - Adding basic validations. Module name