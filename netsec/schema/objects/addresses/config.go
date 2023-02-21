package addresses

/*
Config ""

Parent chains:
* "Config"

Args:

Param Description (string): The Description param. String lengh must be between 0 and 1023 characters.

Param Fqdn (string): The Fqdn param. String lengh must be between 1 and 255 characters. One of these params should be specified:  fqdn, ip_netmask, ip_range, or ip_wildcard.

Param ObjectId (string, read-only): The ObjectId param. Example: "abcd-1234".

Param IpNetmask (string): The IpNetmask param. One of these params should be specified:  fqdn, ip_netmask, ip_range, or ip_wildcard.

Param IpRange (string): The IpRange param. One of these params should be specified:  fqdn, ip_netmask, ip_range, or ip_wildcard.

Param IpWildcard (string): The IpWildcard param. One of these params should be specified:  fqdn, ip_netmask, ip_range, or ip_wildcard.

Param Name (required, string): The Name param. String length must not exceed 63 characters.

Param Tag (list of strings): The Tag param. Array should contain at most 64 items.

Param Type (string, read-only): The Type param. Example: "container".
*/
type Config struct {
    Description string `json:"description,omitempty"`
    Fqdn string `json:"fqdn,omitempty"`
    ObjectId string `json:"id,omitempty"`
    IpNetmask string `json:"ip_netmask,omitempty"`
    IpRange string `json:"ip_range,omitempty"`
    IpWildcard string `json:"ip_wildcard,omitempty"`
    Name string `json:"name"`
    Tag []string `json:"tag,omitempty"`
    Type string `json:"type,omitempty"`
}