package bgprouting

// This package is created automatically.
// Any manual changes will be overwritten when this is generated.
// []string{"netsec", "service", "v1", "bgprouting"}

import (
	"context"
	"net/url"
	"strconv"

	"github.com/paloaltonetworks/sase-go/api"
	aUYGFNI "github.com/paloaltonetworks/sase-go/netsec/schema/bgp/routing"
)

// Client is the client for this namespace.
type Client struct {
	client api.Client
}

// NewClient returns a new client for this namespace.
func NewClient(client api.Client) *Client {
	return &Client{client: client}
}

// BgpRoutingPutInput takes some input.
// name:"BgpRoutingPut" nsfName:"BgpRoutingPut" param:0 query:0
// path: []string(nil)
// query: []string(nil)
type BgpRoutingPutInput struct {
	Config aUYGFNI.Config
}

// BgpRoutingPut performs a the given operation.
//
// Method: put
// URI: /sse/config/v1/bgp-routing
func (c *Client) BgpRoutingPut(ctx context.Context, input BgpRoutingPutInput) error {
	// Variables.
	var err error
	path := "/sse/config/v1/bgp-routing"

	// Execute the command.
	_, err = c.client.Do(ctx, "PUT", path, nil, input.Config, nil)

	// Done.
	return err
}

// ListInput takes some input.
// name:"List" nsfName:"List" param:0 query:3
// path: []string(nil)
// query: []string{"limit-optional", "offset-optional", "folder"}
type ListInput struct {
	Limit  *int64
	Offset *int64
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
	Data   []aUYGFNI.Config `json:"data,omitempty"`
	Limit  int64            `json:"limit,omitempty"`
	Offset int64            `json:"offset,omitempty"`
	Total  int64            `json:"total,omitempty"`
}

// List gets a list of objects back.
//
// Method: get
// URI: /sse/config/v1/bgp-routing
func (c *Client) List(ctx context.Context, input ListInput) (ListOutput, error) {
	// Variables.
	var err error
	var ans ListOutput
	path := "/sse/config/v1/bgp-routing"

	// Query parameter handling.
	uv := url.Values{}
	if input.Limit != nil {
		uv.Set("limit", strconv.Itoa(int(*input.Limit)))
	}
	if input.Offset != nil {
		uv.Set("offset", strconv.Itoa(int(*input.Offset)))
	}
	uv.Set("folder", input.Folder)

	// Execute the command.
	_, err = c.client.Do(ctx, "GET", path, uv, nil, &ans)

	// Done.
	return ans, err
}
