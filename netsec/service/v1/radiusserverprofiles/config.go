package radiusserverprofiles

// This package is created automatically.
// Any manual changes will be overwritten when this is generated.
// []string{"netsec", "service", "v1", "radiusserverprofiles"}

import (
    "context"
    "math"
    "net/url"
    "strconv"
    "strings"

    "github.com/paloaltonetworks/sase-go/api"
    ieMayhq "github.com/paloaltonetworks/sase-go/netsec/schema/radius/server/profiles"
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
// name:"Create" nsfName:"Create" param:0 query:1
// path: []string{}
// query: []string{"folder"}
type CreateInput struct {
    Folder string
    Config ieMayhq.Config
}

// Create creates the specified object.
//
// Method: post
// URI: /sse/config/v1/radius-server-profiles
func (c *Client) Create(ctx context.Context, input CreateInput)  (ieMayhq.Config, error) {
    // Variables.
    var err error
    var ans ieMayhq.Config
    path := "/sse/config/v1/radius-server-profiles"

    // Query parameter handling.
    uv := url.Values{}
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
// URI: /sse/config/v1/radius-server-profiles/{id}
func (c *Client) Delete(ctx context.Context, input DeleteInput)  (ieMayhq.Config, error) {
    // Variables.
    var err error
    var ans ieMayhq.Config
    path := "/sse/config/v1/radius-server-profiles/{id}"

    // Path param handling.
    path = strings.ReplaceAll(path, "{id}", input.ObjectId)

    // Execute the command.
    _, err = c.client.Do(ctx, "DELETE", path, nil, nil, &ans)

    // Done.
    return ans, err
}

// ListInput takes some input.
// name:"List" nsfName:"List" param:0 query:4
// path: []string{}
// query: []string{"limit-optional", "offset-optional", "folder", "name-optional"}
type ListInput struct {
    Limit *int64
    Offset *int64
    Folder string
    Name *string
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
    Data []ieMayhq.Config `json:"data,omitempty"`
    Limit int64 `json:"limit,omitempty"`
    Offset int64 `json:"offset,omitempty"`
    Total int64 `json:"total,omitempty"`
}

// List gets a list of objects back.
//
// Method: get
// URI: /sse/config/v1/radius-server-profiles
func (c *Client) List(ctx context.Context, input ListInput)  (ListOutput, error) {
    // Variables.
    var err error
    var ans ListOutput
    path := "/sse/config/v1/radius-server-profiles"

    // Query parameter handling.
    uv := url.Values{}
    if input.Limit != nil {
        uv.Set("limit", strconv.Itoa(int(*input.Limit)))
    }
    if input.Offset != nil {
        uv.Set("offset", strconv.Itoa(int(*input.Offset)))
    }
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
// URI: /sse/config/v1/radius-server-profiles/{id}
func (c *Client) Read(ctx context.Context, input ReadInput)  (ieMayhq.Config, error) {
    // Variables.
    var err error
    var ans ieMayhq.Config
    path := "/sse/config/v1/radius-server-profiles/{id}"

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
    Config ieMayhq.Config
}

// Update modifies the configuration of the given object.
//
// Method: put
// URI: /sse/config/v1/radius-server-profiles/{id}
func (c *Client) Update(ctx context.Context, input UpdateInput)  (ieMayhq.Config, error) {
    // Variables.
    var err error
    var ans ieMayhq.Config
    path := "/sse/config/v1/radius-server-profiles/{id}"

    // Path param handling.
    path = strings.ReplaceAll(path, "{id}", input.ObjectId)

    // Execute the command.
    _, err = c.client.Do(ctx, "PUT", path, nil, input.Config, &ans)

    // Done.
    return ans, err
}
