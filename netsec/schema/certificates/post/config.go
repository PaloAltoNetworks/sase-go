package post

/*
Config ""

Parent chains:
* "Config"

Args:

Param Algorithm (required, AlgorithmObject): The Algorithm param. AlgorithmObject instance.

Param AlternateEmail (list of strings): The AlternateEmail param.

Param CertificateName (required, string): The CertificateName param. String length must exceed 1 characters.

Param CommonName (required, string): The CommonName param. String length must exceed 1 characters.

Param CountryCode (string): The CountryCode param.

Param DayTillExpiration (int64): The DayTillExpiration param.

Param Department (list of strings): The Department param.

Param Digest (required, string): The Digest param. String values: []string{"sha1", "sha256", "sha384", "sha512", "md5"}

Param Email (string): The Email param. String length must not exceed 255 characters.

Param Hostname (list of strings): The Hostname param.

Param Ip (list of strings): The Ip param.

Param IsBlockPrivateKey (bool): The IsBlockPrivateKey param.

Param IsCertificateAuthority (bool): The IsCertificateAuthority param.

Param Locality (string): The Locality param. String length must not exceed 64 characters.

Param OcspResponderUrl (string): The OcspResponderUrl param. String length must not exceed 64 characters.

Param SignedBy (string): The SignedBy param. String length must not exceed 64 characters.

Param State (string): The State param. String length must not exceed 32 characters.
*/
type Config struct {
	Algorithm              AlgorithmObject `json:"algorithm"`
	AlternateEmail         []string        `json:"alternate_email,omitempty"`
	CertificateName        string          `json:"certificate_name"`
	CommonName             string          `json:"common_name"`
	CountryCode            string          `json:"country_code,omitempty"`
	DayTillExpiration      int64           `json:"day_till_expiration,omitempty"`
	Department             []string        `json:"department,omitempty"`
	Digest                 string          `json:"digest"`
	Email                  string          `json:"email,omitempty"`
	Hostname               []string        `json:"hostname,omitempty"`
	Ip                     []string        `json:"ip,omitempty"`
	IsBlockPrivateKey      bool            `json:"is_block_privateKey,omitempty"`
	IsCertificateAuthority bool            `json:"is_certificate_authority,omitempty"`
	Locality               string          `json:"locality,omitempty"`
	OcspResponderUrl       string          `json:"ocsp_responder_url,omitempty"`
	SignedBy               string          `json:"signed_by,omitempty"`
	State                  string          `json:"state,omitempty"`
}

/*
AlgorithmObject ""

Parent chains:
* "Config algorithm"

Args:

Param EcdsaNumberOfBits (int64): The EcdsaNumberOfBits param. One of these params should be specified:  ecdsa_number_of_bits or rsa_number_of_bits.

Param RsaNumberOfBits (int64): The RsaNumberOfBits param. One of these params should be specified:  ecdsa_number_of_bits or rsa_number_of_bits.
*/
type AlgorithmObject struct {
	EcdsaNumberOfBits int64 `json:"ecdsa_number_of_bits,omitempty"`
	RsaNumberOfBits   int64 `json:"rsa_number_of_bits,omitempty"`
}
