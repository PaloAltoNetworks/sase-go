package groups

/*
Config ""

Parent chains:
* "Config"

Args:

Param DisableSnat (bool): The DisableSnat param.

Param Name (required, string): The Name param.

Param PbfOnly (bool): The PbfOnly param.

Param Target (required, list of strings): The Target param.
*/
type Config struct {
	DisableSnat bool     `json:"disable_snat,omitempty"`
	Name        string   `json:"name"`
	PbfOnly     bool     `json:"pbf_only,omitempty"`
	Target      []string `json:"target"`
}
