// errors2.go
package main
import (
	"errors"
	"fmt"
	"math"
)
func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math - square root of negative number")
	}
	return math.Sqrt(f), nil
}
func main() {
	f, err := Sqrt(-1)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	fmt.Println(f)


	
}
var (
    EPERM       Error = Errno(syscall.EPERM)
    ENOENT      Error = Errno(syscall.ENOENT)
    ESRCH       Error = Errno(syscall.ESRCH)
    EINTR       Error = Errno(syscall.EINTR)
    EIO         Error = Errno(syscall.EIO)
    ...
)