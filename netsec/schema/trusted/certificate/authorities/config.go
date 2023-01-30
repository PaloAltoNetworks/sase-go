package authorities

/*
Config ""

Parent chains:
* "Config"

Args:

Param CommonName (string): The CommonName param. String length must not exceed 255 characters.

Param ExpiryEpoch (string): The ExpiryEpoch param.

Param Filename (string): The Filename param.

Param ObjectId (string, read-only): The ObjectId param. Example: "abcd-1234".

Param Issuer (string): The Issuer param.

Param Name (string): The Name param. String length must not exceed 63 characters.

Param NotValidAfter (string): The NotValidAfter param.

Param NotValidBefore (string): The NotValidBefore param.

Param SerialNumber (string): The SerialNumber param.

Param Subject (string): The Subject param.
*/
type Config struct {
	CommonName     string `json:"common_name,omitempty"`
	ExpiryEpoch    string `json:"expiry_epoch,omitempty"`
	Filename       string `json:"filename,omitempty"`
	ObjectId       string `json:"id,omitempty"`
	Issuer         string `json:"issuer,omitempty"`
	Name           string `json:"name,omitempty"`
	NotValidAfter  string `json:"not_valid_after,omitempty"`
	NotValidBefore string `json:"not_valid_before,omitempty"`
	SerialNumber   string `json:"serial_number,omitempty"`
	Subject        string `json:"subject,omitempty"`
}
