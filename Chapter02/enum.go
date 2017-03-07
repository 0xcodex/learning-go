package main

type color byte

const (
	black color = iota
	red
	blue
)

func test(c color) {
	println(c)
}

func main() {
	const (
		x = iota
		y
		z
	)

	println(x,y,z)

	println("======================")
	

	test(red)
	test(100)

}