package main

import(
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("/dev/random")

	buf := make([]byte, 1024)
	n, err := f.Read(buf)

	fmt.Println(n,err)
}