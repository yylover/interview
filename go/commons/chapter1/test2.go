package main

import "fmt"

func main() {
	t := []int{1, 2, 3}
	for d, w := range t {
		fmt.Println(d, w)
	}
}
