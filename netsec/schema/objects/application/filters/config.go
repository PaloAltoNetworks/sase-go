package filters

/*
Config ""

Parent chains:
* "Config"

Args:

Param Category (list of strings): The Category param.

Param Evasive (bool): The Evasive param.

Param ExcessiveBandwidthUse (bool): The ExcessiveBandwidthUse param.

Param Exclude (list of strings): The Exclude param.

Param HasKnownVulnerabilities (bool): The HasKnownVulnerabilities param.

Param ObjectId (string, read-only): The ObjectId param. Example: "abcd-1234".

Param IsSaas (bool): The IsSaas param.

Param Name (required, string): The Name param. String length must not exceed 31 characters.

Param NewAppid (bool): The NewAppid param.

Param Pervasive (bool): The Pervasive param.

Param ProneToMisuse (bool): The ProneToMisuse param.

Param Risk (list of ints): The Risk param.

Param SaasCertifications (list of strings): The SaasCertifications param.

Param SaasRisk (list of strings): The SaasRisk param.

Param Subcategory (list of strings): The Subcategory param.

Param Tagging (TaggingObject): The Tagging param. TaggingObject instance.

Param Technology (list of strings): The Technology param.

Param TransfersFiles (bool): The TransfersFiles param.

Param TunnelsOtherApps (bool): The TunnelsOtherApps param.

Param UsedByMalware (bool): The UsedByMalware param.
*/
type Config struct {
	Category                []string       `json:"category,omitempty"`
	Evasive                 bool           `json:"evasive,omitempty"`
	ExcessiveBandwidthUse   bool           `json:"excessive_bandwidth_use,omitempty"`
	Exclude                 []string       `json:"exclude,omitempty"`
	HasKnownVulnerabilities bool           `json:"has_known_vulnerabilities,omitempty"`
	ObjectId                string         `json:"id,omitempty"`
	IsSaas                  bool           `json:"is_saas,omitempty"`
	Name                    string         `json:"name"`
	NewAppid                bool           `json:"new_appid,omitempty"`
	Pervasive               bool           `json:"pervasive,omitempty"`
	ProneToMisuse           bool           `json:"prone_to_misuse,omitempty"`
	Risk                    []int64        `json:"risk,omitempty"`
	SaasCertifications      []string       `json:"saas_certifications,omitempty"`
	SaasRisk                []string       `json:"saas_risk,omitempty"`
	Subcategory             []string       `json:"subcategory,omitempty"`
	Tagging                 *TaggingObject `json:"tagging,omitempty"`
	Technology              []string       `json:"technology,omitempty"`
	TransfersFiles          bool           `json:"transfers_files,omitempty"`
	TunnelsOtherApps        bool           `json:"tunnels_other_apps,omitempty"`
	UsedByMalware           bool           `json:"used_by_malware,omitempty"`
}

/*
TaggingObject ""

Parent chains:
* "Config tagging"

Args:

Param NoTag (bool): The NoTag param.

Param Tag (list of strings): The Tag param.
*/
type TaggingObject struct {
	NoTag bool     `json:"no_tag,omitempty"`
	Tag   []string `json:"tag,omitempty"`
}
