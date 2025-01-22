// Copyright (c) Digitec Galaxus AG
// SPDX-License-Identifier: MIT

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

type ServerDataSource struct {
	client *MyClient
}

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &ServerDataSource{}

func NewServerDataSource() datasource.DataSource {
	return &ServerDataSource{}
}

type ServersDatasourceModel struct {
	Servers []ServerDatasourceModel `tfsdk:"servers"`
}

type ServerDatasourceModel struct {
	ServerId   types.Int32  `tfsdk:"server_number"`
	ServerName types.String `tfsdk:"server_name"`
	ServerIpv4 types.String `tfsdk:"server_ipv4"`
	ServerIpv6 types.String `tfsdk:"server_ipv6"`
	Product    types.String `tfsdk:"product"`
	Datacenter types.String `tfsdk:"datacenter"`
	Traffic    types.String `tfsdk:"traffic"`
	Status     types.String `tfsdk:"status"`
	Cancelled  types.Bool   `tfsdk:"cancelled"`
	PaidUntil  types.String `tfsdk:"paid_until"`
}

func (d *ServerDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_servers"
}

func (d *ServerDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.

		Attributes: map[string]schema.Attribute{
			"servers": schema.ListNestedAttribute{
				MarkdownDescription: "the list of all your hetzner servers. calls GET /server.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"server_number": schema.Int32Attribute{Computed: true},
						"server_name":   schema.StringAttribute{Computed: true},
						"server_ipv4":   schema.StringAttribute{Computed: true},
						"server_ipv6":   schema.StringAttribute{Computed: true},
						"product":       schema.StringAttribute{Computed: true},
						"datacenter":    schema.StringAttribute{Computed: true},
						"traffic":       schema.StringAttribute{Computed: true},
						"status":        schema.StringAttribute{Computed: true},
						"cancelled":     schema.BoolAttribute{Computed: true},
						"paid_until":    schema.StringAttribute{Computed: true},
					},
				},
			},
		},
	}
}

func (d *ServerDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*MyClient)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *hashicups.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}

func (d *ServerDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data ServersDatasourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Trace(ctx, "read servers endpoint of hetzner robot username: "+d.client.Username)

	servers, err := d.client.GetServers(ctx)
	if err != nil {
		resp.Diagnostics.AddError("reading servers failed", err.Error())
	}

	if servers == nil {
		resp.Diagnostics.AddError("reading servers failed", "servers is nil for some reason")
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.

	// For the purposes of this example code, hardcoding a response value to
	// save into the Terraform state.

	for _, server := range *servers {
		serverState := ServerDatasourceModel{
			ServerId:   types.Int32Value(server.Server.ServerId),
			ServerName: types.StringValue(server.Server.ServerName),
			ServerIpv4: types.StringValue(server.Server.ServerIpv4),
			ServerIpv6: types.StringValue(server.Server.ServerIpv6),
			Product:    types.StringValue(server.Server.Product),
			Datacenter: types.StringValue(server.Server.Datacenter),
			Traffic:    types.StringValue(server.Server.Traffic),
			Status:     types.StringValue(server.Server.Status),
			Cancelled:  types.BoolValue(server.Server.Cancelled),
			PaidUntil:  types.StringValue(server.Server.PaidUntil),
		}

		data.Servers = append(data.Servers, serverState)
	}

	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
