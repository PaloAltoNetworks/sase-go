package quarantineddevices

// This package is created automatically.
// Any manual changes will be overwritten when this is generated.
// []string{"netsec", "service", "v1", "quarantineddevices"}

import (
	"context"
	"net/url"

	"github.com/paloaltonetworks/sase-go/api"
	qGeZssX "github.com/paloaltonetworks/sase-go/netsec/schema/objects/quarantined/devices"
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
// name:"Create" nsfName:"Create" param:0 query:0
// path:
// query:
type CreateInput struct {
	Config qGeZssX.Config
}

// Create creates the specified object.
//
// Method: post
// URI: /sse/config/v1/quarantined-devices
func (c *Client) Create(ctx context.Context, input CreateInput) (qGeZssX.Config, error) {
	// Variables.
	var err error
	var ans qGeZssX.Config
	path := "/sse/config/v1/quarantined-devices"

	// Execute the command.
	_, err = c.client.Do(ctx, "POST", path, nil, input.Config, &ans)

	// Done.
	return ans, err
}

// QuarantinedDevicesDeleteInput takes some input.
// name:"QuarantinedDevicesDelete" nsfName:"QuarantinedDevicesDelete" param:0 query:3
// path:
// query: ServiceType | ConnectionName | HostId
type QuarantinedDevicesDeleteInput struct {
	ServiceType    string
	ConnectionName string
	HostId         string
}

// QuarantinedDevicesDelete performs a the given operation.
//
// Method: delete
// URI: /sse/config/v1/quarantined-devices
func (c *Client) QuarantinedDevicesDelete(ctx context.Context, input QuarantinedDevicesDeleteInput) error {
	// Variables.
	var err error
	path := "/sse/config/v1/quarantined-devices"

	// Query parameter handling.
	uv := url.Values{}
	uv.Set("service_type", input.ServiceType)
	uv.Set("connection_name", input.ConnectionName)
	uv.Set("host_id", input.HostId)

	// Execute the command.
	_, err = c.client.Do(ctx, "DELETE", path, uv, nil, nil)

	// Done.
	return err
}

// QuarantinedDevicesGetInput takes some input.
// name:"QuarantinedDevicesGet" nsfName:"QuarantinedDevicesGet" param:0 query:4
// path:
// query: ServiceType | ConnectionName | HostId | SerialNumber
type QuarantinedDevicesGetInput struct {
	ServiceType    string
	ConnectionName string
	HostId         *string
	SerialNumber   *string
}

// QuarantinedDevicesGet performs a the given operation.
//
// Method: get
// URI: /sse/config/v1/quarantined-devices
func (c *Client) QuarantinedDevicesGet(ctx context.Context, input QuarantinedDevicesGetInput) (qGeZssX.Config, error) {
	// Variables.
	var err error
	var ans qGeZssX.Config
	path := "/sse/config/v1/quarantined-devices"

	// Query parameter handling.
	uv := url.Values{}
	uv.Set("service_type", input.ServiceType)
	uv.Set("connection_name", input.ConnectionName)
	if input.HostId != nil {
		uv.Set("host_id", *input.HostId)
	}
	if input.SerialNumber != nil {
		uv.Set("serial_number", *input.SerialNumber)
	}

	// Execute the command.
	_, err = c.client.Do(ctx, "GET", path, uv, nil, &ans)

	// Done.
	return ans, err
}
