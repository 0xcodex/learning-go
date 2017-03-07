package main

func main() {
	for i := 0; i < 5; i++ {
		println(i)
	}

	for i := 4; i >= 0; i-- {
		println(i)
	}

	x := 0
	for x < 5 {
		println(x)
		x++
	}

	y := 4

	for{
		println(y)
		y--

		if y < 0 {
			break
		}
	}
}