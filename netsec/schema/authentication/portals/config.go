package portals

/*
Config ""

Parent chains:
* "Config"

Args:

Param AuthenticationProfile (string): The AuthenticationProfile param.

Param CertificateProfile (string): The CertificateProfile param.

Param GpUdpPort (int64): The GpUdpPort param. Value must be between 1 and 65535.

Param ObjectId (string, read-only): The ObjectId param. Example: "abcd-1234".

Param IdleTimer (int64): The IdleTimer param. Value must be between 1 and 1440.

Param RedirectHost (string): The RedirectHost param.

Param Timer (int64): The Timer param. Value must be between 1 and 1440.

Param TlsServiceProfile (string): The TlsServiceProfile param.
*/
type Config struct {
    AuthenticationProfile string `json:"authentication_profile,omitempty"`
    CertificateProfile string `json:"certificate_profile,omitempty"`
    GpUdpPort int64 `json:"gp_udp_port,omitempty"`
    ObjectId string `json:"id,omitempty"`
    IdleTimer int64 `json:"idle_timer,omitempty"`
    RedirectHost string `json:"redirect_host,omitempty"`
    Timer int64 `json:"timer,omitempty"`
    TlsServiceProfile string `json:"tls_service_profile,omitempty"`
}