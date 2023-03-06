package attribute

/*
Config ""

Parent chains:
* "Config"

Args:

Param Code (string, read-only): The Code param. Example: "19".

Param Status (string, read-only): The Status param. Example: "success".
*/
type Config struct {
	Code   string `json:"code,omitempty"`
	Status string `json:"status,omitempty"`
}
