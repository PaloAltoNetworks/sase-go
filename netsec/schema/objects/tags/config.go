package tags

/*
Config ""

Parent chains:
* "Config"

Args:

Param Color (string): The Color param. String values: []string{"Red", "Green", "Blue", "Yellow", "Copper", "Orange", "Purple", "Gray", "Light Green", "Cyan", "Light Gray", "Blue Gray", "Lime", "Black", "Gold", "Brown", "Olive", "Maroon", "Red-Orange", "Yellow-Orange", "Forest Green", "Turquoise Blue", "Azure Blue", "Cerulean Blue", "Midnight Blue", "Medium Blue", "Cobalt Blue", "Violet Blue", "Blue Violet", "Medium Violet", "Medium Rose", "Lavender", "Orchid", "Thistle", "Peach", "Salmon", "Magenta", "Red Violet", "Mahogany", "Burnt Sienna", "Chestnut"}

Param Comments (string): The Comments param. String lengh must be between 0 and 1023 characters.

Param ObjectId (string, read-only): The ObjectId param. Example: "abcd-1234".

Param Name (required, string): The Name param. String length must not exceed 127 characters.
*/
type Config struct {
    Color string `json:"color,omitempty"`
    Comments string `json:"comments,omitempty"`
    ObjectId string `json:"id,omitempty"`
    Name string `json:"name"`
}