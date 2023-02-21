package profiles

/*
Config ""

Parent chains:
* "Config"

Args:

Param Alert (list of strings): The Alert param.

Param Allow (list of strings): The Allow param.

Param Block (list of strings): The Block param.

Param Continue (list of strings): The Continue param.

Param CredentialEnforcement (CredentialEnforcementObject): The CredentialEnforcement param. CredentialEnforcementObject instance.

Param Description (string): The Description param. String lengh must be between 0 and 255 characters.

Param ObjectId (string, read-only): The ObjectId param. Example: "abcd-1234".

Param LogContainerPageOnly (bool): The LogContainerPageOnly param. Default: true

Param LogHttpHdrReferer (bool): The LogHttpHdrReferer param. Default: false

Param LogHttpHdrUserAgent (bool): The LogHttpHdrUserAgent param. Default: false

Param LogHttpHdrXff (bool): The LogHttpHdrXff param. Default: false

Param MlavCategoryException (list of strings): The MlavCategoryException param.

Param MlavEngineUrlbasedEnabled (list of MlavEngineUrlbasedEnabledObject objects): The MlavEngineUrlbasedEnabled param.

Param Name (required, string): The Name param.

Param SafeSearchEnforcement (bool): The SafeSearchEnforcement param. Default: false
*/
type Config struct {
    Alert []string `json:"alert,omitempty"`
    Allow []string `json:"allow,omitempty"`
    Block []string `json:"block,omitempty"`
    Continue []string `json:"continue,omitempty"`
    CredentialEnforcement *CredentialEnforcementObject `json:"credential_enforcement,omitempty"`
    Description string `json:"description,omitempty"`
    ObjectId string `json:"id,omitempty"`
    LogContainerPageOnly bool `json:"log_container_page_only,omitempty"`
    LogHttpHdrReferer bool `json:"log_http_hdr_referer,omitempty"`
    LogHttpHdrUserAgent bool `json:"log_http_hdr_user_agent,omitempty"`
    LogHttpHdrXff bool `json:"log_http_hdr_xff,omitempty"`
    MlavCategoryException []string `json:"mlav_category_exception,omitempty"`
    MlavEngineUrlbasedEnabled []MlavEngineUrlbasedEnabledObject `json:"mlav_engine_urlbased_enabled,omitempty"`
    Name string `json:"name"`
    SafeSearchEnforcement bool `json:"safe_search_enforcement,omitempty"`
}

/*
CredentialEnforcementObject ""

Parent chains:
* "Config credential_enforcement"

Args:

Param Alert (list of strings): The Alert param.

Param Allow (list of strings): The Allow param.

Param Block (list of strings): The Block param.

Param Continue (list of strings): The Continue param.

Param LogSeverity (string): The LogSeverity param. Default: "medium".

Param Mode (ModeObject): The Mode param. ModeObject instance.
*/
type CredentialEnforcementObject struct {
    Alert []string `json:"alert,omitempty"`
    Allow []string `json:"allow,omitempty"`
    Block []string `json:"block,omitempty"`
    Continue []string `json:"continue,omitempty"`
    LogSeverity string `json:"log_severity,omitempty"`
    Mode *ModeObject `json:"mode,omitempty"`
}

/*
ModeObject ""

Parent chains:
* "Config credential_enforcement mode"

Args:

Param Disabled (interface{}): The Disabled param. interface{} instance.

Param DomainCredentials (interface{}): The DomainCredentials param. interface{} instance.

Param GroupMapping (string): The GroupMapping param.

Param IpUser (interface{}): The IpUser param. interface{} instance.
*/
type ModeObject struct {
    Disabled interface{} `json:"disabled,omitempty"`
    DomainCredentials interface{} `json:"domain_credentials,omitempty"`
    GroupMapping string `json:"group_mapping,omitempty"`
    IpUser interface{} `json:"ip_user,omitempty"`
}

/*
MlavEngineUrlbasedEnabledObject ""

Parent chains:
* "Config mlav_engine_urlbased_enabled"

Args:

Param MlavPolicyAction (string): The MlavPolicyAction param. String values: []string{"allow", "alert", "block"}

Param Name (string): The Name param.
*/
type MlavEngineUrlbasedEnabledObject struct {
    MlavPolicyAction string `json:"mlav_policy_action,omitempty"`
    Name string `json:"name,omitempty"`
}