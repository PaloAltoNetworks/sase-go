package settings

/*
Config ""

Parent chains:
* "Config"

Args:

Param DnsServers (required, list of DnsServersObject objects): The DnsServers param.

Param EnableWins (EnableWinsObject): The EnableWins param. EnableWinsObject instance.

Param IpPools (required, list of IpPoolsObject objects): The IpPools param.

Param Ipv6 (bool): The Ipv6 param.

Param Name (required, string): The Name param.

Param PortalHostname (required, PortalHostnameObject): The PortalHostname param. PortalHostnameObject instance.

Param RegionIpv6 (required, RegionIpv6Object): The RegionIpv6 param. RegionIpv6Object instance.

Param UdpQueries (UdpQueriesObject): The UdpQueries param. UdpQueriesObject instance.
*/
type Config struct {
    DnsServers []DnsServersObject `json:"dns_servers"`
    EnableWins *EnableWinsObject `json:"enable_wins,omitempty"`
    IpPools []IpPoolsObject `json:"ip_pools"`
    Ipv6 bool `json:"ipv6,omitempty"`
    Name string `json:"name"`
    PortalHostname PortalHostnameObject `json:"portal_hostname"`
    RegionIpv6 RegionIpv6Object `json:"region_ipv6"`
    UdpQueries *UdpQueriesObject `json:"udp_queries,omitempty"`
}

/*
DnsServersObject ""

Parent chains:
* "Config dns_servers"

Args:

Param DnsSuffix (list of strings): The DnsSuffix param.

Param InternalDnsMatch (list of InternalDnsMatchObject objects): The InternalDnsMatch param.

Param Name (string): The Name param.

Param PrimaryPublicDns (PrimaryPublicDnsObject): The PrimaryPublicDns param. PrimaryPublicDnsObject instance.

Param SecondaryPublicDns (SecondaryPublicDnsObject): The SecondaryPublicDns param. SecondaryPublicDnsObject instance.
*/
type DnsServersObject struct {
    DnsSuffix []string `json:"dns_suffix,omitempty"`
    InternalDnsMatch []InternalDnsMatchObject `json:"internal_dns_match,omitempty"`
    Name string `json:"name,omitempty"`
    PrimaryPublicDns *PrimaryPublicDnsObject `json:"primary_public_dns,omitempty"`
    SecondaryPublicDns *SecondaryPublicDnsObject `json:"secondary_public_dns,omitempty"`
}

/*
InternalDnsMatchObject ""

Parent chains:
* "Config dns_servers internal_dns_match"

Args:

Param DomainList (list of strings): The DomainList param.

Param Name (string): The Name param.

Param Primary (PrimaryObject): The Primary param. PrimaryObject instance.

Param Secondary (SecondaryObject): The Secondary param. SecondaryObject instance.
*/
type InternalDnsMatchObject struct {
    DomainList []string `json:"domain_list,omitempty"`
    Name string `json:"name,omitempty"`
    Primary *PrimaryObject `json:"primary,omitempty"`
    Secondary *SecondaryObject `json:"secondary,omitempty"`
}

/*
PrimaryObject ""

Parent chains:
* "Config dns_servers internal_dns_match primary"

Args:

Param DnsServer (interface{}): The DnsServer param. interface{} instance.

Param UseCloudDefault (interface{}): The UseCloudDefault param. interface{} instance.
*/
type PrimaryObject struct {
    DnsServer interface{} `json:"dns_server,omitempty"`
    UseCloudDefault interface{} `json:"use_cloud_default,omitempty"`
}

/*
SecondaryObject ""

Parent chains:
* "Config dns_servers internal_dns_match secondary"

Args:

Param DnsServer (interface{}): The DnsServer param. interface{} instance.

Param UseCloudDefault (interface{}): The UseCloudDefault param. interface{} instance.
*/
type SecondaryObject struct {
    DnsServer interface{} `json:"dns_server,omitempty"`
    UseCloudDefault interface{} `json:"use_cloud_default,omitempty"`
}

/*
PrimaryPublicDnsObject ""

Parent chains:
* "Config dns_servers primary_public_dns"

Args:

Param DnsServer (string): The DnsServer param.
*/
type PrimaryPublicDnsObject struct {
    DnsServer string `json:"dns_server,omitempty"`
}

