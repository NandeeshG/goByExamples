package main

import "fmt"

func main() {
	//message := make(chan string, 1) //makes unbuffered chan. Requires a receiver before passing to channel
	message := make(chan string, 1) //makes buffered channel

	select {
	case msg := <-message:
		fmt.Println("10 ", msg)
	default:
		fmt.Println("12 ")
	}

	select {
	case message <- "16":
		fmt.Println("17")
	default:
		fmt.Println("19")
	}

	select {
	case msg := <-message:
		fmt.Println("24 ", msg)
	default:
		fmt.Println("26 ")
	}
}
