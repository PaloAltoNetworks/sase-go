package versions

/*
Config ""

Parent chains:
* "Config"

Args:

Param Date (string): The Date param.

Param Device (string): The Device param.

Param Version (int64): The Version param.
*/
type Config struct {
    Date string `json:"date,omitempty"`
    Device string `json:"device,omitempty"`
    Version int64 `json:"version,omitempty"`
}