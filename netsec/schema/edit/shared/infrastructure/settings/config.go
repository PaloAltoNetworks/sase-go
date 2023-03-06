package settings

/*
Config ""

Parent chains:
* "Config"

Args:

Param EgressIpNotificationUrl (string): The EgressIpNotificationUrl param.

Param InfraBgpAs (string): The InfraBgpAs param.

Param InfrastructureSubnet (string): The InfrastructureSubnet param.

Param InfrastructureSubnetIpv6 (string): The InfrastructureSubnetIpv6 param.
*/
type Config struct {
	EgressIpNotificationUrl  string `json:"egress_ip_notification_url,omitempty"`
	InfraBgpAs               string `json:"infra_bgp_as,omitempty"`
	InfrastructureSubnet     string `json:"infrastructure_subnet,omitempty"`
	InfrastructureSubnetIpv6 string `json:"infrastructure_subnet_ipv6,omitempty"`
}
