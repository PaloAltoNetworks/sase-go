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

Param ThreatException (list of ThreatExceptionObject objects): The ThreatException param.
*/
type Config struct {
    Description string `json:"description,omitempty"`
    ObjectId string `json:"id,omitempty"`
    Name string `json:"name"`
    Rules []RulesObject `json:"rules,omitempty"`
    ThreatException []ThreatExceptionObject `json:"threat_exception,omitempty"`
}

/*
RulesObject ""

Parent chains:
* "Config rules"

Args:

Param Action (ActionObject): The Action param. ActionObject instance.

Param Category (string): The Category param. String values: []string{"dns-proxy", "backdoor", "data-theft", "autogen", "spyware", "dns-security", "downloader", "dns-phishing", "phishing-kit", "cryptominer", "hacktool", "dns-benign", "dns-wildfire", "botnet", "dns-grayware", "inline-cloud-c2", "keylogger", "p2p-communication", "domain-edl", "webshell", "command-and-control", "dns-ddns", "net-worm", "any", "tls-fingerprint", "dns-new-domain", "dns", "fraud", "dns-c2", "adware", "post-exploitation", "dns-malware", "browser-hijack", "dns-parked"}

Param Name (string): The Name param.

Param PacketCapture (string): The PacketCapture param. String values: []string{"disable", "single-packet", "extended-capture"}

Param Severity (list of strings): The Severity param.

Param ThreatName (string): The ThreatName param. String length must exceed 4 characters.
*/
type RulesObject struct {
    Action *ActionObject `json:"action,omitempty"`
    Category string `json:"category,omitempty"`
    Name string `json:"name,omitempty"`
    PacketCapture string `json:"packet_capture,omitempty"`
    Severity []string `json:"severity,omitempty"`
    ThreatName string `json:"threat_name,omitempty"`
}

/*
ActionObject ""

Parent chains:
* "Config rules action"

Args:

Param Alert (interface{}): The Alert param. interface{} instance.

Param Allow (interface{}): The Allow param. interface{} instance.

Param BlockIp (BlockIpObject): The BlockIp param. BlockIpObject instance.

Param Drop (interface{}): The Drop param. interface{} instance.

Param ResetBoth (interface{}): The ResetBoth param. interface{} instance.

Param ResetClient (interface{}): The ResetClient param. interface{} instance.

Param ResetServer (interface{}): The ResetServer param. interface{} instance.
*/
type ActionObject struct {
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
* "Config rules action block_ip"
* "Config threat_exception action block_ip"

Args:

Param Duration (int64): The Duration param. Value must be between 1 and 3600.

Param TrackBy (string): The TrackBy param. String values: []string{"source-and-destination", "source"}
*/
type BlockIpObject struct {
    Duration int64 `json:"duration,omitempty"`
    TrackBy string `json:"track_by,omitempty"`
}

/*
ThreatExceptionObject ""

Parent chains:
* "Config threat_exception"

Args:

Param Action (ActionObject1): The Action param. ActionObject1 instance.

Param ExemptIp (list of ExemptIpObject objects): The ExemptIp param.

Param Name (string): The Name param.

Param Notes (string): The Notes param.

Param PacketCapture (string): The PacketCapture param. String values: []string{"disable", "single-packet", "extended-capture"}
*/
type ThreatExceptionObject struct {
    Action *ActionObject1 `json:"action,omitempty"`
    ExemptIp []ExemptIpObject `json:"exempt_ip,omitempty"`
    Name string `json:"name,omitempty"`
    Notes string `json:"notes,omitempty"`
    PacketCapture string `json:"packet_capture,omitempty"`
}

/*
ActionObject1 ""

Parent chains:
* "Config threat_exception action"

Args:

Param Alert (interface{}): The Alert param. interface{} instance.

Param Allow (interface{}): The Allow param. interface{} instance.

Param BlockIp (BlockIpObject): The BlockIp param. BlockIpObject instance.

Param Default (interface{}): The Default param. interface{} instance.

Param Drop (interface{}): The Drop param. interface{} instance.

Param ResetBoth (interface{}): The ResetBoth param. interface{} instance.

Param ResetClient (interface{}): The ResetClient param. interface{} instance.

Param ResetServer (interface{}): The ResetServer param. interface{} instance.
*/
type ActionObject1 struct {
    Alert interface{} `json:"alert,omitempty"`
    Allow interface{} `json:"allow,omitempty"`
    BlockIp *BlockIpObject `json:"block_ip,omitempty"`
    Default interface{} `json:"default,omitempty"`
    Drop interface{} `json:"drop,omitempty"`
    ResetBoth interface{} `json:"reset_both,omitempty"`
    ResetClient interface{} `json:"reset_client,omitempty"`
    ResetServer interface{} `json:"reset_server,omitempty"`
}

/*
ExemptIpObject ""

Parent chains:
* "Config threat_exception exempt_ip"

Args:

Param Name (required, string): The Name param.
*/
type ExemptIpObject struct {
    Name string `json:"name"`
}