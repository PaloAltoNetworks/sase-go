package groups

/*
Config ""

Parent chains:
* "Config"

Args:

Param Description (string): The Description param. String lengh must be between 0 and 1023 characters.

Param DynamicValue (DynamicObject): The DynamicValue param. DynamicObject instance. One of these params should be specified:  dynamic or static.

Param ObjectId (string, read-only): The ObjectId param. Example: "abcd-1234".

Param Name (required, string): The Name param. String length must not exceed 63 characters.

Param Static (list of strings): The Static param. One of these params should be specified:  dynamic or static.

Param Tag (list of strings): The Tag param. Array should contain at most 64 items.
*/
type Config struct {
	Description  string         `json:"description,omitempty"`
	DynamicValue *DynamicObject `json:"dynamic,omitempty"`
	ObjectId     string         `json:"id,omitempty"`
	Name         string         `json:"name"`
	Static       []string       `json:"static,omitempty"`
	Tag          []string       `json:"tag,omitempty"`
}

/*
DynamicObject ""

Parent chains:
* "Config dynamic"

Args:

Param Filter (required, string): The Filter param. String length must not exceed 2047 characters.
*/
type DynamicObject struct {
	Filter string `json:"filter"`
}
