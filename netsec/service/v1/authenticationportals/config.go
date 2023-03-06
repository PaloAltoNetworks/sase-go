package authenticationportals

// This package is created automatically.
// Any manual changes will be overwritten when this is generated.
// []string{"netsec", "service", "v1", "authenticationportals"}

import (
    "context"
    "math"
    "net/url"
    "strconv"
    "strings"

    "github.com/paloaltonetworks/sase-go/api"
    xUpBHLw "github.com/paloaltonetworks/sase-go/netsec/schema/authentication/portals"
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
    Config xUpBHLw.Config
}

// Create creates the specified object.
//
// Method: post
// URI: /sse/config/v1/authentication-portals
func (c *Client) Create(ctx context.Context, input CreateInput)  (xUpBHLw.Config, error) {
    // Variables.
    var err error
    var ans xUpBHLw.Config
    path := "/sse/config/v1/authentication-portals"

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
// URI: /sse/config/v1/authentication-portals/{id}
func (c *Client) Delete(ctx context.Context, input DeleteInput)  (xUpBHLw.Config, error) {
    // Variables.
    var err error
    var ans xUpBHLw.Config
    path := "/sse/config/v1/authentication-portals/{id}"

    // Path param handling.
    path = strings.ReplaceAll(path, "{id}", input.ObjectId)

    // Execute the command.
    _, err = c.client.Do(ctx, "DELETE", path, nil, nil, &ans)

    // Done.
    return ans, err
}

// ListInput takes some input.
// name:"List" nsfName:"List" param:0 query:3
// path: []string{}
// query: []string{"limit-optional", "offset-optional", "folder"}
type ListInput struct {
    Limit *int64
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
    Data []xUpBHLw.Config `json:"data,omitempty"`
    Limit int64 `json:"limit,omitempty"`
    Offset int64 `json:"offset,omitempty"`
    Total int64 `json:"total,omitempty"`
}

// List gets a list of objects back.
//
// Method: get
// URI: /sse/config/v1/authentication-portals
func (c *Client) List(ctx context.Context, input ListInput)  (ListOutput, error) {
    // Variables.
    var err error
    var ans ListOutput
    path := "/sse/config/v1/authentication-portals"

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

// UpdateInput takes some input.
// name:"Update" nsfName:"Update" param:1 query:0
// path: []string{"uuid-required"}
// query: []string{}
type UpdateInput struct {
    ObjectId string
    Config xUpBHLw.Config
}

// Update modifies the configuration of the given object.
//
// Method: put
// URI: /sse/config/v1/authentication-portals/{id}
func (c *Client) Update(ctx context.Context, input UpdateInput)  (xUpBHLw.Config, error) {
    // Variables.
    var err error
    var ans xUpBHLw.Config
    path := "/sse/config/v1/authentication-portals/{id}"

    // Path param handling.
    path = strings.ReplaceAll(path, "{id}", input.ObjectId)

    // Execute the command.
    _, err = c.client.Do(ctx, "PUT", path, nil, input.Config, &ans)

    // Done.
    return ans, err
}
