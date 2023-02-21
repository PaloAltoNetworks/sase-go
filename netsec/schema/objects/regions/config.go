package regions

/*
Config ""

Parent chains:
* "Config"

Args:

Param Address (list of strings): The Address param.

Param GeoLocation (GeoLocationObject): The GeoLocation param. GeoLocationObject instance.

Param ObjectId (string, read-only): The ObjectId param. Example: "abcd-1234".

Param Name (required, string): The Name param. String length must not exceed 31 characters.
*/
type Config struct {
    Address []string `json:"address,omitempty"`
    GeoLocation *GeoLocationObject `json:"geo_location,omitempty"`
    ObjectId string `json:"id,omitempty"`
    Name string `json:"name"`
}

/*
GeoLocationObject ""

Parent chains:
* "Config geo_location"

Args:

Param Latitude (required, float64): The Latitude param. Value must be between -90.000000 and 90.000000.

Param Longitude (required, float64): The Longitude param. Value must be between -180.000000 and 180.000000.
*/
type GeoLocationObject struct {
    Latitude float64 `json:"latitude"`
    Longitude float64 `json:"longitude"`
}