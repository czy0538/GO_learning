// ping.go
package main

import (
	"os"
	"os/exec"
)

func main() {
	c := exec.Command("ping", "www.hitwh.edu.cn")
	c.Stdout = os.Stdout
	c.Run()

}
