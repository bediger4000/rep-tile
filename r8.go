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

	rand.Seed(time.Now().UnixNano() + int64(os.Getpid()))

	fmt.Println(`/Courier 20 selectfont        
150 450 translate
75 75 scale
.01 setlinewidth`)
	drawReptile(plies, M, N)

	fmt.Println("gsave\n1 -1 translate")
	drawReptile(plies, M, N)
	fmt.Println("grestore")

	fmt.Println("gsave\n2 -2 translate")
	drawReptile(plies, M, N)
	fmt.Println("grestore")

	fmt.Println("gsave\n2 1 translate")
	drawReptile(plies, M, N)
	fmt.Println("grestore")

	fmt.Println("gsave\n3 0 translate")
	drawReptile(plies, M, N)
	fmt.Println("grestore")

	fmt.Println("gsave\n4 -1 translate")
	drawReptile(plies, M, N)
	fmt.Println("grestore")

	fmt.Println("showpage")
}

func drawReptile(ply, M, N int) {

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
		if rand.Intn(N) < M {
			fmt.Println("gsave")
			fmt.Println(".0  .0 translate")
			fmt.Println(".5  .5 scale")
			drawReptile(ply, M, N)
			fmt.Println("grestore")
		}

		if rand.Intn(N) < M {
			fmt.Println("gsave")
			fmt.Println("-.5 .5 translate")
			fmt.Println(".5  .5 scale")
			drawReptile(ply, M, N)
			fmt.Println("grestore")
		}

		if rand.Intn(N) < M {
			fmt.Println("gsave")
			fmt.Println(".5 .5 translate")
			fmt.Println("-.5  .5 scale")
			drawReptile(ply, M, N)
			fmt.Println("grestore")
		}

		if rand.Intn(N) < M {
			fmt.Println("gsave")
			fmt.Println("-.5 -.5 translate")
			fmt.Println(" .5 -.5 scale")
			drawReptile(ply, M, N)
			fmt.Println("grestore")
		}
	}
}
