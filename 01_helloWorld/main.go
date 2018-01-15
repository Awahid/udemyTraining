package main

import (
	"fmt"
)

const p string = "constant"

const (
	PI = 3.14
)

func main() {
	a := 10
	b := "golang"
	c := 4.17
	d := true

	fmt.Printf("%v %T \n", a, a)
	fmt.Printf("%v %T \n", b, b)
	fmt.Printf("%v %T\n", c, c)
	fmt.Printf("%v %T \n", d, d)
	increament := wrapper()
	fmt.Println(increament())
	fmt.Println(p)
	fmt.Println(PI)
	withParam := withParameters(1, 2)
	fmt.Println(withParam)
	n := average(43, 56, 87, 12, 45, 57)
	fmt.Println(n)
	fmt.Println(method1(2))
	fmt.Println(max(4,7,9,123,543,200,43,125))
}

func wrapper() func() int {
	return func() int {
		return 2
	}
}

func withParameters(a, b int) int {
	return a + b
}

func average(sf ...float64) float64 {
	fmt.Println(sf)
	fmt.Printf("%T \n", sf)
	total := 0.0
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}

func method1(value int) (float64, bool) {
	return float64(value),value%2 == 0
}

func max(numbers ...int) int{
	var largest int
	for _, v:= range numbers  {
		if v > largest{
			largest = v
		}
	}
	return largest
}
