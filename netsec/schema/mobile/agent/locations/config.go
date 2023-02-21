package locations

/*
Config ""

Parent chains:
* "Config"

Args:

Param Region (list of RegionObject objects): The Region param.
*/
type Config struct {
    Region []RegionObject `json:"region,omitempty"`
}

/*
RegionObject ""

Parent chains:
* "Config region"

Args:

Param Locations (list of strings): The Locations param.

Param Name (string): The Name param.
*/
type RegionObject struct {
    Locations []string `json:"locations,omitempty"`
    Name string `json:"name,omitempty"`
}