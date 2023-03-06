package configversions

// This package is created automatically.
// Any manual changes will be overwritten when this is generated.
// []string{"netsec", "service", "v1", "configversions"}

import (
	"context"
	"net/url"
	"strconv"
	"strings"

	"github.com/paloaltonetworks/sase-go/api"
	bUPXiuP "github.com/paloaltonetworks/sase-go/netsec/schema/candidate/config/versions"
	gjFlAYT "github.com/paloaltonetworks/sase-go/netsec/schema/load/config"
)

// Client is the client for this namespace.
type Client struct {
	client api.Client
}

// NewClient returns a new client for this namespace.
func NewClient(client api.Client) *Client {
	return &Client{client: client}
}

// ConfigVersionsGetInput takes some input.
// name:"ConfigVersionsGet" nsfName:"ConfigVersionsGet" param:0 query:2
// path: []string{}
// query: []string{"limit-optional", "offset-optional"}
type ConfigVersionsGetInput struct {
	Limit  *int64
	Offset *int64
}

// ConfigVersionsGet performs a the given operation.
//
// Method: get
// URI: /sse/config/v1/config-versions
func (c *Client) ConfigVersionsGet(ctx context.Context, input ConfigVersionsGetInput) (bUPXiuP.Config, error) {
	// Variables.
	var err error
	var ans bUPXiuP.Config
	path := "/sse/config/v1/config-versions"

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

// CreateInput takes some input.
// name:"Create" nsfName:"Create" param:0 query:0
// path: []string(nil)
// query: []string(nil)
type CreateInput struct {
	Config gjFlAYT.Config
}

// Create creates the specified object.
//
// Method: post
// URI: /sse/config/v1/config-versions:load
func (c *Client) Create(ctx context.Context, input CreateInput) (gjFlAYT.Config, error) {
	// Variables.
	var err error
	var ans gjFlAYT.Config
	path := "/sse/config/v1/config-versions:load"

	// Execute the command.
	_, err = c.client.Do(ctx, "POST", path, nil, input.Config, &ans)

	// Done.
	return ans, err
}

// ReadInput takes some input.
// name:"Read" nsfName:"Read" param:1 query:0
// path: []string{"version-required"}
// query: []string{}
type ReadInput struct {
	Version string
}

// Read returns the configuration of the specified object.
//
// Method: get
// URI: /sse/config/v1/config-versions/{version}
func (c *Client) Read(ctx context.Context, input ReadInput) (bUPXiuP.Config, error) {
	// Variables.
	var err error
	var ans bUPXiuP.Config
	path := "/sse/config/v1/config-versions/{version}"

	// Path param handling.
	path = strings.ReplaceAll(path, "{version}", input.Version)

	// Execute the command.
	_, err = c.client.Do(ctx, "GET", path, nil, nil, &ans)

	// Done.
	return ans, err
}
