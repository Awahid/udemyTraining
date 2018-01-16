package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	fmt.Println(max(4, 7, 9, 123, 543, 200, 43, 125))
	myMap()
	stringToRune()
	num := hashBucket("Go", 12)
	fmt.Println(num)
	myScanner()
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
	return float64(value), value%2 == 0
}

func max(numbers ...int) int {
	var largest int
	for _, v := range numbers {
		if v > largest {
			largest = v
		}
	}
	return largest
}

func myMap() {
	myGreetings := map[int]string{
		0: "Good morning",
		1: "Hello",
		2: "Bye",
	}

	myGreetings[3] = "Hi"

	for k, v := range myGreetings {
		fmt.Println(k, " - ", v)
	}
}

func stringToRune() {
	var letter = rune("A"[0])
	fmt.Println(letter)
}

func hashBucket(words string, buckets int) int {
	letter := int(words[0])
	bucket := letter % buckets
	return bucket
}

func myScanner() {
	const input = "The optional second return value when getting a value from a map indicates if the key was present in the map. This can be used to disambiguate between missing keys and keys with zero values like 0 or." +
		" Here we didn’t need the value itself, so we ignored it with the blank identifier"
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
}
