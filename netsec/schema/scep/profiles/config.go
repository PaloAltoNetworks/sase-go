package profiles

/*
Config ""

Parent chains:
* "Config"

Args:

Param Algorithm (AlgorithmObject): The Algorithm param. AlgorithmObject instance.

Param CaIdentityName (required, string): The CaIdentityName param.

Param CertificateAttributes (CertificateAttributesObject): The CertificateAttributes param. CertificateAttributesObject instance.

Param Digest (required, string): The Digest param.

Param Fingerprint (string): The Fingerprint param.

Param ObjectId (string, read-only): The ObjectId param. Example: "abcd-1234".

Param Name (required, string): The Name param. String length must not exceed 31 characters.

Param ScepCaCert (string): The ScepCaCert param.

Param ScepChallenge (ScepChallengeObject): The ScepChallenge param. ScepChallengeObject instance.

Param ScepClientCert (string): The ScepClientCert param.

Param ScepUrl (required, string): The ScepUrl param.

Param Subject (string): The Subject param.

Param UseAsDigitalSignature (bool): The UseAsDigitalSignature param.

Param UseForKeyEncipherment (bool): The UseForKeyEncipherment param.
*/
type Config struct {
    Algorithm *AlgorithmObject `json:"algorithm,omitempty"`
    CaIdentityName string `json:"ca_identity_name"`
    CertificateAttributes *CertificateAttributesObject `json:"certificate_attributes,omitempty"`
    Digest string `json:"digest"`
    Fingerprint string `json:"fingerprint,omitempty"`
    ObjectId string `json:"id,omitempty"`
    Name string `json:"name"`
    ScepCaCert string `json:"scep_ca_cert,omitempty"`
    ScepChallenge *ScepChallengeObject `json:"scep_challenge,omitempty"`
    ScepClientCert string `json:"scep_client_cert,omitempty"`
    ScepUrl string `json:"scep_url"`
    Subject string `json:"subject,omitempty"`
    UseAsDigitalSignature bool `json:"use_as_digital_signature,omitempty"`
    UseForKeyEncipherment bool `json:"use_for_key_encipherment,omitempty"`
}

/*
AlgorithmObject ""

Parent chains:
* "Config algorithm"

Args:

Param Rsa (RsaObject): The Rsa param. RsaObject instance.
*/
type AlgorithmObject struct {
    Rsa *RsaObject `json:"rsa,omitempty"`
}

/*
RsaObject ""

Parent chains:
* "Config algorithm rsa"

Args:

Param RsaNbits (string): The RsaNbits param.
*/
type RsaObject struct {
    RsaNbits string `json:"rsa_nbits,omitempty"`
}

/*
CertificateAttributesObject ""

Parent chains:
* "Config certificate_attributes"

Args:

Param Dnsname (string): The Dnsname param.

Param Rfc822name (string): The Rfc822name param.

Param UniformResourceIdentifier (string): The UniformResourceIdentifier param.
*/
type CertificateAttributesObject struct {
    Dnsname string `json:"dnsname,omitempty"`
    Rfc822name string `json:"rfc822name,omitempty"`
    UniformResourceIdentifier string `json:"uniform_resource_identifier,omitempty"`
}

/*
ScepChallengeObject ""

Parent chains:
* "Config scep_challenge"

Args:

Param DynamicValue (DynamicObject): The DynamicValue param. DynamicObject instance.

Param Fixed (string): The Fixed param. String lengh must be between 0 and 1024 characters.

Param None (string): The None param. String values: []string{""} Default: "".
*/
type ScepChallengeObject struct {
    DynamicValue *DynamicObject `json:"dynamic,omitempty"`
    Fixed string `json:"fixed,omitempty"`
    None string `json:"none,omitempty"`
}

/*
DynamicObject ""

Parent chains:
* "Config scep_challenge dynamic"

Args:

Param OtpServerUrl (string): The OtpServerUrl param. String lengh must be between 0 and 255 characters.

Param Password (string): The Password param. String lengh must be between 0 and 255 characters.

Param Username (string): The Username param. String lengh must be between 0 and 255 characters.
*/
type DynamicObject struct {
    OtpServerUrl string `json:"otp_server_url,omitempty"`
    Password string `json:"password,omitempty"`
    Username string `json:"username,omitempty"`
}