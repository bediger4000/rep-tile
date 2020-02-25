package main

import "fmt"

func main() {
	fmt.Println(`/Courier 20 selectfont        
150 450 translate
50 50 scale
.01 setlinewidth`)
	drawReptile(1, 0., 0., 1.0, 1.0, 1.0)
	// drawReptile(1, -1., 1., 1.0, 1.0, 1.0)
	// drawReptile(1, 1., 1., 1.0, -1.0, 1.0)
	// drawReptile(1, -1., -1., 1.0, 1.0, -1.0)

	//awReptile(1, 1., 1., 1.0, -1.0, 1.0)
	drawReptile(1, .5, .5, .50, -1.0, 1.0)
	drawReptile(1, .25, .75, .25, 1.0, 1.0)
	drawReptile(1, .375, .875, .125, -1.0, 1.0)

	fmt.Println("showpage")
}

type point struct {
	x float64
	y float64
}

var tilePoints = []point{
	point{0., 0.},
	point{1., 0.},
	point{1., 1.},
	point{-1., 1.},
	point{-1., -1.},
	point{0., -1.},
	point{0., 0.},
}

func drawReptile(ply int, dx, dy float64, scale float64, xFlip, yFlip float64) {

	if ply == 0 {
		return
	}

	fmt.Println("newpath\ngsave")
	fmt.Printf("%f %f moveto\n", dx, dy)
	for _, pt := range tilePoints {
		x := scale*(xFlip*pt.x) + dx
		y := scale*(yFlip*pt.y) + dy
		fmt.Printf("%f %f lineto\n", x, y)
	}

	// Redraw inside
	scale /= 2.0
	ddx := dx * scale
	ddy := dy * scale
	drawReptile(ply-1, dx, dy, scale, xFlip, yFlip)
	drawReptile(ply-1, dx-ddx, dy+ddy, scale, xFlip, yFlip)
	drawReptile(ply-1, dx+ddx, dy+ddy, scale, -xFlip, yFlip)
	drawReptile(ply-1, dx-ddx, dy-ddy, scale, xFlip, -yFlip)

	fmt.Println("stroke\ngrestore")
}
