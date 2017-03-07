package main

import (
	"fmt"
)

type User struct{
	name string
	age byte
}

func (u User) Print() {
	fmt.Printf("%+v\n", u)
}

type Printer interface {
	Print()
}

func main() {
	var u User
	u.name = "Tom"
	u.age = 40

	var p Printer = u
	p.Print()
}