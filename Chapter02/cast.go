package main

import "fmt"

func main() {
	x := 100

	p := (*int)(&x)

	fmt.Println(p)
}