package gateways

/*
Config ""

Parent chains:
* "Config"

Args:

Param Authentication (required, AuthenticationObject): The Authentication param. AuthenticationObject instance.

Param ObjectId (string, read-only): The ObjectId param. Example: "abcd-1234".

Param LocalId (LocalIdObject): The LocalId param. LocalIdObject instance.

Param Name (required, string): The Name param. String length must not exceed 63 characters.

Param PeerAddress (required, PeerAddressObject): The PeerAddress param. PeerAddressObject instance.

Param PeerId (PeerIdObject): The PeerId param. PeerIdObject instance.

Param Protocol (required, ProtocolObject): The Protocol param. ProtocolObject instance.

Param ProtocolCommon (ProtocolCommonObject): The ProtocolCommon param. ProtocolCommonObject instance.
*/
type Config struct {
	Authentication AuthenticationObject  `json:"authentication"`
	ObjectId       string                `json:"id,omitempty"`
	LocalId        *LocalIdObject        `json:"local_id,omitempty"`
	Name           string                `json:"name"`
	PeerAddress    PeerAddressObject     `json:"peer_address"`
	PeerId         *PeerIdObject         `json:"peer_id,omitempty"`
	Protocol       ProtocolObject        `json:"protocol"`
	ProtocolCommon *ProtocolCommonObject `json:"protocol_common,omitempty"`
}

/*
AuthenticationObject ""

Parent chains:
* "Config authentication"

Args:

Param AllowIdPayloadMismatch (bool): The AllowIdPayloadMismatch param.

Param CertificateProfile (string): The CertificateProfile param.

Param LocalCertificate (LocalCertificateObject): The LocalCertificate param. LocalCertificateObject instance.

Param PreSharedKey (PreSharedKeyObject): The PreSharedKey param. PreSharedKeyObject instance.

Param StrictValidationRevocation (bool): The StrictValidationRevocation param.

Param UseManagementAsSource (bool): The UseManagementAsSource param.
*/
type AuthenticationObject struct {
	AllowIdPayloadMismatch     bool                    `json:"allow_id_payload_mismatch,omitempty"`
	CertificateProfile         string                  `json:"certificate_profile,omitempty"`
	LocalCertificate           *LocalCertificateObject `json:"local_certificate,omitempty"`
	PreSharedKey               *PreSharedKeyObject     `json:"pre_shared_key,omitempty"`
	StrictValidationRevocation bool                    `json:"strict_validation_revocation,omitempty"`
	UseManagementAsSource      bool                    `json:"use_management_as_source,omitempty"`
}

/*
LocalCertificateObject ""

Parent chains:
* "Config authentication local_certificate"

Args:

Param LocalCertificateName (string): The LocalCertificateName param.
*/
type LocalCertificateObject struct {
	LocalCertificateName string `json:"local_certificate_name,omitempty"`
}

/*
PreSharedKeyObject ""

Parent chains:
* "Config authentication pre_shared_key"

Args:

Param Key (string): The Key param.
*/
type PreSharedKeyObject struct {
	Key string `json:"key,omitempty"`
}

/*
LocalIdObject ""

Parent chains:
* "Config local_id"

Args:

Param ObjectId (string): The ObjectId param. String lengh must be between 1 and 1024 characters.

Param Type (string): The Type param.
*/
type LocalIdObject struct {
	ObjectId string `json:"id,omitempty"`
	Type     string `json:"type,omitempty"`
}

/*
PeerAddressObject ""

Parent chains:
* "Config peer_address"

Args:

Param DynamicValue (interface{}): The DynamicValue param. interface{} instance.

Param Fqdn (string): The Fqdn param. String length must not exceed 255 characters.

Param Ip (string): The Ip param.
*/
type PeerAddressObject struct {
	DynamicValue interface{} `json:"dynamic,omitempty"`
	Fqdn         string      `json:"fqdn,omitempty"`
	Ip           string      `json:"ip,omitempty"`
}

/*
PeerIdObject ""

Parent chains:
* "Config peer_id"

Args:

Param ObjectId (string): The ObjectId param. String lengh must be between 1 and 1024 characters.

Param Type (string): The Type param. String values: []string{"ipaddr", "keyid", "fqdn", "ufqdn"}
*/
type PeerIdObject struct {
	ObjectId string `json:"id,omitempty"`
	Type     string `json:"type,omitempty"`
}

/*
ProtocolObject ""

Parent chains:
* "Config protocol"

Args:

Param Ikev1 (Ikev1Object): The Ikev1 param. Ikev1Object instance.

Param Ikev2 (Ikev2Object): The Ikev2 param. Ikev2Object instance.

Param Version (string): The Version param. String values: []string{"ikev2-preferred", "ikev1", "ikev2"} Default: "ikev2-preferred".
*/
type ProtocolObject struct {
	Ikev1   *Ikev1Object `json:"ikev1,omitempty"`
	Ikev2   *Ikev2Object `json:"ikev2,omitempty"`
	Version string       `json:"version,omitempty"`
}

/*
Ikev1Object ""

Parent chains:
* "Config protocol ikev1"

Args:

Param Dpd (DpdObject): The Dpd param. DpdObject instance.

Param IkeCryptoProfile (string): The IkeCryptoProfile param.
*/
type Ikev1Object struct {
	Dpd              *DpdObject `json:"dpd,omitempty"`
	IkeCryptoProfile string     `json:"ike_crypto_profile,omitempty"`
}

/*
DpdObject ""

Parent chains:
* "Config protocol ikev1 dpd"
* "Config protocol ikev2 dpd"

Args:

Param Enable (bool): The Enable param.
*/
type DpdObject struct {
	Enable bool `json:"enable,omitempty"`
}

/*
Ikev2Object ""

Parent chains:
* "Config protocol ikev2"

Args:

Param Dpd (DpdObject): The Dpd param. DpdObject instance.

Param IkeCryptoProfile (string): The IkeCryptoProfile param.
*/
type Ikev2Object struct {
	Dpd              *DpdObject `json:"dpd,omitempty"`
	IkeCryptoProfile string     `json:"ike_crypto_profile,omitempty"`
}

/*
ProtocolCommonObject ""

Parent chains:
* "Config protocol_common"

Args:

Param Fragmentation (FragmentationObject): The Fragmentation param. FragmentationObject instance.

Param NatTraversal (NatTraversalObject): The NatTraversal param. NatTraversalObject instance.

Param PassiveMode (bool): The PassiveMode param.
*/
type ProtocolCommonObject struct {
	Fragmentation *FragmentationObject `json:"fragmentation,omitempty"`
	NatTraversal  *NatTraversalObject  `json:"nat_traversal,omitempty"`
	PassiveMode   bool                 `json:"passive_mode,omitempty"`
}

/*
FragmentationObject ""

Parent chains:
* "Config protocol_common fragmentation"

Args:

Param Enable (bool): The Enable param. Default: false
*/
type FragmentationObject struct {
	Enable bool `json:"enable,omitempty"`
}

/*
NatTraversalObject ""

Parent chains:
* "Config protocol_common nat_traversal"

Args:

Param Enable (bool): The Enable param.
*/
type NatTraversalObject struct {
	Enable bool `json:"enable,omitempty"`
}
