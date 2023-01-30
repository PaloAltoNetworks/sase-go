package attribute

/*
Config ""

Parent chains:
* "Config"

Args:

Param Count (string, read-only): The Count param. Example: "1".
*/
type Config struct {
	Count string `json:"count,omitempty"`
}
