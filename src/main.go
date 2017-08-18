package main

import (
	"fmt"
	"gopkg.in/couchbase/gocb.v1"
	"pims/system"
	//"pims/tools"
	"strconv"
)

func addAttributes(bucket *gocb.Bucket, ids []int, names []string) {

	for _, id := range ids {
		fmt.Println(id)
		name := names[id-1]
		attr := pims.Attribute{id, "attribute", "name ", "decimal"}
		attr.Name = name
		bucket.Insert(strconv.Itoa(id), attr, 0)
	}
}

func addAttributeSet(bucket *gocb.Bucket, name string, attributeIds []int) {
	attributeSet := pims.AttributeSet{1, "attributeSet", name, &attributeIds}
	bucket.Insert("a1", attributeSet, 0)
}




func addProduct(bucket *gocb.Bucket, entities[] pims.Entity)  {
		for i,e:= range entities{
			bucket.Insert(strconv.Itoa(i), e, 0)
		}
	//product.Attributes = ([]map[string]interface{})data["attributes"]
	//product.AttributeSets = data["attributeSets"]

}



func main() {
	cluster, _ := gocb.Connect("couchbase://timms.turbointernational.com")
	bucket, _ := cluster.OpenBucket("pims", "")

	if false {
		ids := []int{1, 2, 3, 4}
		names := []string{"A", "B", "C", "D"}
		addAttributes(bucket, ids, names)
		addAttributeSet(bucket, "Default", ids)
	}
	entities:= pims.ReadMigration("var/products.json")
	fmt.Println(entities)
	addProduct(bucket, entities)


}
