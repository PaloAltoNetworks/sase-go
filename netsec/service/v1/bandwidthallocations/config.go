package bandwidthallocations

// This package is created automatically.
// Any manual changes will be overwritten when this is generated.
// []string{"netsec", "service", "v1", "bandwidthallocations"}

import (
    "context"
    "math"
    "net/url"
    "strconv"

    "github.com/paloaltonetworks/sase-go/api"
    wYKgwLF "github.com/paloaltonetworks/sase-go/netsec/schema/bandwidth/allocations"
)

// Client is the client for this namespace.
type Client struct {
    client api.Client
}

// NewClient returns a new client for this namespace.
func NewClient(client api.Client) *Client {
    return &Client{client: client}
}

// BandwidthAllocationsDeleteInput takes some input.
// name:"BandwidthAllocationsDelete" nsfName:"BandwidthAllocationsDelete" param:0 query:2
// path: []string{}
// query: []string{"aggregated-bandwidth-region-name-required", "aggregated-bandwidth-spn-name-list-required"}
type BandwidthAllocationsDeleteInput struct {
    Name string
    SpnNameList string
}

// BandwidthAllocationsDelete performs a the given operation.
//
// Method: delete
// URI: /sse/config/v1/bandwidth-allocations
func (c *Client) BandwidthAllocationsDelete(ctx context.Context, input BandwidthAllocationsDeleteInput)  (error) {
    // Variables.
    var err error
    path := "/sse/config/v1/bandwidth-allocations"

    // Query parameter handling.
    uv := url.Values{}
    uv.Set("name", input.Name)
    uv.Set("spn_name_list", input.SpnNameList)

    // Execute the command.
    _, err = c.client.Do(ctx, "DELETE", path, uv, nil, nil)

    // Done.
    return err
}

// BandwidthAllocationsPutInput takes some input.
// name:"BandwidthAllocationsPut" nsfName:"BandwidthAllocationsPut" param:0 query:0
// path: []string(nil)
// query: []string(nil)
type BandwidthAllocationsPutInput struct {
    Config wYKgwLF.Config
}

// BandwidthAllocationsPut performs a the given operation.
//
// Method: put
// URI: /sse/config/v1/bandwidth-allocations
func (c *Client) BandwidthAllocationsPut(ctx context.Context, input BandwidthAllocationsPutInput)  (error) {
    // Variables.
    var err error
    path := "/sse/config/v1/bandwidth-allocations"


    // Execute the command.
    _, err = c.client.Do(ctx, "PUT", path, nil, input.Config, nil)

    // Done.
    return err
}

// CreateInput takes some input.
// name:"Create" nsfName:"Create" param:0 query:0
// path: []string(nil)
// query: []string(nil)
type CreateInput struct {
    Config wYKgwLF.Config
}

// Create creates the specified object.
//
// Method: post
// URI: /sse/config/v1/bandwidth-allocations
func (c *Client) Create(ctx context.Context, input CreateInput)  (wYKgwLF.Config, error) {
    // Variables.
    var err error
    var ans wYKgwLF.Config
    path := "/sse/config/v1/bandwidth-allocations"


    // Execute the command.
    _, err = c.client.Do(ctx, "POST", path, nil, input.Config, &ans)

    // Done.
    return ans, err
}

// ListInput takes some input.
// name:"List" nsfName:"List" param:0 query:2
// path: []string{}
// query: []string{"limit-optional", "offset-optional"}
type ListInput struct {
    Limit *int64
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
    Data []wYKgwLF.Config `json:"data,omitempty"`
    Limit int64 `json:"limit,omitempty"`
    Offset int64 `json:"offset,omitempty"`
    Total int64 `json:"total,omitempty"`
}

// List gets a list of objects back.
//
// Method: get
// URI: /sse/config/v1/bandwidth-allocations
func (c *Client) List(ctx context.Context, input ListInput)  (ListOutput, error) {
    // Variables.
    var err error
    var ans ListOutput
    path := "/sse/config/v1/bandwidth-allocations"

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
