package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	person := Person{"Jora", 13}

	//person2 := Person{
	//	name: "Jora",
	//	age:  13,
	//}
	person.age = 9
	//word := person2.name

	p := Person{name: "Bob", age: 25}
	p.birthdayValue()
	fmt.Println(p.age) // 25

	p.birthdayPointer()
	fmt.Println(p.age) // 26

	aas()
}

func (p Person) hello() {
	fmt.Println(p.name)
}

func (p Person) birthdayValue() {
	p.age++
}

func (p *Person) birthdayPointer() {
	p.age++
}

type Account struct {
	name string
	age  int
}

type ManageAccount struct {
	Account
}

func (a Account) String() string {
	return "Standart"
}

func (m ManageAccount) String() string {
	return "Manage Account"
}

func aas() {
	myAcc := ManageAccount{Account{"dsgdg", 34}}
	myAcc2 := Account{"dsgdg", 34}
	fmt.Println(myAcc)
	fmt.Println(myAcc2)
}
