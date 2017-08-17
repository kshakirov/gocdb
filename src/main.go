package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/couchbase/gocb.v1"
	"io/ioutil"
	"pims/system"
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

func main() {
	cluster, _ := gocb.Connect("couchbase://timms.turbointernational.com")
	bucket, _ := cluster.OpenBucket("pims", "")

	if false {
		ids := []int{1, 2, 3, 4}
		names := []string{"A", "B", "C", "D"}
		addAttributes(bucket, ids, names)
		addAttributeSet(bucket, "Default", ids)
	}
	js, err := ioutil.ReadFile("var/products.json")
	//js = []byte(`{"num":6.13,"strs":["a","b"]}`)
	var dat []map[string]interface{}
	if err == nil {
		err1 := json.Unmarshal(js, &dat)
		if err1 == nil {
			fmt.Println(dat)
			//fmt.Println(num)
		} else {
			fmt.Println(err1)
		}

	} else {
		fmt.Print(err)
	}

}
