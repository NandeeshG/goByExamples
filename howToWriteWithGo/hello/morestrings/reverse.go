package morestrings

import "fmt"

func ReverseRunes(s string) string {
	r := []rune(s)
	fmt.Println("rune: ", r)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
