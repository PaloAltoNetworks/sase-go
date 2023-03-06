package versions

/*
Config ""

Parent chains:
* "Config"

Args:

Param Admin (string): The Admin param.

Param Created (int64): The Created param.

Param Date (string): The Date param.

Param Deleted (int64): The Deleted param.

Param Description (string): The Description param.

Param ObjectId (int64): The ObjectId param.

Param Scope (string): The Scope param.

Param SwgConfig (string): The SwgConfig param.

Param Updated (int64): The Updated param.

Param Version (string): The Version param.
*/
type Config struct {
	Admin       string `json:"admin,omitempty"`
	Created     int64  `json:"created,omitempty"`
	Date        string `json:"date,omitempty"`
	Deleted     int64  `json:"deleted,omitempty"`
	Description string `json:"description,omitempty"`
	ObjectId    int64  `json:"id,omitempty"`
	Scope       string `json:"scope,omitempty"`
	SwgConfig   string `json:"swg_config,omitempty"`
	Updated     int64  `json:"updated,omitempty"`
	Version     string `json:"version,omitempty"`
}
