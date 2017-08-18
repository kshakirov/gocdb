package pims

import (
	"io/ioutil"
	"encoding/json"
	"fmt"
)

func ReadMigration (fileName string)  []Entity{

	js, err := ioutil.ReadFile(fileName)
	var data[]  Entity
	if err == nil {
		err1 := json.Unmarshal(js, &data)
		if err1 == nil {
			return data
		} else {
			fmt.Println(err1)
		}

	} else {
		fmt.Print(err)
	}
	return data
}