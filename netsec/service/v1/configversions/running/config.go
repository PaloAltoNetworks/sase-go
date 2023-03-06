package running

// This package is created automatically.
// Any manual changes will be overwritten when this is generated.
// []string{"netsec", "service", "v1", "configversions", "running"}

import (
    "context"

    "github.com/paloaltonetworks/sase-go/api"
    epjnjbE "github.com/paloaltonetworks/sase-go/netsec/schema/running/config/versions"
)

// Client is the client for this namespace.
type Client struct {
    client api.Client
}

// NewClient returns a new client for this namespace.
func NewClient(client api.Client) *Client {
    return &Client{client: client}
}

// RunningGet performs a the given operation.
//
// Method: get
// URI: /sse/config/v1/config-versions/running
func (c *Client) RunningGet(ctx context.Context)  (epjnjbE.Config, error) {
    // Variables.
    var err error
    var ans epjnjbE.Config
    path := "/sse/config/v1/config-versions/running"


    // Execute the command.
    _, err = c.client.Do(ctx, "GET", path, nil, nil, &ans)

    // Done.
    return ans, err
}
