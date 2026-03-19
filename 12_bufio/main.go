package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// bufio.NewReader(os.Stdin) введите что-то
// bufio.NewWriter(file) записать в файл
// bufio.NewScanner(file) считать с файла
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите ваше имя: ")

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	name := strings.TrimSpace(input)
	fmt.Printf("Привет, %s!\n", name)

	file, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	//записывает в буфер
	_, err = writer.WriteString("Это первая строка.\n") // _ кол-во записанных байт
	if err != nil {
		log.Fatal(err)
	}
	_, err = writer.WriteString("А это вторая строка.\n")
	if err != nil {
		log.Fatal(err)
	}

	// Сбрасываем буфер в файл!
	err = writer.Flush()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Данные успешно записаны в output.txt")

	file, err := os.Open("output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// разделения на слова
	scanner.Split(bufio.ScanWords)

	wordCount := 0
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		wordCount++
	}

	// Всегда проверяем на ошибки после цикла!
	if err := scanner.Err(); err != nil {
		log.Fatal("Ошибка при сканировании:", err)
	}

	fmt.Printf("Найдено слов: %d\n", wordCount)
}
