package users

/*
Config ""

Parent chains:
* "Config"

Args:

Param ObjectId (string, read-only): The ObjectId param. Example: "abcd-1234".

Param Name (required, string): The Name param. String length must not exceed 31 characters.

Param Password (required, string): The Password param. String length must not exceed 63 characters.
*/
type Config struct {
    ObjectId string `json:"id,omitempty"`
    Name string `json:"name"`
    Password string `json:"password"`
}