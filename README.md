# Terraform Provider Hetzner Robot (Webservice)

This terraform provider can be used to manage hetzner bare metal servers over their webservice api. The implemented functionality is limited for the use cases that we need and is not a full implementation over all their apis.

## Requirements

- [Terraform](https://developer.hashicorp.com/terraform/downloads) >= 1.0
- [Go](https://golang.org/doc/install) >= 1.22

## Building The Provider

1. Clone the repository
1. Enter the repository directory
1. Build the provider using the Go `install` command:

```shell
go install
```

## Adding Dependencies

This provider uses [Go modules](https://github.com/golang/go/wiki/Modules).
Please see the Go documentation for the most up to date information about using Go modules.

To add a new dependency `github.com/author/dependency` to your Terraform provider:

```shell
go get github.com/author/dependency
go mod tidy
```

Then commit the changes to `go.mod` and `go.sum`.

## Using the provider

[docs](https://github.com/DigitecGalaxus/terraform-provider-hetznerrobot/tree/master/docs)

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (see [Requirements](#requirements) above).

To compile the provider, run `go install`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

To generate or update documentation, run `make generate`.

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```shell
make testacc
```

## Releasing a new version

The action in this repository is triggered by creating a git tag with a vX.Y.Z format. There is also a github action setup in the org that communicates to the terraform registry via a [webhook](https://developer.hashicorp.com/terraform/registry/providers/publishing#webhooks)
