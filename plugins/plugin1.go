package main

import "fmt"

func Run() {
	fmt.Printf("(plugin1) Running a function in a Go plugin!\n", )
}

var I = 7

func PrintI() {
	fmt.Printf("(plugin1) Printing new value: %d\n", I)
}
