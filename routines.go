/*
package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	f("direct")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	g := [8][2]int{
		{},
		{2, 3},
		{4, 5},
		{6, 7},
	}
	fmt.Printf("g: %v\n", g)
	fmt.Println(time.Second)

	dfs(1, 0, g)
	//time.Sleep(10 * time.Second)
}

func dfs(nd, par int, g [8][2]int) {
	time.Sleep(10 * time.Second)
	fmt.Println("Entered: ", nd)
	for _, v := range g[nd] {
		if v > 0 && v != par {
			//fmt.Println("Spawned for: ", nd)
			go dfs(v, nd, g)
		}
	}
	fmt.Println("Exited: ", nd)
}
