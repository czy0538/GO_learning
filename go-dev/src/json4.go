// json4.go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	file, _ := os.Open("vcard.json")
	defer file.Close()
	dec := json.NewDecoder(file) //建立解码器
	var f map[string]interface{} //建立存储JSON对象的map变量 f
	err := dec.Decode(&f)        //解码到f中
	if err != nil {
		log.Println("Error in decoding json")
	}
	for k, v := range f { //遍历map变量f
		switch vv := v.(type) { //类型判断switch
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
