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
// path:
// query: Name | SpnNameList
type BandwidthAllocationsDeleteInput struct {
	Name        string
	SpnNameList string
}

// BandwidthAllocationsDelete performs a the given operation.
//
// Method: delete
// URI: /sse/config/v1/bandwidth-allocations
func (c *Client) BandwidthAllocationsDelete(ctx context.Context, input BandwidthAllocationsDeleteInput) error {
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
// path:
// query:
type BandwidthAllocationsPutInput struct {
	Config wYKgwLF.Config
}

// BandwidthAllocationsPut performs a the given operation.
//
// Method: put
// URI: /sse/config/v1/bandwidth-allocations
func (c *Client) BandwidthAllocationsPut(ctx context.Context, input BandwidthAllocationsPutInput) error {
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
// path:
// query:
type CreateInput struct {
	Config wYKgwLF.Config
}

// Create creates the specified object.
//
// Method: post
// URI: /sse/config/v1/bandwidth-allocations
func (c *Client) Create(ctx context.Context, input CreateInput) (wYKgwLF.Config, error) {
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
// path:
// query: Limit | Offset
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
	Data   []wYKgwLF.Config `json:"data,omitempty"`
	Limit  int64            `json:"limit,omitempty"`
	Offset int64            `json:"offset,omitempty"`
	Total  int64            `json:"total,omitempty"`
}

// List gets a list of objects back.
//
// Method: get
// URI: /sse/config/v1/bandwidth-allocations
func (c *Client) List(ctx context.Context, input ListInput) (ListOutput, error) {
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

	// Optional: retrieve everything if limit is -1.
	if input.Limit != nil && *input.Limit == -1 {
		return c.listAll(ctx, input)
	}

	// Execute the command.
	_, err = c.client.Do(ctx, "GET", path, uv, nil, &ans)

	// Done.
	return ans, err
}

type listResponse struct {
	Output ListOutput
	Error  error
}

func (c *Client) listAll(ctx context.Context, input ListInput) (ListOutput, error) {
	var ans ListOutput
	var err error
	var items map[string]wYKgwLF.Config
	everything := ListInput{
		Limit: api.Int(api.MaxLimit),
	}

	times := 0
	for {
		// Get the total number of things.
		ans, err = c.List(ctx, everything)
		if err != nil || len(ans.Data) == int(ans.Total) {
			return ans, err
		}

		total := int(ans.Total)
		items = make(map[string]wYKgwLF.Config)
		numRetrievers := int(math.Ceil(float64(total) / float64(api.MaxLimit)))
		responses := make(chan listResponse, numRetrievers)

		for i := 0; i < numRetrievers; i++ {
			ri := ListInput{
				Offset: api.Int(int64(i * api.MaxLimit)),
				Limit:  api.Int(int64(api.MaxLimit)),
			}
			go func() {
				rout, rerr := c.List(ctx, ri)
				responses <- listResponse{
					Output: rout,
					Error:  rerr,
				}
			}()
		}

		var totalChanged bool
		for i := 0; i < numRetrievers; i++ {
			resp := <-responses
			if resp.Error != nil {
				return resp.Output, resp.Error
			}
			if ans.Total != resp.Output.Total {
				totalChanged = true
				continue
			}
			for j := 0; j < len(resp.Output.Data); j++ {
				if _, ok := items[resp.Output.Data[j].ObjectId]; !ok {
					items[resp.Output.Data[j].ObjectId] = resp.Output.Data[j]
				}
			}
		}

		if !totalChanged && len(items) == total {
			break
		}

		times++
		if times >= 5 {
			return ListOutput{}, api.TooManyRetriesError
		}
	}

	listing := make([]wYKgwLF.Config, 0, len(items))
	for key := range items {
		listing = append(listing, items[key])
	}

	ans = ListOutput{
		Data:  listing,
		Total: int64(len(listing)),
	}

	return ans, nil
}
