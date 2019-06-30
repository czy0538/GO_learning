package sum

/************************************
等差数列求和:首项为a,公差为d,项数为n
************************************/
func Sum(a1, num, dis int) int {
	return (num*a1 + num*(num-1)*dis/2)
}
