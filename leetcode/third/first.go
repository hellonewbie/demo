package main

import "fmt"

func reverseLeftWords(s string, n int) string {
	return s[n:] + s[0:n]
}

func main() {
	var s string
	var n int
	fmt.Scan(&s, &n)
	fmt.Println(reverseLeftWords(s, n))
}
