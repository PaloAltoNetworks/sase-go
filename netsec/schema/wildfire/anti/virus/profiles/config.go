package profiles

/*
Config ""

Parent chains:
* "Config"

Args:

Param Description (string): The Description param.

Param ObjectId (string, read-only): The ObjectId param. Example: "abcd-1234".

Param MlavException (list of MlavExceptionObject objects): The MlavException param.

Param Name (required, string): The Name param.

Param PacketCapture (bool): The PacketCapture param.

Param Rules (list of RulesObject objects): The Rules param.

Param ThreatException (list of ThreatExceptionObject objects): The ThreatException param.
*/
type Config struct {
    Description string `json:"description,omitempty"`
    ObjectId string `json:"id,omitempty"`
    MlavException []MlavExceptionObject `json:"mlav_exception,omitempty"`
    Name string `json:"name"`
    PacketCapture bool `json:"packet_capture,omitempty"`
    Rules []RulesObject `json:"rules,omitempty"`
    ThreatException []ThreatExceptionObject `json:"threat_exception,omitempty"`
}

/*
MlavExceptionObject ""

Parent chains:
* "Config mlav_exception"

Args:

Param Description (string): The Description param.

Param Filename (string): The Filename param.

Param Name (string): The Name param.
*/
type MlavExceptionObject struct {
    Description string `json:"description,omitempty"`
    Filename string `json:"filename,omitempty"`
    Name string `json:"name,omitempty"`
}

/*
RulesObject ""

Parent chains:
* "Config rules"

Args:

Param Analysis (string): The Analysis param. String values: []string{"public-cloud", "private-cloud"}

Param Application (list of strings): The Application param.

Param Direction (string): The Direction param. String values: []string{"download", "upload", "both"}

Param FileType (list of strings): The FileType param.

Param Name (string): The Name param.
*/
type RulesObject struct {
    Analysis string `json:"analysis,omitempty"`
    Application []string `json:"application,omitempty"`
    Direction string `json:"direction,omitempty"`
    FileType []string `json:"file_type,omitempty"`
    Name string `json:"name,omitempty"`
}

/*
ThreatExceptionObject ""

Parent chains:
* "Config threat_exception"

Args:

Param Name (string): The Name param.

Param Notes (string): The Notes param.
*/
type ThreatExceptionObject struct {
    Name string `json:"name,omitempty"`
    Notes string `json:"notes,omitempty"`
}