package services

/*
Config ""

Parent chains:
* "Config"

Args:

Param Description (string): The Description param. String lengh must be between 0 and 1023 characters.

Param ObjectId (string, read-only): The ObjectId param. Example: "abcd-1234".

Param Name (required, string): The Name param. String length must not exceed 63 characters.

Param Protocol (required, ProtocolObject): The Protocol param. ProtocolObject instance.

Param Tag (list of strings): The Tag param. Array should contain at most 64 items.
*/
type Config struct {
	Description string         `json:"description,omitempty"`
	ObjectId    string         `json:"id,omitempty"`
	Name        string         `json:"name"`
	Protocol    ProtocolObject `json:"protocol"`
	Tag         []string       `json:"tag,omitempty"`
}

/*
ProtocolObject ""

Parent chains:
* "Config protocol"

Args:

Param Tcp (TcpObject): The Tcp param. TcpObject instance.

Param Udp (UdpObject): The Udp param. UdpObject instance.
*/
type ProtocolObject struct {
	Tcp *TcpObject `json:"tcp,omitempty"`
	Udp *UdpObject `json:"udp,omitempty"`
}

/*
TcpObject ""

Parent chains:
* "Config protocol tcp"

Args:

Param Override (OverrideObject): The Override param. OverrideObject instance.

Param Port (required, string): The Port param. String lengh must be between 1 and 1023 characters.

Param SourcePort (string): The SourcePort param. String lengh must be between 1 and 1023 characters.
*/
type TcpObject struct {
	Override   *OverrideObject `json:"override,omitempty"`
	Port       string          `json:"port"`
	SourcePort string          `json:"source_port,omitempty"`
}

/*
OverrideObject ""

Parent chains:
* "Config protocol tcp override"

Args:

Param HalfcloseTimeout (int64): The HalfcloseTimeout param. Value must be between 1 and 604800. Default: 120

Param Timeout (int64): The Timeout param. Value must be between 1 and 604800. Default: 3600

Param TimewaitTimeout (int64): The TimewaitTimeout param. Value must be between 1 and 600. Default: 15
*/
type OverrideObject struct {
	HalfcloseTimeout int64 `json:"halfclose_timeout,omitempty"`
	Timeout          int64 `json:"timeout,omitempty"`
	TimewaitTimeout  int64 `json:"timewait_timeout,omitempty"`
}

/*
UdpObject ""

Parent chains:
* "Config protocol udp"

Args:

Param Override (OverrideObject1): The Override param. OverrideObject1 instance.

Param Port (required, string): The Port param. String lengh must be between 1 and 1023 characters.

Param SourcePort (string): The SourcePort param. String lengh must be between 1 and 1023 characters.
*/
type UdpObject struct {
	Override   *OverrideObject1 `json:"override,omitempty"`
	Port       string           `json:"port"`
	SourcePort string           `json:"source_port,omitempty"`
}

/*
OverrideObject1 ""

Parent chains:
* "Config protocol udp override"

Args:

Param Timeout (int64): The Timeout param. Value must be between 1 and 604800. Default: 30
*/
type OverrideObject1 struct {
	Timeout int64 `json:"timeout,omitempty"`
}
