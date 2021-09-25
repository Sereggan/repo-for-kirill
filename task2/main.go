package main

import (
	"fmt"
	"math"
	"strings"
)

type IPAddr [4]byte
type ErrNegativeSqrt float64
type MyReader struct{}

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Number should be positive: %v", float64(e))
}

func main() {
	task1()
	task2()
	task3()
}

// https://tour.golang.org/moretypes/23
func task1() map[string]int {
	kirill := "Hello Kirill156rus hehe!"

	fmt.Println("\nTask 1:")
	wordsMap := make(map[string]int)

	for _, v := range strings.Fields(kirill) {
		c := string(v)
		_, ok := wordsMap[c]
		if ok {
			wordsMap[c] += 1
		} else {
			wordsMap[c] = 1
		}

	}
	return wordsMap
}

// https://tour.golang.org/methods/18
func task2() {
	fmt.Println("\nTask 2:")

	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

func (ip IPAddr) String() string {
	ipAddr := []string{}
	for _, v := range ip {
		ipAddr = append(ipAddr, fmt.Sprint(int(v)))
	}
	return strings.Join(ipAddr, ".")
}

// https://tour.golang.org/methods/20
func task3() {
	fmt.Println("\nTask 3:")

	x := 4.
	result, _ := Sqrt(x)
	fmt.Println(result)
	fmt.Println(math.Sqrt(x) == result)
	x = -4
	result, err := Sqrt(x)
	fmt.Println(err)
	if err != nil {
		fmt.Println(true)
	}
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.
	y := 0.

	for {
		z, y = z-(z*z-x)/(2*z), z
		if math.Abs(y-z) < 1e-8 {
			return z, nil
		}
	}

}
