package pims

type Entity struct {
	uuid int
	Name string
	Description string
	Attributes map[string]interface{}
	AttributeSets *[]int
}