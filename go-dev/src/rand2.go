package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	fmt.Println("The rand num is", rand.Intn(100), "OK hah")
}
