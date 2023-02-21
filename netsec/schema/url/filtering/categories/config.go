package categories

/*
Config ""

Parent chains:
* "Config"

Args:

Param Type (string): The Type param.

Param Value (string): The Value param.
*/
type Config struct {
    Type string `json:"type,omitempty"`
    Value string `json:"value,omitempty"`
}