package groups

/*
Config ""

Parent chains:
* "Config"

Args:

Param ObjectId (string, read-only): The ObjectId param. Example: "abcd-1234".

Param Members (required, list of strings): The Members param.

Param Name (required, string): The Name param. String length must not exceed 63 characters.

Param Tag (list of strings): The Tag param. Array should contain at most 64 items.
*/
type Config struct {
	ObjectId string   `json:"id,omitempty"`
	Members  []string `json:"members"`
	Name     string   `json:"name"`
	Tag      []string `json:"tag,omitempty"`
}
