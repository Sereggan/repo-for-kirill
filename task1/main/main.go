package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.
	y := 0.

	for {
		fmt.Println(z, y)
		z, y = z-(z*z-x)/(2*z), z

		if math.Abs(y-z) < 1e-8 {
			return z
		}
	}

}

func main() {
	x := 4.
	fmt.Println(Sqrt(x))
	fmt.Println(math.Sqrt(x) == Sqrt(x))
}
