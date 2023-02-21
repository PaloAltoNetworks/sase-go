package exclusions

/*
Config ""

Parent chains:
* "Config"

Args:

Param Description (string): The Description param.

Param ObjectId (string, read-only): The ObjectId param. Example: "abcd-1234".

Param Name (required, string): The Name param.
*/
type Config struct {
    Description string `json:"description,omitempty"`
    ObjectId string `json:"id,omitempty"`
    Name string `json:"name"`
}