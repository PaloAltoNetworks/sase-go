package sequences

/*
Config ""

Parent chains:
* "Config"

Args:

Param AuthenticationProfiles (list of strings): The AuthenticationProfiles param.

Param ObjectId (string, read-only): The ObjectId param. Example: "abcd-1234".

Param Name (required, string): The Name param.

Param UseDomainFindProfile (bool): The UseDomainFindProfile param. Default: true
*/
type Config struct {
    AuthenticationProfiles []string `json:"authentication_profiles,omitempty"`
    ObjectId string `json:"id,omitempty"`
    Name string `json:"name"`
    UseDomainFindProfile bool `json:"use_domain_find_profile,omitempty"`
}