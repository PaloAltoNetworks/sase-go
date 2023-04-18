package candidate

// This package is created automatically.
// Any manual changes will be overwritten when this is generated.
// []string{"netsec", "service", "v1", "configversions", "candidate"}

import (
	"context"

	"github.com/paloaltonetworks/sase-go/api"
)

// Client is the client for this namespace.
type Client struct {
	client api.Client
}

// NewClient returns a new client for this namespace.
func NewClient(client api.Client) *Client {
	return &Client{client: client}
}

// CandidateDelete performs a the given operation.
//
// Method: delete
// URI: /sse/config/v1/config-versions/candidate
func (c *Client) CandidateDelete(ctx context.Context) error {
	// Variables.
	var err error
	path := "/sse/config/v1/config-versions/candidate"

	// Execute the command.
	_, err = c.client.Do(ctx, "DELETE", path, nil, nil, nil)

	// Done.
	return err
}

// CandidatePushPostInput takes some input.
// name:"CandidatePushPost" nsfName:"CandidatePushPost" param:0 query:0
// path:
// query:
type CandidatePushPostInput struct {
	Config CandidatePushPostInputConfig
}

/*
CandidatePushPostInputConfig ""

Parent chains:
* "_input"

Args:

Param Description (string): The Description param.

Param Folders (required, list of strings): The Folders param.
*/
type CandidatePushPostInputConfig struct {
	Description string   `json:"description,omitempty"`
	Folders     []string `json:"folders"`
}

// CandidatePushPost performs a the given operation.
//
// Method: post
// URI: /sse/config/v1/config-versions/candidate:push
func (c *Client) CandidatePushPost(ctx context.Context, input CandidatePushPostInput) error {
	// Variables.
	var err error
	path := "/sse/config/v1/config-versions/candidate:push"

	// Execute the command.
	_, err = c.client.Do(ctx, "POST", path, nil, input.Config, nil)

	// Done.
	return err
}
