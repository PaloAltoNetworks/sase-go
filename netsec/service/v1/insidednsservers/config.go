package insidednsservers

// This package is created automatically.
// Any manual changes will be overwritten when this is generated.
// []string{"netsec", "service", "v1", "insidednsservers"}

import (
    "context"
    "net/url"
    "strconv"
    "strings"

    "github.com/paloaltonetworks/sase-go/api"
    yYYdAbD "github.com/paloaltonetworks/sase-go/netsec/schema/inside/dns/servers"
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
    Config yYYdAbD.Config
}

// Create creates the specified object.
//
// Method: post
// URI: /sse/config/v1/internal-dns-servers
func (c *Client) Create(ctx context.Context, input CreateInput)  (yYYdAbD.Config, error) {
    // Variables.
    var err error
    var ans yYYdAbD.Config
    path := "/sse/config/v1/internal-dns-servers"

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
// URI: /sse/config/v1/internal-dns-servers/{id}
func (c *Client) Delete(ctx context.Context, input DeleteInput)  (yYYdAbD.Config, error) {
    // Variables.
    var err error
    var ans yYYdAbD.Config
    path := "/sse/config/v1/internal-dns-servers/{id}"

    // Path param handling.
    path = strings.ReplaceAll(path, "{id}", input.ObjectId)

    // Execute the command.
    _, err = c.client.Do(ctx, "DELETE", path, nil, nil, &ans)

    // Done.
    return ans, err
}

// InternalDnsServersGetInput takes some input.
// name:"InternalDnsServersGet" nsfName:"InternalDnsServersGet" param:0 query:4
// path: []string{}
// query: []string{"limit-optional", "offset-optional", "name-optional", "folder"}
type InternalDnsServersGetInput struct {
    Limit *int64
    Offset *int64
    Name *string
    Folder string
}

// InternalDnsServersGet performs a the given operation.
//
// Method: get
// URI: /sse/config/v1/internal-dns-servers
func (c *Client) InternalDnsServersGet(ctx context.Context, input InternalDnsServersGetInput)  (yYYdAbD.Config, error) {
    // Variables.
    var err error
    var ans yYYdAbD.Config
    path := "/sse/config/v1/internal-dns-servers"

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
// URI: /sse/config/v1/internal-dns-servers/{id}
func (c *Client) Read(ctx context.Context, input ReadInput)  (yYYdAbD.Config, error) {
    // Variables.
    var err error
    var ans yYYdAbD.Config
    path := "/sse/config/v1/internal-dns-servers/{id}"

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
    Config yYYdAbD.Config
}

// Update modifies the configuration of the given object.
//
// Method: put
// URI: /sse/config/v1/internal-dns-servers/{id}
func (c *Client) Update(ctx context.Context, input UpdateInput)  (yYYdAbD.Config, error) {
    // Variables.
    var err error
    var ans yYYdAbD.Config
    path := "/sse/config/v1/internal-dns-servers/{id}"

    // Path param handling.
    path = strings.ReplaceAll(path, "{id}", input.ObjectId)

    // Execute the command.
    _, err = c.client.Do(ctx, "PUT", path, nil, input.Config, &ans)

    // Done.
    return ans, err
}
