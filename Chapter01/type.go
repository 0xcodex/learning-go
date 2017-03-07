package main

import(
	"fmt"
)

type X int

func (x *X) inc() {
	*x ++
}

type User struct{
	name string
	age int
}

func (user *User) p(){
	fmt.Println(*user)
}

func main() {
	var x X
	x.inc()

	println(x)

	var user User
	user.age = 30
	user.name = "Tom"

	user.p()

}