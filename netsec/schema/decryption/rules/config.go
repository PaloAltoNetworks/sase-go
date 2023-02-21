package rules

/*
Config ""

Parent chains:
* "Config"

Args:

Param Action (required, string): The Action param. String values: []string{"decrypt", "no-decrypt"}

Param Category (required, list of strings): The Category param.

Param Description (string): The Description param.

Param Destination (required, list of strings): The Destination param.

Param DestinationHip (list of strings): The DestinationHip param.

Param Disabled (bool): The Disabled param.

Param From (required, list of strings): The From param.

Param ObjectId (string, read-only): The ObjectId param. Example: "abcd-1234".

Param LogFail (bool): The LogFail param.

Param LogSetting (string): The LogSetting param.

Param LogSuccess (bool): The LogSuccess param.

Param Name (required, string): The Name param.

Param NegateDestination (bool): The NegateDestination param.

Param NegateSource (bool): The NegateSource param.

Param Profile (string): The Profile param.

Param Service (required, list of strings): The Service param.

Param Source (required, list of strings): The Source param.

Param SourceHip (list of strings): The SourceHip param.

Param SourceUser (required, list of strings): The SourceUser param.

Param Tag (list of strings): The Tag param.

Param To (required, list of strings): The To param.

Param Type (TypeObject): The Type param. TypeObject instance.
*/
type Config struct {
    Action string `json:"action"`
    Category []string `json:"category"`
    Description string `json:"description,omitempty"`
    Destination []string `json:"destination"`
    DestinationHip []string `json:"destination_hip,omitempty"`
    Disabled bool `json:"disabled,omitempty"`
    From []string `json:"from"`
    ObjectId string `json:"id,omitempty"`
    LogFail bool `json:"log_fail,omitempty"`
    LogSetting string `json:"log_setting,omitempty"`
    LogSuccess bool `json:"log_success,omitempty"`
    Name string `json:"name"`
    NegateDestination bool `json:"negate_destination,omitempty"`
    NegateSource bool `json:"negate_source,omitempty"`
    Profile string `json:"profile,omitempty"`
    Service []string `json:"service"`
    Source []string `json:"source"`
    SourceHip []string `json:"source_hip,omitempty"`
    SourceUser []string `json:"source_user"`
    Tag []string `json:"tag,omitempty"`
    To []string `json:"to"`
    Type *TypeObject `json:"type,omitempty"`
}

/*
TypeObject ""

Parent chains:
* "Config type"

Args:

Param SslForwardProxy (interface{}): The SslForwardProxy param. interface{} instance.

Param SslInboundInspection (string): The SslInboundInspection param. One of these params should be specified:  ssl_inbound_inspection.
*/
type TypeObject struct {
    SslForwardProxy interface{} `json:"ssl_forward_proxy,omitempty"`
    SslInboundInspection string `json:"ssl_inbound_inspection,omitempty"`
}