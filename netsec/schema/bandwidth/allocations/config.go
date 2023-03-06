package allocations

/*
Config ""

Parent chains:
* "Config"

Args:

Param AllocatedBandwidth (required, int64): The AllocatedBandwidth param.

Param ObjectId (string, read-only): The ObjectId param. Example: "abcd-1234".

Param Name (required, string): The Name param.

Param Qos (QosObject): The Qos param. QosObject instance.

Param SpnNameList (list of strings): The SpnNameList param.
*/
type Config struct {
	AllocatedBandwidth int64      `json:"allocated_bandwidth"`
	ObjectId           string     `json:"id,omitempty"`
	Name               string     `json:"name"`
	Qos                *QosObject `json:"qos,omitempty"`
	SpnNameList        []string   `json:"spn_name_list,omitempty"`
}

/*
QosObject ""

Parent chains:
* "Config qos"

Args:

Param Customized (bool): The Customized param.

Param Enabled (bool): The Enabled param.

Param GuaranteedRatio (int64): The GuaranteedRatio param.

Param Profile (string): The Profile param.
*/
type QosObject struct {
	Customized      bool   `json:"customized,omitempty"`
	Enabled         bool   `json:"enabled,omitempty"`
	GuaranteedRatio int64  `json:"guaranteed_ratio,omitempty"`
	Profile         string `json:"profile,omitempty"`
}
