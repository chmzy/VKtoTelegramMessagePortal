package main

import "fmt"

func pick[T any](cond bool, a, b T) T {
	if cond {
		return a
	} else {
		return b
	}
}

func main() {
	x := 10
	y := 20
	fmt.Println(pick((x > y), x, y))
}
