package main

import "fmt"

func main() {
	queue := make(chan int, 4)
	for i := 0; i < 4; i++ {
		queue <- i
	}
	close(queue)

	for q := range queue {
		fmt.Println(q)
	}

	//for {
	//	q, more := <-queue
	//	if more {
	//		fmt.Println(q)
	//	} else {
	//		fmt.Println("Done")
	//		return
	//	}
	//}
}
