package responder

/*
Config ""

Parent chains:
* "Config"

Args:

Param HostName (required, string): The HostName param. String lengh must be between 1 and 255 characters.

Param ObjectId (string, read-only): The ObjectId param. Example: "abcd-1234".

Param Name (required, string): The Name param. String length must not exceed 63 characters.
*/
type Config struct {
    HostName string `json:"host_name"`
    ObjectId string `json:"id,omitempty"`
    Name string `json:"name"`
}