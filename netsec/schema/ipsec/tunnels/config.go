package tunnels

/*
Config ""

Parent chains:
* "Config"

Args:

Param AntiReplay (bool): The AntiReplay param.

Param AutoKey (required, AutoKeyObject): The AutoKey param. AutoKeyObject instance.

Param CopyTos (bool): The CopyTos param. Default: false

Param EnableGreEncapsulation (bool): The EnableGreEncapsulation param. Default: false

Param ObjectId (string, read-only): The ObjectId param. Example: "abcd-1234".

Param Name (required, string): The Name param. String length must not exceed 63 characters.

Param TunnelMonitor (TunnelMonitorObject): The TunnelMonitor param. TunnelMonitorObject instance.
*/
type Config struct {
	AntiReplay             bool                 `json:"anti_replay,omitempty"`
	AutoKey                AutoKeyObject        `json:"auto_key"`
	CopyTos                bool                 `json:"copy_tos,omitempty"`
	EnableGreEncapsulation bool                 `json:"enable_gre_encapsulation,omitempty"`
	ObjectId               string               `json:"id,omitempty"`
	Name                   string               `json:"name"`
	TunnelMonitor          *TunnelMonitorObject `json:"tunnel_monitor,omitempty"`
}

/*
AutoKeyObject ""

Parent chains:
* "Config auto_key"

Args:

Param IkeGateway (required, list of IkeGatewayObject objects): The IkeGateway param.

Param IpsecCryptoProfile (required, string): The IpsecCryptoProfile param.

Param ProxyId (list of ProxyIdObject objects): The ProxyId param.
*/
type AutoKeyObject struct {
	IkeGateway         []IkeGatewayObject `json:"ike_gateway"`
	IpsecCryptoProfile string             `json:"ipsec_crypto_profile"`
	ProxyId            []ProxyIdObject    `json:"proxy_id,omitempty"`
}

/*
IkeGatewayObject ""

Parent chains:
* "Config auto_key ike_gateway"

Args:

Param Name (string): The Name param.
*/
type IkeGatewayObject struct {
	Name string `json:"name,omitempty"`
}

/*
ProxyIdObject ""

Parent chains:
* "Config auto_key proxy_id"

Args:

Param Local (string): The Local param.

Param Name (required, string): The Name param.

Param Protocol (ProtocolObject): The Protocol param. ProtocolObject instance.

Param Remote (string): The Remote param.
*/
type ProxyIdObject struct {
	Local    string          `json:"local,omitempty"`
	Name     string          `json:"name"`
	Protocol *ProtocolObject `json:"protocol,omitempty"`
	Remote   string          `json:"remote,omitempty"`
}

/*
ProtocolObject ""

Parent chains:
* "Config auto_key proxy_id protocol"

Args:

Param Number (int64): The Number param. Value must be between 1 and 254.

Param Tcp (TcpObject): The Tcp param. TcpObject instance.

Param Udp (UdpObject): The Udp param. UdpObject instance.
*/
type ProtocolObject struct {
	Number int64      `json:"number,omitempty"`
	Tcp    *TcpObject `json:"tcp,omitempty"`
	Udp    *UdpObject `json:"udp,omitempty"`
}

/*
TcpObject ""

Parent chains:
* "Config auto_key proxy_id protocol tcp"

Args:

Param LocalPort (int64): The LocalPort param. Value must be between 0 and 65535. Default: 0

Param RemotePort (int64): The RemotePort param. Value must be between 0 and 65535. Default: 0
*/
type TcpObject struct {
	LocalPort  int64 `json:"local_port,omitempty"`
	RemotePort int64 `json:"remote_port,omitempty"`
}

/*
UdpObject ""

Parent chains:
* "Config auto_key proxy_id protocol udp"

Args:

Param LocalPort (int64): The LocalPort param. Value must be between 0 and 65535. Default: 0

Param RemotePort (int64): The RemotePort param. Value must be between 0 and 65535. Default: 0
*/
type UdpObject struct {
	LocalPort  int64 `json:"local_port,omitempty"`
	RemotePort int64 `json:"remote_port,omitempty"`
}

/*
TunnelMonitorObject ""

Parent chains:
* "Config tunnel_monitor"

Args:

Param DestinationIp (required, string): The DestinationIp param.

Param Enable (bool): The Enable param. Default: true

Param ProxyId (string): The ProxyId param.
*/
type TunnelMonitorObject struct {
	DestinationIp string `json:"destination_ip"`
	Enable        bool   `json:"enable,omitempty"`
	ProxyId       string `json:"proxy_id,omitempty"`
}
