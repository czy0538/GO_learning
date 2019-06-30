package sum

import "testing"

func Test_Sum(t *testing.T) {
	r := Sum(100)
	if r != 5050 {
		t.Error("结果不对")

	} else {
		t.Log("功能正确")
	}

}
