package infrastructuresettings

// This package is created automatically.
// Any manual changes will be overwritten when this is generated.
// []string{"netsec", "service", "v1", "mobileagent", "infrastructuresettings"}

import (
	"context"
	"net/url"

	"github.com/paloaltonetworks/sase-go/api"
	gUGwYZv "github.com/paloaltonetworks/sase-go/netsec/schema/mobile/agent/infrastructure/settings"
)

// Client is the client for this namespace.
type Client struct {
	client api.Client
}

// NewClient returns a new client for this namespace.
func NewClient(client api.Client) *Client {
	return &Client{client: client}
}

// CreateInput takes some input.
// name:"Create" nsfName:"Create" param:0 query:1
// path:
// query: Folder
type CreateInput struct {
	Folder string
	Config gUGwYZv.Config
}

// Create creates the specified object.
//
// Method: post
// URI: /sse/config/v1/mobile-agent/infrastructure-settings
func (c *Client) Create(ctx context.Context, input CreateInput) (gUGwYZv.Config, error) {
	// Variables.
	var err error
	var ans gUGwYZv.Config
	path := "/sse/config/v1/mobile-agent/infrastructure-settings"

	// Query parameter handling.
	uv := url.Values{}
	uv.Set("folder", input.Folder)

	// Execute the command.
	_, err = c.client.Do(ctx, "POST", path, uv, input.Config, &ans)

	// Done.
	return ans, err
}

// InfrastructureSettingsDeleteInput takes some input.
// name:"InfrastructureSettingsDelete" nsfName:"InfrastructureSettingsDelete" param:0 query:2
// path:
// query: Name | Folder
type InfrastructureSettingsDeleteInput struct {
	Name   string
	Folder string
}

// InfrastructureSettingsDelete performs a the given operation.
//
// Method: delete
// URI: /sse/config/v1/mobile-agent/infrastructure-settings
func (c *Client) InfrastructureSettingsDelete(ctx context.Context, input InfrastructureSettingsDeleteInput) error {
	// Variables.
	var err error
	path := "/sse/config/v1/mobile-agent/infrastructure-settings"

	// Query parameter handling.
	uv := url.Values{}
	uv.Set("name", input.Name)
	uv.Set("folder", input.Folder)

	// Execute the command.
	_, err = c.client.Do(ctx, "DELETE", path, uv, nil, nil)

	// Done.
	return err
}

// InfrastructureSettingsPutInput takes some input.
// name:"InfrastructureSettingsPut" nsfName:"InfrastructureSettingsPut" param:0 query:1
// path:
// query: Folder
type InfrastructureSettingsPutInput struct {
	Folder string
	Config gUGwYZv.Config
}

// InfrastructureSettingsPut performs a the given operation.
//
// Method: put
// URI: /sse/config/v1/mobile-agent/infrastructure-settings
func (c *Client) InfrastructureSettingsPut(ctx context.Context, input InfrastructureSettingsPutInput) error {
	// Variables.
	var err error
	path := "/sse/config/v1/mobile-agent/infrastructure-settings"

	// Query parameter handling.
	uv := url.Values{}
	uv.Set("folder", input.Folder)

	// Execute the command.
	_, err = c.client.Do(ctx, "PUT", path, uv, input.Config, nil)

	// Done.
	return err
}

// ListInput takes some input.
// name:"List" nsfName:"List" param:0 query:1
// path:
// query: Folder
type ListInput struct {
	Folder string
}

/*
ListOutput ""

Parent chains:
* "_output"

Args:

Param Data (unordered, list of Config objects): The Data param.

Param Limit (int64): The Limit param. Default: 200

Param Offset (int64): The Offset param. Default: 0

Param Total (int64): The Total param.
*/
type ListOutput struct {
	Data   []gUGwYZv.Config `json:"data,omitempty"`
	Limit  int64            `json:"limit,omitempty"`
	Offset int64            `json:"offset,omitempty"`
	Total  int64            `json:"total,omitempty"`
}

// List gets a list of objects back.
//
// Method: get
// URI: /sse/config/v1/mobile-agent/infrastructure-settings
func (c *Client) List(ctx context.Context, input ListInput) (ListOutput, error) {
	// Variables.
	var err error
	var ans ListOutput
	path := "/sse/config/v1/mobile-agent/infrastructure-settings"

	// Query parameter handling.
	uv := url.Values{}
	uv.Set("folder", input.Folder)

	// Execute the command.
	_, err = c.client.Do(ctx, "GET", path, uv, nil, &ans)

	// Done.
	return ans, err
}
