package sum

func Sum(n int) int {
	var r int
	for i := 0; i <= n; i++ {
		r += i
	}
	return r
}
