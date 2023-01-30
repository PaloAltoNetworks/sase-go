package profiles

/*
Config ""

Parent chains:
* "Config"

Args:

Param Description (string): The Description param.

Param HttpHeaderInsertion (list of HttpHeaderInsertionObject objects): The HttpHeaderInsertion param.

Param ObjectId (string, read-only): The ObjectId param. Example: "abcd-1234".

Param Name (required, string): The Name param.
*/
type Config struct {
	Description         string                      `json:"description,omitempty"`
	HttpHeaderInsertion []HttpHeaderInsertionObject `json:"http_header_insertion,omitempty"`
	ObjectId            string                      `json:"id,omitempty"`
	Name                string                      `json:"name"`
}

/*
HttpHeaderInsertionObject ""

Parent chains:
* "Config http_header_insertion"

Args:

Param Name (required, string): The Name param.

Param Type (required, list of TypeObject objects): The Type param.
*/
type HttpHeaderInsertionObject struct {
	Name string       `json:"name"`
	Type []TypeObject `json:"type"`
}

/*
TypeObject ""

Parent chains:
* "Config http_header_insertion type"

Args:

Param Domains (required, list of strings): The Domains param.

Param Headers (required, list of HeadersObject objects): The Headers param.

Param Name (required, string): The Name param.
*/
type TypeObject struct {
	Domains []string        `json:"domains"`
	Headers []HeadersObject `json:"headers"`
	Name    string          `json:"name"`
}

/*
HeadersObject ""

Parent chains:
* "Config http_header_insertion type headers"

Args:

Param Header (required, string): The Header param.

Param Log (bool): The Log param. Default: false

Param Name (required, string): The Name param.

Param Value (required, string): The Value param.
*/
type HeadersObject struct {
	Header string `json:"header"`
	Log    bool   `json:"log,omitempty"`
	Name   string `json:"name"`
	Value  string `json:"value"`
}
