package signatures

/*
Config ""

Parent chains:
* "Config"

Args:

Param Bugtraq (list of strings): The Bugtraq param.

Param Comment (string): The Comment param. String lengh must be between 0 and 256 characters.

Param Cve (list of strings): The Cve param.

Param DefaultAction (DefaultActionObject): The DefaultAction param. DefaultActionObject instance.

Param Direction (string): The Direction param. String values: []string{"client2server", "server2client", "both"}

Param ObjectId (string, read-only): The ObjectId param. Example: "abcd-1234".

Param Reference (list of strings): The Reference param.

Param Severity (string): The Severity param. String values: []string{"critical", "low", "high", "medium", "informational"}

Param Signature (SignatureObject): The Signature param. SignatureObject instance.

Param ThreatId (required, int64): The ThreatId param. Value must be between 15000 and 70000000.

Param Threatname (required, string): The Threatname param. String length must not exceed 1024 characters.

Param Vendor (list of strings): The Vendor param.
*/
type Config struct {
    Bugtraq []string `json:"bugtraq,omitempty"`
    Comment string `json:"comment,omitempty"`
    Cve []string `json:"cve,omitempty"`
    DefaultAction *DefaultActionObject `json:"default_action,omitempty"`
    Direction string `json:"direction,omitempty"`
    ObjectId string `json:"id,omitempty"`
    Reference []string `json:"reference,omitempty"`
    Severity string `json:"severity,omitempty"`
    Signature *SignatureObject `json:"signature,omitempty"`
    ThreatId int64 `json:"threat_id"`
    Threatname string `json:"threatname"`
    Vendor []string `json:"vendor,omitempty"`
}

/*
DefaultActionObject ""

Parent chains:
* "Config default_action"

Args:

Param Alert (interface{}): The Alert param. interface{} instance.

Param Allow (interface{}): The Allow param. interface{} instance.

Param BlockIp (BlockIpObject): The BlockIp param. BlockIpObject instance.

Param Drop (interface{}): The Drop param. interface{} instance.

Param ResetBoth (interface{}): The ResetBoth param. interface{} instance.

Param ResetClient (interface{}): The ResetClient param. interface{} instance.

Param ResetServer (interface{}): The ResetServer param. interface{} instance.
*/
type DefaultActionObject struct {
    Alert interface{} `json:"alert,omitempty"`
    Allow interface{} `json:"allow,omitempty"`
    BlockIp *BlockIpObject `json:"block_ip,omitempty"`
    Drop interface{} `json:"drop,omitempty"`
    ResetBoth interface{} `json:"reset_both,omitempty"`
    ResetClient interface{} `json:"reset_client,omitempty"`
    ResetServer interface{} `json:"reset_server,omitempty"`
}

/*
BlockIpObject ""

Parent chains:
* "Config default_action block_ip"

Args:

Param Duration (int64): The Duration param. Value must be between 1 and 3600.

Param TrackBy (string): The TrackBy param. String values: []string{"source-and-destination", "source"}
*/
type BlockIpObject struct {
    Duration int64 `json:"duration,omitempty"`
    TrackBy string `json:"track_by,omitempty"`
}

/*
SignatureObject ""

Parent chains:
* "Config signature"

Args:

Param Combination (CombinationObject): The Combination param. CombinationObject instance.

Param Standard (list of StandardObject objects): The Standard param.
*/
type SignatureObject struct {
    Combination *CombinationObject `json:"combination,omitempty"`
    Standard []StandardObject `json:"standard,omitempty"`
}

/*
CombinationObject ""

Parent chains:
* "Config signature combination"

Args:

Param AndCondition (list of AndConditionObject objects): The AndCondition param.

Param OrderFree (bool): The OrderFree param. Default: false

Param TimeAttribute (TimeAttributeObject): The TimeAttribute param. TimeAttributeObject instance.
*/
type CombinationObject struct {
    AndCondition []AndConditionObject `json:"and_condition,omitempty"`
    OrderFree bool `json:"order_free,omitempty"`
    TimeAttribute *TimeAttributeObject `json:"time_attribute,omitempty"`
}

/*
AndConditionObject ""

Parent chains:
* "Config signature combination and_condition"
* "Config signature standard and_condition"

Args:

Param Name (string): The Name param.

Param OrCondition (list of OrConditionObject objects): The OrCondition param.
*/
type AndConditionObject struct {
    Name string `json:"name,omitempty"`
    OrCondition []OrConditionObject `json:"or_condition,omitempty"`
}

/*
OrConditionObject ""

Parent chains:
* "Config signature combination and_condition or_condition"

Args:

Param Name (string): The Name param.

Param ThreatId (string): The ThreatId param.
*/
type OrConditionObject struct {
    Name string `json:"name,omitempty"`
    ThreatId string `json:"threat_id,omitempty"`
}

/*
TimeAttributeObject ""

Parent chains:
* "Config signature combination time_attribute"

Args:

Param Interval (int64): The Interval param. Value must be between 1 and 3600.

Param Threshold (int64): The Threshold param. Value must be between 1 and 255.

Param TrackBy (string): The TrackBy param. String values: []string{"source-and-destination", "source", "destination"}
*/
type TimeAttributeObject struct {
    Interval int64 `json:"interval,omitempty"`
    Threshold int64 `json:"threshold,omitempty"`
    TrackBy string `json:"track_by,omitempty"`
}

/*
StandardObject ""

Parent chains:
* "Config signature standard"

Args:

Param AndCondition (list of AndConditionObject objects): The AndCondition param.

Param Comment (string): The Comment param. String lengh must be between 0 and 256 characters.

Param Name (required, string): The Name param.

Param OrderFree (bool): The OrderFree param. Default: false

Param Scope (string): The Scope param. String values: []string{"protocol-data-unit", "session"}
*/
type StandardObject struct {
    AndCondition []AndConditionObject `json:"and_condition,omitempty"`
    Comment string `json:"comment,omitempty"`
    Name string `json:"name"`
    OrderFree bool `json:"order_free,omitempty"`
    Scope string `json:"scope,omitempty"`
}