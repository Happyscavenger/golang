package main

import (
	"fmt"
	"sync"
)

type rect struct {
	length int
	width  int
}

func (r rect) area(wg *sync.WaitGroup) {
	defer wg.Done()
	if r.length <= 0 {
		fmt.Println("length of r must greater than zero")
		return
	}
	if r.width <= 0 {
		fmt.Println("width of r must greater than zero")
		return
	}
	area := r.length * r.width
	fmt.Printf("area of r is %d\n", area)
}

func main() {
	var wg sync.WaitGroup
	r1 := rect{10, -8}
	r2 := rect{-10, 20}
	r3 := rect{10, 2}
	rects := []rect{r1, r2, r3}
	for _, v := range rects {
		wg.Add(1)
		go v.area(&wg)
	}
	wg.Wait()
}
