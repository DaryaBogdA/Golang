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
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)

	}

	for age < 10 { //while
		fmt.Println(age)
		age++
	}

	for { //вечный цикл
		if age < 10 {
			break
		}
	}
}
