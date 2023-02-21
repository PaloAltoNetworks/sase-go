package connections

/*
Config ""

Parent chains:
* "Config"

Args:

Param BackupSC (string): The BackupSC param.

Param BgpPeer (BgpPeerObject): The BgpPeer param. BgpPeerObject instance.

Param IpsecTunnel (required, string): The IpsecTunnel param.

Param Name (required, string): The Name param.

Param NatPool (string): The NatPool param.

Param NoExportCommunity (string): The NoExportCommunity param. String values: []string{"Disabled", "Enabled-In", "Enabled-Out", "Enabled-Both"}

Param OnboardingType (string): The OnboardingType param. String values: []string{"classic"} Default: "classic".

Param Protocol (ProtocolObject): The Protocol param. ProtocolObject instance.

Param Qos (QosObject): The Qos param. QosObject instance.

Param Region (required, string): The Region param.

Param SecondaryIpsecTunnel (string): The SecondaryIpsecTunnel param.

Param SourceNat (bool): The SourceNat param.

Param Subnets (list of strings): The Subnets param.
*/
type Config struct {
    BackupSC string `json:"backup_SC,omitempty"`
    BgpPeer *BgpPeerObject `json:"bgp_peer,omitempty"`
    IpsecTunnel string `json:"ipsec_tunnel"`
    Name string `json:"name"`
    NatPool string `json:"nat_pool,omitempty"`
    NoExportCommunity string `json:"no_export_community,omitempty"`
    OnboardingType string `json:"onboarding_type,omitempty"`
    Protocol *ProtocolObject `json:"protocol,omitempty"`
    Qos *QosObject `json:"qos,omitempty"`
    Region string `json:"region"`
    SecondaryIpsecTunnel string `json:"secondary_ipsec_tunnel,omitempty"`
    SourceNat bool `json:"source_nat,omitempty"`
    Subnets []string `json:"subnets,omitempty"`
}

/*
BgpPeerObject ""

Parent chains:
* "Config bgp_peer"

Args:

Param LocalIpAddress (string): The LocalIpAddress param.

Param LocalIpv6Address (string): The LocalIpv6Address param.

Param PeerIpAddress (string): The PeerIpAddress param.

Param PeerIpv6Address (string): The PeerIpv6Address param.

Param SameAsPrimary (bool): The SameAsPrimary param.

Param Secret (string): The Secret param.
*/
type BgpPeerObject struct {
    LocalIpAddress string `json:"local_ip_address,omitempty"`
    LocalIpv6Address string `json:"local_ipv6_address,omitempty"`
    PeerIpAddress string `json:"peer_ip_address,omitempty"`
    PeerIpv6Address string `json:"peer_ipv6_address,omitempty"`
    SameAsPrimary bool `json:"same_as_primary,omitempty"`
    Secret string `json:"secret,omitempty"`
}

/*
ProtocolObject ""

Parent chains:
* "Config protocol"

Args:

Param Bgp (BgpObject): The Bgp param. BgpObject instance.
*/
type ProtocolObject struct {
    Bgp *BgpObject `json:"bgp,omitempty"`
}

/*
BgpObject ""

Parent chains:
* "Config protocol bgp"

Args:

Param DoNotExportRoutes (bool): The DoNotExportRoutes param.

Param Enable (bool): The Enable param.

Param FastFailover (bool): The FastFailover param.

Param LocalIpAddress (string): The LocalIpAddress param.

Param OriginateDefaultRoute (bool): The OriginateDefaultRoute param.

Param PeerAs (string): The PeerAs param.

Param PeerIpAddress (string): The PeerIpAddress param.

Param Secret (string): The Secret param.

Param SummarizeMobileUserRoutes (bool): The SummarizeMobileUserRoutes param.
*/
type BgpObject struct {
    DoNotExportRoutes bool `json:"do_not_export_routes,omitempty"`
    Enable bool `json:"enable,omitempty"`
    FastFailover bool `json:"fast_failover,omitempty"`
    LocalIpAddress string `json:"local_ip_address,omitempty"`
    OriginateDefaultRoute bool `json:"originate_default_route,omitempty"`
    PeerAs string `json:"peer_as,omitempty"`
    PeerIpAddress string `json:"peer_ip_address,omitempty"`
    Secret string `json:"secret,omitempty"`
    SummarizeMobileUserRoutes bool `json:"summarize_mobile_user_routes,omitempty"`
}

/*
QosObject ""

Parent chains:
* "Config qos"

Args:

Param Enable (bool): The Enable param.

Param QosProfile (string): The QosProfile param.
*/
type QosObject struct {
    Enable bool `json:"enable,omitempty"`
    QosProfile string `json:"qos_profile,omitempty"`
}