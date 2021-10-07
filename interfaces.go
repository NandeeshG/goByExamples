package main

import "fmt"

type dog struct {
	legs int
}
type hen struct {
	legs int
}
type animal interface {
	speak()
}

func (d *dog) speak() *dog {
	fmt.Println("Woof!")
	return d
}
func (h *hen) speak() *hen {
	fmt.Println("Kukdookoo!")
	return h
}

func editArr(arr []int) {
	arr[0] = 1
	fmt.Printf("arr: %v, type: %T\n", arr, arr)
	return
}

func main() {
	//var d animal = &dog{}
	//var h animal = &hen{}
	d := &dog{}
	h := &hen{}
	d.speak().speak()
	h.speak().speak()

	arr := make([]int, 4)
	editArr(arr)
	fmt.Printf("arr: %v\n", arr)
}
