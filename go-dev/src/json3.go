// json3.go
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	b := []byte(`{"FirstName":"Jan","LastName":"Kersschot",
		"Addresses":[{"Type":"private","City":"Aartselaar",
		"Country":"Belgium"},{"Type":"work","City":"Boom",
		"Country":"Belgium"}],
		"Remark":"none"}`) //定义切片b，存放一段JSON文本
	var f interface{}
	err := json.Unmarshal(b, &f) //反序列化b到f中
	m := f.(map[string]interface{})
	if err == nil {
		for k, v := range m {
			switch vv := v.(type) {
			case string:
				fmt.Println(k, "is string", vv)
			case int:
				fmt.Println(k, "is int", vv)
			case []interface{}:
				fmt.Println(k, "is an array:")
				for i, u := range vv {
					fmt.Println(i, u)
				}
			default:
				fmt.Println(k, ",I don’t know how to handle")
			}
		}
	}
}
