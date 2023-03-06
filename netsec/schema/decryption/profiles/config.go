package profiles

/*
Config ""

Parent chains:
* "Config"

Args:

Param ObjectId (string, read-only): The ObjectId param. Example: "abcd-1234".

Param Name (required, string): The Name param.

Param SslForwardProxy (SslForwardProxyObject): The SslForwardProxy param. SslForwardProxyObject instance.

Param SslInboundProxy (SslInboundProxyObject): The SslInboundProxy param. SslInboundProxyObject instance.

Param SslNoProxy (SslNoProxyObject): The SslNoProxy param. SslNoProxyObject instance.

Param SslProtocolSettings (SslProtocolSettingsObject): The SslProtocolSettings param. SslProtocolSettingsObject instance.
*/
type Config struct {
	ObjectId            string                     `json:"id,omitempty"`
	Name                string                     `json:"name"`
	SslForwardProxy     *SslForwardProxyObject     `json:"ssl_forward_proxy,omitempty"`
	SslInboundProxy     *SslInboundProxyObject     `json:"ssl_inbound_proxy,omitempty"`
	SslNoProxy          *SslNoProxyObject          `json:"ssl_no_proxy,omitempty"`
	SslProtocolSettings *SslProtocolSettingsObject `json:"ssl_protocol_settings,omitempty"`
}

/*
SslForwardProxyObject ""

Parent chains:
* "Config ssl_forward_proxy"

Args:

Param AutoIncludeAltname (bool): The AutoIncludeAltname param. Default: false

Param BlockClientCert (bool): The BlockClientCert param. Default: false

Param BlockExpiredCertificate (bool): The BlockExpiredCertificate param. Default: false

Param BlockTimeoutCert (bool): The BlockTimeoutCert param. Default: false

Param BlockTls13DowngradeNoResource (bool): The BlockTls13DowngradeNoResource param. Default: false

Param BlockUnknownCert (bool): The BlockUnknownCert param. Default: false

Param BlockUnsupportedCipher (bool): The BlockUnsupportedCipher param. Default: false

Param BlockUnsupportedVersion (bool): The BlockUnsupportedVersion param. Default: false

Param BlockUntrustedIssuer (bool): The BlockUntrustedIssuer param. Default: false

Param RestrictCertExts (bool): The RestrictCertExts param. Default: false

Param StripAlpn (bool): The StripAlpn param. Default: false
*/
type SslForwardProxyObject struct {
	AutoIncludeAltname            bool `json:"auto_include_altname,omitempty"`
	BlockClientCert               bool `json:"block_client_cert,omitempty"`
	BlockExpiredCertificate       bool `json:"block_expired_certificate,omitempty"`
	BlockTimeoutCert              bool `json:"block_timeout_cert,omitempty"`
	BlockTls13DowngradeNoResource bool `json:"block_tls13_downgrade_no_resource,omitempty"`
	BlockUnknownCert              bool `json:"block_unknown_cert,omitempty"`
	BlockUnsupportedCipher        bool `json:"block_unsupported_cipher,omitempty"`
	BlockUnsupportedVersion       bool `json:"block_unsupported_version,omitempty"`
	BlockUntrustedIssuer          bool `json:"block_untrusted_issuer,omitempty"`
	RestrictCertExts              bool `json:"restrict_cert_exts,omitempty"`
	StripAlpn                     bool `json:"strip_alpn,omitempty"`
}

/*
SslInboundProxyObject ""

Parent chains:
* "Config ssl_inbound_proxy"

Args:

Param BlockIfHsmUnavailable (bool): The BlockIfHsmUnavailable param. Default: false

Param BlockIfNoResource (bool): The BlockIfNoResource param. Default: false

Param BlockUnsupportedCipher (bool): The BlockUnsupportedCipher param. Default: false

Param BlockUnsupportedVersion (bool): The BlockUnsupportedVersion param. Default: false
*/
type SslInboundProxyObject struct {
	BlockIfHsmUnavailable   bool `json:"block_if_hsm_unavailable,omitempty"`
	BlockIfNoResource       bool `json:"block_if_no_resource,omitempty"`
	BlockUnsupportedCipher  bool `json:"block_unsupported_cipher,omitempty"`
	BlockUnsupportedVersion bool `json:"block_unsupported_version,omitempty"`
}

