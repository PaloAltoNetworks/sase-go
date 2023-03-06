package profiles

/*
Config ""

Parent chains:
* "Config"

Args:

Param AggregateBandwidth (AggregateBandwidthObject): The AggregateBandwidth param. AggregateBandwidthObject instance.

Param ClassBandwidthType (ClassBandwidthTypeObject): The ClassBandwidthType param. ClassBandwidthTypeObject instance.

Param ObjectId (string, read-only): The ObjectId param. Example: "abcd-1234".

Param Name (required, string): The Name param. String length must not exceed 31 characters.
*/
type Config struct {
	AggregateBandwidth *AggregateBandwidthObject `json:"aggregate_bandwidth,omitempty"`
	ClassBandwidthType *ClassBandwidthTypeObject `json:"class_bandwidth_type,omitempty"`
	ObjectId           string                    `json:"id,omitempty"`
	Name               string                    `json:"name"`
}

/*
AggregateBandwidthObject ""

Parent chains:
* "Config aggregate_bandwidth"

Args:

Param EgressGuaranteed (int64): The EgressGuaranteed param. Value must be between 0 and 16000.

Param EgressMax (int64): The EgressMax param. Value must be between 0 and 60000.
*/
type AggregateBandwidthObject struct {
	EgressGuaranteed int64 `json:"egress_guaranteed,omitempty"`
	EgressMax        int64 `json:"egress_max,omitempty"`
}

/*
ClassBandwidthTypeObject ""

Parent chains:
* "Config class_bandwidth_type"

Args:

Param Mbps (MbpsObject): The Mbps param. MbpsObject instance.

Param Percentage (PercentageObject): The Percentage param. PercentageObject instance.
*/
type ClassBandwidthTypeObject struct {
	Mbps       *MbpsObject       `json:"mbps,omitempty"`
	Percentage *PercentageObject `json:"percentage,omitempty"`
}

/*
MbpsObject ""

Parent chains:
* "Config class_bandwidth_type mbps"

Args:

Param Class (list of ClassObject objects): The Class param.
*/
type MbpsObject struct {
	Class []ClassObject `json:"class,omitempty"`
}

/*
ClassObject ""

Parent chains:
* "Config class_bandwidth_type mbps class"
* "Config class_bandwidth_type percentage class"

Args:

Param ClassBandwidth (ClassBandwidthObject): The ClassBandwidth param. ClassBandwidthObject instance.

Param Name (string): The Name param. String length must not exceed 31 characters.

Param Priority (string): The Priority param. String values: []string{"real-time", "high", "medium", "low"} Default: "medium".
*/
type ClassObject struct {
	ClassBandwidth *ClassBandwidthObject `json:"class_bandwidth,omitempty"`
	Name           string                `json:"name,omitempty"`
	Priority       string                `json:"priority,omitempty"`
}

/*
ClassBandwidthObject ""

Parent chains:
* "Config class_bandwidth_type mbps class class_bandwidth"

Args:

Param EgressGuaranteed (int64): The EgressGuaranteed param. Value must be between 0 and 60000.

Param EgressMax (int64): The EgressMax param. Value must be between 0 and 60000.
*/
type ClassBandwidthObject struct {
	EgressGuaranteed int64 `json:"egress_guaranteed,omitempty"`
	EgressMax        int64 `json:"egress_max,omitempty"`
}

/*
PercentageObject ""

Parent chains:
* "Config class_bandwidth_type percentage"

Args:

Param Class (list of ClassObject objects): The Class param.
*/
type PercentageObject struct {
	Class []ClassObject `json:"class,omitempty"`
}
