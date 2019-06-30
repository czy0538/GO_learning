package split

import (
	"testing"
)

/*func Test_Split(t *testing.T) {

	r1, r2 := Split(9)
	if r1 != 9 && r2 != 5 {
		t.Error("The result is wrong")

	}
}*/
func Test_Split(t *testing.T) {
	a, b := Split(17)
	if a == 7 && b == 10 {
		t.Log("ok,right")
	} else {
		t.Error("wrong")
	}
}
