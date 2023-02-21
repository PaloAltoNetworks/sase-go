package licensetypes

// This package is created automatically.
// Any manual changes will be overwritten when this is generated.
// []string{"netsec", "service", "v1", "licensetypes"}

import (
    "context"

    "github.com/paloaltonetworks/sase-go/api"
    prfBhvs "github.com/paloaltonetworks/sase-go/netsec/schema/license/types"
)

// Client is the client for this namespace.
type Client struct {
    client api.Client
}

// NewClient returns a new client for this namespace.
func NewClient(client api.Client) *Client {
    return &Client{client: client}
}

// LicenseTypesGet performs a the given operation.
//
// Method: get
// URI: /sse/config/v1/license-types
func (c *Client) LicenseTypesGet(ctx context.Context)  (prfBhvs.Config, error) {
    // Variables.
    var err error
    var ans prfBhvs.Config
    path := "/sse/config/v1/license-types"


    // Execute the command.
    _, err = c.client.Do(ctx, "GET", path, nil, nil, &ans)

    // Done.
    return ans, err
}
