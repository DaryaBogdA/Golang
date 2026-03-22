package main

import (
	"errors"
	"fmt"
	"os"
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

func readConfig(path string) error {
	data, err := os.ReadFile(path)

	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("файл не найден")
	}

	if err != nil {
		return fmt.Errorf("read config %s: %w", path, err)
	}

	fmt.Println(string(data))
	return nil
}