/*
SslNoProxyObject ""

Parent chains:
* "Config ssl_no_proxy"

Args:

Param BlockExpiredCertificate (bool): The BlockExpiredCertificate param. Default: false

Param BlockUntrustedIssuer (bool): The BlockUntrustedIssuer param. Default: false
*/
type SslNoProxyObject struct {
	BlockExpiredCertificate bool `json:"block_expired_certificate,omitempty"`
	BlockUntrustedIssuer    bool `json:"block_untrusted_issuer,omitempty"`
}

/*
SslProtocolSettingsObject ""

Parent chains:
* "Config ssl_protocol_settings"

Args:

Param AuthAlgoMd5 (bool): The AuthAlgoMd5 param. Default: true

Param AuthAlgoSha1 (bool): The AuthAlgoSha1 param. Default: true

Param AuthAlgoSha256 (bool): The AuthAlgoSha256 param. Default: true

Param AuthAlgoSha384 (bool): The AuthAlgoSha384 param. Default: true

Param EncAlgo3des (bool): The EncAlgo3des param. Default: true

Param EncAlgoAes128Cbc (bool): The EncAlgoAes128Cbc param. Default: true

Param EncAlgoAes128Gcm (bool): The EncAlgoAes128Gcm param. Default: true

Param EncAlgoAes256Cbc (bool): The EncAlgoAes256Cbc param. Default: true

Param EncAlgoAes256Gcm (bool): The EncAlgoAes256Gcm param. Default: true

Param EncAlgoChacha20Poly1305 (bool): The EncAlgoChacha20Poly1305 param. Default: true

Param EncAlgoRc4 (bool): The EncAlgoRc4 param. Default: true

Param KeyxchgAlgoDhe (bool): The KeyxchgAlgoDhe param. Default: true

Param KeyxchgAlgoEcdhe (bool): The KeyxchgAlgoEcdhe param. Default: true

Param KeyxchgAlgoRsa (bool): The KeyxchgAlgoRsa param. Default: true

Param MaxVersion (string): The MaxVersion param. String values: []string{"sslv3", "tls1-0", "tls1-1", "tls1-2", "tls1-3", "max"} Default: "tls1-2".

Param MinVersion (string): The MinVersion param. String values: []string{"sslv3", "tls1-0", "tls1-1", "tls1-2", "tls1-3"} Default: "tls1-0".
*/
type SslProtocolSettingsObject struct {
	AuthAlgoMd5             bool   `json:"auth_algo_md5,omitempty"`
	AuthAlgoSha1            bool   `json:"auth_algo_sha1,omitempty"`
	AuthAlgoSha256          bool   `json:"auth_algo_sha256,omitempty"`
	AuthAlgoSha384          bool   `json:"auth_algo_sha384,omitempty"`
	EncAlgo3des             bool   `json:"enc_algo_3des,omitempty"`
	EncAlgoAes128Cbc        bool   `json:"enc_algo_aes_128_cbc,omitempty"`
	EncAlgoAes128Gcm        bool   `json:"enc_algo_aes_128_gcm,omitempty"`
	EncAlgoAes256Cbc        bool   `json:"enc_algo_aes_256_cbc,omitempty"`
	EncAlgoAes256Gcm        bool   `json:"enc_algo_aes_256_gcm,omitempty"`
	EncAlgoChacha20Poly1305 bool   `json:"enc_algo_chacha20_poly1305,omitempty"`
	EncAlgoRc4              bool   `json:"enc_algo_rc4,omitempty"`
	KeyxchgAlgoDhe          bool   `json:"keyxchg_algo_dhe,omitempty"`
	KeyxchgAlgoEcdhe        bool   `json:"keyxchg_algo_ecdhe,omitempty"`
	KeyxchgAlgoRsa          bool   `json:"keyxchg_algo_rsa,omitempty"`
	MaxVersion              string `json:"max_version,omitempty"`
	MinVersion              string `json:"min_version,omitempty"`
}
