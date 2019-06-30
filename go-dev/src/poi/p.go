package poi

type Point struct {
	X int
	Y int
}

func (p *Point) GetX(a int) {
	p.X = a
}
func (p *Point) GetY(b int) {
	p.Y = b
}
