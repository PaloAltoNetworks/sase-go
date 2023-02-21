package jobs

/*
Config ""

Parent chains:
* "Config"

Args:

Param Details (string): The Details param.

Param EndTs (string): The EndTs param.

Param ObjectId (string): The ObjectId param.

Param InsertTs (string): The InsertTs param.

Param JobResult (string): The JobResult param.

Param JobStatus (string): The JobStatus param.

Param JobType (string): The JobType param.

Param LastUpdate (string): The LastUpdate param.

Param OpaqueInt (string): The OpaqueInt param.

Param OpaqueStr (string): The OpaqueStr param.

Param Owner (string): The Owner param.

Param ParentId (string): The ParentId param.

Param Percent (string): The Percent param.

Param ResultI (string): The ResultI param.

Param ResultStr (string): The ResultStr param.

Param SessionId (string): The SessionId param.

Param StartTs (string): The StartTs param.

Param StatusI (string): The StatusI param.

Param StatusStr (string): The StatusStr param.

Param Summary (string): The Summary param.

Param TypeI (string): The TypeI param.

Param TypeStr (string): The TypeStr param.

Param Uname (string): The Uname param.
*/
type Config struct {
    Details string `json:"details,omitempty"`
    EndTs string `json:"end_ts,omitempty"`
    ObjectId string `json:"id,omitempty"`
    InsertTs string `json:"insert_ts,omitempty"`
    JobResult string `json:"job_result,omitempty"`
    JobStatus string `json:"job_status,omitempty"`
    JobType string `json:"job_type,omitempty"`
    LastUpdate string `json:"last_update,omitempty"`
    OpaqueInt string `json:"opaque_int,omitempty"`
    OpaqueStr string `json:"opaque_str,omitempty"`
    Owner string `json:"owner,omitempty"`
    ParentId string `json:"parent_id,omitempty"`
    Percent string `json:"percent,omitempty"`
    ResultI string `json:"result_i,omitempty"`
    ResultStr string `json:"result_str,omitempty"`
    SessionId string `json:"session_id,omitempty"`
    StartTs string `json:"start_ts,omitempty"`
    StatusI string `json:"status_i,omitempty"`
    StatusStr string `json:"status_str,omitempty"`
    Summary string `json:"summary,omitempty"`
    TypeI string `json:"type_i,omitempty"`
    TypeStr string `json:"type_str,omitempty"`
    Uname string `json:"uname,omitempty"`
}