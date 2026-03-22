package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Stats struct {
	Sum       int
	Count     int
	Average   float64
	EvenCount int
	OddCount  int
}

func main() {
	numbers, err := readNumbers("numbers.txt")
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	if len(numbers) == 0 {
		log.Fatal("Файл не содержит чисел")
	}

	show(numbers)

	sumCh := make(chan int)
	evenCh := make(chan int)
	oddCh := make(chan int)

	go func() {
		odd, even := OddEven(numbers)
		oddCh <- odd
		evenCh <- even
	}()

	go func() {
		sumCh <- summa(numbers)
	}()

	sum := <-sumCh
	odd := <-oddCh
	even := <-evenCh

	avg := float64(sum) / float64(len(numbers))

	stats := Stats{
		Sum:       summa(numbers),
		Count:     len(numbers),
		Average:   avg,
		EvenCount: even,
		OddCount:  odd,
	}
	if err := saveStats("stats.json", stats); err != nil {
		log.Fatal(err)
	}
}

func readNumbers(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("Error opening file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var numbers []int

	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, number)
	}

	return numbers, nil
}

func show(arr []int) {
	for _, val := range arr {
		fmt.Println(val)
	}
}

func summa(arr []int) (sum int) {
	for _, v := range arr {
		sum += v
	}
	return sum
}

func OddEven(arr []int) (oddCount, evenCount int) {
	for _, v := range arr {
		if v%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
	}
	return oddCount, evenCount
}

func saveStats(filename string, stats Stats) error {
	data, err := json.MarshalIndent(stats, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
