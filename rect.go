package main

import "fmt"

type rect struct {
	l, b int
}

func (r *rect) area() int {
	return r.l * r.b
}

func (r rect) incL() {
	r.l++
}

func (r *rect) incB() {
	r.b++
}

func main() {
	r := rect{l: 10, b: 100}

	fmt.Println(r)
	fmt.Println(r.area())
	r.incL()
	fmt.Println(r)
	r.incB()
	fmt.Println(r)
}
