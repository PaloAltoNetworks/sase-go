package profiles

/*
Config ""

Parent chains:
* "Config"

Args:

Param AuthenticationMultiple (int64): The AuthenticationMultiple param. Value must be less than or equal to 50. Default: 0

Param DhGroup (required, list of strings): The DhGroup param.

Param Encryption (required, list of strings): The Encryption param.

Param Hash (required, list of strings): The Hash param.

Param ObjectId (string, read-only): The ObjectId param. Example: "abcd-1234".

Param Lifetime (LifetimeObject): The Lifetime param. LifetimeObject instance.

Param Name (required, string): The Name param. String length must not exceed 31 characters.
*/
type Config struct {
	AuthenticationMultiple int64           `json:"authentication_multiple,omitempty"`
	DhGroup                []string        `json:"dh_group"`
	Encryption             []string        `json:"encryption"`
	Hash                   []string        `json:"hash"`
	ObjectId               string          `json:"id,omitempty"`
	Lifetime               *LifetimeObject `json:"lifetime,omitempty"`
	Name                   string          `json:"name"`
}

/*
LifetimeObject ""

Parent chains:
* "Config lifetime"

Args:

Param Days (int64): The Days param. Value must be between 1 and 365.

Param Hours (int64): The Hours param. Value must be between 1 and 65535.

Param Minutes (int64): The Minutes param. Value must be between 3 and 65535.

Param Seconds (int64): The Seconds param. Value must be between 180 and 65535.
*/
type LifetimeObject struct {
	Days    int64 `json:"days,omitempty"`
	Hours   int64 `json:"hours,omitempty"`
	Minutes int64 `json:"minutes,omitempty"`
	Seconds int64 `json:"seconds,omitempty"`
}
