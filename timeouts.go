package main

import (
	"fmt"
	"time"
)

func main() {
	str := make(chan string)

	r1 := func() {
		time.Sleep(3 * time.Second)
		str <- "r1 ends"
	}

	go r1()

	for i := 0; i < 2; i++ {
		select {
		case mess := <-str:
			fmt.Println(mess)
		case <-time.After(4 * time.Second):
			fmt.Println("Timeout 1")
		}
	}
}
