package jobs

// This package is created automatically.
// Any manual changes will be overwritten when this is generated.
// []string{"netsec", "service", "v1", "jobs"}

import (
	"context"
	"strings"

	"github.com/paloaltonetworks/sase-go/api"
	sPPuDTU "github.com/paloaltonetworks/sase-go/netsec/schema/jobs"
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
	Data   []sPPuDTU.Config `json:"data,omitempty"`
	Limit  int64            `json:"limit,omitempty"`
	Offset int64            `json:"offset,omitempty"`
	Total  int64            `json:"total,omitempty"`
}

// List gets a list of objects back.
//
// Method: get
// URI: /sse/config/v1/jobs
func (c *Client) List(ctx context.Context) (ListOutput, error) {
	// Variables.
	var err error
	var ans ListOutput
	path := "/sse/config/v1/jobs"

	// Execute the command.
	_, err = c.client.Do(ctx, "GET", path, nil, nil, &ans)

	// Done.
	return ans, err
}

// ReadInput takes some input.
// name:"Read" nsfName:"Read" param:1 query:0
// path: JobId
// query:
type ReadInput struct {
	JobId string
}

// Read returns the configuration of the specified object.
//
// Method: get
// URI: /sse/config/v1/jobs/{id}
func (c *Client) Read(ctx context.Context, input ReadInput) (sPPuDTU.Config, error) {
	// Variables.
	var err error
	var ans sPPuDTU.Config
	path := "/sse/config/v1/jobs/{id}"

	// Path param handling.
	path = strings.ReplaceAll(path, "{id}", input.JobId)

	// Execute the command.
	_, err = c.client.Do(ctx, "GET", path, nil, nil, &ans)

	// Done.
	return ans, err
}
