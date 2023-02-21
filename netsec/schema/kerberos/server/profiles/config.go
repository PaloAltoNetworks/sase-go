package profiles

/*
Config ""

Parent chains:
* "Config"

Args:

Param ObjectId (string, read-only): The ObjectId param. Example: "abcd-1234".

Param Server (required, list of ServerObject objects): The Server param.
*/
type Config struct {
    ObjectId string `json:"id,omitempty"`
    Server []ServerObject `json:"server"`
}

/*
ServerObject ""

Parent chains:
* "Config server"

Args:

Param Host (string): The Host param.

Param Name (string): The Name param.

Param Port (int64): The Port param. Value must be between 1 and 65535.
*/
type ServerObject struct {
    Host string `json:"host,omitempty"`
    Name string `json:"name,omitempty"`
    Port int64 `json:"port,omitempty"`
}