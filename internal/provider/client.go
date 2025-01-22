// Copyright (c) Digitec Galaxus AG
// SPDX-License-Identifier: MIT

package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"io"
	"net/http"
)

const HostURL string = "https://robot-ws.your-server.de"

type MyClient struct {
	HttpClient *http.Client
	Username   string
	Password   string
}

func NewClient(username *string, password *string) *MyClient {
	c := MyClient{
		HttpClient: &http.Client{},
		Username:   *username,
		Password:   *password,
	}

	return &c
}

func (c *MyClient) GetServers(ctx context.Context) (*[]ServerWrapper, error) {
	r, err := http.NewRequest("GET", HostURL+"/server/", nil)
	if err != nil {
		return nil, err
	}

	r.SetBasicAuth(c.Username, c.Password)

	client := http.Client{}

	response, err := client.Do(r)
	if err != nil {
		return nil, err
	}

	if response.StatusCode == 404 {
		return &[]ServerWrapper{}, nil
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected status code %d", response.StatusCode)
	}

	bodyBytes, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	tflog.Trace(ctx, "content of json: "+string(bodyBytes))

	var servers []ServerWrapper
	err = json.Unmarshal(bodyBytes, &servers)

	if err != nil {
		return nil, err
	}

	return &servers, nil
}

type ServerWrapper struct {
	Server HetznerRobotServer `json:"server"`
}

type HetznerRobotServer struct {
	ServerId   int32  `json:"server_number"`
	ServerName string `json:"server_name"`
	ServerIpv4 string `json:"server_ip"`
	ServerIpv6 string `json:"server_ipv6_net"`
	Product    string `json:"product"`
	Datacenter string `json:"dc"`
	Traffic    string `json:"traffic"`
	Status     string `json:"status"`
	Cancelled  bool   `json:"cancelled"`
	PaidUntil  string `json:"paid_until"`
}
