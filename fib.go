package main

import "fmt"

func main() {
	x := 0
	y := 1
	z := 0
	counter := 0

	for counter < 50 {
		fmt.Println(x)
		z = x + y
		x = y
		y = z
		counter++
	}
}
