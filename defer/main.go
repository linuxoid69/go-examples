package main

import (
	"fmt"
)

func main() {
	printDefer()
	deferOrder()
}

func printDefer() {
	foo := "hello"
	// defer хоть и исполняется вконце функции,
	// но запоминает значение foo на тот момент где объявлен
	defer fmt.Printf("Foo from deffer: %s\n", foo)
	// defer func () {fmt.Printf("Foo from deffer: %s\n", foo)}() // имеет больший приоритет чем print

	foo = "bay"

	fmt.Printf("Foo from end function: %s\n", foo)
}

func deferOrder() {
	// Порядок вызова меняется
	defer fmt.Println("1")
	defer fmt.Println("2")
}
