package process

import "testing"

func Test_Process(t *testing.T) {
	r := Process(1, 100, 1)
	if r != 5050 {
		t.Error("wrong")
	}
}
