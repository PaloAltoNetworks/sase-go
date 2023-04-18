package locations

// This package is created automatically.
// Any manual changes will be overwritten when this is generated.
// []string{"netsec", "service", "v1", "mobileagent", "locations"}

import (
	"context"
	"net/url"

	"github.com/paloaltonetworks/sase-go/api"
	xJJfmrr "github.com/paloaltonetworks/sase-go/netsec/schema/mobile/agent/locations"
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
	Data   []xJJfmrr.Config `json:"data,omitempty"`
	Limit  int64            `json:"limit,omitempty"`
	Offset int64            `json:"offset,omitempty"`
	Total  int64            `json:"total,omitempty"`
}

// List gets a list of objects back.
//
// Method: get
// URI: /sse/config/v1/mobile-agent/locations
func (c *Client) List(ctx context.Context, input ListInput) (ListOutput, error) {
	// Variables.
	var err error
	var ans ListOutput
	path := "/sse/config/v1/mobile-agent/locations"

	// Query parameter handling.
	uv := url.Values{}
	uv.Set("folder", input.Folder)

	// Execute the command.
	_, err = c.client.Do(ctx, "GET", path, uv, nil, &ans)

	// Done.
	return ans, err
}

// LocationsPutInput takes some input.
// name:"LocationsPut" nsfName:"LocationsPut" param:0 query:1
// path:
// query: Folder
type LocationsPutInput struct {
	Folder string
	Config xJJfmrr.Config
}

// LocationsPut performs a the given operation.
//
// Method: put
// URI: /sse/config/v1/mobile-agent/locations
func (c *Client) LocationsPut(ctx context.Context, input LocationsPutInput) error {
	// Variables.
	var err error
	path := "/sse/config/v1/mobile-agent/locations"

	// Query parameter handling.
	uv := url.Values{}
	uv.Set("folder", input.Folder)

	// Execute the command.
	_, err = c.client.Do(ctx, "PUT", path, uv, input.Config, nil)

	// Done.
	return err
}
