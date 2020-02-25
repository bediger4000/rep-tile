package main

import "fmt"

const (
	alpha int = 0
	beta  int = 1
	gamma int = 2
)

func main() {
	repTile(1, 0.0, 0.0, 1.0)
}

func repTile(ply int, x, y, side float64) {
	if ply == 0 {
		return
	}
	printTile(x, y, side, alpha)
	printTile(x-side/2.0, y+side/2.0, side, alpha)
	printTile(x+side/2.0, y+side/2.0, side, beta)
	printTile(x-side/2.0, y-side/2.0, side, gamma)
}

func printTile(x, y, side float64, which int) {
	for _, increment := range increments[which] {
		fmt.Printf("%f\t%f\n", x+side*increment[0], y+side*increment[1])
	}
	fmt.Printf("%f\t%f\n", x+side*increments[0][0], y+side*increments[0][1])
}

var incrementsAlpha [6][2]float64 = [6][2]float64{
	[2]float64{0.0, 0.0},
	[2]float64{1.0, 0.0},
	[2]float64{1.0, 1.0},
	[2]float64{-1.0, 1.0},
	[2]float64{-1.0, -1.0},
	[2]float64{0.0, -1.0},
}
var incrementsBeta [6][2]float64 = [6][2]float64{
	[2]float64{0.0, 0.0},
	[2]float64{0.0, -1.0},
	[2]float64{1.0, -1.0},
	[2]float64{1.0, 1.0},
	[2]float64{-1.0, 1.0},
	[2]float64{-1.0, 0.0},
}
var incrementsGamma [6][2]float64 = [6][2]float64{
	[2]float64{0.0, 0.0},
	[2]float64{1.0, 0.0},
	[2]float64{1.0, -1.0},
	[2]float64{-1.0, -1.0},
	[2]float64{-1.0, 1.0},
	[2]float64{0.0, 1.0},
}

var increments [3][6][2]float64 = [3][6][2]float64{incrementsAlpha, incrementsBeta, incrementsGamma}
