package decryptionexclusions

// This package is created automatically.
// Any manual changes will be overwritten when this is generated.
// []string{"netsec", "service", "v1", "decryptionexclusions"}

import (
    "context"
    "net/url"
    "strings"

    "github.com/paloaltonetworks/sase-go/api"
    esegVXo "github.com/paloaltonetworks/sase-go/netsec/schema/decryption/exclusions"
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
    Config esegVXo.Config
}

// Create creates the specified object.
//
// Method: post
// URI: /sse/config/v1/decryption-exclusions
func (c *Client) Create(ctx context.Context, input CreateInput)  (esegVXo.Config, error) {
    // Variables.
    var err error
    var ans esegVXo.Config
    path := "/sse/config/v1/decryption-exclusions"

    // Query parameter handling.
    uv := url.Values{}
    uv.Set("folder", input.Folder)

    // Execute the command.
    _, err = c.client.Do(ctx, "POST", path, uv, input.Config, &ans)

    // Done.
    return ans, err
}

// DecryptionExclusionsGetInput takes some input.
// name:"DecryptionExclusionsGet" nsfName:"DecryptionExclusionsGet" param:0 query:2
// path: []string{}
// query: []string{"name-optional", "folder"}
type DecryptionExclusionsGetInput struct {
    Name *string
    Folder string
}

// DecryptionExclusionsGet performs a the given operation.
//
// Method: get
// URI: /sse/config/v1/decryption-exclusions
func (c *Client) DecryptionExclusionsGet(ctx context.Context, input DecryptionExclusionsGetInput)  (esegVXo.Config, error) {
    // Variables.
    var err error
    var ans esegVXo.Config
    path := "/sse/config/v1/decryption-exclusions"

    // Query parameter handling.
    uv := url.Values{}
    if input.Name != nil {
        uv.Set("name", *input.Name)
    }
    uv.Set("folder", input.Folder)

    // Execute the command.
    _, err = c.client.Do(ctx, "GET", path, uv, nil, &ans)

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
// URI: /sse/config/v1/decryption-exclusions/{id}
func (c *Client) Delete(ctx context.Context, input DeleteInput)  (esegVXo.Config, error) {
    // Variables.
    var err error
    var ans esegVXo.Config
    path := "/sse/config/v1/decryption-exclusions/{id}"

    // Path param handling.
    path = strings.ReplaceAll(path, "{id}", input.ObjectId)

    // Execute the command.
    _, err = c.client.Do(ctx, "DELETE", path, nil, nil, &ans)

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
// URI: /sse/config/v1/decryption-exclusions/{id}
func (c *Client) Read(ctx context.Context, input ReadInput)  (esegVXo.Config, error) {
    // Variables.
    var err error
    var ans esegVXo.Config
    path := "/sse/config/v1/decryption-exclusions/{id}"

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
    Config esegVXo.Config
}

// Update modifies the configuration of the given object.
//
// Method: put
// URI: /sse/config/v1/decryption-exclusions/{id}
func (c *Client) Update(ctx context.Context, input UpdateInput)  (esegVXo.Config, error) {
    // Variables.
    var err error
    var ans esegVXo.Config
    path := "/sse/config/v1/decryption-exclusions/{id}"

    // Path param handling.
    path = strings.ReplaceAll(path, "{id}", input.ObjectId)

    // Execute the command.
    _, err = c.client.Do(ctx, "PUT", path, nil, input.Config, &ans)

    // Done.
    return ans, err
}
