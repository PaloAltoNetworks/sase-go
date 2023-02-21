package actions

/*
Config ""

Parent chains:
* "Config"

Args:

Param Actions (list of ActionsObject objects): The Actions param.

Param Description (string): The Description param. String lengh must be between 0 and 1023 characters.

Param Filter (required, string): The Filter param. String length must not exceed 2047 characters.

Param ObjectId (string, read-only): The ObjectId param. Example: "abcd-1234".

Param LogType (required, string, read-only): The LogType param. Example: "container".

Param Name (required, string): The Name param. String length must not exceed 63 characters.

Param Quarantine (bool): The Quarantine param.

Param SendToPanorama (bool): The SendToPanorama param.
*/
type Config struct {
    Actions []ActionsObject `json:"actions,omitempty"`
    Description string `json:"description,omitempty"`
    Filter string `json:"filter"`
    ObjectId string `json:"id,omitempty"`
    LogType string `json:"log_type"`
    Name string `json:"name"`
    Quarantine bool `json:"quarantine,omitempty"`
    SendToPanorama bool `json:"send_to_panorama,omitempty"`
}

/*
ActionsObject ""

Parent chains:
* "Config actions"

Args:

Param Name (required, string): The Name param.

Param Type (required, TypeObject): The Type param. TypeObject instance.
*/
type ActionsObject struct {
    Name string `json:"name"`
    Type TypeObject `json:"type"`
}

/*
TypeObject ""

Parent chains:
* "Config actions type"

Args:

Param Tagging (required, TaggingObject): The Tagging param. TaggingObject instance.
*/
type TypeObject struct {
    Tagging TaggingObject `json:"tagging"`
}

/*
TaggingObject ""

Parent chains:
* "Config actions type tagging"

Args:

Param Action (required, string): The Action param. String values: []string{"add-tag", "remove-tag"}

Param Tags (list of strings): The Tags param. Array should contain at most 64 items.

Param Target (required, string): The Target param.

Param Timeout (int64): The Timeout param.
*/
type TaggingObject struct {
    Action string `json:"action"`
    Tags []string `json:"tags,omitempty"`
    Target string `json:"target"`
    Timeout int64 `json:"timeout,omitempty"`
}