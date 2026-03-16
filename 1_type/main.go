package main

import (
	"fmt"
)

var globalStr = 9

//Супер глабальный переременный надо с большой пуквы Name

func main() {
	fmt.Println("Hello hgfjfgjWorld")

	var age int
	age = 4
	fmt.Println(age)

	var age2 = 10
	fmt.Println(age2)

	age3 := 20
	fmt.Println(age3)

	//age4 := float64(10)
	age4 := 10.
	fmt.Println(age4)

	x, y, z := 1, 2, 3
	fmt.Println(x, y, z)

	const pi = 3.14

	//int
	//uint
	//float64
	age5 := 6
	isAge := bool(age5 > 13)
	fmt.Println(isAge)
	name := "mashf"
	fmt.Println(name)
	fmt.Println(name[3])
	fmt.Println(string(name[3]))
	//byte
	m := byte('m')
	fmt.Println(m)
	//m2 := byte('т')
	m2 := rune('т')
	fmt.Println(m2)

}
