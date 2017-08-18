package main

import (
	"fmt"
	"gopkg.in/couchbase/gocb.v1"
	"pims/system"
	"github.com/satori/go.uuid"
)

func addAttributes(bucket *gocb.Bucket, names []string) []string {
	var uuids = make([]string, 4, 8)
	for _, name := range names {
		id := uuid.NewV4().String()
		fmt.Println(id)
		attr := pims.Attribute{id, "attribute", "name ", "decimal"}
		attr.Name = name
		bucket.Insert(id, attr, 0)
		uuids = append(uuids, id)
	}
	return uuids
}

func addAttributeSet(bucket *gocb.Bucket, name string, attributeIds []string) {
	attributeSet := pims.AttributeSet{uuid.NewV4().String(), "attributeSet", name, &attributeIds}
	bucket.Insert("a1", attributeSet, 0)
}

func addProduct(bucket *gocb.Bucket, entities[] pims.Entity) {
	for _, e := range entities {
		bucket.Insert(uuid.NewV4().String(), e, 0)
	}
	//product.Attributes = ([]map[string]interface{})data["attributes"]
	//product.AttributeSets = data["attributeSets"]

}

func main() {
	cluster, _ := gocb.Connect("couchbase://timms.turbointernational.com")
	bucket, _ := cluster.OpenBucket("pims", "")

	names := []string{"A", "B", "C", "D"}
	ids := addAttributes(bucket, names)
	addAttributeSet(bucket, "Default", ids)

	entities := pims.ReadMigration("var/products.json")
	fmt.Println(entities)
	addProduct(bucket, entities)

}
