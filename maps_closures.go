package main

import "fmt"

func main() {
	//nodes := 5
	edges := 6
	edges_ip := [...][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}, {2, 4}, {1, 3}}

	graph := make(map[int][]int)
	fmt.Printf("graph: %v\n", graph)

	for i := 0; i < edges; i++ {
		a, b := edges_ip[i][0], edges_ip[i][1]
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	fmt.Printf("graph: %v\n", graph)

	fmt.Println(sum(1, 2, 3))

	//arr := [...]int{1, 2, 3}
	//fmt.Println(sum(arr...))

	sl := []int{1, 2, 3}
	fmt.Println(sum(sl...))

	counter1 := counter(11)
	fmt.Println(counter1())
	fmt.Println(counter1())
	fmt.Println(counter1())

	counter2 := counter(27)
	fmt.Println(counter2())
	fmt.Println(counter2())
	fmt.Println(counter2())

	person := Person{name: "Nan", age: 10}
	fmt.Println(person)

	changePerson(person)
	fmt.Println(person)

	changePersonPtr(&person)
	fmt.Println(person)

	person2 := PersonConstr("yo", 20)
	fmt.Printf("person2: %v\n", person2)

	person3 := PersonConstrPtr("yo3", 30)
	fmt.Printf("person3: %v\n", person3.age)
}

type Person struct {
	name string
	age  int
}

func PersonConstr(name string, age int) Person {
	newPerson := Person{name, age}
	return newPerson
}

func PersonConstrPtr(name string, age int) *Person {
	newPerson := Person{name, age}
	return &newPerson
}

func changePerson(person Person) {
	person.age++
}

func changePersonPtr(person *Person) {
	(*person).age++
}

func counter(inc int) func() int {
	ctr := 0
	return func() int {
		ctr += inc
		return ctr
	}
}

func sum(vals ...int) int {
	s := 0
	for _, v := range vals {
		s += v
	}
	return s
}
