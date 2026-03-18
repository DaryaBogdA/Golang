package main

import "fmt"

func main() {
	x := 4
	p := &x
	fmt.Println(*p)
	*p = 9
	fmt.Println(x)

	a := 10
	byValue(a)
	fmt.Println(a) // 10

	byPointer(&a)
	fmt.Println(a) // 100
}
func byValue(x int) {
	x = 100
}

func byPointer(x *int) {
	*x = 100
}
