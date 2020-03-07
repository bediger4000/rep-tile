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
100 100 scale
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
		fmt.Println("stroke")
	}

	ply--
	if ply > 0 {
		if rand.Intn(N) < P {
			fmt.Println("gsave")
			fmt.Println(".0  .0 translate")
			fmt.Println(".5  .5 scale")
			drawReptile(ply, M, N, P)
			fmt.Println("grestore")
		}

		if rand.Intn(N) < P {
			fmt.Println("gsave")
			fmt.Println("-.5 .5 translate")
			fmt.Println(".5  .5 scale")
			drawReptile(ply, M, N, P)
			fmt.Println("grestore")
		}

		if rand.Intn(N) < P {
			fmt.Println("gsave")
			fmt.Println(".5 .5 translate")
			fmt.Println("-.5  .5 scale")
			drawReptile(ply, M, N, P)
			fmt.Println("grestore")
		}

		if rand.Intn(N) < P {
			fmt.Println("gsave")
			fmt.Println("-.5 -.5 translate")
			fmt.Println(" .5 -.5 scale")
			drawReptile(ply, M, N, P)
			fmt.Println("grestore")
		}
	}
}
