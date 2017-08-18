package pims

type Entity struct {
	Uuid  string  `json:"uuid"`
	Name string  `json:"name"`
	Description string `json:"string"`
	Attributes[] map[string]interface{}`json:"attributes"`
	AttributeSets []int `json:"attrbibuteSets"`
}