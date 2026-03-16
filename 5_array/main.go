package main

import "fmt"

func main() {
	arr := [3]int{1, 2, 3}

	for i, val := range arr {
		fmt.Println(i, val)
	}
}
