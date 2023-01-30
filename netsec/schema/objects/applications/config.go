package applications

/*
Config ""

Parent chains:
* "Config"

Args:

Param AbleToTransferFile (bool): The AbleToTransferFile param.

Param AlgDisableCapability (string): The AlgDisableCapability param. String length must not exceed 127 characters.

Param Category (required, string): The Category param.

Param ConsumeBigBandwidth (bool): The ConsumeBigBandwidth param.

Param DataIdent (bool): The DataIdent param.

Param Default (DefaultObject): The Default param. DefaultObject instance.

Param Description (string): The Description param. String lengh must be between 0 and 1023 characters.

Param EvasiveBehavior (bool): The EvasiveBehavior param.

Param FileTypeIdent (bool): The FileTypeIdent param.

Param HasKnownVulnerability (bool): The HasKnownVulnerability param.

Param ObjectId (string, read-only): The ObjectId param. Example: "abcd-1234".

Param Name (required, string): The Name param. String length must not exceed 31 characters.

Param NoAppidCaching (bool): The NoAppidCaching param.

Param ParentApp (string): The ParentApp param. String length must not exceed 127 characters.

Param PervasiveUse (bool): The PervasiveUse param.

Param ProneToMisuse (bool): The ProneToMisuse param.

Param Risk (required, int64): The Risk param. Value must be between 1 and 5.

Param Signature (list of SignatureObject objects): The Signature param.

Param Subcategory (required, string): The Subcategory param. String length must not exceed 63 characters.

Param TcpHalfClosedTimeout (int64): The TcpHalfClosedTimeout param. Value must be between 1 and 604800.

Param TcpTimeWaitTimeout (int64): The TcpTimeWaitTimeout param. Value must be between 1 and 600.

Param TcpTimeout (int64): The TcpTimeout param. Value must be between 0 and 604800.

Param Technology (required, string): The Technology param. String length must not exceed 63 characters.

Param Timeout (int64): The Timeout param. Value must be between 0 and 604800.

Param TunnelApplications (bool): The TunnelApplications param.

Param TunnelOtherApplication (bool): The TunnelOtherApplication param.

Param UdpTimeout (int64): The UdpTimeout param. Value must be between 0 and 604800.

Param UsedByMalware (bool): The UsedByMalware param.

Param VirusIdent (bool): The VirusIdent param.
*/
type Config struct {
	AbleToTransferFile     bool              `json:"able_to_transfer_file,omitempty"`
	AlgDisableCapability   string            `json:"alg_disable_capability,omitempty"`
	Category               string            `json:"category"`
	ConsumeBigBandwidth    bool              `json:"consume_big_bandwidth,omitempty"`
	DataIdent              bool              `json:"data_ident,omitempty"`
	Default                *DefaultObject    `json:"default,omitempty"`
	Description            string            `json:"description,omitempty"`
	EvasiveBehavior        bool              `json:"evasive_behavior,omitempty"`
	FileTypeIdent          bool              `json:"file_type_ident,omitempty"`
	HasKnownVulnerability  bool              `json:"has_known_vulnerability,omitempty"`
	ObjectId               string            `json:"id,omitempty"`
	Name                   string            `json:"name"`
	NoAppidCaching         bool              `json:"no_appid_caching,omitempty"`
	ParentApp              string            `json:"parent_app,omitempty"`
	PervasiveUse           bool              `json:"pervasive_use,omitempty"`
	ProneToMisuse          bool              `json:"prone_to_misuse,omitempty"`
	Risk                   int64             `json:"risk"`
	Signature              []SignatureObject `json:"signature,omitempty"`
	Subcategory            string            `json:"subcategory"`
	TcpHalfClosedTimeout   int64             `json:"tcp_half_closed_timeout,omitempty"`
	TcpTimeWaitTimeout     int64             `json:"tcp_time_wait_timeout,omitempty"`
	TcpTimeout             int64             `json:"tcp_timeout,omitempty"`
	Technology             string            `json:"technology"`
	Timeout                int64             `json:"timeout,omitempty"`
	TunnelApplications     bool              `json:"tunnel_applications,omitempty"`
	TunnelOtherApplication bool              `json:"tunnel_other_application,omitempty"`
	UdpTimeout             int64             `json:"udp_timeout,omitempty"`
	UsedByMalware          bool              `json:"used_by_malware,omitempty"`
	VirusIdent             bool              `json:"virus_ident,omitempty"`
}

/*
DefaultObject ""

Parent chains:
* "Config default"

Args:

Param IdentByIcmp6Type (IdentByIcmp6TypeObject): The IdentByIcmp6Type param. IdentByIcmp6TypeObject instance.

Param IdentByIcmpType (IdentByIcmpTypeObject): The IdentByIcmpType param. IdentByIcmpTypeObject instance.

Param IdentByIpProtocol (string): The IdentByIpProtocol param. Example: "0,1-255".

Param Port (list of strings): The Port param.
*/
type DefaultObject struct {
	IdentByIcmp6Type  *IdentByIcmp6TypeObject `json:"ident_by_icmp6_type,omitempty"`
	IdentByIcmpType   *IdentByIcmpTypeObject  `json:"ident_by_icmp_type,omitempty"`
	IdentByIpProtocol string                  `json:"ident_by_ip_protocol,omitempty"`
	Port              []string                `json:"port,omitempty"`
}

