package profiles

/*
Config ""

Parent chains:
* "Config"

Args:

Param BotnetDomains (BotnetDomainsObject): The BotnetDomains param. BotnetDomainsObject instance.

Param Description (string): The Description param.

Param ObjectId (string, read-only): The ObjectId param. Example: "abcd-1234".

Param Name (string): The Name param.
*/
type Config struct {
	BotnetDomains *BotnetDomainsObject `json:"botnet_domains,omitempty"`
	Description   string               `json:"description,omitempty"`
	ObjectId      string               `json:"id,omitempty"`
	Name          string               `json:"name,omitempty"`
}

/*
BotnetDomainsObject ""

Parent chains:
* "Config botnet_domains"

Args:

Param DnsSecurityCategories (list of DnsSecurityCategoriesObject objects): The DnsSecurityCategories param.

Param Lists (list of ListsObject objects): The Lists param.

Param Sinkhole (SinkholeObject): The Sinkhole param. SinkholeObject instance.

Param Whitelist (list of WhitelistObject objects): The Whitelist param.
*/
type BotnetDomainsObject struct {
	DnsSecurityCategories []DnsSecurityCategoriesObject `json:"dns_security_categories,omitempty"`
	Lists                 []ListsObject                 `json:"lists,omitempty"`
	Sinkhole              *SinkholeObject               `json:"sinkhole,omitempty"`
	Whitelist             []WhitelistObject             `json:"whitelist,omitempty"`
}

/*
DnsSecurityCategoriesObject ""

Parent chains:
* "Config botnet_domains dns_security_categories"

Args:

Param Action (string): The Action param. String values: []string{"default", "allow", "block", "sinkhole"} Default: "default".

Param LogLevel (string): The LogLevel param. String values: []string{"default", "none", "low", "informational", "medium", "high", "critical"} Default: "default".

Param Name (string): The Name param.

Param PacketCapture (string): The PacketCapture param. String values: []string{"disable", "single-packet", "extended-capture"}
*/
type DnsSecurityCategoriesObject struct {
	Action        string `json:"action,omitempty"`
	LogLevel      string `json:"log_level,omitempty"`
	Name          string `json:"name,omitempty"`
	PacketCapture string `json:"packet_capture,omitempty"`
}

/*
ListsObject ""

Parent chains:
* "Config botnet_domains lists"

Args:

Param Action (ActionObject): The Action param. ActionObject instance.

Param Name (required, string): The Name param.

Param PacketCapture (string): The PacketCapture param. String values: []string{"disable", "single-packet", "extended-capture"}
*/
type ListsObject struct {
	Action        *ActionObject `json:"action,omitempty"`
	Name          string        `json:"name"`
	PacketCapture string        `json:"packet_capture,omitempty"`
}

/*
ActionObject ""

Parent chains:
* "Config botnet_domains lists action"

Args:

Param Alert (interface{}): The Alert param. interface{} instance.

Param Allow (interface{}): The Allow param. interface{} instance.

Param Block (interface{}): The Block param. interface{} instance.

Param Sinkhole (interface{}): The Sinkhole param. interface{} instance.
*/
type ActionObject struct {
	Alert    interface{} `json:"alert,omitempty"`
	Allow    interface{} `json:"allow,omitempty"`
	Block    interface{} `json:"block,omitempty"`
	Sinkhole interface{} `json:"sinkhole,omitempty"`
}

/*
SinkholeObject ""

Parent chains:
* "Config botnet_domains sinkhole"

Args:

Param Ipv4Address (string): The Ipv4Address param. String values: []string{"127.0.0.1", "pan-sinkhole-default-ip"}

Param Ipv6Address (string): The Ipv6Address param. String values: []string{"::1"}
*/
type SinkholeObject struct {
	Ipv4Address string `json:"ipv4_address,omitempty"`
	Ipv6Address string `json:"ipv6_address,omitempty"`
}

/*
WhitelistObject ""

Parent chains:
* "Config botnet_domains whitelist"

Args:

Param Description (string): The Description param.

Param Name (required, string): The Name param.
*/
type WhitelistObject struct {
	Description string `json:"description,omitempty"`
	Name        string `json:"name"`
}
