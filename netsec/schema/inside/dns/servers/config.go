package servers

/*
Config ""

Parent chains:
* "Config"

Args:

Param DomainName (required, list of strings): The DomainName param.

Param Name (required, string): The Name param.

Param Primary (required, string): The Primary param.

Param Secondary (string): The Secondary param.
*/
type Config struct {
    DomainName []string `json:"domain_name"`
    Name string `json:"name"`
    Primary string `json:"primary"`
    Secondary string `json:"secondary,omitempty"`
}