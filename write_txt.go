package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("lines.txt")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	d := []string{"hello world", "hello GO", "hello china"}
	for _, v := range d {
		fmt.Fprintln(f, v)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file written successfully")
}
