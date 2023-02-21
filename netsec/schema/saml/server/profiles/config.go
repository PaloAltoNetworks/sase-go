package profiles

/*
Config ""

Parent chains:
* "Config"

Args:

Param Certificate (required, string): The Certificate param. String length must not exceed 63 characters.

Param EntityId (string): The EntityId param. String lengh must be between 1 and 1024 characters.

Param ObjectId (string, read-only): The ObjectId param. Example: "abcd-1234".

Param MaxClockSkew (int64): The MaxClockSkew param. Value must be between 1 and 900.

Param SloBindings (string): The SloBindings param. String values: []string{"post", "redirect"}

Param SsoBindings (string): The SsoBindings param. String values: []string{"post", "redirect"}

Param SsoUrl (string): The SsoUrl param. String lengh must be between 1 and 255 characters.

Param ValidateIdpCertificate (bool): The ValidateIdpCertificate param.

Param WantAuthRequestsSigned (bool): The WantAuthRequestsSigned param.
*/
type Config struct {
    Certificate string `json:"certificate"`
    EntityId string `json:"entity_id,omitempty"`
    ObjectId string `json:"id,omitempty"`
    MaxClockSkew int64 `json:"max_clock_skew,omitempty"`
    SloBindings string `json:"slo_bindings,omitempty"`
    SsoBindings string `json:"sso_bindings,omitempty"`
    SsoUrl string `json:"sso_url,omitempty"`
    ValidateIdpCertificate bool `json:"validate_idp_certificate,omitempty"`
    WantAuthRequestsSigned bool `json:"want_auth_requests_signed,omitempty"`
}