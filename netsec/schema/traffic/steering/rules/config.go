package rules

/*
Config ""

Parent chains:
* "Config"

Args:

Param Action (ActionObject): The Action param. ActionObject instance.

Param Category (list of strings): The Category param.

Param Destination (list of strings): The Destination param.

Param Name (required, string): The Name param.

Param Service (required, list of strings): The Service param.

Param Source (required, list of strings): The Source param.

Param SourceUser (list of strings): The SourceUser param.
*/
type Config struct {
    Action *ActionObject `json:"action,omitempty"`
    Category []string `json:"category,omitempty"`
    Destination []string `json:"destination,omitempty"`
    Name string `json:"name"`
    Service []string `json:"service"`
    Source []string `json:"source"`
    SourceUser []string `json:"source_user,omitempty"`
}

/*
ActionObject ""

Parent chains:
* "Config action"

Args:

Param Forward (ForwardObject): The Forward param. ForwardObject instance.

Param NoPbf (interface{}): The NoPbf param. interface{} instance.
*/
type ActionObject struct {
    Forward *ForwardObject `json:"forward,omitempty"`
    NoPbf interface{} `json:"no-pbf,omitempty"`
}

/*
ForwardObject ""

Parent chains:
* "Config action forward"

Args:

Param Target (string): The Target param.
*/
type ForwardObject struct {
    Target string `json:"target,omitempty"`
}