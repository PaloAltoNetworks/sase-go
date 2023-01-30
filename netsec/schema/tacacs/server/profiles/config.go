package profiles

/*
Config ""

Parent chains:
* "Config"

Args:

Param ObjectId (string, read-only): The ObjectId param. Example: "abcd-1234".

Param Protocol (required, string): The Protocol param. String values: []string{"CHAP", "PAP"}

Param Server (required, list of ServerObject objects): The Server param.

Param Timeout (int64): The Timeout param. Value must be between 1 and 30.

Param UseSingleConnection (bool): The UseSingleConnection param.
*/
type Config struct {
	ObjectId            string         `json:"id,omitempty"`
	Protocol            string         `json:"protocol"`
	Server              []ServerObject `json:"server"`
	Timeout             int64          `json:"timeout,omitempty"`
	UseSingleConnection bool           `json:"use_single_connection,omitempty"`
}

/*
ServerObject ""

Parent chains:
* "Config server"

Args:

Param Address (string): The Address param.

Param Name (string): The Name param.

Param Port (int64): The Port param. Value must be between 1 and 65535.

Param Secret (string): The Secret param. String length must not exceed 64 characters.
*/
type ServerObject struct {
	Address string `json:"address,omitempty"`
	Name    string `json:"name,omitempty"`
	Port    int64  `json:"port,omitempty"`
	Secret  string `json:"secret,omitempty"`
}
