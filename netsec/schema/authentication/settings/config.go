package settings

/*
Config ""

Parent chains:
* "Config"

Args:

Param AuthenticationProfile (required, string): The AuthenticationProfile param.

Param Os (required, string): The Os param. String values: []string{"Any", "Android", "Browser", "Chrome", "IoT", "Linux", "Mac", "Satellite", "Windows", "WindowsUWP", "iOS"} Default: "Any".

Param UserCredentialOrClientCertRequired (required, bool): The UserCredentialOrClientCertRequired param.
*/
type Config struct {
    AuthenticationProfile string `json:"authentication_profile"`
    Os string `json:"os"`
    UserCredentialOrClientCertRequired bool `json:"user_credential_or_client_cert_required"`
}