package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	str := "DS;GLDKSHG"
	fmt.Println(len(str))
	num := 3.436532462
	fmt.Printf("%.3f \n", num)
	fmt.Printf("%T \n", num)

	str2 := "Привет"
	fmt.Println(len(str2)) //12 кирилица x2

	strings.EqualFold("Go", "go") // true

	strings.Contains("hello world", "world")       // true
	strings.Index("hello", "l")                    // 2
	strings.ToUpper("go")                          // "GO"
	strings.ReplaceAll("oink oink", "oink", "moo") // "moo moo"
	strings.Split("a,b,c", ",")                    // ["a", "b", "c"]
	strings.Join([]string{"a", "b"}, "-")          // "a-b"

	numStr := "2"
	num1, err := strconv.Atoi(numStr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(num1)

	//numFloat, err := strconv.ParseFloat(numStr)

}
