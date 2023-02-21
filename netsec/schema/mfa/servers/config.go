package servers

/*
Config ""

Parent chains:
* "Config"

Args:

Param ObjectId (string, read-only): The ObjectId param. Example: "abcd-1234".

Param MfaCertProfile (required, string): The MfaCertProfile param.

Param MfaVendorType (MfaVendorTypeObject): The MfaVendorType param. MfaVendorTypeObject instance.

Param Name (required, string): The Name param.
*/
type Config struct {
    ObjectId string `json:"id,omitempty"`
    MfaCertProfile string `json:"mfa_cert_profile"`
    MfaVendorType *MfaVendorTypeObject `json:"mfa_vendor_type,omitempty"`
    Name string `json:"name"`
}

/*
MfaVendorTypeObject ""

Parent chains:
* "Config mfa_vendor_type"

Args:

Param DuoSecurityV2 (DuoSecurityV2Object): The DuoSecurityV2 param. DuoSecurityV2Object instance.

Param OktaAdaptiveV1 (OktaAdaptiveV1Object): The OktaAdaptiveV1 param. OktaAdaptiveV1Object instance.

Param PingIdentityV1 (PingIdentityV1Object): The PingIdentityV1 param. PingIdentityV1Object instance.

Param RsaSecuridAccessV1 (RsaSecuridAccessV1Object): The RsaSecuridAccessV1 param. RsaSecuridAccessV1Object instance.
*/
type MfaVendorTypeObject struct {
    DuoSecurityV2 *DuoSecurityV2Object `json:"duo_security_v2,omitempty"`
    OktaAdaptiveV1 *OktaAdaptiveV1Object `json:"okta_adaptive_v1,omitempty"`
    PingIdentityV1 *PingIdentityV1Object `json:"ping_identity_v1,omitempty"`
    RsaSecuridAccessV1 *RsaSecuridAccessV1Object `json:"rsa_securid_access_v1,omitempty"`
}

/*
DuoSecurityV2Object ""

Parent chains:
* "Config mfa_vendor_type duo_security_v2"

Args:

Param DuoApiHost (string): The DuoApiHost param.

Param DuoBaseuri (string): The DuoBaseuri param.

Param DuoIntegrationKey (string): The DuoIntegrationKey param.

Param DuoSecretKey (string): The DuoSecretKey param.

Param DuoTimeout (string): The DuoTimeout param.
*/
type DuoSecurityV2Object struct {
    DuoApiHost string `json:"duo_api_host,omitempty"`
    DuoBaseuri string `json:"duo_baseuri,omitempty"`
    DuoIntegrationKey string `json:"duo_integration_key,omitempty"`
    DuoSecretKey string `json:"duo_secret_key,omitempty"`
    DuoTimeout string `json:"duo_timeout,omitempty"`
}

/*
OktaAdaptiveV1Object ""

Parent chains:
* "Config mfa_vendor_type okta_adaptive_v1"

Args:

Param OktaApiHost (string): The OktaApiHost param.

Param OktaBaseuri (string): The OktaBaseuri param.

Param OktaOrg (string): The OktaOrg param.

Param OktaTimeout (string): The OktaTimeout param.

Param OktaToken (string): The OktaToken param.
*/
type OktaAdaptiveV1Object struct {
    OktaApiHost string `json:"okta_api_host,omitempty"`
    OktaBaseuri string `json:"okta_baseuri,omitempty"`
    OktaOrg string `json:"okta_org,omitempty"`
    OktaTimeout string `json:"okta_timeout,omitempty"`
    OktaToken string `json:"okta_token,omitempty"`
}

/*
PingIdentityV1Object ""

Parent chains:
* "Config mfa_vendor_type ping_identity_v1"

Args:

Param PingApiHost (string): The PingApiHost param.

Param PingBaseuri (string): The PingBaseuri param.

Param PingOrg (string): The PingOrg param.

Param PingOrgAlias (string): The PingOrgAlias param.

Param PingTimeout (string): The PingTimeout param.

Param PingToken (string): The PingToken param.
*/
type PingIdentityV1Object struct {
    PingApiHost string `json:"ping_api_host,omitempty"`
    PingBaseuri string `json:"ping_baseuri,omitempty"`
    PingOrg string `json:"ping_org,omitempty"`
    PingOrgAlias string `json:"ping_org_alias,omitempty"`
    PingTimeout string `json:"ping_timeout,omitempty"`
    PingToken string `json:"ping_token,omitempty"`
}

/*
RsaSecuridAccessV1Object ""

Parent chains:
* "Config mfa_vendor_type rsa_securid_access_v1"

Args:

Param RsaAccessid (string): The RsaAccessid param.

Param RsaAccesskey (string): The RsaAccesskey param.

Param RsaApiHost (string): The RsaApiHost param.

Param RsaAssurancepolicyid (string): The RsaAssurancepolicyid param.

Param RsaBaseuri (string): The RsaBaseuri param.

Param RsaTimeout (string): The RsaTimeout param.
*/
type RsaSecuridAccessV1Object struct {
    RsaAccessid string `json:"rsa_accessid,omitempty"`
    RsaAccesskey string `json:"rsa_accesskey,omitempty"`
    RsaApiHost string `json:"rsa_api_host,omitempty"`
    RsaAssurancepolicyid string `json:"rsa_assurancepolicyid,omitempty"`
    RsaBaseuri string `json:"rsa_baseuri,omitempty"`
    RsaTimeout string `json:"rsa_timeout,omitempty"`
}