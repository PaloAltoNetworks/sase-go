package profiles

/*
Config ""

Parent chains:
* "Config"

Args:

Param Certificate (required, string): The Certificate param. String length must not exceed 255 characters.

Param ObjectId (string, read-only): The ObjectId param. Example: "abcd-1234".

Param Name (required, string): The Name param. String length must not exceed 127 characters.

Param ProtocolSettings (required, ProtocolSettingsObject): The ProtocolSettings param. ProtocolSettingsObject instance.
*/
type Config struct {
    Certificate string `json:"certificate"`
    ObjectId string `json:"id,omitempty"`
    Name string `json:"name"`
    ProtocolSettings ProtocolSettingsObject `json:"protocol_settings"`
}

/*
ProtocolSettingsObject ""

Parent chains:
* "Config protocol_settings"

Args:

Param AuthAlgoSha1 (bool): The AuthAlgoSha1 param.

Param AuthAlgoSha256 (bool): The AuthAlgoSha256 param.

Param AuthAlgoSha384 (bool): The AuthAlgoSha384 param.

Param EncAlgo3des (bool): The EncAlgo3des param.

Param EncAlgoAes128Cbc (bool): The EncAlgoAes128Cbc param.

Param EncAlgoAes128Gcm (bool): The EncAlgoAes128Gcm param.

Param EncAlgoAes256Cbc (bool): The EncAlgoAes256Cbc param.

Param EncAlgoAes256Gcm (bool): The EncAlgoAes256Gcm param.

Param EncAlgoRc4 (bool): The EncAlgoRc4 param.

Param KeyxchgAlgoDhe (bool): The KeyxchgAlgoDhe param.

Param KeyxchgAlgoEcdhe (bool): The KeyxchgAlgoEcdhe param.

Param KeyxchgAlgoRsa (bool): The KeyxchgAlgoRsa param.

Param MaxVersion (string): The MaxVersion param. String values: []string{"tls1-0", "tls1-1", "tls1-2", "tls1-3", "max"} Default: "max".

Param MinVersion (string): The MinVersion param. String values: []string{"tls1-0", "tls1-1", "tls1-2"} Default: "tls1-0".
*/
type ProtocolSettingsObject struct {
    AuthAlgoSha1 bool `json:"auth_algo_sha1,omitempty"`
    AuthAlgoSha256 bool `json:"auth_algo_sha256,omitempty"`
    AuthAlgoSha384 bool `json:"auth_algo_sha384,omitempty"`
    EncAlgo3des bool `json:"enc_algo_3des,omitempty"`
    EncAlgoAes128Cbc bool `json:"enc_algo_aes_128_cbc,omitempty"`
    EncAlgoAes128Gcm bool `json:"enc_algo_aes_128_gcm,omitempty"`
    EncAlgoAes256Cbc bool `json:"enc_algo_aes_256_cbc,omitempty"`
    EncAlgoAes256Gcm bool `json:"enc_algo_aes_256_gcm,omitempty"`
    EncAlgoRc4 bool `json:"enc_algo_rc4,omitempty"`
    KeyxchgAlgoDhe bool `json:"keyxchg_algo_dhe,omitempty"`
    KeyxchgAlgoEcdhe bool `json:"keyxchg_algo_ecdhe,omitempty"`
    KeyxchgAlgoRsa bool `json:"keyxchg_algo_rsa,omitempty"`
    MaxVersion string `json:"max_version,omitempty"`
    MinVersion string `json:"min_version,omitempty"`
}