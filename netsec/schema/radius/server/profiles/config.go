package profiles

/*
Config ""

Parent chains:
* "Config"

Args:

Param ObjectId (string, read-only): The ObjectId param. Example: "abcd-1234".

Param Protocol (ProtocolObject): The Protocol param. ProtocolObject instance.

Param Retries (int64): The Retries param. Value must be between 1 and 5.

Param Server (required, list of ServerObject objects): The Server param.

Param Timeout (int64): The Timeout param. Value must be between 1 and 120.
*/
type Config struct {
	ObjectId string          `json:"id,omitempty"`
	Protocol *ProtocolObject `json:"protocol,omitempty"`
	Retries  int64           `json:"retries,omitempty"`
	Server   []ServerObject  `json:"server"`
	Timeout  int64           `json:"timeout,omitempty"`
}

/*
ProtocolObject ""

Parent chains:
* "Config protocol"

Args:

Param CHAP (interface{}): The CHAP param. interface{} instance.

Param EAPTTLSWithPAP (EAPTTLSWithPAPObject): The EAPTTLSWithPAP param. EAPTTLSWithPAPObject instance.

Param PAP (interface{}): The PAP param. interface{} instance.

Param PEAPMSCHAPv2 (PEAPMSCHAPv2Object): The PEAPMSCHAPv2 param. PEAPMSCHAPv2Object instance.

Param PEAPWithGTC (PEAPWithGTCObject): The PEAPWithGTC param. PEAPWithGTCObject instance.
*/
type ProtocolObject struct {
	CHAP           interface{}           `json:"CHAP,omitempty"`
	EAPTTLSWithPAP *EAPTTLSWithPAPObject `json:"EAP_TTLS_with_PAP,omitempty"`
	PAP            interface{}           `json:"PAP,omitempty"`
	PEAPMSCHAPv2   *PEAPMSCHAPv2Object   `json:"PEAP_MSCHAPv2,omitempty"`
	PEAPWithGTC    *PEAPWithGTCObject    `json:"PEAP_with_GTC,omitempty"`
}

/*
EAPTTLSWithPAPObject ""

Parent chains:
* "Config protocol EAP_TTLS_with_PAP"

Args:

Param AnonOuterId (bool): The AnonOuterId param.

Param RadiusCertProfile (string): The RadiusCertProfile param.
*/
type EAPTTLSWithPAPObject struct {
	AnonOuterId       bool   `json:"anon_outer_id,omitempty"`
	RadiusCertProfile string `json:"radius_cert_profile,omitempty"`
}

/*
PEAPMSCHAPv2Object ""

Parent chains:
* "Config protocol PEAP_MSCHAPv2"

Args:

Param AllowPwdChange (bool): The AllowPwdChange param.

Param AnonOuterId (bool): The AnonOuterId param.

Param RadiusCertProfile (string): The RadiusCertProfile param.
*/
type PEAPMSCHAPv2Object struct {
	AllowPwdChange    bool   `json:"allow_pwd_change,omitempty"`
	AnonOuterId       bool   `json:"anon_outer_id,omitempty"`
	RadiusCertProfile string `json:"radius_cert_profile,omitempty"`
}

/*
PEAPWithGTCObject ""

Parent chains:
* "Config protocol PEAP_with_GTC"

Args:

Param AnonOuterId (bool): The AnonOuterId param.

Param RadiusCertProfile (string): The RadiusCertProfile param.
*/
type PEAPWithGTCObject struct {
	AnonOuterId       bool   `json:"anon_outer_id,omitempty"`
	RadiusCertProfile string `json:"radius_cert_profile,omitempty"`
}

/*
ServerObject ""

Parent chains:
* "Config server"

Args:

Param IpAddress (string): The IpAddress param.

Param Name (string): The Name param.

Param Port (int64): The Port param. Value must be between 1 and 65535.

Param Secret (string): The Secret param. String length must not exceed 64 characters.
*/
type ServerObject struct {
	IpAddress string `json:"ip_address,omitempty"`
	Name      string `json:"name,omitempty"`
	Port      int64  `json:"port,omitempty"`
	Secret    string `json:"secret,omitempty"`
}
