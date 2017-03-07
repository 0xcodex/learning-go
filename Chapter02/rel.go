package main

import "fmt"

func mkslice() []int {
	s := make([]int, 0, 10)
	s = append(s,100)
	return s
}

func mkmap() map[string]int {
	m := make(map[string]int)
	m["a"] = 1
	return m
}

func main() {
	m := mkmap()
	fmt.Println(m)

	s := mkslice()
	fmt.Println(s)
}

// err
// func main() {
// 	p := new(map[string]int)
// 	m := *p
// 	m["a"] = 1
// 	fmt.Println(m)
// }