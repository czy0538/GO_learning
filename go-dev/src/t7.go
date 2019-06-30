package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10; i++ {
		fmt.Println(rnd.Int())
	}
}
