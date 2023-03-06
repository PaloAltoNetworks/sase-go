package rules

/*
Config ""

Parent chains:
* "Config"

Args:

Param AuthenticationEnforcement (string): The AuthenticationEnforcement param.

Param Category (list of strings): The Category param.

Param Description (string): The Description param.

Param Destination (list of strings): The Destination param.

Param DestinationHip (list of strings): The DestinationHip param.

Param Disabled (bool): The Disabled param. Default: false

Param From (list of strings): The From param.

Param GroupTag (string): The GroupTag param.

Param HipProfiles (list of strings): The HipProfiles param.

Param ObjectId (string, read-only): The ObjectId param. Example: "abcd-1234".

Param LogAuthenticationTimeout (bool): The LogAuthenticationTimeout param. Default: false

Param LogSetting (string): The LogSetting param.

Param Name (string): The Name param.

Param NegateDestination (bool): The NegateDestination param. Default: false

Param NegateSource (bool): The NegateSource param. Default: false

Param Service (list of strings): The Service param.

Param Source (list of strings): The Source param.

Param SourceHip (list of strings): The SourceHip param.

Param SourceUser (list of strings): The SourceUser param.

Param Tag (list of strings): The Tag param.

Param Timeout (int64): The Timeout param. Value must be between 1 and 1440.

Param To (list of strings): The To param.
*/
type Config struct {
	AuthenticationEnforcement string   `json:"authentication_enforcement,omitempty"`
	Category                  []string `json:"category,omitempty"`
	Description               string   `json:"description,omitempty"`
	Destination               []string `json:"destination,omitempty"`
	DestinationHip            []string `json:"destination_hip,omitempty"`
	Disabled                  bool     `json:"disabled,omitempty"`
	From                      []string `json:"from,omitempty"`
	GroupTag                  string   `json:"group_tag,omitempty"`
	HipProfiles               []string `json:"hip_profiles,omitempty"`
	ObjectId                  string   `json:"id,omitempty"`
	LogAuthenticationTimeout  bool     `json:"log_authentication_timeout,omitempty"`
	LogSetting                string   `json:"log_setting,omitempty"`
	Name                      string   `json:"name,omitempty"`
	NegateDestination         bool     `json:"negate_destination,omitempty"`
	NegateSource              bool     `json:"negate_source,omitempty"`
	Service                   []string `json:"service,omitempty"`
	Source                    []string `json:"source,omitempty"`
	SourceHip                 []string `json:"source_hip,omitempty"`
	SourceUser                []string `json:"source_user,omitempty"`
	Tag                       []string `json:"tag,omitempty"`
	Timeout                   int64    `json:"timeout,omitempty"`
	To                        []string `json:"to,omitempty"`
}
