package profiles

/*
Config ""

Parent chains:
* "Config"

Args:

Param AllowList (list of strings): The AllowList param.

Param ObjectId (string, read-only): The ObjectId param. Example: "abcd-1234".

Param Lockout (LockoutObject): The Lockout param. LockoutObject instance.

Param Method (MethodObject): The Method param. MethodObject instance.

Param MultiFactorAuth (MultiFactorAuthObject): The MultiFactorAuth param. MultiFactorAuthObject instance.

Param Name (required, string): The Name param.

Param SingleSignOn (SingleSignOnObject): The SingleSignOn param. SingleSignOnObject instance.

Param UserDomain (string): The UserDomain param. String length must not exceed 63 characters.

Param UsernameModifier (string): The UsernameModifier param. String values: []string{"%USERINPUT%", "%USERINPUT%@%USERDOMAIN%", "%USERDOMAIN%\\\\%USERINPUT%"}
*/
type Config struct {
    AllowList []string `json:"allow_list,omitempty"`
    ObjectId string `json:"id,omitempty"`
    Lockout *LockoutObject `json:"lockout,omitempty"`
    Method *MethodObject `json:"method,omitempty"`
    MultiFactorAuth *MultiFactorAuthObject `json:"multi_factor_auth,omitempty"`
    Name string `json:"name"`
    SingleSignOn *SingleSignOnObject `json:"single_sign_on,omitempty"`
    UserDomain string `json:"user_domain,omitempty"`
    UsernameModifier string `json:"username_modifier,omitempty"`
}

/*
LockoutObject ""

Parent chains:
* "Config lockout"

Args:

Param FailedAttempts (int64): The FailedAttempts param. Value must be between 0 and 10.

Param LockoutTime (int64): The LockoutTime param. Value must be between 0 and 60.
*/
type LockoutObject struct {
    FailedAttempts int64 `json:"failed_attempts,omitempty"`
    LockoutTime int64 `json:"lockout_time,omitempty"`
}

/*
MethodObject ""

Parent chains:
* "Config method"

Args:

Param Kerberos (KerberosObject): The Kerberos param. KerberosObject instance.

Param Ldap (LdapObject): The Ldap param. LdapObject instance.

Param LocalDatabase (interface{}): The LocalDatabase param. interface{} instance.

Param Radius (RadiusObject): The Radius param. RadiusObject instance.

Param SamlIdp (SamlIdpObject): The SamlIdp param. SamlIdpObject instance.

Param Tacplus (TacplusObject): The Tacplus param. TacplusObject instance.
*/
type MethodObject struct {
    Kerberos *KerberosObject `json:"kerberos,omitempty"`
    Ldap *LdapObject `json:"ldap,omitempty"`
    LocalDatabase interface{} `json:"local_database,omitempty"`
    Radius *RadiusObject `json:"radius,omitempty"`
    SamlIdp *SamlIdpObject `json:"saml_idp,omitempty"`
    Tacplus *TacplusObject `json:"tacplus,omitempty"`
}

/*
KerberosObject ""

Parent chains:
* "Config method kerberos"

Args:

Param Realm (string): The Realm param.

Param ServerProfile (string): The ServerProfile param.
*/
type KerberosObject struct {
    Realm string `json:"realm,omitempty"`
    ServerProfile string `json:"server_profile,omitempty"`
}

/*
LdapObject ""

Parent chains:
* "Config method ldap"

Args:

Param LoginAttribute (string): The LoginAttribute param.

Param PasswdExpDays (int64): The PasswdExpDays param.

Param ServerProfile (string): The ServerProfile param.
*/
type LdapObject struct {
    LoginAttribute string `json:"login_attribute,omitempty"`
    PasswdExpDays int64 `json:"passwd_exp_days,omitempty"`
    ServerProfile string `json:"server_profile,omitempty"`
}

/*
RadiusObject ""

Parent chains:
* "Config method radius"

Args:

Param Checkgroup (bool): The Checkgroup param.

Param ServerProfile (string): The ServerProfile param.
*/
type RadiusObject struct {
    Checkgroup bool `json:"checkgroup,omitempty"`
    ServerProfile string `json:"server_profile,omitempty"`
}

/*
SamlIdpObject ""

Parent chains:
* "Config method saml_idp"

Args:

Param AttributeNameUsergroup (string): The AttributeNameUsergroup param. String lengh must be between 1 and 63 characters.

Param AttributeNameUsername (string): The AttributeNameUsername param. String lengh must be between 1 and 63 characters.

Param CertificateProfile (string): The CertificateProfile param. String length must not exceed 31 characters.

Param EnableSingleLogout (bool): The EnableSingleLogout param.

Param RequestSigningCertificate (string): The RequestSigningCertificate param. String length must not exceed 64 characters.

Param ServerProfile (string): The ServerProfile param. String length must not exceed 63 characters.
*/
type SamlIdpObject struct {
    AttributeNameUsergroup string `json:"attribute_name_usergroup,omitempty"`
    AttributeNameUsername string `json:"attribute_name_username,omitempty"`
    CertificateProfile string `json:"certificate_profile,omitempty"`
    EnableSingleLogout bool `json:"enable_single_logout,omitempty"`
    RequestSigningCertificate string `json:"request_signing_certificate,omitempty"`
    ServerProfile string `json:"server_profile,omitempty"`
}

/*
TacplusObject ""

Parent chains:
* "Config method tacplus"

Args:

Param Checkgroup (bool): The Checkgroup param.

Param ServerProfile (string): The ServerProfile param.
*/
type TacplusObject struct {
    Checkgroup bool `json:"checkgroup,omitempty"`
    ServerProfile string `json:"server_profile,omitempty"`
}

/*
MultiFactorAuthObject ""

Parent chains:
* "Config multi_factor_auth"

Args:

Param Factors (list of strings): The Factors param.

Param MfaEnable (bool): The MfaEnable param.
*/
type MultiFactorAuthObject struct {
    Factors []string `json:"factors,omitempty"`
    MfaEnable bool `json:"mfa_enable,omitempty"`
}

/*
SingleSignOnObject ""

Parent chains:
* "Config single_sign_on"

Args:

Param KerberosKeytab (string): The KerberosKeytab param. String length must not exceed 8192 characters.

Param Realm (string): The Realm param. String length must not exceed 127 characters.
*/
type SingleSignOnObject struct {
    KerberosKeytab string `json:"kerberos_keytab,omitempty"`
    Realm string `json:"realm,omitempty"`
}