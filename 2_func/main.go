package main

import "fmt"

func main() {

}

func square(x int) int {
	sqr := x * x
	return sqr
}

func summ(x int, y int, z int) int {
	return x + y + z
}

func printAll(nums ...int) {
	fmt.Println(nums)
}
