package sharedinfrastructuresettings

// This package is created automatically.
// Any manual changes will be overwritten when this is generated.
// []string{"netsec", "service", "v1", "sharedinfrastructuresettings"}

import (
	"context"
	"net/url"
	"strconv"

	"github.com/paloaltonetworks/sase-go/api"
	ebBZfXA "github.com/paloaltonetworks/sase-go/netsec/schema/edit/shared/infrastructure/settings"
	gulqfOv "github.com/paloaltonetworks/sase-go/netsec/schema/shared/infrastructure/settings"
)

// Client is the client for this namespace.
type Client struct {
	client api.Client
}

// NewClient returns a new client for this namespace.
func NewClient(client api.Client) *Client {
	return &Client{client: client}
}

// ListInput takes some input.
// name:"List" nsfName:"List" param:0 query:2
// path: []string(nil)
// query: []string{"limit-optional", "offset-optional"}
type ListInput struct {
	Limit  *int64
	Offset *int64
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
	Data   []gulqfOv.Config `json:"data,omitempty"`
	Limit  int64            `json:"limit,omitempty"`
	Offset int64            `json:"offset,omitempty"`
	Total  int64            `json:"total,omitempty"`
}

// List gets a list of objects back.
//
// Method: get
// URI: /sse/config/v1/shared-infrastructure-settings
func (c *Client) List(ctx context.Context, input ListInput) (ListOutput, error) {
	// Variables.
	var err error
	var ans ListOutput
	path := "/sse/config/v1/shared-infrastructure-settings"

	// Query parameter handling.
	uv := url.Values{}
	if input.Limit != nil {
		uv.Set("limit", strconv.Itoa(int(*input.Limit)))
	}
	if input.Offset != nil {
		uv.Set("offset", strconv.Itoa(int(*input.Offset)))
	}

	// Execute the command.
	_, err = c.client.Do(ctx, "GET", path, uv, nil, &ans)

	// Done.
	return ans, err
}

// SharedInfrastructureSettingsPutInput takes some input.
// name:"SharedInfrastructureSettingsPut" nsfName:"SharedInfrastructureSettingsPut" param:0 query:0
// path: []string(nil)
// query: []string(nil)
type SharedInfrastructureSettingsPutInput struct {
	Config ebBZfXA.Config
}

// SharedInfrastructureSettingsPut performs a the given operation.
//
// Method: put
// URI: /sse/config/v1/shared-infrastructure-settings
func (c *Client) SharedInfrastructureSettingsPut(ctx context.Context, input SharedInfrastructureSettingsPutInput) error {
	// Variables.
	var err error
	path := "/sse/config/v1/shared-infrastructure-settings"

	// Execute the command.
	_, err = c.client.Do(ctx, "PUT", path, nil, input.Config, nil)

	// Done.
	return err
}
