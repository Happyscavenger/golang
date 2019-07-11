package main

import "fmt"

//封装的python range
func Range(n int) []int {
	var rng []int
	for i := 0; i < n; i++ {
		rng = append(rng, i)
	}
	return rng
}

func main() {
	var a = 10
	b := Range(a)
	fmt.Println(b)
	for _, i := range b {
		fmt.Println(i)
	}
}
