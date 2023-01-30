package rules

/*
Config ""

Parent chains:
* "Config"

Args:

Param Action (required, ActionObject): The Action param. ActionObject instance.

Param Description (string): The Description param.

Param DscpTos (DscpTosObject): The DscpTos param. DscpTosObject instance.

Param ObjectId (string, read-only): The ObjectId param. Example: "abcd-1234".

Param Name (required, string): The Name param.

Param Schedule (string): The Schedule param.
*/
type Config struct {
	Action      ActionObject   `json:"action"`
	Description string         `json:"description,omitempty"`
	DscpTos     *DscpTosObject `json:"dscp_tos,omitempty"`
	ObjectId    string         `json:"id,omitempty"`
	Name        string         `json:"name"`
	Schedule    string         `json:"schedule,omitempty"`
}

/*
ActionObject ""

Parent chains:
* "Config action"

Args:

Param Class (string): The Class param.
*/
type ActionObject struct {
	Class string `json:"class,omitempty"`
}

/*
DscpTosObject ""

Parent chains:
* "Config dscp_tos"

Args:

Param Codepoints (list of CodepointsObject objects): The Codepoints param.
*/
type DscpTosObject struct {
	Codepoints []CodepointsObject `json:"codepoints,omitempty"`
}

/*
CodepointsObject ""

Parent chains:
* "Config dscp_tos codepoints"

Args:

Param Name (string): The Name param.

Param Type (TypeObject): The Type param. TypeObject instance.
*/
type CodepointsObject struct {
	Name string      `json:"name,omitempty"`
	Type *TypeObject `json:"type,omitempty"`
}

/*
TypeObject ""

Parent chains:
* "Config dscp_tos codepoints type"

Args:

Param Af (AfObject): The Af param. AfObject instance.

Param Cs (CsObject): The Cs param. CsObject instance.

Param Custom (CustomObject): The Custom param. CustomObject instance.

Param Ef (interface{}): The Ef param. interface{} instance.

Param Tos (TosObject): The Tos param. TosObject instance.
*/
type TypeObject struct {
	Af     *AfObject     `json:"af,omitempty"`
	Cs     *CsObject     `json:"cs,omitempty"`
	Custom *CustomObject `json:"custom,omitempty"`
	Ef     interface{}   `json:"ef,omitempty"`
	Tos    *TosObject    `json:"tos,omitempty"`
}

/*
AfObject ""

Parent chains:
* "Config dscp_tos codepoints type af"

Args:

Param Codepoint (string): The Codepoint param.
*/
type AfObject struct {
	Codepoint string `json:"codepoint,omitempty"`
}

/*
CsObject ""

Parent chains:
* "Config dscp_tos codepoints type cs"

Args:

Param Codepoint (string): The Codepoint param.
*/
type CsObject struct {
	Codepoint string `json:"codepoint,omitempty"`
}

/*
CustomObject ""

Parent chains:
* "Config dscp_tos codepoints type custom"

Args:

Param Codepoint (CodepointObject): The Codepoint param. CodepointObject instance.
*/
type CustomObject struct {
	Codepoint *CodepointObject `json:"codepoint,omitempty"`
}

/*
CodepointObject ""

Parent chains:
* "Config dscp_tos codepoints type custom codepoint"

Args:

Param BinaryValue (string): The BinaryValue param.

Param CodepointName (string): The CodepointName param.
*/
type CodepointObject struct {
	BinaryValue   string `json:"binary_value,omitempty"`
	CodepointName string `json:"codepoint_name,omitempty"`
}

/*
TosObject ""

Parent chains:
* "Config dscp_tos codepoints type tos"

Args:

Param Codepoint (string): The Codepoint param.
*/
type TosObject struct {
	Codepoint string `json:"codepoint,omitempty"`
}
