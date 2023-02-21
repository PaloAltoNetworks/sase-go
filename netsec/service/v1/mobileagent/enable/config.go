package enable

// This package is created automatically.
// Any manual changes will be overwritten when this is generated.
// []string{"netsec", "service", "v1", "mobileagent", "enable"}

import (
    "context"

    "github.com/paloaltonetworks/sase-go/api"
)

// Client is the client for this namespace.
type Client struct {
    client api.Client
}

// NewClient returns a new client for this namespace.
func NewClient(client api.Client) *Client {
    return &Client{client: client}
}

/*
EnableGetOutput ""

Parent chains:
* "_output"

Args:

Param Data (DataObject): The Data param. DataObject instance.
*/
type EnableGetOutput struct {
    Data *DataObject `json:"data,omitempty"`
}

/*
DataObject ""

Parent chains:
* "_output data"

Args:

Param Enabled (bool): The Enabled param.
*/
type DataObject struct {
    Enabled bool `json:"enabled,omitempty"`
}

// EnableGet performs a the given operation.
//
// Method: get
// URI: /sse/config/v1/mobile-agent/enable
func (c *Client) EnableGet(ctx context.Context)  (EnableGetOutput, error) {
    // Variables.
    var err error
    var ans EnableGetOutput
    path := "/sse/config/v1/mobile-agent/enable"


    // Execute the command.
    _, err = c.client.Do(ctx, "GET", path, nil, nil, &ans)

    // Done.
    return ans, err
}

// EnablePost performs a the given operation.
//
// Method: post
// URI: /sse/config/v1/mobile-agent/enable
func (c *Client) EnablePost(ctx context.Context)  (error) {
    // Variables.
    var err error
    path := "/sse/config/v1/mobile-agent/enable"


    // Execute the command.
    _, err = c.client.Do(ctx, "POST", path, nil, nil, nil)

    // Done.
    return err
}
