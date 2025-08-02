package main

import (
	"fmt"
)

func main() {
	exampleDefer()
}

func exampleDefer() {
	foo := "hello"
	// defer хоть и исполняется вконце функции,
	// но запоминает значение foo на тот момент где объявлен
	defer fmt.Printf("Foo from deffer: %s\n", foo)

	foo = "bay"

	fmt.Printf("Foo from end function: %s\n", foo)
}
