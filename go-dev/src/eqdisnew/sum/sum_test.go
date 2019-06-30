package sum

import "testing"

func Test_Sum(t *testing.T) {
	r := Sum(1, 100, 1)
	if r != 5050 {
		t.Error("the Sum func is wrong")
	}

}
