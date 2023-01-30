package categories

/*
Config ""

Parent chains:
* "Config"

Args:

Param Description (string): The Description param.

Param List (list of strings): The List param.

Param Name (required, string): The Name param.

Param Type (string): The Type param. String values: []string{"URL List", "Category Match"} Default: "URL List".
*/
type Config struct {
	Description string   `json:"description,omitempty"`
	List        []string `json:"list,omitempty"`
	Name        string   `json:"name"`
	Type        string   `json:"type,omitempty"`
}
