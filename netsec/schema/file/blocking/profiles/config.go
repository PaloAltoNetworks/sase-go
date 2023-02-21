package profiles

/*
Config ""

Parent chains:
* "Config"

Args:

Param Description (string): The Description param.

Param ObjectId (string, read-only): The ObjectId param. Example: "abcd-1234".

Param Name (required, string): The Name param.

Param Rules (list of RulesObject objects): The Rules param.
*/
type Config struct {
    Description string `json:"description,omitempty"`
    ObjectId string `json:"id,omitempty"`
    Name string `json:"name"`
    Rules []RulesObject `json:"rules,omitempty"`
}

/*
RulesObject ""

Parent chains:
* "Config rules"

Args:

Param Action (required, string): The Action param. String values: []string{"alert", "block", "continue"} Default: "alert".

Param Application (required, list of strings): The Application param. Array should contain at least 1 items.

Param Direction (required, string): The Direction param. String values: []string{"download", "upload", "both"} Default: "both".

Param FileType (required, list of strings): The FileType param. Array should contain at least 1 items.

Param Name (required, string): The Name param.
*/
type RulesObject struct {
    Action string `json:"action"`
    Application []string `json:"application"`
    Direction string `json:"direction"`
    FileType []string `json:"file_type"`
    Name string `json:"name"`
}