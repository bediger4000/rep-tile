package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	plies, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	M, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	N, err := strconv.Atoi(os.Args[3])
	if err != nil {
		log.Fatal(err)
	}
	P, err := strconv.Atoi(os.Args[4])
	if err != nil {
		log.Fatal(err)
	}

	rand.Seed(time.Now().UnixNano() + int64(os.Getpid()))

	fmt.Println(`150 450 translate
75 75 scale
.001 setlinewidth`)

	var centers = [][2]int{
		[2]int{0, 0},
		[2]int{1, -1},
		[2]int{2, -2},
		[2]int{2, 1},
		[2]int{3, 0},
		[2]int{3, -3},
		[2]int{5, -2},
	}

	for _, center := range centers {

		fmt.Printf("gsave\n%d %d translate\n", center[0], center[1])
		drawReptile(plies, M, N, P)
		fmt.Println("grestore")
	}

	fmt.Println("showpage")
}

func drawReptile(ply, M, N, P int) {

	if rand.Intn(N) < M {
		fmt.Println("newpath")
		fmt.Println(" 0.  0. moveto")
		fmt.Println(" 1.  0. lineto")
		fmt.Println(" 1.  1. lineto")
		fmt.Println("-1.  1. lineto")
		fmt.Println("-1. -1. lineto")
		fmt.Println(" 0. -1. lineto")
		fmt.Println(" 0.  0. lineto")
		fmt.Println("closepath")
		if rand.Intn(20) <= 1 {
			switch rand.Intn(5) {
			case 0:
				fmt.Println("0.5 setgray fill")
			case 1:
				fmt.Println("1 0 0 setrgbcolor fill")
			case 2:
				fmt.Println("0 1 0 setrgbcolor fill")
			case 3:
				fmt.Println("0 0 1 setrgbcolor fill")
			}
		}
		fmt.Println("stroke")
	}

	subtiles := [][4]float64{
		[4]float64{0.0, 0.0, 0.5, 0.5},
		[4]float64{-.5, 0.5, 0.5, 0.5},
		[4]float64{0.5, 0.5, -.5, 0.5},
		[4]float64{-.5, -.5, 0.5, -.5},
	}

	ply--
	if ply > 0 {
		for _, f := range subtiles {
			if rand.Intn(N) < P {
				fmt.Printf("gsave\n%.01f %.01f translate\n%.01f %.01f scale\n", f[0], f[1], f[2], f[3])
				drawReptile(ply, M, N, P)
				fmt.Println("grestore")
			}
		}
	}
}
