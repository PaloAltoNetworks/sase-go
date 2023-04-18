package globalsettings

// This package is created automatically.
// Any manual changes will be overwritten when this is generated.
// []string{"netsec", "service", "v1", "mobileagent", "globalsettings"}

import (
	"context"

	"github.com/paloaltonetworks/sase-go/api"
	wajgxii "github.com/paloaltonetworks/sase-go/netsec/schema/mobile/agent/global/settings"
)

// Client is the client for this namespace.
type Client struct {
	client api.Client
}

// NewClient returns a new client for this namespace.
func NewClient(client api.Client) *Client {
	return &Client{client: client}
}

// GlobalSettingsGet performs a the given operation.
//
// Method: get
// URI: /sse/config/v1/mobile-agent/global-settings
func (c *Client) GlobalSettingsGet(ctx context.Context) (wajgxii.Config, error) {
	// Variables.
	var err error
	var ans wajgxii.Config
	path := "/sse/config/v1/mobile-agent/global-settings"

	// Execute the command.
	_, err = c.client.Do(ctx, "GET", path, nil, nil, &ans)

	// Done.
	return ans, err
}

// GlobalSettingsPutInput takes some input.
// name:"GlobalSettingsPut" nsfName:"GlobalSettingsPut" param:0 query:0
// path:
// query:
type GlobalSettingsPutInput struct {
	Config wajgxii.Config
}

// GlobalSettingsPut performs a the given operation.
//
// Method: put
// URI: /sse/config/v1/mobile-agent/global-settings
func (c *Client) GlobalSettingsPut(ctx context.Context, input GlobalSettingsPutInput) error {
	// Variables.
	var err error
	path := "/sse/config/v1/mobile-agent/global-settings"

	// Execute the command.
	_, err = c.client.Do(ctx, "PUT", path, nil, input.Config, nil)

	// Done.
	return err
}