/*
IdentByIcmp6TypeObject ""

Parent chains:
* "Config default ident_by_icmp6_type"

Args:

Param Code (string): The Code param. Example: "0,1-255".

Param Type (required, string): The Type param. Example: "0,1-255".
*/
type IdentByIcmp6TypeObject struct {
	Code string `json:"code,omitempty"`
	Type string `json:"type"`
}

/*
IdentByIcmpTypeObject ""

Parent chains:
* "Config default ident_by_icmp_type"

Args:

Param Code (string): The Code param. Example: "0,1-255".

Param Type (required, string): The Type param. Example: "0,1-255".
*/
type IdentByIcmpTypeObject struct {
	Code string `json:"code,omitempty"`
	Type string `json:"type"`
}

/*
SignatureObject ""

Parent chains:
* "Config signature"

Args:

Param AndCondition (list of AndConditionObject objects): The AndCondition param.

Param Comment (string): The Comment param. String lengh must be between 0 and 256 characters.

Param Name (required, string): The Name param. String length must not exceed 31 characters.

Param OrderFree (bool): The OrderFree param. Default: false

Param Scope (string): The Scope param. String values: []string{"protocol-data-unit", "session"} Default: "protocol-data-unit".
*/
type SignatureObject struct {
	AndCondition []AndConditionObject `json:"and_condition,omitempty"`
	Comment      string               `json:"comment,omitempty"`
	Name         string               `json:"name"`
	OrderFree    bool                 `json:"order_free,omitempty"`
	Scope        string               `json:"scope,omitempty"`
}

/*
AndConditionObject ""

Parent chains:
* "Config signature and_condition"

Args:

Param Name (required, string): The Name param. String length must not exceed 31 characters.

Param OrCondition (list of OrConditionObject objects): The OrCondition param.
*/
type AndConditionObject struct {
	Name        string              `json:"name"`
	OrCondition []OrConditionObject `json:"or_condition,omitempty"`
}

/*
OrConditionObject ""

Parent chains:
* "Config signature and_condition or_condition"

Args:

Param Name (required, string): The Name param. String length must not exceed 31 characters.

Param Operator (required, OperatorObject): The Operator param. OperatorObject instance.
*/
type OrConditionObject struct {
	Name     string         `json:"name"`
	Operator OperatorObject `json:"operator"`
}

/*
OperatorObject ""

Parent chains:
* "Config signature and_condition or_condition operator"

Args:

Param EqualTo (EqualToObject): The EqualTo param. EqualToObject instance.

Param GreaterThan (GreaterThanObject): The GreaterThan param. GreaterThanObject instance.

Param LessThan (LessThanObject): The LessThan param. LessThanObject instance.

Param PatternMatch (PatternMatchObject): The PatternMatch param. PatternMatchObject instance.
*/
type OperatorObject struct {
	EqualTo      *EqualToObject      `json:"equal_to,omitempty"`
	GreaterThan  *GreaterThanObject  `json:"greater_than,omitempty"`
	LessThan     *LessThanObject     `json:"less_than,omitempty"`
	PatternMatch *PatternMatchObject `json:"pattern_match,omitempty"`
}

/*
EqualToObject ""

Parent chains:
* "Config signature and_condition or_condition operator equal_to"

Args:

Param Context (required, string): The Context param.

Param Mask (string): The Mask param. String length must not exceed 10 characters.

Param Position (string): The Position param. String length must not exceed 127 characters.

Param Value (required, string): The Value param. String length must not exceed 10 characters.
*/
type EqualToObject struct {
	Context  string `json:"context"`
	Mask     string `json:"mask,omitempty"`
	Position string `json:"position,omitempty"`
	Value    string `json:"value"`
}

/*
GreaterThanObject ""

Parent chains:
* "Config signature and_condition or_condition operator greater_than"

Args:

Param Context (required, string): The Context param. String length must not exceed 127 characters.

Param Qualifier (list of QualifierObject objects): The Qualifier param.

Param Value (required, int64): The Value param. Value must be between 0 and 4294967295.
*/
type GreaterThanObject struct {
	Context   string            `json:"context"`
	Qualifier []QualifierObject `json:"qualifier,omitempty"`
	Value     int64             `json:"value"`
}

/*
QualifierObject ""

Parent chains:
* "Config signature and_condition or_condition operator greater_than qualifier"
* "Config signature and_condition or_condition operator less_than qualifier"
* "Config signature and_condition or_condition operator pattern_match qualifier"

Args:

Param Name (required, string): The Name param. String length must not exceed 31 characters.

Param Value (required, string): The Value param.
*/
type QualifierObject struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

/*
LessThanObject ""

Parent chains:
* "Config signature and_condition or_condition operator less_than"

Args:

Param Context (required, string): The Context param. String length must not exceed 127 characters.

Param Qualifier (list of QualifierObject objects): The Qualifier param.

Param Value (required, int64): The Value param. Value must be between 0 and 4294967295.
*/
type LessThanObject struct {
	Context   string            `json:"context"`
	Qualifier []QualifierObject `json:"qualifier,omitempty"`
	Value     int64             `json:"value"`
}

/*
PatternMatchObject ""

Parent chains:
* "Config signature and_condition or_condition operator pattern_match"

Args:

Param Context (required, string): The Context param. String length must not exceed 127 characters.

Param Pattern (required, string): The Pattern param. String length must not exceed 127 characters.

Param Qualifier (list of QualifierObject objects): The Qualifier param.
*/
type PatternMatchObject struct {
	Context   string            `json:"context"`
	Pattern   string            `json:"pattern"`
	Qualifier []QualifierObject `json:"qualifier,omitempty"`
}
