package main

import (
	"example/nandeeshG/hello/morestrings"
	"fmt"

	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println("Hi")
	fmt.Println(morestrings.ReverseRunes("Hi"))
	fmt.Println("Hi")
	fmt.Println(morestrings.ReverseRunes("Hi"))
	fmt.Println(cmp.Diff("ABC", "CBA"))
}
