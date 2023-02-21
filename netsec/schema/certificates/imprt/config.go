package imprt

/*
Config ""

Parent chains:
* "Config"

Args:

Param CertificateFile (required, string): The CertificateFile param.

Param Format (required, string): The Format param. String values: []string{"pem", "pkcs12", "der"} Default: "pem".

Param KeyFile (string): The KeyFile param.

Param Name (required, string): The Name param. String length must exceed 1 characters.

Param Passphrase (string): The Passphrase param.
*/
type Config struct {
    CertificateFile string `json:"certificate_file"`
    Format string `json:"format"`
    KeyFile string `json:"key_file,omitempty"`
    Name string `json:"name"`
    Passphrase string `json:"passphrase,omitempty"`
}