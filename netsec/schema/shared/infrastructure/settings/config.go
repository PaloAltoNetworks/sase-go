package settings

/*
Config ""

Parent chains:
* "Config"

Args:

Param ApiKey (string): The ApiKey param.

Param CaptivePortalRedirectIpAddress (string): The CaptivePortalRedirectIpAddress param.

Param EgressIpNotificationUrl (string): The EgressIpNotificationUrl param.

Param InfraBgpAs (string): The InfraBgpAs param.

Param InfrastructureSubnet (string): The InfrastructureSubnet param.

Param InfrastructureSubnetIpv6 (string): The InfrastructureSubnetIpv6 param.

Param Ipv6 (bool): The Ipv6 param.

Param LoopbackIps (list of strings): The LoopbackIps param.

Param TunnelMonitorIpAddress (string): The TunnelMonitorIpAddress param.
*/
type Config struct {
	ApiKey                         string   `json:"api_key,omitempty"`
	CaptivePortalRedirectIpAddress string   `json:"captive_portal_redirect_ip_address,omitempty"`
	EgressIpNotificationUrl        string   `json:"egress_ip_notification_url,omitempty"`
	InfraBgpAs                     string   `json:"infra_bgp_as,omitempty"`
	InfrastructureSubnet           string   `json:"infrastructure_subnet,omitempty"`
	InfrastructureSubnetIpv6       string   `json:"infrastructure_subnet_ipv6,omitempty"`
	Ipv6                           bool     `json:"ipv6,omitempty"`
	LoopbackIps                    []string `json:"loopback_ips,omitempty"`
	TunnelMonitorIpAddress         string   `json:"tunnel_monitor_ip_address,omitempty"`
}
