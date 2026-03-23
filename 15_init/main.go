package main

import "fmt"

func init() {
	fmt.Println("init1") // нужно для поключения к бд например
}

func init() {
	fmt.Println("init2")
}

func main() {
	fmt.Println("main")
}

// init1
// init2
// main
