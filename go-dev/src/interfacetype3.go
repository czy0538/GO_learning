package main

import "fmt"

func main() {
	people := People{"张三", 20}
	it := []interface{}{1, "hello", people, true}
	for i, e := range it {
		switch val := e.(type) {
		case int:
			fmt.Printf("it[%d] type is int,val=%d\n", i, val)
		case string:
			fmt.Printf("it[%d] type is string,val=%s\n", i, val)
		case People:
			fmt.Printf("it[%d] type is People,val=%v\n", i, val)
		case bool:
			fmt.Printf("it[%d] type is bool,val=%v\n", i, val)
		default:
			fmt.Println("Unknown type of it[", i, "]")
		}
	}

}
