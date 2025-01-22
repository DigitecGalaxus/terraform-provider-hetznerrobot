// Copyright (c) Digitec Galaxus AG
// SPDX-License-Identifier: MIT

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/ephemeral"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure HetznerRobotProvider satisfies various provider interfaces.
var _ provider.Provider = &HetznerRobotProvider{}
var _ provider.ProviderWithFunctions = &HetznerRobotProvider{}
var _ provider.ProviderWithEphemeralResources = &HetznerRobotProvider{}

// HetznerRobotProvider defines the provider implementation.
type HetznerRobotProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// HetznerRobotProviderModel describes the provider data model.
type HetznerRobotProviderModel struct {
	Username types.String `tfsdk:"username"`
	Password types.String `tfsdk:"password"`
}

func (p *HetznerRobotProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "hetznerrobot"
	resp.Version = p.version
}

func (p *HetznerRobotProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"username": schema.StringAttribute{
				MarkdownDescription: "Username for the robot webservice",
				Required:            true,
			},
			"password": schema.StringAttribute{
				MarkdownDescription: "Password for the robot webservice",
				Required:            true,
			},
		},
	}
}

func (p *HetznerRobotProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config HetznerRobotProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if config.Username.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("hetznerrobot_username"),
			"No username for hetznerrobot.",
			"specify a username for hetznerrobot in your provider configuration. https://robot.hetzner.com/preferences/index",
		)
	}
	if config.Password.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("hetznerrobot_password"),
			"No password for hetznerrobot.",
			"specify a password for hetznerrobot in your provider configuration. https://robot.hetzner.com/preferences/index",
		)
	}

	// Configuration values are now available.
	// if data.Endpoint.IsNull() { /* ... */ }

	// Example client configuration for data sources and resources

	client := NewClient(config.Username.ValueStringPointer(), config.Password.ValueStringPointer())

	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *HetznerRobotProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}

func (p *HetznerRobotProvider) EphemeralResources(ctx context.Context) []func() ephemeral.EphemeralResource {
	return []func() ephemeral.EphemeralResource{}
}

func (p *HetznerRobotProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewServerDataSource,
	}
}

func (p *HetznerRobotProvider) Functions(ctx context.Context) []func() function.Function {
	return []func() function.Function{}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &HetznerRobotProvider{
			version: version,
		}
	}
}
