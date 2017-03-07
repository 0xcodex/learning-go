package main

import "fmt"
import "time"

func main() {
	const (
		x uint16 = 120
		y
		s string = "abc"
		z
	)

	fmt.Printf("%T, %v\n", y,y)
	fmt.Printf("%T, %v\n", z,z)

	time.Sleep(time.Second)
}