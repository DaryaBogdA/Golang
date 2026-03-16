package main

import "fmt"

func main() {
	age := 4
	if age > 10 {
		fmt.Println("Вам больше 10")
	} else if age == 4 {
		fmt.Println("Вам 4")
	} else {
		fmt.Println("Хз сколько вам")
	}

	switch age {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("---")
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for age < 10 {
		fmt.Println(age)
		age++
	}
}
