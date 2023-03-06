package profiles

/*
Config ""

Parent chains:
* "Config"

Args:

Param BlockExpiredCert (bool): The BlockExpiredCert param.

Param BlockTimeoutCert (bool): The BlockTimeoutCert param.

Param BlockUnauthenticatedCert (bool): The BlockUnauthenticatedCert param.

Param BlockUnknownCert (bool): The BlockUnknownCert param.

Param CaCertificates (required, list of CaCertificatesObject objects): The CaCertificates param.

Param CertStatusTimeout (string): The CertStatusTimeout param.

Param CrlReceiveTimeout (string): The CrlReceiveTimeout param.

Param Domain (string): The Domain param.

Param ObjectId (string, read-only): The ObjectId param. Example: "abcd-1234".

Param Name (required, string): The Name param. String length must not exceed 63 characters.

Param OcspReceiveTimeout (string): The OcspReceiveTimeout param.

Param UseCrl (bool): The UseCrl param.

Param UseOcsp (bool): The UseOcsp param.

Param UsernameField (UsernameFieldObject): The UsernameField param. UsernameFieldObject instance.
*/
type Config struct {
	BlockExpiredCert         bool                   `json:"block_expired_cert,omitempty"`
	BlockTimeoutCert         bool                   `json:"block_timeout_cert,omitempty"`
	BlockUnauthenticatedCert bool                   `json:"block_unauthenticated_cert,omitempty"`
	BlockUnknownCert         bool                   `json:"block_unknown_cert,omitempty"`
	CaCertificates           []CaCertificatesObject `json:"ca_certificates"`
	CertStatusTimeout        string                 `json:"cert_status_timeout,omitempty"`
	CrlReceiveTimeout        string                 `json:"crl_receive_timeout,omitempty"`
	Domain                   string                 `json:"domain,omitempty"`
	ObjectId                 string                 `json:"id,omitempty"`
	Name                     string                 `json:"name"`
	OcspReceiveTimeout       string                 `json:"ocsp_receive_timeout,omitempty"`
	UseCrl                   bool                   `json:"use_crl,omitempty"`
	UseOcsp                  bool                   `json:"use_ocsp,omitempty"`
	UsernameField            *UsernameFieldObject   `json:"username_field,omitempty"`
}

/*
CaCertificatesObject ""

Parent chains:
* "Config ca_certificates"

Args:

Param DefaultOcspUrl (string): The DefaultOcspUrl param.

Param Name (string): The Name param.

Param OcspVerifyCert (string): The OcspVerifyCert param.

Param TemplateName (string): The TemplateName param.
*/
type CaCertificatesObject struct {
	DefaultOcspUrl string `json:"default_ocsp_url,omitempty"`
	Name           string `json:"name,omitempty"`
	OcspVerifyCert string `json:"ocsp_verify_cert,omitempty"`
	TemplateName   string `json:"template_name,omitempty"`
}

/*
UsernameFieldObject ""

Parent chains:
* "Config username_field"

Args:

Param Subject (string): The Subject param. String values: []string{"common-name"}

Param SubjectAlt (string): The SubjectAlt param. String values: []string{"email"}
*/
type UsernameFieldObject struct {
	Subject    string `json:"subject,omitempty"`
	SubjectAlt string `json:"subject_alt,omitempty"`
}
