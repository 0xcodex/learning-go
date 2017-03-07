package main

import "fmt"

func main() {
	type (
		User struct {
			name string
			age uint8
		}

		event func(string) bool
	)

	u := User{"Tom",20}

	fmt.Println(u)

	var f event = func(s string) bool {
		println(s)
		return s != ""
	}

	f("abc")

	
}