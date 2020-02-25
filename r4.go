package main

import "fmt"

func main() {
	fmt.Println(`/Courier 20 selectfont        
150 450 translate
50 50 scale
.01 setlinewidth`)
	drawReptile(1, 0., 0., 1.0, 1.0, 1.0)

	/*
		drawReptile(1, .0, .0, .50, 1.0, 1.0)    // A (0., 0.) 0.5 <1, 1>
		drawReptile(1, .5, .5, .50, -1.0, 1.0)   // B (.5, .5) 0.5 <-1, 1>
		drawReptile(1, -.5, .5, .50, 1.0, 1.0)   // C (-.5,.5) 0.5 <1, 1>
		drawReptile(1, -.5, -.5, .50, 1.0, -1.0) // D (-.5,-.5) 0.5 <1, -1>
	*/

	/* d = scale/2 */

	// A filled in,  (0., 0.) 0.5 <1, 1>
	drawReptile(1, .00, .00, .25, 1.0, 1.0)    // x     y
	drawReptile(1, .25, .25, .25, -1.0, 1.0)   // x+d   y+d
	drawReptile(1, -.25, .25, .25, 1.0, 1.0)   // x-d   y+d
	drawReptile(1, -.25, -.25, .25, 1.0, -1.0) // x-d   y-d

	// B filled in,  (.5, .5) 0.5 <-1, 1>
	drawReptile(1, .50, .50, .25, -1.0, 1.0)  // x     y
	drawReptile(1, .25, .75, .25, 1.0, 1.0)   // x-d   y+d
	drawReptile(1, .75, .75, .25, -1.0, 1.0)  // x+d   y+d
	drawReptile(1, .75, .25, .25, -1.0, -1.0) // x+d   y-d

	// C filled in,  (-.5,.5) 0.5 <1, 1>
	drawReptile(1, -.5, .5, .25, 1.0, 1.0)    // x      y
	drawReptile(1, -.25, .75, .25, -1.0, 1.0) // x+d    y+d
	drawReptile(1, -.75, .75, .25, 1.0, 1.0)  // x-d    y+d
	drawReptile(1, -.75, .25, .25, 1.0, -1.0) // x-d    y-d

	// D filled in  (-.5,-.5) 0.5 <1, -1>
	drawReptile(1, -.5, -.5, .25, 1.0, -1.0)    // x    y
	drawReptile(1, -.25, -.75, .25, -1.0, -1.0) // x+d  y-d
	drawReptile(1, -.75, -.75, .25, 1.0, -1.0)  // x-d  y-d
	drawReptile(1, -.75, -.25, .25, 1.0, 1.0)   // x-d  y+d

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

	xincr := scale * xFlip
	yincr := scale * yFlip

	fmt.Println("newpath\ngsave")
	fmt.Printf("%f %f moveto\n", x, y)
	fmt.Printf("%f %f lineto\n", x+xincr, y)
	fmt.Printf("%f %f lineto\n", x+xincr, y+yincr)
	fmt.Printf("%f %f lineto\n", x-xincr, y+yincr)
	fmt.Printf("%f %f lineto\n", x-xincr, y-yincr)
	fmt.Printf("%f %f lineto\n", x, y-yincr)
	fmt.Printf("%f %f lineto\n", x, y)

	fmt.Println("stroke\ngrestore")
}
