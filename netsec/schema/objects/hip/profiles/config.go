package profiles

/*
Config ""

Parent chains:
* "Config"

Args:

Param Description (string): The Description param. String lengh must be between 0 and 255 characters.

Param ObjectId (string, read-only): The ObjectId param. Example: "abcd-1234".

Param Match (required, string): The Match param. String length must not exceed 2048 characters.

Param Name (required, string): The Name param. String length must not exceed 31 characters.
*/
type Config struct {
    Description string `json:"description,omitempty"`
    ObjectId string `json:"id,omitempty"`
    Match string `json:"match"`
    Name string `json:"name"`
}