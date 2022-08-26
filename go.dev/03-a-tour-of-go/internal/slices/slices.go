package slices

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var picture = make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		picture[y] = make([]uint8, dx)
		for x := 0; x < dx; x++ {
			picture[y][x] = uint8(x ^ y)
		}
	}
	return picture
}

func Run() {
	pic.Show(Pic)
}
