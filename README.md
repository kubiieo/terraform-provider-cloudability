# kubiieo Terraform Provider for Cloudability

[![Go Reference](https://pkg.go.dev/badge/github.com/kubiieo/terraform-provider-cloudability.svg)](https://pkg.go.dev/github.com/kubiieo/terraform-provider-cloudability)
[![Go Report Card](https://goreportcard.com/badge/github.com/kubiieo/terraform-provider-cloudability)](https://goreportcard.com/report/github.com/kubiieo/terraform-provider-cloudability)
![Github Actions Workflow](https://github.com/kubiieo/terraform-provider-cloudability/actions/workflows/release.yml/badge.svg)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/kubiieo/terraform-provider-cloudability)
![License](https://img.shields.io/dub/l/vibe-d.svg)

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) 0.12+
- [Go](https://golang.org/doc/install) 1.22.7 (to build the provider plugin)
- [cloudability-sdk-go](https://github.com/kubiieo/cloudability-sdk-go) 

## Installing the Provider

The provider is registered in the official [terraform registry](https://registry.terraform.io/providers/kubiieo/cloudability/latest) 

This enables the provider to be auto-installed when you run ```terraform init```

You can also download the latest binary for your target platform from the [releases](https://github.com/kubiieo/terraform-provider-cloudability/releases) tab.

## Building the Provider

- Clone the repo:
    ```sh
    $ mkdir -p terraform-provider-cloudability
    $ cd terraform-provider-cloudability
    $ git clone https://github.com/kubiieo/terraform-provider-cloudability
    ```

- Build the provider: (NOTE: the install directory will allow using this provider by the current user)
    ```sh
    $ go build -o ~/.terraform.d/plugins/terraform-provider-cloudability
    ```

## Local testing

- Clone the repo
- Build in a specific folder
  ```sh
  $ mkdir -p ~/.terraform/plugins/local.providers/local/cloudability/<version>/<darwin_amd64>
  $ go build -o ~/.terraform.d/plugins/local.providers/local/cloudability/<version>/<darwin_arm64>/terraform-provider-cloudability_v<version>
  ```
- Create .terraformrc in root
    ```sh
        provider_installation {
            filesystem_mirror {
                path = "/Users/<user_name>/.terraform.d/plugins"
                include = ["local.providers/*/*"]
            }
            direct {
                exclude = ["local.providers/*/*"]
            }
        }

        disable_checkpoint = true
    ```
- cd to project, set provider as:
  ```sh
    terraform {
        required_providers {
            cloudability = {
                source = "local.providers/local/cloudability"
                version = "1.0.0"
            }
        }
    }

    provider cloudability {
        apikey = "<api_key>"
        region = "au"
    }
  ```
