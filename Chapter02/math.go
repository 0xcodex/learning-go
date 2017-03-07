package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	a, b, c := 100, 0144, 0x64

	fmt.Println(a, b, c)
	fmt.Printf("0b%b, %#o, %#x\n",a ,a ,a)

	fmt.Println(math.MinInt8, math.MaxInt8)

	aa, _ := strconv.ParseInt("1100100", 2, 32)
	bb, _ := strconv.ParseInt("0144", 8, 32)
	cc, _ := strconv.ParseInt("64", 16,32)

	println(aa,bb,cc)
}