package main

import (
	"errors"
	"fmt"
)

func divide(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("divide by zero")
	} else {
		return x / y, nil
	}
}
func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Перехвачена паника:", r)
		}
	}()

	fmt.Println("до паники")
	panic("авария!")
	fmt.Println("после паники")

	panic("something bad happened")
	//далшье не выполнится
}
