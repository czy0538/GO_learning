// json2.go
package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	Type    string
	City    string
	Country string
}
type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

func main() {
	b := []byte(`{"FirstName":"Jan","LastName":"Kersschot",
		"Addresses":[{"Type":"private","City":"Aartselaar",
		"Country":"Belgium"},{"Type":"work","City":"Boom",
		"Country":"Belgium"}],
		"Remark":"none"}`) //定义切片b，存放一段JSON文本
	var vc VCard
	err := json.Unmarshal(b, &vc) //反序列化b到vc中
	if err == nil {
		fmt.Println(vc.FirstName)
		fmt.Println(vc.LastName)
		for _, v := range vc.Addresses {
			fmt.Println(*v)
		}
		fmt.Println(vc.Remark)
	}
}
