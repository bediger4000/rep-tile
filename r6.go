package main

import "fmt"

func main() {
	fmt.Println(`/Courier 20 selectfont        
150 450 translate
100 100 scale
.01 setlinewidth`)
	drawReptile(4)

	fmt.Println("showpage")
}

func drawReptile(ply int) {

	fmt.Println("newpath")
	fmt.Println(" 0.  0. moveto")
	fmt.Println(" 1.  0. lineto")
	fmt.Println(" 1.  1. lineto")
	fmt.Println("-1.  1. lineto")
	fmt.Println("-1. -1. lineto")
	fmt.Println(" 0. -1. lineto")
	fmt.Println(" 0.  0. lineto")
	fmt.Println("stroke")

	ply--
	if ply > 0 {
		fmt.Println("gsave")
		fmt.Println(".0  .0 translate")
		fmt.Println(".5  .5 scale")
		drawReptile(ply)
		fmt.Println("grestore")

		fmt.Println("gsave")
		fmt.Println("-.5 .5 translate")
		fmt.Println(".5  .5 scale")
		drawReptile(ply)
		fmt.Println("grestore")

		fmt.Println("gsave")
		fmt.Println(".5 .5 translate")
		fmt.Println("-.5  .5 scale")
		drawReptile(ply)
		fmt.Println("grestore")

		fmt.Println("gsave")
		fmt.Println("-.5 -.5 translate")
		fmt.Println(" .5 -.5 scale")
		drawReptile(ply)
		fmt.Println("grestore")
	}

	// fmt.Println("stroke\ngrestore")
}
