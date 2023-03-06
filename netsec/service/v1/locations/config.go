package locations

// This package is created automatically.
// Any manual changes will be overwritten when this is generated.
// []string{"netsec", "service", "v1", "locations"}

import (
    "context"

    "github.com/paloaltonetworks/sase-go/api"
    eItpily "github.com/paloaltonetworks/sase-go/netsec/schema/locations"
)

// Client is the client for this namespace.
type Client struct {
    client api.Client
}

// NewClient returns a new client for this namespace.
func NewClient(client api.Client) *Client {
    return &Client{client: client}
}

// LocationsGet performs a the given operation.
//
// Method: get
// URI: /sse/config/v1/locations
func (c *Client) LocationsGet(ctx context.Context)  (eItpily.Config, error) {
    // Variables.
    var err error
    var ans eItpily.Config
    path := "/sse/config/v1/locations"


    // Execute the command.
    _, err = c.client.Do(ctx, "GET", path, nil, nil, &ans)

    // Done.
    return ans, err
}
