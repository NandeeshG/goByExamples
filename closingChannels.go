package main

import "fmt"

func main() {
	jobs := make(chan int)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			fmt.Println(j, more)
			if more {
				fmt.Println("Rec ", j)
			} else {
				fmt.Println("All done")
				done <- true
				return
			}
		}
	}()

	for i := 0; i < 3; i++ {
		fmt.Println("sending ", i)
		jobs <- i
	}
	fmt.Println("SENT all")
	close(jobs)

	<-done
}
