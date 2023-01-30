package get

/*
Config ""

Parent chains:
* "Config"

Args:

Param Algorithm (string): The Algorithm param.

Param Ca (bool): The Ca param.

Param CommonName (string): The CommonName param.

Param CommonNameInt (string): The CommonNameInt param.

Param ExpiryEpoch (string): The ExpiryEpoch param.

Param ObjectId (string, read-only): The ObjectId param. Example: "abcd-1234".

Param Issuer (string): The Issuer param.

Param IssuerHash (string): The IssuerHash param.

Param NotValidAfter (string): The NotValidAfter param.

Param NotValidBefore (string): The NotValidBefore param.

Param PublicKey (string): The PublicKey param.

Param Subject (string): The Subject param.

Param SubjectHash (string): The SubjectHash param.

Param SubjectInt (string): The SubjectInt param.
*/
type Config struct {
	Algorithm      string `json:"algorithm,omitempty"`
	Ca             bool   `json:"ca,omitempty"`
	CommonName     string `json:"common_name,omitempty"`
	CommonNameInt  string `json:"common_name_int,omitempty"`
	ExpiryEpoch    string `json:"expiry_epoch,omitempty"`
	ObjectId       string `json:"id,omitempty"`
	Issuer         string `json:"issuer,omitempty"`
	IssuerHash     string `json:"issuer_hash,omitempty"`
	NotValidAfter  string `json:"not_valid_after,omitempty"`
	NotValidBefore string `json:"not_valid_before,omitempty"`
	PublicKey      string `json:"public_key,omitempty"`
	Subject        string `json:"subject,omitempty"`
	SubjectHash    string `json:"subject_hash,omitempty"`
	SubjectInt     string `json:"subject_int,omitempty"`
}
