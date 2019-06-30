package sum

import "testing"

func Test_Sum(t *testing.T) {
	result := Sum(1, 100, 1)
	if result != 5050 {
		t.Error("The result is wrong")
	}
}
