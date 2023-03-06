package autotagactions

// This package is created automatically.
// Any manual changes will be overwritten when this is generated.
// []string{"netsec", "service", "v1", "autotagactions"}

import (
    "context"
    "math"
    "net/url"
    "strconv"

    "github.com/paloaltonetworks/sase-go/api"
    kJVbXva "github.com/paloaltonetworks/sase-go/netsec/schema/auto/tag/actions"
)

// Client is the client for this namespace.
type Client struct {
    client api.Client
}

// NewClient returns a new client for this namespace.
func NewClient(client api.Client) *Client {
    return &Client{client: client}
}

// AutoTagActionsDeleteInput takes some input.
// name:"AutoTagActionsDelete" nsfName:"AutoTagActionsDelete" param:0 query:1
// path: []string{}
// query: []string{"name"}
type AutoTagActionsDeleteInput struct {
    Name string
}

// AutoTagActionsDelete performs a the given operation.
//
// Method: delete
// URI: /sse/config/v1/auto-tag-actions
func (c *Client) AutoTagActionsDelete(ctx context.Context, input AutoTagActionsDeleteInput)  (error) {
    // Variables.
    var err error
    path := "/sse/config/v1/auto-tag-actions"

    // Query parameter handling.
    uv := url.Values{}
    uv.Set("name", input.Name)

    // Execute the command.
    _, err = c.client.Do(ctx, "DELETE", path, uv, nil, nil)

    // Done.
    return err
}

// AutoTagActionsPutInput takes some input.
// name:"AutoTagActionsPut" nsfName:"AutoTagActionsPut" param:0 query:1
// path: []string{}
// query: []string{"folder-shared"}
type AutoTagActionsPutInput struct {
    Folder string
    Config kJVbXva.Config
}

// AutoTagActionsPut performs a the given operation.
//
// Method: put
// URI: /sse/config/v1/auto-tag-actions
func (c *Client) AutoTagActionsPut(ctx context.Context, input AutoTagActionsPutInput)  (error) {
    // Variables.
    var err error
    path := "/sse/config/v1/auto-tag-actions"

    // Query parameter handling.
    uv := url.Values{}
    uv.Set("folder", input.Folder)

    // Execute the command.
    _, err = c.client.Do(ctx, "PUT", path, uv, input.Config, nil)

    // Done.
    return err
}

// CreateInput takes some input.
// name:"Create" nsfName:"Create" param:0 query:1
// path: []string{}
// query: []string{"folder-shared"}
type CreateInput struct {
    Folder string
    Config kJVbXva.Config
}

// Create creates the specified object.
//
// Method: post
// URI: /sse/config/v1/auto-tag-actions
func (c *Client) Create(ctx context.Context, input CreateInput)  (kJVbXva.Config, error) {
    // Variables.
    var err error
    var ans kJVbXva.Config
    path := "/sse/config/v1/auto-tag-actions"

    // Query parameter handling.
    uv := url.Values{}
    uv.Set("folder", input.Folder)

    // Execute the command.
    _, err = c.client.Do(ctx, "POST", path, uv, input.Config, &ans)

    // Done.
    return ans, err
}

// ListInput takes some input.
// name:"List" nsfName:"List" param:0 query:4
// path: []string{}
// query: []string{"limit-optional", "offset-optional", "name-optional", "folder-shared"}
type ListInput struct {
    Limit *int64
    Offset *int64
    Name *string
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
    Data []kJVbXva.Config `json:"data,omitempty"`
    Limit int64 `json:"limit,omitempty"`
    Offset int64 `json:"offset,omitempty"`
    Total int64 `json:"total,omitempty"`
}

// List gets a list of objects back.
//
// Method: get
// URI: /sse/config/v1/auto-tag-actions
func (c *Client) List(ctx context.Context, input ListInput)  (ListOutput, error) {
    // Variables.
    var err error
    var ans ListOutput
    path := "/sse/config/v1/auto-tag-actions"

    // Query parameter handling.
    uv := url.Values{}
    if input.Limit != nil {
        uv.Set("limit", strconv.Itoa(int(*input.Limit)))
    }
    if input.Offset != nil {
        uv.Set("offset", strconv.Itoa(int(*input.Offset)))
    }
    if input.Name != nil {
        uv.Set("name", *input.Name)
    }
    uv.Set("folder", input.Folder)

    // Execute the command.
    _, err = c.client.Do(ctx, "GET", path, uv, nil, &ans)

    // Done.
    return ans, err
}
