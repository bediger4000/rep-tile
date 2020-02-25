package main

import "fmt"

func main() {
	fmt.Println(`/Courier 20 selectfont        
150 450 translate
50 50 scale
.01 setlinewidth`)
	drawReptile(1, 0., 0., 1.0, 1.0, 1.0)

	fmt.Println("showpage")
}

type point struct {
	x float64
	y float64
}

func drawReptile(ply int, x, y float64, scale float64, xFlip, yFlip float64) {

	if ply == 0 {
		return
	}

	fmt.Println("newpath\ngsave")
	fmt.Printf("%f %f moveto\n", x, y)
	fmt.Printf("%f %f lineto\n", x+scale, y)
	fmt.Printf("%f %f lineto\n", x+scale, y+scale)
	fmt.Printf("%f %f lineto\n", x-scale, y+scale)
	fmt.Printf("%f %f lineto\n", x-scale, y-scale)
	fmt.Printf("%f %f lineto\n", x, y-scale)
	fmt.Printf("%f %f lineto\n", x, y)

	fmt.Println("stroke\ngrestore")
}
