package src

import (
	"testing"
)

func Test_Split(t *testing.T) {
	a, b := Split(17)
	if a == 7 && b == 10 {
		t.Log("ok")
	} else {
		t.Error("wrong")
	}
}
