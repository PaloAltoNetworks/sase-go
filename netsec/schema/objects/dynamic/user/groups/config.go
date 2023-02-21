package groups

/*
Config ""

Parent chains:
* "Config"

Args:

Param Description (string): The Description param. String lengh must be between 0 and 1023 characters.

Param Filter (required, string): The Filter param. String length must not exceed 2047 characters.

Param ObjectId (string, read-only): The ObjectId param. Example: "abcd-1234".

Param Name (required, string): The Name param. String length must not exceed 63 characters.

Param Tag (list of strings): The Tag param. Array should contain at most 64 items.
*/
type Config struct {
    Description string `json:"description,omitempty"`
    Filter string `json:"filter"`
    ObjectId string `json:"id,omitempty"`
    Name string `json:"name"`
    Tag []string `json:"tag,omitempty"`
}