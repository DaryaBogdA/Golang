package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // сигнал о завершении
	fmt.Printf("Worker %d started\n", id)
	// имитация работы
	fmt.Printf("Worker %d finished\n", id)
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait() // ждем конец горутин
	fmt.Println("All workers done")

	arr := [5]int{1, 2, 3, 4, 5}
	s := arr[1:3:4]             // элементы [2,3], len=2, cap=4-1=3
	fmt.Println(len(s), cap(s)) // 2 3

	s = append(s, 100) // [2,3,100]
	s = append(s, 435) // не добавит так как cap ьаксимаоьный
	fmt.Println(arr)   // [1,2,3,100,5]
}
