package s

import "testing"

func Test_S(t *testing.T) {
	r := S(1, 100, 1)
	if r == 5050 {
		t.Log("功能正确")
	} else {
		t.Error("功能错误")
	}
}
