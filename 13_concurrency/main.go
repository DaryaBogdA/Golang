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

func asdf() { //Небуферизированные каналы
	ch := make(chan string)

	go func() {
		ch <- "привет" // отправитель блокируется, пока main не прочитает
	}()

	msg := <-ch // получатель блокируется, пока горутина не отправит
	fmt.Println(msg)

	// Буферизированные каналы
	ch1 := make(chan int, 3)
	ch1 <- 1 // не блокируется
	ch1 <- 2 // не блокируется
	ch1 <- 3 // не блокируется
	// ch1 <- 4 // заблокируется, так как буфер заполнен

	// Закрытие канала
	ch3 := make(chan int, 2)
	ch3 <- 1
	ch3 <- 2
	close(ch3)

	v, ok := <-ch3 // v = 1, ok = true
	v, ok = <-ch3  // v = 2, ok = true
	v, ok = <-ch3  // v = 0, ok = false (канал закрыт и пуст)

	// чтение канала через For range
	ch4 := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch4 <- i
		}
		close(ch4)
	}()
	for v := range ch4 {
		fmt.Println(v)
	}
	// после закрытия и чтения всех значений цикл завершится

	//select помогает ожидать операции оть разных горутин
	select {
	case msg1 := <-ch1:
		fmt.Println("получено из ch1:", msg1)
	case msg2 := <-ch4:
		fmt.Println("получено из ch2:", msg2)
	case ch3 <- 42:
		fmt.Println("отправлено в ch3")
	default:
		fmt.Println("ни одна операция не готова")
	}
}