/*
SecondaryPublicDnsObject ""

Parent chains:
* "Config dns_servers secondary_public_dns"

Args:

Param DnsServer (string): The DnsServer param.
*/
type SecondaryPublicDnsObject struct {
    DnsServer string `json:"dns_server,omitempty"`
}

/*
EnableWinsObject ""

Parent chains:
* "Config enable_wins"

Args:

Param No (interface{}): The No param. interface{} instance.

Param Yes (YesObject): The Yes param. YesObject instance.
*/
type EnableWinsObject struct {
    No interface{} `json:"no,omitempty"`
    Yes *YesObject `json:"yes,omitempty"`
}

/*
YesObject ""

Parent chains:
* "Config enable_wins yes"

Args:

Param WinsServers (list of WinsServersObject objects): The WinsServers param.
*/
type YesObject struct {
    WinsServers []WinsServersObject `json:"wins_servers,omitempty"`
}

/*
WinsServersObject ""

Parent chains:
* "Config enable_wins yes wins_servers"

Args:

Param Name (string): The Name param.

Param Primary (string): The Primary param.

Param Secondary (string): The Secondary param.
*/
type WinsServersObject struct {
    Name string `json:"name,omitempty"`
    Primary string `json:"primary,omitempty"`
    Secondary string `json:"secondary,omitempty"`
}

/*
IpPoolsObject ""

Parent chains:
* "Config ip_pools"

Args:

Param IpPool (list of strings): The IpPool param.

Param Name (string): The Name param.
*/
type IpPoolsObject struct {
    IpPool []string `json:"ip_pool,omitempty"`
    Name string `json:"name,omitempty"`
}

/*
PortalHostnameObject ""

Parent chains:
* "Config portal_hostname"

Args:

Param CustomDomain (CustomDomainObject): The CustomDomain param. CustomDomainObject instance.

Param DefaultDomain (DefaultDomainObject): The DefaultDomain param. DefaultDomainObject instance.
*/
type PortalHostnameObject struct {
    CustomDomain *CustomDomainObject `json:"custom_domain,omitempty"`
    DefaultDomain *DefaultDomainObject `json:"default_domain,omitempty"`
}

/*
CustomDomainObject ""

Parent chains:
* "Config portal_hostname custom_domain"

Args:

Param Cname (string): The Cname param.

Param Hostname (string): The Hostname param.

Param SslTlsServiceProfile (string): The SslTlsServiceProfile param.
*/
type CustomDomainObject struct {
    Cname string `json:"cname,omitempty"`
    Hostname string `json:"hostname,omitempty"`
    SslTlsServiceProfile string `json:"ssl_tls_service_profile,omitempty"`
}

/*
DefaultDomainObject ""

Parent chains:
* "Config portal_hostname default_domain"

Args:

Param Hostname (string): The Hostname param.
*/
type DefaultDomainObject struct {
    Hostname string `json:"hostname,omitempty"`
}

/*
RegionIpv6Object ""

Parent chains:
* "Config region_ipv6"

Args:

Param Region (list of RegionObject objects): The Region param.
*/
type RegionIpv6Object struct {
    Region []RegionObject `json:"region,omitempty"`
}

/*
RegionObject ""

Parent chains:
* "Config region_ipv6 region"

Args:

Param Locations (list of strings): The Locations param.

Param Name (string): The Name param.
*/
type RegionObject struct {
    Locations []string `json:"locations,omitempty"`
    Name string `json:"name,omitempty"`
}

/*
UdpQueriesObject ""

Parent chains:
* "Config udp_queries"

Args:

Param Retries (RetriesObject): The Retries param. RetriesObject instance.
*/
type UdpQueriesObject struct {
    Retries *RetriesObject `json:"retries,omitempty"`
}

/*
RetriesObject ""

Parent chains:
* "Config udp_queries retries"

Args:

Param Attempts (int64): The Attempts param. Value must be between 1 and 30.

Param Interval (int64): The Interval param. Value must be between 1 and 30.
*/
type RetriesObject struct {
    Attempts int64 `json:"attempts,omitempty"`
    Interval int64 `json:"interval,omitempty"`
}