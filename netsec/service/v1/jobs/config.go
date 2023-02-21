package jobs

// This package is created automatically.
// Any manual changes will be overwritten when this is generated.
// []string{"netsec", "service", "v1", "jobs"}

import (
    "context"
    "math"
    "strings"

    "github.com/paloaltonetworks/sase-go/api"
    nTYZGXc "github.com/paloaltonetworks/sase-go/netsec/schema/jobs"
)

// Client is the client for this namespace.
type Client struct {
    client api.Client
}

// NewClient returns a new client for this namespace.
func NewClient(client api.Client) *Client {
    return &Client{client: client}
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
    Data []nTYZGXc.Config `json:"data,omitempty"`
    Limit int64 `json:"limit,omitempty"`
    Offset int64 `json:"offset,omitempty"`
    Total int64 `json:"total,omitempty"`
}

// List gets a list of objects back.
//
// Method: get
// URI: /sse/config/v1/jobs
func (c *Client) List(ctx context.Context)  (ListOutput, error) {
    // Variables.
    var err error
    var ans ListOutput
    path := "/sse/config/v1/jobs"


    // Optional: retrieve everything if limit is -1.
    if input.Limit != nil && *input.Limit == -1 {
        return c.listAll(ctx, input)
    }

    // Execute the command.
    _, err = c.client.Do(ctx, "GET", path, nil, nil, &ans)

    // Done.
    return ans, err
}

type listResponse struct {
    Output ListOutput
    Error error
}

func (c *Client) listAll(ctx context.Context, input ListInput) (ListOutput, error) {
    var ans ListOutput
    var err error
    var items map[string] nTYZGXc.Config
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
        items = make(map[string] nTYZGXc.Config)
        numRetrievers := int(math.Ceil(float64(total)/float64(api.MaxLimit)))
        responses := make(chan listResponse, numRetrievers)

        for i := 0; i < numRetrievers; i++ {
            ri := ListInput{
                Offset: api.Int(int64(i*api.MaxLimit)),
                Limit: api.Int(int64(api.MaxLimit)),
            }
            go func(){
                rout, rerr := c.List(ctx, ri)
                responses <- listResponse{
                    Output: rout,
                    Error: rerr,
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

    listing := make([]nTYZGXc.Config, 0, len(items))
    for key := range items {
        listing = append(listing, items[key])
    }

    ans = ListOutput{
        Data: listing,
        Total: int64(len(listing)),
    }

    return ans, nil
}

// ReadInput takes some input.
// name:"Read" nsfName:"Read" param:1 query:0
// path: []string{"jobid-required"}
// query: []string{}
type ReadInput struct {
    ObjectId string
}

// Read returns the configuration of the specified object.
//
// Method: get
// URI: /sse/config/v1/jobs/{id}
func (c *Client) Read(ctx context.Context, input ReadInput)  (nTYZGXc.Config, error) {
    // Variables.
    var err error
    var ans nTYZGXc.Config
    path := "/sse/config/v1/jobs/{id}"

    // Path param handling.
    path = strings.ReplaceAll(path, "{id}", input.ObjectId)

    // Execute the command.
    _, err = c.client.Do(ctx, "GET", path, nil, nil, &ans)

    // Done.
    return ans, err
}
