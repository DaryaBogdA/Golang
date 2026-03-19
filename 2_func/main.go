package main

import (
	"errors"
	"fmt"
)

func main() {
	//Анонимная функция
	inc := func(x int) int { return x + 1 }
	fmt.Println(inc(5)) //6

	var a MyInt = 10
	fmt.Println(a.minus()) // -10
}

func square(x int) int {
	sqr := x * x
	return sqr
}

type MyInt int

func (x MyInt) minus() MyInt {
	return -x
}

func summ(x int, y int, z int) int {
	return x + y + z
}

func printAll(nums ...int) {
	fmt.Println(nums)
}

func fd() {
	x := 10
	defer func() {
		fmt.Println(x) //  20
	}()
	x = 20
}

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func stats(nums []int) (min, max int) {
	min, max = nums[0], nums[0]
	for _, v := range nums {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func zamekanie() {
	a := adder()
	fmt.Println(a(5)) // 5
	fmt.Println(a(3)) // 8
}
