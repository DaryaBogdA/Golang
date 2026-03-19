package main

import "fmt"

func main() {
	map1 := make(map[string]int)

	map1["a"] = 1

	x := map1["fds"] // 0

	delete(map1, "a")

	a, found := map1["a"]
	if found {
		fmt.Println(a)
	}

	for key, value := range map1 {
		fmt.Println(key, value)
	}
}
