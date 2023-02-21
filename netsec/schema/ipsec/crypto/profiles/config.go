package profiles

/*
Config ""

Parent chains:
* "Config"

Args:

Param Ah (AhObject): The Ah param. AhObject instance. One of these params should be specified:  ah or esp.

Param DhGroup (string): The DhGroup param. String values: []string{"no-pfs", "group1", "group2", "group5", "group14", "group19", "group20"} Default: "group2".

Param Esp (EspObject): The Esp param. EspObject instance. One of these params should be specified:  ah or esp.

Param ObjectId (string, read-only): The ObjectId param. Example: "abcd-1234".

Param Lifesize (LifesizeObject): The Lifesize param. LifesizeObject instance.

Param Lifetime (required, LifetimeObject): The Lifetime param. LifetimeObject instance.

Param Name (required, string): The Name param. String length must not exceed 31 characters.
*/
type Config struct {
    Ah *AhObject `json:"ah,omitempty"`
    DhGroup string `json:"dh_group,omitempty"`
    Esp *EspObject `json:"esp,omitempty"`
    ObjectId string `json:"id,omitempty"`
    Lifesize *LifesizeObject `json:"lifesize,omitempty"`
    Lifetime LifetimeObject `json:"lifetime"`
    Name string `json:"name"`
}

/*
AhObject ""

Parent chains:
* "Config ah"

Args:

Param Authentication (required, list of strings): The Authentication param.
*/
type AhObject struct {
    Authentication []string `json:"authentication"`
}

/*
EspObject ""

Parent chains:
* "Config esp"

Args:

Param Authentication (required, list of strings): The Authentication param.

Param Encryption (required, list of strings): The Encryption param.
*/
type EspObject struct {
    Authentication []string `json:"authentication"`
    Encryption []string `json:"encryption"`
}

/*
LifesizeObject ""

Parent chains:
* "Config lifesize"

Args:

Param Gb (int64): The Gb param. Value must be between 1 and 65535.

Param Kb (int64): The Kb param. Value must be between 1 and 65535.

Param Mb (int64): The Mb param. Value must be between 1 and 65535.

Param Tb (int64): The Tb param. Value must be between 1 and 65535.
*/
type LifesizeObject struct {
    Gb int64 `json:"gb,omitempty"`
    Kb int64 `json:"kb,omitempty"`
    Mb int64 `json:"mb,omitempty"`
    Tb int64 `json:"tb,omitempty"`
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
    Days int64 `json:"days,omitempty"`
    Hours int64 `json:"hours,omitempty"`
    Minutes int64 `json:"minutes,omitempty"`
    Seconds int64 `json:"seconds,omitempty"`
}