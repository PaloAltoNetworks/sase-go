package rules

/*
Config ""

Parent chains:
* "Config"

Args:

Param Action (required, string): The Action param.

Param Application (required, list of strings): The Application param.

Param Category (required, list of strings): The Category param.

Param Description (string): The Description param.

Param Destination (required, list of strings): The Destination param.

Param DestinationHip (list of strings): The DestinationHip param.

Param Disabled (bool): The Disabled param.

Param From (required, list of strings): The From param.

Param ObjectId (string, read-only): The ObjectId param. Example: "abcd-1234".

Param LogSetting (string): The LogSetting param.

Param Name (required, string): The Name param.

Param NegateDestination (bool): The NegateDestination param.

Param NegateSource (bool): The NegateSource param.

Param ProfileSetting (ProfileSettingObject): The ProfileSetting param. ProfileSettingObject instance.

Param Service (required, list of strings): The Service param.

Param Source (required, list of strings): The Source param.

Param SourceHip (list of strings): The SourceHip param.

Param SourceUser (required, list of strings): The SourceUser param.

Param Tag (list of strings): The Tag param.

Param To (required, list of strings): The To param.
*/
type Config struct {
	Action            string                `json:"action"`
	Application       []string              `json:"application"`
	Category          []string              `json:"category"`
	Description       string                `json:"description,omitempty"`
	Destination       []string              `json:"destination"`
	DestinationHip    []string              `json:"destination_hip,omitempty"`
	Disabled          bool                  `json:"disabled,omitempty"`
	From              []string              `json:"from"`
	ObjectId          string                `json:"id,omitempty"`
	LogSetting        string                `json:"log_setting,omitempty"`
	Name              string                `json:"name"`
	NegateDestination bool                  `json:"negate_destination,omitempty"`
	NegateSource      bool                  `json:"negate_source,omitempty"`
	ProfileSetting    *ProfileSettingObject `json:"profile_setting,omitempty"`
	Service           []string              `json:"service"`
	Source            []string              `json:"source"`
	SourceHip         []string              `json:"source_hip,omitempty"`
	SourceUser        []string              `json:"source_user"`
	Tag               []string              `json:"tag,omitempty"`
	To                []string              `json:"to"`
}

/*
ProfileSettingObject ""

Parent chains:
* "Config profile_setting"

Args:

Param Group (list of strings): The Group param.
*/
type ProfileSettingObject struct {
	Group []string `json:"group,omitempty"`
}
