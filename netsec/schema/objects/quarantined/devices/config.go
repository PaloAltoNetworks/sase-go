package devices

/*
Config ""

Parent chains:
* "Config"

Args:

Param ConnectionName (required, string): The ConnectionName param.

Param HostId (required, string): The HostId param.

Param Ip (string): The Ip param.

Param Ipv6 (string): The Ipv6 param.

Param SerialNumber (string): The SerialNumber param.

Param ServiceType (required, string): The ServiceType param.
*/
type Config struct {
	ConnectionName string `json:"connection_name"`
	HostId         string `json:"host_id"`
	Ip             string `json:"ip,omitempty"`
	Ipv6           string `json:"ipv6,omitempty"`
	SerialNumber   string `json:"serial_number,omitempty"`
	ServiceType    string `json:"service_type"`
}
