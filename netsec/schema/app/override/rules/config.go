package rules

/*
Config ""

Parent chains:
* "Config"

Args:

Param Application (required, string): The Application param.

Param Description (string): The Description param. String length must not exceed 1024 characters.

Param Destination (required, list of strings): The Destination param.

Param Disabled (bool): The Disabled param. Default: false

Param From (required, list of strings): The From param.

Param GroupTag (string): The GroupTag param.

Param ObjectId (string, read-only): The ObjectId param. Example: "abcd-1234".

Param Name (required, string): The Name param. String length must not exceed 63 characters.

Param NegateDestination (bool): The NegateDestination param. Default: false

Param NegateSource (bool): The NegateSource param. Default: false

Param Port (required, int64): The Port param. Value must be between 0 and 65535.

Param Protocol (required, string): The Protocol param. String values: []string{"tcp", "udp"}

Param Source (required, list of strings): The Source param.

Param Tag (list of strings): The Tag param.

Param To (required, list of strings): The To param.
*/
type Config struct {
	Application       string   `json:"application"`
	Description       string   `json:"description,omitempty"`
	Destination       []string `json:"destination"`
	Disabled          bool     `json:"disabled,omitempty"`
	From              []string `json:"from"`
	GroupTag          string   `json:"group_tag,omitempty"`
	ObjectId          string   `json:"id,omitempty"`
	Name              string   `json:"name"`
	NegateDestination bool     `json:"negate_destination,omitempty"`
	NegateSource      bool     `json:"negate_source,omitempty"`
	Port              int64    `json:"port"`
	Protocol          string   `json:"protocol"`
	Source            []string `json:"source"`
	Tag               []string `json:"tag,omitempty"`
	To                []string `json:"to"`
}
