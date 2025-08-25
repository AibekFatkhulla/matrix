package main

import "fmt"

func main() {
	fmt.Println(hash(29))
}

func hash(k int) int {
	k = k * 10
	k = k * 3
	k = k * 232
	k = k % 215
	return k
}
