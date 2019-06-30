package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
for i := 0; i < 5; i++ {
	var v int 
	fmt.Println("%d ",v)
	v = 5
}

s := ""
for s != "aaaaa"; {
    fmt.Println("Value of s:", s)
    s = s + "a"
}

for i := 0; i < 3; {
    fmt.Println("Value of i:", i)
}

for i := 0; ; i++ {
    fmt.Println("Value of i is now:", i)
}
