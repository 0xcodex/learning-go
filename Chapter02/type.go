package main


var x = 100

func main() {
	println(&x, x)

	x := "abc"

	println(&x, x)	

	println("===================")

	y := 100

	println(&y, y)

	y,z := 10, "abc"

	println(&y, y)
	println(&z, z)

	println("===================")

	xx := 100

	println(&xx, xx)

	{
		xx, xy := "xx", "xy"
		println(&xx, xx,xy)
	}
}