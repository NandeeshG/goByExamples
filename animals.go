package main

import "fmt"

type Dog struct{}
type Cat struct{}
type Mouse struct{}

func (t *Dog) speak() string {
	return "Woof!"
}
func (t *Cat) speak() string {
	return "Meow!"
}
func (t *Mouse) speak() string {
	return "Squeak!"
}

type Animal interface {
	speak() string
}

func main() {
	A := []Animal{&Cat{}, &Dog{}, &Mouse{}}
	for _, a := range A {
		fmt.Println(a.speak())
	}
}
