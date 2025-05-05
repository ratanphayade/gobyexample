package main

import (
	"fmt"
)

func main() {
	a := 10
	square(a)
	fmt.Println(a)
}

func square(x int) int {
	return x * x
}
