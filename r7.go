package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	plies, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(`/Courier 20 selectfont        
150 450 translate
100 100 scale
.01 setlinewidth`)
	drawReptile(plies)

	fmt.Println("gsave\n1 -1 translate")
	drawReptile(plies + 1)
	fmt.Println("grestore")

	fmt.Println("gsave\n2 1 translate")
	drawReptile(plies - 1)
	fmt.Println("grestore")

	fmt.Println("gsave\n3 0 translate")
	drawReptile(plies)
	fmt.Println("grestore")

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
