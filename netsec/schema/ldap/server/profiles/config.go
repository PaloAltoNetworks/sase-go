package profiles

/*
Config ""

Parent chains:
* "Config"

Args:

Param Base (string): The Base param. String length must not exceed 255 characters.

Param BindDn (string): The BindDn param. String length must not exceed 255 characters.

Param BindPassword (string): The BindPassword param. String length must not exceed 121 characters.

Param BindTimelimit (string): The BindTimelimit param.

Param ObjectId (string, read-only): The ObjectId param. Example: "abcd-1234".

Param LdapType (string): The LdapType param. String values: []string{"active-directory", "e-directory", "sun", "other"}

Param RetryInterval (int64): The RetryInterval param.

Param Server (required, list of ServerObject objects): The Server param.

Param Ssl (bool): The Ssl param.

Param Timelimit (int64): The Timelimit param.

Param VerifyServerCertificate (bool): The VerifyServerCertificate param.
*/
type Config struct {
	Base                    string         `json:"base,omitempty"`
	BindDn                  string         `json:"bind_dn,omitempty"`
	BindPassword            string         `json:"bind_password,omitempty"`
	BindTimelimit           string         `json:"bind_timelimit,omitempty"`
	ObjectId                string         `json:"id,omitempty"`
	LdapType                string         `json:"ldap_type,omitempty"`
	RetryInterval           int64          `json:"retry_interval,omitempty"`
	Server                  []ServerObject `json:"server"`
	Ssl                     bool           `json:"ssl,omitempty"`
	Timelimit               int64          `json:"timelimit,omitempty"`
	VerifyServerCertificate bool           `json:"verify_server_certificate,omitempty"`
}

/*
ServerObject ""

Parent chains:
* "Config server"

Args:

Param Address (string): The Address param.

Param Name (string): The Name param.

Param Port (int64): The Port param. Value must be between 1 and 65535.
*/
type ServerObject struct {
	Address string `json:"address,omitempty"`
	Name    string `json:"name,omitempty"`
	Port    int64  `json:"port,omitempty"`
}
