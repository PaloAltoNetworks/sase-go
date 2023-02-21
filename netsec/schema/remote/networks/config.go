package networks

/*
Config ""

Parent chains:
* "Config"

Args:

Param EcmpLoadBalancing (string): The EcmpLoadBalancing param. String values: []string{"enable", "disable"} Default: "disable".

Param EcmpTunnels (list of EcmpTunnelsObject objects): The EcmpTunnels param.

Param ObjectId (string, read-only): The ObjectId param. Example: "abcd-1234".

Param IpsecTunnel (string): The IpsecTunnel param.

Param LicenseType (required, string): The LicenseType param. String length must exceed 1 characters. Default: "FWAAS-AGGREGATE".

Param Name (required, string): The Name param. String length must not exceed 63 characters.

Param Protocol (ProtocolObject): The Protocol param. ProtocolObject instance.

Param Region (required, string): The Region param. String length must exceed 1 characters.

Param SecondaryIpsecTunnel (string): The SecondaryIpsecTunnel param.

Param SpnName (string): The SpnName param.

Param Subnets (list of strings): The Subnets param.
*/
type Config struct {
    EcmpLoadBalancing string `json:"ecmp_load_balancing,omitempty"`
    EcmpTunnels []EcmpTunnelsObject `json:"ecmp_tunnels,omitempty"`
    ObjectId string `json:"id,omitempty"`
    IpsecTunnel string `json:"ipsec_tunnel,omitempty"`
    LicenseType string `json:"license_type"`
    Name string `json:"name"`
    Protocol *ProtocolObject `json:"protocol,omitempty"`
    Region string `json:"region"`
    SecondaryIpsecTunnel string `json:"secondary_ipsec_tunnel,omitempty"`
    SpnName string `json:"spn_name,omitempty"`
    Subnets []string `json:"subnets,omitempty"`
}

/*
EcmpTunnelsObject ""

Parent chains:
* "Config ecmp_tunnels"

Args:

Param DoNotExportRoutes (bool): The DoNotExportRoutes param.

Param IpsecTunnel (required, string): The IpsecTunnel param.

Param LocalIpAddress (string): The LocalIpAddress param.

Param Name (required, string): The Name param.

Param OriginateDefaultRoute (bool): The OriginateDefaultRoute param.

Param PeerAs (string): The PeerAs param.

Param PeerIpAddress (string): The PeerIpAddress param.

Param PeeringType (string): The PeeringType param. String values: []string{"exchange-v4-over-v4", "exchange-v4-v6-over-v4", "exchange-v4-over-v4-v6-over-v6", "exchange-v6-over-v6"}

Param Secret (string): The Secret param.

Param SummarizeMobileUserRoutes (bool): The SummarizeMobileUserRoutes param.
*/
type EcmpTunnelsObject struct {
    DoNotExportRoutes bool `json:"do_not_export_routes,omitempty"`
    IpsecTunnel string `json:"ipsec_tunnel"`
    LocalIpAddress string `json:"local_ip_address,omitempty"`
    Name string `json:"name"`
    OriginateDefaultRoute bool `json:"originate_default_route,omitempty"`
    PeerAs string `json:"peer_as,omitempty"`
    PeerIpAddress string `json:"peer_ip_address,omitempty"`
    PeeringType string `json:"peering_type,omitempty"`
    Secret string `json:"secret,omitempty"`
    SummarizeMobileUserRoutes bool `json:"summarize_mobile_user_routes,omitempty"`
}

/*
ProtocolObject ""

Parent chains:
* "Config protocol"

Args:

Param Bgp (BgpObject): The Bgp param. BgpObject instance.

Param BgpPeer (BgpPeerObject): The BgpPeer param. BgpPeerObject instance.
*/
type ProtocolObject struct {
    Bgp *BgpObject `json:"bgp,omitempty"`
    BgpPeer *BgpPeerObject `json:"bgp_peer,omitempty"`
}

/*
BgpObject ""

Parent chains:
* "Config protocol bgp"

Args:

Param DoNotExportRoutes (bool): The DoNotExportRoutes param.

Param Enable (bool): The Enable param.

Param LocalIpAddress (string): The LocalIpAddress param.

Param OriginateDefaultRoute (bool): The OriginateDefaultRoute param.

Param PeerAs (string): The PeerAs param.

Param PeerIpAddress (string): The PeerIpAddress param.

Param PeeringType (string): The PeeringType param. String values: []string{"exchange-v4-over-v4", "exchange-v4-v6-over-v4", "exchange-v4-over-v4-v6-over-v6", "exchange-v6-over-v6"}

Param Secret (string): The Secret param.

Param SummarizeMobileUserRoutes (bool): The SummarizeMobileUserRoutes param.
*/
type BgpObject struct {
    DoNotExportRoutes bool `json:"do_not_export_routes,omitempty"`
    Enable bool `json:"enable,omitempty"`
    LocalIpAddress string `json:"local_ip_address,omitempty"`
    OriginateDefaultRoute bool `json:"originate_default_route,omitempty"`
    PeerAs string `json:"peer_as,omitempty"`
    PeerIpAddress string `json:"peer_ip_address,omitempty"`
    PeeringType string `json:"peering_type,omitempty"`
    Secret string `json:"secret,omitempty"`
    SummarizeMobileUserRoutes bool `json:"summarize_mobile_user_routes,omitempty"`
}

/*
BgpPeerObject ""

Parent chains:
* "Config protocol bgp_peer"

Args:

Param LocalIpAddress (string): The LocalIpAddress param.

Param PeerIpAddress (string): The PeerIpAddress param.

Param Secret (string): The Secret param.
*/
type BgpPeerObject struct {
    LocalIpAddress string `json:"local_ip_address,omitempty"`
    PeerIpAddress string `json:"peer_ip_address,omitempty"`
    Secret string `json:"secret,omitempty"`
}