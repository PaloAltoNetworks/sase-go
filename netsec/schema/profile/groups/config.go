package groups

/*
Config ""

Parent chains:
* "Config"

Args:

Param DnsSecurity (list of strings): The DnsSecurity param.

Param FileBlocking (list of strings): The FileBlocking param.

Param ObjectId (string, read-only): The ObjectId param. Example: "abcd-1234".

Param Name (required, string): The Name param.

Param SaasSecurity (list of strings): The SaasSecurity param.

Param Spyware (list of strings): The Spyware param.

Param UrlFiltering (list of strings): The UrlFiltering param.

Param VirusAndWildfireAnalysis (list of strings): The VirusAndWildfireAnalysis param.

Param Vulnerability (list of strings): The Vulnerability param.
*/
type Config struct {
    DnsSecurity []string `json:"dns_security,omitempty"`
    FileBlocking []string `json:"file_blocking,omitempty"`
    ObjectId string `json:"id,omitempty"`
    Name string `json:"name"`
    SaasSecurity []string `json:"saas_security,omitempty"`
    Spyware []string `json:"spyware,omitempty"`
    UrlFiltering []string `json:"url_filtering,omitempty"`
    VirusAndWildfireAnalysis []string `json:"virus_and_wildfire_analysis,omitempty"`
    Vulnerability []string `json:"vulnerability,omitempty"`
}