package pims

type Entity struct {
	Id  int  `json:"id"`
	Name string  `json:"name"`
	Description string `json:"string"`
	Attributes[] map[string]interface{}`json:"attributes"`
	AttributeSets []int `json:"attrbibuteSets"`
}