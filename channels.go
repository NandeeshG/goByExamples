package main

import (
	"fmt"
	"time"
)

func main() {
	chn := make(chan int)
	go rou(chn)
	go rou(chn)
	time.Sleep(5 * time.Second)
	fmt.Printf("chn: %v\n", <-chn)
}

func rou(chn chan int) {
	time.Sleep(1 * time.Second)
	chn <- <-chn + 1
}
