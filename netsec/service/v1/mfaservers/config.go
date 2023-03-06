package mfaservers

// This package is created automatically.
// Any manual changes will be overwritten when this is generated.
// []string{"netsec", "service", "v1", "mfaservers"}

import (
	"context"
	"net/url"
	"strings"

	"github.com/paloaltonetworks/sase-go/api"
	deRyMEf "github.com/paloaltonetworks/sase-go/netsec/schema/mfa/servers"
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
// name:"Create" nsfName:"Create" param:0 query:2
// path: []string{}
// query: []string{"position", "folder"}
type CreateInput struct {
	Position string
	Folder   string
	Config   deRyMEf.Config
}

// Create creates the specified object.
//
// Method: post
// URI: /sse/config/v1/mfa-servers
func (c *Client) Create(ctx context.Context, input CreateInput) (deRyMEf.Config, error) {
	// Variables.
	var err error
	var ans deRyMEf.Config
	path := "/sse/config/v1/mfa-servers"

	// Query parameter handling.
	uv := url.Values{}
	uv.Set("position", input.Position)
	uv.Set("folder", input.Folder)

	// Execute the command.
	_, err = c.client.Do(ctx, "POST", path, uv, input.Config, &ans)

	// Done.
	return ans, err
}

// DeleteInput takes some input.
// name:"Delete" nsfName:"Delete" param:1 query:0
// path: []string{"uuid-required"}
// query: []string{}
type DeleteInput struct {
	ObjectId string
}

// Delete removes the specified configuration.
//
// Method: delete
// URI: /sse/config/v1/mfa-servers/{id}
func (c *Client) Delete(ctx context.Context, input DeleteInput) (deRyMEf.Config, error) {
	// Variables.
	var err error
	var ans deRyMEf.Config
	path := "/sse/config/v1/mfa-servers/{id}"

	// Path param handling.
	path = strings.ReplaceAll(path, "{id}", input.ObjectId)

	// Execute the command.
	_, err = c.client.Do(ctx, "DELETE", path, nil, nil, &ans)

	// Done.
	return ans, err
}

// MfaServersGetInput takes some input.
// name:"MfaServersGet" nsfName:"MfaServersGet" param:0 query:2
// path: []string{}
// query: []string{"folder", "name-optional"}
type MfaServersGetInput struct {
	Folder string
	Name   *string
}

// MfaServersGet performs a the given operation.
//
// Method: get
// URI: /sse/config/v1/mfa-servers
func (c *Client) MfaServersGet(ctx context.Context, input MfaServersGetInput) (deRyMEf.Config, error) {
	// Variables.
	var err error
	var ans deRyMEf.Config
	path := "/sse/config/v1/mfa-servers"

	// Query parameter handling.
	uv := url.Values{}
	uv.Set("folder", input.Folder)
	if input.Name != nil {
		uv.Set("name", *input.Name)
	}

	// Execute the command.
	_, err = c.client.Do(ctx, "GET", path, uv, nil, &ans)

	// Done.
	return ans, err
}

// ReadInput takes some input.
// name:"Read" nsfName:"Read" param:1 query:0
// path: []string{"uuid-required"}
// query: []string{}
type ReadInput struct {
	ObjectId string
}

// Read returns the configuration of the specified object.
//
// Method: get
// URI: /sse/config/v1/mfa-servers/{id}
func (c *Client) Read(ctx context.Context, input ReadInput) (deRyMEf.Config, error) {
	// Variables.
	var err error
	var ans deRyMEf.Config
	path := "/sse/config/v1/mfa-servers/{id}"

	// Path param handling.
	path = strings.ReplaceAll(path, "{id}", input.ObjectId)

	// Execute the command.
	_, err = c.client.Do(ctx, "GET", path, nil, nil, &ans)

	// Done.
	return ans, err
}

// UpdateInput takes some input.
// name:"Update" nsfName:"Update" param:1 query:0
// path: []string{"uuid-required"}
// query: []string{}
type UpdateInput struct {
	ObjectId string
	Config   deRyMEf.Config
}

// Update modifies the configuration of the given object.
//
// Method: put
// URI: /sse/config/v1/mfa-servers/{id}
func (c *Client) Update(ctx context.Context, input UpdateInput) (deRyMEf.Config, error) {
	// Variables.
	var err error
	var ans deRyMEf.Config
	path := "/sse/config/v1/mfa-servers/{id}"

	// Path param handling.
	path = strings.ReplaceAll(path, "{id}", input.ObjectId)

	// Execute the command.
	_, err = c.client.Do(ctx, "PUT", path, nil, input.Config, &ans)

	// Done.
	return ans, err
}
